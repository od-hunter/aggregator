package tasks

import (
	"context"
	"fmt"
	"math"
	"strings"
	"sync"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/go-co-op/gocron"
	"github.com/google/uuid"
	fastshot "github.com/opus-domini/fast-shot"
	"github.com/paycrest/aggregator/config"
	"github.com/paycrest/aggregator/ent"
	"github.com/paycrest/aggregator/ent/fiatcurrency"
	"github.com/paycrest/aggregator/ent/lockorderfulfillment"
	"github.com/paycrest/aggregator/ent/lockpaymentorder"
	networkent "github.com/paycrest/aggregator/ent/network"
	"github.com/paycrest/aggregator/ent/paymentorder"
	"github.com/paycrest/aggregator/ent/paymentorderrecipient"
	"github.com/paycrest/aggregator/ent/providerordertoken"
	"github.com/paycrest/aggregator/ent/providerprofile"
	"github.com/paycrest/aggregator/ent/receiveaddress"
	"github.com/paycrest/aggregator/ent/senderprofile"
	tokenent "github.com/paycrest/aggregator/ent/token"
	"github.com/paycrest/aggregator/ent/transactionlog"
	"github.com/paycrest/aggregator/ent/webhookretryattempt"
	"github.com/paycrest/aggregator/services"
	orderService "github.com/paycrest/aggregator/services/order"
	"github.com/paycrest/aggregator/storage"
	"github.com/paycrest/aggregator/types"
	"github.com/paycrest/aggregator/utils"
	"github.com/paycrest/aggregator/utils/logger"
	"github.com/redis/go-redis/v9"
	"github.com/shopspring/decimal"
)

var orderConf = config.OrderConfig()
var serverConf = config.ServerConfig()

var rpcClients = map[string]types.RPCClient{}

// setRPCClients connects to the RPC endpoints of all networks
func setRPCClients(ctx context.Context) ([]*ent.Network, error) {
	isTestnet := false
	if serverConf.Environment != "production" {
		isTestnet = true
	}

	networks, err := storage.Client.Network.
		Query().
		Where(networkent.IsTestnetEQ(isTestnet)).
		All(ctx)
	if err != nil {
		return nil, fmt.Errorf("setRPCClients.fetchNetworks: %w", err)
	}

	// Connect to RPC endpoint
	var client types.RPCClient
	for _, network := range networks {
		if rpcClients[network.Identifier] == nil && !strings.HasPrefix(network.Identifier, "tron") {
			retryErr := utils.Retry(3, 1*time.Second, func() error {
				client, err = types.NewEthClient(network.RPCEndpoint)
				return err
			})
			if retryErr != nil {
				logger.Errorf("failed to connect to %s RPC %v", network.Identifier, retryErr)
				continue
			}

			rpcClients[network.Identifier] = client
		}
	}

	return networks, nil
}

