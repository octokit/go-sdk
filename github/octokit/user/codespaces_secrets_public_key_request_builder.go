package user

import (
    "context"
    i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e "github.com/octokit/go-sdk/github/octokit/models"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// CodespacesSecretsPublicKeyRequestBuilder builds and executes requests for operations under \user\codespaces\secrets\public-key
type CodespacesSecretsPublicKeyRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// CodespacesSecretsPublicKeyRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CodespacesSecretsPublicKeyRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewCodespacesSecretsPublicKeyRequestBuilderInternal instantiates a new PublicKeyRequestBuilder and sets the default values.
func NewCodespacesSecretsPublicKeyRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CodespacesSecretsPublicKeyRequestBuilder) {
    m := &CodespacesSecretsPublicKeyRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/user/codespaces/secrets/public-key", pathParameters),
    }
    return m
}
// NewCodespacesSecretsPublicKeyRequestBuilder instantiates a new PublicKeyRequestBuilder and sets the default values.
func NewCodespacesSecretsPublicKeyRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CodespacesSecretsPublicKeyRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewCodespacesSecretsPublicKeyRequestBuilderInternal(urlParams, requestAdapter)
}
// Get gets your public key, which you need to encrypt secrets. You need to encrypt a secret before you can create or update secrets.You must authenticate using an access token with the `codespace` or `codespace:secrets` scope to use this endpoint. User must have Codespaces access to use this endpoint.GitHub Apps must have read access to the `codespaces_user_secrets` user permission to use this endpoint.
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/codespaces/secrets#get-public-key-for-the-authenticated-user
func (m *CodespacesSecretsPublicKeyRequestBuilder) Get(ctx context.Context, requestConfiguration *CodespacesSecretsPublicKeyRequestBuilderGetRequestConfiguration)(i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.CodespacesUserPublicKeyable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.CreateCodespacesUserPublicKeyFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.CodespacesUserPublicKeyable), nil
}
// ToGetRequestInformation gets your public key, which you need to encrypt secrets. You need to encrypt a secret before you can create or update secrets.You must authenticate using an access token with the `codespace` or `codespace:secrets` scope to use this endpoint. User must have Codespaces access to use this endpoint.GitHub Apps must have read access to the `codespaces_user_secrets` user permission to use this endpoint.
func (m *CodespacesSecretsPublicKeyRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *CodespacesSecretsPublicKeyRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
func (m *CodespacesSecretsPublicKeyRequestBuilder) WithUrl(rawUrl string)(*CodespacesSecretsPublicKeyRequestBuilder) {
    return NewCodespacesSecretsPublicKeyRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
