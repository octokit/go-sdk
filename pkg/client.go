package pkg

import (
	"fmt"
	"net/http"
	"time"

	"github.com/kfcampbell/ghinstallation"
	kiotaHttp "github.com/microsoft/kiota-http-go"
	auth "github.com/octokit/go-sdk/pkg/authentication"
	"github.com/octokit/go-sdk/pkg/github"
	"github.com/octokit/go-sdk/pkg/handlers"
)

// NewApiClient is a convenience constructor to create a new instance of a
// Client (wrapper of *github.ApiClient) with the provided option functions.
// By default, it includes a rate limiting middleware.
func NewApiClient(optionFuncs ...ClientOptionFunc) (*Client, error) {
	options := GetDefaultClientOptions()
	for _, optionFunc := range optionFuncs {
		optionFunc(options)
	}

	rateLimitHandler := handlers.NewRateLimitHandler()
	middlewares := options.Middleware
	middlewares = append(middlewares, rateLimitHandler)
	defaultTransport := kiotaHttp.GetDefaultTransport()
	netHttpClient := &http.Client{
		Transport: defaultTransport,
	}

	if options.RequestTimeout != 0 {
		netHttpClient.Timeout = options.RequestTimeout
	}

	// Configure GitHub App authentication if required fields are provided
	if options.GitHubAppID != 0 && options.GitHubAppInstallationID != 0 && options.GitHubAppPemFilePath != "" {
		existingTransport := netHttpClient.Transport
		appTransport, err := ghinstallation.NewKeyFromFile(existingTransport, options.GitHubAppID, options.GitHubAppInstallationID, options.GitHubAppPemFilePath)
		if err != nil {
			return nil, fmt.Errorf("failed to create transport from GitHub App: %v", err)
		}
		netHttpClient.Transport = appTransport
	}

	// Middleware must be applied after App transport is set, otherwise App token will fail to be
	// renewed with a 400 Bad Request error (even though the request is identical to a successful one.)
	finalTransport := kiotaHttp.NewCustomTransportWithParentTransport(netHttpClient.Transport, middlewares...)
	netHttpClient.Transport = finalTransport

	tokenProviderOptions := []auth.TokenProviderOption{
		auth.WithAPIVersion(options.APIVersion),
		auth.WithUserAgent(options.UserAgent),
	}

	// If a PAT is provided and GitHub App information is not, configure token authentication
	if options.Token != "" && (options.GitHubAppID == 0 && options.GitHubAppInstallationID == 0 && options.GitHubAppPemFilePath == "") {
		tokenProviderOptions = append(tokenProviderOptions, auth.WithTokenAuthentication(options.Token))
	}

	tokenProvider := auth.NewTokenProvider(tokenProviderOptions...)

	adapter, err := kiotaHttp.NewNetHttpRequestAdapterWithParseNodeFactoryAndSerializationWriterFactoryAndHttpClient(tokenProvider, nil, nil, netHttpClient)
	if err != nil {
		return nil, fmt.Errorf("failed to create request adapter: %v", err)
	}
	if options.BaseURL != "" {
		adapter.SetBaseUrl(options.BaseURL)
	}

	client := github.NewApiClient(adapter)
	sdkClient := &Client{
		ApiClient: client,
	}

	return sdkClient, nil
}

// Client wraps github.ApiClient so that we may provide neater constructors and ease of use
type Client struct {
	*github.ApiClient
}

// ClientOptions contains every setting we could possibly want to set for the token provider,
// the netHttpClient, the middleware, and the adapter. If we can possibly override it, it should
// be in this struct.
// In the constructor, when helper functions apply options, they'll be applied to this struct.
// Then later in the constructor when that chain of objects is put together, all configuration
// will be drawn from this (hydrated) struct.
type ClientOptions struct {
	UserAgent      string
	APIVersion     string
	RequestTimeout time.Duration
	Middleware     []kiotaHttp.Middleware
	BaseURL        string

	// Token should be left blank if GitHub App auth or an unauthenticated client is desired.
	Token string

	// GitHubAppPemFilePath should be left blank if token auth or an unauthenticated client is desired.
	GitHubAppPemFilePath string
	// GitHubAppID should be left blank if token auth or an unauthenticated client is desired.
	GitHubAppID int64
	// GitHubAppInstallationID should be left blank if token auth or an unauthenticated client is desired.
	GitHubAppInstallationID int64
}

// GetDefaultClientOptions returns a new instance of ClientOptions with default values.
// This is used in the NewApiClient constructor before applying user-defined custom options.
func GetDefaultClientOptions() *ClientOptions {
	return &ClientOptions{
		UserAgent:  "octokit/go-sdk",
		APIVersion: "2022-11-28",
		Middleware: kiotaHttp.GetDefaultMiddlewares(),
	}
}

// ClientOptionFunc provides a functional pattern for client configuration
type ClientOptionFunc func(*ClientOptions)

// WithUserAgent configures the client with the given user agent string.
func WithUserAgent(userAgent string) ClientOptionFunc {
	return func(c *ClientOptions) {
		c.UserAgent = userAgent
	}
}

// WithRequestTimeout configures the client with the given request timeout.
func WithRequestTimeout(timeout time.Duration) ClientOptionFunc {
	return func(c *ClientOptions) {
		c.RequestTimeout = timeout
	}
}

// WithBaseUrl configures the client with the given base URL.
func WithBaseUrl(baseURL string) ClientOptionFunc {
	return func(c *ClientOptions) {
		c.BaseURL = baseURL
	}
}

// WithAuthorizationToken configures the client with the given
// Personal Authorization Token.
func WithAuthorizationToken(token string) ClientOptionFunc {
	return func(c *ClientOptions) {
		c.Token = token
	}
}

// WithAPIVersion configures the client with the given API version.
func WithAPIVersion(version string) ClientOptionFunc {
	return func(c *ClientOptions) {
		c.APIVersion = version
	}
}

func WithGitHubAppAuthentication(GitHubAppPemFilePath string, GitHubAppID int64, GitHubAppInstallationID int64) ClientOptionFunc {
	return func(c *ClientOptions) {
		c.GitHubAppPemFilePath = GitHubAppPemFilePath
		c.GitHubAppID = GitHubAppID
		c.GitHubAppInstallationID = GitHubAppInstallationID
	}
}