// RetryStaleUserOperations retries stale user operations
func RetryStaleUserOperations() error {
	ctx := context.Background()
	var wg sync.WaitGroup

	// Establish RPC connections
	_, err := setRPCClients(ctx)
	if err != nil {
		return fmt.Errorf("RetryStaleUserOperations: %w", err)
	}

	// Process initiated orders
	orders, err := storage.Client.PaymentOrder.
		Query().
		Where(func(s *sql.Selector) {
			ra := sql.Table(receiveaddress.Table)
			s.LeftJoin(ra).On(s.C(paymentorder.FieldReceiveAddressText), ra.C(receiveaddress.FieldAddress)).
				Where(sql.And(
					sql.EQ(s.C(paymentorder.FieldStatus), paymentorder.StatusInitiated),
					sql.EQ(ra.C(receiveaddress.FieldStatus), receiveaddress.StatusUsed),
					sql.IsNull(s.C(paymentorder.FieldGatewayID)),
				))
		}).
		Where(
			paymentorder.Or(
				paymentorder.UpdatedAtGTE(time.Now().Add(-10*time.Minute)),
				paymentorder.HasRecipientWith(
					paymentorderrecipient.MemoHasPrefix("P#P"),
				),
			)).
		WithToken(func(tq *ent.TokenQuery) {
			tq.WithNetwork()
		}).
		All(ctx)
	if err != nil {
		return fmt.Errorf("RetryStaleUserOperations: %w", err)
	}

	wg.Add(1)
	go func() {
		defer wg.Done()
		for _, order := range orders {
			orderAmountWithFees := order.Amount.Add(order.NetworkFee).Add(order.SenderFee).Add(order.ProtocolFee)
			if order.AmountPaid.GreaterThanOrEqual(orderAmountWithFees) {
				var service types.OrderService
				if strings.HasPrefix(order.Edges.Token.Edges.Network.Identifier, "tron") {
					service = orderService.NewOrderTron()
				} else {
					service = orderService.NewOrderEVM()
				}
				err := service.CreateOrder(ctx, rpcClients[order.Edges.Token.Edges.Network.Identifier], order.ID)
				if err != nil {
					logger.Errorf("RetryStaleUserOperations.CreateOrder %v", err)
				}
			}
		}
	}()

	// Settle order process
	lockOrders, err := storage.Client.LockPaymentOrder.
		Query().
		Where(
			lockpaymentorder.StatusEQ(lockpaymentorder.StatusValidated),
			lockpaymentorder.HasFulfillmentsWith(
				lockorderfulfillment.ValidationStatusEQ(lockorderfulfillment.ValidationStatusSuccess),
			),
			lockpaymentorder.UpdatedAtLT(time.Now().Add(-5*time.Minute)),
		).
		WithToken(func(tq *ent.TokenQuery) {
			tq.WithNetwork()
		}).
		All(ctx)
	if err != nil {
		return fmt.Errorf("RetryStaleUserOperations: %w", err)
	}

	wg.Add(1)
	go func() {
		defer wg.Done()
		for _, order := range lockOrders {
			var service types.OrderService
			if strings.HasPrefix(order.Edges.Token.Edges.Network.Identifier, "tron") {
				service = orderService.NewOrderTron()
			} else {
				service = orderService.NewOrderEVM()
			}
			err := service.SettleOrder(ctx, rpcClients[order.Edges.Token.Edges.Network.Identifier], order.ID)
			if err != nil {
				logger.Errorf("RetryStaleUserOperations.SettleOrder: %v", err)
			}
		}
	}()

	// Refund order process
	lockOrders, err = storage.Client.LockPaymentOrder.
		Query().
		Where(
			lockpaymentorder.GatewayIDNEQ(""),
			lockpaymentorder.Or(
				lockpaymentorder.And(
					lockpaymentorder.Or(
						lockpaymentorder.StatusEQ(lockpaymentorder.StatusPending),
						lockpaymentorder.StatusEQ(lockpaymentorder.StatusCancelled),
					),
					lockpaymentorder.CreatedAtLTE(time.Now().Add(-10*time.Minute)),
					lockpaymentorder.Or(
						lockpaymentorder.Not(lockpaymentorder.HasFulfillments()),
						lockpaymentorder.HasFulfillmentsWith(
							lockorderfulfillment.Not(lockorderfulfillment.ValidationStatusEQ(lockorderfulfillment.ValidationStatusPending)),
						),
						lockpaymentorder.HasFulfillmentsWith(
							lockorderfulfillment.ValidationStatusEQ(lockorderfulfillment.ValidationStatusFailed),
							lockorderfulfillment.Not(lockorderfulfillment.ValidationStatusEQ(lockorderfulfillment.ValidationStatusSuccess)),
							lockorderfulfillment.Not(lockorderfulfillment.ValidationStatusEQ(lockorderfulfillment.ValidationStatusPending)),
						),
					),
				),
				lockpaymentorder.And(
					lockpaymentorder.HasProviderWith(
						providerprofile.VisibilityModeEQ(providerprofile.VisibilityModePrivate),
					),
					lockpaymentorder.StatusEQ(lockpaymentorder.StatusFulfilled),
					lockpaymentorder.HasFulfillmentsWith(
						lockorderfulfillment.ValidationStatusEQ(lockorderfulfillment.ValidationStatusFailed),
						lockorderfulfillment.Not(lockorderfulfillment.ValidationStatusEQ(lockorderfulfillment.ValidationStatusSuccess)),
						lockorderfulfillment.Not(lockorderfulfillment.ValidationStatusEQ(lockorderfulfillment.ValidationStatusPending)),
					),
				),
			),
		).
		WithToken(func(tq *ent.TokenQuery) {
			tq.WithNetwork()
		}).
		All(ctx)
	if err != nil {
		return fmt.Errorf("RetryStaleUserOperations: %w", err)
	}

	wg.Add(1)
	go func() {
		defer wg.Done()
		for _, order := range lockOrders {
			var service types.OrderService
			if strings.HasPrefix(order.Edges.Token.Edges.Network.Identifier, "tron") {
				service = orderService.NewOrderTron()
			} else {
				service = orderService.NewOrderEVM()
			}
			err := service.RefundOrder(ctx, rpcClients[order.Edges.Token.Edges.Network.Identifier], order.Edges.Token.Edges.Network, order.GatewayID)
			if err != nil {
				logger.Errorf("RetryStaleUserOperations.RefundOrder: %v", err)
			}
		}
	}()

	// Retry refunded linked address deposits
	orders, err = storage.Client.PaymentOrder.
		Query().
		Where(
			paymentorder.StatusEQ(paymentorder.StatusRefunded),
			paymentorder.HasLinkedAddress(),
		).
		WithToken(func(tq *ent.TokenQuery) {
			tq.WithNetwork()
		}).
		All(ctx)
	if err != nil {
		return fmt.Errorf("RetryStaleUserOperations: %w", err)
	}

	wg.Add(1)
	go func() {
		defer wg.Done()
		for _, order := range orders {
			service := orderService.NewOrderEVM()
			err = service.CreateOrder(ctx, rpcClients[order.Edges.Token.Edges.Network.Identifier], order.ID)
			if err != nil {
				logger.Errorf("RetryStaleUserOperations.RetryLinkedAddress: %v", err)
			}
		}
	}()

	return nil
}

