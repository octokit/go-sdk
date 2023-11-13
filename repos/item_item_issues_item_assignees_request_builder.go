package repos

import (
    "context"
    i33bcfa37677dfd8b71eaf22ad3bfd140ba1702e7a8418558499c22418bf2fef2 "github.com/octokit/go-sdk/models"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemItemIssuesItemAssigneesRequestBuilder builds and executes requests for operations under \repos\{owner}\{repo}\issues\{issue_number}\assignees
type ItemItemIssuesItemAssigneesRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemItemIssuesItemAssigneesRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemItemIssuesItemAssigneesRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ItemItemIssuesItemAssigneesRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemItemIssuesItemAssigneesRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ByAssignee gets an item from the go-sdk.repos.item.item.issues.item.assignees.item collection
func (m *ItemItemIssuesItemAssigneesRequestBuilder) ByAssignee(assignee string)(*ItemItemIssuesItemAssigneesWithAssigneeItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if assignee != "" {
        urlTplParams["assignee"] = assignee
    }
    return NewItemItemIssuesItemAssigneesWithAssigneeItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewItemItemIssuesItemAssigneesRequestBuilderInternal instantiates a new AssigneesRequestBuilder and sets the default values.
func NewItemItemIssuesItemAssigneesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemIssuesItemAssigneesRequestBuilder) {
    m := &ItemItemIssuesItemAssigneesRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/repos/{owner}/{repo}/issues/{issue_number}/assignees", pathParameters),
    }
    return m
}
// NewItemItemIssuesItemAssigneesRequestBuilder instantiates a new AssigneesRequestBuilder and sets the default values.
func NewItemItemIssuesItemAssigneesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemIssuesItemAssigneesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemItemIssuesItemAssigneesRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete removes one or more assignees from an issue.
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/issues/assignees#remove-assignees-from-an-issue
func (m *ItemItemIssuesItemAssigneesRequestBuilder) Delete(ctx context.Context, body ItemItemIssuesItemAssigneesDeleteRequestBodyable, requestConfiguration *ItemItemIssuesItemAssigneesRequestBuilderDeleteRequestConfiguration)(i33bcfa37677dfd8b71eaf22ad3bfd140ba1702e7a8418558499c22418bf2fef2.Issueable, error) {
    requestInfo, err := m.ToDeleteRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i33bcfa37677dfd8b71eaf22ad3bfd140ba1702e7a8418558499c22418bf2fef2.CreateIssueFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i33bcfa37677dfd8b71eaf22ad3bfd140ba1702e7a8418558499c22418bf2fef2.Issueable), nil
}
// Post adds up to 10 assignees to an issue. Users already assigned to an issue are not replaced.
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/issues/assignees#add-assignees-to-an-issue
func (m *ItemItemIssuesItemAssigneesRequestBuilder) Post(ctx context.Context, body ItemItemIssuesItemAssigneesPostRequestBodyable, requestConfiguration *ItemItemIssuesItemAssigneesRequestBuilderPostRequestConfiguration)(i33bcfa37677dfd8b71eaf22ad3bfd140ba1702e7a8418558499c22418bf2fef2.Issueable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i33bcfa37677dfd8b71eaf22ad3bfd140ba1702e7a8418558499c22418bf2fef2.CreateIssueFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i33bcfa37677dfd8b71eaf22ad3bfd140ba1702e7a8418558499c22418bf2fef2.Issueable), nil
}
// ToDeleteRequestInformation removes one or more assignees from an issue.
func (m *ItemItemIssuesItemAssigneesRequestBuilder) ToDeleteRequestInformation(ctx context.Context, body ItemItemIssuesItemAssigneesDeleteRequestBodyable, requestConfiguration *ItemItemIssuesItemAssigneesRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE
    requestInfo.Headers.TryAdd("Accept", "application/json;q=1")
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    return requestInfo, nil
}
// ToPostRequestInformation adds up to 10 assignees to an issue. Users already assigned to an issue are not replaced.
func (m *ItemItemIssuesItemAssigneesRequestBuilder) ToPostRequestInformation(ctx context.Context, body ItemItemIssuesItemAssigneesPostRequestBodyable, requestConfiguration *ItemItemIssuesItemAssigneesRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
    requestInfo.Headers.TryAdd("Accept", "application/json;q=1")
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
func (m *ItemItemIssuesItemAssigneesRequestBuilder) WithUrl(rawUrl string)(*ItemItemIssuesItemAssigneesRequestBuilder) {
    return NewItemItemIssuesItemAssigneesRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
