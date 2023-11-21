package repos

import (
    "context"
    i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e "github.com/octokit/go-sdk/github/octokit/models"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemItemIssuesItemEventsRequestBuilder builds and executes requests for operations under \repos\{owner}\{repo}\issues\{issue_number}\events
type ItemItemIssuesItemEventsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemItemIssuesItemEventsRequestBuilderGetQueryParameters lists all events for an issue.
type ItemItemIssuesItemEventsRequestBuilderGetQueryParameters struct {
    // Page number of the results to fetch.
    Page *int32 `uriparametername:"page"`
    // The number of results per page (max 100).
    Per_page *int32 `uriparametername:"per_page"`
}
// ItemItemIssuesItemEventsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemItemIssuesItemEventsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemItemIssuesItemEventsRequestBuilderGetQueryParameters
}
// NewItemItemIssuesItemEventsRequestBuilderInternal instantiates a new EventsRequestBuilder and sets the default values.
func NewItemItemIssuesItemEventsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemIssuesItemEventsRequestBuilder) {
    m := &ItemItemIssuesItemEventsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/repos/{owner}/{repo}/issues/{issue_number}/events{?per_page*,page*}", pathParameters),
    }
    return m
}
// NewItemItemIssuesItemEventsRequestBuilder instantiates a new EventsRequestBuilder and sets the default values.
func NewItemItemIssuesItemEventsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemIssuesItemEventsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemItemIssuesItemEventsRequestBuilderInternal(urlParams, requestAdapter)
}
// Get lists all events for an issue.
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/issues/events#list-issue-events
func (m *ItemItemIssuesItemEventsRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemItemIssuesItemEventsRequestBuilderGetRequestConfiguration)([]i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.IssueEventForIssueable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "410": i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.CreateBasicErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.SendCollection(ctx, requestInfo, i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.CreateIssueEventForIssueFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    val := make([]i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.IssueEventForIssueable, len(res))
    for i, v := range res {
        if v != nil {
            val[i] = v.(i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.IssueEventForIssueable)
        }
    }
    return val, nil
}
// ToGetRequestInformation lists all events for an issue.
func (m *ItemItemIssuesItemEventsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemItemIssuesItemEventsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *ItemItemIssuesItemEventsRequestBuilder) WithUrl(rawUrl string)(*ItemItemIssuesItemEventsRequestBuilder) {
    return NewItemItemIssuesItemEventsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