// IndexBlockchainEvents indexes missed blocks
func IndexBlockchainEvents() error {
	ctx := context.Background()
	var wg sync.WaitGroup

	time.Sleep(100 * time.Millisecond) // to keep out of sync with other tasks

	// Establish RPC connections
	networks, err := setRPCClients(ctx)
	if err != nil {
		return fmt.Errorf("IndexBlockchainEvents: %w", err)
	}

	// Index ERC20 transfer events
	wg.Add(1)
	go func() {
		defer wg.Done()
		_ = utils.Retry(10, 2*time.Second, func() error {
			orders, err := storage.Client.PaymentOrder.
				Query().
				Where(
					paymentorder.StatusEQ(paymentorder.StatusInitiated),
					paymentorder.HasReceiveAddressWith(
						receiveaddress.StatusEQ(receiveaddress.StatusUnused),
						receiveaddress.ValidUntilGT(time.Now()),
					),
				).
				WithToken(func(tq *ent.TokenQuery) {
					tq.WithNetwork()
				}).
				WithReceiveAddress().
				WithRecipient().
				All(ctx)
			if err != nil {
				logger.Errorf("IndexBlockchainEvents: %v", err)
			}

			if len(orders) > 0 {
				for _, order := range orders {
					if strings.HasPrefix(order.Edges.Token.Edges.Network.Identifier, "tron") {
						indexerService := services.NewIndexerService(orderService.NewOrderTron())
						err := indexerService.IndexTRC20Transfer(ctx, order)
						if err != nil {
							continue
						}
					} else {
						indexerService := services.NewIndexerService(orderService.NewOrderEVM())
						err := indexerService.IndexERC20Transfer(ctx, rpcClients[order.Edges.Token.Edges.Network.Identifier], order, nil, 0)
						if err != nil {
							continue
						}
					}
				}
			}

			return fmt.Errorf("trigger retry")
		})
	}()

	// Index OrderCreated events
	wg.Add(1)
	go func() {
		defer wg.Done()
		time.Sleep(500 * time.Millisecond)
		_ = utils.Retry(10, 2*time.Second, func() error {
			for _, network := range networks {
				if strings.HasPrefix(network.Identifier, "tron") {
					orders, err := storage.Client.PaymentOrder.
						Query().
						Where(func(s *sql.Selector) {
							lpo := sql.Table(lockpaymentorder.Table)
							s.Where(sql.And(
								sql.EQ(s.C(paymentorder.FieldStatus), paymentorder.StatusPending),
								sql.Or(
									sql.NotExists(
										sql.Select().
											From(lpo).
											Where(sql.ColumnsEQ(s.C(paymentorder.FieldGatewayID), lpo.C(lockpaymentorder.FieldGatewayID))),
									),
									sql.IsNull(s.C(paymentorder.FieldGatewayID)),
								),
								sql.GT(s.C(paymentorder.FieldBlockNumber), 0),
							))
						}).
						Where(
							paymentorder.HasTokenWith(
								tokenent.HasNetworkWith(networkent.IDEQ(network.ID)),
							),
						).
						WithReceiveAddress().
						WithToken(func(tq *ent.TokenQuery) {
							tq.WithNetwork()
						}).
						Order(ent.Asc(paymentorder.FieldBlockNumber)).
						All(ctx)
					if err != nil {
						continue
					}

					if len(orders) > 0 {
						for _, order := range orders {
							indexerService := services.NewIndexerService(orderService.NewOrderTron())
							err := indexerService.IndexOrderCreatedTron(ctx, order)
							if err != nil {
								continue
							}
						}
					}

				} else {
					indexerService := services.NewIndexerService(orderService.NewOrderEVM())
					err = indexerService.IndexOrderCreated(ctx, rpcClients[network.Identifier], network)
					if err != nil {
						continue
					}
				}
			}

			return fmt.Errorf("trigger retry")
		})
	}()

	// Index OrderSettled events
	wg.Add(1)
	go func() {
		defer wg.Done()
		time.Sleep(1000 * time.Millisecond)
		_ = utils.Retry(10, 2*time.Second, func() error {
			for _, network := range networks {
				if strings.HasPrefix(network.Identifier, "tron") {
					lockOrders, err := storage.Client.LockPaymentOrder.
						Query().
						Where(func(s *sql.Selector) {
							po := sql.Table(paymentorder.Table)
							s.LeftJoin(po).On(s.C(lockpaymentorder.FieldGatewayID), po.C(paymentorder.FieldGatewayID)).
								Where(sql.Or(
									sql.EQ(s.C(lockpaymentorder.FieldStatus), lockpaymentorder.StatusValidated),
									sql.And(
										sql.EQ(po.C(paymentorder.FieldStatus), paymentorder.StatusPending),
										sql.EQ(s.C(lockpaymentorder.FieldStatus), lockpaymentorder.StatusSettled),
									)),
								)
						}).
						Where(
							lockpaymentorder.HasTokenWith(
								tokenent.HasNetworkWith(networkent.IDEQ(network.ID)),
							),
						).
						WithToken(func(tq *ent.TokenQuery) {
							tq.WithNetwork()
						}).
						Order(ent.Asc(lockpaymentorder.FieldBlockNumber)).
						All(ctx)
					if err != nil {
						logger.Errorf("IndexBlockchainEvents: %v", err)
					}

					if len(lockOrders) > 0 {
						for _, order := range lockOrders {
							indexerService := services.NewIndexerService(orderService.NewOrderTron())
							err := indexerService.IndexOrderSettledTron(ctx, order)
							if err != nil {
								continue
							}
						}
					}
				} else {
					indexerService := services.NewIndexerService(orderService.NewOrderEVM())
					err = indexerService.IndexOrderSettled(ctx, rpcClients[network.Identifier], network)
					if err != nil {
						continue
					}
				}
			}

			return fmt.Errorf("trigger retry")
		})
	}()

	// Index OrderRefunded events
	wg.Add(1)
	go func() {
		defer wg.Done()
		time.Sleep(1500 * time.Millisecond)
		_ = utils.Retry(10, 2*time.Second, func() error {
			for _, network := range networks {
				if strings.HasPrefix(network.Identifier, "tron") {
					lockOrders, err := storage.Client.LockPaymentOrder.
						Query().
						Where(func(s *sql.Selector) {
							po := sql.Table(paymentorder.Table)
							s.LeftJoin(po).On(s.C(lockpaymentorder.FieldGatewayID), po.C(paymentorder.FieldGatewayID)).
								Where(sql.Or(
									sql.And(
										sql.EQ(s.C(lockpaymentorder.FieldStatus), lockpaymentorder.StatusPending),
										sql.LT(s.C(lockpaymentorder.FieldCreatedAt), time.Now().Add(-35*time.Minute)),
									),
									sql.And(
										sql.EQ(po.C(paymentorder.FieldStatus), paymentorder.StatusPending),
										sql.EQ(s.C(lockpaymentorder.FieldStatus), lockpaymentorder.StatusRefunded),
									),
								))
						}).
						Where(
							lockpaymentorder.HasTokenWith(
								tokenent.HasNetworkWith(networkent.IDEQ(network.ID)),
							),
						).
						WithToken(func(tq *ent.TokenQuery) {
							tq.WithNetwork()
						}).
						Order(ent.Asc(lockpaymentorder.FieldBlockNumber)).
						All(ctx)
					if err != nil {
						logger.Errorf("IndexBlockchainEvents: %v", err)
					}

					if len(lockOrders) > 0 {
						for _, order := range lockOrders {
							indexerService := services.NewIndexerService(orderService.NewOrderTron())
							err := indexerService.IndexOrderRefundedTron(ctx, order)
							if err != nil {
								continue
							}
						}
					}
				} else {
					indexerService := services.NewIndexerService(orderService.NewOrderEVM())
					err = indexerService.IndexOrderRefunded(ctx, rpcClients[network.Identifier], network)
					if err != nil {
						continue
					}
				}
			}

			return fmt.Errorf("trigger retry")
		})
	}()

	return nil
}

