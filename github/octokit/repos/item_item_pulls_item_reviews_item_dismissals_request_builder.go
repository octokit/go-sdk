package repos

import (
    "context"
    i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e "github.com/octokit/go-sdk/github/octokit/models"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemItemPullsItemReviewsItemDismissalsRequestBuilder builds and executes requests for operations under \repos\{owner}\{repo}\pulls\{pull_number}\reviews\{review_id}\dismissals
type ItemItemPullsItemReviewsItemDismissalsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemItemPullsItemReviewsItemDismissalsRequestBuilderPutRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemItemPullsItemReviewsItemDismissalsRequestBuilderPutRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemItemPullsItemReviewsItemDismissalsRequestBuilderInternal instantiates a new DismissalsRequestBuilder and sets the default values.
func NewItemItemPullsItemReviewsItemDismissalsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemPullsItemReviewsItemDismissalsRequestBuilder) {
    m := &ItemItemPullsItemReviewsItemDismissalsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/repos/{owner}/{repo}/pulls/{pull_number}/reviews/{review_id}/dismissals", pathParameters),
    }
    return m
}
// NewItemItemPullsItemReviewsItemDismissalsRequestBuilder instantiates a new DismissalsRequestBuilder and sets the default values.
func NewItemItemPullsItemReviewsItemDismissalsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemPullsItemReviewsItemDismissalsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemItemPullsItemReviewsItemDismissalsRequestBuilderInternal(urlParams, requestAdapter)
}
// Put **Note:** To dismiss a pull request review on a [protected branch](https://docs.github.com/rest/branches/branch-protection), you must be a repository administrator or be included in the list of people or teams who can dismiss pull request reviews.
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/pulls/reviews#dismiss-a-review-for-a-pull-request
func (m *ItemItemPullsItemReviewsItemDismissalsRequestBuilder) Put(ctx context.Context, body ItemItemPullsItemReviewsItemDismissalsPutRequestBodyable, requestConfiguration *ItemItemPullsItemReviewsItemDismissalsRequestBuilderPutRequestConfiguration)(i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.PullRequestReviewable, error) {
    requestInfo, err := m.ToPutRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "404": i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.CreateBasicErrorFromDiscriminatorValue,
        "422": i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.CreateValidationErrorSimpleFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.CreatePullRequestReviewFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.PullRequestReviewable), nil
}
// ToPutRequestInformation **Note:** To dismiss a pull request review on a [protected branch](https://docs.github.com/rest/branches/branch-protection), you must be a repository administrator or be included in the list of people or teams who can dismiss pull request reviews.
func (m *ItemItemPullsItemReviewsItemDismissalsRequestBuilder) ToPutRequestInformation(ctx context.Context, body ItemItemPullsItemReviewsItemDismissalsPutRequestBodyable, requestConfiguration *ItemItemPullsItemReviewsItemDismissalsRequestBuilderPutRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PUT
    requestInfo.Headers.TryAdd("Accept", "application/json;q=1")
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
func (m *ItemItemPullsItemReviewsItemDismissalsRequestBuilder) WithUrl(rawUrl string)(*ItemItemPullsItemReviewsItemDismissalsRequestBuilder) {
    return NewItemItemPullsItemReviewsItemDismissalsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
