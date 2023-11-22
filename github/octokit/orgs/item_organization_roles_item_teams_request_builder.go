package orgs

import (
    "context"
    i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e "github.com/octokit/go-sdk/github/octokit/models"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemOrganizationRolesItemTeamsRequestBuilder builds and executes requests for operations under \orgs\{org}\organization-roles\{role_id}\teams
type ItemOrganizationRolesItemTeamsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemOrganizationRolesItemTeamsRequestBuilderGetQueryParameters lists the teams that are assigned to an organization role.To use this endpoint, you must be an administrator for the organization, and you must use an access token with the `admin:org` scope.GitHub Apps must have the `members` organization read permission to use this endpoint.For more information on organization roles, see "[Managing people's access to your organization with roles](https://docs.github.com/organizations/managing-peoples-access-to-your-organization-with-roles/about-custom-organization-roles)."
type ItemOrganizationRolesItemTeamsRequestBuilderGetQueryParameters struct {
    // Page number of the results to fetch.
    Page *int32 `uriparametername:"page"`
    // The number of results per page (max 100).
    Per_page *int32 `uriparametername:"per_page"`
}
// ItemOrganizationRolesItemTeamsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemOrganizationRolesItemTeamsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemOrganizationRolesItemTeamsRequestBuilderGetQueryParameters
}
// NewItemOrganizationRolesItemTeamsRequestBuilderInternal instantiates a new TeamsRequestBuilder and sets the default values.
func NewItemOrganizationRolesItemTeamsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemOrganizationRolesItemTeamsRequestBuilder) {
    m := &ItemOrganizationRolesItemTeamsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/orgs/{org}/organization-roles/{role_id}/teams{?per_page*,page*}", pathParameters),
    }
    return m
}
// NewItemOrganizationRolesItemTeamsRequestBuilder instantiates a new TeamsRequestBuilder and sets the default values.
func NewItemOrganizationRolesItemTeamsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemOrganizationRolesItemTeamsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemOrganizationRolesItemTeamsRequestBuilderInternal(urlParams, requestAdapter)
}
// Get lists the teams that are assigned to an organization role.To use this endpoint, you must be an administrator for the organization, and you must use an access token with the `admin:org` scope.GitHub Apps must have the `members` organization read permission to use this endpoint.For more information on organization roles, see "[Managing people's access to your organization with roles](https://docs.github.com/organizations/managing-peoples-access-to-your-organization-with-roles/about-custom-organization-roles)."
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/orgs/organization-roles#list-teams-that-are-assigned-to-an-organization-role
func (m *ItemOrganizationRolesItemTeamsRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemOrganizationRolesItemTeamsRequestBuilderGetRequestConfiguration)([]i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.Teamable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.SendCollection(ctx, requestInfo, i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.CreateTeamFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    val := make([]i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.Teamable, len(res))
    for i, v := range res {
        if v != nil {
            val[i] = v.(i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.Teamable)
        }
    }
    return val, nil
}
// ToGetRequestInformation lists the teams that are assigned to an organization role.To use this endpoint, you must be an administrator for the organization, and you must use an access token with the `admin:org` scope.GitHub Apps must have the `members` organization read permission to use this endpoint.For more information on organization roles, see "[Managing people's access to your organization with roles](https://docs.github.com/organizations/managing-peoples-access-to-your-organization-with-roles/about-custom-organization-roles)."
func (m *ItemOrganizationRolesItemTeamsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemOrganizationRolesItemTeamsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *ItemOrganizationRolesItemTeamsRequestBuilder) WithUrl(rawUrl string)(*ItemOrganizationRolesItemTeamsRequestBuilder) {
    return NewItemOrganizationRolesItemTeamsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