// IndexLinkAddresses indexes ERC20 transfer events to linked addresses
func IndexLinkedAddresses() error {
	ctx := context.Background()

	// Establish RPC connections
	_, err := setRPCClients(ctx)
	if err != nil {
		return fmt.Errorf("IndexLinkedAddresses: %w", err)
	}

	go func() {
		time.Sleep(500 * time.Millisecond)
		_ = utils.Retry(8, 2*time.Second, func() error {
			tokens, err := storage.Client.Token.
				Query().
				Where(
					tokenent.IsEnabled(true),
				).
				WithNetwork().
				All(ctx)
			if err != nil {
				logger.Errorf("IndexLinkedAddresses: %v", err)
			}

			if len(tokens) > 0 {
				for _, token := range tokens {
					indexerService := services.NewIndexerService(orderService.NewOrderEVM())
					err = indexerService.IndexERC20Transfer(ctx, rpcClients[token.Edges.Network.Identifier], nil, token, 0)
					if err != nil {
						continue
					}
				}
			}

			return fmt.Errorf("trigger retry")
		})
	}()

	return nil
}

// ReassignPendingOrders reassigns declined order requests to providers
func ReassignPendingOrders() {
	ctx := context.Background()

	// Remove provider id from pending lock orders
	_, err := storage.Client.LockPaymentOrder.
		Update().
		Where(
			lockpaymentorder.StatusEQ(lockpaymentorder.StatusPending),
			lockpaymentorder.Not(lockpaymentorder.HasFulfillments()),
		).
		ClearProvider().
		Save(ctx)
	if err != nil {
		logger.Errorf("ReassignPendingOrders.db: %v", err)
		return
	}

	// Query pending lock orders
	lockOrders, err := storage.Client.LockPaymentOrder.
		Query().
		Where(
			lockpaymentorder.StatusEQ(lockpaymentorder.StatusPending),
			lockpaymentorder.Not(lockpaymentorder.HasFulfillments()),
		).
		WithToken().
		WithProvider().
		WithProvisionBucket(
			func(pbq *ent.ProvisionBucketQuery) {
				pbq.WithCurrency()
			},
		).
		All(ctx)
	if err != nil {
		logger.Errorf("ReassignPendingOrders.db: %v", err)
		return
	}

	// Check if order_request_<order_id> exists in Redis
	for _, order := range lockOrders {
		orderKey := fmt.Sprintf("order_request_%s", order.ID)
		exists, err := storage.RedisClient.Exists(ctx, orderKey).Result()
		if err != nil {
			logger.Errorf("ReassignPendingOrders.redis: %v", err)
			continue
		}

		if exists == 0 {
			// Order request doesn't exist in Redis, reassign the order
			lockPaymentOrder := types.LockPaymentOrderFields{
				ID:                order.ID,
				Token:             order.Edges.Token,
				GatewayID:         order.GatewayID,
				Amount:            order.Amount,
				Rate:              order.Rate,
				BlockNumber:       order.BlockNumber,
				Institution:       order.Institution,
				AccountIdentifier: order.AccountIdentifier,
				AccountName:       order.AccountName,
				Memo:              order.Memo,
				ProvisionBucket:   order.Edges.ProvisionBucket,
			}

			if order.Edges.Provider != nil {
				lockPaymentOrder.ProviderID = order.Edges.Provider.ID
			}

			err := services.NewPriorityQueueService().AssignLockPaymentOrder(ctx, lockPaymentOrder)
			if err != nil {
				logger.Errorf("failed to reassign declined order request: %v", err)
			}
		}
	}
}

// ReassignUnfulfilledLockOrders reassigns lockOrder unfulfilled within a time frame.
func ReassignUnfulfilledLockOrders() {
	ctx := context.Background()

	// Unassign unfulfilled lock orders.
	_, err := storage.Client.LockPaymentOrder.
		Update().
		Where(
			lockpaymentorder.Or(
				lockpaymentorder.And(
					lockpaymentorder.StatusEQ(lockpaymentorder.StatusProcessing),
					lockpaymentorder.UpdatedAtLTE(time.Now().Add(-orderConf.OrderFulfillmentValidity*time.Minute)),
				),
				lockpaymentorder.StatusEQ(lockpaymentorder.StatusCancelled),
			),
			lockpaymentorder.Or(
				lockpaymentorder.Not(lockpaymentorder.HasFulfillments()),
				lockpaymentorder.HasFulfillmentsWith(
					lockorderfulfillment.ValidationStatusEQ(lockorderfulfillment.ValidationStatusFailed),
					lockorderfulfillment.Not(lockorderfulfillment.ValidationStatusEQ(lockorderfulfillment.ValidationStatusSuccess)),
					lockorderfulfillment.Not(lockorderfulfillment.ValidationStatusEQ(lockorderfulfillment.ValidationStatusPending)),
				),
			),
		).
		SetStatus(lockpaymentorder.StatusPending).
		ClearProvider().
		Save(ctx)
	if err != nil {
		logger.Errorf("ReassignUnfulfilledLockOrders: %v", err)
		return
	}

	// Query unfulfilled lock orders.
	lockOrders, err := storage.Client.LockPaymentOrder.
		Query().
		Where(
			lockpaymentorder.Or(
				lockpaymentorder.Not(lockpaymentorder.HasFulfillments()),
				lockpaymentorder.HasFulfillmentsWith(
					lockorderfulfillment.ValidationStatusEQ(lockorderfulfillment.ValidationStatusFailed),
					lockorderfulfillment.Not(lockorderfulfillment.ValidationStatusEQ(lockorderfulfillment.ValidationStatusSuccess)),
					lockorderfulfillment.Not(lockorderfulfillment.ValidationStatusEQ(lockorderfulfillment.ValidationStatusPending)),
				),
			),
			lockpaymentorder.Or(
				lockpaymentorder.StatusEQ(lockpaymentorder.StatusProcessing),
				lockpaymentorder.StatusEQ(lockpaymentorder.StatusCancelled),
			),
			lockpaymentorder.Or(
				lockpaymentorder.Or(
					lockpaymentorder.And(
						lockpaymentorder.StatusEQ(lockpaymentorder.StatusProcessing),
						lockpaymentorder.UpdatedAtLTE(time.Now().Add(-orderConf.OrderFulfillmentValidity*time.Minute)),
					),
					lockpaymentorder.StatusEQ(lockpaymentorder.StatusCancelled),
				),
				lockpaymentorder.HasFulfillmentsWith(
					lockorderfulfillment.CreatedAtLTE(time.Now().Add(-orderConf.OrderFulfillmentValidity*time.Minute)),
				),
			),
		).
		WithToken().
		WithProvider().
		WithProvisionBucket(func(pbq *ent.ProvisionBucketQuery) {
			pbq.WithCurrency()
		}).
		All(ctx)
	if err != nil {
		logger.Errorf("ReassignUnfulfilledLockOrders: %v", err)
		return
	}

	for _, order := range lockOrders {
		lockPaymentOrder := types.LockPaymentOrderFields{
			ID:                order.ID,
			Token:             order.Edges.Token,
			GatewayID:         order.GatewayID,
			Amount:            order.Amount,
			Rate:              order.Rate,
			BlockNumber:       order.BlockNumber,
			Institution:       order.Institution,
			AccountIdentifier: order.AccountIdentifier,
			AccountName:       order.AccountName,
			Memo:              order.Memo,
			ProvisionBucket:   order.Edges.ProvisionBucket,
		}

		if order.Edges.Provider != nil {
			lockPaymentOrder.ProviderID = order.Edges.Provider.ID
		}

		err := services.NewPriorityQueueService().AssignLockPaymentOrder(ctx, lockPaymentOrder)
		if err != nil {
			logger.Errorf("ReassignUnfulfilledLockOrders.AssignLockPaymentOrder: %s => %v", order.GatewayID, err)
		}
	}
}

