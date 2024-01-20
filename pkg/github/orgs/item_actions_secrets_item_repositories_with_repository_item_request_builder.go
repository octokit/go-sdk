package orgs

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemActionsSecretsItemRepositoriesWithRepository_ItemRequestBuilder builds and executes requests for operations under \orgs\{org}\actions\secrets\{secret_name}\repositories\{repository_id}
type ItemActionsSecretsItemRepositoriesWithRepository_ItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// NewItemActionsSecretsItemRepositoriesWithRepository_ItemRequestBuilderInternal instantiates a new WithRepository_ItemRequestBuilder and sets the default values.
func NewItemActionsSecretsItemRepositoriesWithRepository_ItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemActionsSecretsItemRepositoriesWithRepository_ItemRequestBuilder) {
    m := &ItemActionsSecretsItemRepositoriesWithRepository_ItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/orgs/{org}/actions/secrets/{secret_name}/repositories/{repository_id}", pathParameters),
    }
    return m
}
// NewItemActionsSecretsItemRepositoriesWithRepository_ItemRequestBuilder instantiates a new WithRepository_ItemRequestBuilder and sets the default values.
func NewItemActionsSecretsItemRepositoriesWithRepository_ItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemActionsSecretsItemRepositoriesWithRepository_ItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemActionsSecretsItemRepositoriesWithRepository_ItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete removes a repository from an organization secret when the `visibility`for repository access is set to `selected`. The visibility is set when you [Createor update an organization secret](https://docs.github.com/rest/actions/secrets#create-or-update-an-organization-secret).You must authenticate using an access token with the `admin:org` scope to use this endpoint.If the repository is private, you must use an access token with the `repo` scope.GitHub Apps must have the `secrets` organization permission to use this endpoint.Authenticated users must have collaborator access to a repository to create, update, or read secrets.
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/actions/secrets#remove-selected-repository-from-an-organization-secret
func (m *ItemActionsSecretsItemRepositoriesWithRepository_ItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(error) {
    requestInfo, err := m.ToDeleteRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    err = m.BaseRequestBuilder.RequestAdapter.SendNoContent(ctx, requestInfo, nil)
    if err != nil {
        return err
    }
    return nil
}
// Put adds a repository to an organization secret when the `visibility` forrepository access is set to `selected`. The visibility is set when you [Create orupdate an organization secret](https://docs.github.com/rest/actions/secrets#create-or-update-an-organization-secret).You must authenticate using an access token with the `admin:org` scope to use this endpoint.If the repository is private, you must use an access token with the `repo` scope.GitHub Apps must have the `secrets` organization permission to use this endpoint.Authenticated users must have collaborator access to a repository to create, update, or read secrets.
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/actions/secrets#add-selected-repository-to-an-organization-secret
func (m *ItemActionsSecretsItemRepositoriesWithRepository_ItemRequestBuilder) Put(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(error) {
    requestInfo, err := m.ToPutRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    err = m.BaseRequestBuilder.RequestAdapter.SendNoContent(ctx, requestInfo, nil)
    if err != nil {
        return err
    }
    return nil
}
// ToDeleteRequestInformation removes a repository from an organization secret when the `visibility`for repository access is set to `selected`. The visibility is set when you [Createor update an organization secret](https://docs.github.com/rest/actions/secrets#create-or-update-an-organization-secret).You must authenticate using an access token with the `admin:org` scope to use this endpoint.If the repository is private, you must use an access token with the `repo` scope.GitHub Apps must have the `secrets` organization permission to use this endpoint.Authenticated users must have collaborator access to a repository to create, update, or read secrets.
func (m *ItemActionsSecretsItemRepositoriesWithRepository_ItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ConfigureRequestInformation(requestInfo, requestConfiguration)
    return requestInfo, nil
}
// ToPutRequestInformation adds a repository to an organization secret when the `visibility` forrepository access is set to `selected`. The visibility is set when you [Create orupdate an organization secret](https://docs.github.com/rest/actions/secrets#create-or-update-an-organization-secret).You must authenticate using an access token with the `admin:org` scope to use this endpoint.If the repository is private, you must use an access token with the `repo` scope.GitHub Apps must have the `secrets` organization permission to use this endpoint.Authenticated users must have collaborator access to a repository to create, update, or read secrets.
func (m *ItemActionsSecretsItemRepositoriesWithRepository_ItemRequestBuilder) ToPutRequestInformation(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PUT, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ConfigureRequestInformation(requestInfo, requestConfiguration)
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
func (m *ItemActionsSecretsItemRepositoriesWithRepository_ItemRequestBuilder) WithUrl(rawUrl string)(*ItemActionsSecretsItemRepositoriesWithRepository_ItemRequestBuilder) {
    return NewItemActionsSecretsItemRepositoriesWithRepository_ItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
