package user

import (
    "context"
    i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e "github.com/octokit/go-sdk/github/octokit/models"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// CodespacesItemExportsRequestBuilder builds and executes requests for operations under \user\codespaces\{codespace_name}\exports
type CodespacesItemExportsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// CodespacesItemExportsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CodespacesItemExportsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ByExport_id gets an item from the octokit.user.codespaces.item.exports.item collection
func (m *CodespacesItemExportsRequestBuilder) ByExport_id(export_id string)(*CodespacesItemExportsWithExport_ItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if export_id != "" {
        urlTplParams["export_id"] = export_id
    }
    return NewCodespacesItemExportsWithExport_ItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewCodespacesItemExportsRequestBuilderInternal instantiates a new ExportsRequestBuilder and sets the default values.
func NewCodespacesItemExportsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CodespacesItemExportsRequestBuilder) {
    m := &CodespacesItemExportsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/user/codespaces/{codespace_name}/exports", pathParameters),
    }
    return m
}
// NewCodespacesItemExportsRequestBuilder instantiates a new ExportsRequestBuilder and sets the default values.
func NewCodespacesItemExportsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CodespacesItemExportsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewCodespacesItemExportsRequestBuilderInternal(urlParams, requestAdapter)
}
// Post triggers an export of the specified codespace and returns a URL and ID where the status of the export can be monitored.If changes cannot be pushed to the codespace's repository, they will be pushed to a new or previously-existing fork instead.You must authenticate using a personal access token with the `codespace` scope to use this endpoint.GitHub Apps must have write access to the `codespaces_lifecycle_admin` repository permission to use this endpoint.
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/codespaces/codespaces#export-a-codespace-for-the-authenticated-user
func (m *CodespacesItemExportsRequestBuilder) Post(ctx context.Context, requestConfiguration *CodespacesItemExportsRequestBuilderPostRequestConfiguration)(i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.CodespaceExportDetailsable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "401": i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.CreateBasicErrorFromDiscriminatorValue,
        "403": i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.CreateBasicErrorFromDiscriminatorValue,
        "404": i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.CreateBasicErrorFromDiscriminatorValue,
        "422": i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.CreateValidationErrorFromDiscriminatorValue,
        "500": i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.CreateBasicErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.CreateCodespaceExportDetailsFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.CodespaceExportDetailsable), nil
}
// ToPostRequestInformation triggers an export of the specified codespace and returns a URL and ID where the status of the export can be monitored.If changes cannot be pushed to the codespace's repository, they will be pushed to a new or previously-existing fork instead.You must authenticate using a personal access token with the `codespace` scope to use this endpoint.GitHub Apps must have write access to the `codespaces_lifecycle_admin` repository permission to use this endpoint.
func (m *CodespacesItemExportsRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *CodespacesItemExportsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *CodespacesItemExportsRequestBuilder) WithUrl(rawUrl string)(*CodespacesItemExportsRequestBuilder) {
    return NewCodespacesItemExportsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
