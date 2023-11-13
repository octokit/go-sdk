package user

import (
    "context"
    i33bcfa37677dfd8b71eaf22ad3bfd140ba1702e7a8418558499c22418bf2fef2 "github.com/octokit/go-sdk/models"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i64e43db87b8155dc2264806ccd11f1ba3028759adf030ed8e2d0bb52f83589be "github.com/octokit/go-sdk/user/packages"
)

// PackagesRequestBuilder builds and executes requests for operations under \user\packages
type PackagesRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// PackagesRequestBuilderGetQueryParameters lists packages owned by the authenticated user within the user's namespace.To use this endpoint, you must authenticate using an access token with the `read:packages` scope. If the `package_type` belongs to a GitHub Packages registry that only supports repository-scoped permissions, your token must also include the `repo` scope. For the list of GitHub Packages registries that only support repository-scoped permissions, see "[About permissions for GitHub Packages](https://docs.github.com/packages/learn-github-packages/about-permissions-for-github-packages#permissions-for-repository-scoped-packages)."
type PackagesRequestBuilderGetQueryParameters struct {
    // The type of supported package. Packages in GitHub's Gradle registry have the type `maven`. Docker images pushed to GitHub's Container registry (`ghcr.io`) have the type `container`. You can use the type `docker` to find images that were pushed to GitHub's Docker registry (`docker.pkg.github.com`), even if these have now been migrated to the Container registry.
    // Deprecated: This property is deprecated, use package_typeAsGetPackage_typeQueryParameterType instead
    Package_type *string `uriparametername:"package_type"`
    // The type of supported package. Packages in GitHub's Gradle registry have the type `maven`. Docker images pushed to GitHub's Container registry (`ghcr.io`) have the type `container`. You can use the type `docker` to find images that were pushed to GitHub's Docker registry (`docker.pkg.github.com`), even if these have now been migrated to the Container registry.
    Package_typeAsGetPackage_typeQueryParameterType *i64e43db87b8155dc2264806ccd11f1ba3028759adf030ed8e2d0bb52f83589be.GetPackage_typeQueryParameterType `uriparametername:"package_type"`
    // Page number of the results to fetch.
    Page *int32 `uriparametername:"page"`
    // The number of results per page (max 100).
    Per_page *int32 `uriparametername:"per_page"`
    // The selected visibility of the packages.  This parameter is optional and only filters an existing result set.The `internal` visibility is only supported for GitHub Packages registries that allow for granular permissions. For other ecosystems `internal` is synonymous with `private`.For the list of GitHub Packages registries that support granular permissions, see "[About permissions for GitHub Packages](https://docs.github.com/packages/learn-github-packages/about-permissions-for-github-packages#granular-permissions-for-userorganization-scoped-packages)."
    // Deprecated: This property is deprecated, use visibilityAsGetVisibilityQueryParameterType instead
    Visibility *string `uriparametername:"visibility"`
    // The selected visibility of the packages.  This parameter is optional and only filters an existing result set.The `internal` visibility is only supported for GitHub Packages registries that allow for granular permissions. For other ecosystems `internal` is synonymous with `private`.For the list of GitHub Packages registries that support granular permissions, see "[About permissions for GitHub Packages](https://docs.github.com/packages/learn-github-packages/about-permissions-for-github-packages#granular-permissions-for-userorganization-scoped-packages)."
    VisibilityAsGetVisibilityQueryParameterType *i64e43db87b8155dc2264806ccd11f1ba3028759adf030ed8e2d0bb52f83589be.GetVisibilityQueryParameterType `uriparametername:"visibility"`
}
// PackagesRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type PackagesRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *PackagesRequestBuilderGetQueryParameters
}
// ByPackage_type gets an item from the go-sdk.user.packages.item collection
func (m *PackagesRequestBuilder) ByPackage_type(package_type string)(*PackagesWithPackage_typeItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if package_type != "" {
        urlTplParams["package_type"] = package_type
    }
    return NewPackagesWithPackage_typeItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewPackagesRequestBuilderInternal instantiates a new PackagesRequestBuilder and sets the default values.
func NewPackagesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*PackagesRequestBuilder) {
    m := &PackagesRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/user/packages{?package_type*,visibility*,page*,per_page*}", pathParameters),
    }
    return m
}
// NewPackagesRequestBuilder instantiates a new PackagesRequestBuilder and sets the default values.
func NewPackagesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*PackagesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewPackagesRequestBuilderInternal(urlParams, requestAdapter)
}
// Get lists packages owned by the authenticated user within the user's namespace.To use this endpoint, you must authenticate using an access token with the `read:packages` scope. If the `package_type` belongs to a GitHub Packages registry that only supports repository-scoped permissions, your token must also include the `repo` scope. For the list of GitHub Packages registries that only support repository-scoped permissions, see "[About permissions for GitHub Packages](https://docs.github.com/packages/learn-github-packages/about-permissions-for-github-packages#permissions-for-repository-scoped-packages)."
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/packages/packages#list-packages-for-the-authenticated-users-namespace
func (m *PackagesRequestBuilder) Get(ctx context.Context, requestConfiguration *PackagesRequestBuilderGetRequestConfiguration)([]i33bcfa37677dfd8b71eaf22ad3bfd140ba1702e7a8418558499c22418bf2fef2.PackageEscapedable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.SendCollection(ctx, requestInfo, i33bcfa37677dfd8b71eaf22ad3bfd140ba1702e7a8418558499c22418bf2fef2.CreatePackageEscapedFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    val := make([]i33bcfa37677dfd8b71eaf22ad3bfd140ba1702e7a8418558499c22418bf2fef2.PackageEscapedable, len(res))
    for i, v := range res {
        if v != nil {
            val[i] = v.(i33bcfa37677dfd8b71eaf22ad3bfd140ba1702e7a8418558499c22418bf2fef2.PackageEscapedable)
        }
    }
    return val, nil
}
// ToGetRequestInformation lists packages owned by the authenticated user within the user's namespace.To use this endpoint, you must authenticate using an access token with the `read:packages` scope. If the `package_type` belongs to a GitHub Packages registry that only supports repository-scoped permissions, your token must also include the `repo` scope. For the list of GitHub Packages registries that only support repository-scoped permissions, see "[About permissions for GitHub Packages](https://docs.github.com/packages/learn-github-packages/about-permissions-for-github-packages#permissions-for-repository-scoped-packages)."
func (m *PackagesRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *PackagesRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *PackagesRequestBuilder) WithUrl(rawUrl string)(*PackagesRequestBuilder) {
    return NewPackagesRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