// ReassignUnvalidatedLockOrders reassigns or refunds unvalidated lock orders to providers
func ReassignUnvalidatedLockOrders() {
	ctx := context.Background()

	// Query unvalidated lock orders.
	lockOrders, err := storage.Client.LockPaymentOrder.
		Query().
		Where(
			lockpaymentorder.Or(
				lockpaymentorder.StatusEQ(lockpaymentorder.StatusFulfilled),
				lockpaymentorder.And(
					lockpaymentorder.StatusEQ(lockpaymentorder.StatusCancelled),
					lockpaymentorder.HasFulfillmentsWith(
						lockorderfulfillment.ValidationStatusEQ(lockorderfulfillment.ValidationStatusPending),
					),
				),
			),
			lockpaymentorder.Or(
				lockpaymentorder.HasFulfillmentsWith(
					lockorderfulfillment.ValidationStatusEQ(lockorderfulfillment.ValidationStatusFailed),
					lockorderfulfillment.Not(lockorderfulfillment.ValidationStatusEQ(lockorderfulfillment.ValidationStatusSuccess)),
					lockorderfulfillment.Not(lockorderfulfillment.ValidationStatusEQ(lockorderfulfillment.ValidationStatusPending)),
				),
				lockpaymentorder.And(
					lockpaymentorder.HasFulfillmentsWith(
						lockorderfulfillment.ValidationStatusEQ(lockorderfulfillment.ValidationStatusPending),
					),
					lockpaymentorder.HasFulfillmentsWith(
						lockorderfulfillment.UpdatedAtLTE(time.Now().Add(-orderConf.OrderFulfillmentValidity*time.Minute)),
						lockorderfulfillment.Not(lockorderfulfillment.UpdatedAtGT(time.Now().Add(-orderConf.OrderFulfillmentValidity*time.Minute))),
					),
				),
				lockpaymentorder.HasFulfillmentsWith(
					lockorderfulfillment.ValidationStatusEQ(lockorderfulfillment.ValidationStatusSuccess),
				),
			),
		).
		WithToken(func(tq *ent.TokenQuery) {
			tq.WithNetwork()
		}).
		WithProvider(func(pq *ent.ProviderProfileQuery) {
			pq.WithAPIKey()
		}).
		WithFulfillments().
		WithProvisionBucket(func(pb *ent.ProvisionBucketQuery) {
			pb.WithCurrency()
		}).
		All(ctx)
	if err != nil {
		logger.Errorf("ReassignUnvalidatedLockOrders.db: %v", err)
		return
	}

	for _, order := range lockOrders {
		for _, fulfillment := range order.Edges.Fulfillments {
			if fulfillment.ValidationStatus == lockorderfulfillment.ValidationStatusPending {
				// TODO: use auth
				// // Compute HMAC
				// decodedSecret, err := base64.StdEncoding.DecodeString(order.Edges.Provider.Edges.APIKey.Secret)
				// if err != nil {
				// 	logger.Errorf("ReassignUnvalidatedLockOrders: %v", err)
				// 	return
				// }
				// decryptedSecret, err := cryptoUtils.DecryptPlain(decodedSecret)
				// if err != nil {
				// 	logger.Errorf("ReassignUnvalidatedLockOrders: %v", err)
				// 	return
				// }

				// payload := map[string]interface{}{}

				// signature := tokenUtils.GenerateHMACSignature(payload, string(decryptedSecret))

				// Send GET request to the provider's node
				res, err := fastshot.NewClient(order.Edges.Provider.HostIdentifier).
					Config().SetTimeout(30 * time.Second).
					// Header().Add("X-Request-Signature", signature).
					Build().GET(fmt.Sprintf("/tx_status/%s/%s", fulfillment.Psp, fulfillment.TxID)).
					Send()
				if err != nil {
					logger.Errorf("ReassignUnvalidatedLockOrders: %v", err)
					continue
				}

				data, err := utils.ParseJSONResponse(res.RawResponse)
				if err != nil {
					logger.Errorf("ReassignUnvalidatedLockOrders: %v %v", err, data)
					continue
				}

				status := data["data"].(map[string]interface{})["status"].(string)

				if status == "failed" {
					_, err = storage.Client.LockOrderFulfillment.
						UpdateOneID(fulfillment.ID).
						SetValidationStatus(lockorderfulfillment.ValidationStatusFailed).
						SetValidationError(data["data"].(map[string]interface{})["error"].(string)).
						Save(ctx)
					if err != nil {
						logger.Errorf("ReassignUnvalidatedLockOrders.UpdateFulfillmentStatusFailed: %v", err)
						continue
					}

					_, err = order.Update().
						SetStatus(lockpaymentorder.StatusFulfilled).
						Save(ctx)
					if err != nil {
						logger.Errorf("ReassignUnvalidatedLockOrders.UpdateOrderStatusFulfilled: %v", err)
						continue
					}

				} else if status == "success" {
					_, err = storage.Client.LockOrderFulfillment.
						UpdateOneID(fulfillment.ID).
						SetValidationStatus(lockorderfulfillment.ValidationStatusSuccess).
						Save(ctx)
					if err != nil {
						logger.Errorf("ReassignUnvalidatedLockOrders.UpdateFulfillmentStatusSuccess: %v", err)
						continue
					}

					transactionLog, err := storage.Client.TransactionLog.
						Create().
						SetStatus(transactionlog.StatusOrderValidated).
						SetNetwork(order.Edges.Token.Edges.Network.Identifier).
						SetMetadata(map[string]interface{}{
							"TransactionID": fulfillment.TxID,
							"PSP":           fulfillment.Psp,
						}).
						Save(ctx)
					if err != nil {
						logger.Errorf("ReassignUnvalidatedLockOrders.CreateTransactionLog: %v", err)
						continue
					}

					_, err = storage.Client.LockPaymentOrder.
						UpdateOneID(order.ID).
						SetStatus(lockpaymentorder.StatusValidated).
						AddTransactions(transactionLog).
						Save(ctx)
					if err != nil {
						logger.Errorf("ReassignUnvalidatedLockOrders.UpdateOrderStatusValidated: %v", err)
						continue
					}
				}

			} else if fulfillment.ValidationStatus == lockorderfulfillment.ValidationStatusFailed {
				if order.Edges.Provider.VisibilityMode != providerprofile.VisibilityModePrivate {
					lockPaymentOrder := types.LockPaymentOrderFields{
						ID:                order.ID,
						Token:             order.Edges.Token,
						GatewayID:         order.GatewayID,
						Amount:            order.Amount,
						Rate:              order.Rate,
						BlockNumber:       order.BlockNumber,
						Institution:       order.Institution,
						AccountIdentifier: order.AccountIdentifier,
						AccountName:       order.AccountName,
						ProviderID:        "",
						Memo:              order.Memo,
						ProvisionBucket:   order.Edges.ProvisionBucket,
					}

					err := services.NewPriorityQueueService().AssignLockPaymentOrder(ctx, lockPaymentOrder)
					if err != nil {
						logger.Errorf("ReassignUnvalidatedLockOrders.AssignLockPaymentOrder: %v", err)
					}
				}
			} else if fulfillment.ValidationStatus == lockorderfulfillment.ValidationStatusSuccess {
				transactionLog, err := storage.Client.TransactionLog.
					Create().
					SetStatus(transactionlog.StatusOrderValidated).
					SetNetwork(order.Edges.Token.Edges.Network.Identifier).
					SetMetadata(map[string]interface{}{
						"TransactionID": fulfillment.TxID,
						"PSP":           fulfillment.Psp,
					}).
					Save(ctx)
				if err != nil {
					logger.Errorf("ReassignUnvalidatedLockOrders.CreateTransactionLog: %v", err)
					continue
				}

				_, err = storage.Client.LockPaymentOrder.
					UpdateOneID(order.ID).
					SetStatus(lockpaymentorder.StatusValidated).
					AddTransactions(transactionLog).
					Save(ctx)
				if err != nil {
					logger.Errorf("ReassignUnvalidatedLockOrders.UpdateOrderStatusValidated: %v", err)
					continue
				}
			}
		}
	}
}

