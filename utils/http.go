package utils

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/paycrest/protocol/types"
)

// APIResponse is a helper function to return an API response
func APIResponse(ctx *gin.Context, httpCode int, status string, message string, data interface{}) {
	ctx.JSON(httpCode, types.Response{
		Status:  status,
		Message: message,
		Data:    data,
	})
}

// GetErrorMsg returns a list of meaningful error messages from binding tags.
// Reference: https://blog.logrocket.com/gin-binding-in-go-a-tutorial-with-examples/
func GetErrorMsg(fe validator.FieldError) string {
	switch fe.Tag() {
	case "required":
		return "This field is required"
	case "lte":
		return "Should be less than " + fe.Param()
	case "gte":
		return "Should be greater than " + fe.Param()
	case "email":
		return "Must be a valid email address"
	case "min":
		return "Should be at least " + fe.Param() + " characters"
	case "max":
		return "Should be at most " + fe.Param() + " characters"
	case "oneof":
		options := strings.Split(fe.Param(), ",")
		return "Must be one of " + strings.Join(options, ", ")
	}
	return "Unknown error"
}

// GetErrorData returns a list of error data
func GetErrorData(err error) []types.ErrorData {
	var errorData []types.ErrorData
	for _, fe := range err.(validator.ValidationErrors) {
		errorData = append(errorData, types.ErrorData{
			Field:   fe.Field(),
			Message: GetErrorMsg(fe),
		})
	}
	return errorData
}

// MakeJSONRequest makes a JSON request
func MakeJSONRequest(ctx context.Context, method, url string, payload map[string]interface{}, headers map[string]string) (responseData map[string]interface{}, err error) {
	if !ContainsString([]string{"GET", "POST", "PUT", "PATCH", "DELETE"}, method) {
		return nil, errors.New("invalid method")
	}

	// Create a new request
	requestBody, _ := json.Marshal(payload)
	req, err := http.NewRequest(method, url, bytes.NewReader(requestBody))
	if err != nil {
		return nil, err
	}

	// Set headers
	for key, value := range headers {
		req.Header.Set(key, value)
	}
	req.Header.Set("Content-Type", "application/json")

	// Create a new context and add the request to it
	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	// Make the request
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	// Decode the response body into a map
	responseBody, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var body map[string]interface{}
	err = json.Unmarshal(responseBody, &body)
	if err != nil {
		return nil, err
	}

	if res.StatusCode >= 400 {
		return body, errors.New(fmt.Sprint(res.StatusCode))
	}

	return body, nil
}
