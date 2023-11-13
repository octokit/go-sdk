package orgs

import (
    "context"
    i33bcfa37677dfd8b71eaf22ad3bfd140ba1702e7a8418558499c22418bf2fef2 "github.com/octokit/go-sdk/models"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemSettingsBillingPackagesRequestBuilder builds and executes requests for operations under \orgs\{org}\settings\billing\packages
type ItemSettingsBillingPackagesRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemSettingsBillingPackagesRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemSettingsBillingPackagesRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemSettingsBillingPackagesRequestBuilderInternal instantiates a new PackagesRequestBuilder and sets the default values.
func NewItemSettingsBillingPackagesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemSettingsBillingPackagesRequestBuilder) {
    m := &ItemSettingsBillingPackagesRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/orgs/{org}/settings/billing/packages", pathParameters),
    }
    return m
}
// NewItemSettingsBillingPackagesRequestBuilder instantiates a new PackagesRequestBuilder and sets the default values.
func NewItemSettingsBillingPackagesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemSettingsBillingPackagesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemSettingsBillingPackagesRequestBuilderInternal(urlParams, requestAdapter)
}
// Get gets the free and paid storage used for GitHub Packages in gigabytes.Paid minutes only apply to packages stored for private repositories. For more information, see "[Managing billing for GitHub Packages](https://docs.github.com/github/setting-up-and-managing-billing-and-payments-on-github/managing-billing-for-github-packages)."Access tokens must have the `repo` or `admin:org` scope.
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/billing/billing#get-github-packages-billing-for-an-organization
func (m *ItemSettingsBillingPackagesRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemSettingsBillingPackagesRequestBuilderGetRequestConfiguration)(i33bcfa37677dfd8b71eaf22ad3bfd140ba1702e7a8418558499c22418bf2fef2.PackagesBillingUsageable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i33bcfa37677dfd8b71eaf22ad3bfd140ba1702e7a8418558499c22418bf2fef2.CreatePackagesBillingUsageFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i33bcfa37677dfd8b71eaf22ad3bfd140ba1702e7a8418558499c22418bf2fef2.PackagesBillingUsageable), nil
}
// ToGetRequestInformation gets the free and paid storage used for GitHub Packages in gigabytes.Paid minutes only apply to packages stored for private repositories. For more information, see "[Managing billing for GitHub Packages](https://docs.github.com/github/setting-up-and-managing-billing-and-payments-on-github/managing-billing-for-github-packages)."Access tokens must have the `repo` or `admin:org` scope.
func (m *ItemSettingsBillingPackagesRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemSettingsBillingPackagesRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    if requestConfiguration != nil {
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
func (m *ItemSettingsBillingPackagesRequestBuilder) WithUrl(rawUrl string)(*ItemSettingsBillingPackagesRequestBuilder) {
    return NewItemSettingsBillingPackagesRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
