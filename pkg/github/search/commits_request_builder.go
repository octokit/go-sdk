package search

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i625e45c268d4ab4a42a97d02ce2d6719584691128398ca7ec6cca25ecc82eda8 "github.com/octokit/go-sdk/pkg/github/search/commits"
)

// CommitsRequestBuilder builds and executes requests for operations under \search\commits
type CommitsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// CommitsRequestBuilderGetQueryParameters find commits via various criteria on the default branch (usually `main`). This method returns up to 100 results [per page](https://docs.github.com/rest/guides/using-pagination-in-the-rest-api).When searching for commits, you can get text match metadata for the **message** field when you provide the `text-match` media type. For more details about how to receive highlighted search results, see [Text matchmetadata](https://docs.github.com/rest/search/search#text-match-metadata).For example, if you want to find commits related to CSS in the [octocat/Spoon-Knife](https://github.com/octocat/Spoon-Knife) repository. Your query would look something like this:`q=repo:octocat/Spoon-Knife+css`
type CommitsRequestBuilderGetQueryParameters struct {
    // Determines whether the first search result returned is the highest number of matches (`desc`) or lowest number of matches (`asc`). This parameter is ignored unless you provide `sort`.
    // Deprecated: This property is deprecated, use orderAsGetOrderQueryParameterType instead
    Order *string `uriparametername:"order"`
    // Determines whether the first search result returned is the highest number of matches (`desc`) or lowest number of matches (`asc`). This parameter is ignored unless you provide `sort`.
    OrderAsGetOrderQueryParameterType *i625e45c268d4ab4a42a97d02ce2d6719584691128398ca7ec6cca25ecc82eda8.GetOrderQueryParameterType `uriparametername:"order"`
    // Page number of the results to fetch.
    Page *int32 `uriparametername:"page"`
    // The number of results per page (max 100).
    Per_page *int32 `uriparametername:"per_page"`
    // The query contains one or more search keywords and qualifiers. Qualifiers allow you to limit your search to specific areas of GitHub. The REST API supports the same qualifiers as the web interface for GitHub. To learn more about the format of the query, see [Constructing a search query](https://docs.github.com/rest/search/search#constructing-a-search-query). See "[Searching commits](https://docs.github.com/search-github/searching-on-github/searching-commits)" for a detailed list of qualifiers.
    Q *string `uriparametername:"q"`
    // Sorts the results of your query by `author-date` or `committer-date`. Default: [best match](https://docs.github.com/rest/search/search#ranking-search-results)
    // Deprecated: This property is deprecated, use sortAsGetSortQueryParameterType instead
    Sort *string `uriparametername:"sort"`
    // Sorts the results of your query by `author-date` or `committer-date`. Default: [best match](https://docs.github.com/rest/search/search#ranking-search-results)
    SortAsGetSortQueryParameterType *i625e45c268d4ab4a42a97d02ce2d6719584691128398ca7ec6cca25ecc82eda8.GetSortQueryParameterType `uriparametername:"sort"`
}
// CommitsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CommitsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *CommitsRequestBuilderGetQueryParameters
}
// NewCommitsRequestBuilderInternal instantiates a new CommitsRequestBuilder and sets the default values.
func NewCommitsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CommitsRequestBuilder) {
    m := &CommitsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/search/commits{?q*,sort*,order*,per_page*,page*}", pathParameters),
    }
    return m
}
// NewCommitsRequestBuilder instantiates a new CommitsRequestBuilder and sets the default values.
func NewCommitsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CommitsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewCommitsRequestBuilderInternal(urlParams, requestAdapter)
}
// Get find commits via various criteria on the default branch (usually `main`). This method returns up to 100 results [per page](https://docs.github.com/rest/guides/using-pagination-in-the-rest-api).When searching for commits, you can get text match metadata for the **message** field when you provide the `text-match` media type. For more details about how to receive highlighted search results, see [Text matchmetadata](https://docs.github.com/rest/search/search#text-match-metadata).For example, if you want to find commits related to CSS in the [octocat/Spoon-Knife](https://github.com/octocat/Spoon-Knife) repository. Your query would look something like this:`q=repo:octocat/Spoon-Knife+css`
// Deprecated: This method is obsolete. Use GetAsCommitsGetResponse instead.
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/search/search#search-commits
func (m *CommitsRequestBuilder) Get(ctx context.Context, requestConfiguration *CommitsRequestBuilderGetRequestConfiguration)(CommitsResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateCommitsResponseFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(CommitsResponseable), nil
}
// GetAsCommitsGetResponse find commits via various criteria on the default branch (usually `main`). This method returns up to 100 results [per page](https://docs.github.com/rest/guides/using-pagination-in-the-rest-api).When searching for commits, you can get text match metadata for the **message** field when you provide the `text-match` media type. For more details about how to receive highlighted search results, see [Text matchmetadata](https://docs.github.com/rest/search/search#text-match-metadata).For example, if you want to find commits related to CSS in the [octocat/Spoon-Knife](https://github.com/octocat/Spoon-Knife) repository. Your query would look something like this:`q=repo:octocat/Spoon-Knife+css`
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/search/search#search-commits
func (m *CommitsRequestBuilder) GetAsCommitsGetResponse(ctx context.Context, requestConfiguration *CommitsRequestBuilderGetRequestConfiguration)(CommitsGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateCommitsGetResponseFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(CommitsGetResponseable), nil
}
// ToGetRequestInformation find commits via various criteria on the default branch (usually `main`). This method returns up to 100 results [per page](https://docs.github.com/rest/guides/using-pagination-in-the-rest-api).When searching for commits, you can get text match metadata for the **message** field when you provide the `text-match` media type. For more details about how to receive highlighted search results, see [Text matchmetadata](https://docs.github.com/rest/search/search#text-match-metadata).For example, if you want to find commits related to CSS in the [octocat/Spoon-Knife](https://github.com/octocat/Spoon-Knife) repository. Your query would look something like this:`q=repo:octocat/Spoon-Knife+css`
func (m *CommitsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *CommitsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *CommitsRequestBuilder) WithUrl(rawUrl string)(*CommitsRequestBuilder) {
    return NewCommitsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
