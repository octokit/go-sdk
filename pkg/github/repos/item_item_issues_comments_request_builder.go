package repos

import (
    "context"
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i53ac87e8cb3cc9276228f74d38694a208cacb99bb8ceb705eeae99fb88d4d274 "strconv"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6 "github.com/octokit/go-sdk/pkg/github/models"
    if9fe7687d66ccbe48e81b41eb88b87988bd006a467b7e33a09cc95d9ac51aa3f "github.com/octokit/go-sdk/pkg/github/repos/item/item/issues/comments"
)

// ItemItemIssuesCommentsRequestBuilder builds and executes requests for operations under \repos\{repos-id}\{Owner-id}\issues\comments
type ItemItemIssuesCommentsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemItemIssuesCommentsRequestBuilderGetQueryParameters you can use the REST API to list comments on issues and pull requests for a repository. Every pull request is an issue, but not every issue is a pull request.By default, issue comments are ordered by ascending ID.
type ItemItemIssuesCommentsRequestBuilderGetQueryParameters struct {
    // Either `asc` or `desc`. Ignored without the `sort` parameter.
    // Deprecated: This property is deprecated, use directionAsGetDirectionQueryParameterType instead
    Direction *string `uriparametername:"direction"`
    // Either `asc` or `desc`. Ignored without the `sort` parameter.
    DirectionAsGetDirectionQueryParameterType *if9fe7687d66ccbe48e81b41eb88b87988bd006a467b7e33a09cc95d9ac51aa3f.GetDirectionQueryParameterType `uriparametername:"direction"`
    // Page number of the results to fetch.
    Page *int32 `uriparametername:"page"`
    // The number of results per page (max 100).
    Per_page *int32 `uriparametername:"per_page"`
    // Only show results that were last updated after the given time. This is a timestamp in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format: `YYYY-MM-DDTHH:MM:SSZ`.
    Since *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time `uriparametername:"since"`
    // The property to sort the results by.
    // Deprecated: This property is deprecated, use sortAsGetSortQueryParameterType instead
    Sort *string `uriparametername:"sort"`
    // The property to sort the results by.
    SortAsGetSortQueryParameterType *if9fe7687d66ccbe48e81b41eb88b87988bd006a467b7e33a09cc95d9ac51aa3f.GetSortQueryParameterType `uriparametername:"sort"`
}
// ItemItemIssuesCommentsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemItemIssuesCommentsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemItemIssuesCommentsRequestBuilderGetQueryParameters
}
// ByComment_id gets an item from the github.com/octokit/go-sdk/pkg/github/.repos.item.item.issues.comments.item collection
// Deprecated: This indexer is deprecated and will be removed in the next major version. Use the one with the typed parameter instead.
func (m *ItemItemIssuesCommentsRequestBuilder) ByComment_id(comment_id string)(*ItemItemIssuesCommentsWithComment_ItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if comment_id != "" {
        urlTplParams["comment_id"] = comment_id
    }
    return NewItemItemIssuesCommentsWithComment_ItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// ByComment_idInteger gets an item from the github.com/octokit/go-sdk/pkg/github/.repos.item.item.issues.comments.item collection
func (m *ItemItemIssuesCommentsRequestBuilder) ByComment_idInteger(comment_id int32)(*ItemItemIssuesCommentsWithComment_ItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    urlTplParams["comment_id"] = i53ac87e8cb3cc9276228f74d38694a208cacb99bb8ceb705eeae99fb88d4d274.FormatInt(int64(comment_id), 10)
    return NewItemItemIssuesCommentsWithComment_ItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewItemItemIssuesCommentsRequestBuilderInternal instantiates a new CommentsRequestBuilder and sets the default values.
func NewItemItemIssuesCommentsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemIssuesCommentsRequestBuilder) {
    m := &ItemItemIssuesCommentsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/repos/{repos%2Did}/{Owner%2Did}/issues/comments{?sort*,direction*,since*,per_page*,page*}", pathParameters),
    }
    return m
}
// NewItemItemIssuesCommentsRequestBuilder instantiates a new CommentsRequestBuilder and sets the default values.
func NewItemItemIssuesCommentsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemIssuesCommentsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemItemIssuesCommentsRequestBuilderInternal(urlParams, requestAdapter)
}
// Get you can use the REST API to list comments on issues and pull requests for a repository. Every pull request is an issue, but not every issue is a pull request.By default, issue comments are ordered by ascending ID.
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/issues/comments#list-issue-comments-for-a-repository
func (m *ItemItemIssuesCommentsRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemItemIssuesCommentsRequestBuilderGetRequestConfiguration)([]i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.IssueCommentable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "404": i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CreateBasicErrorFromDiscriminatorValue,
        "422": i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CreateValidationErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.SendCollection(ctx, requestInfo, i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CreateIssueCommentFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    val := make([]i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.IssueCommentable, len(res))
    for i, v := range res {
        if v != nil {
            val[i] = v.(i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.IssueCommentable)
        }
    }
    return val, nil
}
// ToGetRequestInformation you can use the REST API to list comments on issues and pull requests for a repository. Every pull request is an issue, but not every issue is a pull request.By default, issue comments are ordered by ascending ID.
func (m *ItemItemIssuesCommentsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemItemIssuesCommentsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *ItemItemIssuesCommentsRequestBuilder) WithUrl(rawUrl string)(*ItemItemIssuesCommentsRequestBuilder) {
    return NewItemItemIssuesCommentsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
