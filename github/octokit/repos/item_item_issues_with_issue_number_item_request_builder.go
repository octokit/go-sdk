package repos

import (
    "context"
    i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e "github.com/octokit/go-sdk/github/octokit/models"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemItemIssuesWithIssue_numberItemRequestBuilder builds and executes requests for operations under \repos\{owner}\{repo}\issues\{issue_number}
type ItemItemIssuesWithIssue_numberItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemItemIssuesWithIssue_numberItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemItemIssuesWithIssue_numberItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ItemItemIssuesWithIssue_numberItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemItemIssuesWithIssue_numberItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// Assignees the assignees property
func (m *ItemItemIssuesWithIssue_numberItemRequestBuilder) Assignees()(*ItemItemIssuesItemAssigneesRequestBuilder) {
    return NewItemItemIssuesItemAssigneesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Comments the comments property
func (m *ItemItemIssuesWithIssue_numberItemRequestBuilder) Comments()(*ItemItemIssuesItemCommentsRequestBuilder) {
    return NewItemItemIssuesItemCommentsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewItemItemIssuesWithIssue_numberItemRequestBuilderInternal instantiates a new WithIssue_numberItemRequestBuilder and sets the default values.
func NewItemItemIssuesWithIssue_numberItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemIssuesWithIssue_numberItemRequestBuilder) {
    m := &ItemItemIssuesWithIssue_numberItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/repos/{owner}/{repo}/issues/{issue_number}", pathParameters),
    }
    return m
}
// NewItemItemIssuesWithIssue_numberItemRequestBuilder instantiates a new WithIssue_numberItemRequestBuilder and sets the default values.
func NewItemItemIssuesWithIssue_numberItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemIssuesWithIssue_numberItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemItemIssuesWithIssue_numberItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Events the events property
func (m *ItemItemIssuesWithIssue_numberItemRequestBuilder) Events()(*ItemItemIssuesItemEventsRequestBuilder) {
    return NewItemItemIssuesItemEventsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get the API returns a [`301 Moved Permanently` status](https://docs.github.com/rest/guides/best-practices-for-using-the-rest-api#follow-redirects) if the issue was[transferred](https://docs.github.com/articles/transferring-an-issue-to-another-repository/) to another repository. Ifthe issue was transferred to or deleted from a repository where the authenticated user lacks read access, the APIreturns a `404 Not Found` status. If the issue was deleted from a repository where the authenticated user has readaccess, the API returns a `410 Gone` status. To receive webhook events for transferred and deleted issues, subscribeto the [`issues`](https://docs.github.com/webhooks/event-payloads/#issues) webhook.**Note**: GitHub's REST API considers every pull request an issue, but not every issue is a pull request. For thisreason, "Issues" endpoints may return both issues and pull requests in the response. You can identify pull requests bythe `pull_request` key. Be aware that the `id` of a pull request returned from "Issues" endpoints will be an _issue id_. To find out the pullrequest id, use the "[List pull requests](https://docs.github.com/rest/pulls/pulls#list-pull-requests)" endpoint.
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/issues/issues#get-an-issue
func (m *ItemItemIssuesWithIssue_numberItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemItemIssuesWithIssue_numberItemRequestBuilderGetRequestConfiguration)(i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.Issueable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "404": i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.CreateBasicErrorFromDiscriminatorValue,
        "410": i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.CreateBasicErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.CreateIssueFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.Issueable), nil
}
// Labels the labels property
func (m *ItemItemIssuesWithIssue_numberItemRequestBuilder) Labels()(*ItemItemIssuesItemLabelsRequestBuilder) {
    return NewItemItemIssuesItemLabelsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Lock the lock property
func (m *ItemItemIssuesWithIssue_numberItemRequestBuilder) Lock()(*ItemItemIssuesItemLockRequestBuilder) {
    return NewItemItemIssuesItemLockRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Patch issue owners and users with push access can edit an issue.
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/issues/issues#update-an-issue
func (m *ItemItemIssuesWithIssue_numberItemRequestBuilder) Patch(ctx context.Context, body ItemItemIssuesItemWithIssue_numberPatchRequestBodyable, requestConfiguration *ItemItemIssuesWithIssue_numberItemRequestBuilderPatchRequestConfiguration)(i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.Issueable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "403": i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.CreateBasicErrorFromDiscriminatorValue,
        "404": i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.CreateBasicErrorFromDiscriminatorValue,
        "410": i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.CreateBasicErrorFromDiscriminatorValue,
        "422": i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.CreateValidationErrorFromDiscriminatorValue,
        "503": i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.CreateIssue503ErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.CreateIssueFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.Issueable), nil
}
// Reactions the reactions property
func (m *ItemItemIssuesWithIssue_numberItemRequestBuilder) Reactions()(*ItemItemIssuesItemReactionsRequestBuilder) {
    return NewItemItemIssuesItemReactionsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Timeline the timeline property
func (m *ItemItemIssuesWithIssue_numberItemRequestBuilder) Timeline()(*ItemItemIssuesItemTimelineRequestBuilder) {
    return NewItemItemIssuesItemTimelineRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToGetRequestInformation the API returns a [`301 Moved Permanently` status](https://docs.github.com/rest/guides/best-practices-for-using-the-rest-api#follow-redirects) if the issue was[transferred](https://docs.github.com/articles/transferring-an-issue-to-another-repository/) to another repository. Ifthe issue was transferred to or deleted from a repository where the authenticated user lacks read access, the APIreturns a `404 Not Found` status. If the issue was deleted from a repository where the authenticated user has readaccess, the API returns a `410 Gone` status. To receive webhook events for transferred and deleted issues, subscribeto the [`issues`](https://docs.github.com/webhooks/event-payloads/#issues) webhook.**Note**: GitHub's REST API considers every pull request an issue, but not every issue is a pull request. For thisreason, "Issues" endpoints may return both issues and pull requests in the response. You can identify pull requests bythe `pull_request` key. Be aware that the `id` of a pull request returned from "Issues" endpoints will be an _issue id_. To find out the pullrequest id, use the "[List pull requests](https://docs.github.com/rest/pulls/pulls#list-pull-requests)" endpoint.
func (m *ItemItemIssuesWithIssue_numberItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemItemIssuesWithIssue_numberItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToPatchRequestInformation issue owners and users with push access can edit an issue.
func (m *ItemItemIssuesWithIssue_numberItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ItemItemIssuesItemWithIssue_numberPatchRequestBodyable, requestConfiguration *ItemItemIssuesWithIssue_numberItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
    requestInfo.Headers.TryAdd("Accept", "application/json")
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
func (m *ItemItemIssuesWithIssue_numberItemRequestBuilder) WithUrl(rawUrl string)(*ItemItemIssuesWithIssue_numberItemRequestBuilder) {
    return NewItemItemIssuesWithIssue_numberItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
