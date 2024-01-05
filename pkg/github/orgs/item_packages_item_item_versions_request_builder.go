package orgs

import (
    "context"
    i53ac87e8cb3cc9276228f74d38694a208cacb99bb8ceb705eeae99fb88d4d274 "strconv"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6 "github.com/octokit/go-sdk/pkg/github/models"
    i9a34b8e68f433e2618d3e623e60543e9c60f8c189bf5031f9180afed55466b3f "github.com/octokit/go-sdk/pkg/github/orgs/item/packages/item/item/versions"
)

// ItemPackagesItemItemVersionsRequestBuilder builds and executes requests for operations under \orgs\{org}\packages\{package_type}\{package_name}\versions
type ItemPackagesItemItemVersionsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemPackagesItemItemVersionsRequestBuilderGetQueryParameters lists package versions for a package owned by an organization.If the `package_type` belongs to a GitHub Packages registry that only supports repository-scoped permissions, your token must also include the `repo` scope. For the list of GitHub Packages registries that only support repository-scoped permissions, see "[About permissions for GitHub Packages](https://docs.github.com/packages/learn-github-packages/about-permissions-for-github-packages#permissions-for-repository-scoped-packages)."
type ItemPackagesItemItemVersionsRequestBuilderGetQueryParameters struct {
    // The page number of the results to fetch. For more information, see "[Using pagination in the REST API](https://docs.github.com/rest/using-the-rest-api/using-pagination-in-the-rest-api)."
    Page *int32 `uriparametername:"page"`
    // The number of results per page (max 100). For more information, see "[Using pagination in the REST API](https://docs.github.com/rest/using-the-rest-api/using-pagination-in-the-rest-api)."
    Per_page *int32 `uriparametername:"per_page"`
    // The state of the package, either active or deleted.
    // Deprecated: This property is deprecated, use stateAsGetStateQueryParameterType instead
    State *string `uriparametername:"state"`
    // The state of the package, either active or deleted.
    StateAsGetStateQueryParameterType *i9a34b8e68f433e2618d3e623e60543e9c60f8c189bf5031f9180afed55466b3f.GetStateQueryParameterType `uriparametername:"state"`
}
// ItemPackagesItemItemVersionsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemPackagesItemItemVersionsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemPackagesItemItemVersionsRequestBuilderGetQueryParameters
}
// ByPackage_version_id gets an item from the github.com/octokit/go-sdk/pkg/github/.orgs.item.packages.item.item.versions.item collection
// Deprecated: This indexer is deprecated and will be removed in the next major version. Use the one with the typed parameter instead.
func (m *ItemPackagesItemItemVersionsRequestBuilder) ByPackage_version_id(package_version_id string)(*ItemPackagesItemItemVersionsWithPackage_version_ItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if package_version_id != "" {
        urlTplParams["package_version_id"] = package_version_id
    }
    return NewItemPackagesItemItemVersionsWithPackage_version_ItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// ByPackage_version_idInteger gets an item from the github.com/octokit/go-sdk/pkg/github/.orgs.item.packages.item.item.versions.item collection
func (m *ItemPackagesItemItemVersionsRequestBuilder) ByPackage_version_idInteger(package_version_id int32)(*ItemPackagesItemItemVersionsWithPackage_version_ItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    urlTplParams["package_version_id"] = i53ac87e8cb3cc9276228f74d38694a208cacb99bb8ceb705eeae99fb88d4d274.FormatInt(int64(package_version_id), 10)
    return NewItemPackagesItemItemVersionsWithPackage_version_ItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewItemPackagesItemItemVersionsRequestBuilderInternal instantiates a new VersionsRequestBuilder and sets the default values.
func NewItemPackagesItemItemVersionsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemPackagesItemItemVersionsRequestBuilder) {
    m := &ItemPackagesItemItemVersionsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/orgs/{org}/packages/{package_type}/{package_name}/versions{?page*,per_page*,state*}", pathParameters),
    }
    return m
}
// NewItemPackagesItemItemVersionsRequestBuilder instantiates a new VersionsRequestBuilder and sets the default values.
func NewItemPackagesItemItemVersionsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemPackagesItemItemVersionsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemPackagesItemItemVersionsRequestBuilderInternal(urlParams, requestAdapter)
}
// Get lists package versions for a package owned by an organization.If the `package_type` belongs to a GitHub Packages registry that only supports repository-scoped permissions, your token must also include the `repo` scope. For the list of GitHub Packages registries that only support repository-scoped permissions, see "[About permissions for GitHub Packages](https://docs.github.com/packages/learn-github-packages/about-permissions-for-github-packages#permissions-for-repository-scoped-packages)."
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/packages/packages#list-package-versions-for-a-package-owned-by-an-organization
func (m *ItemPackagesItemItemVersionsRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemPackagesItemItemVersionsRequestBuilderGetRequestConfiguration)([]i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.PackageVersionable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "401": i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CreateBasicErrorFromDiscriminatorValue,
        "403": i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CreateBasicErrorFromDiscriminatorValue,
        "404": i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CreateBasicErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.SendCollection(ctx, requestInfo, i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CreatePackageVersionFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    val := make([]i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.PackageVersionable, len(res))
    for i, v := range res {
        if v != nil {
            val[i] = v.(i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.PackageVersionable)
        }
    }
    return val, nil
}
// ToGetRequestInformation lists package versions for a package owned by an organization.If the `package_type` belongs to a GitHub Packages registry that only supports repository-scoped permissions, your token must also include the `repo` scope. For the list of GitHub Packages registries that only support repository-scoped permissions, see "[About permissions for GitHub Packages](https://docs.github.com/packages/learn-github-packages/about-permissions-for-github-packages#permissions-for-repository-scoped-packages)."
func (m *ItemPackagesItemItemVersionsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemPackagesItemItemVersionsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
func (m *ItemPackagesItemItemVersionsRequestBuilder) WithUrl(rawUrl string)(*ItemPackagesItemItemVersionsRequestBuilder) {
    return NewItemPackagesItemItemVersionsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
