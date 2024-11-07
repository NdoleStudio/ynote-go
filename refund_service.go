package ynote

import (
	"context"
	"encoding/json"
	"net/http"
)

// RefundService is the API client for the `/prod/refund` endpoint
type RefundService service

//// Status returns the status of an initiated transaction
//func (service *RefundService) Status(ctx context.Context, payToken *string) (*OrangeResponse[MerchantPaymentTransaction], *Response, error) {
//	err := service.client.refreshToken(ctx)
//	if err != nil {
//		return nil, nil, err
//	}
//
//	request, err := service.client.newRequest(ctx, http.MethodGet, "/omcoreapis/1.0.2/mp/paymentstatus/"+*payToken, nil)
//	if err != nil {
//		return nil, nil, err
//	}
//
//	response, err := service.client.do(request)
//	if err != nil {
//		return nil, response, err
//	}
//
//	transaction := new(OrangeResponse[MerchantPaymentTransaction])
//	if err = json.Unmarshal(*response.Body, transaction); err != nil {
//		return nil, response, err
//	}
//
//	return transaction, response, nil
//}

// Refund executes an initiated transaction
func (service *RefundService) Refund(ctx context.Context, params *RefundParams) (*map[string]any, *Response, error) {
	err := service.client.refreshToken(ctx)
	if err != nil {
		return nil, nil, err
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
		"fees_included":                params.FeesIncluded,
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

	transaction := new(map[string]any)
	if err = json.Unmarshal(*response.Body, transaction); err != nil {
		return nil, response, err
	}

	return transaction, response, nil
}
