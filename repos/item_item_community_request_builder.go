package repos

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemItemCommunityRequestBuilder builds and executes requests for operations under \repos\{owner}\{repo}\community
type ItemItemCommunityRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// NewItemItemCommunityRequestBuilderInternal instantiates a new CommunityRequestBuilder and sets the default values.
func NewItemItemCommunityRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemCommunityRequestBuilder) {
    m := &ItemItemCommunityRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/repos/{owner}/{repo}/community", pathParameters),
    }
    return m
}
// NewItemItemCommunityRequestBuilder instantiates a new CommunityRequestBuilder and sets the default values.
func NewItemItemCommunityRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemCommunityRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemItemCommunityRequestBuilderInternal(urlParams, requestAdapter)
}
// Profile the profile property
func (m *ItemItemCommunityRequestBuilder) Profile()(*ItemItemCommunityProfileRequestBuilder) {
    return NewItemItemCommunityProfileRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
