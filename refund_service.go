package ynote

import (
	"context"
	"encoding/json"
	"net/http"
)

// RefundService is the API client for the `/prod/refund` endpoint
type RefundService service

// Status returns the status of an initiated transaction
func (service *RefundService) Status(ctx context.Context, transactionID string) (*RefundTransactionStatus, *Response, error) {
	err := service.client.refreshToken(ctx)
	if err != nil {
		return nil, nil, err
	}

	payload := map[string]any{
		"customerkey":    service.client.customerKey,
		"customersecret": service.client.customerSecret,
	}

	request, err := service.client.newRequest(ctx, http.MethodGet, "/prod/refund/status/"+transactionID, payload)
	if err != nil {
		return nil, nil, err
	}

	response, err := service.client.do(request)
	if err != nil {
		return nil, response, err
	}

	transaction := new(RefundTransactionStatus)
	if err = json.Unmarshal(*response.Body, transaction); err != nil {
		return nil, response, err
	}

	return transaction, response, nil
}

// Refund executes an initiated transaction
func (service *RefundService) Refund(ctx context.Context, params *RefundParams) (*RefundTransaction, *Response, error) {
	err := service.client.refreshToken(ctx)
	if err != nil {
		return nil, nil, err
	}

	feesIncluded := "No"
	if params.FeesIncluded {
		feesIncluded = "Yes"
	}

	payload := map[string]any{
		"customerkey":                  service.client.customerKey,
		"customersecret":               service.client.customerSecret,
		"channelUserMsisdn":            params.ChannelUserMsisdn,
		"pin":                          params.Pin,
		"webhook":                      params.Webhook,
		"amount":                       params.Amount,
		"final_customer_phone":         params.FinalCustomerPhone,
		"final_customer_name":          params.FinalCustomerName,
		"refund_method":                params.RefundMethod,
		"fees_included":                feesIncluded,
		"maximum_retries":              params.MaximumRetries,
		"final_customer_name_accuracy": params.FinalCustomerNameAccuracy,
	}

	request, err := service.client.newRequest(ctx, http.MethodPost, "/prod/refund", payload)
	if err != nil {
		return nil, nil, err
	}

	response, err := service.client.do(request)
	if err != nil {
		return nil, response, err
	}

	transaction := new(RefundTransaction)
	if err = json.Unmarshal(*response.Body, transaction); err != nil {
		return nil, response, err
	}

	return transaction, response, nil
}
