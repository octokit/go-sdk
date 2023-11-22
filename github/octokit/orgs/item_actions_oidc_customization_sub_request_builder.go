package orgs

import (
    "context"
    i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e "github.com/octokit/go-sdk/github/octokit/models"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemActionsOidcCustomizationSubRequestBuilder builds and executes requests for operations under \orgs\{org}\actions\oidc\customization\sub
type ItemActionsOidcCustomizationSubRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemActionsOidcCustomizationSubRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemActionsOidcCustomizationSubRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ItemActionsOidcCustomizationSubRequestBuilderPutRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemActionsOidcCustomizationSubRequestBuilderPutRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemActionsOidcCustomizationSubRequestBuilderInternal instantiates a new SubRequestBuilder and sets the default values.
func NewItemActionsOidcCustomizationSubRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemActionsOidcCustomizationSubRequestBuilder) {
    m := &ItemActionsOidcCustomizationSubRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/orgs/{org}/actions/oidc/customization/sub", pathParameters),
    }
    return m
}
// NewItemActionsOidcCustomizationSubRequestBuilder instantiates a new SubRequestBuilder and sets the default values.
func NewItemActionsOidcCustomizationSubRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemActionsOidcCustomizationSubRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemActionsOidcCustomizationSubRequestBuilderInternal(urlParams, requestAdapter)
}
// Get gets the customization template for an OpenID Connect (OIDC) subject claim.You must authenticate using an access token with the `read:org` scope to use this endpoint.GitHub Apps must have the `organization_administration:write` permission to use this endpoint.
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/actions/oidc#get-the-customization-template-for-an-oidc-subject-claim-for-an-organization
func (m *ItemActionsOidcCustomizationSubRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemActionsOidcCustomizationSubRequestBuilderGetRequestConfiguration)(i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.OidcCustomSubable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.CreateOidcCustomSubFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.OidcCustomSubable), nil
}
// Put creates or updates the customization template for an OpenID Connect (OIDC) subject claim.You must authenticate using an access token with the `write:org` scope to use this endpoint.GitHub Apps must have the `admin:org` permission to use this endpoint.
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/actions/oidc#set-the-customization-template-for-an-oidc-subject-claim-for-an-organization
func (m *ItemActionsOidcCustomizationSubRequestBuilder) Put(ctx context.Context, body i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.OidcCustomSubable, requestConfiguration *ItemActionsOidcCustomizationSubRequestBuilderPutRequestConfiguration)(i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.EmptyObjectable, error) {
    requestInfo, err := m.ToPutRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "403": i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.CreateBasicErrorFromDiscriminatorValue,
        "404": i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.CreateBasicErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.CreateEmptyObjectFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.EmptyObjectable), nil
}
// ToGetRequestInformation gets the customization template for an OpenID Connect (OIDC) subject claim.You must authenticate using an access token with the `read:org` scope to use this endpoint.GitHub Apps must have the `organization_administration:write` permission to use this endpoint.
func (m *ItemActionsOidcCustomizationSubRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemActionsOidcCustomizationSubRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToPutRequestInformation creates or updates the customization template for an OpenID Connect (OIDC) subject claim.You must authenticate using an access token with the `write:org` scope to use this endpoint.GitHub Apps must have the `admin:org` permission to use this endpoint.
func (m *ItemActionsOidcCustomizationSubRequestBuilder) ToPutRequestInformation(ctx context.Context, body i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.OidcCustomSubable, requestConfiguration *ItemActionsOidcCustomizationSubRequestBuilderPutRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PUT
    requestInfo.Headers.TryAdd("Accept", "application/json")
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
func (m *ItemActionsOidcCustomizationSubRequestBuilder) WithUrl(rawUrl string)(*ItemActionsOidcCustomizationSubRequestBuilder) {
    return NewItemActionsOidcCustomizationSubRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
