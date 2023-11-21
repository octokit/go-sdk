package teams

import (
    "context"
    i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e "github.com/octokit/go-sdk/github/octokit/models"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i949f6571ab906194a9314d01e930f960285837b45292e75beed92b54c60429af "github.com/octokit/go-sdk/github/octokit/teams/item/members"
)

// ItemMembersRequestBuilder builds and executes requests for operations under \teams\{team_id}\members
type ItemMembersRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemMembersRequestBuilderGetQueryParameters **Deprecation Notice:** This endpoint route is deprecated and will be removed from the Teams API. We recommend migrating your existing code to use the new [`List team members`](https://docs.github.com/rest/teams/members#list-team-members) endpoint.Team members will include the members of child teams.
type ItemMembersRequestBuilderGetQueryParameters struct {
    // Page number of the results to fetch.
    Page *int32 `uriparametername:"page"`
    // The number of results per page (max 100).
    Per_page *int32 `uriparametername:"per_page"`
    // Filters members returned by their role in the team.
    // Deprecated: This property is deprecated, use roleAsGetRoleQueryParameterType instead
    Role *string `uriparametername:"role"`
    // Filters members returned by their role in the team.
    RoleAsGetRoleQueryParameterType *i949f6571ab906194a9314d01e930f960285837b45292e75beed92b54c60429af.GetRoleQueryParameterType `uriparametername:"role"`
}
// ItemMembersRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemMembersRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemMembersRequestBuilderGetQueryParameters
}
// ByUsername gets an item from the octokit.teams.item.members.item collection
// Deprecated: 
func (m *ItemMembersRequestBuilder) ByUsername(username string)(*ItemMembersWithUsernameItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if username != "" {
        urlTplParams["username"] = username
    }
    return NewItemMembersWithUsernameItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewItemMembersRequestBuilderInternal instantiates a new MembersRequestBuilder and sets the default values.
func NewItemMembersRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemMembersRequestBuilder) {
    m := &ItemMembersRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/teams/{team_id}/members{?role*,per_page*,page*}", pathParameters),
    }
    return m
}
// NewItemMembersRequestBuilder instantiates a new MembersRequestBuilder and sets the default values.
func NewItemMembersRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemMembersRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemMembersRequestBuilderInternal(urlParams, requestAdapter)
}
// Get **Deprecation Notice:** This endpoint route is deprecated and will be removed from the Teams API. We recommend migrating your existing code to use the new [`List team members`](https://docs.github.com/rest/teams/members#list-team-members) endpoint.Team members will include the members of child teams.
// Deprecated: 
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/teams/members#list-team-members-legacy
func (m *ItemMembersRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemMembersRequestBuilderGetRequestConfiguration)([]i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.SimpleUserable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "404": i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.CreateBasicErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.SendCollection(ctx, requestInfo, i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.CreateSimpleUserFromDiscriminatorValue, errorMapping)
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
// ToGetRequestInformation **Deprecation Notice:** This endpoint route is deprecated and will be removed from the Teams API. We recommend migrating your existing code to use the new [`List team members`](https://docs.github.com/rest/teams/members#list-team-members) endpoint.Team members will include the members of child teams.
// Deprecated: 
func (m *ItemMembersRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemMembersRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Deprecated: 
func (m *ItemMembersRequestBuilder) WithUrl(rawUrl string)(*ItemMembersRequestBuilder) {
    return NewItemMembersRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
