package repos

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i000736ae6dd74f01081193e4f903216bc2bd2954ed818433b986f45d581ed035 "github.com/octokit/go-sdk/github/octokit/models"
)

// ItemItemActionsArtifactsItemWithArchive_formatItemRequestBuilder builds and executes requests for operations under \repos\{repos-id}\{Owner-id}\actions\artifacts\{artifact_id}\{archive_format}
type ItemItemActionsArtifactsItemWithArchive_formatItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemItemActionsArtifactsItemWithArchive_formatItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemItemActionsArtifactsItemWithArchive_formatItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemItemActionsArtifactsItemWithArchive_formatItemRequestBuilderInternal instantiates a new WithArchive_formatItemRequestBuilder and sets the default values.
func NewItemItemActionsArtifactsItemWithArchive_formatItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemActionsArtifactsItemWithArchive_formatItemRequestBuilder) {
    m := &ItemItemActionsArtifactsItemWithArchive_formatItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/repos/{repos%2Did}/{Owner%2Did}/actions/artifacts/{artifact_id}/{archive_format}", pathParameters),
    }
    return m
}
// NewItemItemActionsArtifactsItemWithArchive_formatItemRequestBuilder instantiates a new WithArchive_formatItemRequestBuilder and sets the default values.
func NewItemItemActionsArtifactsItemWithArchive_formatItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemActionsArtifactsItemWithArchive_formatItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemItemActionsArtifactsItemWithArchive_formatItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Get gets a redirect URL to download an archive for a repository. This URL expires after 1 minute. Look for `Location:` inthe response header to find the URL for the download. The `:archive_format` must be `zip`.You must authenticate using an access token with the `repo` scope to use this endpoint.GitHub Apps must have the `actions:read` permission to use this endpoint.
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/actions/artifacts#download-an-artifact
func (m *ItemItemActionsArtifactsItemWithArchive_formatItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemItemActionsArtifactsItemWithArchive_formatItemRequestBuilderGetRequestConfiguration)([]byte, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "410": i000736ae6dd74f01081193e4f903216bc2bd2954ed818433b986f45d581ed035.CreateBasicErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.SendPrimitive(ctx, requestInfo, "[]byte", errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.([]byte), nil
}
// ToGetRequestInformation gets a redirect URL to download an archive for a repository. This URL expires after 1 minute. Look for `Location:` inthe response header to find the URL for the download. The `:archive_format` must be `zip`.You must authenticate using an access token with the `repo` scope to use this endpoint.GitHub Apps must have the `actions:read` permission to use this endpoint.
func (m *ItemItemActionsArtifactsItemWithArchive_formatItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemItemActionsArtifactsItemWithArchive_formatItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
func (m *ItemItemActionsArtifactsItemWithArchive_formatItemRequestBuilder) WithUrl(rawUrl string)(*ItemItemActionsArtifactsItemWithArchive_formatItemRequestBuilder) {
    return NewItemItemActionsArtifactsItemWithArchive_formatItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
