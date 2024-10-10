package pkg

import (
	"context"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	abs "github.com/microsoft/kiota-abstractions-go"
	"github.com/octokit/go-sdk/pkg/github/installation"
	"github.com/octokit/go-sdk/pkg/headers"
)

const (
	installationID = 1234
	clientID       = "testClientID"
	pemFileName    = "testPemFile"
	token          = "testToken"
	appID          = 123
)

// test key taken from Go source code at
// https://go.dev/src/crypto/rsa/rsa_test.go
var key = []byte(`-----BEGIN TESTING KEY-----
MIIEvQIBADANBgkqhkiG9w0BAQEFAASCBKcwggSjAgEAAoIBAQDNoyFUYeDuqw+k
iyv47iBy/udbWmQdpbUZ8JobHv8uQrvL7sQN6l83teHgNJsXqtiLF3MC+K+XI6Dq
hxUWfQwLip8WEnv7Jx/+53S8yp/CS4Jw86Q1bQHbZjFDpcoqSuwAxlegw18HNZCY
fpipYnA1lYCm+MTjtgXJQbjA0dwUGCf4BDMqt+76Jk3XZF5975rftbkGoT9eu8Jt
Xs5F5Xkwd8q3fkQz+fpLW4u9jrfFyQ61RRFkYrCjlhtGjYIzBHGgQM4n/sNXhiy5
h0tA7Xa6NyYrN/OXe/Y1K8Rz/tzlvbMoxgZgtBuKo1N3m8ckFi7hUVK2eNv7GoAb
teTTPrg/AgMBAAECggEAAnfsVpmsL3R0Bh4gXRpPeM63H6e1a8B8kyVwiO9o0cXX
gKp9+P39izfB0Kt6lyCj/Wg+wOQT7rg5qy1yIw7fBHGmcjquxh3uN0s3YZ+Vcym6
SAY5f0vh/OyJN9r3Uv8+Pc4jtb7So7QDzdWeZurssBmUB0avAMRdGNFGP5SyILcz
l3Q59hTxQ4czRHKjZ06L1/sA+tFVbO1j39FN8nMOU/ovLF4lAmZTkQ6AP6n6XPHP
B8Nq7jSYz6RDO200jzp6UsdrnjjkJRbzOxN/fn+ckCP+WYuq+y/d05ET9PdVa4qI
Jyr80D9QgHmfztcecvYwoskGnkb2F4Tmp0WnAj/xVQKBgQD4TrMLyyHdbAr5hoSi
p+r7qBQxnHxPe2FKO7aqagi4iPEHauEDgwPIcsOYota1ACiSs3BaESdJAClbqPYd
HDI4c2DZ6opux6WYkSju+tVXYW6qarR3fzrP3fUCdz2c2NfruWOqq8YmjzAhTNPm
YzvtzTdwheNYV0Vi71t1SfZmfQKBgQDUAgSUcrgXdGDnSbaNe6KwjY5oZWOQfZe2
DUhqfN/JRFZj+EMfIIh6OQXnZqkp0FeRdfRAFl8Yz8ESHEs4j+TikLJEeOdfmYLS
TWxlMPDTUGbUvSf4g358NJ8TlfYA7dYpSTNPXMRSLtsz1palmaDBTE/V2xKtTH6p
VglRNRUKawKBgCPqBh2TkN9czC2RFkgMb4FcqycN0jEQ0F6TSnVVhtNiAzKmc8s1
POvWJZJDIzjkv/mP+JUeXAdD/bdjNc26EU126rA6KzGgsMPjYv9FymusDPybGGUc
Qt5j5RcpNgEkn/5ZPyAlXjCfjz+RxChTfAyGHRmqU9qoLMIFir3pJ7llAoGBAMNH
sIxENwlzqyafoUUlEq/pU7kZWuJmrO2FwqRDraYoCiM/NCRhxRQ/ng6NY1gejepw
abD2alXiV4alBSxubne6rFmhvA00y2mG40c6Ezmxn2ZpbX3dMQ6bMcPKp7QnXtLc
mCSL4FGK02ImUNDsd0RVVFw51DRId4rmsuJYMK9NAoGAKlYdc4784ixTD2ZICIOC
ZWPxPAyQUEA7EkuUhAX1bVNG6UJTYA8kmGcUCG4jPTgWzi00IyUUr8jK7efyU/zs
qiJuVs1bia+flYIQpysMl1VzZh8gW1nkB4SVPm5l2wBvVJDIr9Mc6rueC/oVNkh2
fLVGuFoTVIu2bF0cWAjNNMg=
-----END TESTING KEY-----`)

