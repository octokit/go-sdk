package orgs

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i000736ae6dd74f01081193e4f903216bc2bd2954ed818433b986f45d581ed035 "github.com/octokit/go-sdk/github/octokit/models"
    i1f1fe171ed5d5d45f1b30c33864c24b3379e2f44c334236ed14a16053c18b7e3 "github.com/octokit/go-sdk/github/octokit/orgs/item/teams/item/members"
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
    RoleAsGetRoleQueryParameterType *i1f1fe171ed5d5d45f1b30c33864c24b3379e2f44c334236ed14a16053c18b7e3.GetRoleQueryParameterType `uriparametername:"role"`
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
func (m *ItemTeamsItemMembersRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemTeamsItemMembersRequestBuilderGetRequestConfiguration)([]i000736ae6dd74f01081193e4f903216bc2bd2954ed818433b986f45d581ed035.SimpleUserable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.SendCollection(ctx, requestInfo, i000736ae6dd74f01081193e4f903216bc2bd2954ed818433b986f45d581ed035.CreateSimpleUserFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    val := make([]i000736ae6dd74f01081193e4f903216bc2bd2954ed818433b986f45d581ed035.SimpleUserable, len(res))
    for i, v := range res {
        if v != nil {
            val[i] = v.(i000736ae6dd74f01081193e4f903216bc2bd2954ed818433b986f45d581ed035.SimpleUserable)
        }
    }
    return val, nil
}
// ToGetRequestInformation team members will include the members of child teams.To list members in a team, the team must be visible to the authenticated user.
func (m *ItemTeamsItemMembersRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemTeamsItemMembersRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
func (m *ItemTeamsItemMembersRequestBuilder) WithUrl(rawUrl string)(*ItemTeamsItemMembersRequestBuilder) {
    return NewItemTeamsItemMembersRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
