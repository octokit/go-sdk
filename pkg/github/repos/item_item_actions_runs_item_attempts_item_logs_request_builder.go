package repos

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemItemActionsRunsItemAttemptsItemLogsRequestBuilder builds and executes requests for operations under \repos\{repos-id}\{Owner-id}\actions\runs\{run_id}\attempts\{attempt_number}\logs
type ItemItemActionsRunsItemAttemptsItemLogsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// NewItemItemActionsRunsItemAttemptsItemLogsRequestBuilderInternal instantiates a new LogsRequestBuilder and sets the default values.
func NewItemItemActionsRunsItemAttemptsItemLogsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemActionsRunsItemAttemptsItemLogsRequestBuilder) {
    m := &ItemItemActionsRunsItemAttemptsItemLogsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/repos/{repos%2Did}/{Owner%2Did}/actions/runs/{run_id}/attempts/{attempt_number}/logs", pathParameters),
    }
    return m
}
// NewItemItemActionsRunsItemAttemptsItemLogsRequestBuilder instantiates a new LogsRequestBuilder and sets the default values.
func NewItemItemActionsRunsItemAttemptsItemLogsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemActionsRunsItemAttemptsItemLogsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemItemActionsRunsItemAttemptsItemLogsRequestBuilderInternal(urlParams, requestAdapter)
}
// Get gets a redirect URL to download an archive of log files for a specific workflow run attempt. This link expires after1 minute. Look for `Location:` in the response header to find the URL for the download. Anyone with read access tothe repository can use this endpoint. If the repository is private you must use an access token with the `repo` scope.GitHub Apps must have the `actions:read` permission to use this endpoint.
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/actions/workflow-runs#download-workflow-run-attempt-logs
func (m *ItemItemActionsRunsItemAttemptsItemLogsRequestBuilder) Get(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])([]byte, error) {
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
// ToGetRequestInformation gets a redirect URL to download an archive of log files for a specific workflow run attempt. This link expires after1 minute. Look for `Location:` in the response header to find the URL for the download. Anyone with read access tothe repository can use this endpoint. If the repository is private you must use an access token with the `repo` scope.GitHub Apps must have the `actions:read` permission to use this endpoint.
func (m *ItemItemActionsRunsItemAttemptsItemLogsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ConfigureRequestInformation(requestInfo, requestConfiguration)
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
func (m *ItemItemActionsRunsItemAttemptsItemLogsRequestBuilder) WithUrl(rawUrl string)(*ItemItemActionsRunsItemAttemptsItemLogsRequestBuilder) {
    return NewItemItemActionsRunsItemAttemptsItemLogsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
