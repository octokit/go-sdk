package repos

import (
    "context"
    i53ac87e8cb3cc9276228f74d38694a208cacb99bb8ceb705eeae99fb88d4d274 "strconv"
    i33bcfa37677dfd8b71eaf22ad3bfd140ba1702e7a8418558499c22418bf2fef2 "github.com/octokit/go-sdk/models"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemItemCommentsRequestBuilder builds and executes requests for operations under \repos\{owner}\{repo}\comments
type ItemItemCommentsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemItemCommentsRequestBuilderGetQueryParameters commit Comments use [these custom media types](https://docs.github.com/rest/overview/media-types). You can read more about the use of media types in the API [here](https://docs.github.com/rest/overview/media-types/).Comments are ordered by ascending ID.
type ItemItemCommentsRequestBuilderGetQueryParameters struct {
    // Page number of the results to fetch.
    Page *int32 `uriparametername:"page"`
    // The number of results per page (max 100).
    Per_page *int32 `uriparametername:"per_page"`
}
// ItemItemCommentsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemItemCommentsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemItemCommentsRequestBuilderGetQueryParameters
}
// ByComment_id gets an item from the go-sdk.repos.item.item.comments.item collection
// Deprecated: This indexer is deprecated and will be removed in the next major version. Use the one with the typed parameter instead.
func (m *ItemItemCommentsRequestBuilder) ByComment_id(comment_id string)(*ItemItemCommentsWithComment_ItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if comment_id != "" {
        urlTplParams["comment_id"] = comment_id
    }
    return NewItemItemCommentsWithComment_ItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// ByComment_idInteger gets an item from the go-sdk.repos.item.item.comments.item collection
func (m *ItemItemCommentsRequestBuilder) ByComment_idInteger(comment_id int32)(*ItemItemCommentsWithComment_ItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    urlTplParams["comment_id"] = i53ac87e8cb3cc9276228f74d38694a208cacb99bb8ceb705eeae99fb88d4d274.FormatInt(int64(comment_id), 10)
    return NewItemItemCommentsWithComment_ItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewItemItemCommentsRequestBuilderInternal instantiates a new CommentsRequestBuilder and sets the default values.
func NewItemItemCommentsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemCommentsRequestBuilder) {
    m := &ItemItemCommentsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/repos/{owner}/{repo}/comments{?per_page*,page*}", pathParameters),
    }
    return m
}
// NewItemItemCommentsRequestBuilder instantiates a new CommentsRequestBuilder and sets the default values.
func NewItemItemCommentsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemCommentsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemItemCommentsRequestBuilderInternal(urlParams, requestAdapter)
}
// Get commit Comments use [these custom media types](https://docs.github.com/rest/overview/media-types). You can read more about the use of media types in the API [here](https://docs.github.com/rest/overview/media-types/).Comments are ordered by ascending ID.
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/commits/comments#list-commit-comments-for-a-repository
func (m *ItemItemCommentsRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemItemCommentsRequestBuilderGetRequestConfiguration)([]i33bcfa37677dfd8b71eaf22ad3bfd140ba1702e7a8418558499c22418bf2fef2.CommitCommentable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.SendCollection(ctx, requestInfo, i33bcfa37677dfd8b71eaf22ad3bfd140ba1702e7a8418558499c22418bf2fef2.CreateCommitCommentFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    val := make([]i33bcfa37677dfd8b71eaf22ad3bfd140ba1702e7a8418558499c22418bf2fef2.CommitCommentable, len(res))
    for i, v := range res {
        if v != nil {
            val[i] = v.(i33bcfa37677dfd8b71eaf22ad3bfd140ba1702e7a8418558499c22418bf2fef2.CommitCommentable)
        }
    }
    return val, nil
}
// ToGetRequestInformation commit Comments use [these custom media types](https://docs.github.com/rest/overview/media-types). You can read more about the use of media types in the API [here](https://docs.github.com/rest/overview/media-types/).Comments are ordered by ascending ID.
func (m *ItemItemCommentsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemItemCommentsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *ItemItemCommentsRequestBuilder) WithUrl(rawUrl string)(*ItemItemCommentsRequestBuilder) {
    return NewItemItemCommentsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
