package orgs

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6 "github.com/octokit/go-sdk/pkg/github/models"
)

// ItemCodeSecurityConfigurationsDetachRequestBuilder builds and executes requests for operations under \orgs\{org}\code-security\configurations\detach
type ItemCodeSecurityConfigurationsDetachRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// NewItemCodeSecurityConfigurationsDetachRequestBuilderInternal instantiates a new ItemCodeSecurityConfigurationsDetachRequestBuilder and sets the default values.
func NewItemCodeSecurityConfigurationsDetachRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemCodeSecurityConfigurationsDetachRequestBuilder) {
    m := &ItemCodeSecurityConfigurationsDetachRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/orgs/{org}/code-security/configurations/detach", pathParameters),
    }
    return m
}
// NewItemCodeSecurityConfigurationsDetachRequestBuilder instantiates a new ItemCodeSecurityConfigurationsDetachRequestBuilder and sets the default values.
func NewItemCodeSecurityConfigurationsDetachRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemCodeSecurityConfigurationsDetachRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemCodeSecurityConfigurationsDetachRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete detach code security configuration(s) from a set of repositories.Repositories will retain their settings but will no longer be associated with the configuration.The authenticated user must be an administrator or security manager for the organization to use this endpoint.OAuth app tokens and personal access tokens (classic) need the `write:org` scope to use this endpoint.
// returns a BasicError error when the service returns a 400 status code
// returns a BasicError error when the service returns a 403 status code
// returns a BasicError error when the service returns a 404 status code
// returns a BasicError error when the service returns a 409 status code
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/code-security/configurations#detach-configurations-from-repositories
func (m *ItemCodeSecurityConfigurationsDetachRequestBuilder) Delete(ctx context.Context, body ItemCodeSecurityConfigurationsDetachDeleteRequestBodyable, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(error) {
    requestInfo, err := m.ToDeleteRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "400": i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CreateBasicErrorFromDiscriminatorValue,
        "403": i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CreateBasicErrorFromDiscriminatorValue,
        "404": i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CreateBasicErrorFromDiscriminatorValue,
        "409": i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CreateBasicErrorFromDiscriminatorValue,
    }
    err = m.BaseRequestBuilder.RequestAdapter.SendNoContent(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// ToDeleteRequestInformation detach code security configuration(s) from a set of repositories.Repositories will retain their settings but will no longer be associated with the configuration.The authenticated user must be an administrator or security manager for the organization to use this endpoint.OAuth app tokens and personal access tokens (classic) need the `write:org` scope to use this endpoint.
// returns a *RequestInformation when successful
func (m *ItemCodeSecurityConfigurationsDetachRequestBuilder) ToDeleteRequestInformation(ctx context.Context, body ItemCodeSecurityConfigurationsDetachDeleteRequestBodyable, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ConfigureRequestInformation(requestInfo, requestConfiguration)
    requestInfo.Headers.TryAdd("Accept", "application/json, application/scim+json")
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemCodeSecurityConfigurationsDetachRequestBuilder when successful
func (m *ItemCodeSecurityConfigurationsDetachRequestBuilder) WithUrl(rawUrl string)(*ItemCodeSecurityConfigurationsDetachRequestBuilder) {
    return NewItemCodeSecurityConfigurationsDetachRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
