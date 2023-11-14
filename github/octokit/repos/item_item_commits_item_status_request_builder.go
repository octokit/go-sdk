package repos

import (
    "context"
    i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e "github.com/octokit/go-sdk/github/octokit/models"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemItemCommitsItemStatusRequestBuilder builds and executes requests for operations under \repos\{owner}\{repo}\commits\{commit_sha}\status
type ItemItemCommitsItemStatusRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemItemCommitsItemStatusRequestBuilderGetQueryParameters users with pull access in a repository can access a combined view of commit statuses for a given ref. The ref can be a SHA, a branch name, or a tag name.Additionally, a combined `state` is returned. The `state` is one of:*   **failure** if any of the contexts report as `error` or `failure`*   **pending** if there are no statuses or a context is `pending`*   **success** if the latest status for all contexts is `success`
type ItemItemCommitsItemStatusRequestBuilderGetQueryParameters struct {
    // Page number of the results to fetch.
    Page *int32 `uriparametername:"page"`
    // The number of results per page (max 100).
    Per_page *int32 `uriparametername:"per_page"`
}
// ItemItemCommitsItemStatusRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemItemCommitsItemStatusRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemItemCommitsItemStatusRequestBuilderGetQueryParameters
}
// NewItemItemCommitsItemStatusRequestBuilderInternal instantiates a new StatusRequestBuilder and sets the default values.
func NewItemItemCommitsItemStatusRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemCommitsItemStatusRequestBuilder) {
    m := &ItemItemCommitsItemStatusRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/repos/{owner}/{repo}/commits/{commit_sha}/status{?per_page*,page*}", pathParameters),
    }
    return m
}
// NewItemItemCommitsItemStatusRequestBuilder instantiates a new StatusRequestBuilder and sets the default values.
func NewItemItemCommitsItemStatusRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemCommitsItemStatusRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemItemCommitsItemStatusRequestBuilderInternal(urlParams, requestAdapter)
}
// Get users with pull access in a repository can access a combined view of commit statuses for a given ref. The ref can be a SHA, a branch name, or a tag name.Additionally, a combined `state` is returned. The `state` is one of:*   **failure** if any of the contexts report as `error` or `failure`*   **pending** if there are no statuses or a context is `pending`*   **success** if the latest status for all contexts is `success`
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/commits/statuses#get-the-combined-status-for-a-specific-reference
func (m *ItemItemCommitsItemStatusRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemItemCommitsItemStatusRequestBuilderGetRequestConfiguration)(i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.CombinedCommitStatusable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "404": i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.CreateBasicErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.CreateCombinedCommitStatusFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.CombinedCommitStatusable), nil
}
// ToGetRequestInformation users with pull access in a repository can access a combined view of commit statuses for a given ref. The ref can be a SHA, a branch name, or a tag name.Additionally, a combined `state` is returned. The `state` is one of:*   **failure** if any of the contexts report as `error` or `failure`*   **pending** if there are no statuses or a context is `pending`*   **success** if the latest status for all contexts is `success`
func (m *ItemItemCommitsItemStatusRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemItemCommitsItemStatusRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *ItemItemCommitsItemStatusRequestBuilder) WithUrl(rawUrl string)(*ItemItemCommitsItemStatusRequestBuilder) {
    return NewItemItemCommitsItemStatusRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
