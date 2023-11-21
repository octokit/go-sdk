package repos

import (
    "context"
    i53ac87e8cb3cc9276228f74d38694a208cacb99bb8ceb705eeae99fb88d4d274 "strconv"
    i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e "github.com/octokit/go-sdk/github/octokit/models"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i6e404a3ca2d7392b40d50aad905e1c0783f3d21e55eec49e137f3612396d9938 "github.com/octokit/go-sdk/github/octokit/repos/item/item/pulls/comments/item/reactions"
)

// ItemItemPullsCommentsItemReactionsRequestBuilder builds and executes requests for operations under \repos\{owner}\{repo}\pulls\comments\{comment_id}\reactions
type ItemItemPullsCommentsItemReactionsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemItemPullsCommentsItemReactionsRequestBuilderGetQueryParameters list the reactions to a [pull request review comment](https://docs.github.com/rest/pulls/comments#get-a-review-comment-for-a-pull-request).
type ItemItemPullsCommentsItemReactionsRequestBuilderGetQueryParameters struct {
    // Returns a single [reaction type](https://docs.github.com/rest/reactions/reactions#about-reactions). Omit this parameter to list all reactions to a pull request review comment.
    // Deprecated: This property is deprecated, use contentAsGetContentQueryParameterType instead
    Content *string `uriparametername:"content"`
    // Returns a single [reaction type](https://docs.github.com/rest/reactions/reactions#about-reactions). Omit this parameter to list all reactions to a pull request review comment.
    ContentAsGetContentQueryParameterType *i6e404a3ca2d7392b40d50aad905e1c0783f3d21e55eec49e137f3612396d9938.GetContentQueryParameterType `uriparametername:"content"`
    // Page number of the results to fetch.
    Page *int32 `uriparametername:"page"`
    // The number of results per page (max 100).
    Per_page *int32 `uriparametername:"per_page"`
}
// ItemItemPullsCommentsItemReactionsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemItemPullsCommentsItemReactionsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemItemPullsCommentsItemReactionsRequestBuilderGetQueryParameters
}
// ItemItemPullsCommentsItemReactionsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemItemPullsCommentsItemReactionsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ByReaction_id gets an item from the octokit.repos.item.item.pulls.comments.item.reactions.item collection
// Deprecated: This indexer is deprecated and will be removed in the next major version. Use the one with the typed parameter instead.
func (m *ItemItemPullsCommentsItemReactionsRequestBuilder) ByReaction_id(reaction_id string)(*ItemItemPullsCommentsItemReactionsWithReaction_ItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if reaction_id != "" {
        urlTplParams["reaction_id"] = reaction_id
    }
    return NewItemItemPullsCommentsItemReactionsWithReaction_ItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// ByReaction_idInteger gets an item from the octokit.repos.item.item.pulls.comments.item.reactions.item collection
func (m *ItemItemPullsCommentsItemReactionsRequestBuilder) ByReaction_idInteger(reaction_id int32)(*ItemItemPullsCommentsItemReactionsWithReaction_ItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    urlTplParams["reaction_id"] = i53ac87e8cb3cc9276228f74d38694a208cacb99bb8ceb705eeae99fb88d4d274.FormatInt(int64(reaction_id), 10)
    return NewItemItemPullsCommentsItemReactionsWithReaction_ItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewItemItemPullsCommentsItemReactionsRequestBuilderInternal instantiates a new ReactionsRequestBuilder and sets the default values.
func NewItemItemPullsCommentsItemReactionsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemPullsCommentsItemReactionsRequestBuilder) {
    m := &ItemItemPullsCommentsItemReactionsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/repos/{owner}/{repo}/pulls/comments/{comment_id}/reactions{?content*,per_page*,page*}", pathParameters),
    }
    return m
}
// NewItemItemPullsCommentsItemReactionsRequestBuilder instantiates a new ReactionsRequestBuilder and sets the default values.
func NewItemItemPullsCommentsItemReactionsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemPullsCommentsItemReactionsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemItemPullsCommentsItemReactionsRequestBuilderInternal(urlParams, requestAdapter)
}
// Get list the reactions to a [pull request review comment](https://docs.github.com/rest/pulls/comments#get-a-review-comment-for-a-pull-request).
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/reactions/reactions#list-reactions-for-a-pull-request-review-comment
func (m *ItemItemPullsCommentsItemReactionsRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemItemPullsCommentsItemReactionsRequestBuilderGetRequestConfiguration)([]i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.Reactionable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "404": i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.CreateBasicErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.SendCollection(ctx, requestInfo, i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.CreateReactionFromDiscriminatorValue, errorMapping)
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
// Post create a reaction to a [pull request review comment](https://docs.github.com/rest/pulls/comments#get-a-review-comment-for-a-pull-request). A response with an HTTP `200` status means that you already added the reaction type to this pull request review comment.
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/reactions/reactions#create-reaction-for-a-pull-request-review-comment
func (m *ItemItemPullsCommentsItemReactionsRequestBuilder) Post(ctx context.Context, body ItemItemPullsCommentsItemReactionsPostRequestBodyable, requestConfiguration *ItemItemPullsCommentsItemReactionsRequestBuilderPostRequestConfiguration)(i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.Reactionable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "422": i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.CreateValidationErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.CreateReactionFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.Reactionable), nil
}
// ToGetRequestInformation list the reactions to a [pull request review comment](https://docs.github.com/rest/pulls/comments#get-a-review-comment-for-a-pull-request).
func (m *ItemItemPullsCommentsItemReactionsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemItemPullsCommentsItemReactionsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation create a reaction to a [pull request review comment](https://docs.github.com/rest/pulls/comments#get-a-review-comment-for-a-pull-request). A response with an HTTP `200` status means that you already added the reaction type to this pull request review comment.
func (m *ItemItemPullsCommentsItemReactionsRequestBuilder) ToPostRequestInformation(ctx context.Context, body ItemItemPullsCommentsItemReactionsPostRequestBodyable, requestConfiguration *ItemItemPullsCommentsItemReactionsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *ItemItemPullsCommentsItemReactionsRequestBuilder) WithUrl(rawUrl string)(*ItemItemPullsCommentsItemReactionsRequestBuilder) {
    return NewItemItemPullsCommentsItemReactionsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
