package authentication

import (
	"context"
	"fmt"

	abs "github.com/microsoft/kiota-abstractions-go"
)

// TODO(kfcampbell): should these constants be centralized somewhere?
const AuthorizationKey = "Authorization"
const AuthType = "token"
const UserAgentKey = "User-Agent"

// TODO(kfcampbell): get the version and binary name from build settings rather than hard-coding
const UserAgentValue = "go-sdk@v0.0.0"

const APIVersionKey = "X-GitHub-Api-Version"

// TODO(kfcampbell): get the version from the generated code somehow
const APIVersionValue = "2022-11-28"

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
