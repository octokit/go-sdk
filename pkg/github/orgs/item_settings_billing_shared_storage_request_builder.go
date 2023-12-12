package orgs

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6 "github.com/octokit/go-sdk/pkg/github/models"
)

// ItemSettingsBillingSharedStorageRequestBuilder builds and executes requests for operations under \orgs\{org}\settings\billing\shared-storage
type ItemSettingsBillingSharedStorageRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemSettingsBillingSharedStorageRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemSettingsBillingSharedStorageRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemSettingsBillingSharedStorageRequestBuilderInternal instantiates a new SharedStorageRequestBuilder and sets the default values.
func NewItemSettingsBillingSharedStorageRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemSettingsBillingSharedStorageRequestBuilder) {
    m := &ItemSettingsBillingSharedStorageRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/orgs/{org}/settings/billing/shared-storage", pathParameters),
    }
    return m
}
// NewItemSettingsBillingSharedStorageRequestBuilder instantiates a new SharedStorageRequestBuilder and sets the default values.
func NewItemSettingsBillingSharedStorageRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemSettingsBillingSharedStorageRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemSettingsBillingSharedStorageRequestBuilderInternal(urlParams, requestAdapter)
}
// Get gets the estimated paid and estimated total storage used for GitHub Actions and GitHub Packages.Paid minutes only apply to packages stored for private repositories. For more information, see "[Managing billing for GitHub Packages](https://docs.github.com/github/setting-up-and-managing-billing-and-payments-on-github/managing-billing-for-github-packages)."Access tokens must have the `repo` or `admin:org` scope.
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/billing/billing#get-shared-storage-billing-for-an-organization
func (m *ItemSettingsBillingSharedStorageRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemSettingsBillingSharedStorageRequestBuilderGetRequestConfiguration)(i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CombinedBillingUsageable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CreateCombinedBillingUsageFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CombinedBillingUsageable), nil
}
// ToGetRequestInformation gets the estimated paid and estimated total storage used for GitHub Actions and GitHub Packages.Paid minutes only apply to packages stored for private repositories. For more information, see "[Managing billing for GitHub Packages](https://docs.github.com/github/setting-up-and-managing-billing-and-payments-on-github/managing-billing-for-github-packages)."Access tokens must have the `repo` or `admin:org` scope.
func (m *ItemSettingsBillingSharedStorageRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemSettingsBillingSharedStorageRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
func (m *ItemSettingsBillingSharedStorageRequestBuilder) WithUrl(rawUrl string)(*ItemSettingsBillingSharedStorageRequestBuilder) {
    return NewItemSettingsBillingSharedStorageRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
