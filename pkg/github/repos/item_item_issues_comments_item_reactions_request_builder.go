package repos

import (
    "context"
    i53ac87e8cb3cc9276228f74d38694a208cacb99bb8ceb705eeae99fb88d4d274 "strconv"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6 "github.com/octokit/go-sdk/pkg/github/models"
    i4962a446a8e590313347ef7c73ec2d481be83dcd51a243c337358fc09ae22a4e "github.com/octokit/go-sdk/pkg/github/repos/item/item/issues/comments/item/reactions"
)

// ItemItemIssuesCommentsItemReactionsRequestBuilder builds and executes requests for operations under \repos\{repos-id}\{Owner-id}\issues\comments\{comment_id}\reactions
type ItemItemIssuesCommentsItemReactionsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemItemIssuesCommentsItemReactionsRequestBuilderGetQueryParameters list the reactions to an [issue comment](https://docs.github.com/rest/issues/comments#get-an-issue-comment).
type ItemItemIssuesCommentsItemReactionsRequestBuilderGetQueryParameters struct {
    // Returns a single [reaction type](https://docs.github.com/rest/reactions/reactions#about-reactions). Omit this parameter to list all reactions to an issue comment.
    // Deprecated: This property is deprecated, use contentAsGetContentQueryParameterType instead
    Content *string `uriparametername:"content"`
    // Returns a single [reaction type](https://docs.github.com/rest/reactions/reactions#about-reactions). Omit this parameter to list all reactions to an issue comment.
    ContentAsGetContentQueryParameterType *i4962a446a8e590313347ef7c73ec2d481be83dcd51a243c337358fc09ae22a4e.GetContentQueryParameterType `uriparametername:"content"`
    // Page number of the results to fetch.
    Page *int32 `uriparametername:"page"`
    // The number of results per page (max 100).
    Per_page *int32 `uriparametername:"per_page"`
}
// ItemItemIssuesCommentsItemReactionsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemItemIssuesCommentsItemReactionsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemItemIssuesCommentsItemReactionsRequestBuilderGetQueryParameters
}
// ItemItemIssuesCommentsItemReactionsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemItemIssuesCommentsItemReactionsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ByReaction_id gets an item from the github.com/octokit/go-sdk/pkg/github/.repos.item.item.issues.comments.item.reactions.item collection
// Deprecated: This indexer is deprecated and will be removed in the next major version. Use the one with the typed parameter instead.
func (m *ItemItemIssuesCommentsItemReactionsRequestBuilder) ByReaction_id(reaction_id string)(*ItemItemIssuesCommentsItemReactionsWithReaction_ItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if reaction_id != "" {
        urlTplParams["reaction_id"] = reaction_id
    }
    return NewItemItemIssuesCommentsItemReactionsWithReaction_ItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// ByReaction_idInteger gets an item from the github.com/octokit/go-sdk/pkg/github/.repos.item.item.issues.comments.item.reactions.item collection
func (m *ItemItemIssuesCommentsItemReactionsRequestBuilder) ByReaction_idInteger(reaction_id int32)(*ItemItemIssuesCommentsItemReactionsWithReaction_ItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    urlTplParams["reaction_id"] = i53ac87e8cb3cc9276228f74d38694a208cacb99bb8ceb705eeae99fb88d4d274.FormatInt(int64(reaction_id), 10)
    return NewItemItemIssuesCommentsItemReactionsWithReaction_ItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewItemItemIssuesCommentsItemReactionsRequestBuilderInternal instantiates a new ReactionsRequestBuilder and sets the default values.
func NewItemItemIssuesCommentsItemReactionsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemIssuesCommentsItemReactionsRequestBuilder) {
    m := &ItemItemIssuesCommentsItemReactionsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/repos/{repos%2Did}/{Owner%2Did}/issues/comments/{comment_id}/reactions{?content*,per_page*,page*}", pathParameters),
    }
    return m
}
// NewItemItemIssuesCommentsItemReactionsRequestBuilder instantiates a new ReactionsRequestBuilder and sets the default values.
func NewItemItemIssuesCommentsItemReactionsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemIssuesCommentsItemReactionsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemItemIssuesCommentsItemReactionsRequestBuilderInternal(urlParams, requestAdapter)
}
// Get list the reactions to an [issue comment](https://docs.github.com/rest/issues/comments#get-an-issue-comment).
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/reactions/reactions#list-reactions-for-an-issue-comment
func (m *ItemItemIssuesCommentsItemReactionsRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemItemIssuesCommentsItemReactionsRequestBuilderGetRequestConfiguration)([]i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.Reactionable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "404": i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CreateBasicErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.SendCollection(ctx, requestInfo, i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CreateReactionFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    val := make([]i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.Reactionable, len(res))
    for i, v := range res {
        if v != nil {
            val[i] = v.(i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.Reactionable)
        }
    }
    return val, nil
}
// Post create a reaction to an [issue comment](https://docs.github.com/rest/issues/comments#get-an-issue-comment). A response with an HTTP `200` status means that you already added the reaction type to this issue comment.
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/reactions/reactions#create-reaction-for-an-issue-comment
func (m *ItemItemIssuesCommentsItemReactionsRequestBuilder) Post(ctx context.Context, body ItemItemIssuesCommentsItemReactionsPostRequestBodyable, requestConfiguration *ItemItemIssuesCommentsItemReactionsRequestBuilderPostRequestConfiguration)(i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.Reactionable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "422": i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CreateValidationErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CreateReactionFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.Reactionable), nil
}
// ToGetRequestInformation list the reactions to an [issue comment](https://docs.github.com/rest/issues/comments#get-an-issue-comment).
func (m *ItemItemIssuesCommentsItemReactionsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemItemIssuesCommentsItemReactionsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation create a reaction to an [issue comment](https://docs.github.com/rest/issues/comments#get-an-issue-comment). A response with an HTTP `200` status means that you already added the reaction type to this issue comment.
func (m *ItemItemIssuesCommentsItemReactionsRequestBuilder) ToPostRequestInformation(ctx context.Context, body ItemItemIssuesCommentsItemReactionsPostRequestBodyable, requestConfiguration *ItemItemIssuesCommentsItemReactionsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
func (m *ItemItemIssuesCommentsItemReactionsRequestBuilder) WithUrl(rawUrl string)(*ItemItemIssuesCommentsItemReactionsRequestBuilder) {
    return NewItemItemIssuesCommentsItemReactionsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
