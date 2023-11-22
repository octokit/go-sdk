package repos

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemItemActionsOrganizationSecretsRequestBuilder builds and executes requests for operations under \repos\{owner}\{repo}\actions\organization-secrets
type ItemItemActionsOrganizationSecretsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemItemActionsOrganizationSecretsRequestBuilderGetQueryParameters lists all organization secrets shared with a repository without revealing their encryptedvalues.You must authenticate using an access token with the `repo` scope to use this endpoint.GitHub Apps must have the `secrets` repository permission to use this endpoint.Authenticated users must have collaborator access to a repository to create, update, or read secrets.
type ItemItemActionsOrganizationSecretsRequestBuilderGetQueryParameters struct {
    // Page number of the results to fetch.
    Page *int32 `uriparametername:"page"`
    // The number of results per page (max 100).
    Per_page *int32 `uriparametername:"per_page"`
}
// ItemItemActionsOrganizationSecretsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemItemActionsOrganizationSecretsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemItemActionsOrganizationSecretsRequestBuilderGetQueryParameters
}
// NewItemItemActionsOrganizationSecretsRequestBuilderInternal instantiates a new OrganizationSecretsRequestBuilder and sets the default values.
func NewItemItemActionsOrganizationSecretsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemActionsOrganizationSecretsRequestBuilder) {
    m := &ItemItemActionsOrganizationSecretsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/repos/{owner}/{repo}/actions/organization-secrets{?per_page*,page*}", pathParameters),
    }
    return m
}
// NewItemItemActionsOrganizationSecretsRequestBuilder instantiates a new OrganizationSecretsRequestBuilder and sets the default values.
func NewItemItemActionsOrganizationSecretsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemActionsOrganizationSecretsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemItemActionsOrganizationSecretsRequestBuilderInternal(urlParams, requestAdapter)
}
// Get lists all organization secrets shared with a repository without revealing their encryptedvalues.You must authenticate using an access token with the `repo` scope to use this endpoint.GitHub Apps must have the `secrets` repository permission to use this endpoint.Authenticated users must have collaborator access to a repository to create, update, or read secrets.
// Deprecated: This method is obsolete. Use GetAsOrganizationSecretsGetResponse instead.
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/actions/secrets#list-repository-organization-secrets
func (m *ItemItemActionsOrganizationSecretsRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemItemActionsOrganizationSecretsRequestBuilderGetRequestConfiguration)(ItemItemActionsOrganizationSecretsResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateItemItemActionsOrganizationSecretsResponseFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ItemItemActionsOrganizationSecretsResponseable), nil
}
// GetAsOrganizationSecretsGetResponse lists all organization secrets shared with a repository without revealing their encryptedvalues.You must authenticate using an access token with the `repo` scope to use this endpoint.GitHub Apps must have the `secrets` repository permission to use this endpoint.Authenticated users must have collaborator access to a repository to create, update, or read secrets.
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/actions/secrets#list-repository-organization-secrets
func (m *ItemItemActionsOrganizationSecretsRequestBuilder) GetAsOrganizationSecretsGetResponse(ctx context.Context, requestConfiguration *ItemItemActionsOrganizationSecretsRequestBuilderGetRequestConfiguration)(ItemItemActionsOrganizationSecretsGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateItemItemActionsOrganizationSecretsGetResponseFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ItemItemActionsOrganizationSecretsGetResponseable), nil
}
// ToGetRequestInformation lists all organization secrets shared with a repository without revealing their encryptedvalues.You must authenticate using an access token with the `repo` scope to use this endpoint.GitHub Apps must have the `secrets` repository permission to use this endpoint.Authenticated users must have collaborator access to a repository to create, update, or read secrets.
func (m *ItemItemActionsOrganizationSecretsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemItemActionsOrganizationSecretsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *ItemItemActionsOrganizationSecretsRequestBuilder) WithUrl(rawUrl string)(*ItemItemActionsOrganizationSecretsRequestBuilder) {
    return NewItemItemActionsOrganizationSecretsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
