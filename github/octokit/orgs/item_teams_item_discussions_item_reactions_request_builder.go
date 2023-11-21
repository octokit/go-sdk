package orgs

import (
    "context"
    i53ac87e8cb3cc9276228f74d38694a208cacb99bb8ceb705eeae99fb88d4d274 "strconv"
    i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e "github.com/octokit/go-sdk/github/octokit/models"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i7f5abcf179a2f55135a6269163bcfc47401ccddcc80c79364cb305a20e8a88de "github.com/octokit/go-sdk/github/octokit/orgs/item/teams/item/discussions/item/reactions"
)

// ItemTeamsItemDiscussionsItemReactionsRequestBuilder builds and executes requests for operations under \orgs\{org}\teams\{team_slug}\discussions\{discussion_number}\reactions
type ItemTeamsItemDiscussionsItemReactionsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemTeamsItemDiscussionsItemReactionsRequestBuilderGetQueryParameters list the reactions to a [team discussion](https://docs.github.com/rest/teams/discussions#get-a-discussion). OAuth access tokens require the `read:discussion` [scope](https://docs.github.com/apps/building-oauth-apps/understanding-scopes-for-oauth-apps/).**Note:** You can also specify a team by `org_id` and `team_id` using the route `GET /organizations/:org_id/team/:team_id/discussions/:discussion_number/reactions`.
type ItemTeamsItemDiscussionsItemReactionsRequestBuilderGetQueryParameters struct {
    // Returns a single [reaction type](https://docs.github.com/rest/reactions/reactions#about-reactions). Omit this parameter to list all reactions to a team discussion.
    // Deprecated: This property is deprecated, use contentAsGetContentQueryParameterType instead
    Content *string `uriparametername:"content"`
    // Returns a single [reaction type](https://docs.github.com/rest/reactions/reactions#about-reactions). Omit this parameter to list all reactions to a team discussion.
    ContentAsGetContentQueryParameterType *i7f5abcf179a2f55135a6269163bcfc47401ccddcc80c79364cb305a20e8a88de.GetContentQueryParameterType `uriparametername:"content"`
    // Page number of the results to fetch.
    Page *int32 `uriparametername:"page"`
    // The number of results per page (max 100).
    Per_page *int32 `uriparametername:"per_page"`
}
// ItemTeamsItemDiscussionsItemReactionsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemTeamsItemDiscussionsItemReactionsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemTeamsItemDiscussionsItemReactionsRequestBuilderGetQueryParameters
}
// ItemTeamsItemDiscussionsItemReactionsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemTeamsItemDiscussionsItemReactionsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ByReaction_id gets an item from the octokit.orgs.item.teams.item.discussions.item.reactions.item collection
// Deprecated: This indexer is deprecated and will be removed in the next major version. Use the one with the typed parameter instead.
func (m *ItemTeamsItemDiscussionsItemReactionsRequestBuilder) ByReaction_id(reaction_id string)(*ItemTeamsItemDiscussionsItemReactionsWithReaction_ItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if reaction_id != "" {
        urlTplParams["reaction_id"] = reaction_id
    }
    return NewItemTeamsItemDiscussionsItemReactionsWithReaction_ItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// ByReaction_idInteger gets an item from the octokit.orgs.item.teams.item.discussions.item.reactions.item collection
func (m *ItemTeamsItemDiscussionsItemReactionsRequestBuilder) ByReaction_idInteger(reaction_id int32)(*ItemTeamsItemDiscussionsItemReactionsWithReaction_ItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    urlTplParams["reaction_id"] = i53ac87e8cb3cc9276228f74d38694a208cacb99bb8ceb705eeae99fb88d4d274.FormatInt(int64(reaction_id), 10)
    return NewItemTeamsItemDiscussionsItemReactionsWithReaction_ItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewItemTeamsItemDiscussionsItemReactionsRequestBuilderInternal instantiates a new ReactionsRequestBuilder and sets the default values.
func NewItemTeamsItemDiscussionsItemReactionsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemTeamsItemDiscussionsItemReactionsRequestBuilder) {
    m := &ItemTeamsItemDiscussionsItemReactionsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/orgs/{org}/teams/{team_slug}/discussions/{discussion_number}/reactions{?content*,per_page*,page*}", pathParameters),
    }
    return m
}
// NewItemTeamsItemDiscussionsItemReactionsRequestBuilder instantiates a new ReactionsRequestBuilder and sets the default values.
func NewItemTeamsItemDiscussionsItemReactionsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemTeamsItemDiscussionsItemReactionsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemTeamsItemDiscussionsItemReactionsRequestBuilderInternal(urlParams, requestAdapter)
}
// Get list the reactions to a [team discussion](https://docs.github.com/rest/teams/discussions#get-a-discussion). OAuth access tokens require the `read:discussion` [scope](https://docs.github.com/apps/building-oauth-apps/understanding-scopes-for-oauth-apps/).**Note:** You can also specify a team by `org_id` and `team_id` using the route `GET /organizations/:org_id/team/:team_id/discussions/:discussion_number/reactions`.
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/reactions/reactions#list-reactions-for-a-team-discussion
func (m *ItemTeamsItemDiscussionsItemReactionsRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemTeamsItemDiscussionsItemReactionsRequestBuilderGetRequestConfiguration)([]i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.Reactionable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.SendCollection(ctx, requestInfo, i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.CreateReactionFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    val := make([]i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.Reactionable, len(res))
    for i, v := range res {
        if v != nil {
            val[i] = v.(i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.Reactionable)
        }
    }
    return val, nil
}
// Post create a reaction to a [team discussion](https://docs.github.com/rest/teams/discussions#get-a-discussion). OAuth access tokens require the `write:discussion` [scope](https://docs.github.com/apps/building-oauth-apps/understanding-scopes-for-oauth-apps/). A response with an HTTP `200` status means that you already added the reaction type to this team discussion.**Note:** You can also specify a team by `org_id` and `team_id` using the route `POST /organizations/:org_id/team/:team_id/discussions/:discussion_number/reactions`.
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/reactions/reactions#create-reaction-for-a-team-discussion
func (m *ItemTeamsItemDiscussionsItemReactionsRequestBuilder) Post(ctx context.Context, body ItemTeamsItemDiscussionsItemReactionsPostRequestBodyable, requestConfiguration *ItemTeamsItemDiscussionsItemReactionsRequestBuilderPostRequestConfiguration)(i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.Reactionable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.CreateReactionFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.Reactionable), nil
}
// ToGetRequestInformation list the reactions to a [team discussion](https://docs.github.com/rest/teams/discussions#get-a-discussion). OAuth access tokens require the `read:discussion` [scope](https://docs.github.com/apps/building-oauth-apps/understanding-scopes-for-oauth-apps/).**Note:** You can also specify a team by `org_id` and `team_id` using the route `GET /organizations/:org_id/team/:team_id/discussions/:discussion_number/reactions`.
func (m *ItemTeamsItemDiscussionsItemReactionsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemTeamsItemDiscussionsItemReactionsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation create a reaction to a [team discussion](https://docs.github.com/rest/teams/discussions#get-a-discussion). OAuth access tokens require the `write:discussion` [scope](https://docs.github.com/apps/building-oauth-apps/understanding-scopes-for-oauth-apps/). A response with an HTTP `200` status means that you already added the reaction type to this team discussion.**Note:** You can also specify a team by `org_id` and `team_id` using the route `POST /organizations/:org_id/team/:team_id/discussions/:discussion_number/reactions`.
func (m *ItemTeamsItemDiscussionsItemReactionsRequestBuilder) ToPostRequestInformation(ctx context.Context, body ItemTeamsItemDiscussionsItemReactionsPostRequestBodyable, requestConfiguration *ItemTeamsItemDiscussionsItemReactionsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
    requestInfo.Headers.TryAdd("Accept", "application/json")
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
func (m *ItemTeamsItemDiscussionsItemReactionsRequestBuilder) WithUrl(rawUrl string)(*ItemTeamsItemDiscussionsItemReactionsRequestBuilder) {
    return NewItemTeamsItemDiscussionsItemReactionsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
