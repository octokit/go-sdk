package orgs

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6 "github.com/octokit/go-sdk/pkg/github/models"
)

// ItemPrivateRegistriesWithSecret_nameItemRequestBuilder builds and executes requests for operations under \orgs\{org}\private-registries\{secret_name}
type ItemPrivateRegistriesWithSecret_nameItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// NewItemPrivateRegistriesWithSecret_nameItemRequestBuilderInternal instantiates a new ItemPrivateRegistriesWithSecret_nameItemRequestBuilder and sets the default values.
func NewItemPrivateRegistriesWithSecret_nameItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemPrivateRegistriesWithSecret_nameItemRequestBuilder) {
    m := &ItemPrivateRegistriesWithSecret_nameItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/orgs/{org}/private-registries/{secret_name}", pathParameters),
    }
    return m
}
// NewItemPrivateRegistriesWithSecret_nameItemRequestBuilder instantiates a new ItemPrivateRegistriesWithSecret_nameItemRequestBuilder and sets the default values.
func NewItemPrivateRegistriesWithSecret_nameItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemPrivateRegistriesWithSecret_nameItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemPrivateRegistriesWithSecret_nameItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete > [!NOTE]> This endpoint is in public preview and is subject to change.Delete a private registry configuration at the organization-level.OAuth app tokens and personal access tokens (classic) need the `admin:org` scope to use this endpoint.
// returns a BasicError error when the service returns a 400 status code
// returns a BasicError error when the service returns a 404 status code
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/private-registries/organization-configurations#delete-a-private-registry-for-an-organization
func (m *ItemPrivateRegistriesWithSecret_nameItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(error) {
    requestInfo, err := m.ToDeleteRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "400": i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CreateBasicErrorFromDiscriminatorValue,
        "404": i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CreateBasicErrorFromDiscriminatorValue,
    }
    err = m.BaseRequestBuilder.RequestAdapter.SendNoContent(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Get > [!NOTE]> This endpoint is in public preview and is subject to change.Get the configuration of a single private registry defined for an organization, omitting its encrypted value.OAuth app tokens and personal access tokens (classic) need the `admin:org` scope to use this endpoint.
// returns a OrgPrivateRegistryConfigurationable when successful
// returns a BasicError error when the service returns a 404 status code
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/private-registries/organization-configurations#get-a-private-registry-for-an-organization
func (m *ItemPrivateRegistriesWithSecret_nameItemRequestBuilder) Get(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.OrgPrivateRegistryConfigurationable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "404": i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CreateBasicErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CreateOrgPrivateRegistryConfigurationFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.OrgPrivateRegistryConfigurationable), nil
}
// Patch > [!NOTE]> This endpoint is in public preview and is subject to change.Updates a private registry configuration with an encrypted value for an organization. Encrypt your secret using [LibSodium](https://libsodium.gitbook.io/doc/bindings_for_other_languages). For more information, see "[Encrypting secrets for the REST API](https://docs.github.com/rest/guides/encrypting-secrets-for-the-rest-api)."OAuth app tokens and personal access tokens (classic) need the `admin:org` scope to use this endpoint.
// returns a BasicError error when the service returns a 404 status code
// returns a ValidationError error when the service returns a 422 status code
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/private-registries/organization-configurations#update-a-private-registry-for-an-organization
func (m *ItemPrivateRegistriesWithSecret_nameItemRequestBuilder) Patch(ctx context.Context, body ItemPrivateRegistriesItemWithSecret_namePatchRequestBodyable, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "404": i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CreateBasicErrorFromDiscriminatorValue,
        "422": i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CreateValidationErrorFromDiscriminatorValue,
    }
    err = m.BaseRequestBuilder.RequestAdapter.SendNoContent(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// ToDeleteRequestInformation > [!NOTE]> This endpoint is in public preview and is subject to change.Delete a private registry configuration at the organization-level.OAuth app tokens and personal access tokens (classic) need the `admin:org` scope to use this endpoint.
// returns a *RequestInformation when successful
func (m *ItemPrivateRegistriesWithSecret_nameItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ConfigureRequestInformation(requestInfo, requestConfiguration)
    requestInfo.Headers.TryAdd("Accept", "application/json, application/scim+json")
    return requestInfo, nil
}
// ToGetRequestInformation > [!NOTE]> This endpoint is in public preview and is subject to change.Get the configuration of a single private registry defined for an organization, omitting its encrypted value.OAuth app tokens and personal access tokens (classic) need the `admin:org` scope to use this endpoint.
// returns a *RequestInformation when successful
func (m *ItemPrivateRegistriesWithSecret_nameItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ConfigureRequestInformation(requestInfo, requestConfiguration)
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToPatchRequestInformation > [!NOTE]> This endpoint is in public preview and is subject to change.Updates a private registry configuration with an encrypted value for an organization. Encrypt your secret using [LibSodium](https://libsodium.gitbook.io/doc/bindings_for_other_languages). For more information, see "[Encrypting secrets for the REST API](https://docs.github.com/rest/guides/encrypting-secrets-for-the-rest-api)."OAuth app tokens and personal access tokens (classic) need the `admin:org` scope to use this endpoint.
// returns a *RequestInformation when successful
func (m *ItemPrivateRegistriesWithSecret_nameItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ItemPrivateRegistriesItemWithSecret_namePatchRequestBodyable, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ConfigureRequestInformation(requestInfo, requestConfiguration)
    requestInfo.Headers.TryAdd("Accept", "application/json")
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemPrivateRegistriesWithSecret_nameItemRequestBuilder when successful
func (m *ItemPrivateRegistriesWithSecret_nameItemRequestBuilder) WithUrl(rawUrl string)(*ItemPrivateRegistriesWithSecret_nameItemRequestBuilder) {
    return NewItemPrivateRegistriesWithSecret_nameItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