func TestNewApiClientUnauthenticatedHappyPath(t *testing.T) {
	client, err := NewApiClient(
		WithAPIVersion(headers.APIVersionValue),
	)
	if err != nil {
		t.Fatalf("error creating client: %v", err)
	}
	if client == nil {
		t.Fatalf("client is nil")
	}
}

func TestNewApiClientTokenAuthHappyPath(t *testing.T) {
	client, err := NewApiClient(WithTokenAuthentication(token))
	if err != nil {
		t.Fatalf("error creating client: %v", err)
	}
	if client == nil {
		t.Fatalf("client is nil")
	}
}

func TestNewApiClientAppAuthHappyPath(t *testing.T) {
	tmpfile, err := os.CreateTemp("", pemFileName)
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(tmpfile.Name())

	if _, err := tmpfile.Write(key); err != nil {
		t.Fatal(err)
	}
	if err := tmpfile.Close(); err != nil {
		t.Fatal(err)
	}

	client, err := NewApiClient(WithGitHubAppAuthentication(tmpfile.Name(), clientID, installationID))
	if err != nil {
		t.Fatalf("error creating client: %v", err)
	}
	if client == nil {
		t.Fatalf("client is nil")
	}
}

func TestNewApiClientAppAuthBaseUrl(t *testing.T) {
	tmpfile, err := os.CreateTemp("", pemFileName)
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(tmpfile.Name())

	if _, err := tmpfile.Write(key); err != nil {
		t.Fatal(err)
	}
	if err := tmpfile.Close(); err != nil {
		t.Fatal(err)
	}

	expectedCall := false
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if strings.Contains(r.URL.Path, "/app/installations") {
			expectedCall = true
		}
	}))

	client, err := NewApiClient(
		WithGitHubAppAuthentication(tmpfile.Name(), clientID, installationID),
		WithBaseUrl(server.URL),
	)
	if err != nil {
		t.Fatalf("error creating client: %v", err)
	}
	if client == nil {
		t.Fatalf("client is nil")
	}
	queryParams := &installation.RepositoriesRequestBuilderGetQueryParameters{}
	requestConfig := &abs.RequestConfiguration[installation.RepositoriesRequestBuilderGetQueryParameters]{
		QueryParameters: queryParams,
	}

	// trigger a refresh of the installation token to the expected url
	_, _ = client.Installation().Repositories().Get(context.Background(), requestConfig)
	if !expectedCall {
		t.Errorf("installation token endpoint not called")
	}
}

func TestNewApiClientAppAuthErrorGettingKeyFromFile(t *testing.T) {
	client, err := NewApiClient(WithGitHubAppAuthentication("pem/file/does/not/exist.pem", clientID, installationID))
	if err == nil {
		t.Fatalf("expected error creating client")
	}
	if client != nil {
		t.Fatalf("client is not nil")
	}
}

func TestNewApiClientAppAuthWithAppIDHappyPath(t *testing.T) {
	tmpfile, err := os.CreateTemp("", pemFileName)
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(tmpfile.Name())

	if _, err := tmpfile.Write(key); err != nil {
		t.Fatal(err)
	}
	if err := tmpfile.Close(); err != nil {
		t.Fatal(err)
	}

	client, err := NewApiClient(WithGitHubAppAuthenticationUsingAppID(tmpfile.Name(), appID, installationID))
	if err != nil {
		t.Fatalf("error creating client: %v", err)
	}
	if client == nil {
		t.Fatalf("client is nil")
	}
}

func TestNewApiClientAppAuthWithAppIDErrorGettingKeyFromFile(t *testing.T) {
	client, err := NewApiClient(WithGitHubAppAuthenticationUsingAppID("pem/file/does/not/exist.pem", appID, installationID))
	if err == nil {
		t.Fatalf("expected error creating client")
	}
	if client != nil {
		t.Fatalf("client is not nil")
	}
}
