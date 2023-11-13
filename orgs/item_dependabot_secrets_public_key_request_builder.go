package orgs

import (
    "context"
    i33bcfa37677dfd8b71eaf22ad3bfd140ba1702e7a8418558499c22418bf2fef2 "github.com/octokit/go-sdk/models"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemDependabotSecretsPublicKeyRequestBuilder builds and executes requests for operations under \orgs\{org}\dependabot\secrets\public-key
type ItemDependabotSecretsPublicKeyRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemDependabotSecretsPublicKeyRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemDependabotSecretsPublicKeyRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemDependabotSecretsPublicKeyRequestBuilderInternal instantiates a new PublicKeyRequestBuilder and sets the default values.
func NewItemDependabotSecretsPublicKeyRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemDependabotSecretsPublicKeyRequestBuilder) {
    m := &ItemDependabotSecretsPublicKeyRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/orgs/{org}/dependabot/secrets/public-key", pathParameters),
    }
    return m
}
// NewItemDependabotSecretsPublicKeyRequestBuilder instantiates a new PublicKeyRequestBuilder and sets the default values.
func NewItemDependabotSecretsPublicKeyRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemDependabotSecretsPublicKeyRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemDependabotSecretsPublicKeyRequestBuilderInternal(urlParams, requestAdapter)
}
// Get gets your public key, which you need to encrypt secrets. You need to encrypt a secret before you can create or update secrets. You must authenticate using an access token with the `admin:org` scope to use this endpoint. GitHub Apps must have the `dependabot_secrets` organization permission to use this endpoint.
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/dependabot/secrets#get-an-organization-public-key
func (m *ItemDependabotSecretsPublicKeyRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemDependabotSecretsPublicKeyRequestBuilderGetRequestConfiguration)(i33bcfa37677dfd8b71eaf22ad3bfd140ba1702e7a8418558499c22418bf2fef2.DependabotPublicKeyable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i33bcfa37677dfd8b71eaf22ad3bfd140ba1702e7a8418558499c22418bf2fef2.CreateDependabotPublicKeyFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i33bcfa37677dfd8b71eaf22ad3bfd140ba1702e7a8418558499c22418bf2fef2.DependabotPublicKeyable), nil
}
// ToGetRequestInformation gets your public key, which you need to encrypt secrets. You need to encrypt a secret before you can create or update secrets. You must authenticate using an access token with the `admin:org` scope to use this endpoint. GitHub Apps must have the `dependabot_secrets` organization permission to use this endpoint.
func (m *ItemDependabotSecretsPublicKeyRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemDependabotSecretsPublicKeyRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers.TryAdd("Accept", "application/json;q=1")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
func (m *ItemDependabotSecretsPublicKeyRequestBuilder) WithUrl(rawUrl string)(*ItemDependabotSecretsPublicKeyRequestBuilder) {
    return NewItemDependabotSecretsPublicKeyRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
