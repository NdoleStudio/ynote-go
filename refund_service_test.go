package ynote

import (
	"context"
	"net/http"
	"testing"

	"github.com/NdoleStudio/ynote-go/internal/helpers"
	"github.com/NdoleStudio/ynote-go/internal/stubs"
	"github.com/stretchr/testify/assert"
)

func TestRefundService_Refund(t *testing.T) {
	// Setup
	t.Parallel()

	// Arrange
	requests := make([]http.Request, 0)
	responses := [][]byte{stubs.TokenResponse(), stubs.RefundResponse()}
	server := helpers.MakeRequestCapturingTestServer([]int{http.StatusOK, http.StatusOK}, responses, &requests)
	client := New(
		WithTokenURL(server.URL),
		WithAPIURL(server.URL),
		WithClientID(testClientID),
		WithClientSecret(testClientSecret),
	)

	payload := &RefundParams{
		ChannelUserMsisdn:         "699999999",
		Pin:                       "0000",
		Webhook:                   "https://example.com/webhook",
		Amount:                    "100",
		FinalCustomerPhone:        "699999999",
		FinalCustomerName:         "",
		RefundMethod:              "OrangeMoney",
		FeesIncluded:              false,
		FinalCustomerNameAccuracy: "0",
	}

	// Act
	transaction, response, err := client.Refund.Refund(context.Background(), payload)

	// Assert
	assert.Nil(t, err)

	assert.GreaterOrEqual(t, len(requests), 2)
	request := requests[len(requests)-1]

	assert.Equal(t, "/prod/refund", request.URL.Path)
	assert.Equal(t, "Bearer 19077204-9d0a-31fa-85cf-xxxxxxxxxx", request.Header.Get("Authorization"))

	assert.Equal(t, http.StatusOK, response.HTTPResponse.StatusCode)
	assert.Equal(t, "993764f9-6b1f-41bd-a7ca-97b8b2167ed7", transaction.MessageID)

	// Teardown
	server.Close()
}

func TestRefundService_RefundWithInvalidClient(t *testing.T) {
	// Setup
	t.Parallel()

	// Arrange
	requests := make([]http.Request, 0)
	responses := [][]byte{stubs.TokenResponse(), stubs.RefundInvalidClientResponse()}
	server := helpers.MakeRequestCapturingTestServer([]int{http.StatusOK, http.StatusBadRequest}, responses, &requests)
	client := New(
		WithTokenURL(server.URL),
		WithAPIURL(server.URL),
		WithClientID(testClientID),
		WithClientSecret(testClientSecret),
	)

	payload := &RefundParams{
		ChannelUserMsisdn:         "699999999",
		Pin:                       "0000",
		Webhook:                   "https://api.nyangapay.com/v1/y-note",
		Amount:                    "100",
		FinalCustomerPhone:        "699999999",
		FinalCustomerName:         "",
		RefundMethod:              "OrangeMoney",
		FeesIncluded:              false,
		FinalCustomerNameAccuracy: "0",
	}

	// Act
	transaction, response, err := client.Refund.Refund(context.Background(), payload)

	// Assert
	assert.Nil(t, transaction)
	assert.NotNil(t, err)

	assert.Equal(t, http.StatusBadRequest, response.HTTPResponse.StatusCode)

	// Teardown
	server.Close()
}

func TestRefundService_Status(t *testing.T) {
	// Setup
	t.Parallel()

	// Arrange
	requests := make([]http.Request, 0)
	responses := [][]byte{stubs.TokenResponse(), stubs.RefundStatusResponse()}
	server := helpers.MakeRequestCapturingTestServer([]int{http.StatusOK, http.StatusOK}, responses, &requests)
	client := New(
		WithTokenURL(server.URL),
		WithAPIURL(server.URL),
		WithClientID(testClientID),
		WithClientSecret(testClientSecret),
	)

	// Act
	transaction, response, err := client.Refund.Status(context.Background(), "")

	// Assert
	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, response.HTTPResponse.StatusCode)
	assert.Equal(t, "CI24120168FBF65A909F588B4480", transaction.Result.Data.PayToken)
	assert.True(t, transaction.IsSuccessful())
	assert.False(t, transaction.IsFailed())
	assert.False(t, transaction.IsPending())

	// Teardown
	server.Close()
}

func TestRefundService_StatusWithFailure(t *testing.T) {
	// Setup
	t.Parallel()

	// Arrange
	requests := make([]http.Request, 0)
	responses := [][]byte{stubs.TokenResponse(), stubs.RefundStatusResponseWithFailure()}
	server := helpers.MakeRequestCapturingTestServer([]int{http.StatusOK, http.StatusOK}, responses, &requests)
	client := New(
		WithTokenURL(server.URL),
		WithAPIURL(server.URL),
		WithClientID(testClientID),
		WithClientSecret(testClientSecret),
	)

	// Act
	transaction, response, err := client.Refund.Status(context.Background(), "")

	// Assert
	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, response.HTTPResponse.StatusCode)
	assert.Nil(t, transaction.Result)
	assert.True(t, transaction.IsFailed())
	assert.False(t, transaction.IsSuccessful())
	assert.False(t, transaction.IsPending())

	// Teardown
	server.Close()
}

func TestRefundService_StatusWithError(t *testing.T) {
	// Setup
	t.Parallel()

	// Arrange
	requests := make([]http.Request, 0)
	responses := [][]byte{stubs.TokenResponse(), []byte("Transactions not found")}
	server := helpers.MakeRequestCapturingTestServer([]int{http.StatusOK, http.StatusOK}, responses, &requests)
	client := New(
		WithTokenURL(server.URL),
		WithAPIURL(server.URL),
		WithClientID(testClientID),
		WithClientSecret(testClientSecret),
	)

	// Act
	_, response, err := client.Refund.Status(context.Background(), "ddd")

	// Assert
	assert.NotNil(t, err)
	assert.Equal(t, http.StatusOK, response.HTTPResponse.StatusCode)

	// Teardown
	server.Close()
}

func TestRefundService_StatusWithMaxTransactionExceeded(t *testing.T) {
	// Setup
	t.Parallel()

	// Arrange
	requests := make([]http.Request, 0)
	responses := [][]byte{stubs.TokenResponse(), stubs.RefundStatusWithMaxRetryExceeded()}
	server := helpers.MakeRequestCapturingTestServer([]int{http.StatusOK, http.StatusOK}, responses, &requests)
	client := New(
		WithTokenURL(server.URL),
		WithAPIURL(server.URL),
		WithClientID(testClientID),
		WithClientSecret(testClientSecret),
	)

	// Act
	transaction, response, err := client.Refund.Status(context.Background(), "")

	// Assert
	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, response.HTTPResponse.StatusCode)
	assert.True(t, transaction.IsFailed())
	assert.False(t, transaction.IsSuccessful())
	assert.False(t, transaction.IsPending())

	// Teardown
	server.Close()
}
