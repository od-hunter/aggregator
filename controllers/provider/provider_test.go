package provider

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"testing"
	"time"

	"github.com/google/uuid"
	_ "github.com/mattn/go-sqlite3"
	"github.com/paycrest/protocol/ent"
	"github.com/paycrest/protocol/routers/middleware"
	"github.com/paycrest/protocol/services"
	db "github.com/paycrest/protocol/storage"
	"github.com/paycrest/protocol/types"

	"github.com/gin-gonic/gin"
	"github.com/paycrest/protocol/ent/enttest"
	"github.com/paycrest/protocol/utils/test"
	"github.com/paycrest/protocol/utils/token"
	"github.com/stretchr/testify/assert"
)

var testCtx = struct {
	user         *ent.User
	apiKey       *ent.APIKey
	apiKeySecret string
	lockOrder    *ent.LockPaymentOrder
}{}

func setup() error {
	// Set up test data
	user, err := test.CreateTestUser(map[string]string{
		"scope": "provider"})
	if err != nil {
		return err
	}
	testCtx.user = user

	currency, err := test.CreateTestFiatCurrency(nil)
	if err != nil {
		return err
	}

	provderProfile, err := test.CreateTestProviderProfile(nil, testCtx.user, currency)
	if err != nil {
		return err
	}

	for i := 0; i < 10; i++ {
		_, err := test.CreateTestLockPaymentOrder(map[string]interface{}{
			"order_id": uuid.New().String(),
			"provider": provderProfile,
		})
		if err != nil {
			return err
		}
	}

	apiKeyService := services.NewAPIKeyService()
	apiKey, secretKey, err := apiKeyService.GenerateAPIKey(
		context.Background(),
		nil,
		nil,
		provderProfile,
	)
	if err != nil {
		return err
	}

	testCtx.apiKey = apiKey
	testCtx.apiKeySecret = secretKey

	return nil
}

func TestProvider(t *testing.T) {

	// Set up test database client
	client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&_fk=1")
	defer client.Close()

	db.Client = client

	// Setup test data
	err := setup()
	assert.NoError(t, err)

	// Set up test routers
	router := gin.New()
	router.Use(middleware.DynamicAuthMiddleware)
	router.Use(middleware.OnlyProviderMiddleware)

	// Create a new instance of the SenderController with the mock service
	ctrl := NewProviderController()
	router.GET("/orders/", ctrl.GetLockPaymentOrders)

	t.Run("GetLockPaymentOrders", func(t *testing.T) {
		var payload = map[string]interface{}{
			"timestamp": time.Now().Unix(),
		}

		signature := token.GenerateHMACSignature(payload, testCtx.apiKeySecret)

		headers := map[string]string{
			"Authorization": "HMAC " + testCtx.apiKey.ID.String() + ":" + signature,
			"Client-Type":   "backend",
		}

		//query params
		page := 1
		pageSize := 5
		status := "pending"
		ordering := "desc"

		res, err := test.PerformRequest(t, "GET", fmt.Sprintf("/orders/?page=%s&pageSize=%s&status=%s&ordering=%s", strconv.Itoa(page), strconv.Itoa(pageSize), status, ordering), payload, headers, router)
		assert.NoError(t, err)

		// Assert the response body
		assert.Equal(t, http.StatusOK, res.Code)

		var response types.Response
		err = json.Unmarshal(res.Body.Bytes(), &response)
		assert.NoError(t, err)
		assert.Equal(t, "Orders successfully retrieved", response.Message)
		data, ok := response.Data.(map[string]interface{})
		assert.True(t, ok, "response.Data is of not type map[string]interface{}")
		assert.NotNil(t, data, "response.Data is nil")

		assert.Equal(t, int(data["page"].(float64)), page)
		assert.Equal(t, int(data["pageSize"].(float64)), pageSize)
		assert.NotNil(t, data["total"])
		assert.NotEmpty(t, data["orders"])
	})
}
