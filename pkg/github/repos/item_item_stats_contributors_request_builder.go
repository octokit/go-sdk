package repos

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6 "github.com/octokit/go-sdk/pkg/github/models"
)

// ItemItemStatsContributorsRequestBuilder builds and executes requests for operations under \repos\{repos-id}\{Owner-id}\stats\contributors
type ItemItemStatsContributorsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemItemStatsContributorsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemItemStatsContributorsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemItemStatsContributorsRequestBuilderInternal instantiates a new ContributorsRequestBuilder and sets the default values.
func NewItemItemStatsContributorsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemStatsContributorsRequestBuilder) {
    m := &ItemItemStatsContributorsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/repos/{repos%2Did}/{Owner%2Did}/stats/contributors", pathParameters),
    }
    return m
}
// NewItemItemStatsContributorsRequestBuilder instantiates a new ContributorsRequestBuilder and sets the default values.
func NewItemItemStatsContributorsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemStatsContributorsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemItemStatsContributorsRequestBuilderInternal(urlParams, requestAdapter)
}
// Get returns the `total` number of commits authored by the contributor. In addition, the response includes a Weekly Hash (`weeks` array) with the following information:*   `w` - Start of the week, given as a [Unix timestamp](https://en.wikipedia.org/wiki/Unix_time).*   `a` - Number of additions*   `d` - Number of deletions*   `c` - Number of commits**Note:** This endpoint will return `0` values for all addition and deletion counts in repositories with 10,000 or more commits.
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/metrics/statistics#get-all-contributor-commit-activity
func (m *ItemItemStatsContributorsRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemItemStatsContributorsRequestBuilderGetRequestConfiguration)([]i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.ContributorActivityable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.SendCollection(ctx, requestInfo, i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CreateContributorActivityFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    val := make([]i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.ContributorActivityable, len(res))
    for i, v := range res {
        if v != nil {
            val[i] = v.(i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.ContributorActivityable)
        }
    }
    return val, nil
}
// ToGetRequestInformation returns the `total` number of commits authored by the contributor. In addition, the response includes a Weekly Hash (`weeks` array) with the following information:*   `w` - Start of the week, given as a [Unix timestamp](https://en.wikipedia.org/wiki/Unix_time).*   `a` - Number of additions*   `d` - Number of deletions*   `c` - Number of commits**Note:** This endpoint will return `0` values for all addition and deletion counts in repositories with 10,000 or more commits.
func (m *ItemItemStatsContributorsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemItemStatsContributorsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
func (m *ItemItemStatsContributorsRequestBuilder) WithUrl(rawUrl string)(*ItemItemStatsContributorsRequestBuilder) {
    return NewItemItemStatsContributorsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
