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
	FinalCustomerNameAccuracy string `json:"final_customer_name_accuracy"`
}

// RefundTransaction is the response from a refund transaction
type RefundTransaction struct {
	MD5OfMessageBody       string `json:"MD5OfMessageBody"`
	MD5OfMessageAttributes string `json:"MD5OfMessageAttributes"`
	MessageID              string `json:"MessageId"`
	ResponseMetadata       struct {
		RequestID      string `json:"RequestId"`
		HTTPStatusCode int    `json:"HTTPStatusCode"`
		HTTPHeaders    struct {
			XAmznRequestid string `json:"x-amzn-requestid"`
			Date           string `json:"date"`
			ContentType    string `json:"content-type"`
			ContentLength  string `json:"content-length"`
			Connection     string `json:"connection"`
		} `json:"HTTPHeaders"`
		RetryAttempts int `json:"RetryAttempts"`
	} `json:"ResponseMetadata"`
}

// RefundTransactionStatus is the response from a refund transaction status
type RefundTransactionStatus struct {
	Status       *string `json:"status"`
	ErrorCode    *int    `json:"ErrorCode"`
	ErrorMessage *string `json:"ErrorMessage"`
	Result       *struct {
		Message string `json:"message"`
		Data    struct {
			CreatedAt         string `json:"createtime"`
			SubscriberMsisdn  string `json:"subscriberMsisdn"`
			Amount            int    `json:"amount"`
			PayToken          string `json:"payToken"`
			TransactionID     string `json:"txnid"`
			TransactionMode   string `json:"txnmode"`
			TransactionStatus string `json:"txnstatus"`
			OrderID           string `json:"orderId"`
			Status            string `json:"status"`
			ChannelUserMsisdn string `json:"channelUserMsisdn"`
			Description       string `json:"description"`
		} `json:"data"`
	} `json:"result"`
	Parameters struct {
		Amount                    string `json:"amount"`
		Xauth                     string `json:"xauth"`
		ChannelUserMsisdn         string `json:"channel_user_msisdn"`
		CustomerKey               string `json:"customer_key"`
		CustomerSecret            string `json:"customer_secret"`
		FinalCustomerName         string `json:"final_customer_name"`
		FinalCustomerPhone        string `json:"final_customer_phone"`
		FinalCustomerNameAccuracy any    `json:"final_customer_name_accuracy"`
	} `json:"parameters"`
	CreatedAt  string `json:"CreateAt"`
	MessageID  string `json:"MessageId"`
	RefundStep string `json:"RefundStep"`
}

// IsPending checks if the refund transaction is pending
func (status *RefundTransactionStatus) IsPending() bool {
	return status.Status == nil && (status.Result == nil || status.Result.Data.Status == "" || status.Result.Data.Status == "PENDING" || status.Result.Data.Status == "INITIATED")
}

// IsSuccessful checks if the refund transaction is successful
func (status *RefundTransactionStatus) IsSuccessful() bool {
	return status.Result != nil && (status.Result.Data.Status == "SUCCESSFULL" || status.Result.Data.Status == "SUCCESSFUL")
}

// IsFailed checks if the refund transaction is failed
func (status *RefundTransactionStatus) IsFailed() bool {
	return (status.Status != nil && (*status.Status == "FAILED" || *status.ErrorCode == 5019)) || (status.Result != nil && (status.Result.Data.Status == "FAILED" || status.Result.Data.Status == "EXPIRED"))
}
