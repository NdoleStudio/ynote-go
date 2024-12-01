package ynote

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWithHTTPClient(t *testing.T) {
	t.Run("httpClient is not set when the httpClient is nil", func(t *testing.T) {
		// Setup
		t.Parallel()

		// Arrange
		config := defaultClientConfig()

		// Act
		WithHTTPClient(nil).apply(config)

		// Assert
		assert.NotNil(t, config.httpClient)
	})

	t.Run("httpClient is set when the httpClient is not nil", func(t *testing.T) {
		// Setup
		t.Parallel()

		// Arrange
		config := defaultClientConfig()
		newClient := &http.Client{Timeout: 300}

		// Act
		WithHTTPClient(newClient).apply(config)

		// Assert
		assert.NotNil(t, config.httpClient)
		assert.Equal(t, newClient.Timeout, config.httpClient.Timeout)
	})
}

func TestWithApiURL(t *testing.T) {
	t.Run("apiURL is set successfully", func(t *testing.T) {
		// Setup
		t.Parallel()

		// Arrange
		apiURL := "https://example.com"
		config := defaultClientConfig()

		// Act
		WithApiURL(apiURL).apply(config)

		// Assert
		assert.Equal(t, apiURL, config.apiURL)
	})

	t.Run("tailing / is trimmed from apiURL", func(t *testing.T) {
		// Setup
		t.Parallel()

		// Arrange
		apiURL := "https://example.com/"
		config := defaultClientConfig()

		// Act
		WithApiURL(apiURL).apply(config)

		// Assert
		assert.Equal(t, "https://example.com", config.apiURL)
	})
}

func TestWithClientID(t *testing.T) {
	t.Run("clientID is set successfully", func(t *testing.T) {
		// Setup
		t.Parallel()

		// Arrange
		username := "client-id"
		config := defaultClientConfig()

		// Act
		WithClientID(username).apply(config)

		// Assert
		assert.Equal(t, username, config.clientID)
	})
}

func TestWithClientSecret(t *testing.T) {
	t.Run("clientSecret is set successfully", func(t *testing.T) {
		// Setup
		t.Parallel()

		// Arrange
		password := "client-secret"
		config := defaultClientConfig()

		// Act
		WithClientSecret(password).apply(config)

		// Assert
		assert.Equal(t, password, config.clientSecret)
	})
}

func TestWithCustomerKey(t *testing.T) {
	t.Run("customerKey is set successfully", func(t *testing.T) {
		// Setup
		t.Parallel()

		// Arrange
		customerKey := "key-1"
		config := defaultClientConfig()

		// Act
		WithCustomerKey(customerKey).apply(config)

		// Assert
		assert.Equal(t, customerKey, config.customerKey)
	})
}

func TestWith(t *testing.T) {
	t.Run("customerSecret is set successfully", func(t *testing.T) {
		// Setup
		t.Parallel()

		// Arrange
		customerSecret := "secret-1"
		config := defaultClientConfig()

		// Act
		WithCustomerSecret(customerSecret).apply(config)

		// Assert
		assert.Equal(t, customerSecret, config.customerSecret)
	})
}
