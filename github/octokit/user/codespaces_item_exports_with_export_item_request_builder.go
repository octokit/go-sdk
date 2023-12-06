package user

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i000736ae6dd74f01081193e4f903216bc2bd2954ed818433b986f45d581ed035 "github.com/octokit/go-sdk/github/octokit/models"
)

// CodespacesItemExportsWithExport_ItemRequestBuilder builds and executes requests for operations under \user\codespaces\{codespace_name}\exports\{export_id}
type CodespacesItemExportsWithExport_ItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// CodespacesItemExportsWithExport_ItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CodespacesItemExportsWithExport_ItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewCodespacesItemExportsWithExport_ItemRequestBuilderInternal instantiates a new WithExport_ItemRequestBuilder and sets the default values.
func NewCodespacesItemExportsWithExport_ItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CodespacesItemExportsWithExport_ItemRequestBuilder) {
    m := &CodespacesItemExportsWithExport_ItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/user/codespaces/{codespace_name}/exports/{export_id}", pathParameters),
    }
    return m
}
// NewCodespacesItemExportsWithExport_ItemRequestBuilder instantiates a new WithExport_ItemRequestBuilder and sets the default values.
func NewCodespacesItemExportsWithExport_ItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CodespacesItemExportsWithExport_ItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewCodespacesItemExportsWithExport_ItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Get gets information about an export of a codespace.You must authenticate using a personal access token with the `codespace` scope to use this endpoint.GitHub Apps must have read access to the `codespaces_lifecycle_admin` repository permission to use this endpoint.
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/codespaces/codespaces#get-details-about-a-codespace-export
func (m *CodespacesItemExportsWithExport_ItemRequestBuilder) Get(ctx context.Context, requestConfiguration *CodespacesItemExportsWithExport_ItemRequestBuilderGetRequestConfiguration)(i000736ae6dd74f01081193e4f903216bc2bd2954ed818433b986f45d581ed035.CodespaceExportDetailsable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "404": i000736ae6dd74f01081193e4f903216bc2bd2954ed818433b986f45d581ed035.CreateBasicErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i000736ae6dd74f01081193e4f903216bc2bd2954ed818433b986f45d581ed035.CreateCodespaceExportDetailsFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i000736ae6dd74f01081193e4f903216bc2bd2954ed818433b986f45d581ed035.CodespaceExportDetailsable), nil
}
// ToGetRequestInformation gets information about an export of a codespace.You must authenticate using a personal access token with the `codespace` scope to use this endpoint.GitHub Apps must have read access to the `codespaces_lifecycle_admin` repository permission to use this endpoint.
func (m *CodespacesItemExportsWithExport_ItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *CodespacesItemExportsWithExport_ItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
func (m *CodespacesItemExportsWithExport_ItemRequestBuilder) WithUrl(rawUrl string)(*CodespacesItemExportsWithExport_ItemRequestBuilder) {
    return NewCodespacesItemExportsWithExport_ItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
