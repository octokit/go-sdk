package repos

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6 "github.com/octokit/go-sdk/pkg/github/models"
)

// ItemItemStatsCommit_activityRequestBuilder builds and executes requests for operations under \repos\{repos-id}\{Owner-id}\stats\commit_activity
type ItemItemStatsCommit_activityRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// NewItemItemStatsCommit_activityRequestBuilderInternal instantiates a new ItemItemStatsCommit_activityRequestBuilder and sets the default values.
func NewItemItemStatsCommit_activityRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemStatsCommit_activityRequestBuilder) {
    m := &ItemItemStatsCommit_activityRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/repos/{repos%2Did}/{Owner%2Did}/stats/commit_activity", pathParameters),
    }
    return m
}
// NewItemItemStatsCommit_activityRequestBuilder instantiates a new ItemItemStatsCommit_activityRequestBuilder and sets the default values.
func NewItemItemStatsCommit_activityRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemStatsCommit_activityRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemItemStatsCommit_activityRequestBuilderInternal(urlParams, requestAdapter)
}
// Get returns the last year of commit activity grouped by week. The `days` array is a group of commits per day, starting on `Sunday`.
// returns a []CommitActivityable when successful
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/metrics/statistics#get-the-last-year-of-commit-activity
func (m *ItemItemStatsCommit_activityRequestBuilder) Get(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])([]i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CommitActivityable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.SendCollection(ctx, requestInfo, i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CreateCommitActivityFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    val := make([]i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CommitActivityable, len(res))
    for i, v := range res {
        if v != nil {
            val[i] = v.(i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CommitActivityable)
        }
    }
    return val, nil
}
// ToGetRequestInformation returns the last year of commit activity grouped by week. The `days` array is a group of commits per day, starting on `Sunday`.
// returns a *RequestInformation when successful
func (m *ItemItemStatsCommit_activityRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ConfigureRequestInformation(requestInfo, requestConfiguration)
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemItemStatsCommit_activityRequestBuilder when successful
func (m *ItemItemStatsCommit_activityRequestBuilder) WithUrl(rawUrl string)(*ItemItemStatsCommit_activityRequestBuilder) {
    return NewItemItemStatsCommit_activityRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
