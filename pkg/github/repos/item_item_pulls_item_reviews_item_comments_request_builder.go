package repos

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6 "github.com/octokit/go-sdk/pkg/github/models"
)

// ItemItemPullsItemReviewsItemCommentsRequestBuilder builds and executes requests for operations under \repos\{repos-id}\{Owner-id}\pulls\{pull_number}\reviews\{review_id}\comments
type ItemItemPullsItemReviewsItemCommentsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemItemPullsItemReviewsItemCommentsRequestBuilderGetQueryParameters list comments for a specific pull request review.
type ItemItemPullsItemReviewsItemCommentsRequestBuilderGetQueryParameters struct {
    // Page number of the results to fetch.
    Page *int32 `uriparametername:"page"`
    // The number of results per page (max 100).
    Per_page *int32 `uriparametername:"per_page"`
}
// ItemItemPullsItemReviewsItemCommentsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemItemPullsItemReviewsItemCommentsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemItemPullsItemReviewsItemCommentsRequestBuilderGetQueryParameters
}
// NewItemItemPullsItemReviewsItemCommentsRequestBuilderInternal instantiates a new CommentsRequestBuilder and sets the default values.
func NewItemItemPullsItemReviewsItemCommentsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemPullsItemReviewsItemCommentsRequestBuilder) {
    m := &ItemItemPullsItemReviewsItemCommentsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/repos/{repos%2Did}/{Owner%2Did}/pulls/{pull_number}/reviews/{review_id}/comments{?per_page*,page*}", pathParameters),
    }
    return m
}
// NewItemItemPullsItemReviewsItemCommentsRequestBuilder instantiates a new CommentsRequestBuilder and sets the default values.
func NewItemItemPullsItemReviewsItemCommentsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemPullsItemReviewsItemCommentsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemItemPullsItemReviewsItemCommentsRequestBuilderInternal(urlParams, requestAdapter)
}
// Get list comments for a specific pull request review.
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/pulls/reviews#list-comments-for-a-pull-request-review
func (m *ItemItemPullsItemReviewsItemCommentsRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemItemPullsItemReviewsItemCommentsRequestBuilderGetRequestConfiguration)([]i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.ReviewCommentable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "404": i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CreateBasicErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.SendCollection(ctx, requestInfo, i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CreateReviewCommentFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    val := make([]i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.ReviewCommentable, len(res))
    for i, v := range res {
        if v != nil {
            val[i] = v.(i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.ReviewCommentable)
        }
    }
    return val, nil
}
// ToGetRequestInformation list comments for a specific pull request review.
func (m *ItemItemPullsItemReviewsItemCommentsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemItemPullsItemReviewsItemCommentsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *ItemItemPullsItemReviewsItemCommentsRequestBuilder) WithUrl(rawUrl string)(*ItemItemPullsItemReviewsItemCommentsRequestBuilder) {
    return NewItemItemPullsItemReviewsItemCommentsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
