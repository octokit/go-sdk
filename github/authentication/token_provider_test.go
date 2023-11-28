package authentication_test

import (
	"context"
	"fmt"
	"log"
	"os"
	"testing"

	abstractions "github.com/microsoft/kiota-abstractions-go"
	http "github.com/microsoft/kiota-http-go"
	"github.com/octokit/go-sdk/github/authentication"
	"github.com/octokit/go-sdk/github/octokit"
	"github.com/octokit/go-sdk/github/octokit/user"
)

// TODO(kfcampbell): additional test cases
// - error on instantiation when an empty token is given
// - validate existing authorization header does _not_ get overwritten

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
	expectedToken := fmt.Sprintf("token %v", token)
	receivedToken := reqInfo.Headers.Get(authentication.HeaderKey)
	if expectedToken != receivedToken[0] {
		t.Errorf("received token doesn't match up with given token")
	}
}

func TestDoesItReallyWork(t *testing.T) {
	token := os.Getenv("GITHUB_TOKEN")
	if token == "" {
		t.Errorf("GITHUB_TOKEN must exist")
	}

	provider, err := authentication.NewTokenProvider(token)
	if err != nil {
		t.Error("blah blah blah")
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
