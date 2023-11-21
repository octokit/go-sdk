package user

import (
    "context"
    i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e "github.com/octokit/go-sdk/github/octokit/models"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// CodespacesItemPublishRequestBuilder builds and executes requests for operations under \user\codespaces\{codespace_name}\publish
type CodespacesItemPublishRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// CodespacesItemPublishRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CodespacesItemPublishRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewCodespacesItemPublishRequestBuilderInternal instantiates a new PublishRequestBuilder and sets the default values.
func NewCodespacesItemPublishRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CodespacesItemPublishRequestBuilder) {
    m := &CodespacesItemPublishRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/user/codespaces/{codespace_name}/publish", pathParameters),
    }
    return m
}
// NewCodespacesItemPublishRequestBuilder instantiates a new PublishRequestBuilder and sets the default values.
func NewCodespacesItemPublishRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CodespacesItemPublishRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewCodespacesItemPublishRequestBuilderInternal(urlParams, requestAdapter)
}
// Post publishes an unpublished codespace, creating a new repository and assigning it to the codespace.The codespace's token is granted write permissions to the repository, allowing the user to push their changes.This will fail for a codespace that is already published, meaning it has an associated repository.You must authenticate using a personal access token with the `codespace` scope to use this endpoint.GitHub Apps must have write access to the `codespaces` repository permission to use this endpoint.
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/codespaces/codespaces#create-a-repository-from-an-unpublished-codespace
func (m *CodespacesItemPublishRequestBuilder) Post(ctx context.Context, body CodespacesItemPublishPostRequestBodyable, requestConfiguration *CodespacesItemPublishRequestBuilderPostRequestConfiguration)(i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.CodespaceWithFullRepositoryable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "401": i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.CreateBasicErrorFromDiscriminatorValue,
        "403": i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.CreateBasicErrorFromDiscriminatorValue,
        "404": i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.CreateBasicErrorFromDiscriminatorValue,
        "422": i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.CreateValidationErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.CreateCodespaceWithFullRepositoryFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.CodespaceWithFullRepositoryable), nil
}
// ToPostRequestInformation publishes an unpublished codespace, creating a new repository and assigning it to the codespace.The codespace's token is granted write permissions to the repository, allowing the user to push their changes.This will fail for a codespace that is already published, meaning it has an associated repository.You must authenticate using a personal access token with the `codespace` scope to use this endpoint.GitHub Apps must have write access to the `codespaces` repository permission to use this endpoint.
func (m *CodespacesItemPublishRequestBuilder) ToPostRequestInformation(ctx context.Context, body CodespacesItemPublishPostRequestBodyable, requestConfiguration *CodespacesItemPublishRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
    requestInfo.Headers.TryAdd("Accept", "application/json")
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
func (m *CodespacesItemPublishRequestBuilder) WithUrl(rawUrl string)(*CodespacesItemPublishRequestBuilder) {
    return NewCodespacesItemPublishRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