// ReassignStaleOrderRequest reassigns expired order requests to providers
func ReassignStaleOrderRequest(ctx context.Context, orderRequestChan <-chan *redis.Message) {
	for msg := range orderRequestChan {
		key := strings.Split(msg.Payload, "_")
		orderID := key[len(key)-1]

		orderUUID, err := uuid.Parse(orderID)
		if err != nil {
			logger.Errorf("ReassignStaleOrderRequest: %v", err)
			continue
		}

		// Get the order from the database
		order, err := storage.Client.LockPaymentOrder.
			Query().
			Where(
				lockpaymentorder.IDEQ(orderUUID),
			).
			WithProvisionBucket().
			Only(ctx)
		if err != nil {
			logger.Errorf("ReassignStaleOrderRequest: %v", err)
			continue
		}

		orderFields := types.LockPaymentOrderFields{
			ID:                order.ID,
			GatewayID:         order.GatewayID,
			Amount:            order.Amount,
			Rate:              order.Rate,
			BlockNumber:       order.BlockNumber,
			Institution:       order.Institution,
			AccountIdentifier: order.AccountIdentifier,
			AccountName:       order.AccountName,
			Memo:              order.Memo,
			ProvisionBucket:   order.Edges.ProvisionBucket,
		}

		// Assign the order to a provider
		err = services.NewPriorityQueueService().AssignLockPaymentOrder(ctx, orderFields)
		if err != nil {
			logger.Errorf("ReassignStaleOrderRequest.AssignLockPaymentOrder: %v", err)
		}
	}
}

func FixDatabaseMisHap() error {
	ctx := context.Background()

	// parse string to uuid
	orderUUID, err := uuid.Parse("14baa582-84d9-40bf-96b8-94601d6ffe2b")
	if err != nil {
		logger.Errorf("FixDatabaseMisHap: %v", err)
		return nil
	}

	order, err := storage.Client.PaymentOrder.
		Query().
		Where(
			paymentorder.IDEQ(orderUUID),
			paymentorder.StatusEQ(paymentorder.StatusInitiated),
		).
		WithToken(func(tq *ent.TokenQuery) {
			tq.WithNetwork()
		}).
		WithRecipient().
		Only(ctx)
	if err != nil {
		logger.Errorf("FixDatabaseMisHap: %v", err)
	}

	service := orderService.NewOrderEVM()
	err = service.CreateOrder(ctx, rpcClients[order.Edges.Token.Edges.Network.Identifier], order.ID)
	if err != nil {
		logger.Errorf("FixDatabaseMisHap: %v", err)
	}

	return nil
}

