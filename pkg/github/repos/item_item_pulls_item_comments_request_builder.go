package repos

import (
    "context"
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i53ac87e8cb3cc9276228f74d38694a208cacb99bb8ceb705eeae99fb88d4d274 "strconv"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6 "github.com/octokit/go-sdk/pkg/github/models"
    i7e815420148810f5a54942bff4de28f1f9dc6ef1e9fcbdfdb2255d0329e2272e "github.com/octokit/go-sdk/pkg/github/repos/item/item/pulls/item/comments"
)

// ItemItemPullsItemCommentsRequestBuilder builds and executes requests for operations under \repos\{repos-id}\{Owner-id}\pulls\{pull_number}\comments
type ItemItemPullsItemCommentsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemItemPullsItemCommentsRequestBuilderGetQueryParameters lists all review comments for a pull request. By default, review comments are in ascending order by ID.
type ItemItemPullsItemCommentsRequestBuilderGetQueryParameters struct {
    // The direction to sort results. Ignored without `sort` parameter.
    // Deprecated: This property is deprecated, use directionAsGetDirectionQueryParameterType instead
    Direction *string `uriparametername:"direction"`
    // The direction to sort results. Ignored without `sort` parameter.
    DirectionAsGetDirectionQueryParameterType *i7e815420148810f5a54942bff4de28f1f9dc6ef1e9fcbdfdb2255d0329e2272e.GetDirectionQueryParameterType `uriparametername:"direction"`
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
    SortAsGetSortQueryParameterType *i7e815420148810f5a54942bff4de28f1f9dc6ef1e9fcbdfdb2255d0329e2272e.GetSortQueryParameterType `uriparametername:"sort"`
}
// ItemItemPullsItemCommentsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemItemPullsItemCommentsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemItemPullsItemCommentsRequestBuilderGetQueryParameters
}
// ItemItemPullsItemCommentsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemItemPullsItemCommentsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ByComment_id gets an item from the github.com/octokit/go-sdk/pkg/github/.repos.item.item.pulls.item.comments.item collection
// Deprecated: This indexer is deprecated and will be removed in the next major version. Use the one with the typed parameter instead.
func (m *ItemItemPullsItemCommentsRequestBuilder) ByComment_id(comment_id string)(*ItemItemPullsItemCommentsWithComment_ItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if comment_id != "" {
        urlTplParams["comment_id"] = comment_id
    }
    return NewItemItemPullsItemCommentsWithComment_ItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// ByComment_idInteger gets an item from the github.com/octokit/go-sdk/pkg/github/.repos.item.item.pulls.item.comments.item collection
func (m *ItemItemPullsItemCommentsRequestBuilder) ByComment_idInteger(comment_id int32)(*ItemItemPullsItemCommentsWithComment_ItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    urlTplParams["comment_id"] = i53ac87e8cb3cc9276228f74d38694a208cacb99bb8ceb705eeae99fb88d4d274.FormatInt(int64(comment_id), 10)
    return NewItemItemPullsItemCommentsWithComment_ItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewItemItemPullsItemCommentsRequestBuilderInternal instantiates a new CommentsRequestBuilder and sets the default values.
func NewItemItemPullsItemCommentsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemPullsItemCommentsRequestBuilder) {
    m := &ItemItemPullsItemCommentsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/repos/{repos%2Did}/{Owner%2Did}/pulls/{pull_number}/comments{?sort*,direction*,since*,per_page*,page*}", pathParameters),
    }
    return m
}
// NewItemItemPullsItemCommentsRequestBuilder instantiates a new CommentsRequestBuilder and sets the default values.
func NewItemItemPullsItemCommentsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemPullsItemCommentsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemItemPullsItemCommentsRequestBuilderInternal(urlParams, requestAdapter)
}
// Get lists all review comments for a pull request. By default, review comments are in ascending order by ID.
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/pulls/comments#list-review-comments-on-a-pull-request
func (m *ItemItemPullsItemCommentsRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemItemPullsItemCommentsRequestBuilderGetRequestConfiguration)([]i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.PullRequestReviewCommentable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.SendCollection(ctx, requestInfo, i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CreatePullRequestReviewCommentFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    val := make([]i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.PullRequestReviewCommentable, len(res))
    for i, v := range res {
        if v != nil {
            val[i] = v.(i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.PullRequestReviewCommentable)
        }
    }
    return val, nil
}
// Post creates a review comment in the pull request diff. To add a regular comment to a pull request timeline, see "[Create an issue comment](https://docs.github.com/rest/issues/comments#create-an-issue-comment)." We recommend creating a review comment using `line`, `side`, and optionally `start_line` and `start_side` if your comment applies to more than one line in the pull request diff.The `position` parameter is deprecated. If you use `position`, the `line`, `side`, `start_line`, and `start_side` parameters are not required.**Note:** The position value equals the number of lines down from the first "@@" hunk header in the file you want to add a comment. The line just below the "@@" line is position 1, the next line is position 2, and so on. The position in the diff continues to increase through lines of whitespace and additional hunks until the beginning of a new file.This endpoint triggers [notifications](https://docs.github.com/github/managing-subscriptions-and-notifications-on-github/about-notifications). Creating content too quickly using this endpoint may result in secondary rate limiting. For more information, see "[Rate limits for the API](https://docs.github.com/rest/overview/rate-limits-for-the-rest-api#about-secondary-rate-limits)"and "[Best practices for using the REST API](https://docs.github.com/rest/guides/best-practices-for-using-the-rest-api)."
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/pulls/comments#create-a-review-comment-for-a-pull-request
func (m *ItemItemPullsItemCommentsRequestBuilder) Post(ctx context.Context, body ItemItemPullsItemCommentsPostRequestBodyable, requestConfiguration *ItemItemPullsItemCommentsRequestBuilderPostRequestConfiguration)(i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.PullRequestReviewCommentable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "403": i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CreateBasicErrorFromDiscriminatorValue,
        "422": i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CreateValidationErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CreatePullRequestReviewCommentFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.PullRequestReviewCommentable), nil
}
// ToGetRequestInformation lists all review comments for a pull request. By default, review comments are in ascending order by ID.
func (m *ItemItemPullsItemCommentsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemItemPullsItemCommentsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation creates a review comment in the pull request diff. To add a regular comment to a pull request timeline, see "[Create an issue comment](https://docs.github.com/rest/issues/comments#create-an-issue-comment)." We recommend creating a review comment using `line`, `side`, and optionally `start_line` and `start_side` if your comment applies to more than one line in the pull request diff.The `position` parameter is deprecated. If you use `position`, the `line`, `side`, `start_line`, and `start_side` parameters are not required.**Note:** The position value equals the number of lines down from the first "@@" hunk header in the file you want to add a comment. The line just below the "@@" line is position 1, the next line is position 2, and so on. The position in the diff continues to increase through lines of whitespace and additional hunks until the beginning of a new file.This endpoint triggers [notifications](https://docs.github.com/github/managing-subscriptions-and-notifications-on-github/about-notifications). Creating content too quickly using this endpoint may result in secondary rate limiting. For more information, see "[Rate limits for the API](https://docs.github.com/rest/overview/rate-limits-for-the-rest-api#about-secondary-rate-limits)"and "[Best practices for using the REST API](https://docs.github.com/rest/guides/best-practices-for-using-the-rest-api)."
func (m *ItemItemPullsItemCommentsRequestBuilder) ToPostRequestInformation(ctx context.Context, body ItemItemPullsItemCommentsPostRequestBodyable, requestConfiguration *ItemItemPullsItemCommentsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
func (m *ItemItemPullsItemCommentsRequestBuilder) WithUrl(rawUrl string)(*ItemItemPullsItemCommentsRequestBuilder) {
    return NewItemItemPullsItemCommentsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
