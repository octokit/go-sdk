package teams

import (
    "context"
    i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e "github.com/octokit/go-sdk/github/octokit/models"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i7e5e1db7d8027c2959637d9fcad216ac57d9c6a5d4fc194d579a8dbcc9cf8d32 "github.com/octokit/go-sdk/github/octokit/teams/item/discussions/item/comments/item/reactions"
)

// ItemDiscussionsItemCommentsItemReactionsRequestBuilder builds and executes requests for operations under \teams\{team_id}\discussions\{discussion_number}\comments\{comment_number}\reactions
type ItemDiscussionsItemCommentsItemReactionsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemDiscussionsItemCommentsItemReactionsRequestBuilderGetQueryParameters **Deprecation Notice:** This endpoint route is deprecated and will be removed from the Teams API. We recommend migrating your existing code to use the new [`List reactions for a team discussion comment`](https://docs.github.com/rest/reactions/reactions#list-reactions-for-a-team-discussion-comment) endpoint.List the reactions to a [team discussion comment](https://docs.github.com/rest/teams/discussion-comments#get-a-discussion-comment). OAuth access tokens require the `read:discussion` [scope](https://docs.github.com/apps/building-oauth-apps/understanding-scopes-for-oauth-apps/).
type ItemDiscussionsItemCommentsItemReactionsRequestBuilderGetQueryParameters struct {
    // Returns a single [reaction type](https://docs.github.com/rest/reactions/reactions#about-reactions). Omit this parameter to list all reactions to a team discussion comment.
    // Deprecated: This property is deprecated, use contentAsGetContentQueryParameterType instead
    Content *string `uriparametername:"content"`
    // Returns a single [reaction type](https://docs.github.com/rest/reactions/reactions#about-reactions). Omit this parameter to list all reactions to a team discussion comment.
    ContentAsGetContentQueryParameterType *i7e5e1db7d8027c2959637d9fcad216ac57d9c6a5d4fc194d579a8dbcc9cf8d32.GetContentQueryParameterType `uriparametername:"content"`
    // Page number of the results to fetch.
    Page *int32 `uriparametername:"page"`
    // The number of results per page (max 100).
    Per_page *int32 `uriparametername:"per_page"`
}
// ItemDiscussionsItemCommentsItemReactionsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemDiscussionsItemCommentsItemReactionsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemDiscussionsItemCommentsItemReactionsRequestBuilderGetQueryParameters
}
// ItemDiscussionsItemCommentsItemReactionsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemDiscussionsItemCommentsItemReactionsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemDiscussionsItemCommentsItemReactionsRequestBuilderInternal instantiates a new ReactionsRequestBuilder and sets the default values.
func NewItemDiscussionsItemCommentsItemReactionsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemDiscussionsItemCommentsItemReactionsRequestBuilder) {
    m := &ItemDiscussionsItemCommentsItemReactionsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/teams/{team_id}/discussions/{discussion_number}/comments/{comment_number}/reactions{?content*,per_page*,page*}", pathParameters),
    }
    return m
}
// NewItemDiscussionsItemCommentsItemReactionsRequestBuilder instantiates a new ReactionsRequestBuilder and sets the default values.
func NewItemDiscussionsItemCommentsItemReactionsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemDiscussionsItemCommentsItemReactionsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemDiscussionsItemCommentsItemReactionsRequestBuilderInternal(urlParams, requestAdapter)
}
// Get **Deprecation Notice:** This endpoint route is deprecated and will be removed from the Teams API. We recommend migrating your existing code to use the new [`List reactions for a team discussion comment`](https://docs.github.com/rest/reactions/reactions#list-reactions-for-a-team-discussion-comment) endpoint.List the reactions to a [team discussion comment](https://docs.github.com/rest/teams/discussion-comments#get-a-discussion-comment). OAuth access tokens require the `read:discussion` [scope](https://docs.github.com/apps/building-oauth-apps/understanding-scopes-for-oauth-apps/).
// Deprecated: 
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/reactions/reactions#list-reactions-for-a-team-discussion-comment-legacy
func (m *ItemDiscussionsItemCommentsItemReactionsRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemDiscussionsItemCommentsItemReactionsRequestBuilderGetRequestConfiguration)([]i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.Reactionable, error) {
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
// Post **Deprecation Notice:** This endpoint route is deprecated and will be removed from the Teams API. We recommend migrating your existing code to use the new "[Create reaction for a team discussion comment](https://docs.github.com/rest/reactions/reactions#create-reaction-for-a-team-discussion-comment)" endpoint.Create a reaction to a [team discussion comment](https://docs.github.com/rest/teams/discussion-comments#get-a-discussion-comment). OAuth access tokens require the `write:discussion` [scope](https://docs.github.com/apps/building-oauth-apps/understanding-scopes-for-oauth-apps/). A response with an HTTP `200` status means that you already added the reaction type to this team discussion comment.
// Deprecated: 
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/reactions/reactions#create-reaction-for-a-team-discussion-comment-legacy
func (m *ItemDiscussionsItemCommentsItemReactionsRequestBuilder) Post(ctx context.Context, body ItemDiscussionsItemCommentsItemReactionsPostRequestBodyable, requestConfiguration *ItemDiscussionsItemCommentsItemReactionsRequestBuilderPostRequestConfiguration)(i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.Reactionable, error) {
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
// ToGetRequestInformation **Deprecation Notice:** This endpoint route is deprecated and will be removed from the Teams API. We recommend migrating your existing code to use the new [`List reactions for a team discussion comment`](https://docs.github.com/rest/reactions/reactions#list-reactions-for-a-team-discussion-comment) endpoint.List the reactions to a [team discussion comment](https://docs.github.com/rest/teams/discussion-comments#get-a-discussion-comment). OAuth access tokens require the `read:discussion` [scope](https://docs.github.com/apps/building-oauth-apps/understanding-scopes-for-oauth-apps/).
// Deprecated: 
func (m *ItemDiscussionsItemCommentsItemReactionsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemDiscussionsItemCommentsItemReactionsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation **Deprecation Notice:** This endpoint route is deprecated and will be removed from the Teams API. We recommend migrating your existing code to use the new "[Create reaction for a team discussion comment](https://docs.github.com/rest/reactions/reactions#create-reaction-for-a-team-discussion-comment)" endpoint.Create a reaction to a [team discussion comment](https://docs.github.com/rest/teams/discussion-comments#get-a-discussion-comment). OAuth access tokens require the `write:discussion` [scope](https://docs.github.com/apps/building-oauth-apps/understanding-scopes-for-oauth-apps/). A response with an HTTP `200` status means that you already added the reaction type to this team discussion comment.
// Deprecated: 
func (m *ItemDiscussionsItemCommentsItemReactionsRequestBuilder) ToPostRequestInformation(ctx context.Context, body ItemDiscussionsItemCommentsItemReactionsPostRequestBodyable, requestConfiguration *ItemDiscussionsItemCommentsItemReactionsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Deprecated: 
func (m *ItemDiscussionsItemCommentsItemReactionsRequestBuilder) WithUrl(rawUrl string)(*ItemDiscussionsItemCommentsItemReactionsRequestBuilder) {
    return NewItemDiscussionsItemCommentsItemReactionsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
