package orgs

import (
    "context"
    i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e "github.com/octokit/go-sdk/github/octokit/models"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemOrganizationRolesItemUsersRequestBuilder builds and executes requests for operations under \orgs\{org}\organization-roles\{role_id}\users
type ItemOrganizationRolesItemUsersRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemOrganizationRolesItemUsersRequestBuilderGetQueryParameters lists organization members that are assigned to an organization role.To use this endpoint, you must be an administrator for the organization, and you must use an access token with the `admin:org` scope.GitHub Apps must have the `members` organization read permission to use this endpoint.For more information on organization roles, see "[Managing people's access to your organization with roles](https://docs.github.com/organizations/managing-peoples-access-to-your-organization-with-roles/about-custom-organization-roles)."
type ItemOrganizationRolesItemUsersRequestBuilderGetQueryParameters struct {
    // Page number of the results to fetch.
    Page *int32 `uriparametername:"page"`
    // The number of results per page (max 100).
    Per_page *int32 `uriparametername:"per_page"`
}
// ItemOrganizationRolesItemUsersRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemOrganizationRolesItemUsersRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemOrganizationRolesItemUsersRequestBuilderGetQueryParameters
}
// NewItemOrganizationRolesItemUsersRequestBuilderInternal instantiates a new UsersRequestBuilder and sets the default values.
func NewItemOrganizationRolesItemUsersRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemOrganizationRolesItemUsersRequestBuilder) {
    m := &ItemOrganizationRolesItemUsersRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/orgs/{org}/organization-roles/{role_id}/users{?per_page*,page*}", pathParameters),
    }
    return m
}
// NewItemOrganizationRolesItemUsersRequestBuilder instantiates a new UsersRequestBuilder and sets the default values.
func NewItemOrganizationRolesItemUsersRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemOrganizationRolesItemUsersRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemOrganizationRolesItemUsersRequestBuilderInternal(urlParams, requestAdapter)
}
// Get lists organization members that are assigned to an organization role.To use this endpoint, you must be an administrator for the organization, and you must use an access token with the `admin:org` scope.GitHub Apps must have the `members` organization read permission to use this endpoint.For more information on organization roles, see "[Managing people's access to your organization with roles](https://docs.github.com/organizations/managing-peoples-access-to-your-organization-with-roles/about-custom-organization-roles)."
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/orgs/organization-roles#list-users-that-are-assigned-to-an-organization-role
func (m *ItemOrganizationRolesItemUsersRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemOrganizationRolesItemUsersRequestBuilderGetRequestConfiguration)([]i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.SimpleUserable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.SendCollection(ctx, requestInfo, i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.CreateSimpleUserFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    val := make([]i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.SimpleUserable, len(res))
    for i, v := range res {
        if v != nil {
            val[i] = v.(i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.SimpleUserable)
        }
    }
    return val, nil
}
// ToGetRequestInformation lists organization members that are assigned to an organization role.To use this endpoint, you must be an administrator for the organization, and you must use an access token with the `admin:org` scope.GitHub Apps must have the `members` organization read permission to use this endpoint.For more information on organization roles, see "[Managing people's access to your organization with roles](https://docs.github.com/organizations/managing-peoples-access-to-your-organization-with-roles/about-custom-organization-roles)."
func (m *ItemOrganizationRolesItemUsersRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemOrganizationRolesItemUsersRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *ItemOrganizationRolesItemUsersRequestBuilder) WithUrl(rawUrl string)(*ItemOrganizationRolesItemUsersRequestBuilder) {
    return NewItemOrganizationRolesItemUsersRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
