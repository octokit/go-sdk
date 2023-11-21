package repos

import (
    "context"
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i53ac87e8cb3cc9276228f74d38694a208cacb99bb8ceb705eeae99fb88d4d274 "strconv"
    i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e "github.com/octokit/go-sdk/github/octokit/models"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    idebe7880df40ec16823f14ee325804cbf75205b53f25102fecc906d91efe63ad "github.com/octokit/go-sdk/github/octokit/repos/item/item/pulls/comments"
)

// ItemItemPullsCommentsRequestBuilder builds and executes requests for operations under \repos\{owner}\{repo}\pulls\comments
type ItemItemPullsCommentsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemItemPullsCommentsRequestBuilderGetQueryParameters lists review comments for all pull requests in a repository. By default, review comments are in ascending order by ID.
type ItemItemPullsCommentsRequestBuilderGetQueryParameters struct {
    // The direction to sort results. Ignored without `sort` parameter.
    // Deprecated: This property is deprecated, use directionAsGetDirectionQueryParameterType instead
    Direction *string `uriparametername:"direction"`
    // The direction to sort results. Ignored without `sort` parameter.
    DirectionAsGetDirectionQueryParameterType *idebe7880df40ec16823f14ee325804cbf75205b53f25102fecc906d91efe63ad.GetDirectionQueryParameterType `uriparametername:"direction"`
    // Page number of the results to fetch.
    Page *int32 `uriparametername:"page"`
    // The number of results per page (max 100).
    Per_page *int32 `uriparametername:"per_page"`
    // Only show results that were last updated after the given time. This is a timestamp in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format: `YYYY-MM-DDTHH:MM:SSZ`.
    Since *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time `uriparametername:"since"`
    // 
    // Deprecated: This property is deprecated, use sortAsGetSortQueryParameterType instead
    Sort *string `uriparametername:"sort"`
    // 
    SortAsGetSortQueryParameterType *idebe7880df40ec16823f14ee325804cbf75205b53f25102fecc906d91efe63ad.GetSortQueryParameterType `uriparametername:"sort"`
}
// ItemItemPullsCommentsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemItemPullsCommentsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemItemPullsCommentsRequestBuilderGetQueryParameters
}
// ByComment_id gets an item from the octokit.repos.item.item.pulls.comments.item collection
// Deprecated: This indexer is deprecated and will be removed in the next major version. Use the one with the typed parameter instead.
func (m *ItemItemPullsCommentsRequestBuilder) ByComment_id(comment_id string)(*ItemItemPullsCommentsWithComment_ItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if comment_id != "" {
        urlTplParams["comment_id"] = comment_id
    }
    return NewItemItemPullsCommentsWithComment_ItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// ByComment_idInteger gets an item from the octokit.repos.item.item.pulls.comments.item collection
func (m *ItemItemPullsCommentsRequestBuilder) ByComment_idInteger(comment_id int32)(*ItemItemPullsCommentsWithComment_ItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    urlTplParams["comment_id"] = i53ac87e8cb3cc9276228f74d38694a208cacb99bb8ceb705eeae99fb88d4d274.FormatInt(int64(comment_id), 10)
    return NewItemItemPullsCommentsWithComment_ItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewItemItemPullsCommentsRequestBuilderInternal instantiates a new CommentsRequestBuilder and sets the default values.
func NewItemItemPullsCommentsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemPullsCommentsRequestBuilder) {
    m := &ItemItemPullsCommentsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/repos/{owner}/{repo}/pulls/comments{?sort*,direction*,since*,per_page*,page*}", pathParameters),
    }
    return m
}
// NewItemItemPullsCommentsRequestBuilder instantiates a new CommentsRequestBuilder and sets the default values.
func NewItemItemPullsCommentsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemPullsCommentsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemItemPullsCommentsRequestBuilderInternal(urlParams, requestAdapter)
}
// Get lists review comments for all pull requests in a repository. By default, review comments are in ascending order by ID.
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/pulls/comments#list-review-comments-in-a-repository
func (m *ItemItemPullsCommentsRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemItemPullsCommentsRequestBuilderGetRequestConfiguration)([]i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.PullRequestReviewCommentable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.SendCollection(ctx, requestInfo, i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.CreatePullRequestReviewCommentFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    val := make([]i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.PullRequestReviewCommentable, len(res))
    for i, v := range res {
        if v != nil {
            val[i] = v.(i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.PullRequestReviewCommentable)
        }
    }
    return val, nil
}
// ToGetRequestInformation lists review comments for all pull requests in a repository. By default, review comments are in ascending order by ID.
func (m *ItemItemPullsCommentsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemItemPullsCommentsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *ItemItemPullsCommentsRequestBuilder) WithUrl(rawUrl string)(*ItemItemPullsCommentsRequestBuilder) {
    return NewItemItemPullsCommentsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
