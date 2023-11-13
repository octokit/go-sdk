package repos

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemItemTrafficRequestBuilder builds and executes requests for operations under \repos\{owner}\{repo}\traffic
type ItemItemTrafficRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// Clones the clones property
func (m *ItemItemTrafficRequestBuilder) Clones()(*ItemItemTrafficClonesRequestBuilder) {
    return NewItemItemTrafficClonesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewItemItemTrafficRequestBuilderInternal instantiates a new TrafficRequestBuilder and sets the default values.
func NewItemItemTrafficRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemTrafficRequestBuilder) {
    m := &ItemItemTrafficRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/repos/{owner}/{repo}/traffic", pathParameters),
    }
    return m
}
// NewItemItemTrafficRequestBuilder instantiates a new TrafficRequestBuilder and sets the default values.
func NewItemItemTrafficRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemTrafficRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemItemTrafficRequestBuilderInternal(urlParams, requestAdapter)
}
// Popular the popular property
func (m *ItemItemTrafficRequestBuilder) Popular()(*ItemItemTrafficPopularRequestBuilder) {
    return NewItemItemTrafficPopularRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Views the views property
func (m *ItemItemTrafficRequestBuilder) Views()(*ItemItemTrafficViewsRequestBuilder) {
    return NewItemItemTrafficViewsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
