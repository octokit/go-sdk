package orgs

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemActionsRequestBuilder builds and executes requests for operations under \orgs\{org}\actions
type ItemActionsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// Cache the cache property
func (m *ItemActionsRequestBuilder) Cache()(*ItemActionsCacheRequestBuilder) {
    return NewItemActionsCacheRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewItemActionsRequestBuilderInternal instantiates a new ActionsRequestBuilder and sets the default values.
func NewItemActionsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemActionsRequestBuilder) {
    m := &ItemActionsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/orgs/{org}/actions", pathParameters),
    }
    return m
}
// NewItemActionsRequestBuilder instantiates a new ActionsRequestBuilder and sets the default values.
func NewItemActionsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemActionsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemActionsRequestBuilderInternal(urlParams, requestAdapter)
}
// Oidc the oidc property
func (m *ItemActionsRequestBuilder) Oidc()(*ItemActionsOidcRequestBuilder) {
    return NewItemActionsOidcRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Permissions the permissions property
func (m *ItemActionsRequestBuilder) Permissions()(*ItemActionsPermissionsRequestBuilder) {
    return NewItemActionsPermissionsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Runners the runners property
func (m *ItemActionsRequestBuilder) Runners()(*ItemActionsRunnersRequestBuilder) {
    return NewItemActionsRunnersRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Secrets the secrets property
func (m *ItemActionsRequestBuilder) Secrets()(*ItemActionsSecretsRequestBuilder) {
    return NewItemActionsSecretsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Variables the variables property
func (m *ItemActionsRequestBuilder) Variables()(*ItemActionsVariablesRequestBuilder) {
    return NewItemActionsVariablesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
