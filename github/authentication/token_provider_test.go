package authentication_test

import (
	"context"
	"testing"

	abstractions "github.com/microsoft/kiota-abstractions-go"
	"github.com/octokit/go-sdk/github/authentication"
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
	receivedToken := reqInfo.Headers.Get(authentication.HeaderKey)
	if token != receivedToken[0] {
		t.Errorf("received token doesn't match up with given token")
	}
}
