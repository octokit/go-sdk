package orgs

import (
    "context"
    i33bcfa37677dfd8b71eaf22ad3bfd140ba1702e7a8418558499c22418bf2fef2 "github.com/octokit/go-sdk/models"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    if3a3ea04c37eb0319e9abad1f386089babaae23ab1f5a4346d1b1ef34f3d07ba "github.com/octokit/go-sdk/orgs/item/teams/item/members"
)

// ItemTeamsItemMembersRequestBuilder builds and executes requests for operations under \orgs\{org}\teams\{team_slug}\members
type ItemTeamsItemMembersRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemTeamsItemMembersRequestBuilderGetQueryParameters team members will include the members of child teams.To list members in a team, the team must be visible to the authenticated user.
type ItemTeamsItemMembersRequestBuilderGetQueryParameters struct {
    // Page number of the results to fetch.
    Page *int32 `uriparametername:"page"`
    // The number of results per page (max 100).
    Per_page *int32 `uriparametername:"per_page"`
    // Filters members returned by their role in the team.
    // Deprecated: This property is deprecated, use roleAsGetRoleQueryParameterType instead
    Role *string `uriparametername:"role"`
    // Filters members returned by their role in the team.
    RoleAsGetRoleQueryParameterType *if3a3ea04c37eb0319e9abad1f386089babaae23ab1f5a4346d1b1ef34f3d07ba.GetRoleQueryParameterType `uriparametername:"role"`
}
// ItemTeamsItemMembersRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemTeamsItemMembersRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemTeamsItemMembersRequestBuilderGetQueryParameters
}
// NewItemTeamsItemMembersRequestBuilderInternal instantiates a new MembersRequestBuilder and sets the default values.
func NewItemTeamsItemMembersRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemTeamsItemMembersRequestBuilder) {
    m := &ItemTeamsItemMembersRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/orgs/{org}/teams/{team_slug}/members{?role*,per_page*,page*}", pathParameters),
    }
    return m
}
// NewItemTeamsItemMembersRequestBuilder instantiates a new MembersRequestBuilder and sets the default values.
func NewItemTeamsItemMembersRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemTeamsItemMembersRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemTeamsItemMembersRequestBuilderInternal(urlParams, requestAdapter)
}
// Get team members will include the members of child teams.To list members in a team, the team must be visible to the authenticated user.
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/teams/members#list-team-members
func (m *ItemTeamsItemMembersRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemTeamsItemMembersRequestBuilderGetRequestConfiguration)([]i33bcfa37677dfd8b71eaf22ad3bfd140ba1702e7a8418558499c22418bf2fef2.SimpleUserable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.SendCollection(ctx, requestInfo, i33bcfa37677dfd8b71eaf22ad3bfd140ba1702e7a8418558499c22418bf2fef2.CreateSimpleUserFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    val := make([]i33bcfa37677dfd8b71eaf22ad3bfd140ba1702e7a8418558499c22418bf2fef2.SimpleUserable, len(res))
    for i, v := range res {
        if v != nil {
            val[i] = v.(i33bcfa37677dfd8b71eaf22ad3bfd140ba1702e7a8418558499c22418bf2fef2.SimpleUserable)
        }
    }
    return val, nil
}
// ToGetRequestInformation team members will include the members of child teams.To list members in a team, the team must be visible to the authenticated user.
func (m *ItemTeamsItemMembersRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemTeamsItemMembersRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *ItemTeamsItemMembersRequestBuilder) WithUrl(rawUrl string)(*ItemTeamsItemMembersRequestBuilder) {
    return NewItemTeamsItemMembersRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
