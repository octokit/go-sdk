package repos

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemItemActionsOrganizationVariablesRequestBuilder builds and executes requests for operations under \repos\{owner}\{repo}\actions\organization-variables
type ItemItemActionsOrganizationVariablesRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemItemActionsOrganizationVariablesRequestBuilderGetQueryParameters lists all organiation variables shared with a repository.You must authenticate using an access token with the `repo` scope to use this endpoint.GitHub Apps must have the `actions_variables:read` repository permission to use this endpoint.Authenticated users must have collaborator access to a repository to create, update, or read variables.
type ItemItemActionsOrganizationVariablesRequestBuilderGetQueryParameters struct {
    // Page number of the results to fetch.
    Page *int32 `uriparametername:"page"`
    // The number of results per page (max 30).
    Per_page *int32 `uriparametername:"per_page"`
}
// ItemItemActionsOrganizationVariablesRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemItemActionsOrganizationVariablesRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemItemActionsOrganizationVariablesRequestBuilderGetQueryParameters
}
// NewItemItemActionsOrganizationVariablesRequestBuilderInternal instantiates a new OrganizationVariablesRequestBuilder and sets the default values.
func NewItemItemActionsOrganizationVariablesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemActionsOrganizationVariablesRequestBuilder) {
    m := &ItemItemActionsOrganizationVariablesRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/repos/{owner}/{repo}/actions/organization-variables{?per_page*,page*}", pathParameters),
    }
    return m
}
// NewItemItemActionsOrganizationVariablesRequestBuilder instantiates a new OrganizationVariablesRequestBuilder and sets the default values.
func NewItemItemActionsOrganizationVariablesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemActionsOrganizationVariablesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemItemActionsOrganizationVariablesRequestBuilderInternal(urlParams, requestAdapter)
}
// Get lists all organiation variables shared with a repository.You must authenticate using an access token with the `repo` scope to use this endpoint.GitHub Apps must have the `actions_variables:read` repository permission to use this endpoint.Authenticated users must have collaborator access to a repository to create, update, or read variables.
// Deprecated: This method is obsolete. Use GetAsOrganizationVariablesGetResponse instead.
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/actions/variables#list-repository-organization-variables
func (m *ItemItemActionsOrganizationVariablesRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemItemActionsOrganizationVariablesRequestBuilderGetRequestConfiguration)(ItemItemActionsOrganizationVariablesResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateItemItemActionsOrganizationVariablesResponseFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ItemItemActionsOrganizationVariablesResponseable), nil
}
// GetAsOrganizationVariablesGetResponse lists all organiation variables shared with a repository.You must authenticate using an access token with the `repo` scope to use this endpoint.GitHub Apps must have the `actions_variables:read` repository permission to use this endpoint.Authenticated users must have collaborator access to a repository to create, update, or read variables.
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/actions/variables#list-repository-organization-variables
func (m *ItemItemActionsOrganizationVariablesRequestBuilder) GetAsOrganizationVariablesGetResponse(ctx context.Context, requestConfiguration *ItemItemActionsOrganizationVariablesRequestBuilderGetRequestConfiguration)(ItemItemActionsOrganizationVariablesGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateItemItemActionsOrganizationVariablesGetResponseFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ItemItemActionsOrganizationVariablesGetResponseable), nil
}
// ToGetRequestInformation lists all organiation variables shared with a repository.You must authenticate using an access token with the `repo` scope to use this endpoint.GitHub Apps must have the `actions_variables:read` repository permission to use this endpoint.Authenticated users must have collaborator access to a repository to create, update, or read variables.
func (m *ItemItemActionsOrganizationVariablesRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemItemActionsOrganizationVariablesRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *ItemItemActionsOrganizationVariablesRequestBuilder) WithUrl(rawUrl string)(*ItemItemActionsOrganizationVariablesRequestBuilder) {
    return NewItemItemActionsOrganizationVariablesRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
