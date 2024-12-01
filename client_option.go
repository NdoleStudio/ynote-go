package ynote

import (
	"net/http"
	"strings"
)

// Option is options for constructing a client
type Option interface {
	apply(config *clientConfig)
}

type clientOptionFunc func(config *clientConfig)

func (fn clientOptionFunc) apply(config *clientConfig) {
	fn(config)
}

// WithHTTPClient sets the underlying HTTP client used for API requests.
// By default, http.DefaultClient is used.
func WithHTTPClient(httpClient *http.Client) Option {
	return clientOptionFunc(func(config *clientConfig) {
		if httpClient != nil {
			config.httpClient = httpClient
		}
	})
}

// WithTokenURL set's the token URL for the Y-Note API
func WithTokenURL(tokenURL string) Option {
	return clientOptionFunc(func(config *clientConfig) {
		if tokenURL != "" {
			config.tokenURL = strings.TrimRight(tokenURL, "/")
		}
	})
}

// WithAPIURL set's the api URL for the Y-Note API
func WithAPIURL(apiURL string) Option {
	return clientOptionFunc(func(config *clientConfig) {
		if apiURL != "" {
			config.apiURL = strings.TrimRight(apiURL, "/")
		}
	})
}

// WithClientID sets the Y-Note API clientID used to fetch the access token
func WithClientID(clientID string) Option {
	return clientOptionFunc(func(config *clientConfig) {
		config.clientID = clientID
	})
}

// WithClientSecret sets the Y-Note API client secret used to fetch the access token
func WithClientSecret(clientSecret string) Option {
	return clientOptionFunc(func(config *clientConfig) {
		config.clientSecret = clientSecret
	})
}

// WithCustomerKey sets the customer key used to make API requests
func WithCustomerKey(customerKey string) Option {
	return clientOptionFunc(func(config *clientConfig) {
		config.customerKey = customerKey
	})
}

// WithCustomerSecret sets the customer secret used to make API requests
func WithCustomerSecret(customerSecret string) Option {
	return clientOptionFunc(func(config *clientConfig) {
		config.customerSecret = customerSecret
	})
}
