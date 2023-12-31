package repos

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ReposRequestBuilder builds and executes requests for operations under \repos
type ReposRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ByReposId gets an item from the github.com/octokit/go-sdk/pkg/github/.repos.item collection
func (m *ReposRequestBuilder) ByReposId(reposId string)(*ReposItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if reposId != "" {
        urlTplParams["repos%2Did"] = reposId
    }
    return NewReposItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewReposRequestBuilderInternal instantiates a new ReposRequestBuilder and sets the default values.
func NewReposRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ReposRequestBuilder) {
    m := &ReposRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/repos", pathParameters),
    }
    return m
}
// NewReposRequestBuilder instantiates a new ReposRequestBuilder and sets the default values.
func NewReposRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ReposRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewReposRequestBuilderInternal(urlParams, requestAdapter)
}
