package repos

import (
    "context"
    i33bcfa37677dfd8b71eaf22ad3bfd140ba1702e7a8418558499c22418bf2fef2 "github.com/octokit/go-sdk/models"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemItemActionsRunnersGenerateJitconfigRequestBuilder builds and executes requests for operations under \repos\{owner}\{repo}\actions\runners\generate-jitconfig
type ItemItemActionsRunnersGenerateJitconfigRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemItemActionsRunnersGenerateJitconfigRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemItemActionsRunnersGenerateJitconfigRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemItemActionsRunnersGenerateJitconfigRequestBuilderInternal instantiates a new GenerateJitconfigRequestBuilder and sets the default values.
func NewItemItemActionsRunnersGenerateJitconfigRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemActionsRunnersGenerateJitconfigRequestBuilder) {
    m := &ItemItemActionsRunnersGenerateJitconfigRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/repos/{owner}/{repo}/actions/runners/generate-jitconfig", pathParameters),
    }
    return m
}
// NewItemItemActionsRunnersGenerateJitconfigRequestBuilder instantiates a new GenerateJitconfigRequestBuilder and sets the default values.
func NewItemItemActionsRunnersGenerateJitconfigRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemActionsRunnersGenerateJitconfigRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemItemActionsRunnersGenerateJitconfigRequestBuilderInternal(urlParams, requestAdapter)
}
// Post generates a configuration that can be passed to the runner application at startup.You must authenticate using an access token with the `repo` scope to use this endpoint.GitHub Apps must have the `administration` permission for repositories and the `organization_self_hosted_runners` permission for organizations.Authenticated users must have admin access to repositories or organizations, or the `manage_runners:enterprise` scope for enterprises, to use these endpoints.
// Deprecated: This method is obsolete. Use PostAsGenerateJitconfigPostResponse instead.
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/actions/self-hosted-runners#create-configuration-for-a-just-in-time-runner-for-a-repository
func (m *ItemItemActionsRunnersGenerateJitconfigRequestBuilder) Post(ctx context.Context, body ItemItemActionsRunnersGenerateJitconfigPostRequestBodyable, requestConfiguration *ItemItemActionsRunnersGenerateJitconfigRequestBuilderPostRequestConfiguration)(ItemItemActionsRunnersGenerateJitconfigResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "404": i33bcfa37677dfd8b71eaf22ad3bfd140ba1702e7a8418558499c22418bf2fef2.CreateBasicErrorFromDiscriminatorValue,
        "422": i33bcfa37677dfd8b71eaf22ad3bfd140ba1702e7a8418558499c22418bf2fef2.CreateValidationErrorSimpleFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateItemItemActionsRunnersGenerateJitconfigResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ItemItemActionsRunnersGenerateJitconfigResponseable), nil
}
// PostAsGenerateJitconfigPostResponse generates a configuration that can be passed to the runner application at startup.You must authenticate using an access token with the `repo` scope to use this endpoint.GitHub Apps must have the `administration` permission for repositories and the `organization_self_hosted_runners` permission for organizations.Authenticated users must have admin access to repositories or organizations, or the `manage_runners:enterprise` scope for enterprises, to use these endpoints.
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/actions/self-hosted-runners#create-configuration-for-a-just-in-time-runner-for-a-repository
func (m *ItemItemActionsRunnersGenerateJitconfigRequestBuilder) PostAsGenerateJitconfigPostResponse(ctx context.Context, body ItemItemActionsRunnersGenerateJitconfigPostRequestBodyable, requestConfiguration *ItemItemActionsRunnersGenerateJitconfigRequestBuilderPostRequestConfiguration)(ItemItemActionsRunnersGenerateJitconfigPostResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "404": i33bcfa37677dfd8b71eaf22ad3bfd140ba1702e7a8418558499c22418bf2fef2.CreateBasicErrorFromDiscriminatorValue,
        "422": i33bcfa37677dfd8b71eaf22ad3bfd140ba1702e7a8418558499c22418bf2fef2.CreateValidationErrorSimpleFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateItemItemActionsRunnersGenerateJitconfigPostResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ItemItemActionsRunnersGenerateJitconfigPostResponseable), nil
}
// ToPostRequestInformation generates a configuration that can be passed to the runner application at startup.You must authenticate using an access token with the `repo` scope to use this endpoint.GitHub Apps must have the `administration` permission for repositories and the `organization_self_hosted_runners` permission for organizations.Authenticated users must have admin access to repositories or organizations, or the `manage_runners:enterprise` scope for enterprises, to use these endpoints.
func (m *ItemItemActionsRunnersGenerateJitconfigRequestBuilder) ToPostRequestInformation(ctx context.Context, body ItemItemActionsRunnersGenerateJitconfigPostRequestBodyable, requestConfiguration *ItemItemActionsRunnersGenerateJitconfigRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
    requestInfo.Headers.TryAdd("Accept", "application/json;q=1")
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
func (m *ItemItemActionsRunnersGenerateJitconfigRequestBuilder) WithUrl(rawUrl string)(*ItemItemActionsRunnersGenerateJitconfigRequestBuilder) {
    return NewItemItemActionsRunnersGenerateJitconfigRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
