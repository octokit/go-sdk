package repos

import (
    "context"
    i53ac87e8cb3cc9276228f74d38694a208cacb99bb8ceb705eeae99fb88d4d274 "strconv"
    i33bcfa37677dfd8b71eaf22ad3bfd140ba1702e7a8418558499c22418bf2fef2 "github.com/octokit/go-sdk/models"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i00dc8773358332755701f6884003d901f883528e958db93a2a08b1012ef42805 "github.com/octokit/go-sdk/repos/item/item/comments/item/reactions"
)

// ItemItemCommentsItemReactionsRequestBuilder builds and executes requests for operations under \repos\{owner}\{repo}\comments\{comment_id}\reactions
type ItemItemCommentsItemReactionsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemItemCommentsItemReactionsRequestBuilderGetQueryParameters list the reactions to a [commit comment](https://docs.github.com/rest/commits/comments#get-a-commit-comment).
type ItemItemCommentsItemReactionsRequestBuilderGetQueryParameters struct {
    // Returns a single [reaction type](https://docs.github.com/rest/reactions/reactions#about-reactions). Omit this parameter to list all reactions to a commit comment.
    // Deprecated: This property is deprecated, use contentAsGetContentQueryParameterType instead
    Content *string `uriparametername:"content"`
    // Returns a single [reaction type](https://docs.github.com/rest/reactions/reactions#about-reactions). Omit this parameter to list all reactions to a commit comment.
    ContentAsGetContentQueryParameterType *i00dc8773358332755701f6884003d901f883528e958db93a2a08b1012ef42805.GetContentQueryParameterType `uriparametername:"content"`
    // Page number of the results to fetch.
    Page *int32 `uriparametername:"page"`
    // The number of results per page (max 100).
    Per_page *int32 `uriparametername:"per_page"`
}
// ItemItemCommentsItemReactionsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemItemCommentsItemReactionsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemItemCommentsItemReactionsRequestBuilderGetQueryParameters
}
// ItemItemCommentsItemReactionsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemItemCommentsItemReactionsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ByReaction_id gets an item from the go-sdk.repos.item.item.comments.item.reactions.item collection
// Deprecated: This indexer is deprecated and will be removed in the next major version. Use the one with the typed parameter instead.
func (m *ItemItemCommentsItemReactionsRequestBuilder) ByReaction_id(reaction_id string)(*ItemItemCommentsItemReactionsWithReaction_ItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if reaction_id != "" {
        urlTplParams["reaction_id"] = reaction_id
    }
    return NewItemItemCommentsItemReactionsWithReaction_ItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// ByReaction_idInteger gets an item from the go-sdk.repos.item.item.comments.item.reactions.item collection
func (m *ItemItemCommentsItemReactionsRequestBuilder) ByReaction_idInteger(reaction_id int32)(*ItemItemCommentsItemReactionsWithReaction_ItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    urlTplParams["reaction_id"] = i53ac87e8cb3cc9276228f74d38694a208cacb99bb8ceb705eeae99fb88d4d274.FormatInt(int64(reaction_id), 10)
    return NewItemItemCommentsItemReactionsWithReaction_ItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewItemItemCommentsItemReactionsRequestBuilderInternal instantiates a new ReactionsRequestBuilder and sets the default values.
func NewItemItemCommentsItemReactionsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemCommentsItemReactionsRequestBuilder) {
    m := &ItemItemCommentsItemReactionsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/repos/{owner}/{repo}/comments/{comment_id}/reactions{?content*,per_page*,page*}", pathParameters),
    }
    return m
}
// NewItemItemCommentsItemReactionsRequestBuilder instantiates a new ReactionsRequestBuilder and sets the default values.
func NewItemItemCommentsItemReactionsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemCommentsItemReactionsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemItemCommentsItemReactionsRequestBuilderInternal(urlParams, requestAdapter)
}
// Get list the reactions to a [commit comment](https://docs.github.com/rest/commits/comments#get-a-commit-comment).
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/reactions/reactions#list-reactions-for-a-commit-comment
func (m *ItemItemCommentsItemReactionsRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemItemCommentsItemReactionsRequestBuilderGetRequestConfiguration)([]i33bcfa37677dfd8b71eaf22ad3bfd140ba1702e7a8418558499c22418bf2fef2.Reactionable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "404": i33bcfa37677dfd8b71eaf22ad3bfd140ba1702e7a8418558499c22418bf2fef2.CreateBasicErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.SendCollection(ctx, requestInfo, i33bcfa37677dfd8b71eaf22ad3bfd140ba1702e7a8418558499c22418bf2fef2.CreateReactionFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    val := make([]i33bcfa37677dfd8b71eaf22ad3bfd140ba1702e7a8418558499c22418bf2fef2.Reactionable, len(res))
    for i, v := range res {
        if v != nil {
            val[i] = v.(i33bcfa37677dfd8b71eaf22ad3bfd140ba1702e7a8418558499c22418bf2fef2.Reactionable)
        }
    }
    return val, nil
}
// Post create a reaction to a [commit comment](https://docs.github.com/rest/commits/comments#get-a-commit-comment). A response with an HTTP `200` status means that you already added the reaction type to this commit comment.
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/reactions/reactions#create-reaction-for-a-commit-comment
func (m *ItemItemCommentsItemReactionsRequestBuilder) Post(ctx context.Context, body ItemItemCommentsItemReactionsPostRequestBodyable, requestConfiguration *ItemItemCommentsItemReactionsRequestBuilderPostRequestConfiguration)(i33bcfa37677dfd8b71eaf22ad3bfd140ba1702e7a8418558499c22418bf2fef2.Reactionable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "422": i33bcfa37677dfd8b71eaf22ad3bfd140ba1702e7a8418558499c22418bf2fef2.CreateValidationErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i33bcfa37677dfd8b71eaf22ad3bfd140ba1702e7a8418558499c22418bf2fef2.CreateReactionFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i33bcfa37677dfd8b71eaf22ad3bfd140ba1702e7a8418558499c22418bf2fef2.Reactionable), nil
}
// ToGetRequestInformation list the reactions to a [commit comment](https://docs.github.com/rest/commits/comments#get-a-commit-comment).
func (m *ItemItemCommentsItemReactionsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemItemCommentsItemReactionsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation create a reaction to a [commit comment](https://docs.github.com/rest/commits/comments#get-a-commit-comment). A response with an HTTP `200` status means that you already added the reaction type to this commit comment.
func (m *ItemItemCommentsItemReactionsRequestBuilder) ToPostRequestInformation(ctx context.Context, body ItemItemCommentsItemReactionsPostRequestBodyable, requestConfiguration *ItemItemCommentsItemReactionsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
    requestInfo.Headers.TryAdd("Accept", "application/json;q=1")
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
func (m *ItemItemCommentsItemReactionsRequestBuilder) WithUrl(rawUrl string)(*ItemItemCommentsItemReactionsRequestBuilder) {
    return NewItemItemCommentsItemReactionsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
