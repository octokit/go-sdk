package repos

import (
    "context"
    i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e "github.com/octokit/go-sdk/github/octokit/models"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemItemPullsItemFilesRequestBuilder builds and executes requests for operations under \repos\{owner}\{repo}\pulls\{pull_number}\files
type ItemItemPullsItemFilesRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemItemPullsItemFilesRequestBuilderGetQueryParameters **Note:** Responses include a maximum of 3000 files. The paginated response returns 30 files per page by default.
type ItemItemPullsItemFilesRequestBuilderGetQueryParameters struct {
    // Page number of the results to fetch.
    Page *int32 `uriparametername:"page"`
    // The number of results per page (max 100).
    Per_page *int32 `uriparametername:"per_page"`
}
// ItemItemPullsItemFilesRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemItemPullsItemFilesRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemItemPullsItemFilesRequestBuilderGetQueryParameters
}
// NewItemItemPullsItemFilesRequestBuilderInternal instantiates a new FilesRequestBuilder and sets the default values.
func NewItemItemPullsItemFilesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemPullsItemFilesRequestBuilder) {
    m := &ItemItemPullsItemFilesRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/repos/{owner}/{repo}/pulls/{pull_number}/files{?per_page*,page*}", pathParameters),
    }
    return m
}
// NewItemItemPullsItemFilesRequestBuilder instantiates a new FilesRequestBuilder and sets the default values.
func NewItemItemPullsItemFilesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemPullsItemFilesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemItemPullsItemFilesRequestBuilderInternal(urlParams, requestAdapter)
}
// Get **Note:** Responses include a maximum of 3000 files. The paginated response returns 30 files per page by default.
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/pulls/pulls#list-pull-requests-files
func (m *ItemItemPullsItemFilesRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemItemPullsItemFilesRequestBuilderGetRequestConfiguration)([]i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.DiffEntryable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "422": i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.CreateValidationErrorFromDiscriminatorValue,
        "500": i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.CreateBasicErrorFromDiscriminatorValue,
        "503": i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.CreateFiles503ErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.SendCollection(ctx, requestInfo, i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.CreateDiffEntryFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    val := make([]i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.DiffEntryable, len(res))
    for i, v := range res {
        if v != nil {
            val[i] = v.(i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.DiffEntryable)
        }
    }
    return val, nil
}
// ToGetRequestInformation **Note:** Responses include a maximum of 3000 files. The paginated response returns 30 files per page by default.
func (m *ItemItemPullsItemFilesRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemItemPullsItemFilesRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
func (m *ItemItemPullsItemFilesRequestBuilder) WithUrl(rawUrl string)(*ItemItemPullsItemFilesRequestBuilder) {
    return NewItemItemPullsItemFilesRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
