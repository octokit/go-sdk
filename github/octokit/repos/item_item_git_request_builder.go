package repos

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemItemGitRequestBuilder builds and executes requests for operations under \repos\{repos-id}\{Owner-id}\git
type ItemItemGitRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// Blobs the blobs property
func (m *ItemItemGitRequestBuilder) Blobs()(*ItemItemGitBlobsRequestBuilder) {
    return NewItemItemGitBlobsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Commits the commits property
func (m *ItemItemGitRequestBuilder) Commits()(*ItemItemGitCommitsRequestBuilder) {
    return NewItemItemGitCommitsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewItemItemGitRequestBuilderInternal instantiates a new GitRequestBuilder and sets the default values.
func NewItemItemGitRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemGitRequestBuilder) {
    m := &ItemItemGitRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/repos/{repos%2Did}/{Owner%2Did}/git", pathParameters),
    }
    return m
}
// NewItemItemGitRequestBuilder instantiates a new GitRequestBuilder and sets the default values.
func NewItemItemGitRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemGitRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemItemGitRequestBuilderInternal(urlParams, requestAdapter)
}
// MatchingRefs the matchingRefs property
func (m *ItemItemGitRequestBuilder) MatchingRefs()(*ItemItemGitMatchingRefsRequestBuilder) {
    return NewItemItemGitMatchingRefsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Ref the ref property
func (m *ItemItemGitRequestBuilder) Ref()(*ItemItemGitRefRequestBuilder) {
    return NewItemItemGitRefRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Refs the refs property
func (m *ItemItemGitRequestBuilder) Refs()(*ItemItemGitRefsRequestBuilder) {
    return NewItemItemGitRefsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Tags the tags property
func (m *ItemItemGitRequestBuilder) Tags()(*ItemItemGitTagsRequestBuilder) {
    return NewItemItemGitTagsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Trees the trees property
func (m *ItemItemGitRequestBuilder) Trees()(*ItemItemGitTreesRequestBuilder) {
    return NewItemItemGitTreesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
