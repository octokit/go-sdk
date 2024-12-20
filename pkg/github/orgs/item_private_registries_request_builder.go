package orgs

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6 "github.com/octokit/go-sdk/pkg/github/models"
)

// ItemPrivateRegistriesRequestBuilder builds and executes requests for operations under \orgs\{org}\private-registries
type ItemPrivateRegistriesRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemPrivateRegistriesRequestBuilderGetQueryParameters > [!NOTE]> This endpoint is in public preview and is subject to change.Lists all private registry configurations available at the organization-level without revealing their encryptedvalues.OAuth app tokens and personal access tokens (classic) need the `admin:org` scope to use this endpoint.
type ItemPrivateRegistriesRequestBuilderGetQueryParameters struct {
    // The page number of the results to fetch. For more information, see "[Using pagination in the REST API](https://docs.github.com/rest/using-the-rest-api/using-pagination-in-the-rest-api)."
    Page *int32 `uriparametername:"page"`
    // The number of results per page (max 100). For more information, see "[Using pagination in the REST API](https://docs.github.com/rest/using-the-rest-api/using-pagination-in-the-rest-api)."
    Per_page *int32 `uriparametername:"per_page"`
}
// BySecret_name gets an item from the github.com/octokit/go-sdk/pkg/github.orgs.item.privateRegistries.item collection
// returns a *ItemPrivateRegistriesWithSecret_nameItemRequestBuilder when successful
func (m *ItemPrivateRegistriesRequestBuilder) BySecret_name(secret_name string)(*ItemPrivateRegistriesWithSecret_nameItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if secret_name != "" {
        urlTplParams["secret_name"] = secret_name
    }
    return NewItemPrivateRegistriesWithSecret_nameItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewItemPrivateRegistriesRequestBuilderInternal instantiates a new ItemPrivateRegistriesRequestBuilder and sets the default values.
func NewItemPrivateRegistriesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemPrivateRegistriesRequestBuilder) {
    m := &ItemPrivateRegistriesRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/orgs/{org}/private-registries{?page*,per_page*}", pathParameters),
    }
    return m
}
// NewItemPrivateRegistriesRequestBuilder instantiates a new ItemPrivateRegistriesRequestBuilder and sets the default values.
func NewItemPrivateRegistriesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemPrivateRegistriesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemPrivateRegistriesRequestBuilderInternal(urlParams, requestAdapter)
}
// Get > [!NOTE]> This endpoint is in public preview and is subject to change.Lists all private registry configurations available at the organization-level without revealing their encryptedvalues.OAuth app tokens and personal access tokens (classic) need the `admin:org` scope to use this endpoint.
// returns a ItemPrivateRegistriesGetResponseable when successful
// returns a BasicError error when the service returns a 400 status code
// returns a BasicError error when the service returns a 404 status code
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/private-registries/organization-configurations#list-private-registries-for-an-organization
func (m *ItemPrivateRegistriesRequestBuilder) Get(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[ItemPrivateRegistriesRequestBuilderGetQueryParameters])(ItemPrivateRegistriesGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "400": i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CreateBasicErrorFromDiscriminatorValue,
        "404": i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CreateBasicErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateItemPrivateRegistriesGetResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ItemPrivateRegistriesGetResponseable), nil
}
// Post > [!NOTE]> This endpoint is in public preview and is subject to change.Creates a private registry configuration with an encrypted value for an organization. Encrypt your secret using [LibSodium](https://libsodium.gitbook.io/doc/bindings_for_other_languages). For more information, see "[Encrypting secrets for the REST API](https://docs.github.com/rest/guides/encrypting-secrets-for-the-rest-api)."OAuth app tokens and personal access tokens (classic) need the `admin:org` scope to use this endpoint.
// returns a OrgPrivateRegistryConfigurationWithSelectedRepositoriesable when successful
// returns a BasicError error when the service returns a 404 status code
// returns a ValidationError error when the service returns a 422 status code
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/private-registries/organization-configurations#create-a-private-registry-for-an-organization
func (m *ItemPrivateRegistriesRequestBuilder) Post(ctx context.Context, body ItemPrivateRegistriesPostRequestBodyable, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.OrgPrivateRegistryConfigurationWithSelectedRepositoriesable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "404": i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CreateBasicErrorFromDiscriminatorValue,
        "422": i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CreateValidationErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CreateOrgPrivateRegistryConfigurationWithSelectedRepositoriesFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.OrgPrivateRegistryConfigurationWithSelectedRepositoriesable), nil
}
// PublicKey the publicKey property
// returns a *ItemPrivateRegistriesPublicKeyRequestBuilder when successful
func (m *ItemPrivateRegistriesRequestBuilder) PublicKey()(*ItemPrivateRegistriesPublicKeyRequestBuilder) {
    return NewItemPrivateRegistriesPublicKeyRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToGetRequestInformation > [!NOTE]> This endpoint is in public preview and is subject to change.Lists all private registry configurations available at the organization-level without revealing their encryptedvalues.OAuth app tokens and personal access tokens (classic) need the `admin:org` scope to use this endpoint.
// returns a *RequestInformation when successful
func (m *ItemPrivateRegistriesRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[ItemPrivateRegistriesRequestBuilderGetQueryParameters])(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ConfigureRequestInformation(requestInfo, requestConfiguration)
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToPostRequestInformation > [!NOTE]> This endpoint is in public preview and is subject to change.Creates a private registry configuration with an encrypted value for an organization. Encrypt your secret using [LibSodium](https://libsodium.gitbook.io/doc/bindings_for_other_languages). For more information, see "[Encrypting secrets for the REST API](https://docs.github.com/rest/guides/encrypting-secrets-for-the-rest-api)."OAuth app tokens and personal access tokens (classic) need the `admin:org` scope to use this endpoint.
// returns a *RequestInformation when successful
func (m *ItemPrivateRegistriesRequestBuilder) ToPostRequestInformation(ctx context.Context, body ItemPrivateRegistriesPostRequestBodyable, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemPrivateRegistriesRequestBuilder when successful
func (m *ItemPrivateRegistriesRequestBuilder) WithUrl(rawUrl string)(*ItemPrivateRegistriesRequestBuilder) {
    return NewItemPrivateRegistriesRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
