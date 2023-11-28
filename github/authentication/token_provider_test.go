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
	"github.com/octokit/go-sdk/github/authentication"
	"github.com/octokit/go-sdk/github/octokit"
	"github.com/octokit/go-sdk/github/octokit/user"
)

func TestTokenIsSetInAuthenticatedRequest(t *testing.T) {
	token := "help i'm trapped in a Go binary"
	provider, err := authentication.NewTokenProvider(token)
	if err != nil {
		t.Errorf("should be no error when instantiating provider")
	}

	reqInfo := &abstractions.RequestInformation{}
	addtlContext := make(map[string]interface{})

	err = provider.AuthenticateRequest(context.Background(), reqInfo, addtlContext)
	if err != nil {
		t.Errorf("there should be no error when calling AuthenticateRequest")
	}

	if len(reqInfo.Headers.Get("Authorization")) != 1 {
		t.Errorf("exactly one authorization header should be set")
	}
	receivedToken := reqInfo.Headers.Get(authentication.HeaderKey)[0]
	if !strings.Contains(receivedToken, token) {
		t.Errorf("received token doesn't match up with given token")
	}
}

func TestNewTokenProviderErrorsWithBlankToken(t *testing.T) {
	_, err := authentication.NewTokenProvider("")
	if err == nil {
		t.Errorf("NewTokenProvider should error when a blank token is not given")
	}
}

func TestTokenSetInRequestIsNotOverwritten(t *testing.T) {
	providerToken := "dit dit dit / dat dat dat / dit dit dit"
	provider, err := authentication.NewTokenProvider(providerToken)
	if err != nil {
		t.Errorf("should be no error when instantiating provider")
	}

	requestToken := "dit dit dit dit / dit / dit dat dit dit / dit dat dat dit"
	headers := abstractions.NewRequestHeaders()
	headers.Add(authentication.AuthType, requestToken)

	reqInfo := abstractions.NewRequestInformation()
	reqInfo.Headers = headers
	addtlContext := make(map[string]interface{})

	err = provider.AuthenticateRequest(context.Background(), reqInfo, addtlContext)
	if err != nil {
		t.Errorf("AuthenticateRequest should not error")
	}
	reqInfoToken := reqInfo.Headers.Get(authentication.HeaderKey)[0]

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

	provider, err := authentication.NewTokenProvider(token)
	if err != nil {
		t.Error("instantiating TokenProvider should not error")
	}

	adapter, err := http.NewNetHttpRequestAdapter(provider)
	if err != nil {
		log.Fatalf("Error creating request adapter: %v", err)
	}
	headers := abstractions.NewRequestHeaders()
	_ = headers.TryAdd("Accept", "application/vnd.github.v3+json")

	client := octokit.NewApiClient(adapter)
	emailsRequestConfig := &user.EmailsRequestBuilderGetRequestConfiguration{
		Headers: headers,
	}
	userEmails, err := client.User().Emails().Get(context.Background(), emailsRequestConfig)
	if err != nil {
		log.Fatalf("%v\n", err)
	}

	for _, v := range userEmails {
		fmt.Printf("%v\n", *v.GetEmail())
	}
}
