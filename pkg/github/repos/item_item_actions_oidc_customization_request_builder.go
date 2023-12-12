package repos

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemItemActionsOidcCustomizationRequestBuilder builds and executes requests for operations under \repos\{repos-id}\{Owner-id}\actions\oidc\customization
type ItemItemActionsOidcCustomizationRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// NewItemItemActionsOidcCustomizationRequestBuilderInternal instantiates a new CustomizationRequestBuilder and sets the default values.
func NewItemItemActionsOidcCustomizationRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemActionsOidcCustomizationRequestBuilder) {
    m := &ItemItemActionsOidcCustomizationRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/repos/{repos%2Did}/{Owner%2Did}/actions/oidc/customization", pathParameters),
    }
    return m
}
// NewItemItemActionsOidcCustomizationRequestBuilder instantiates a new CustomizationRequestBuilder and sets the default values.
func NewItemItemActionsOidcCustomizationRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemActionsOidcCustomizationRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemItemActionsOidcCustomizationRequestBuilderInternal(urlParams, requestAdapter)
}
// Sub the sub property
func (m *ItemItemActionsOidcCustomizationRequestBuilder) Sub()(*ItemItemActionsOidcCustomizationSubRequestBuilder) {
    return NewItemItemActionsOidcCustomizationSubRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
