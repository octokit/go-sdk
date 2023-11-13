package orgs

import (
    "context"
    i33bcfa37677dfd8b71eaf22ad3bfd140ba1702e7a8418558499c22418bf2fef2 "github.com/octokit/go-sdk/models"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemActionsRunnersGenerateJitconfigRequestBuilder builds and executes requests for operations under \orgs\{org}\actions\runners\generate-jitconfig
type ItemActionsRunnersGenerateJitconfigRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemActionsRunnersGenerateJitconfigRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemActionsRunnersGenerateJitconfigRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemActionsRunnersGenerateJitconfigRequestBuilderInternal instantiates a new GenerateJitconfigRequestBuilder and sets the default values.
func NewItemActionsRunnersGenerateJitconfigRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemActionsRunnersGenerateJitconfigRequestBuilder) {
    m := &ItemActionsRunnersGenerateJitconfigRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/orgs/{org}/actions/runners/generate-jitconfig", pathParameters),
    }
    return m
}
// NewItemActionsRunnersGenerateJitconfigRequestBuilder instantiates a new GenerateJitconfigRequestBuilder and sets the default values.
func NewItemActionsRunnersGenerateJitconfigRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemActionsRunnersGenerateJitconfigRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemActionsRunnersGenerateJitconfigRequestBuilderInternal(urlParams, requestAdapter)
}
// Post generates a configuration that can be passed to the runner application at startup.You must authenticate using an access token with the `admin:org` scope to use this endpoint.If the repository is private, you must use an access token with the `repo` scope.GitHub Apps must have the `administration` permission for repositories and the `organization_self_hosted_runners` permission for organizations.Authenticated users must have admin access to repositories or organizations, or the `manage_runners:enterprise` scope for enterprises, to use these endpoints.
// Deprecated: This method is obsolete. Use PostAsGenerateJitconfigPostResponse instead.
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/actions/self-hosted-runners#create-configuration-for-a-just-in-time-runner-for-an-organization
func (m *ItemActionsRunnersGenerateJitconfigRequestBuilder) Post(ctx context.Context, body ItemActionsRunnersGenerateJitconfigPostRequestBodyable, requestConfiguration *ItemActionsRunnersGenerateJitconfigRequestBuilderPostRequestConfiguration)(ItemActionsRunnersGenerateJitconfigResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "404": i33bcfa37677dfd8b71eaf22ad3bfd140ba1702e7a8418558499c22418bf2fef2.CreateBasicErrorFromDiscriminatorValue,
        "422": i33bcfa37677dfd8b71eaf22ad3bfd140ba1702e7a8418558499c22418bf2fef2.CreateValidationErrorSimpleFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateItemActionsRunnersGenerateJitconfigResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ItemActionsRunnersGenerateJitconfigResponseable), nil
}
// PostAsGenerateJitconfigPostResponse generates a configuration that can be passed to the runner application at startup.You must authenticate using an access token with the `admin:org` scope to use this endpoint.If the repository is private, you must use an access token with the `repo` scope.GitHub Apps must have the `administration` permission for repositories and the `organization_self_hosted_runners` permission for organizations.Authenticated users must have admin access to repositories or organizations, or the `manage_runners:enterprise` scope for enterprises, to use these endpoints.
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/actions/self-hosted-runners#create-configuration-for-a-just-in-time-runner-for-an-organization
func (m *ItemActionsRunnersGenerateJitconfigRequestBuilder) PostAsGenerateJitconfigPostResponse(ctx context.Context, body ItemActionsRunnersGenerateJitconfigPostRequestBodyable, requestConfiguration *ItemActionsRunnersGenerateJitconfigRequestBuilderPostRequestConfiguration)(ItemActionsRunnersGenerateJitconfigPostResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "404": i33bcfa37677dfd8b71eaf22ad3bfd140ba1702e7a8418558499c22418bf2fef2.CreateBasicErrorFromDiscriminatorValue,
        "422": i33bcfa37677dfd8b71eaf22ad3bfd140ba1702e7a8418558499c22418bf2fef2.CreateValidationErrorSimpleFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateItemActionsRunnersGenerateJitconfigPostResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ItemActionsRunnersGenerateJitconfigPostResponseable), nil
}
// ToPostRequestInformation generates a configuration that can be passed to the runner application at startup.You must authenticate using an access token with the `admin:org` scope to use this endpoint.If the repository is private, you must use an access token with the `repo` scope.GitHub Apps must have the `administration` permission for repositories and the `organization_self_hosted_runners` permission for organizations.Authenticated users must have admin access to repositories or organizations, or the `manage_runners:enterprise` scope for enterprises, to use these endpoints.
func (m *ItemActionsRunnersGenerateJitconfigRequestBuilder) ToPostRequestInformation(ctx context.Context, body ItemActionsRunnersGenerateJitconfigPostRequestBodyable, requestConfiguration *ItemActionsRunnersGenerateJitconfigRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *ItemActionsRunnersGenerateJitconfigRequestBuilder) WithUrl(rawUrl string)(*ItemActionsRunnersGenerateJitconfigRequestBuilder) {
    return NewItemActionsRunnersGenerateJitconfigRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
