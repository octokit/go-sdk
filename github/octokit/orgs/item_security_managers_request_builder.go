package orgs

import (
    "context"
    i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e "github.com/octokit/go-sdk/github/octokit/models"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemSecurityManagersRequestBuilder builds and executes requests for operations under \orgs\{org}\security-managers
type ItemSecurityManagersRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemSecurityManagersRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemSecurityManagersRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemSecurityManagersRequestBuilderInternal instantiates a new SecurityManagersRequestBuilder and sets the default values.
func NewItemSecurityManagersRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemSecurityManagersRequestBuilder) {
    m := &ItemSecurityManagersRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/orgs/{org}/security-managers", pathParameters),
    }
    return m
}
// NewItemSecurityManagersRequestBuilder instantiates a new SecurityManagersRequestBuilder and sets the default values.
func NewItemSecurityManagersRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemSecurityManagersRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemSecurityManagersRequestBuilderInternal(urlParams, requestAdapter)
}
// Get lists teams that are security managers for an organization. For more information, see "[Managing security managers in your organization](https://docs.github.com/organizations/managing-peoples-access-to-your-organization-with-roles/managing-security-managers-in-your-organization)."To use this endpoint, you must be an administrator or security manager for the organization, and you must use an access token with the `read:org` scope.GitHub Apps must have the `administration` organization read permission to use this endpoint.
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/orgs/security-managers#list-security-manager-teams
func (m *ItemSecurityManagersRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemSecurityManagersRequestBuilderGetRequestConfiguration)([]i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.TeamSimpleable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.SendCollection(ctx, requestInfo, i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.CreateTeamSimpleFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    val := make([]i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.TeamSimpleable, len(res))
    for i, v := range res {
        if v != nil {
            val[i] = v.(i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.TeamSimpleable)
        }
    }
    return val, nil
}
// Teams the teams property
func (m *ItemSecurityManagersRequestBuilder) Teams()(*ItemSecurityManagersTeamsRequestBuilder) {
    return NewItemSecurityManagersTeamsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToGetRequestInformation lists teams that are security managers for an organization. For more information, see "[Managing security managers in your organization](https://docs.github.com/organizations/managing-peoples-access-to-your-organization-with-roles/managing-security-managers-in-your-organization)."To use this endpoint, you must be an administrator or security manager for the organization, and you must use an access token with the `read:org` scope.GitHub Apps must have the `administration` organization read permission to use this endpoint.
func (m *ItemSecurityManagersRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemSecurityManagersRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *ItemSecurityManagersRequestBuilder) WithUrl(rawUrl string)(*ItemSecurityManagersRequestBuilder) {
    return NewItemSecurityManagersRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
