package repos

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6 "github.com/octokit/go-sdk/pkg/github/models"
)

// ItemItemPullsCommentsWithComment_ItemRequestBuilder builds and executes requests for operations under \repos\{repos-id}\{Owner-id}\pulls\comments\{comment_id}
type ItemItemPullsCommentsWithComment_ItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemItemPullsCommentsWithComment_ItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemItemPullsCommentsWithComment_ItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ItemItemPullsCommentsWithComment_ItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemItemPullsCommentsWithComment_ItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ItemItemPullsCommentsWithComment_ItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemItemPullsCommentsWithComment_ItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemItemPullsCommentsWithComment_ItemRequestBuilderInternal instantiates a new WithComment_ItemRequestBuilder and sets the default values.
func NewItemItemPullsCommentsWithComment_ItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemPullsCommentsWithComment_ItemRequestBuilder) {
    m := &ItemItemPullsCommentsWithComment_ItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/repos/{repos%2Did}/{Owner%2Did}/pulls/comments/{comment_id}", pathParameters),
    }
    return m
}
// NewItemItemPullsCommentsWithComment_ItemRequestBuilder instantiates a new WithComment_ItemRequestBuilder and sets the default values.
func NewItemItemPullsCommentsWithComment_ItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemPullsCommentsWithComment_ItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemItemPullsCommentsWithComment_ItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete deletes a review comment.
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/pulls/comments#delete-a-review-comment-for-a-pull-request
func (m *ItemItemPullsCommentsWithComment_ItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *ItemItemPullsCommentsWithComment_ItemRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.ToDeleteRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "404": i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CreateBasicErrorFromDiscriminatorValue,
    }
    err = m.BaseRequestBuilder.RequestAdapter.SendNoContent(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Get provides details for a review comment.
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/pulls/comments#get-a-review-comment-for-a-pull-request
func (m *ItemItemPullsCommentsWithComment_ItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemItemPullsCommentsWithComment_ItemRequestBuilderGetRequestConfiguration)(i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.PullRequestReviewCommentable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "404": i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CreateBasicErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CreatePullRequestReviewCommentFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.PullRequestReviewCommentable), nil
}
// Patch enables you to edit a review comment.
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/pulls/comments#update-a-review-comment-for-a-pull-request
func (m *ItemItemPullsCommentsWithComment_ItemRequestBuilder) Patch(ctx context.Context, body ItemItemPullsCommentsItemWithComment_PatchRequestBodyable, requestConfiguration *ItemItemPullsCommentsWithComment_ItemRequestBuilderPatchRequestConfiguration)(i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.PullRequestReviewCommentable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CreatePullRequestReviewCommentFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.PullRequestReviewCommentable), nil
}
// Reactions the reactions property
func (m *ItemItemPullsCommentsWithComment_ItemRequestBuilder) Reactions()(*ItemItemPullsCommentsItemReactionsRequestBuilder) {
    return NewItemItemPullsCommentsItemReactionsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToDeleteRequestInformation deletes a review comment.
func (m *ItemItemPullsCommentsWithComment_ItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *ItemItemPullsCommentsWithComment_ItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation provides details for a review comment.
func (m *ItemItemPullsCommentsWithComment_ItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemItemPullsCommentsWithComment_ItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToPatchRequestInformation enables you to edit a review comment.
func (m *ItemItemPullsCommentsWithComment_ItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ItemItemPullsCommentsItemWithComment_PatchRequestBodyable, requestConfiguration *ItemItemPullsCommentsWithComment_ItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
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
func (m *ItemItemPullsCommentsWithComment_ItemRequestBuilder) WithUrl(rawUrl string)(*ItemItemPullsCommentsWithComment_ItemRequestBuilder) {
    return NewItemItemPullsCommentsWithComment_ItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
