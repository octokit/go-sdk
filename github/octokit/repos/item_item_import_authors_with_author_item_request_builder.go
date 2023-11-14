package repos

import (
    "context"
    i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e "github.com/octokit/go-sdk/github/octokit/models"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemItemImportAuthorsWithAuthor_ItemRequestBuilder builds and executes requests for operations under \repos\{owner}\{repo}\import\authors\{author_id}
type ItemItemImportAuthorsWithAuthor_ItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemItemImportAuthorsWithAuthor_ItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemItemImportAuthorsWithAuthor_ItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemItemImportAuthorsWithAuthor_ItemRequestBuilderInternal instantiates a new WithAuthor_ItemRequestBuilder and sets the default values.
func NewItemItemImportAuthorsWithAuthor_ItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemImportAuthorsWithAuthor_ItemRequestBuilder) {
    m := &ItemItemImportAuthorsWithAuthor_ItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/repos/{owner}/{repo}/import/authors/{author_id}", pathParameters),
    }
    return m
}
// NewItemItemImportAuthorsWithAuthor_ItemRequestBuilder instantiates a new WithAuthor_ItemRequestBuilder and sets the default values.
func NewItemItemImportAuthorsWithAuthor_ItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemImportAuthorsWithAuthor_ItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemItemImportAuthorsWithAuthor_ItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Patch update an author's identity for the import. Your application can continue updating authors any time before you pushnew commits to the repository.**Warning:** Due to very low levels of usage and available alternatives, this endpoint is deprecated and will no longer be available from 00:00 UTC on April 12, 2024. For more details and alternatives, see the [changelog](https://gh.io/source-imports-api-deprecation).
// Deprecated: 
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/migrations/source-imports#map-a-commit-author
func (m *ItemItemImportAuthorsWithAuthor_ItemRequestBuilder) Patch(ctx context.Context, body ItemItemImportAuthorsItemWithAuthor_PatchRequestBodyable, requestConfiguration *ItemItemImportAuthorsWithAuthor_ItemRequestBuilderPatchRequestConfiguration)(i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.PorterAuthorable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "404": i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.CreateBasicErrorFromDiscriminatorValue,
        "422": i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.CreateValidationErrorFromDiscriminatorValue,
        "503": i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.CreateBasicErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.CreatePorterAuthorFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.PorterAuthorable), nil
}
// ToPatchRequestInformation update an author's identity for the import. Your application can continue updating authors any time before you pushnew commits to the repository.**Warning:** Due to very low levels of usage and available alternatives, this endpoint is deprecated and will no longer be available from 00:00 UTC on April 12, 2024. For more details and alternatives, see the [changelog](https://gh.io/source-imports-api-deprecation).
// Deprecated: 
func (m *ItemItemImportAuthorsWithAuthor_ItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ItemItemImportAuthorsItemWithAuthor_PatchRequestBodyable, requestConfiguration *ItemItemImportAuthorsWithAuthor_ItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
    requestInfo.Headers.TryAdd("Accept", "application/json;q=1")
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// Deprecated: 
func (m *ItemItemImportAuthorsWithAuthor_ItemRequestBuilder) WithUrl(rawUrl string)(*ItemItemImportAuthorsWithAuthor_ItemRequestBuilder) {
    return NewItemItemImportAuthorsWithAuthor_ItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
