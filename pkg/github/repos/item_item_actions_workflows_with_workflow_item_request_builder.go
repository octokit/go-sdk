package repos

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6 "github.com/octokit/go-sdk/pkg/github/models"
)

// ItemItemActionsWorkflowsWithWorkflow_ItemRequestBuilder builds and executes requests for operations under \repos\{repos-id}\{Owner-id}\actions\workflows\{workflow_id}
type ItemItemActionsWorkflowsWithWorkflow_ItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// NewItemItemActionsWorkflowsWithWorkflow_ItemRequestBuilderInternal instantiates a new WithWorkflow_ItemRequestBuilder and sets the default values.
func NewItemItemActionsWorkflowsWithWorkflow_ItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemActionsWorkflowsWithWorkflow_ItemRequestBuilder) {
    m := &ItemItemActionsWorkflowsWithWorkflow_ItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/repos/{repos%2Did}/{Owner%2Did}/actions/workflows/{workflow_id}", pathParameters),
    }
    return m
}
// NewItemItemActionsWorkflowsWithWorkflow_ItemRequestBuilder instantiates a new WithWorkflow_ItemRequestBuilder and sets the default values.
func NewItemItemActionsWorkflowsWithWorkflow_ItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemActionsWorkflowsWithWorkflow_ItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemItemActionsWorkflowsWithWorkflow_ItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Disable the disable property
func (m *ItemItemActionsWorkflowsWithWorkflow_ItemRequestBuilder) Disable()(*ItemItemActionsWorkflowsItemDisableRequestBuilder) {
    return NewItemItemActionsWorkflowsItemDisableRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Dispatches the dispatches property
func (m *ItemItemActionsWorkflowsWithWorkflow_ItemRequestBuilder) Dispatches()(*ItemItemActionsWorkflowsItemDispatchesRequestBuilder) {
    return NewItemItemActionsWorkflowsItemDispatchesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Enable the enable property
func (m *ItemItemActionsWorkflowsWithWorkflow_ItemRequestBuilder) Enable()(*ItemItemActionsWorkflowsItemEnableRequestBuilder) {
    return NewItemItemActionsWorkflowsItemEnableRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get gets a specific workflow. You can replace `workflow_id` with the workflow file name. For example, you could use `main.yaml`. Anyone with read access to the repository can use this endpoint. If the repository is private you must use an access token with the `repo` scope. GitHub Apps must have the `actions:read` permission to use this endpoint.
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/actions/workflows#get-a-workflow
func (m *ItemItemActionsWorkflowsWithWorkflow_ItemRequestBuilder) Get(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.Workflowable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CreateWorkflowFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.Workflowable), nil
}
// Runs the runs property
func (m *ItemItemActionsWorkflowsWithWorkflow_ItemRequestBuilder) Runs()(*ItemItemActionsWorkflowsItemRunsRequestBuilder) {
    return NewItemItemActionsWorkflowsItemRunsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Timing the timing property
func (m *ItemItemActionsWorkflowsWithWorkflow_ItemRequestBuilder) Timing()(*ItemItemActionsWorkflowsItemTimingRequestBuilder) {
    return NewItemItemActionsWorkflowsItemTimingRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToGetRequestInformation gets a specific workflow. You can replace `workflow_id` with the workflow file name. For example, you could use `main.yaml`. Anyone with read access to the repository can use this endpoint. If the repository is private you must use an access token with the `repo` scope. GitHub Apps must have the `actions:read` permission to use this endpoint.
func (m *ItemItemActionsWorkflowsWithWorkflow_ItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ConfigureRequestInformation(requestInfo, requestConfiguration)
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
func (m *ItemItemActionsWorkflowsWithWorkflow_ItemRequestBuilder) WithUrl(rawUrl string)(*ItemItemActionsWorkflowsWithWorkflow_ItemRequestBuilder) {
    return NewItemItemActionsWorkflowsWithWorkflow_ItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
