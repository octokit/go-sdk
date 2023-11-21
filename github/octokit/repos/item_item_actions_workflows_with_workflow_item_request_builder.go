package repos

import (
    "context"
    i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e "github.com/octokit/go-sdk/github/octokit/models"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemItemActionsWorkflowsWithWorkflow_ItemRequestBuilder builds and executes requests for operations under \repos\{owner}\{repo}\actions\workflows\{workflow_id}
type ItemItemActionsWorkflowsWithWorkflow_ItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemItemActionsWorkflowsWithWorkflow_ItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemItemActionsWorkflowsWithWorkflow_ItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemItemActionsWorkflowsWithWorkflow_ItemRequestBuilderInternal instantiates a new WithWorkflow_ItemRequestBuilder and sets the default values.
func NewItemItemActionsWorkflowsWithWorkflow_ItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemActionsWorkflowsWithWorkflow_ItemRequestBuilder) {
    m := &ItemItemActionsWorkflowsWithWorkflow_ItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/repos/{owner}/{repo}/actions/workflows/{workflow_id}", pathParameters),
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
func (m *ItemItemActionsWorkflowsWithWorkflow_ItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemItemActionsWorkflowsWithWorkflow_ItemRequestBuilderGetRequestConfiguration)(i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.Workflowable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.CreateWorkflowFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.Workflowable), nil
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
func (m *ItemItemActionsWorkflowsWithWorkflow_ItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemItemActionsWorkflowsWithWorkflow_ItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *ItemItemActionsWorkflowsWithWorkflow_ItemRequestBuilder) WithUrl(rawUrl string)(*ItemItemActionsWorkflowsWithWorkflow_ItemRequestBuilder) {
    return NewItemItemActionsWorkflowsWithWorkflow_ItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
