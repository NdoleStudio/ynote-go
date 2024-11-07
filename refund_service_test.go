package ynote

import (
	"context"
	"github.com/NdoleStudio/ynote-go/internal/helpers"
	"github.com/NdoleStudio/ynote-go/internal/stubs"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestRefundService_Refund(t *testing.T) {
	// Setup
	t.Parallel()

	// Arrange
	requests := make([]http.Request, 0)
	responses := [][]byte{stubs.TokenResponse(), stubs.RefundInvalidClientResponse()}
	server := helpers.MakeRequestCapturingTestServer([]int{http.StatusOK, http.StatusOK}, responses, &requests)
	client := New(
		WithTokenURL(server.URL),
		WithApiURL(server.URL),
		WithUsername(testUsername),
		WithPassword(testPassword),
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
	_, response, err := client.Refund.Refund(context.Background(), payload)

	// Assert
	assert.Nil(t, err)

	assert.GreaterOrEqual(t, len(requests), 2)
	request := requests[len(requests)-1]

	assert.Equal(t, "/prod/refund", request.URL.Path)
	assert.Equal(t, "Bearer 19077204-9d0a-31fa-85cf-xxxxxxxxxx", request.Header.Get("Authorization"))

	assert.Equal(t, http.StatusOK, response.HTTPResponse.StatusCode)

	//Teardown
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
		WithApiURL(server.URL),
		WithUsername(testUsername),
		WithPassword(testPassword),
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

	//Teardown
	server.Close()
}

//func TestMerchantPaymentService_Pay(t *testing.T) {
//	// Setup
//	t.Parallel()
//
//	// Arrange
//	requests := make([]http.Request, 0)
//	responses := [][]byte{stubs.TokenResponse(), stubs.MerchantPaymentPayResponse()}
//	server := helpers.MakeRequestCapturingTestServer([]int{http.StatusOK, http.StatusOK}, responses, &requests)
//	client := New(
//		WithBaseURL(server.URL),
//		WithUsername(testUsername),
//		WithPassword(testPassword),
//		WithAuthToken(testAuthToken),
//	)
//
//	// Act
//	transaction, response, err := client.MerchantPayment.Pay(context.Background(), &MerchantPaymentPayPrams{
//		SubscriberMSISDN:  "69XXXXXXX",
//		ChannelUserMSISDN: "69XXXXXXX",
//		Amount:            "100",
//		Description:       "Payment Description",
//		OrderID:           "abcdef",
//		Pin:               "123456",
//		PayToken:          "MP22120771FEB7B21FD2381C3786",
//		NotificationURL:   "https://example.com/payment-notification",
//	})
//
//	// Assert
//	assert.Nil(t, err)
//
//	assert.GreaterOrEqual(t, len(requests), 2)
//	request := requests[len(requests)-1]
//
//	assert.Equal(t, "/omcoreapis/1.0.2/mp/pay", request.URL.Path)
//	assert.Equal(t, testAuthToken, request.Header.Get("X-AUTH-TOKEN"))
//	assert.Equal(t, "Bearer 19077204-9d0a-31fa-85cf-xxxxxxxxxx", request.Header.Get("Authorization"))
//
//	assert.Equal(t, http.StatusOK, response.HTTPResponse.StatusCode)
//
//	assert.Equal(t, &OrangeResponse[MerchantPaymentTransaction]{
//		Message: "Merchant payment successfully initiated",
//		Data: MerchantPaymentTransaction{
//			ID:                        48463325,
//			CreatedTime:               "1670442691",
//			SubscriberMSISDN:          "69XXXXXXX",
//			Amount:                    100,
//			PayToken:                  "MP22120771FEB7B21FD2381C3786",
//			TransactionID:             "MP221207.2051.B56929",
//			TransactionMode:           "12345",
//			InitTransactionMessage:    "Paiement e la clientele done.The devrez confirmer le paiement en saisissant son code PIN et vous recevrez alors un SMS. Merci dutiliser des services Orange Money.",
//			InitTransactionStatus:     "200",
//			ConfirmTransactionStatus:  nil,
//			ConfirmTransactionMessage: nil,
//			Status:                    "PENDING",
//			NotificationURL:           "https://example.com/payment-notification",
//			Description:               "Payment Description",
//			ChannelUserMSISDN:         "69XXXXXXX",
//		},
//	}, transaction)
//
//	assert.True(t, transaction.Data.IsPending())
//	assert.False(t, transaction.Data.IsConfirmed())
//	assert.False(t, transaction.Data.IsExpired())
//
//	// Teardown
//	server.Close()
//}
//
//func TestMerchantPaymentService_PayWithInsufficientFunds(t *testing.T) {
//	// Setup
//	t.Parallel()
//
//	// Arrange
//	requests := make([]http.Request, 0)
//	responses := [][]byte{stubs.TokenResponse(), stubs.MerchantPaymentPayResponseWithInsufficientFunds()}
//	server := helpers.MakeRequestCapturingTestServer([]int{http.StatusOK, http.StatusExpectationFailed}, responses, &requests)
//	client := New(
//		WithBaseURL(server.URL),
//		WithUsername(testUsername),
//		WithPassword(testPassword),
//		WithAuthToken(testAuthToken),
//	)
//
//	// Act
//	_, response, err := client.MerchantPayment.Pay(context.Background(), &MerchantPaymentPayPrams{
//		SubscriberMSISDN:  "69XXXXXXX",
//		ChannelUserMSISDN: "69XXXXXXX",
//		Amount:            "100",
//		Description:       "Payment Description",
//		OrderID:           "abcdef",
//		Pin:               "123456",
//		PayToken:          "MP22120771FEB7B21FD2381C3786",
//		NotificationURL:   "https://example.com/payment-notification",
//	})
//
//	// Assert
//	assert.NotNil(t, err)
//	assert.Equal(t, http.StatusExpectationFailed, response.HTTPResponse.StatusCode)
//	assert.True(t, strings.Contains(string(*response.Body), "60019 :: Le solde du compte du payeur est insuffisant"))
//
//	// Teardown
//	server.Close()
//}
//
//func TestMerchantPaymentService_Push(t *testing.T) {
//	// Setup
//	t.Parallel()
//
//	// Arrange
//	requests := make([]http.Request, 0)
//	responses := [][]byte{stubs.TokenResponse(), stubs.MerchantPaymentPushResponse()}
//	server := helpers.MakeRequestCapturingTestServer([]int{http.StatusOK, http.StatusOK}, responses, &requests)
//	client := New(
//		WithBaseURL(server.URL),
//		WithUsername(testUsername),
//		WithPassword(testPassword),
//		WithAuthToken(testAuthToken),
//	)
//	payToken := "MP22120771FEB7B21FD2381C3786"
//
//	// Act
//	transaction, response, err := client.MerchantPayment.Push(context.Background(), &payToken)
//
//	// Assert
//	assert.Nil(t, err)
//
//	assert.GreaterOrEqual(t, len(requests), 2)
//	request := requests[len(requests)-1]
//
//	assert.Equal(t, "/omcoreapis/1.0.2/mp/push/"+payToken, request.URL.Path)
//	assert.Equal(t, testAuthToken, request.Header.Get("X-AUTH-TOKEN"))
//	assert.Equal(t, "Bearer 19077204-9d0a-31fa-85cf-xxxxxxxxxx", request.Header.Get("Authorization"))
//
//	assert.Equal(t, http.StatusOK, response.HTTPResponse.StatusCode)
//
//	assert.Equal(t, &OrangeResponse[MerchantPaymentTransaction]{
//		Message: "Push sent to customer",
//		Data: MerchantPaymentTransaction{
//			ID:                        48463325,
//			CreatedTime:               "1670442691",
//			SubscriberMSISDN:          "69XXXXXXX",
//			Amount:                    100,
//			PayToken:                  "MP22120771FEB7B21FD2381C3786",
//			TransactionID:             "MP221207.2051.B56929",
//			TransactionMode:           "12345",
//			InitTransactionMessage:    "Paiement e la clientele done.The devrez confirmer le paiement en saisissant son code PIN et vous recevrez alors un SMS. Merci dutiliser des services Orange Money.",
//			InitTransactionStatus:     "200",
//			ConfirmTransactionStatus:  nil,
//			ConfirmTransactionMessage: nil,
//			Status:                    "PENDING",
//			NotificationURL:           "https://example.com/payment-notification",
//			Description:               "Payment Description",
//			ChannelUserMSISDN:         "69XXXXXXX",
//		},
//	}, transaction)
//
//	assert.True(t, transaction.Data.IsPending())
//	assert.False(t, transaction.Data.IsConfirmed())
//	assert.False(t, transaction.Data.IsExpired())
//
//	// Teardown
//	server.Close()
//}
//
//func TestMerchantPaymentService_TransactionStatus(t *testing.T) {
//	// Setup
//	t.Parallel()
//
//	// Arrange
//	requests := make([]http.Request, 0)
//	responses := [][]byte{stubs.TokenResponse(), stubs.MerchantPaymentTransactionStatusResponse()}
//	server := helpers.MakeRequestCapturingTestServer([]int{http.StatusOK, http.StatusOK}, responses, &requests)
//	client := New(
//		WithBaseURL(server.URL),
//		WithUsername(testUsername),
//		WithPassword(testPassword),
//		WithAuthToken(testAuthToken),
//	)
//	payToken := "MP22120771FEB7B21FD2381C3786"
//
//	// Act
//	transaction, response, err := client.MerchantPayment.TransactionStatus(context.Background(), &payToken)
//
//	// Assert
//	assert.Nil(t, err)
//
//	assert.GreaterOrEqual(t, len(requests), 2)
//	request := requests[len(requests)-1]
//
//	assert.Equal(t, "/omcoreapis/1.0.2/mp/paymentstatus/"+payToken, request.URL.Path)
//	assert.Equal(t, testAuthToken, request.Header.Get("X-AUTH-TOKEN"))
//	assert.Equal(t, "Bearer 19077204-9d0a-31fa-85cf-xxxxxxxxxx", request.Header.Get("Authorization"))
//
//	assert.Equal(t, http.StatusOK, response.HTTPResponse.StatusCode)
//
//	strPtr := func(val string) *string {
//		return &val
//	}
//
//	assert.Equal(t, &OrangeResponse[MerchantPaymentTransaction]{
//		Message: "Transaction retrieved successfully",
//		Data: MerchantPaymentTransaction{
//			ID:                        48463325,
//			CreatedTime:               "1670442691",
//			SubscriberMSISDN:          "69XXXXXXX",
//			Amount:                    100,
//			PayToken:                  "MP22120771FEB7B21FD2381C3786",
//			TransactionID:             "MP221207.2051.B56929",
//			TransactionMode:           "12345",
//			InitTransactionMessage:    "Paiement e la clientele done.The devrez confirmer le paiement en saisissant son code PIN et vous recevrez alors un SMS. Merci dutiliser des services Orange Money.",
//			InitTransactionStatus:     "200",
//			ConfirmTransactionStatus:  strPtr("200"),
//			ConfirmTransactionMessage: strPtr("Successful Payment of COMPANY_NAME from 69XXXXXXX CUSTOMER_NAME. Transaction ID:MP221207.2051.B56929, Amount:100, New balance:1103.5."),
//			Status:                    "SUCCESSFULL",
//			NotificationURL:           "https://example.com/payment-notification",
//			Description:               "Payment Description",
//			ChannelUserMSISDN:         "69XXXXXXX",
//		},
//	}, transaction)
//
//	assert.False(t, transaction.Data.IsPending())
//	assert.True(t, transaction.Data.IsConfirmed())
//	assert.False(t, transaction.Data.IsExpired())
//
//	// Teardown
//	server.Close()
//}
//
//func TestMerchantPaymentService_TransactionStatusWithExpired(t *testing.T) {
//	// Setup
//	t.Parallel()
//
//	// Arrange
//	requests := make([]http.Request, 0)
//	responses := [][]byte{stubs.TokenResponse(), stubs.MerchantPaymentTransactionStatusResponseWithExpired()}
//	server := helpers.MakeRequestCapturingTestServer([]int{http.StatusOK, http.StatusOK}, responses, &requests)
//	client := New(
//		WithBaseURL(server.URL),
//		WithUsername(testUsername),
//		WithPassword(testPassword),
//		WithAuthToken(testAuthToken),
//	)
//	payToken := "MP22120771FEB7B21FD2381C3786"
//
//	// Act
//	transaction, response, err := client.MerchantPayment.TransactionStatus(context.Background(), &payToken)
//
//	// Assert
//	assert.Nil(t, err)
//
//	assert.GreaterOrEqual(t, len(requests), 2)
//	request := requests[len(requests)-1]
//
//	assert.Equal(t, "/omcoreapis/1.0.2/mp/paymentstatus/"+payToken, request.URL.Path)
//	assert.Equal(t, testAuthToken, request.Header.Get("X-AUTH-TOKEN"))
//	assert.Equal(t, "Bearer 19077204-9d0a-31fa-85cf-xxxxxxxxxx", request.Header.Get("Authorization"))
//
//	assert.Equal(t, http.StatusOK, response.HTTPResponse.StatusCode)
//
//	assert.False(t, transaction.Data.IsPending())
//	assert.False(t, transaction.Data.IsConfirmed())
//	assert.True(t, transaction.Data.IsExpired())
//
//	// Teardown
//	server.Close()
//}
