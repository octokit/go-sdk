package enterprises

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6 "github.com/octokit/go-sdk/pkg/github/models"
)

// ItemCodeSecurityConfigurationsDefaultsRequestBuilder builds and executes requests for operations under \enterprises\{enterprise}\code-security\configurations\defaults
type ItemCodeSecurityConfigurationsDefaultsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// NewItemCodeSecurityConfigurationsDefaultsRequestBuilderInternal instantiates a new ItemCodeSecurityConfigurationsDefaultsRequestBuilder and sets the default values.
func NewItemCodeSecurityConfigurationsDefaultsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemCodeSecurityConfigurationsDefaultsRequestBuilder) {
    m := &ItemCodeSecurityConfigurationsDefaultsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/enterprises/{enterprise}/code-security/configurations/defaults", pathParameters),
    }
    return m
}
// NewItemCodeSecurityConfigurationsDefaultsRequestBuilder instantiates a new ItemCodeSecurityConfigurationsDefaultsRequestBuilder and sets the default values.
func NewItemCodeSecurityConfigurationsDefaultsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemCodeSecurityConfigurationsDefaultsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemCodeSecurityConfigurationsDefaultsRequestBuilderInternal(urlParams, requestAdapter)
}
// Get lists the default code security configurations for an enterprise.The authenticated user must be an administrator of the enterprise in order to use this endpoint.OAuth app tokens and personal access tokens (classic) need the `read:enterprise` scope to use this endpoint.
// returns a []CodeSecurityDefaultConfigurationsable when successful
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/code-security/configurations#get-default-code-security-configurations-for-an-enterprise
func (m *ItemCodeSecurityConfigurationsDefaultsRequestBuilder) Get(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])([]i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CodeSecurityDefaultConfigurationsable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.SendCollection(ctx, requestInfo, i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CreateCodeSecurityDefaultConfigurationsFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    val := make([]i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CodeSecurityDefaultConfigurationsable, len(res))
    for i, v := range res {
        if v != nil {
            val[i] = v.(i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CodeSecurityDefaultConfigurationsable)
        }
    }
    return val, nil
}
// ToGetRequestInformation lists the default code security configurations for an enterprise.The authenticated user must be an administrator of the enterprise in order to use this endpoint.OAuth app tokens and personal access tokens (classic) need the `read:enterprise` scope to use this endpoint.
// returns a *RequestInformation when successful
func (m *ItemCodeSecurityConfigurationsDefaultsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ConfigureRequestInformation(requestInfo, requestConfiguration)
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemCodeSecurityConfigurationsDefaultsRequestBuilder when successful
func (m *ItemCodeSecurityConfigurationsDefaultsRequestBuilder) WithUrl(rawUrl string)(*ItemCodeSecurityConfigurationsDefaultsRequestBuilder) {
    return NewItemCodeSecurityConfigurationsDefaultsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