// HandleReceiveAddressValidity handles receive address validity
func HandleReceiveAddressValidity() error {
	ctx := context.Background()

	// Establish RPC connections
	_, err := setRPCClients(ctx)
	if err != nil {
		return fmt.Errorf("HandleReceiveAddressValidity: %w", err)
	}

	// Fetch expired receive addresses that are due for validity check
	addresses, err := storage.Client.ReceiveAddress.
		Query().
		Where(
			receiveaddress.ValidUntilLTE(time.Now()),
			receiveaddress.Or(
				receiveaddress.StatusNEQ(receiveaddress.StatusUsed),
				receiveaddress.And(
					receiveaddress.StatusEQ(receiveaddress.StatusUsed),
					receiveaddress.HasPaymentOrderWith(
						paymentorder.StatusEQ(paymentorder.StatusInitiated),
					),
				),
			),
			receiveaddress.HasPaymentOrder(),
		).
		WithPaymentOrder(func(po *ent.PaymentOrderQuery) {
			po.WithToken(func(tq *ent.TokenQuery) {
				tq.WithNetwork()
			})
			po.WithRecipient()
		}).
		All(ctx)
	if err != nil {
		return fmt.Errorf("HandleReceiveAddressValidity: %w", err)
	}

	var indexerService services.Indexer
	for _, address := range addresses {
		if strings.HasPrefix(address.Edges.PaymentOrder.Edges.Token.Edges.Network.Identifier, "tron") {
			indexerService = services.NewIndexerService(orderService.NewOrderTron())
		} else {
			indexerService = services.NewIndexerService(orderService.NewOrderEVM())
		}

		err := indexerService.HandleReceiveAddressValidity(ctx, rpcClients[address.Edges.PaymentOrder.Edges.Token.Edges.Network.Identifier], address, address.Edges.PaymentOrder)
		if err != nil {
			continue
		}
	}

	return nil
}

// SubscribeToRedisKeyspaceEvents subscribes to redis keyspace events according to redis.conf settings
func SubscribeToRedisKeyspaceEvents() {
	ctx := context.Background()

	// Handle expired or deleted order request key events
	orderRequest := storage.RedisClient.PSubscribe(
		ctx,
		"__keyevent@0__:expired:order_request_*",
		"__keyevent@0__:del:order_request_*",
	)
	orderRequestChan := orderRequest.Channel()

	go ReassignStaleOrderRequest(ctx, orderRequestChan)
}

// fetchExternalRate fetches the external rate for a fiat currency
func fetchExternalRate(currency string) (decimal.Decimal, error) {
	currency = strings.ToUpper(currency)
	supportedCurrencies := []string{"KES", "NGN", "GHS", "TZS", "UGX", "XOF"}
	isSupported := false
	for _, supported := range supportedCurrencies {
		if currency == supported {
			isSupported = true
			break
		}
	}
	if !isSupported {
		return decimal.Zero, fmt.Errorf("ComputeMarketRate: currency not supported")
	}

	// Fetch rates from third-party APIs
	var price decimal.Decimal
	if currency == "NGN" {
		res, err := fastshot.NewClient("https://www.quidax.com").
			Config().SetTimeout(30*time.Second).
			Build().GET(fmt.Sprintf("/api/v1/markets/tickers/usdt%s", strings.ToLower(currency))).
			Retry().Set(3, 5*time.Second).
			Send()
		if err != nil {
			return decimal.Zero, fmt.Errorf("ComputeMarketRate: %w", err)
		}

		data, err := utils.ParseJSONResponse(res.RawResponse)
		if err != nil {
			return decimal.Zero, fmt.Errorf("ComputeMarketRate: %w %v", err, data)
		}

		price, err = decimal.NewFromString(data["data"].(map[string]interface{})["ticker"].(map[string]interface{})["buy"].(string))
		if err != nil {
			return decimal.Zero, fmt.Errorf("ComputeMarketRate: %w", err)
		}
	} else {
		res, err := fastshot.NewClient("https://p2p.binance.com").
			Config().SetTimeout(30*time.Second).
			Header().Add("Content-Type", "application/json").
			Build().POST("/bapi/c2c/v2/friendly/c2c/adv/search").
			Retry().Set(3, 5*time.Second).
			Body().AsJSON(map[string]interface{}{
			"asset":     "USDT",
			"fiat":      currency,
			"tradeType": "SELL",
			"page":      1,
			"rows":      20,
		}).
			Send()
		if err != nil {
			return decimal.Zero, fmt.Errorf("ComputeMarketRate: %w", err)
		}

		resData, err := utils.ParseJSONResponse(res.RawResponse)
		if err != nil {
			return decimal.Zero, fmt.Errorf("ComputeMarketRate: %w", err)
		}

		// Access the data array
		data, ok := resData["data"].([]interface{})
		if !ok || len(data) == 0 {
			return decimal.Zero, fmt.Errorf("ComputeMarketRate: No data in the response")
		}

		// Loop through the data array and extract prices
		var prices []decimal.Decimal
		for _, item := range data {
			adv, ok := item.(map[string]interface{})["adv"].(map[string]interface{})
			if !ok {
				continue
			}

			price, err := decimal.NewFromString(adv["price"].(string))
			if err != nil {
				continue
			}

			prices = append(prices, price)
		}

		// Calculate and return the median
		price = utils.Median(prices)
	}

	return price, nil
}

// ComputeMarketRate computes the market price for fiat currencies
func ComputeMarketRate() error {
	ctx := context.Background()

	// Fetch all fiat currencies
	currencies, err := storage.Client.FiatCurrency.
		Query().
		Where(fiatcurrency.IsEnabledEQ(true)).
		All(ctx)
	if err != nil {
		return fmt.Errorf("ComputeMarketRate: %w", err)
	}

	for _, currency := range currencies {
		// Fetch external rate
		externalRate, err := fetchExternalRate(currency.Code)
		if err != nil {
			continue
		}

		// Fetch rates from token configs with fixed conversion rate
		tokenConfigs, err := storage.Client.ProviderOrderToken.
			Query().
			Where(
				providerordertoken.SymbolIn("USDT", "USDC"),
				providerordertoken.ConversionRateTypeEQ(providerordertoken.ConversionRateTypeFixed),
			).
			Select(providerordertoken.FieldFixedConversionRate).
			All(ctx)
		if err != nil {
			continue
		}

		var rates []decimal.Decimal
		for _, tokenConfig := range tokenConfigs {
			rates = append(rates, tokenConfig.FixedConversionRate)
		}

		// Calculate median
		median := utils.Median(rates)

		// Check the median rate against the external rate to ensure it's not too far off
		percentDeviation := utils.AbsPercentageDeviation(externalRate, median)
		if percentDeviation.GreaterThan(orderConf.PercentDeviationFromExternalRate) {
			median = externalRate
		}

		// Update currency with median rate
		_, err = storage.Client.FiatCurrency.
			UpdateOneID(currency.ID).
			SetMarketRate(median).
			Save(ctx)
		if err != nil {
			continue
		}
	}

	return nil
}

