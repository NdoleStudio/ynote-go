package stubs

// RefundInvalidClientResponse is the response when refunding with an invalid client
func RefundInvalidClientResponse() []byte {
	return []byte(`{"error":"invalid_client"}`)
}
