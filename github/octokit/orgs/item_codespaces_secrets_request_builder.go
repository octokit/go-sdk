package orgs

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemCodespacesSecretsRequestBuilder builds and executes requests for operations under \orgs\{org}\codespaces\secrets
type ItemCodespacesSecretsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemCodespacesSecretsRequestBuilderGetQueryParameters lists all Codespaces development environment secrets available at the organization-level without revealing their encrypted values.You must authenticate using an access token with the `admin:org` scope to use this endpoint.
type ItemCodespacesSecretsRequestBuilderGetQueryParameters struct {
    // Page number of the results to fetch.
    Page *int32 `uriparametername:"page"`
    // The number of results per page (max 100).
    Per_page *int32 `uriparametername:"per_page"`
}
// ItemCodespacesSecretsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemCodespacesSecretsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemCodespacesSecretsRequestBuilderGetQueryParameters
}
// BySecret_name gets an item from the github.com/octokit/go-sdk/github/octokit/.orgs.item.codespaces.secrets.item collection
func (m *ItemCodespacesSecretsRequestBuilder) BySecret_name(secret_name string)(*ItemCodespacesSecretsWithSecret_nameItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if secret_name != "" {
        urlTplParams["secret_name"] = secret_name
    }
    return NewItemCodespacesSecretsWithSecret_nameItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewItemCodespacesSecretsRequestBuilderInternal instantiates a new SecretsRequestBuilder and sets the default values.
func NewItemCodespacesSecretsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemCodespacesSecretsRequestBuilder) {
    m := &ItemCodespacesSecretsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/orgs/{org}/codespaces/secrets{?per_page*,page*}", pathParameters),
    }
    return m
}
// NewItemCodespacesSecretsRequestBuilder instantiates a new SecretsRequestBuilder and sets the default values.
func NewItemCodespacesSecretsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemCodespacesSecretsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemCodespacesSecretsRequestBuilderInternal(urlParams, requestAdapter)
}
// Get lists all Codespaces development environment secrets available at the organization-level without revealing their encrypted values.You must authenticate using an access token with the `admin:org` scope to use this endpoint.
// Deprecated: This method is obsolete. Use GetAsSecretsGetResponse instead.
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/codespaces/organization-secrets#list-organization-secrets
func (m *ItemCodespacesSecretsRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemCodespacesSecretsRequestBuilderGetRequestConfiguration)(ItemCodespacesSecretsResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateItemCodespacesSecretsResponseFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ItemCodespacesSecretsResponseable), nil
}
// GetAsSecretsGetResponse lists all Codespaces development environment secrets available at the organization-level without revealing their encrypted values.You must authenticate using an access token with the `admin:org` scope to use this endpoint.
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/codespaces/organization-secrets#list-organization-secrets
func (m *ItemCodespacesSecretsRequestBuilder) GetAsSecretsGetResponse(ctx context.Context, requestConfiguration *ItemCodespacesSecretsRequestBuilderGetRequestConfiguration)(ItemCodespacesSecretsGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateItemCodespacesSecretsGetResponseFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ItemCodespacesSecretsGetResponseable), nil
}
// PublicKey the publicKey property
func (m *ItemCodespacesSecretsRequestBuilder) PublicKey()(*ItemCodespacesSecretsPublicKeyRequestBuilder) {
    return NewItemCodespacesSecretsPublicKeyRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToGetRequestInformation lists all Codespaces development environment secrets available at the organization-level without revealing their encrypted values.You must authenticate using an access token with the `admin:org` scope to use this endpoint.
func (m *ItemCodespacesSecretsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemCodespacesSecretsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
func (m *ItemCodespacesSecretsRequestBuilder) WithUrl(rawUrl string)(*ItemCodespacesSecretsRequestBuilder) {
    return NewItemCodespacesSecretsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
