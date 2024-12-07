package organizations

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemSettingsBillingRequestBuilder builds and executes requests for operations under \organizations\{org}\settings\billing
type ItemSettingsBillingRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// NewItemSettingsBillingRequestBuilderInternal instantiates a new ItemSettingsBillingRequestBuilder and sets the default values.
func NewItemSettingsBillingRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemSettingsBillingRequestBuilder) {
    m := &ItemSettingsBillingRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/organizations/{org}/settings/billing", pathParameters),
    }
    return m
}
// NewItemSettingsBillingRequestBuilder instantiates a new ItemSettingsBillingRequestBuilder and sets the default values.
func NewItemSettingsBillingRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemSettingsBillingRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemSettingsBillingRequestBuilderInternal(urlParams, requestAdapter)
}
// Usage the usage property
// returns a *ItemSettingsBillingUsageRequestBuilder when successful
func (m *ItemSettingsBillingRequestBuilder) Usage()(*ItemSettingsBillingUsageRequestBuilder) {
    return NewItemSettingsBillingUsageRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
