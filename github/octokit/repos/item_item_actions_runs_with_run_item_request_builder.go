package repos

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i000736ae6dd74f01081193e4f903216bc2bd2954ed818433b986f45d581ed035 "github.com/octokit/go-sdk/github/octokit/models"
)

// ItemItemActionsRunsWithRun_ItemRequestBuilder builds and executes requests for operations under \repos\{repos-id}\{Owner-id}\actions\runs\{run_id}
type ItemItemActionsRunsWithRun_ItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemItemActionsRunsWithRun_ItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemItemActionsRunsWithRun_ItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ItemItemActionsRunsWithRun_ItemRequestBuilderGetQueryParameters gets a specific workflow run. Anyone with read access to the repository can use this endpoint. If the repository is private you must use an access token with the `repo` scope. GitHub Apps must have the `actions:read` permission to use this endpoint.
type ItemItemActionsRunsWithRun_ItemRequestBuilderGetQueryParameters struct {
    // If `true` pull requests are omitted from the response (empty array).
    Exclude_pull_requests *bool `uriparametername:"exclude_pull_requests"`
}
// ItemItemActionsRunsWithRun_ItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemItemActionsRunsWithRun_ItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemItemActionsRunsWithRun_ItemRequestBuilderGetQueryParameters
}
// Approvals the approvals property
func (m *ItemItemActionsRunsWithRun_ItemRequestBuilder) Approvals()(*ItemItemActionsRunsItemApprovalsRequestBuilder) {
    return NewItemItemActionsRunsItemApprovalsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Approve the approve property
func (m *ItemItemActionsRunsWithRun_ItemRequestBuilder) Approve()(*ItemItemActionsRunsItemApproveRequestBuilder) {
    return NewItemItemActionsRunsItemApproveRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Artifacts the artifacts property
func (m *ItemItemActionsRunsWithRun_ItemRequestBuilder) Artifacts()(*ItemItemActionsRunsItemArtifactsRequestBuilder) {
    return NewItemItemActionsRunsItemArtifactsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Attempts the attempts property
func (m *ItemItemActionsRunsWithRun_ItemRequestBuilder) Attempts()(*ItemItemActionsRunsItemAttemptsRequestBuilder) {
    return NewItemItemActionsRunsItemAttemptsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Cancel the cancel property
func (m *ItemItemActionsRunsWithRun_ItemRequestBuilder) Cancel()(*ItemItemActionsRunsItemCancelRequestBuilder) {
    return NewItemItemActionsRunsItemCancelRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewItemItemActionsRunsWithRun_ItemRequestBuilderInternal instantiates a new WithRun_ItemRequestBuilder and sets the default values.
func NewItemItemActionsRunsWithRun_ItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemActionsRunsWithRun_ItemRequestBuilder) {
    m := &ItemItemActionsRunsWithRun_ItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/repos/{repos%2Did}/{Owner%2Did}/actions/runs/{run_id}{?exclude_pull_requests*}", pathParameters),
    }
    return m
}
// NewItemItemActionsRunsWithRun_ItemRequestBuilder instantiates a new WithRun_ItemRequestBuilder and sets the default values.
func NewItemItemActionsRunsWithRun_ItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemActionsRunsWithRun_ItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemItemActionsRunsWithRun_ItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete a specific workflow run. Anyone with write access to the repository can use this endpoint. If the repository isprivate you must use an access token with the `repo` scope. GitHub Apps must have the `actions:write` permission to usethis endpoint.
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/actions/workflow-runs#delete-a-workflow-run
func (m *ItemItemActionsRunsWithRun_ItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *ItemItemActionsRunsWithRun_ItemRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.ToDeleteRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    err = m.BaseRequestBuilder.RequestAdapter.SendNoContent(ctx, requestInfo, nil)
    if err != nil {
        return err
    }
    return nil
}
// Deployment_protection_rule the deployment_protection_rule property
func (m *ItemItemActionsRunsWithRun_ItemRequestBuilder) Deployment_protection_rule()(*ItemItemActionsRunsItemDeployment_protection_ruleRequestBuilder) {
    return NewItemItemActionsRunsItemDeployment_protection_ruleRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ForceCancel the forceCancel property
func (m *ItemItemActionsRunsWithRun_ItemRequestBuilder) ForceCancel()(*ItemItemActionsRunsItemForceCancelRequestBuilder) {
    return NewItemItemActionsRunsItemForceCancelRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get gets a specific workflow run. Anyone with read access to the repository can use this endpoint. If the repository is private you must use an access token with the `repo` scope. GitHub Apps must have the `actions:read` permission to use this endpoint.
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/actions/workflow-runs#get-a-workflow-run
func (m *ItemItemActionsRunsWithRun_ItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemItemActionsRunsWithRun_ItemRequestBuilderGetRequestConfiguration)(i000736ae6dd74f01081193e4f903216bc2bd2954ed818433b986f45d581ed035.WorkflowRunable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i000736ae6dd74f01081193e4f903216bc2bd2954ed818433b986f45d581ed035.CreateWorkflowRunFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i000736ae6dd74f01081193e4f903216bc2bd2954ed818433b986f45d581ed035.WorkflowRunable), nil
}
// Jobs the jobs property
func (m *ItemItemActionsRunsWithRun_ItemRequestBuilder) Jobs()(*ItemItemActionsRunsItemJobsRequestBuilder) {
    return NewItemItemActionsRunsItemJobsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Logs the logs property
func (m *ItemItemActionsRunsWithRun_ItemRequestBuilder) Logs()(*ItemItemActionsRunsItemLogsRequestBuilder) {
    return NewItemItemActionsRunsItemLogsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Pending_deployments the pending_deployments property
func (m *ItemItemActionsRunsWithRun_ItemRequestBuilder) Pending_deployments()(*ItemItemActionsRunsItemPending_deploymentsRequestBuilder) {
    return NewItemItemActionsRunsItemPending_deploymentsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Rerun the rerun property
func (m *ItemItemActionsRunsWithRun_ItemRequestBuilder) Rerun()(*ItemItemActionsRunsItemRerunRequestBuilder) {
    return NewItemItemActionsRunsItemRerunRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// RerunFailedJobs the rerunFailedJobs property
func (m *ItemItemActionsRunsWithRun_ItemRequestBuilder) RerunFailedJobs()(*ItemItemActionsRunsItemRerunFailedJobsRequestBuilder) {
    return NewItemItemActionsRunsItemRerunFailedJobsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Timing the timing property
func (m *ItemItemActionsRunsWithRun_ItemRequestBuilder) Timing()(*ItemItemActionsRunsItemTimingRequestBuilder) {
    return NewItemItemActionsRunsItemTimingRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToDeleteRequestInformation delete a specific workflow run. Anyone with write access to the repository can use this endpoint. If the repository isprivate you must use an access token with the `repo` scope. GitHub Apps must have the `actions:write` permission to usethis endpoint.
func (m *ItemItemActionsRunsWithRun_ItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *ItemItemActionsRunsWithRun_ItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// ToGetRequestInformation gets a specific workflow run. Anyone with read access to the repository can use this endpoint. If the repository is private you must use an access token with the `repo` scope. GitHub Apps must have the `actions:read` permission to use this endpoint.
func (m *ItemItemActionsRunsWithRun_ItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemItemActionsRunsWithRun_ItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
func (m *ItemItemActionsRunsWithRun_ItemRequestBuilder) WithUrl(rawUrl string)(*ItemItemActionsRunsWithRun_ItemRequestBuilder) {
    return NewItemItemActionsRunsWithRun_ItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
