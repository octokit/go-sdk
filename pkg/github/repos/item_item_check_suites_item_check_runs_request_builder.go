package repos

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i81721daf4afe64f77b2b79b27ec44ce9325b9e20e6f9737c000fdcbda4772f33 "github.com/octokit/go-sdk/pkg/github/repos/item/item/checksuites/item/checkruns"
)

// ItemItemCheckSuitesItemCheckRunsRequestBuilder builds and executes requests for operations under \repos\{repos-id}\{Owner-id}\check-suites\{check_suite_id}\check-runs
type ItemItemCheckSuitesItemCheckRunsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemItemCheckSuitesItemCheckRunsRequestBuilderGetQueryParameters lists check runs for a check suite using its `id`.GitHub Apps must have the `checks:read` permission on a private repository or pull access to a public repository to get check runs. OAuth apps and authenticated users must have the `repo` scope to get check runs in a private repository.**Note:** The endpoints to manage checks only look for pushes in the repository where the check suite or check run were created. Pushes to a branch in a forked repository are not detected and return an empty `pull_requests` array.
type ItemItemCheckSuitesItemCheckRunsRequestBuilderGetQueryParameters struct {
    // Returns check runs with the specified `name`.
    Check_name *string `uriparametername:"check_name"`
    // Filters check runs by their `completed_at` timestamp. `latest` returns the most recent check runs.
    Filter *i81721daf4afe64f77b2b79b27ec44ce9325b9e20e6f9737c000fdcbda4772f33.GetFilterQueryParameterType `uriparametername:"filter"`
    // The page number of the results to fetch. For more information, see "[Using pagination in the REST API](https://docs.github.com/rest/using-the-rest-api/using-pagination-in-the-rest-api)."
    Page *int32 `uriparametername:"page"`
    // The number of results per page (max 100). For more information, see "[Using pagination in the REST API](https://docs.github.com/rest/using-the-rest-api/using-pagination-in-the-rest-api)."
    Per_page *int32 `uriparametername:"per_page"`
    // Returns check runs with the specified `status`.
    Status *i81721daf4afe64f77b2b79b27ec44ce9325b9e20e6f9737c000fdcbda4772f33.GetStatusQueryParameterType `uriparametername:"status"`
}
// NewItemItemCheckSuitesItemCheckRunsRequestBuilderInternal instantiates a new CheckRunsRequestBuilder and sets the default values.
func NewItemItemCheckSuitesItemCheckRunsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemCheckSuitesItemCheckRunsRequestBuilder) {
    m := &ItemItemCheckSuitesItemCheckRunsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/repos/{repos%2Did}/{Owner%2Did}/check-suites/{check_suite_id}/check-runs{?check_name*,filter*,page*,per_page*,status*}", pathParameters),
    }
    return m
}
// NewItemItemCheckSuitesItemCheckRunsRequestBuilder instantiates a new CheckRunsRequestBuilder and sets the default values.
func NewItemItemCheckSuitesItemCheckRunsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemCheckSuitesItemCheckRunsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemItemCheckSuitesItemCheckRunsRequestBuilderInternal(urlParams, requestAdapter)
}
// Get lists check runs for a check suite using its `id`.GitHub Apps must have the `checks:read` permission on a private repository or pull access to a public repository to get check runs. OAuth apps and authenticated users must have the `repo` scope to get check runs in a private repository.**Note:** The endpoints to manage checks only look for pushes in the repository where the check suite or check run were created. Pushes to a branch in a forked repository are not detected and return an empty `pull_requests` array.
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/checks/runs#list-check-runs-in-a-check-suite
func (m *ItemItemCheckSuitesItemCheckRunsRequestBuilder) Get(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[ItemItemCheckSuitesItemCheckRunsRequestBuilderGetQueryParameters])(ItemItemCheckSuitesItemCheckRunsGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateItemItemCheckSuitesItemCheckRunsGetResponseFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ItemItemCheckSuitesItemCheckRunsGetResponseable), nil
}
// ToGetRequestInformation lists check runs for a check suite using its `id`.GitHub Apps must have the `checks:read` permission on a private repository or pull access to a public repository to get check runs. OAuth apps and authenticated users must have the `repo` scope to get check runs in a private repository.**Note:** The endpoints to manage checks only look for pushes in the repository where the check suite or check run were created. Pushes to a branch in a forked repository are not detected and return an empty `pull_requests` array.
func (m *ItemItemCheckSuitesItemCheckRunsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[ItemItemCheckSuitesItemCheckRunsRequestBuilderGetQueryParameters])(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ConfigureRequestInformation(requestInfo, requestConfiguration)
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
func (m *ItemItemCheckSuitesItemCheckRunsRequestBuilder) WithUrl(rawUrl string)(*ItemItemCheckSuitesItemCheckRunsRequestBuilder) {
    return NewItemItemCheckSuitesItemCheckRunsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
