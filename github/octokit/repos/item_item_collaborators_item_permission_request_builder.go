package repos

import (
    "context"
    i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e "github.com/octokit/go-sdk/github/octokit/models"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemItemCollaboratorsItemPermissionRequestBuilder builds and executes requests for operations under \repos\{owner}\{repo}\collaborators\{username}\permission
type ItemItemCollaboratorsItemPermissionRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemItemCollaboratorsItemPermissionRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemItemCollaboratorsItemPermissionRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemItemCollaboratorsItemPermissionRequestBuilderInternal instantiates a new PermissionRequestBuilder and sets the default values.
func NewItemItemCollaboratorsItemPermissionRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemCollaboratorsItemPermissionRequestBuilder) {
    m := &ItemItemCollaboratorsItemPermissionRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/repos/{owner}/{repo}/collaborators/{username}/permission", pathParameters),
    }
    return m
}
// NewItemItemCollaboratorsItemPermissionRequestBuilder instantiates a new PermissionRequestBuilder and sets the default values.
func NewItemItemCollaboratorsItemPermissionRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemCollaboratorsItemPermissionRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemItemCollaboratorsItemPermissionRequestBuilderInternal(urlParams, requestAdapter)
}
// Get checks the repository permission of a collaborator. The possible repositorypermissions are `admin`, `write`, `read`, and `none`.*Note*: The `permission` attribute provides the legacy base roles of `admin`, `write`, `read`, and `none`, where the`maintain` role is mapped to `write` and the `triage` role is mapped to `read`. To determine the role assigned to thecollaborator, see the `role_name` attribute, which will provide the full role name, including custom roles. The`permissions` hash can also be used to determine which base level of access the collaborator has to the repository.
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/collaborators/collaborators#get-repository-permissions-for-a-user
func (m *ItemItemCollaboratorsItemPermissionRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemItemCollaboratorsItemPermissionRequestBuilderGetRequestConfiguration)(i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.RepositoryCollaboratorPermissionable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "404": i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.CreateBasicErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.CreateRepositoryCollaboratorPermissionFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.RepositoryCollaboratorPermissionable), nil
}
// ToGetRequestInformation checks the repository permission of a collaborator. The possible repositorypermissions are `admin`, `write`, `read`, and `none`.*Note*: The `permission` attribute provides the legacy base roles of `admin`, `write`, `read`, and `none`, where the`maintain` role is mapped to `write` and the `triage` role is mapped to `read`. To determine the role assigned to thecollaborator, see the `role_name` attribute, which will provide the full role name, including custom roles. The`permissions` hash can also be used to determine which base level of access the collaborator has to the repository.
func (m *ItemItemCollaboratorsItemPermissionRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemItemCollaboratorsItemPermissionRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
func (m *ItemItemCollaboratorsItemPermissionRequestBuilder) WithUrl(rawUrl string)(*ItemItemCollaboratorsItemPermissionRequestBuilder) {
    return NewItemItemCollaboratorsItemPermissionRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
