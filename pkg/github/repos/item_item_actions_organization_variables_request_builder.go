package repos

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemItemActionsOrganizationVariablesRequestBuilder builds and executes requests for operations under \repos\{repos-id}\{Owner-id}\actions\organization-variables
type ItemItemActionsOrganizationVariablesRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemItemActionsOrganizationVariablesRequestBuilderGetQueryParameters lists all organiation variables shared with a repository.You must authenticate using an access token with the `repo` scope to use this endpoint.GitHub Apps must have the `actions_variables:read` repository permission to use this endpoint.Authenticated users must have collaborator access to a repository to create, update, or read variables.
type ItemItemActionsOrganizationVariablesRequestBuilderGetQueryParameters struct {
    // The page number of the results to fetch. For more information, see "[Using pagination in the REST API](https://docs.github.com/rest/using-the-rest-api/using-pagination-in-the-rest-api)."
    Page *int32 `uriparametername:"page"`
    // The number of results per page (max 30). For more information, see "[Using pagination in the REST API](https://docs.github.com/rest/using-the-rest-api/using-pagination-in-the-rest-api)."
    Per_page *int32 `uriparametername:"per_page"`
}
// NewItemItemActionsOrganizationVariablesRequestBuilderInternal instantiates a new OrganizationVariablesRequestBuilder and sets the default values.
func NewItemItemActionsOrganizationVariablesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemActionsOrganizationVariablesRequestBuilder) {
    m := &ItemItemActionsOrganizationVariablesRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/repos/{repos%2Did}/{Owner%2Did}/actions/organization-variables{?page*,per_page*}", pathParameters),
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
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/actions/variables#list-repository-organization-variables
func (m *ItemItemActionsOrganizationVariablesRequestBuilder) Get(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[ItemItemActionsOrganizationVariablesRequestBuilderGetQueryParameters])(ItemItemActionsOrganizationVariablesGetResponseable, error) {
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
func (m *ItemItemActionsOrganizationVariablesRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[ItemItemActionsOrganizationVariablesRequestBuilderGetQueryParameters])(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ConfigureRequestInformation(requestInfo, requestConfiguration)
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
func (m *ItemItemActionsOrganizationVariablesRequestBuilder) WithUrl(rawUrl string)(*ItemItemActionsOrganizationVariablesRequestBuilder) {
    return NewItemItemActionsOrganizationVariablesRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
