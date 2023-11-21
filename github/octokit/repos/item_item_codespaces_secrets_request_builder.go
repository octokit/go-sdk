package repos

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemItemCodespacesSecretsRequestBuilder builds and executes requests for operations under \repos\{owner}\{repo}\codespaces\secrets
type ItemItemCodespacesSecretsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemItemCodespacesSecretsRequestBuilderGetQueryParameters lists all development environment secrets available in a repository without revealing their encrypted values. You must authenticate using an access token with the `repo` scope to use this endpoint. GitHub Apps must have write access to the `codespaces_secrets` repository permission to use this endpoint.
type ItemItemCodespacesSecretsRequestBuilderGetQueryParameters struct {
    // Page number of the results to fetch.
    Page *int32 `uriparametername:"page"`
    // The number of results per page (max 100).
    Per_page *int32 `uriparametername:"per_page"`
}
// ItemItemCodespacesSecretsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemItemCodespacesSecretsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemItemCodespacesSecretsRequestBuilderGetQueryParameters
}
// BySecret_name gets an item from the octokit.repos.item.item.codespaces.secrets.item collection
func (m *ItemItemCodespacesSecretsRequestBuilder) BySecret_name(secret_name string)(*ItemItemCodespacesSecretsWithSecret_nameItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if secret_name != "" {
        urlTplParams["secret_name"] = secret_name
    }
    return NewItemItemCodespacesSecretsWithSecret_nameItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewItemItemCodespacesSecretsRequestBuilderInternal instantiates a new SecretsRequestBuilder and sets the default values.
func NewItemItemCodespacesSecretsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemCodespacesSecretsRequestBuilder) {
    m := &ItemItemCodespacesSecretsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/repos/{owner}/{repo}/codespaces/secrets{?per_page*,page*}", pathParameters),
    }
    return m
}
// NewItemItemCodespacesSecretsRequestBuilder instantiates a new SecretsRequestBuilder and sets the default values.
func NewItemItemCodespacesSecretsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemCodespacesSecretsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemItemCodespacesSecretsRequestBuilderInternal(urlParams, requestAdapter)
}
// Get lists all development environment secrets available in a repository without revealing their encrypted values. You must authenticate using an access token with the `repo` scope to use this endpoint. GitHub Apps must have write access to the `codespaces_secrets` repository permission to use this endpoint.
// Deprecated: This method is obsolete. Use GetAsSecretsGetResponse instead.
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/codespaces/repository-secrets#list-repository-secrets
func (m *ItemItemCodespacesSecretsRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemItemCodespacesSecretsRequestBuilderGetRequestConfiguration)(ItemItemCodespacesSecretsResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateItemItemCodespacesSecretsResponseFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ItemItemCodespacesSecretsResponseable), nil
}
// GetAsSecretsGetResponse lists all development environment secrets available in a repository without revealing their encrypted values. You must authenticate using an access token with the `repo` scope to use this endpoint. GitHub Apps must have write access to the `codespaces_secrets` repository permission to use this endpoint.
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/codespaces/repository-secrets#list-repository-secrets
func (m *ItemItemCodespacesSecretsRequestBuilder) GetAsSecretsGetResponse(ctx context.Context, requestConfiguration *ItemItemCodespacesSecretsRequestBuilderGetRequestConfiguration)(ItemItemCodespacesSecretsGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateItemItemCodespacesSecretsGetResponseFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ItemItemCodespacesSecretsGetResponseable), nil
}
// PublicKey the publicKey property
func (m *ItemItemCodespacesSecretsRequestBuilder) PublicKey()(*ItemItemCodespacesSecretsPublicKeyRequestBuilder) {
    return NewItemItemCodespacesSecretsPublicKeyRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToGetRequestInformation lists all development environment secrets available in a repository without revealing their encrypted values. You must authenticate using an access token with the `repo` scope to use this endpoint. GitHub Apps must have write access to the `codespaces_secrets` repository permission to use this endpoint.
func (m *ItemItemCodespacesSecretsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemItemCodespacesSecretsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
func (m *ItemItemCodespacesSecretsRequestBuilder) WithUrl(rawUrl string)(*ItemItemCodespacesSecretsRequestBuilder) {
    return NewItemItemCodespacesSecretsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
