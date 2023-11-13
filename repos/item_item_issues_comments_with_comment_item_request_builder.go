package repos

import (
    "context"
    i33bcfa37677dfd8b71eaf22ad3bfd140ba1702e7a8418558499c22418bf2fef2 "github.com/octokit/go-sdk/models"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemItemIssuesCommentsWithComment_ItemRequestBuilder builds and executes requests for operations under \repos\{owner}\{repo}\issues\comments\{comment_id}
type ItemItemIssuesCommentsWithComment_ItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemItemIssuesCommentsWithComment_ItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemItemIssuesCommentsWithComment_ItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ItemItemIssuesCommentsWithComment_ItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemItemIssuesCommentsWithComment_ItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ItemItemIssuesCommentsWithComment_ItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemItemIssuesCommentsWithComment_ItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemItemIssuesCommentsWithComment_ItemRequestBuilderInternal instantiates a new WithComment_ItemRequestBuilder and sets the default values.
func NewItemItemIssuesCommentsWithComment_ItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemIssuesCommentsWithComment_ItemRequestBuilder) {
    m := &ItemItemIssuesCommentsWithComment_ItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/repos/{owner}/{repo}/issues/comments/{comment_id}", pathParameters),
    }
    return m
}
// NewItemItemIssuesCommentsWithComment_ItemRequestBuilder instantiates a new WithComment_ItemRequestBuilder and sets the default values.
func NewItemItemIssuesCommentsWithComment_ItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemIssuesCommentsWithComment_ItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemItemIssuesCommentsWithComment_ItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete you can use the REST API to delete comments on issues and pull requests. Every pull request is an issue, but not every issue is a pull request.
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/issues/comments#delete-an-issue-comment
func (m *ItemItemIssuesCommentsWithComment_ItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *ItemItemIssuesCommentsWithComment_ItemRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.ToDeleteRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    err = m.BaseRequestBuilder.RequestAdapter.SendNoContent(ctx, requestInfo, nil)
    if err != nil {
        return err
    }
    return nil
}
// Get you can use the REST API to get comments on issues and pull requests. Every pull request is an issue, but not every issue is a pull request.
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/issues/comments#get-an-issue-comment
func (m *ItemItemIssuesCommentsWithComment_ItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemItemIssuesCommentsWithComment_ItemRequestBuilderGetRequestConfiguration)(i33bcfa37677dfd8b71eaf22ad3bfd140ba1702e7a8418558499c22418bf2fef2.IssueCommentable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "404": i33bcfa37677dfd8b71eaf22ad3bfd140ba1702e7a8418558499c22418bf2fef2.CreateBasicErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i33bcfa37677dfd8b71eaf22ad3bfd140ba1702e7a8418558499c22418bf2fef2.CreateIssueCommentFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i33bcfa37677dfd8b71eaf22ad3bfd140ba1702e7a8418558499c22418bf2fef2.IssueCommentable), nil
}
// Patch you can use the REST API to update comments on issues and pull requests. Every pull request is an issue, but not every issue is a pull request.
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/issues/comments#update-an-issue-comment
func (m *ItemItemIssuesCommentsWithComment_ItemRequestBuilder) Patch(ctx context.Context, body ItemItemIssuesCommentsItemWithComment_PatchRequestBodyable, requestConfiguration *ItemItemIssuesCommentsWithComment_ItemRequestBuilderPatchRequestConfiguration)(i33bcfa37677dfd8b71eaf22ad3bfd140ba1702e7a8418558499c22418bf2fef2.IssueCommentable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "422": i33bcfa37677dfd8b71eaf22ad3bfd140ba1702e7a8418558499c22418bf2fef2.CreateValidationErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i33bcfa37677dfd8b71eaf22ad3bfd140ba1702e7a8418558499c22418bf2fef2.CreateIssueCommentFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i33bcfa37677dfd8b71eaf22ad3bfd140ba1702e7a8418558499c22418bf2fef2.IssueCommentable), nil
}
// Reactions the reactions property
func (m *ItemItemIssuesCommentsWithComment_ItemRequestBuilder) Reactions()(*ItemItemIssuesCommentsItemReactionsRequestBuilder) {
    return NewItemItemIssuesCommentsItemReactionsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToDeleteRequestInformation you can use the REST API to delete comments on issues and pull requests. Every pull request is an issue, but not every issue is a pull request.
func (m *ItemItemIssuesCommentsWithComment_ItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *ItemItemIssuesCommentsWithComment_ItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE
    return requestInfo, nil
}
// ToGetRequestInformation you can use the REST API to get comments on issues and pull requests. Every pull request is an issue, but not every issue is a pull request.
func (m *ItemItemIssuesCommentsWithComment_ItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemItemIssuesCommentsWithComment_ItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers.TryAdd("Accept", "application/json;q=1")
    return requestInfo, nil
}
// ToPatchRequestInformation you can use the REST API to update comments on issues and pull requests. Every pull request is an issue, but not every issue is a pull request.
func (m *ItemItemIssuesCommentsWithComment_ItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ItemItemIssuesCommentsItemWithComment_PatchRequestBodyable, requestConfiguration *ItemItemIssuesCommentsWithComment_ItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
    requestInfo.Headers.TryAdd("Accept", "application/json;q=1")
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
func (m *ItemItemIssuesCommentsWithComment_ItemRequestBuilder) WithUrl(rawUrl string)(*ItemItemIssuesCommentsWithComment_ItemRequestBuilder) {
    return NewItemItemIssuesCommentsWithComment_ItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
