package repos

import (
    "context"
    i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e "github.com/octokit/go-sdk/github/octokit/models"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemItemCheckSuitesWithCheck_suite_ItemRequestBuilder builds and executes requests for operations under \repos\{owner}\{repo}\check-suites\{check_suite_id}
type ItemItemCheckSuitesWithCheck_suite_ItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemItemCheckSuitesWithCheck_suite_ItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemItemCheckSuitesWithCheck_suite_ItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// CheckRuns the checkRuns property
func (m *ItemItemCheckSuitesWithCheck_suite_ItemRequestBuilder) CheckRuns()(*ItemItemCheckSuitesItemCheckRunsRequestBuilder) {
    return NewItemItemCheckSuitesItemCheckRunsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewItemItemCheckSuitesWithCheck_suite_ItemRequestBuilderInternal instantiates a new WithCheck_suite_ItemRequestBuilder and sets the default values.
func NewItemItemCheckSuitesWithCheck_suite_ItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemCheckSuitesWithCheck_suite_ItemRequestBuilder) {
    m := &ItemItemCheckSuitesWithCheck_suite_ItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/repos/{owner}/{repo}/check-suites/{check_suite_id}", pathParameters),
    }
    return m
}
// NewItemItemCheckSuitesWithCheck_suite_ItemRequestBuilder instantiates a new WithCheck_suite_ItemRequestBuilder and sets the default values.
func NewItemItemCheckSuitesWithCheck_suite_ItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemCheckSuitesWithCheck_suite_ItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemItemCheckSuitesWithCheck_suite_ItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Get **Note:** The Checks API only looks for pushes in the repository where the check suite or check run were created. Pushes to a branch in a forked repository are not detected and return an empty `pull_requests` array and a `null` value for `head_branch`.Gets a single check suite using its `id`. GitHub Apps must have the `checks:read` permission on a private repository or pull access to a public repository to get check suites. OAuth apps and authenticated users must have the `repo` scope to get check suites in a private repository.
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/checks/suites#get-a-check-suite
func (m *ItemItemCheckSuitesWithCheck_suite_ItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemItemCheckSuitesWithCheck_suite_ItemRequestBuilderGetRequestConfiguration)(i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.CheckSuiteable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.CreateCheckSuiteFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.CheckSuiteable), nil
}
// Rerequest the rerequest property
func (m *ItemItemCheckSuitesWithCheck_suite_ItemRequestBuilder) Rerequest()(*ItemItemCheckSuitesItemRerequestRequestBuilder) {
    return NewItemItemCheckSuitesItemRerequestRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToGetRequestInformation **Note:** The Checks API only looks for pushes in the repository where the check suite or check run were created. Pushes to a branch in a forked repository are not detected and return an empty `pull_requests` array and a `null` value for `head_branch`.Gets a single check suite using its `id`. GitHub Apps must have the `checks:read` permission on a private repository or pull access to a public repository to get check suites. OAuth apps and authenticated users must have the `repo` scope to get check suites in a private repository.
func (m *ItemItemCheckSuitesWithCheck_suite_ItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemItemCheckSuitesWithCheck_suite_ItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
func (m *ItemItemCheckSuitesWithCheck_suite_ItemRequestBuilder) WithUrl(rawUrl string)(*ItemItemCheckSuitesWithCheck_suite_ItemRequestBuilder) {
    return NewItemItemCheckSuitesWithCheck_suite_ItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
