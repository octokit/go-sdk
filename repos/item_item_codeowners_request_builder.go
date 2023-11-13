package repos

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemItemCodeownersRequestBuilder builds and executes requests for operations under \repos\{owner}\{repo}\codeowners
type ItemItemCodeownersRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// NewItemItemCodeownersRequestBuilderInternal instantiates a new CodeownersRequestBuilder and sets the default values.
func NewItemItemCodeownersRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemCodeownersRequestBuilder) {
    m := &ItemItemCodeownersRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/repos/{owner}/{repo}/codeowners", pathParameters),
    }
    return m
}
// NewItemItemCodeownersRequestBuilder instantiates a new CodeownersRequestBuilder and sets the default values.
func NewItemItemCodeownersRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemCodeownersRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemItemCodeownersRequestBuilderInternal(urlParams, requestAdapter)
}
// Errors the errors property
func (m *ItemItemCodeownersRequestBuilder) Errors()(*ItemItemCodeownersErrorsRequestBuilder) {
    return NewItemItemCodeownersErrorsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
