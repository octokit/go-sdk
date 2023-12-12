package search

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6 "github.com/octokit/go-sdk/pkg/github/models"
    i7add74210511a3bba07275779ce4518d4e5b8c0b4c74994de39cfb93f5b96ecf "github.com/octokit/go-sdk/pkg/github/search/labels"
)

// LabelsRequestBuilder builds and executes requests for operations under \search\labels
type LabelsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// LabelsRequestBuilderGetQueryParameters find labels in a repository with names or descriptions that match search keywords. Returns up to 100 results [per page](https://docs.github.com/rest/guides/using-pagination-in-the-rest-api).When searching for labels, you can get text match metadata for the label **name** and **description** fields when you pass the `text-match` media type. For more details about how to receive highlighted search results, see [Text match metadata](https://docs.github.com/rest/search/search#text-match-metadata).For example, if you want to find labels in the `linguist` repository that match `bug`, `defect`, or `enhancement`. Your query might look like this:`q=bug+defect+enhancement&repository_id=64778136`The labels that best match the query appear first in the search results.
type LabelsRequestBuilderGetQueryParameters struct {
    // Determines whether the first search result returned is the highest number of matches (`desc`) or lowest number of matches (`asc`). This parameter is ignored unless you provide `sort`.
    // Deprecated: This property is deprecated, use orderAsGetOrderQueryParameterType instead
    Order *string `uriparametername:"order"`
    // Determines whether the first search result returned is the highest number of matches (`desc`) or lowest number of matches (`asc`). This parameter is ignored unless you provide `sort`.
    OrderAsGetOrderQueryParameterType *i7add74210511a3bba07275779ce4518d4e5b8c0b4c74994de39cfb93f5b96ecf.GetOrderQueryParameterType `uriparametername:"order"`
    // Page number of the results to fetch.
    Page *int32 `uriparametername:"page"`
    // The number of results per page (max 100).
    Per_page *int32 `uriparametername:"per_page"`
    // The search keywords. This endpoint does not accept qualifiers in the query. To learn more about the format of the query, see [Constructing a search query](https://docs.github.com/rest/search/search#constructing-a-search-query).
    Q *string `uriparametername:"q"`
    // The id of the repository.
    Repository_id *int32 `uriparametername:"repository_id"`
    // Sorts the results of your query by when the label was `created` or `updated`. Default: [best match](https://docs.github.com/rest/search/search#ranking-search-results)
    // Deprecated: This property is deprecated, use sortAsGetSortQueryParameterType instead
    Sort *string `uriparametername:"sort"`
    // Sorts the results of your query by when the label was `created` or `updated`. Default: [best match](https://docs.github.com/rest/search/search#ranking-search-results)
    SortAsGetSortQueryParameterType *i7add74210511a3bba07275779ce4518d4e5b8c0b4c74994de39cfb93f5b96ecf.GetSortQueryParameterType `uriparametername:"sort"`
}
// LabelsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type LabelsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *LabelsRequestBuilderGetQueryParameters
}
// NewLabelsRequestBuilderInternal instantiates a new LabelsRequestBuilder and sets the default values.
func NewLabelsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*LabelsRequestBuilder) {
    m := &LabelsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/search/labels{?repository_id*,q*,sort*,order*,per_page*,page*}", pathParameters),
    }
    return m
}
// NewLabelsRequestBuilder instantiates a new LabelsRequestBuilder and sets the default values.
func NewLabelsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*LabelsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewLabelsRequestBuilderInternal(urlParams, requestAdapter)
}
// Get find labels in a repository with names or descriptions that match search keywords. Returns up to 100 results [per page](https://docs.github.com/rest/guides/using-pagination-in-the-rest-api).When searching for labels, you can get text match metadata for the label **name** and **description** fields when you pass the `text-match` media type. For more details about how to receive highlighted search results, see [Text match metadata](https://docs.github.com/rest/search/search#text-match-metadata).For example, if you want to find labels in the `linguist` repository that match `bug`, `defect`, or `enhancement`. Your query might look like this:`q=bug+defect+enhancement&repository_id=64778136`The labels that best match the query appear first in the search results.
// Deprecated: This method is obsolete. Use GetAsLabelsGetResponse instead.
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/search/search#search-labels
func (m *LabelsRequestBuilder) Get(ctx context.Context, requestConfiguration *LabelsRequestBuilderGetRequestConfiguration)(LabelsResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "403": i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CreateBasicErrorFromDiscriminatorValue,
        "404": i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CreateBasicErrorFromDiscriminatorValue,
        "422": i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CreateValidationErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateLabelsResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(LabelsResponseable), nil
}
// GetAsLabelsGetResponse find labels in a repository with names or descriptions that match search keywords. Returns up to 100 results [per page](https://docs.github.com/rest/guides/using-pagination-in-the-rest-api).When searching for labels, you can get text match metadata for the label **name** and **description** fields when you pass the `text-match` media type. For more details about how to receive highlighted search results, see [Text match metadata](https://docs.github.com/rest/search/search#text-match-metadata).For example, if you want to find labels in the `linguist` repository that match `bug`, `defect`, or `enhancement`. Your query might look like this:`q=bug+defect+enhancement&repository_id=64778136`The labels that best match the query appear first in the search results.
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/search/search#search-labels
func (m *LabelsRequestBuilder) GetAsLabelsGetResponse(ctx context.Context, requestConfiguration *LabelsRequestBuilderGetRequestConfiguration)(LabelsGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "403": i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CreateBasicErrorFromDiscriminatorValue,
        "404": i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CreateBasicErrorFromDiscriminatorValue,
        "422": i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CreateValidationErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateLabelsGetResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(LabelsGetResponseable), nil
}
// ToGetRequestInformation find labels in a repository with names or descriptions that match search keywords. Returns up to 100 results [per page](https://docs.github.com/rest/guides/using-pagination-in-the-rest-api).When searching for labels, you can get text match metadata for the label **name** and **description** fields when you pass the `text-match` media type. For more details about how to receive highlighted search results, see [Text match metadata](https://docs.github.com/rest/search/search#text-match-metadata).For example, if you want to find labels in the `linguist` repository that match `bug`, `defect`, or `enhancement`. Your query might look like this:`q=bug+defect+enhancement&repository_id=64778136`The labels that best match the query appear first in the search results.
func (m *LabelsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *LabelsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *LabelsRequestBuilder) WithUrl(rawUrl string)(*LabelsRequestBuilder) {
    return NewLabelsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
