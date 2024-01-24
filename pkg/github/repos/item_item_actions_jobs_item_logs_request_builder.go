package repos

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemItemActionsJobsItemLogsRequestBuilder builds and executes requests for operations under \repos\{repos-id}\{Owner-id}\actions\jobs\{job_id}\logs
type ItemItemActionsJobsItemLogsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// NewItemItemActionsJobsItemLogsRequestBuilderInternal instantiates a new LogsRequestBuilder and sets the default values.
func NewItemItemActionsJobsItemLogsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemActionsJobsItemLogsRequestBuilder) {
    m := &ItemItemActionsJobsItemLogsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/repos/{repos%2Did}/{Owner%2Did}/actions/jobs/{job_id}/logs", pathParameters),
    }
    return m
}
// NewItemItemActionsJobsItemLogsRequestBuilder instantiates a new LogsRequestBuilder and sets the default values.
func NewItemItemActionsJobsItemLogsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemActionsJobsItemLogsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemItemActionsJobsItemLogsRequestBuilderInternal(urlParams, requestAdapter)
}
// Get gets a redirect URL to download a plain text file of logs for a workflow job. This link expires after 1 minute. Lookfor `Location:` in the response header to find the URL for the download. Anyone with read access to the repository canuse this endpoint. If the repository is private you must use an access token with the `repo` scope. GitHub Apps musthave the `actions:read` permission to use this endpoint.
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/actions/workflow-jobs#download-job-logs-for-a-workflow-run
func (m *ItemItemActionsJobsItemLogsRequestBuilder) Get(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])([]byte, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.SendPrimitive(ctx, requestInfo, "[]byte", nil)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.([]byte), nil
}
// ToGetRequestInformation gets a redirect URL to download a plain text file of logs for a workflow job. This link expires after 1 minute. Lookfor `Location:` in the response header to find the URL for the download. Anyone with read access to the repository canuse this endpoint. If the repository is private you must use an access token with the `repo` scope. GitHub Apps musthave the `actions:read` permission to use this endpoint.
func (m *ItemItemActionsJobsItemLogsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ConfigureRequestInformation(requestInfo, requestConfiguration)
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
func (m *ItemItemActionsJobsItemLogsRequestBuilder) WithUrl(rawUrl string)(*ItemItemActionsJobsItemLogsRequestBuilder) {
    return NewItemItemActionsJobsItemLogsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
