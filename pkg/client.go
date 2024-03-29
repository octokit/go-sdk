package pkg

import (
	"fmt"
	"time"

	kiotaHttp "github.com/microsoft/kiota-http-go"
	auth "github.com/octokit/go-sdk/pkg/authentication"
	"github.com/octokit/go-sdk/pkg/github"
	"github.com/octokit/go-sdk/pkg/handlers"
)

func NewApiClient(optionFuncs ...ClientOptionFunc) (*Client, error) {
	options := GetDefaultClientOptions()
	for _, optionFunc := range optionFuncs {
		optionFunc(options)
	}

	rateLimitHandler := handlers.NewRateLimitHandler()
	middlewares := kiotaHttp.GetDefaultMiddlewares()
	middlewares = append(middlewares, rateLimitHandler)
	netHttpClient := kiotaHttp.GetDefaultClient(middlewares...)

	tokenProviderOptions := []auth.TokenProviderOption{
		auth.WithAPIVersion(options.APIVersion),
		auth.WithUserAgent(options.UserAgent),
	}
	if options.Token != "" {
		tokenProviderOptions = append(tokenProviderOptions, auth.WithAuthorizationToken(options.Token))
	}

	tokenProvider := auth.NewTokenProvider(tokenProviderOptions...)

	adapter, err := kiotaHttp.NewNetHttpRequestAdapterWithParseNodeFactoryAndSerializationWriterFactoryAndHttpClient(tokenProvider, nil, nil, netHttpClient)
	if err != nil {
		return nil, fmt.Errorf("failed to create request adapter: %v", err)
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
// Then later in the constructor when the chain of objects are put together, all configuration
// will be drawn from this (hydrated) struct.
type ClientOptions struct {
	UserAgent      string
	APIVersion     string
	RequestTimeout time.Duration
	Middleware     []kiotaHttp.Middleware
	BaseURL        string
	Token          string
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

// example of an auth token provider option
func WithUserAgent(userAgent string) ClientOptionFunc {
	return func(c *ClientOptions) {
		c.UserAgent = userAgent
	}
}

// example of a netHttpClient option
func WithRequestTimeout(timeout time.Duration) ClientOptionFunc {
	return func(c *ClientOptions) {
		c.RequestTimeout = timeout
	}
}

// example of a middleware option
func WithRateLimitingMiddleware() ClientOptionFunc {
	return func(c *ClientOptions) {
		rateLimitHandler := handlers.NewRateLimitHandler()
		c.Middleware = append(c.Middleware, rateLimitHandler)
	}
}

// example of an adapter option
func WithBaseUrl(baseURL string) ClientOptionFunc {
	return func(c *ClientOptions) {
		c.BaseURL = baseURL
	}
}

func WithAuthorizationToken(token string) ClientOptionFunc {
	return func(c *ClientOptions) {
		c.Token = token
	}
}

func WithAPIVersion(version string) ClientOptionFunc {
	return func(c *ClientOptions) {
		c.APIVersion = version
	}
}
