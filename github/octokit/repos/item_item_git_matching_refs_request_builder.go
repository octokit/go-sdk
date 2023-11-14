package repos

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemItemGitMatchingRefsRequestBuilder builds and executes requests for operations under \repos\{owner}\{repo}\git\matching-refs
type ItemItemGitMatchingRefsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ByRef gets an item from the octokit.repos.item.item.git.matchingRefs.item collection
func (m *ItemItemGitMatchingRefsRequestBuilder) ByRef(ref string)(*ItemItemGitMatchingRefsWithRefItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if ref != "" {
        urlTplParams["ref"] = ref
    }
    return NewItemItemGitMatchingRefsWithRefItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewItemItemGitMatchingRefsRequestBuilderInternal instantiates a new MatchingRefsRequestBuilder and sets the default values.
func NewItemItemGitMatchingRefsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemGitMatchingRefsRequestBuilder) {
    m := &ItemItemGitMatchingRefsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/repos/{owner}/{repo}/git/matching-refs", pathParameters),
    }
    return m
}
// NewItemItemGitMatchingRefsRequestBuilder instantiates a new MatchingRefsRequestBuilder and sets the default values.
func NewItemItemGitMatchingRefsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemGitMatchingRefsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemItemGitMatchingRefsRequestBuilderInternal(urlParams, requestAdapter)
}
