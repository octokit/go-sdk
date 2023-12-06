package repos

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i000736ae6dd74f01081193e4f903216bc2bd2954ed818433b986f45d581ed035 "github.com/octokit/go-sdk/github/octokit/models"
)

// ItemItemIssuesItemAssigneesWithAssigneeItemRequestBuilder builds and executes requests for operations under \repos\{repos-id}\{Owner-id}\issues\{issue_number}\assignees\{assignee}
type ItemItemIssuesItemAssigneesWithAssigneeItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemItemIssuesItemAssigneesWithAssigneeItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemItemIssuesItemAssigneesWithAssigneeItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemItemIssuesItemAssigneesWithAssigneeItemRequestBuilderInternal instantiates a new WithAssigneeItemRequestBuilder and sets the default values.
func NewItemItemIssuesItemAssigneesWithAssigneeItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemIssuesItemAssigneesWithAssigneeItemRequestBuilder) {
    m := &ItemItemIssuesItemAssigneesWithAssigneeItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/repos/{repos%2Did}/{Owner%2Did}/issues/{issue_number}/assignees/{assignee}", pathParameters),
    }
    return m
}
// NewItemItemIssuesItemAssigneesWithAssigneeItemRequestBuilder instantiates a new WithAssigneeItemRequestBuilder and sets the default values.
func NewItemItemIssuesItemAssigneesWithAssigneeItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemIssuesItemAssigneesWithAssigneeItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemItemIssuesItemAssigneesWithAssigneeItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Get checks if a user has permission to be assigned to a specific issue.If the `assignee` can be assigned to this issue, a `204` status code with no content is returned.Otherwise a `404` status code is returned.
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/issues/assignees#check-if-a-user-can-be-assigned-to-a-issue
func (m *ItemItemIssuesItemAssigneesWithAssigneeItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemItemIssuesItemAssigneesWithAssigneeItemRequestBuilderGetRequestConfiguration)(error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "404": i000736ae6dd74f01081193e4f903216bc2bd2954ed818433b986f45d581ed035.CreateBasicErrorFromDiscriminatorValue,
    }
    err = m.BaseRequestBuilder.RequestAdapter.SendNoContent(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// ToGetRequestInformation checks if a user has permission to be assigned to a specific issue.If the `assignee` can be assigned to this issue, a `204` status code with no content is returned.Otherwise a `404` status code is returned.
func (m *ItemItemIssuesItemAssigneesWithAssigneeItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemItemIssuesItemAssigneesWithAssigneeItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
func (m *ItemItemIssuesItemAssigneesWithAssigneeItemRequestBuilder) WithUrl(rawUrl string)(*ItemItemIssuesItemAssigneesWithAssigneeItemRequestBuilder) {
    return NewItemItemIssuesItemAssigneesWithAssigneeItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
