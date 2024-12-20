package orgs

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6 "github.com/octokit/go-sdk/pkg/github/models"
)

// ItemPrivateRegistriesPublicKeyRequestBuilder builds and executes requests for operations under \orgs\{org}\private-registries\public-key
type ItemPrivateRegistriesPublicKeyRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// NewItemPrivateRegistriesPublicKeyRequestBuilderInternal instantiates a new ItemPrivateRegistriesPublicKeyRequestBuilder and sets the default values.
func NewItemPrivateRegistriesPublicKeyRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemPrivateRegistriesPublicKeyRequestBuilder) {
    m := &ItemPrivateRegistriesPublicKeyRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/orgs/{org}/private-registries/public-key", pathParameters),
    }
    return m
}
// NewItemPrivateRegistriesPublicKeyRequestBuilder instantiates a new ItemPrivateRegistriesPublicKeyRequestBuilder and sets the default values.
func NewItemPrivateRegistriesPublicKeyRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemPrivateRegistriesPublicKeyRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemPrivateRegistriesPublicKeyRequestBuilderInternal(urlParams, requestAdapter)
}
// Get > [!NOTE]> This endpoint is in public preview and is subject to change.Gets the org public key, which is needed to encrypt private registry secrets. You need to encrypt a secret before you can create or update secrets.OAuth tokens and personal access tokens (classic) need the `admin:org` scope to use this endpoint.
// returns a ItemPrivateRegistriesPublicKeyGetResponseable when successful
// returns a BasicError error when the service returns a 404 status code
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/private-registries/organization-configurations#get-private-registries-public-key-for-an-organization
func (m *ItemPrivateRegistriesPublicKeyRequestBuilder) Get(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(ItemPrivateRegistriesPublicKeyGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "404": i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CreateBasicErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateItemPrivateRegistriesPublicKeyGetResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ItemPrivateRegistriesPublicKeyGetResponseable), nil
}
// ToGetRequestInformation > [!NOTE]> This endpoint is in public preview and is subject to change.Gets the org public key, which is needed to encrypt private registry secrets. You need to encrypt a secret before you can create or update secrets.OAuth tokens and personal access tokens (classic) need the `admin:org` scope to use this endpoint.
// returns a *RequestInformation when successful
func (m *ItemPrivateRegistriesPublicKeyRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ConfigureRequestInformation(requestInfo, requestConfiguration)
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemPrivateRegistriesPublicKeyRequestBuilder when successful
func (m *ItemPrivateRegistriesPublicKeyRequestBuilder) WithUrl(rawUrl string)(*ItemPrivateRegistriesPublicKeyRequestBuilder) {
    return NewItemPrivateRegistriesPublicKeyRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
