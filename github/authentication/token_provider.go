package authentication

import (
	"context"
	"fmt"

	abs "github.com/microsoft/kiota-abstractions-go"
)

type TokenProvider struct {
	token   string
	options []TokenProviderOption
}

type TokenProviderOption func(*TokenProvider, *Request)

// WithDefaultUserAgent sets the User-Agent string sent for requests to the default
// for this SDK.
func WithDefaultUserAgent() TokenProviderOption {
	return func(t *TokenProvider, r *Request) {
		r.WithDefaultUserAgent()
	}
}

// WithUserAgent sets the User-Agent string sent with each request.
func WithUserAgent(userAgent string) TokenProviderOption {
	return func(t *TokenProvider, r *Request) {
		r.WithUserAgent(userAgent)
	}
}

// WithDefaultAPIVersion sets the API version header sent with each request.
func WithDefaultAPIVersion() TokenProviderOption {
	return func(t *TokenProvider, r *Request) {
		r.WithDefaultAPIVersion()
	}
}

// WithAPIVersion sets the API version header sent with each request.
func WithAPIVersion(version string) TokenProviderOption {
	return func(t *TokenProvider, r *Request) {
		r.WithAPIVersion(version)
	}
}

// TODO(kfcampbell): implement new constructor with allowedHosts

// NewTokenProvider creates an instance of TokenProvider with the specified token and options.
// An error is returned when the token present is an empty string.
func NewTokenProvider(apiToken string, options ...TokenProviderOption) (*TokenProvider, error) {
	if apiToken == "" {
		return nil, fmt.Errorf("API token must not be an empty string")
	}

	provider := &TokenProvider{
		token:   apiToken,
		options: options,
	}

	return provider, nil
}

var defaultHandlers = []TokenProviderOption{WithDefaultUserAgent(), WithDefaultAPIVersion()}

// AuthenticateRequest applies the default options for each request, then the user's options
// if present in the TokenProvider.
func (t *TokenProvider) AuthenticateRequest(context context.Context, request *abs.RequestInformation, additionalAuthenticationContext map[string]interface{}) error {
	reqWrapper := &Request{RequestInformation: request}

	if reqWrapper.Headers == nil {
		reqWrapper.Headers = abs.NewRequestHeaders()
	}

	for _, option := range defaultHandlers {
		option(t, reqWrapper)
	}

	// apply user options after defaults
	for _, option := range t.options {
		option(t, reqWrapper)
	}

	return nil
}
