package enterprises

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6 "github.com/octokit/go-sdk/pkg/github/models"
)

// ItemCodeSecurityConfigurationsItemAttachRequestBuilder builds and executes requests for operations under \enterprises\{enterprise}\code-security\configurations\{configuration_id}\attach
type ItemCodeSecurityConfigurationsItemAttachRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// NewItemCodeSecurityConfigurationsItemAttachRequestBuilderInternal instantiates a new ItemCodeSecurityConfigurationsItemAttachRequestBuilder and sets the default values.
func NewItemCodeSecurityConfigurationsItemAttachRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemCodeSecurityConfigurationsItemAttachRequestBuilder) {
    m := &ItemCodeSecurityConfigurationsItemAttachRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/enterprises/{enterprise}/code-security/configurations/{configuration_id}/attach", pathParameters),
    }
    return m
}
// NewItemCodeSecurityConfigurationsItemAttachRequestBuilder instantiates a new ItemCodeSecurityConfigurationsItemAttachRequestBuilder and sets the default values.
func NewItemCodeSecurityConfigurationsItemAttachRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemCodeSecurityConfigurationsItemAttachRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemCodeSecurityConfigurationsItemAttachRequestBuilderInternal(urlParams, requestAdapter)
}
// Post attaches an enterprise code security configuration to repositories. If the repositories specified are already attached to a configuration, they will be re-attached to the provided configuration.If insufficient GHAS licenses are available to attach the configuration to a repository, only free features will be enabled.The authenticated user must be an administrator for the enterprise to use this endpoint.OAuth app tokens and personal access tokens (classic) need the `admin:enterprise` scope to use this endpoint.
// returns a ItemCodeSecurityConfigurationsItemAttachPostResponseable when successful
// returns a BasicError error when the service returns a 403 status code
// returns a BasicError error when the service returns a 404 status code
// returns a BasicError error when the service returns a 409 status code
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/code-security/configurations#attach-an-enterprise-configuration-to-repositories
func (m *ItemCodeSecurityConfigurationsItemAttachRequestBuilder) Post(ctx context.Context, body ItemCodeSecurityConfigurationsItemAttachPostRequestBodyable, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(ItemCodeSecurityConfigurationsItemAttachPostResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "403": i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CreateBasicErrorFromDiscriminatorValue,
        "404": i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CreateBasicErrorFromDiscriminatorValue,
        "409": i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CreateBasicErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateItemCodeSecurityConfigurationsItemAttachPostResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ItemCodeSecurityConfigurationsItemAttachPostResponseable), nil
}
// ToPostRequestInformation attaches an enterprise code security configuration to repositories. If the repositories specified are already attached to a configuration, they will be re-attached to the provided configuration.If insufficient GHAS licenses are available to attach the configuration to a repository, only free features will be enabled.The authenticated user must be an administrator for the enterprise to use this endpoint.OAuth app tokens and personal access tokens (classic) need the `admin:enterprise` scope to use this endpoint.
// returns a *RequestInformation when successful
func (m *ItemCodeSecurityConfigurationsItemAttachRequestBuilder) ToPostRequestInformation(ctx context.Context, body ItemCodeSecurityConfigurationsItemAttachPostRequestBodyable, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ConfigureRequestInformation(requestInfo, requestConfiguration)
    requestInfo.Headers.TryAdd("Accept", "application/json")
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemCodeSecurityConfigurationsItemAttachRequestBuilder when successful
func (m *ItemCodeSecurityConfigurationsItemAttachRequestBuilder) WithUrl(rawUrl string)(*ItemCodeSecurityConfigurationsItemAttachRequestBuilder) {
    return NewItemCodeSecurityConfigurationsItemAttachRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
