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

func WithDefaultUserAgent() TokenProviderOption {
	return func(t *TokenProvider, r *Request) {
		r.WithDefaultUserAgent()
	}
}

func WithUserAgent(userAgent string) TokenProviderOption {
	return func(t *TokenProvider, r *Request) {
		r.WithUserAgent(userAgent)
	}
}

// TODO(kfcampbell): implement new constructor with allowedHosts
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

func (t *TokenProvider) AuthenticateRequest(context context.Context, request *abs.RequestInformation, additionalAuthenticationContext map[string]interface{}) error {
	reqWrapper := &Request{RequestInformation: request}

	for _, option := range t.options {
		option(t, reqWrapper)
	}

	if request.Headers == nil {
		request.Headers = abs.NewRequestHeaders()
	}

	// TODO(kfcampbell): do we want to implement some sort of request handler functionality here?
	// perhaps a functional pattern would be better than this chained if approach
	if !request.Headers.ContainsKey(AuthorizationKey) {
		request.Headers.Add(AuthorizationKey, fmt.Sprintf("%v %v", AuthType, t.token))
	}

	if !request.Headers.ContainsKey(UserAgentKey) {
		request.Headers.Add(UserAgentKey, UserAgentValue)
	}

	if !request.Headers.ContainsKey(APIVersionKey) {
		request.Headers.Add(APIVersionKey, APIVersionValue)
	}

	return nil
}
