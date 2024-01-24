package authentication_test

import (
	"context"
	"fmt"
	"log"
	"os"
	"strings"
	"testing"

	abstractions "github.com/microsoft/kiota-abstractions-go"
	http "github.com/microsoft/kiota-http-go"
	"github.com/octokit/go-sdk/pkg/authentication"
	"github.com/octokit/go-sdk/pkg/github"
	"github.com/octokit/go-sdk/pkg/github/user"
	"github.com/octokit/go-sdk/pkg/headers"
)

func TestTokenIsSetInAuthenticatedRequest(t *testing.T) {
	token := "help i'm trapped in a Go binary"
	provider := authentication.NewTokenProvider(authentication.WithAuthorizationToken(token))

	reqInfo := abstractions.NewRequestInformation()
	addtlContext := make(map[string]interface{})

	err := provider.AuthenticateRequest(context.Background(), reqInfo, addtlContext)
	if err != nil {
		t.Errorf("there should be no error when calling AuthenticateRequest")
	}

	if len(reqInfo.Headers.Get(headers.AuthorizationKey)) != 1 {
		t.Errorf("there should be exactly one authorization key")
	}

	receivedToken := reqInfo.Headers.Get(headers.AuthorizationKey)[0]
	if !strings.Contains(receivedToken, token) {
		t.Errorf("received token doesn't match up with given token")
	}
}

// TODO(kfcampbell): this code could be refactored to use table-based tests
func TestDefaultRequestOptions(t *testing.T) {
	token := "this is not the token you're looking for"
	provider := authentication.NewTokenProvider(authentication.WithAuthorizationToken(token))
	reqInfo := abstractions.NewRequestInformation()
	addtlContext := make(map[string]interface{})

	err := provider.AuthenticateRequest(context.Background(), reqInfo, addtlContext)
	if err != nil {
		t.Errorf("there should be no error when calling AuthenticateRequest")
	}

	apiVersions := reqInfo.Headers.Get(headers.APIVersionKey)
	if len(apiVersions) != 1 {
		t.Errorf("exactly one API version should be present in the request")
	}

	if apiVersions[0] != headers.APIVersionValue {
		t.Errorf("default API version is set incorrectly")
	}

	userAgents := reqInfo.Headers.Get(headers.UserAgentKey)
	if len(userAgents) != 1 {
		t.Errorf("exactly one user agent string should be present in the request")
	}

	if userAgents[0] != headers.UserAgentValue {
		t.Errorf("default user agent string is set incorrectly")
	}
}

func TestOverwritingDefaultRequestOptions(t *testing.T) {
	token := "i'm totally a real token"
	apiVersion := "i'm totally a real API version"
	userAgent := "i'm totally a real user agent"
	provider := authentication.NewTokenProvider(
		authentication.WithAuthorizationToken(token),
		authentication.WithAPIVersion(apiVersion),
		authentication.WithUserAgent(userAgent))

	reqInfo := abstractions.NewRequestInformation()
	addtlContext := make(map[string]interface{})

	err := provider.AuthenticateRequest(context.Background(), reqInfo, addtlContext)
	if err != nil {
		t.Errorf("should be no error when calling authenticated request")
	}

	apiVersions := reqInfo.Headers.Get(headers.APIVersionKey)
	if len(apiVersions) != 1 {
		t.Errorf("exactly one API version should be present in the request")
	}

	if apiVersions[0] != apiVersion {
		t.Errorf("default API version is set incorrectly")
	}

	userAgents := reqInfo.Headers.Get(headers.UserAgentKey)
	if len(userAgents) != 1 {
		t.Errorf("exactly one user agent string should be present in the request")
	}

	if userAgents[0] != userAgent {
		t.Errorf("default user agent string is set incorrectly")
	}

}

func TestAnonymousAuthIsAllowed(t *testing.T) {
	provider := authentication.NewTokenProvider()
	reqInfo := abstractions.NewRequestInformation()
	addtlContext := make(map[string]interface{})

	err := provider.AuthenticateRequest(context.Background(), reqInfo, addtlContext)
	if err != nil {
		t.Errorf("should be no error when calling authenticated request")
	}

	authorizations := reqInfo.Headers.Get(headers.AuthorizationKey)
	if len(authorizations) != 0 {
		t.Errorf("no authorization header should be present in the request")
	}
}

func TestTokenSetInRequestIsNotOverwritten(t *testing.T) {
	providerToken := "dit dit dit / dat dat dat / dit dit dit"
	provider := authentication.NewTokenProvider(
		authentication.WithAuthorizationToken(providerToken),
	)

	requestToken := "dit dit dit dit / dit / dit dat dit dit / dit dat dat dit"
	requestHeaders := abstractions.NewRequestHeaders()
	requestHeaders.Add(headers.AuthType, requestToken)

	reqInfo := abstractions.NewRequestInformation()
	reqInfo.Headers = requestHeaders
	addtlContext := make(map[string]interface{})

	err := provider.AuthenticateRequest(context.Background(), reqInfo, addtlContext)
	if err != nil {
		t.Errorf("AuthenticateRequest should not error")
	}
	reqInfoToken := reqInfo.Headers.Get(headers.AuthorizationKey)[0]

	if !strings.Contains(reqInfoToken, providerToken) {
		t.Errorf("received token doesn't match up with given token")
	}
}

// TODO(kfcampbell): make a more permanent decision about how to structure
// and separately run unit vs. integration tests
func TestHappyPathIntegration(t *testing.T) {
	token := os.Getenv("GITHUB_TOKEN")
	if token == "" {
		t.Skip("in order to run integration tests, ensure a valid GITHUB_TOKEN exists in the environment")
	}

	provider := authentication.NewTokenProvider(
		authentication.WithAuthorizationToken(token),
	)

	adapter, err := http.NewNetHttpRequestAdapter(provider)
	if err != nil {
		log.Fatalf("Error creating request adapter: %v", err)
	}
	headers := abstractions.NewRequestHeaders()
	_ = headers.TryAdd("Accept", "application/vnd.github.v3+json")

	client := github.NewApiClient(adapter)

	// Create a new instance of abstractions.RequestConfiguration
	requestConfig := &abstractions.RequestConfiguration[user.EmailsRequestBuilderGetQueryParameters]{
	Headers: headers,
	}



	userEmails, err := client.User().Emails().Get(context.Background(), requestConfig)
	if err != nil {
	log.Fatalf("%v\n", err)
	}

	for _, v := range userEmails {
	fmt.Printf("%v\n", *v.GetEmail())
	}
}
