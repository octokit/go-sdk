package repos

import (
    "context"
    i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e "github.com/octokit/go-sdk/github/octokit/models"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemItemSecurityAdvisoriesWithGhsa_ItemRequestBuilder builds and executes requests for operations under \repos\{owner}\{repo}\security-advisories\{ghsa_id}
type ItemItemSecurityAdvisoriesWithGhsa_ItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemItemSecurityAdvisoriesWithGhsa_ItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemItemSecurityAdvisoriesWithGhsa_ItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ItemItemSecurityAdvisoriesWithGhsa_ItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemItemSecurityAdvisoriesWithGhsa_ItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemItemSecurityAdvisoriesWithGhsa_ItemRequestBuilderInternal instantiates a new WithGhsa_ItemRequestBuilder and sets the default values.
func NewItemItemSecurityAdvisoriesWithGhsa_ItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemSecurityAdvisoriesWithGhsa_ItemRequestBuilder) {
    m := &ItemItemSecurityAdvisoriesWithGhsa_ItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/repos/{owner}/{repo}/security-advisories/{ghsa_id}", pathParameters),
    }
    return m
}
// NewItemItemSecurityAdvisoriesWithGhsa_ItemRequestBuilder instantiates a new WithGhsa_ItemRequestBuilder and sets the default values.
func NewItemItemSecurityAdvisoriesWithGhsa_ItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemSecurityAdvisoriesWithGhsa_ItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemItemSecurityAdvisoriesWithGhsa_ItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Cve the cve property
func (m *ItemItemSecurityAdvisoriesWithGhsa_ItemRequestBuilder) Cve()(*ItemItemSecurityAdvisoriesItemCveRequestBuilder) {
    return NewItemItemSecurityAdvisoriesItemCveRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get get a repository security advisory using its GitHub Security Advisory (GHSA) identifier.You can access any published security advisory on a public repository.You must authenticate using an access token with the `repo` scope or `repository_advisories:read` permissionin order to get a published security advisory in a private repository, or any unpublished security advisory that you have access to.You can access an unpublished security advisory from a repository if you are a security manager or administrator of that repository, or if you are acollaborator on the security advisory.
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/security-advisories/repository-advisories#get-a-repository-security-advisory
func (m *ItemItemSecurityAdvisoriesWithGhsa_ItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemItemSecurityAdvisoriesWithGhsa_ItemRequestBuilderGetRequestConfiguration)(i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.RepositoryAdvisoryable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "403": i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.CreateBasicErrorFromDiscriminatorValue,
        "404": i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.CreateBasicErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.CreateRepositoryAdvisoryFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.RepositoryAdvisoryable), nil
}
// Patch update a repository security advisory using its GitHub Security Advisory (GHSA) identifier.You must authenticate using an access token with the `repo` scope or `repository_advisories:write` permission to use this endpoint.In order to update any security advisory, you must be a security manager or administrator of that repository,or a collaborator on the repository security advisory.
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/security-advisories/repository-advisories#update-a-repository-security-advisory
func (m *ItemItemSecurityAdvisoriesWithGhsa_ItemRequestBuilder) Patch(ctx context.Context, body i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.RepositoryAdvisoryUpdateable, requestConfiguration *ItemItemSecurityAdvisoriesWithGhsa_ItemRequestBuilderPatchRequestConfiguration)(i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.RepositoryAdvisoryable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "403": i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.CreateBasicErrorFromDiscriminatorValue,
        "404": i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.CreateBasicErrorFromDiscriminatorValue,
        "422": i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.CreateValidationErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.CreateRepositoryAdvisoryFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.RepositoryAdvisoryable), nil
}
// ToGetRequestInformation get a repository security advisory using its GitHub Security Advisory (GHSA) identifier.You can access any published security advisory on a public repository.You must authenticate using an access token with the `repo` scope or `repository_advisories:read` permissionin order to get a published security advisory in a private repository, or any unpublished security advisory that you have access to.You can access an unpublished security advisory from a repository if you are a security manager or administrator of that repository, or if you are acollaborator on the security advisory.
func (m *ItemItemSecurityAdvisoriesWithGhsa_ItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemItemSecurityAdvisoriesWithGhsa_ItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update a repository security advisory using its GitHub Security Advisory (GHSA) identifier.You must authenticate using an access token with the `repo` scope or `repository_advisories:write` permission to use this endpoint.In order to update any security advisory, you must be a security manager or administrator of that repository,or a collaborator on the repository security advisory.
func (m *ItemItemSecurityAdvisoriesWithGhsa_ItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.RepositoryAdvisoryUpdateable, requestConfiguration *ItemItemSecurityAdvisoriesWithGhsa_ItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
    requestInfo.Headers.TryAdd("Accept", "application/json")
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
func (m *ItemItemSecurityAdvisoriesWithGhsa_ItemRequestBuilder) WithUrl(rawUrl string)(*ItemItemSecurityAdvisoriesWithGhsa_ItemRequestBuilder) {
    return NewItemItemSecurityAdvisoriesWithGhsa_ItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
