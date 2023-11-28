package authentication

import (
	"context"
	"fmt"

	abs "github.com/microsoft/kiota-abstractions-go"
)

// TODO(kfcampbell): should these constants be centralized somewhere?
const HeaderKey = "Authorization"

// TODO(kfcampbell): implement user-agent string setting, API versioning
// TODO(kfcampbell): do we want to implement some sort of request handler functionality here?
type TokenProvider struct {
	token string
}

// TODO(kfcampbell): implement new constructor with allowedHosts
func NewTokenProvider(apiToken string) (*TokenProvider, error) {
	if apiToken == "" {
		return nil, fmt.Errorf("API token must not be an empty string")
	}

	provider := &TokenProvider{
		token: apiToken,
	}

	return provider, nil
}

func (t *TokenProvider) AuthenticateRequest(context context.Context, request *abs.RequestInformation, additionalAuthenticationContext map[string]interface{}) error {
	if request.Headers == nil {
		request.Headers = &abs.RequestHeaders{}
	}

	if len(request.Headers.Get(HeaderKey)) == 0 {
		request.Headers.Add(HeaderKey, fmt.Sprintf("token %v", t.token))
	}

	return nil
}
