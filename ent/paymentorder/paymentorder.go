// Code generated by ent, DO NOT EDIT.

package paymentorder

import (
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the paymentorder type in the database.
	Label = "payment_order"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldAmount holds the string denoting the amount field in the database.
	FieldAmount = "amount"
	// FieldAmountPaid holds the string denoting the amount_paid field in the database.
	FieldAmountPaid = "amount_paid"
	// FieldAmountReturned holds the string denoting the amount_returned field in the database.
	FieldAmountReturned = "amount_returned"
	// FieldPercentSettled holds the string denoting the percent_settled field in the database.
	FieldPercentSettled = "percent_settled"
	// FieldSenderFee holds the string denoting the sender_fee field in the database.
	FieldSenderFee = "sender_fee"
	// FieldNetworkFee holds the string denoting the network_fee field in the database.
	FieldNetworkFee = "network_fee"
	// FieldProtocolFee holds the string denoting the protocol_fee field in the database.
	FieldProtocolFee = "protocol_fee"
	// FieldRate holds the string denoting the rate field in the database.
	FieldRate = "rate"
	// FieldTxHash holds the string denoting the tx_hash field in the database.
	FieldTxHash = "tx_hash"
	// FieldBlockNumber holds the string denoting the block_number field in the database.
	FieldBlockNumber = "block_number"
	// FieldFromAddress holds the string denoting the from_address field in the database.
	FieldFromAddress = "from_address"
	// FieldReturnAddress holds the string denoting the return_address field in the database.
	FieldReturnAddress = "return_address"
	// FieldReceiveAddressText holds the string denoting the receive_address_text field in the database.
	FieldReceiveAddressText = "receive_address_text"
	// FieldFeePercent holds the string denoting the fee_percent field in the database.
	FieldFeePercent = "fee_percent"
	// FieldFeeAddress holds the string denoting the fee_address field in the database.
	FieldFeeAddress = "fee_address"
	// FieldGatewayID holds the string denoting the gateway_id field in the database.
	FieldGatewayID = "gateway_id"
	// FieldStatus holds the string denoting the status field in the database.
	FieldStatus = "status"
	// EdgeSenderProfile holds the string denoting the sender_profile edge name in mutations.
	EdgeSenderProfile = "sender_profile"
	// EdgeToken holds the string denoting the token edge name in mutations.
	EdgeToken = "token"
	// EdgeReceiveAddress holds the string denoting the receive_address edge name in mutations.
	EdgeReceiveAddress = "receive_address"
	// EdgeRecipient holds the string denoting the recipient edge name in mutations.
	EdgeRecipient = "recipient"
	// EdgeTransactions holds the string denoting the transactions edge name in mutations.
	EdgeTransactions = "transactions"
	// Table holds the table name of the paymentorder in the database.
	Table = "payment_orders"
	// SenderProfileTable is the table that holds the sender_profile relation/edge.
	SenderProfileTable = "payment_orders"
	// SenderProfileInverseTable is the table name for the SenderProfile entity.
	// It exists in this package in order to avoid circular dependency with the "senderprofile" package.
	SenderProfileInverseTable = "sender_profiles"
	// SenderProfileColumn is the table column denoting the sender_profile relation/edge.
	SenderProfileColumn = "sender_profile_payment_orders"
	// TokenTable is the table that holds the token relation/edge.
	TokenTable = "payment_orders"
	// TokenInverseTable is the table name for the Token entity.
	// It exists in this package in order to avoid circular dependency with the "token" package.
	TokenInverseTable = "tokens"
	// TokenColumn is the table column denoting the token relation/edge.
	TokenColumn = "token_payment_orders"
	// ReceiveAddressTable is the table that holds the receive_address relation/edge.
	ReceiveAddressTable = "receive_addresses"
	// ReceiveAddressInverseTable is the table name for the ReceiveAddress entity.
	// It exists in this package in order to avoid circular dependency with the "receiveaddress" package.
	ReceiveAddressInverseTable = "receive_addresses"
	// ReceiveAddressColumn is the table column denoting the receive_address relation/edge.
	ReceiveAddressColumn = "payment_order_receive_address"
	// RecipientTable is the table that holds the recipient relation/edge.
	RecipientTable = "payment_order_recipients"
	// RecipientInverseTable is the table name for the PaymentOrderRecipient entity.
	// It exists in this package in order to avoid circular dependency with the "paymentorderrecipient" package.
	RecipientInverseTable = "payment_order_recipients"
	// RecipientColumn is the table column denoting the recipient relation/edge.
	RecipientColumn = "payment_order_recipient"
	// TransactionsTable is the table that holds the transactions relation/edge.
	TransactionsTable = "transaction_logs"
	// TransactionsInverseTable is the table name for the TransactionLog entity.
	// It exists in this package in order to avoid circular dependency with the "transactionlog" package.
	TransactionsInverseTable = "transaction_logs"
	// TransactionsColumn is the table column denoting the transactions relation/edge.
	TransactionsColumn = "payment_order_transactions"
)

