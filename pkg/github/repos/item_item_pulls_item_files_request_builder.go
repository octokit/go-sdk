package repos

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6 "github.com/octokit/go-sdk/pkg/github/models"
)

// ItemItemPullsItemFilesRequestBuilder builds and executes requests for operations under \repos\{repos-id}\{Owner-id}\pulls\{pull_number}\files
type ItemItemPullsItemFilesRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemItemPullsItemFilesRequestBuilderGetQueryParameters **Note:** Responses include a maximum of 3000 files. The paginated response returns 30 files per page by default.
type ItemItemPullsItemFilesRequestBuilderGetQueryParameters struct {
    // The page number of the results to fetch. For more information, see "[Using pagination in the REST API](https://docs.github.com/rest/using-the-rest-api/using-pagination-in-the-rest-api)."
    Page *int32 `uriparametername:"page"`
    // The number of results per page (max 100). For more information, see "[Using pagination in the REST API](https://docs.github.com/rest/using-the-rest-api/using-pagination-in-the-rest-api)."
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
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/repos/{repos%2Did}/{Owner%2Did}/pulls/{pull_number}/files{?per_page*,page*}", pathParameters),
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
func (m *ItemItemPullsItemFilesRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemItemPullsItemFilesRequestBuilderGetRequestConfiguration)([]i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.DiffEntryable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "422": i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CreateValidationErrorFromDiscriminatorValue,
        "500": i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CreateBasicErrorFromDiscriminatorValue,
        "503": i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CreateFiles503ErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.SendCollection(ctx, requestInfo, i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CreateDiffEntryFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    val := make([]i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.DiffEntryable, len(res))
    for i, v := range res {
        if v != nil {
            val[i] = v.(i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.DiffEntryable)
        }
    }
    return val, nil
}
// ToGetRequestInformation **Note:** Responses include a maximum of 3000 files. The paginated response returns 30 files per page by default.
func (m *ItemItemPullsItemFilesRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemItemPullsItemFilesRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *ItemItemPullsItemFilesRequestBuilder) WithUrl(rawUrl string)(*ItemItemPullsItemFilesRequestBuilder) {
    return NewItemItemPullsItemFilesRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
