package repos

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemItemActionsRequestBuilder builds and executes requests for operations under \repos\{owner}\{repo}\actions
type ItemItemActionsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// Artifacts the artifacts property
func (m *ItemItemActionsRequestBuilder) Artifacts()(*ItemItemActionsArtifactsRequestBuilder) {
    return NewItemItemActionsArtifactsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Cache the cache property
func (m *ItemItemActionsRequestBuilder) Cache()(*ItemItemActionsCacheRequestBuilder) {
    return NewItemItemActionsCacheRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Caches the caches property
func (m *ItemItemActionsRequestBuilder) Caches()(*ItemItemActionsCachesRequestBuilder) {
    return NewItemItemActionsCachesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewItemItemActionsRequestBuilderInternal instantiates a new ActionsRequestBuilder and sets the default values.
func NewItemItemActionsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemActionsRequestBuilder) {
    m := &ItemItemActionsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/repos/{owner}/{repo}/actions", pathParameters),
    }
    return m
}
// NewItemItemActionsRequestBuilder instantiates a new ActionsRequestBuilder and sets the default values.
func NewItemItemActionsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemActionsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemItemActionsRequestBuilderInternal(urlParams, requestAdapter)
}
// Jobs the jobs property
func (m *ItemItemActionsRequestBuilder) Jobs()(*ItemItemActionsJobsRequestBuilder) {
    return NewItemItemActionsJobsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Oidc the oidc property
func (m *ItemItemActionsRequestBuilder) Oidc()(*ItemItemActionsOidcRequestBuilder) {
    return NewItemItemActionsOidcRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// OrganizationSecrets the organizationSecrets property
func (m *ItemItemActionsRequestBuilder) OrganizationSecrets()(*ItemItemActionsOrganizationSecretsRequestBuilder) {
    return NewItemItemActionsOrganizationSecretsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// OrganizationVariables the organizationVariables property
func (m *ItemItemActionsRequestBuilder) OrganizationVariables()(*ItemItemActionsOrganizationVariablesRequestBuilder) {
    return NewItemItemActionsOrganizationVariablesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Permissions the permissions property
func (m *ItemItemActionsRequestBuilder) Permissions()(*ItemItemActionsPermissionsRequestBuilder) {
    return NewItemItemActionsPermissionsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Runners the runners property
func (m *ItemItemActionsRequestBuilder) Runners()(*ItemItemActionsRunnersRequestBuilder) {
    return NewItemItemActionsRunnersRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Runs the runs property
func (m *ItemItemActionsRequestBuilder) Runs()(*ItemItemActionsRunsRequestBuilder) {
    return NewItemItemActionsRunsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Secrets the secrets property
func (m *ItemItemActionsRequestBuilder) Secrets()(*ItemItemActionsSecretsRequestBuilder) {
    return NewItemItemActionsSecretsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Variables the variables property
func (m *ItemItemActionsRequestBuilder) Variables()(*ItemItemActionsVariablesRequestBuilder) {
    return NewItemItemActionsVariablesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Workflows the workflows property
func (m *ItemItemActionsRequestBuilder) Workflows()(*ItemItemActionsWorkflowsRequestBuilder) {
    return NewItemItemActionsWorkflowsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
