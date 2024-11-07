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

// WithApiURL set's the api URL for the Y-Note API
func WithApiURL(apiURL string) Option {
	return clientOptionFunc(func(config *clientConfig) {
		if apiURL != "" {
			config.apiURL = strings.TrimRight(apiURL, "/")
		}
	})
}

// WithUsername sets the Y-Note API Username used to fetch the access token
func WithUsername(username string) Option {
	return clientOptionFunc(func(config *clientConfig) {
		config.username = username
	})
}

// WithPassword sets the Y-Note API password used to fetch the access token
func WithPassword(password string) Option {
	return clientOptionFunc(func(config *clientConfig) {
		config.password = password
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
