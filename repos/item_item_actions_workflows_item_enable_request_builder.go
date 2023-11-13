package repos

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemItemActionsWorkflowsItemEnableRequestBuilder builds and executes requests for operations under \repos\{owner}\{repo}\actions\workflows\{workflow_id}\enable
type ItemItemActionsWorkflowsItemEnableRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemItemActionsWorkflowsItemEnableRequestBuilderPutRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemItemActionsWorkflowsItemEnableRequestBuilderPutRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemItemActionsWorkflowsItemEnableRequestBuilderInternal instantiates a new EnableRequestBuilder and sets the default values.
func NewItemItemActionsWorkflowsItemEnableRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemActionsWorkflowsItemEnableRequestBuilder) {
    m := &ItemItemActionsWorkflowsItemEnableRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/repos/{owner}/{repo}/actions/workflows/{workflow_id}/enable", pathParameters),
    }
    return m
}
// NewItemItemActionsWorkflowsItemEnableRequestBuilder instantiates a new EnableRequestBuilder and sets the default values.
func NewItemItemActionsWorkflowsItemEnableRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemActionsWorkflowsItemEnableRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemItemActionsWorkflowsItemEnableRequestBuilderInternal(urlParams, requestAdapter)
}
// Put enables a workflow and sets the `state` of the workflow to `active`. You can replace `workflow_id` with the workflow file name. For example, you could use `main.yaml`.You must authenticate using an access token with the `repo` scope to use this endpoint. GitHub Apps must have the `actions:write` permission to use this endpoint.
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/actions/workflows#enable-a-workflow
func (m *ItemItemActionsWorkflowsItemEnableRequestBuilder) Put(ctx context.Context, requestConfiguration *ItemItemActionsWorkflowsItemEnableRequestBuilderPutRequestConfiguration)(error) {
    requestInfo, err := m.ToPutRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    err = m.BaseRequestBuilder.RequestAdapter.SendNoContent(ctx, requestInfo, nil)
    if err != nil {
        return err
    }
    return nil
}
// ToPutRequestInformation enables a workflow and sets the `state` of the workflow to `active`. You can replace `workflow_id` with the workflow file name. For example, you could use `main.yaml`.You must authenticate using an access token with the `repo` scope to use this endpoint. GitHub Apps must have the `actions:write` permission to use this endpoint.
func (m *ItemItemActionsWorkflowsItemEnableRequestBuilder) ToPutRequestInformation(ctx context.Context, requestConfiguration *ItemItemActionsWorkflowsItemEnableRequestBuilderPutRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PUT
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
func (m *ItemItemActionsWorkflowsItemEnableRequestBuilder) WithUrl(rawUrl string)(*ItemItemActionsWorkflowsItemEnableRequestBuilder) {
    return NewItemItemActionsWorkflowsItemEnableRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
