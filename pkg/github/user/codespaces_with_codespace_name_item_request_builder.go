package user

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6 "github.com/octokit/go-sdk/pkg/github/models"
)

// CodespacesWithCodespace_nameItemRequestBuilder builds and executes requests for operations under \user\codespaces\{codespace_name}
type CodespacesWithCodespace_nameItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// CodespacesWithCodespace_nameItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CodespacesWithCodespace_nameItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// CodespacesWithCodespace_nameItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CodespacesWithCodespace_nameItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// CodespacesWithCodespace_nameItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CodespacesWithCodespace_nameItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewCodespacesWithCodespace_nameItemRequestBuilderInternal instantiates a new WithCodespace_nameItemRequestBuilder and sets the default values.
func NewCodespacesWithCodespace_nameItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CodespacesWithCodespace_nameItemRequestBuilder) {
    m := &CodespacesWithCodespace_nameItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/user/codespaces/{codespace_name}", pathParameters),
    }
    return m
}
// NewCodespacesWithCodespace_nameItemRequestBuilder instantiates a new WithCodespace_nameItemRequestBuilder and sets the default values.
func NewCodespacesWithCodespace_nameItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CodespacesWithCodespace_nameItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewCodespacesWithCodespace_nameItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete deletes a user's codespace.You must authenticate using an access token with the `codespace` scope to use this endpoint.To use this endpoint with GitHub Apps:- The app must be authenticated on behalf of the user. For more information, see "[Authenticating with a GitHub App on behalf of a user](https://docs.github.com/apps/creating-github-apps/authenticating-with-a-github-app/authenticating-with-a-github-app-on-behalf-of-a-user)."- The app must have write access to the `codespaces` repository permission.
// Deprecated: This method is obsolete. Use DeleteAsWithCodespace_nameDeleteResponse instead.
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/codespaces/codespaces#delete-a-codespace-for-the-authenticated-user
func (m *CodespacesWithCodespace_nameItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *CodespacesWithCodespace_nameItemRequestBuilderDeleteRequestConfiguration)(CodespacesItemWithCodespace_nameResponseable, error) {
    requestInfo, err := m.ToDeleteRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "401": i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CreateBasicErrorFromDiscriminatorValue,
        "403": i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CreateBasicErrorFromDiscriminatorValue,
        "404": i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CreateBasicErrorFromDiscriminatorValue,
        "500": i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CreateBasicErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateCodespacesItemWithCodespace_nameResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(CodespacesItemWithCodespace_nameResponseable), nil
}
// DeleteAsWithCodespace_nameDeleteResponse deletes a user's codespace.You must authenticate using an access token with the `codespace` scope to use this endpoint.To use this endpoint with GitHub Apps:- The app must be authenticated on behalf of the user. For more information, see "[Authenticating with a GitHub App on behalf of a user](https://docs.github.com/apps/creating-github-apps/authenticating-with-a-github-app/authenticating-with-a-github-app-on-behalf-of-a-user)."- The app must have write access to the `codespaces` repository permission.
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/codespaces/codespaces#delete-a-codespace-for-the-authenticated-user
func (m *CodespacesWithCodespace_nameItemRequestBuilder) DeleteAsWithCodespace_nameDeleteResponse(ctx context.Context, requestConfiguration *CodespacesWithCodespace_nameItemRequestBuilderDeleteRequestConfiguration)(CodespacesItemWithCodespace_nameDeleteResponseable, error) {
    requestInfo, err := m.ToDeleteRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "401": i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CreateBasicErrorFromDiscriminatorValue,
        "403": i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CreateBasicErrorFromDiscriminatorValue,
        "404": i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CreateBasicErrorFromDiscriminatorValue,
        "500": i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CreateBasicErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateCodespacesItemWithCodespace_nameDeleteResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(CodespacesItemWithCodespace_nameDeleteResponseable), nil
}
// Exports the exports property
func (m *CodespacesWithCodespace_nameItemRequestBuilder) Exports()(*CodespacesItemExportsRequestBuilder) {
    return NewCodespacesItemExportsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get gets information about a user's codespace.You must authenticate using an access token with the `codespace` scope to use this endpoint.To use this endpoint with GitHub Apps:- The app must be authenticated on behalf of the user. For more information, see "[Authenticating with a GitHub App on behalf of a user](https://docs.github.com/apps/creating-github-apps/authenticating-with-a-github-app/authenticating-with-a-github-app-on-behalf-of-a-user)."- The app must have read access to the `codespaces` repository permission.
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/codespaces/codespaces#get-a-codespace-for-the-authenticated-user
func (m *CodespacesWithCodespace_nameItemRequestBuilder) Get(ctx context.Context, requestConfiguration *CodespacesWithCodespace_nameItemRequestBuilderGetRequestConfiguration)(i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.Codespaceable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "401": i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CreateBasicErrorFromDiscriminatorValue,
        "403": i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CreateBasicErrorFromDiscriminatorValue,
        "404": i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CreateBasicErrorFromDiscriminatorValue,
        "500": i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CreateBasicErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CreateCodespaceFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.Codespaceable), nil
}
// Machines the machines property
func (m *CodespacesWithCodespace_nameItemRequestBuilder) Machines()(*CodespacesItemMachinesRequestBuilder) {
    return NewCodespacesItemMachinesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Patch updates a codespace owned by the authenticated user. Currently only the codespace's machine type and recent folders can be modified using this endpoint.If you specify a new machine type it will be applied the next time your codespace is started.You must authenticate using an access token with the `codespace` scope to use this endpoint.To use this endpoint with GitHub Apps:- The app must be authenticated on behalf of the user. For more information, see "[Authenticating with a GitHub App on behalf of a user](https://docs.github.com/apps/creating-github-apps/authenticating-with-a-github-app/authenticating-with-a-github-app-on-behalf-of-a-user)."- The app must have write access to the `codespaces` repository permission.
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/codespaces/codespaces#update-a-codespace-for-the-authenticated-user
func (m *CodespacesWithCodespace_nameItemRequestBuilder) Patch(ctx context.Context, body CodespacesItemWithCodespace_namePatchRequestBodyable, requestConfiguration *CodespacesWithCodespace_nameItemRequestBuilderPatchRequestConfiguration)(i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.Codespaceable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "401": i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CreateBasicErrorFromDiscriminatorValue,
        "403": i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CreateBasicErrorFromDiscriminatorValue,
        "404": i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CreateBasicErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CreateCodespaceFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.Codespaceable), nil
}
// Publish the publish property
func (m *CodespacesWithCodespace_nameItemRequestBuilder) Publish()(*CodespacesItemPublishRequestBuilder) {
    return NewCodespacesItemPublishRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Start the start property
func (m *CodespacesWithCodespace_nameItemRequestBuilder) Start()(*CodespacesItemStartRequestBuilder) {
    return NewCodespacesItemStartRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Stop the stop property
func (m *CodespacesWithCodespace_nameItemRequestBuilder) Stop()(*CodespacesItemStopRequestBuilder) {
    return NewCodespacesItemStopRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToDeleteRequestInformation deletes a user's codespace.You must authenticate using an access token with the `codespace` scope to use this endpoint.To use this endpoint with GitHub Apps:- The app must be authenticated on behalf of the user. For more information, see "[Authenticating with a GitHub App on behalf of a user](https://docs.github.com/apps/creating-github-apps/authenticating-with-a-github-app/authenticating-with-a-github-app-on-behalf-of-a-user)."- The app must have write access to the `codespaces` repository permission.
func (m *CodespacesWithCodespace_nameItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *CodespacesWithCodespace_nameItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation gets information about a user's codespace.You must authenticate using an access token with the `codespace` scope to use this endpoint.To use this endpoint with GitHub Apps:- The app must be authenticated on behalf of the user. For more information, see "[Authenticating with a GitHub App on behalf of a user](https://docs.github.com/apps/creating-github-apps/authenticating-with-a-github-app/authenticating-with-a-github-app-on-behalf-of-a-user)."- The app must have read access to the `codespaces` repository permission.
func (m *CodespacesWithCodespace_nameItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *CodespacesWithCodespace_nameItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToPatchRequestInformation updates a codespace owned by the authenticated user. Currently only the codespace's machine type and recent folders can be modified using this endpoint.If you specify a new machine type it will be applied the next time your codespace is started.You must authenticate using an access token with the `codespace` scope to use this endpoint.To use this endpoint with GitHub Apps:- The app must be authenticated on behalf of the user. For more information, see "[Authenticating with a GitHub App on behalf of a user](https://docs.github.com/apps/creating-github-apps/authenticating-with-a-github-app/authenticating-with-a-github-app-on-behalf-of-a-user)."- The app must have write access to the `codespaces` repository permission.
func (m *CodespacesWithCodespace_nameItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body CodespacesItemWithCodespace_namePatchRequestBodyable, requestConfiguration *CodespacesWithCodespace_nameItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
func (m *CodespacesWithCodespace_nameItemRequestBuilder) WithUrl(rawUrl string)(*CodespacesWithCodespace_nameItemRequestBuilder) {
    return NewCodespacesWithCodespace_nameItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
