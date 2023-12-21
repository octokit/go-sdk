package repos

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ifb385791f8cac3dd5aa968a1ebb84b55fed8efbc21a8f661021ab21ba709ef32 "github.com/octokit/go-sdk/pkg/github/repos/item/item/commits/item/checkruns"
)

// ItemItemCommitsItemCheckRunsRequestBuilder builds and executes requests for operations under \repos\{repos-id}\{Owner-id}\commits\{commits-id}\check-runs
type ItemItemCommitsItemCheckRunsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemItemCommitsItemCheckRunsRequestBuilderGetQueryParameters lists check runs for a commit ref. The `ref` can be a SHA, branch name, or a tag name.GitHub Apps must have the `checks:read` permission on a private repository or pull access to a public repository to get check runs. OAuth apps and authenticated users must have the `repo` scope to get check runs in a private repository.**Note:** The endpoints to manage checks only look for pushes in the repository where the check suite or check run were created. Pushes to a branch in a forked repository are not detected and return an empty `pull_requests` array.If there are more than 1000 check suites on a single git reference, this endpoint will limit check runs to the 1000 most recent check suites. To iterate over all possible check runs, use the [List check suites for a Git reference](https://docs.github.com/rest/reference/checks#list-check-suites-for-a-git-reference) endpoint and provide the `check_suite_id` parameter to the [List check runs in a check suite](https://docs.github.com/rest/reference/checks#list-check-runs-in-a-check-suite) endpoint.
type ItemItemCommitsItemCheckRunsRequestBuilderGetQueryParameters struct {
    // 
    App_id *int32 `uriparametername:"app_id"`
    // Returns check runs with the specified `name`.
    Check_name *string `uriparametername:"check_name"`
    // Filters check runs by their `completed_at` timestamp. `latest` returns the most recent check runs.
    // Deprecated: This property is deprecated, use filterAsGetFilterQueryParameterType instead
    Filter *string `uriparametername:"filter"`
    // Filters check runs by their `completed_at` timestamp. `latest` returns the most recent check runs.
    FilterAsGetFilterQueryParameterType *ifb385791f8cac3dd5aa968a1ebb84b55fed8efbc21a8f661021ab21ba709ef32.GetFilterQueryParameterType `uriparametername:"filter"`
    // Page number of the results to fetch.
    Page *int32 `uriparametername:"page"`
    // The number of results per page (max 100).
    Per_page *int32 `uriparametername:"per_page"`
    // Returns check runs with the specified `status`.
    // Deprecated: This property is deprecated, use statusAsGetStatusQueryParameterType instead
    Status *string `uriparametername:"status"`
    // Returns check runs with the specified `status`.
    StatusAsGetStatusQueryParameterType *ifb385791f8cac3dd5aa968a1ebb84b55fed8efbc21a8f661021ab21ba709ef32.GetStatusQueryParameterType `uriparametername:"status"`
}
// ItemItemCommitsItemCheckRunsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemItemCommitsItemCheckRunsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemItemCommitsItemCheckRunsRequestBuilderGetQueryParameters
}
// NewItemItemCommitsItemCheckRunsRequestBuilderInternal instantiates a new CheckRunsRequestBuilder and sets the default values.
func NewItemItemCommitsItemCheckRunsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemCommitsItemCheckRunsRequestBuilder) {
    m := &ItemItemCommitsItemCheckRunsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/repos/{repos%2Did}/{Owner%2Did}/commits/{commits%2Did}/check-runs{?check_name*,status*,filter*,per_page*,page*,app_id*}", pathParameters),
    }
    return m
}
// NewItemItemCommitsItemCheckRunsRequestBuilder instantiates a new CheckRunsRequestBuilder and sets the default values.
func NewItemItemCommitsItemCheckRunsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemCommitsItemCheckRunsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemItemCommitsItemCheckRunsRequestBuilderInternal(urlParams, requestAdapter)
}
// Get lists check runs for a commit ref. The `ref` can be a SHA, branch name, or a tag name.GitHub Apps must have the `checks:read` permission on a private repository or pull access to a public repository to get check runs. OAuth apps and authenticated users must have the `repo` scope to get check runs in a private repository.**Note:** The endpoints to manage checks only look for pushes in the repository where the check suite or check run were created. Pushes to a branch in a forked repository are not detected and return an empty `pull_requests` array.If there are more than 1000 check suites on a single git reference, this endpoint will limit check runs to the 1000 most recent check suites. To iterate over all possible check runs, use the [List check suites for a Git reference](https://docs.github.com/rest/reference/checks#list-check-suites-for-a-git-reference) endpoint and provide the `check_suite_id` parameter to the [List check runs in a check suite](https://docs.github.com/rest/reference/checks#list-check-runs-in-a-check-suite) endpoint.
// Deprecated: This method is obsolete. Use GetAsCheckRunsGetResponse instead.
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/checks/runs#list-check-runs-for-a-git-reference
func (m *ItemItemCommitsItemCheckRunsRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemItemCommitsItemCheckRunsRequestBuilderGetRequestConfiguration)(ItemItemCommitsItemCheckRunsResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateItemItemCommitsItemCheckRunsResponseFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ItemItemCommitsItemCheckRunsResponseable), nil
}
// GetAsCheckRunsGetResponse lists check runs for a commit ref. The `ref` can be a SHA, branch name, or a tag name.GitHub Apps must have the `checks:read` permission on a private repository or pull access to a public repository to get check runs. OAuth apps and authenticated users must have the `repo` scope to get check runs in a private repository.**Note:** The endpoints to manage checks only look for pushes in the repository where the check suite or check run were created. Pushes to a branch in a forked repository are not detected and return an empty `pull_requests` array.If there are more than 1000 check suites on a single git reference, this endpoint will limit check runs to the 1000 most recent check suites. To iterate over all possible check runs, use the [List check suites for a Git reference](https://docs.github.com/rest/reference/checks#list-check-suites-for-a-git-reference) endpoint and provide the `check_suite_id` parameter to the [List check runs in a check suite](https://docs.github.com/rest/reference/checks#list-check-runs-in-a-check-suite) endpoint.
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/checks/runs#list-check-runs-for-a-git-reference
func (m *ItemItemCommitsItemCheckRunsRequestBuilder) GetAsCheckRunsGetResponse(ctx context.Context, requestConfiguration *ItemItemCommitsItemCheckRunsRequestBuilderGetRequestConfiguration)(ItemItemCommitsItemCheckRunsGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateItemItemCommitsItemCheckRunsGetResponseFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ItemItemCommitsItemCheckRunsGetResponseable), nil
}
// ToGetRequestInformation lists check runs for a commit ref. The `ref` can be a SHA, branch name, or a tag name.GitHub Apps must have the `checks:read` permission on a private repository or pull access to a public repository to get check runs. OAuth apps and authenticated users must have the `repo` scope to get check runs in a private repository.**Note:** The endpoints to manage checks only look for pushes in the repository where the check suite or check run were created. Pushes to a branch in a forked repository are not detected and return an empty `pull_requests` array.If there are more than 1000 check suites on a single git reference, this endpoint will limit check runs to the 1000 most recent check suites. To iterate over all possible check runs, use the [List check suites for a Git reference](https://docs.github.com/rest/reference/checks#list-check-suites-for-a-git-reference) endpoint and provide the `check_suite_id` parameter to the [List check runs in a check suite](https://docs.github.com/rest/reference/checks#list-check-runs-in-a-check-suite) endpoint.
func (m *ItemItemCommitsItemCheckRunsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemItemCommitsItemCheckRunsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *ItemItemCommitsItemCheckRunsRequestBuilder) WithUrl(rawUrl string)(*ItemItemCommitsItemCheckRunsRequestBuilder) {
    return NewItemItemCommitsItemCheckRunsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
