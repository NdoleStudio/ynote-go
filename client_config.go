package ynote

import "net/http"

type clientConfig struct {
	httpClient     *http.Client
	customerKey    string
	customerSecret string
	clientID       string
	clientSecret   string
	tokenURL       string
	apiURL         string
}

func defaultClientConfig() *clientConfig {
	return &clientConfig{
		httpClient: http.DefaultClient,
		tokenURL:   "https://omapi-token.ynote.africa",
		apiURL:     "https://omapi.ynote.africa",
	}
}
