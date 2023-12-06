package repos

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i000736ae6dd74f01081193e4f903216bc2bd2954ed818433b986f45d581ed035 "github.com/octokit/go-sdk/github/octokit/models"
)

// ItemItemActionsRunsItemAttemptsWithAttempt_numberItemRequestBuilder builds and executes requests for operations under \repos\{repos-id}\{Owner-id}\actions\runs\{run_id}\attempts\{attempt_number}
type ItemItemActionsRunsItemAttemptsWithAttempt_numberItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemItemActionsRunsItemAttemptsWithAttempt_numberItemRequestBuilderGetQueryParameters gets a specific workflow run attempt. Anyone with read access to the repositorycan use this endpoint. If the repository is private you must use an access tokenwith the `repo` scope. GitHub Apps must have the `actions:read` permission touse this endpoint.
type ItemItemActionsRunsItemAttemptsWithAttempt_numberItemRequestBuilderGetQueryParameters struct {
    // If `true` pull requests are omitted from the response (empty array).
    Exclude_pull_requests *bool `uriparametername:"exclude_pull_requests"`
}
// ItemItemActionsRunsItemAttemptsWithAttempt_numberItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemItemActionsRunsItemAttemptsWithAttempt_numberItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemItemActionsRunsItemAttemptsWithAttempt_numberItemRequestBuilderGetQueryParameters
}
// NewItemItemActionsRunsItemAttemptsWithAttempt_numberItemRequestBuilderInternal instantiates a new WithAttempt_numberItemRequestBuilder and sets the default values.
func NewItemItemActionsRunsItemAttemptsWithAttempt_numberItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemActionsRunsItemAttemptsWithAttempt_numberItemRequestBuilder) {
    m := &ItemItemActionsRunsItemAttemptsWithAttempt_numberItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/repos/{repos%2Did}/{Owner%2Did}/actions/runs/{run_id}/attempts/{attempt_number}{?exclude_pull_requests*}", pathParameters),
    }
    return m
}
// NewItemItemActionsRunsItemAttemptsWithAttempt_numberItemRequestBuilder instantiates a new WithAttempt_numberItemRequestBuilder and sets the default values.
func NewItemItemActionsRunsItemAttemptsWithAttempt_numberItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemActionsRunsItemAttemptsWithAttempt_numberItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemItemActionsRunsItemAttemptsWithAttempt_numberItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Get gets a specific workflow run attempt. Anyone with read access to the repositorycan use this endpoint. If the repository is private you must use an access tokenwith the `repo` scope. GitHub Apps must have the `actions:read` permission touse this endpoint.
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/actions/workflow-runs#get-a-workflow-run-attempt
func (m *ItemItemActionsRunsItemAttemptsWithAttempt_numberItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemItemActionsRunsItemAttemptsWithAttempt_numberItemRequestBuilderGetRequestConfiguration)(i000736ae6dd74f01081193e4f903216bc2bd2954ed818433b986f45d581ed035.WorkflowRunable, error) {
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
func (m *ItemItemActionsRunsItemAttemptsWithAttempt_numberItemRequestBuilder) Jobs()(*ItemItemActionsRunsItemAttemptsItemJobsRequestBuilder) {
    return NewItemItemActionsRunsItemAttemptsItemJobsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Logs the logs property
func (m *ItemItemActionsRunsItemAttemptsWithAttempt_numberItemRequestBuilder) Logs()(*ItemItemActionsRunsItemAttemptsItemLogsRequestBuilder) {
    return NewItemItemActionsRunsItemAttemptsItemLogsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToGetRequestInformation gets a specific workflow run attempt. Anyone with read access to the repositorycan use this endpoint. If the repository is private you must use an access tokenwith the `repo` scope. GitHub Apps must have the `actions:read` permission touse this endpoint.
func (m *ItemItemActionsRunsItemAttemptsWithAttempt_numberItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemItemActionsRunsItemAttemptsWithAttempt_numberItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *ItemItemActionsRunsItemAttemptsWithAttempt_numberItemRequestBuilder) WithUrl(rawUrl string)(*ItemItemActionsRunsItemAttemptsWithAttempt_numberItemRequestBuilder) {
    return NewItemItemActionsRunsItemAttemptsWithAttempt_numberItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
