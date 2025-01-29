package stubs

// RefundInvalidClientResponse is the response when refunding with an invalid client
func RefundInvalidClientResponse() []byte {
	return []byte(`{"error":"invalid_client"}`)
}

// RefundResponse is the response when refunding a transaction
func RefundResponse() []byte {
	return []byte(`
{
    "MD5OfMessageBody": "4b55cf6629b5f0ee3c8ac91435a2eb35",
    "MD5OfMessageAttributes": "896f665ac83c778c88943113ee0ccd55",
    "MessageId": "993764f9-6b1f-41bd-a7ca-97b8b2167ed7",
    "ResponseMetadata": {
        "RequestId": "e7b09b6d-0d11-5111-8dc9-c4ab56e3cf7c",
        "HTTPStatusCode": 200,
        "HTTPHeaders": {
            "x-amzn-requestid": "e7b09b6d-0d11-5111-8dc9-c4ab56e3cf7c",
            "date": "Sun, 01 Dec 2024 12:42:26 GMT",
            "content-type": "application/x-amz-json-1.0",
            "content-length": "166",
            "connection": "keep-alive"
        },
        "RetryAttempts": 0
    }
}
`)
}

// RefundStatusResponse is the response when checking the status of a refund transaction
func RefundStatusResponse() []byte {
	return []byte(`
{
    "result": {
        "message": "Cash in performed successfully",
        "data": {
            "createtime": "1733056973",
            "subscriberMsisdn": "695xxxxxx",
            "amount": 98,
            "payToken": "CI24120168FBF65A909F588B4480",
            "txnid": "CI241201.1342.C36820",
            "txnmode": "rembourse",
            "txnstatus": "200",
            "orderId": "rembourse",
            "status": "SUCCESSFULL",
            "channelUserMsisdn": "695xxxxxx",
            "description": "Remboursement"
        }
    },
    "parameters": {
        "amount": "98",
        "xauth": "WU5PVEVIRUFExxxxxxxxxxx",
        "channel_user_msisdn": "69xxxxxx",
        "customer_key": "2fBAAq_xxxxxxx",
        "customer_secret": "34nFkKxxxxxx",
        "final_customer_name": "Arnold",
        "final_customer_phone": "69xxxxxx",
        "final_customer_name_accuracy": "0"
    },
    "CreateAt": "12-01-2024 12:43:00",
    "MessageId": "993764f9-6b1f-41bd-a7ca-97b8b2167ed7",
    "RefundStep": "2"
}
`)
}

// RefundStatusResponseWithFailure represents a refund status response with a failure
func RefundStatusResponseWithFailure() []byte {
	return []byte(`
{
   "ErrorCode":5019,
   "body":"Exception on getToken",
   "ErrorMessage":"[Errno 97] Address family not supported by protocol",
   "status":"FAILED",
   "parameters":{
      "amount":"10182",
      "xauth":"WU5PVEVIRUFEOxxxxxxxxxxxxxx",
      "channel_user_msisdn":"693xxxxxx",
      "customer_key":"2fBAAq_xxxxxxxxxxxxxxxxxxxx",
      "customer_secret":"34nFxxxxxxxxxxxxxxxxxxxxx",
      "final_customer_name":"Arnold",
      "final_customer_phone":"693xxxxxx"
   },
   "CreateAt":"12-20-2024 12:14:38",
   "MessageId":"90be9142-4216-4113-bdc4-d1d11ede2460",
   "RefundStep":"1"
}
`)
}

// RefundStatusWithMaxRetryExceeded represents a refund status response with a max retry exceeded
func RefundStatusWithMaxRetryExceeded() []byte {
	return []byte(`
{
    "CreateAt": "01-29-2025 20:20:08",
    "RefundStep": "1",
    "result": {
        "data": {
            "amount": "",
            "channelUserMsisdn": "",
            "createtime": "",
            "description": "Max retry exceeded",
            "orderId": "",
            "payToken": "",
            "status": "FAILED",
            "subscriberMsisdn": "",
            "txnid": "",
            "txnmessage": "Max retry exceeded",
            "txnmode": "",
            "txnstatus": ""
        },
        "message": "Transaction retrieved successfully"
    }
}
`)
}
