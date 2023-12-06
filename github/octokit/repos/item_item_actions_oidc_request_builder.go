package repos

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemItemActionsOidcRequestBuilder builds and executes requests for operations under \repos\{repos-id}\{Owner-id}\actions\oidc
type ItemItemActionsOidcRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// NewItemItemActionsOidcRequestBuilderInternal instantiates a new OidcRequestBuilder and sets the default values.
func NewItemItemActionsOidcRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemActionsOidcRequestBuilder) {
    m := &ItemItemActionsOidcRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/repos/{repos%2Did}/{Owner%2Did}/actions/oidc", pathParameters),
    }
    return m
}
// NewItemItemActionsOidcRequestBuilder instantiates a new OidcRequestBuilder and sets the default values.
func NewItemItemActionsOidcRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemActionsOidcRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemItemActionsOidcRequestBuilderInternal(urlParams, requestAdapter)
}
// Customization the customization property
func (m *ItemItemActionsOidcRequestBuilder) Customization()(*ItemItemActionsOidcCustomizationRequestBuilder) {
    return NewItemItemActionsOidcCustomizationRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
