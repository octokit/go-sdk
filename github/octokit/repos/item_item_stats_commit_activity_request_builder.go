package repos

import (
    "context"
    i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e "github.com/octokit/go-sdk/github/octokit/models"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemItemStatsCommit_activityRequestBuilder builds and executes requests for operations under \repos\{owner}\{repo}\stats\commit_activity
type ItemItemStatsCommit_activityRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemItemStatsCommit_activityRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemItemStatsCommit_activityRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemItemStatsCommit_activityRequestBuilderInternal instantiates a new Commit_activityRequestBuilder and sets the default values.
func NewItemItemStatsCommit_activityRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemStatsCommit_activityRequestBuilder) {
    m := &ItemItemStatsCommit_activityRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/repos/{owner}/{repo}/stats/commit_activity", pathParameters),
    }
    return m
}
// NewItemItemStatsCommit_activityRequestBuilder instantiates a new Commit_activityRequestBuilder and sets the default values.
func NewItemItemStatsCommit_activityRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemStatsCommit_activityRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemItemStatsCommit_activityRequestBuilderInternal(urlParams, requestAdapter)
}
// Get returns the last year of commit activity grouped by week. The `days` array is a group of commits per day, starting on `Sunday`.
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/metrics/statistics#get-the-last-year-of-commit-activity
func (m *ItemItemStatsCommit_activityRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemItemStatsCommit_activityRequestBuilderGetRequestConfiguration)([]i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.CommitActivityable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.SendCollection(ctx, requestInfo, i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.CreateCommitActivityFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    val := make([]i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.CommitActivityable, len(res))
    for i, v := range res {
        if v != nil {
            val[i] = v.(i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.CommitActivityable)
        }
    }
    return val, nil
}
// ToGetRequestInformation returns the last year of commit activity grouped by week. The `days` array is a group of commits per day, starting on `Sunday`.
func (m *ItemItemStatsCommit_activityRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemItemStatsCommit_activityRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
func (m *ItemItemStatsCommit_activityRequestBuilder) WithUrl(rawUrl string)(*ItemItemStatsCommit_activityRequestBuilder) {
    return NewItemItemStatsCommit_activityRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
