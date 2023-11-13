package repos

import (
    "context"
    i33bcfa37677dfd8b71eaf22ad3bfd140ba1702e7a8418558499c22418bf2fef2 "github.com/octokit/go-sdk/models"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemItemAssigneesWithAssigneeItemRequestBuilder builds and executes requests for operations under \repos\{owner}\{repo}\assignees\{assignee}
type ItemItemAssigneesWithAssigneeItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemItemAssigneesWithAssigneeItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemItemAssigneesWithAssigneeItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemItemAssigneesWithAssigneeItemRequestBuilderInternal instantiates a new WithAssigneeItemRequestBuilder and sets the default values.
func NewItemItemAssigneesWithAssigneeItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemAssigneesWithAssigneeItemRequestBuilder) {
    m := &ItemItemAssigneesWithAssigneeItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/repos/{owner}/{repo}/assignees/{assignee}", pathParameters),
    }
    return m
}
// NewItemItemAssigneesWithAssigneeItemRequestBuilder instantiates a new WithAssigneeItemRequestBuilder and sets the default values.
func NewItemItemAssigneesWithAssigneeItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemAssigneesWithAssigneeItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemItemAssigneesWithAssigneeItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Get checks if a user has permission to be assigned to an issue in this repository.If the `assignee` can be assigned to issues in the repository, a `204` header with no content is returned.Otherwise a `404` status code is returned.
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/issues/assignees#check-if-a-user-can-be-assigned
func (m *ItemItemAssigneesWithAssigneeItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemItemAssigneesWithAssigneeItemRequestBuilderGetRequestConfiguration)(error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "404": i33bcfa37677dfd8b71eaf22ad3bfd140ba1702e7a8418558499c22418bf2fef2.CreateBasicErrorFromDiscriminatorValue,
    }
    err = m.BaseRequestBuilder.RequestAdapter.SendNoContent(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// ToGetRequestInformation checks if a user has permission to be assigned to an issue in this repository.If the `assignee` can be assigned to issues in the repository, a `204` header with no content is returned.Otherwise a `404` status code is returned.
func (m *ItemItemAssigneesWithAssigneeItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemItemAssigneesWithAssigneeItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
func (m *ItemItemAssigneesWithAssigneeItemRequestBuilder) WithUrl(rawUrl string)(*ItemItemAssigneesWithAssigneeItemRequestBuilder) {
    return NewItemItemAssigneesWithAssigneeItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