// Retry failed webhook notifications
func RetryFailedWebhookNotifications() error {
	ctx := context.Background()

	// Fetch failed webhook notifications that are due for retry
	attempts, err := storage.Client.WebhookRetryAttempt.
		Query().
		Where(
			webhookretryattempt.StatusEQ(webhookretryattempt.StatusFailed),
			webhookretryattempt.NextRetryTimeLTE(time.Now()),
		).
		All(ctx)
	if err != nil {
		return fmt.Errorf("RetryFailedWebhookNotifications: %w", err)
	}

	baseDelay := 2 * time.Minute
	maxCumulativeTime := 24 * time.Hour

	for _, attempt := range attempts {
		// Send the webhook notification
		body, err := fastshot.NewClient(attempt.WebhookURL).
			Config().SetTimeout(30*time.Second).
			Header().Add("X-Paycrest-Signature", attempt.Signature).
			Build().POST("").
			Body().AsJSON(attempt.Payload).
			Send()

		if err != nil || (body.StatusCode() >= 205) {
			// Webhook notification failed
			// Update attempt with next retry time
			attemptNumber := attempt.AttemptNumber + 1
			delay := baseDelay * time.Duration(math.Pow(2, float64(attemptNumber-1)))

			nextRetryTime := time.Now().Add(delay)

			attemptUpdate := attempt.Update()

			attemptUpdate.
				AddAttemptNumber(1).
				SetNextRetryTime(nextRetryTime)

			// Set status to expired if cumulative time is greater than 24 hours
			if nextRetryTime.Sub(attempt.CreatedAt.Add(-baseDelay)) > maxCumulativeTime {
				attemptUpdate.SetStatus(webhookretryattempt.StatusExpired)
				uid, err := uuid.Parse(attempt.Payload["data"].(map[string]interface{})["senderId"].(string))
				if err != nil {
					return fmt.Errorf("RetryFailedWebhookNotifications.FailedExtraction: %w", err)
				}
				profile, err := storage.Client.SenderProfile.
					Query().
					Where(
						senderprofile.IDEQ(uid),
					).
					WithUser().Only(ctx)
				if err != nil {
					return fmt.Errorf("RetryFailedWebhookNotifications.CouldNotFetchProfile: %w", err)
				}

				_, err = services.SendTemplateEmail(types.SendEmailPayload{
					ToAddress: profile.Edges.User.Email,
					DynamicData: map[string]interface{}{
						"first_name": profile.Edges.User.FirstName,
					},
				}, "d-da75eee4966544ad92dcd060421d4e12")

				if err != nil {
					return fmt.Errorf("RetryFailedWebhookNotifications.SendTemplateEmail: %w", err)
				}
			}

			_, err := attemptUpdate.Save(ctx)
			if err != nil {
				return fmt.Errorf("RetryFailedWebhookNotifications: %w", err)
			}

			continue
		}

		// Webhook notification was successful
		_, err = attempt.Update().
			SetStatus(webhookretryattempt.StatusSuccess).
			Save(ctx)
		if err != nil {
			return fmt.Errorf("RetryFailedWebhookNotifications: %w", err)
		}
	}

	return nil
}

// StartCronJobs starts cron jobs
func StartCronJobs() {
	scheduler := gocron.NewScheduler(time.UTC)
	priorityQueue := services.NewPriorityQueueService()

	err := ComputeMarketRate()
	if err != nil {
		logger.Errorf("StartCronJobs: %v", err)
	}

	err = priorityQueue.ProcessBucketQueues()
	if err != nil {
		logger.Errorf("StartCronJobs: %v", err)
	}

	// Compute market rate every 30 minutes
	_, err = scheduler.Cron("*/30 * * * *").Do(ComputeMarketRate)
	if err != nil {
		logger.Errorf("StartCronJobs: %v", err)
	}

	// Refresh provision bucket priority queues every X minutes
	_, err = scheduler.Cron(fmt.Sprintf("*/%d * * * *", orderConf.BucketQueueRebuildInterval)).
		Do(priorityQueue.ProcessBucketQueues)
	if err != nil {
		logger.Errorf("StartCronJobs: %v", err)
	}

	// Retry failed webhook notifications every 59 minutes
	_, err = scheduler.Cron("*/59 * * * *").Do(RetryFailedWebhookNotifications)
	if err != nil {
		logger.Errorf("StartCronJobs: %v", err)
	}

	// Reassign pending order requests every 13 minutes
	// _, err = scheduler.Cron("*/13 * * * *").Do(ReassignPendingOrders)
	// if err != nil {
	// 	logger.Errorf("StartCronJobs: %v", err)
	// }

	// Reassign unvalidated order requests every 2 minutes
	_, err = scheduler.Cron("*/2 * * * *").Do(ReassignUnvalidatedLockOrders)
	if err != nil {
		logger.Errorf("StartCronJobs: %v", err)
	}

	// Reassign unfulfilled order requests every 3 minutes
	// _, err = scheduler.Cron("*/3 * * * *").Do(ReassignUnfulfilledLockOrders)
	// if err != nil {
	// 	logger.Errorf("StartCronJobs: %v", err)
	// }

	// Handle receive address validity every 31 minutes
	_, err = scheduler.Cron("*/31 * * * *").Do(HandleReceiveAddressValidity)
	if err != nil {
		logger.Errorf("StartCronJobs: %v", err)
	}

	// Retry stale user operations every 2 minutes
	_, err = scheduler.Cron("*/2 * * * *").Do(RetryStaleUserOperations)
	if err != nil {
		logger.Errorf("StartCronJobs: %v", err)
	}

	// Index blockchain events every 1 minute
	_, err = scheduler.Cron("*/1 * * * *").Do(IndexBlockchainEvents)
	if err != nil {
		logger.Errorf("StartCronJobs: %v", err)
	}

	// Index linked addresses every 1 minute
	_, err = scheduler.Cron("*/1 * * * *").Do(IndexLinkedAddresses)
	if err != nil {
		logger.Errorf("StartCronJobs: %v", err)
	}

	// Start scheduler
	scheduler.StartAsync()
}
