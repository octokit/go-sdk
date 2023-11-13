package orgs

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemActionsSecretsRequestBuilder builds and executes requests for operations under \orgs\{org}\actions\secrets
type ItemActionsSecretsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemActionsSecretsRequestBuilderGetQueryParameters lists all secrets available in an organization without revealing theirencrypted values.You must authenticate using an access token with the `admin:org` scope to use this endpoint.If the repository is private, you must use an access token with the `repo` scope.GitHub Apps must have the `secrets` organization permission to use this endpoint.Authenticated users must have collaborator access to a repository to create, update, or read secrets.
type ItemActionsSecretsRequestBuilderGetQueryParameters struct {
    // Page number of the results to fetch.
    Page *int32 `uriparametername:"page"`
    // The number of results per page (max 100).
    Per_page *int32 `uriparametername:"per_page"`
}
// ItemActionsSecretsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemActionsSecretsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemActionsSecretsRequestBuilderGetQueryParameters
}
// BySecret_name gets an item from the go-sdk.orgs.item.actions.secrets.item collection
func (m *ItemActionsSecretsRequestBuilder) BySecret_name(secret_name string)(*ItemActionsSecretsWithSecret_nameItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if secret_name != "" {
        urlTplParams["secret_name"] = secret_name
    }
    return NewItemActionsSecretsWithSecret_nameItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewItemActionsSecretsRequestBuilderInternal instantiates a new SecretsRequestBuilder and sets the default values.
func NewItemActionsSecretsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemActionsSecretsRequestBuilder) {
    m := &ItemActionsSecretsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/orgs/{org}/actions/secrets{?per_page*,page*}", pathParameters),
    }
    return m
}
// NewItemActionsSecretsRequestBuilder instantiates a new SecretsRequestBuilder and sets the default values.
func NewItemActionsSecretsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemActionsSecretsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemActionsSecretsRequestBuilderInternal(urlParams, requestAdapter)
}
// Get lists all secrets available in an organization without revealing theirencrypted values.You must authenticate using an access token with the `admin:org` scope to use this endpoint.If the repository is private, you must use an access token with the `repo` scope.GitHub Apps must have the `secrets` organization permission to use this endpoint.Authenticated users must have collaborator access to a repository to create, update, or read secrets.
// Deprecated: This method is obsolete. Use GetAsSecretsGetResponse instead.
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/actions/secrets#list-organization-secrets
func (m *ItemActionsSecretsRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemActionsSecretsRequestBuilderGetRequestConfiguration)(ItemActionsSecretsResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateItemActionsSecretsResponseFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ItemActionsSecretsResponseable), nil
}
// GetAsSecretsGetResponse lists all secrets available in an organization without revealing theirencrypted values.You must authenticate using an access token with the `admin:org` scope to use this endpoint.If the repository is private, you must use an access token with the `repo` scope.GitHub Apps must have the `secrets` organization permission to use this endpoint.Authenticated users must have collaborator access to a repository to create, update, or read secrets.
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/actions/secrets#list-organization-secrets
func (m *ItemActionsSecretsRequestBuilder) GetAsSecretsGetResponse(ctx context.Context, requestConfiguration *ItemActionsSecretsRequestBuilderGetRequestConfiguration)(ItemActionsSecretsGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateItemActionsSecretsGetResponseFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ItemActionsSecretsGetResponseable), nil
}
// PublicKey the publicKey property
func (m *ItemActionsSecretsRequestBuilder) PublicKey()(*ItemActionsSecretsPublicKeyRequestBuilder) {
    return NewItemActionsSecretsPublicKeyRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToGetRequestInformation lists all secrets available in an organization without revealing theirencrypted values.You must authenticate using an access token with the `admin:org` scope to use this endpoint.If the repository is private, you must use an access token with the `repo` scope.GitHub Apps must have the `secrets` organization permission to use this endpoint.Authenticated users must have collaborator access to a repository to create, update, or read secrets.
func (m *ItemActionsSecretsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemActionsSecretsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *ItemActionsSecretsRequestBuilder) WithUrl(rawUrl string)(*ItemActionsSecretsRequestBuilder) {
    return NewItemActionsSecretsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
