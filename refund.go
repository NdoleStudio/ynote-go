package ynote

// RefundParams are the parameters for executing a refund transaction
type RefundParams struct {
	ChannelUserMsisdn         string `json:"channelUserMsisdn"`
	Pin                       string `json:"pin"`
	Webhook                   string `json:"webhook"`
	Amount                    string `json:"amount"`
	FinalCustomerPhone        string `json:"final_customer_phone"`
	FinalCustomerName         string `json:"final_customer_name"`
	RefundMethod              string `json:"refund_method"`
	FeesIncluded              bool   `json:"fees_included"`
	DebitPolicy               string `json:"debit_policy"`
	FinalCustomerNameAccuracy string `json:"final_customer_name_accuracy"`
}
