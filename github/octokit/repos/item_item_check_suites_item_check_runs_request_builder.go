package repos

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ic412e7eef8f28635532af2e2d57592e8f211f6a68a1203cad7cc3dd23d9744db "github.com/octokit/go-sdk/github/octokit/repos/item/item/checksuites/item/checkruns"
)

// ItemItemCheckSuitesItemCheckRunsRequestBuilder builds and executes requests for operations under \repos\{owner}\{repo}\check-suites\{check_suite_id}\check-runs
type ItemItemCheckSuitesItemCheckRunsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemItemCheckSuitesItemCheckRunsRequestBuilderGetQueryParameters lists check runs for a check suite using its `id`. GitHub Apps must have the `checks:read` permission on a private repository or pull access to a public repository to get check runs. OAuth apps and authenticated users must have the `repo` scope to get check runs in a private repository.**Note:** The endpoints to manage checks only look for pushes in the repository where the check suite or check run were created. Pushes to a branch in a forked repository are not detected and return an empty `pull_requests` array.
type ItemItemCheckSuitesItemCheckRunsRequestBuilderGetQueryParameters struct {
    // Returns check runs with the specified `name`.
    Check_name *string `uriparametername:"check_name"`
    // Filters check runs by their `completed_at` timestamp. `latest` returns the most recent check runs.
    // Deprecated: This property is deprecated, use filterAsGetFilterQueryParameterType instead
    Filter *string `uriparametername:"filter"`
    // Filters check runs by their `completed_at` timestamp. `latest` returns the most recent check runs.
    FilterAsGetFilterQueryParameterType *ic412e7eef8f28635532af2e2d57592e8f211f6a68a1203cad7cc3dd23d9744db.GetFilterQueryParameterType `uriparametername:"filter"`
    // Page number of the results to fetch.
    Page *int32 `uriparametername:"page"`
    // The number of results per page (max 100).
    Per_page *int32 `uriparametername:"per_page"`
    // Returns check runs with the specified `status`.
    // Deprecated: This property is deprecated, use statusAsGetStatusQueryParameterType instead
    Status *string `uriparametername:"status"`
    // Returns check runs with the specified `status`.
    StatusAsGetStatusQueryParameterType *ic412e7eef8f28635532af2e2d57592e8f211f6a68a1203cad7cc3dd23d9744db.GetStatusQueryParameterType `uriparametername:"status"`
}
// ItemItemCheckSuitesItemCheckRunsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemItemCheckSuitesItemCheckRunsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemItemCheckSuitesItemCheckRunsRequestBuilderGetQueryParameters
}
// NewItemItemCheckSuitesItemCheckRunsRequestBuilderInternal instantiates a new CheckRunsRequestBuilder and sets the default values.
func NewItemItemCheckSuitesItemCheckRunsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemCheckSuitesItemCheckRunsRequestBuilder) {
    m := &ItemItemCheckSuitesItemCheckRunsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/repos/{owner}/{repo}/check-suites/{check_suite_id}/check-runs{?check_name*,status*,filter*,per_page*,page*}", pathParameters),
    }
    return m
}
// NewItemItemCheckSuitesItemCheckRunsRequestBuilder instantiates a new CheckRunsRequestBuilder and sets the default values.
func NewItemItemCheckSuitesItemCheckRunsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemCheckSuitesItemCheckRunsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemItemCheckSuitesItemCheckRunsRequestBuilderInternal(urlParams, requestAdapter)
}
// Get lists check runs for a check suite using its `id`. GitHub Apps must have the `checks:read` permission on a private repository or pull access to a public repository to get check runs. OAuth apps and authenticated users must have the `repo` scope to get check runs in a private repository.**Note:** The endpoints to manage checks only look for pushes in the repository where the check suite or check run were created. Pushes to a branch in a forked repository are not detected and return an empty `pull_requests` array.
// Deprecated: This method is obsolete. Use GetAsCheckRunsGetResponse instead.
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/checks/runs#list-check-runs-in-a-check-suite
func (m *ItemItemCheckSuitesItemCheckRunsRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemItemCheckSuitesItemCheckRunsRequestBuilderGetRequestConfiguration)(ItemItemCheckSuitesItemCheckRunsResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateItemItemCheckSuitesItemCheckRunsResponseFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ItemItemCheckSuitesItemCheckRunsResponseable), nil
}
// GetAsCheckRunsGetResponse lists check runs for a check suite using its `id`. GitHub Apps must have the `checks:read` permission on a private repository or pull access to a public repository to get check runs. OAuth apps and authenticated users must have the `repo` scope to get check runs in a private repository.**Note:** The endpoints to manage checks only look for pushes in the repository where the check suite or check run were created. Pushes to a branch in a forked repository are not detected and return an empty `pull_requests` array.
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/checks/runs#list-check-runs-in-a-check-suite
func (m *ItemItemCheckSuitesItemCheckRunsRequestBuilder) GetAsCheckRunsGetResponse(ctx context.Context, requestConfiguration *ItemItemCheckSuitesItemCheckRunsRequestBuilderGetRequestConfiguration)(ItemItemCheckSuitesItemCheckRunsGetResponseable, error) {
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
// ToGetRequestInformation lists check runs for a check suite using its `id`. GitHub Apps must have the `checks:read` permission on a private repository or pull access to a public repository to get check runs. OAuth apps and authenticated users must have the `repo` scope to get check runs in a private repository.**Note:** The endpoints to manage checks only look for pushes in the repository where the check suite or check run were created. Pushes to a branch in a forked repository are not detected and return an empty `pull_requests` array.
func (m *ItemItemCheckSuitesItemCheckRunsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemItemCheckSuitesItemCheckRunsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers.TryAdd("Accept", "application/json;q=1")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
func (m *ItemItemCheckSuitesItemCheckRunsRequestBuilder) WithUrl(rawUrl string)(*ItemItemCheckSuitesItemCheckRunsRequestBuilder) {
    return NewItemItemCheckSuitesItemCheckRunsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
