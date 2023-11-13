package repositories

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemEnvironmentsItemSecretsRequestBuilder builds and executes requests for operations under \repositories\{repository_id}\environments\{environment_name}\secrets
type ItemEnvironmentsItemSecretsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemEnvironmentsItemSecretsRequestBuilderGetQueryParameters lists all secrets available in an environment without revealing theirencrypted values.You must authenticate using an access token with the `repo` scope to use this endpoint.GitHub Apps must have the `secrets` repository permission to use this endpoint.Authenticated users must have collaborator access to a repository to create, update, or read secrets.
type ItemEnvironmentsItemSecretsRequestBuilderGetQueryParameters struct {
    // Page number of the results to fetch.
    Page *int32 `uriparametername:"page"`
    // The number of results per page (max 100).
    Per_page *int32 `uriparametername:"per_page"`
}
// ItemEnvironmentsItemSecretsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemEnvironmentsItemSecretsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemEnvironmentsItemSecretsRequestBuilderGetQueryParameters
}
// BySecret_name gets an item from the go-sdk.repositories.item.environments.item.secrets.item collection
func (m *ItemEnvironmentsItemSecretsRequestBuilder) BySecret_name(secret_name string)(*ItemEnvironmentsItemSecretsWithSecret_nameItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if secret_name != "" {
        urlTplParams["secret_name"] = secret_name
    }
    return NewItemEnvironmentsItemSecretsWithSecret_nameItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewItemEnvironmentsItemSecretsRequestBuilderInternal instantiates a new SecretsRequestBuilder and sets the default values.
func NewItemEnvironmentsItemSecretsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemEnvironmentsItemSecretsRequestBuilder) {
    m := &ItemEnvironmentsItemSecretsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/repositories/{repository_id}/environments/{environment_name}/secrets{?per_page*,page*}", pathParameters),
    }
    return m
}
// NewItemEnvironmentsItemSecretsRequestBuilder instantiates a new SecretsRequestBuilder and sets the default values.
func NewItemEnvironmentsItemSecretsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemEnvironmentsItemSecretsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemEnvironmentsItemSecretsRequestBuilderInternal(urlParams, requestAdapter)
}
// Get lists all secrets available in an environment without revealing theirencrypted values.You must authenticate using an access token with the `repo` scope to use this endpoint.GitHub Apps must have the `secrets` repository permission to use this endpoint.Authenticated users must have collaborator access to a repository to create, update, or read secrets.
// Deprecated: This method is obsolete. Use GetAsSecretsGetResponse instead.
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/actions/secrets#list-environment-secrets
func (m *ItemEnvironmentsItemSecretsRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemEnvironmentsItemSecretsRequestBuilderGetRequestConfiguration)(ItemEnvironmentsItemSecretsResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateItemEnvironmentsItemSecretsResponseFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ItemEnvironmentsItemSecretsResponseable), nil
}
// GetAsSecretsGetResponse lists all secrets available in an environment without revealing theirencrypted values.You must authenticate using an access token with the `repo` scope to use this endpoint.GitHub Apps must have the `secrets` repository permission to use this endpoint.Authenticated users must have collaborator access to a repository to create, update, or read secrets.
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/actions/secrets#list-environment-secrets
func (m *ItemEnvironmentsItemSecretsRequestBuilder) GetAsSecretsGetResponse(ctx context.Context, requestConfiguration *ItemEnvironmentsItemSecretsRequestBuilderGetRequestConfiguration)(ItemEnvironmentsItemSecretsGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateItemEnvironmentsItemSecretsGetResponseFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ItemEnvironmentsItemSecretsGetResponseable), nil
}
// PublicKey the publicKey property
func (m *ItemEnvironmentsItemSecretsRequestBuilder) PublicKey()(*ItemEnvironmentsItemSecretsPublicKeyRequestBuilder) {
    return NewItemEnvironmentsItemSecretsPublicKeyRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToGetRequestInformation lists all secrets available in an environment without revealing theirencrypted values.You must authenticate using an access token with the `repo` scope to use this endpoint.GitHub Apps must have the `secrets` repository permission to use this endpoint.Authenticated users must have collaborator access to a repository to create, update, or read secrets.
func (m *ItemEnvironmentsItemSecretsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemEnvironmentsItemSecretsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *ItemEnvironmentsItemSecretsRequestBuilder) WithUrl(rawUrl string)(*ItemEnvironmentsItemSecretsRequestBuilder) {
    return NewItemEnvironmentsItemSecretsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
