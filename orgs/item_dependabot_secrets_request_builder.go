package orgs

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemDependabotSecretsRequestBuilder builds and executes requests for operations under \orgs\{org}\dependabot\secrets
type ItemDependabotSecretsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemDependabotSecretsRequestBuilderGetQueryParameters lists all secrets available in an organization without revealing their encrypted values. You must authenticate using an access token with the `admin:org` scope to use this endpoint. GitHub Apps must have the `dependabot_secrets` organization permission to use this endpoint.
type ItemDependabotSecretsRequestBuilderGetQueryParameters struct {
    // Page number of the results to fetch.
    Page *int32 `uriparametername:"page"`
    // The number of results per page (max 100).
    Per_page *int32 `uriparametername:"per_page"`
}
// ItemDependabotSecretsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemDependabotSecretsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemDependabotSecretsRequestBuilderGetQueryParameters
}
// BySecret_name gets an item from the go-sdk.orgs.item.dependabot.secrets.item collection
func (m *ItemDependabotSecretsRequestBuilder) BySecret_name(secret_name string)(*ItemDependabotSecretsWithSecret_nameItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if secret_name != "" {
        urlTplParams["secret_name"] = secret_name
    }
    return NewItemDependabotSecretsWithSecret_nameItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewItemDependabotSecretsRequestBuilderInternal instantiates a new SecretsRequestBuilder and sets the default values.
func NewItemDependabotSecretsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemDependabotSecretsRequestBuilder) {
    m := &ItemDependabotSecretsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/orgs/{org}/dependabot/secrets{?per_page*,page*}", pathParameters),
    }
    return m
}
// NewItemDependabotSecretsRequestBuilder instantiates a new SecretsRequestBuilder and sets the default values.
func NewItemDependabotSecretsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemDependabotSecretsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemDependabotSecretsRequestBuilderInternal(urlParams, requestAdapter)
}
// Get lists all secrets available in an organization without revealing their encrypted values. You must authenticate using an access token with the `admin:org` scope to use this endpoint. GitHub Apps must have the `dependabot_secrets` organization permission to use this endpoint.
// Deprecated: This method is obsolete. Use GetAsSecretsGetResponse instead.
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/dependabot/secrets#list-organization-secrets
func (m *ItemDependabotSecretsRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemDependabotSecretsRequestBuilderGetRequestConfiguration)(ItemDependabotSecretsResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateItemDependabotSecretsResponseFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ItemDependabotSecretsResponseable), nil
}
// GetAsSecretsGetResponse lists all secrets available in an organization without revealing their encrypted values. You must authenticate using an access token with the `admin:org` scope to use this endpoint. GitHub Apps must have the `dependabot_secrets` organization permission to use this endpoint.
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/dependabot/secrets#list-organization-secrets
func (m *ItemDependabotSecretsRequestBuilder) GetAsSecretsGetResponse(ctx context.Context, requestConfiguration *ItemDependabotSecretsRequestBuilderGetRequestConfiguration)(ItemDependabotSecretsGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateItemDependabotSecretsGetResponseFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ItemDependabotSecretsGetResponseable), nil
}
// PublicKey the publicKey property
func (m *ItemDependabotSecretsRequestBuilder) PublicKey()(*ItemDependabotSecretsPublicKeyRequestBuilder) {
    return NewItemDependabotSecretsPublicKeyRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToGetRequestInformation lists all secrets available in an organization without revealing their encrypted values. You must authenticate using an access token with the `admin:org` scope to use this endpoint. GitHub Apps must have the `dependabot_secrets` organization permission to use this endpoint.
func (m *ItemDependabotSecretsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemDependabotSecretsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
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
func (m *ItemDependabotSecretsRequestBuilder) WithUrl(rawUrl string)(*ItemDependabotSecretsRequestBuilder) {
    return NewItemDependabotSecretsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
