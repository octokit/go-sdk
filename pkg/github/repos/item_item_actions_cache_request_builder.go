package repos

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemItemActionsCacheRequestBuilder builds and executes requests for operations under \repos\{repos-id}\{Owner-id}\actions\cache
type ItemItemActionsCacheRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// NewItemItemActionsCacheRequestBuilderInternal instantiates a new CacheRequestBuilder and sets the default values.
func NewItemItemActionsCacheRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemActionsCacheRequestBuilder) {
    m := &ItemItemActionsCacheRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/repos/{repos%2Did}/{Owner%2Did}/actions/cache", pathParameters),
    }
    return m
}
// NewItemItemActionsCacheRequestBuilder instantiates a new CacheRequestBuilder and sets the default values.
func NewItemItemActionsCacheRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemActionsCacheRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemItemActionsCacheRequestBuilderInternal(urlParams, requestAdapter)
}
// Usage the usage property
func (m *ItemItemActionsCacheRequestBuilder) Usage()(*ItemItemActionsCacheUsageRequestBuilder) {
    return NewItemItemActionsCacheUsageRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
