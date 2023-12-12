package orgs

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemCopilotRequestBuilder builds and executes requests for operations under \orgs\{org}\copilot
type ItemCopilotRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// Billing the billing property
func (m *ItemCopilotRequestBuilder) Billing()(*ItemCopilotBillingRequestBuilder) {
    return NewItemCopilotBillingRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewItemCopilotRequestBuilderInternal instantiates a new CopilotRequestBuilder and sets the default values.
func NewItemCopilotRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemCopilotRequestBuilder) {
    m := &ItemCopilotRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/orgs/{org}/copilot", pathParameters),
    }
    return m
}
// NewItemCopilotRequestBuilder instantiates a new CopilotRequestBuilder and sets the default values.
func NewItemCopilotRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemCopilotRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemCopilotRequestBuilderInternal(urlParams, requestAdapter)
}
