package user

import (
    "context"
    i53ac87e8cb3cc9276228f74d38694a208cacb99bb8ceb705eeae99fb88d4d274 "strconv"
    i33bcfa37677dfd8b71eaf22ad3bfd140ba1702e7a8418558499c22418bf2fef2 "github.com/octokit/go-sdk/models"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// InstallationsItemRepositoriesRequestBuilder builds and executes requests for operations under \user\installations\{installation_id}\repositories
type InstallationsItemRepositoriesRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// InstallationsItemRepositoriesRequestBuilderGetQueryParameters list repositories that the authenticated user has explicit permission (`:read`, `:write`, or `:admin`) to access for an installation.The authenticated user has explicit permission to access repositories they own, repositories where they are a collaborator, and repositories that they can access through an organization membership.You must use a [user access token](https://docs.github.com/apps/creating-github-apps/authenticating-with-a-github-app/generating-a-user-access-token-for-a-github-app), created for a user who has authorized your GitHub App, to access this endpoint.The access the user has to each repository is included in the hash under the `permissions` key.
type InstallationsItemRepositoriesRequestBuilderGetQueryParameters struct {
    // Page number of the results to fetch.
    Page *int32 `uriparametername:"page"`
    // The number of results per page (max 100).
    Per_page *int32 `uriparametername:"per_page"`
}
// InstallationsItemRepositoriesRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type InstallationsItemRepositoriesRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *InstallationsItemRepositoriesRequestBuilderGetQueryParameters
}
// ByRepository_id gets an item from the go-sdk.user.installations.item.repositories.item collection
// Deprecated: This indexer is deprecated and will be removed in the next major version. Use the one with the typed parameter instead.
func (m *InstallationsItemRepositoriesRequestBuilder) ByRepository_id(repository_id string)(*InstallationsItemRepositoriesWithRepository_ItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if repository_id != "" {
        urlTplParams["repository_id"] = repository_id
    }
    return NewInstallationsItemRepositoriesWithRepository_ItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// ByRepository_idInteger gets an item from the go-sdk.user.installations.item.repositories.item collection
func (m *InstallationsItemRepositoriesRequestBuilder) ByRepository_idInteger(repository_id int32)(*InstallationsItemRepositoriesWithRepository_ItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    urlTplParams["repository_id"] = i53ac87e8cb3cc9276228f74d38694a208cacb99bb8ceb705eeae99fb88d4d274.FormatInt(int64(repository_id), 10)
    return NewInstallationsItemRepositoriesWithRepository_ItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewInstallationsItemRepositoriesRequestBuilderInternal instantiates a new RepositoriesRequestBuilder and sets the default values.
func NewInstallationsItemRepositoriesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*InstallationsItemRepositoriesRequestBuilder) {
    m := &InstallationsItemRepositoriesRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/user/installations/{installation_id}/repositories{?per_page*,page*}", pathParameters),
    }
    return m
}
// NewInstallationsItemRepositoriesRequestBuilder instantiates a new RepositoriesRequestBuilder and sets the default values.
func NewInstallationsItemRepositoriesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*InstallationsItemRepositoriesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewInstallationsItemRepositoriesRequestBuilderInternal(urlParams, requestAdapter)
}
// Get list repositories that the authenticated user has explicit permission (`:read`, `:write`, or `:admin`) to access for an installation.The authenticated user has explicit permission to access repositories they own, repositories where they are a collaborator, and repositories that they can access through an organization membership.You must use a [user access token](https://docs.github.com/apps/creating-github-apps/authenticating-with-a-github-app/generating-a-user-access-token-for-a-github-app), created for a user who has authorized your GitHub App, to access this endpoint.The access the user has to each repository is included in the hash under the `permissions` key.
// Deprecated: This method is obsolete. Use GetAsRepositoriesGetResponse instead.
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/apps/installations#list-repositories-accessible-to-the-user-access-token
func (m *InstallationsItemRepositoriesRequestBuilder) Get(ctx context.Context, requestConfiguration *InstallationsItemRepositoriesRequestBuilderGetRequestConfiguration)(InstallationsItemRepositoriesResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "403": i33bcfa37677dfd8b71eaf22ad3bfd140ba1702e7a8418558499c22418bf2fef2.CreateBasicErrorFromDiscriminatorValue,
        "404": i33bcfa37677dfd8b71eaf22ad3bfd140ba1702e7a8418558499c22418bf2fef2.CreateBasicErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateInstallationsItemRepositoriesResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(InstallationsItemRepositoriesResponseable), nil
}
// GetAsRepositoriesGetResponse list repositories that the authenticated user has explicit permission (`:read`, `:write`, or `:admin`) to access for an installation.The authenticated user has explicit permission to access repositories they own, repositories where they are a collaborator, and repositories that they can access through an organization membership.You must use a [user access token](https://docs.github.com/apps/creating-github-apps/authenticating-with-a-github-app/generating-a-user-access-token-for-a-github-app), created for a user who has authorized your GitHub App, to access this endpoint.The access the user has to each repository is included in the hash under the `permissions` key.
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/apps/installations#list-repositories-accessible-to-the-user-access-token
func (m *InstallationsItemRepositoriesRequestBuilder) GetAsRepositoriesGetResponse(ctx context.Context, requestConfiguration *InstallationsItemRepositoriesRequestBuilderGetRequestConfiguration)(InstallationsItemRepositoriesGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "403": i33bcfa37677dfd8b71eaf22ad3bfd140ba1702e7a8418558499c22418bf2fef2.CreateBasicErrorFromDiscriminatorValue,
        "404": i33bcfa37677dfd8b71eaf22ad3bfd140ba1702e7a8418558499c22418bf2fef2.CreateBasicErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateInstallationsItemRepositoriesGetResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(InstallationsItemRepositoriesGetResponseable), nil
}
// ToGetRequestInformation list repositories that the authenticated user has explicit permission (`:read`, `:write`, or `:admin`) to access for an installation.The authenticated user has explicit permission to access repositories they own, repositories where they are a collaborator, and repositories that they can access through an organization membership.You must use a [user access token](https://docs.github.com/apps/creating-github-apps/authenticating-with-a-github-app/generating-a-user-access-token-for-a-github-app), created for a user who has authorized your GitHub App, to access this endpoint.The access the user has to each repository is included in the hash under the `permissions` key.
func (m *InstallationsItemRepositoriesRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *InstallationsItemRepositoriesRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers.TryAdd("Accept", "application/json;q=1")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
func (m *InstallationsItemRepositoriesRequestBuilder) WithUrl(rawUrl string)(*InstallationsItemRepositoriesRequestBuilder) {
    return NewInstallationsItemRepositoriesRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
