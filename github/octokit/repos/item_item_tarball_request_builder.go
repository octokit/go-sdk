package repos

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemItemTarballRequestBuilder builds and executes requests for operations under \repos\{repos-id}\{Owner-id}\tarball
type ItemItemTarballRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ByRef gets an item from the github.com/octokit/go-sdk/github/octokit/.repos.item.item.tarball.item collection
func (m *ItemItemTarballRequestBuilder) ByRef(ref string)(*ItemItemTarballWithRefItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if ref != "" {
        urlTplParams["ref"] = ref
    }
    return NewItemItemTarballWithRefItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewItemItemTarballRequestBuilderInternal instantiates a new TarballRequestBuilder and sets the default values.
func NewItemItemTarballRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemTarballRequestBuilder) {
    m := &ItemItemTarballRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/repos/{repos%2Did}/{Owner%2Did}/tarball", pathParameters),
    }
    return m
}
// NewItemItemTarballRequestBuilder instantiates a new TarballRequestBuilder and sets the default values.
func NewItemItemTarballRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemTarballRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemItemTarballRequestBuilderInternal(urlParams, requestAdapter)
}
