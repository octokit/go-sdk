package repos

import (
    "context"
    i33bcfa37677dfd8b71eaf22ad3bfd140ba1702e7a8418558499c22418bf2fef2 "github.com/octokit/go-sdk/models"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemItemActionsRunsItemForceCancelRequestBuilder builds and executes requests for operations under \repos\{owner}\{repo}\actions\runs\{run_id}\force-cancel
type ItemItemActionsRunsItemForceCancelRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemItemActionsRunsItemForceCancelRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemItemActionsRunsItemForceCancelRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemItemActionsRunsItemForceCancelRequestBuilderInternal instantiates a new ForceCancelRequestBuilder and sets the default values.
func NewItemItemActionsRunsItemForceCancelRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemActionsRunsItemForceCancelRequestBuilder) {
    m := &ItemItemActionsRunsItemForceCancelRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/repos/{owner}/{repo}/actions/runs/{run_id}/force-cancel", pathParameters),
    }
    return m
}
// NewItemItemActionsRunsItemForceCancelRequestBuilder instantiates a new ForceCancelRequestBuilder and sets the default values.
func NewItemItemActionsRunsItemForceCancelRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemActionsRunsItemForceCancelRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemItemActionsRunsItemForceCancelRequestBuilderInternal(urlParams, requestAdapter)
}
// Post cancels a workflow run and bypasses conditions that would otherwise cause a workflow execution to continue, such as an `always()` condition on a job.You should only use this endpoint to cancel a workflow run when the workflow run is not responding to [`POST /repos/{owner}/{repo}/actions/runs/{run_id}/cancel`](/rest/actions/workflow-runs#cancel-a-workflow-run).You must authenticate using an access token with the `repo` scope to use this endpoint.GitHub Apps must have the `actions:write` permission to use this endpoint.
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/actions/workflow-runs#force-cancel-a-workflow-run
func (m *ItemItemActionsRunsItemForceCancelRequestBuilder) Post(ctx context.Context, requestConfiguration *ItemItemActionsRunsItemForceCancelRequestBuilderPostRequestConfiguration)(i33bcfa37677dfd8b71eaf22ad3bfd140ba1702e7a8418558499c22418bf2fef2.EmptyObjectable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "409": i33bcfa37677dfd8b71eaf22ad3bfd140ba1702e7a8418558499c22418bf2fef2.CreateBasicErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i33bcfa37677dfd8b71eaf22ad3bfd140ba1702e7a8418558499c22418bf2fef2.CreateEmptyObjectFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i33bcfa37677dfd8b71eaf22ad3bfd140ba1702e7a8418558499c22418bf2fef2.EmptyObjectable), nil
}
// ToPostRequestInformation cancels a workflow run and bypasses conditions that would otherwise cause a workflow execution to continue, such as an `always()` condition on a job.You should only use this endpoint to cancel a workflow run when the workflow run is not responding to [`POST /repos/{owner}/{repo}/actions/runs/{run_id}/cancel`](/rest/actions/workflow-runs#cancel-a-workflow-run).You must authenticate using an access token with the `repo` scope to use this endpoint.GitHub Apps must have the `actions:write` permission to use this endpoint.
func (m *ItemItemActionsRunsItemForceCancelRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *ItemItemActionsRunsItemForceCancelRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
    requestInfo.Headers.TryAdd("Accept", "application/json;q=1")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
func (m *ItemItemActionsRunsItemForceCancelRequestBuilder) WithUrl(rawUrl string)(*ItemItemActionsRunsItemForceCancelRequestBuilder) {
    return NewItemItemActionsRunsItemForceCancelRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