// Columns holds all SQL columns for paymentorder fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldAmount,
	FieldAmountPaid,
	FieldAmountReturned,
	FieldPercentSettled,
	FieldSenderFee,
	FieldNetworkFee,
	FieldProtocolFee,
	FieldRate,
	FieldTxHash,
	FieldBlockNumber,
	FieldFromAddress,
	FieldReturnAddress,
	FieldReceiveAddressText,
	FieldFeePercent,
	FieldFeeAddress,
	FieldGatewayID,
	FieldStatus,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "payment_orders"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"api_key_payment_orders",
	"sender_profile_payment_orders",
	"token_payment_orders",
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() time.Time
	// UpdateDefaultUpdatedAt holds the default value on update for the "updated_at" field.
	UpdateDefaultUpdatedAt func() time.Time
	// TxHashValidator is a validator for the "tx_hash" field. It is called by the builders before save.
	TxHashValidator func(string) error
	// DefaultBlockNumber holds the default value on creation for the "block_number" field.
	DefaultBlockNumber int64
	// FromAddressValidator is a validator for the "from_address" field. It is called by the builders before save.
	FromAddressValidator func(string) error
	// ReturnAddressValidator is a validator for the "return_address" field. It is called by the builders before save.
	ReturnAddressValidator func(string) error
	// ReceiveAddressTextValidator is a validator for the "receive_address_text" field. It is called by the builders before save.
	ReceiveAddressTextValidator func(string) error
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() uuid.UUID
)

// Status defines the type for the "status" enum field.
type Status string

// StatusInitiated is the default value of the Status enum.
const DefaultStatus = StatusInitiated

// Status values.
const (
	StatusInitiated Status = "initiated"
	StatusPending   Status = "pending"
	StatusExpired   Status = "expired"
	StatusSettled   Status = "settled"
	StatusRefunded  Status = "refunded"
)

func (s Status) String() string {
	return string(s)
}

// StatusValidator is a validator for the "status" field enum values. It is called by the builders before save.
func StatusValidator(s Status) error {
	switch s {
	case StatusInitiated, StatusPending, StatusExpired, StatusSettled, StatusRefunded:
		return nil
	default:
		return fmt.Errorf("paymentorder: invalid enum value for status field: %q", s)
	}
}

// OrderOption defines the ordering options for the PaymentOrder queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByCreatedAt orders the results by the created_at field.
func ByCreatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedAt, opts...).ToFunc()
}

// ByUpdatedAt orders the results by the updated_at field.
func ByUpdatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUpdatedAt, opts...).ToFunc()
}

// ByAmount orders the results by the amount field.
func ByAmount(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldAmount, opts...).ToFunc()
}

// ByAmountPaid orders the results by the amount_paid field.
func ByAmountPaid(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldAmountPaid, opts...).ToFunc()
}

// ByAmountReturned orders the results by the amount_returned field.
func ByAmountReturned(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldAmountReturned, opts...).ToFunc()
}

// ByPercentSettled orders the results by the percent_settled field.
func ByPercentSettled(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPercentSettled, opts...).ToFunc()
}

// BySenderFee orders the results by the sender_fee field.
func BySenderFee(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldSenderFee, opts...).ToFunc()
}

// ByNetworkFee orders the results by the network_fee field.
func ByNetworkFee(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldNetworkFee, opts...).ToFunc()
}

// ByProtocolFee orders the results by the protocol_fee field.
func ByProtocolFee(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldProtocolFee, opts...).ToFunc()
}

// ByRate orders the results by the rate field.
func ByRate(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldRate, opts...).ToFunc()
}

// ByTxHash orders the results by the tx_hash field.
func ByTxHash(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldTxHash, opts...).ToFunc()
}

// ByBlockNumber orders the results by the block_number field.
func ByBlockNumber(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldBlockNumber, opts...).ToFunc()
}

// ByFromAddress orders the results by the from_address field.
func ByFromAddress(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldFromAddress, opts...).ToFunc()
}

// ByReturnAddress orders the results by the return_address field.
func ByReturnAddress(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldReturnAddress, opts...).ToFunc()
}

// ByReceiveAddressText orders the results by the receive_address_text field.
func ByReceiveAddressText(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldReceiveAddressText, opts...).ToFunc()
}

// ByFeePercent orders the results by the fee_percent field.
func ByFeePercent(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldFeePercent, opts...).ToFunc()
}

// ByFeeAddress orders the results by the fee_address field.
func ByFeeAddress(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldFeeAddress, opts...).ToFunc()
}

// ByGatewayID orders the results by the gateway_id field.
func ByGatewayID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldGatewayID, opts...).ToFunc()
}

// ByStatus orders the results by the status field.
func ByStatus(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldStatus, opts...).ToFunc()
}

// BySenderProfileField orders the results by sender_profile field.
func BySenderProfileField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newSenderProfileStep(), sql.OrderByField(field, opts...))
	}
}

// ByTokenField orders the results by token field.
func ByTokenField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newTokenStep(), sql.OrderByField(field, opts...))
	}
}

// ByReceiveAddressField orders the results by receive_address field.
func ByReceiveAddressField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newReceiveAddressStep(), sql.OrderByField(field, opts...))
	}
}

// ByRecipientField orders the results by recipient field.
func ByRecipientField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newRecipientStep(), sql.OrderByField(field, opts...))
	}
}

// ByTransactionsCount orders the results by transactions count.
func ByTransactionsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newTransactionsStep(), opts...)
	}
}

// ByTransactions orders the results by transactions terms.
func ByTransactions(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newTransactionsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newSenderProfileStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(SenderProfileInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, SenderProfileTable, SenderProfileColumn),
	)
}
func newTokenStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(TokenInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, TokenTable, TokenColumn),
	)
}
func newReceiveAddressStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(ReceiveAddressInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2O, false, ReceiveAddressTable, ReceiveAddressColumn),
	)
}
func newRecipientStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(RecipientInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2O, false, RecipientTable, RecipientColumn),
	)
}
func newTransactionsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(TransactionsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, TransactionsTable, TransactionsColumn),
	)
}
