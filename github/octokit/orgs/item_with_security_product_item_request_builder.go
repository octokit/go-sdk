package orgs

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemWithSecurity_productItemRequestBuilder builds and executes requests for operations under \orgs\{org}\{security_product}
type ItemWithSecurity_productItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ByEnablement gets an item from the github.com/octokit/go-sdk/github/octokit/.orgs.item.item.item collection
func (m *ItemWithSecurity_productItemRequestBuilder) ByEnablement(enablement string)(*ItemItemWithEnablementItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if enablement != "" {
        urlTplParams["enablement"] = enablement
    }
    return NewItemItemWithEnablementItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewItemWithSecurity_productItemRequestBuilderInternal instantiates a new WithSecurity_productItemRequestBuilder and sets the default values.
func NewItemWithSecurity_productItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemWithSecurity_productItemRequestBuilder) {
    m := &ItemWithSecurity_productItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/orgs/{org}/{security_product}", pathParameters),
    }
    return m
}
// NewItemWithSecurity_productItemRequestBuilder instantiates a new WithSecurity_productItemRequestBuilder and sets the default values.
func NewItemWithSecurity_productItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemWithSecurity_productItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemWithSecurity_productItemRequestBuilderInternal(urlParams, requestAdapter)
}
