package repos

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemItemRulesRequestBuilder builds and executes requests for operations under \repos\{owner}\{repo}\rules
type ItemItemRulesRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// Branches the branches property
func (m *ItemItemRulesRequestBuilder) Branches()(*ItemItemRulesBranchesRequestBuilder) {
    return NewItemItemRulesBranchesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewItemItemRulesRequestBuilderInternal instantiates a new RulesRequestBuilder and sets the default values.
func NewItemItemRulesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemRulesRequestBuilder) {
    m := &ItemItemRulesRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/repos/{owner}/{repo}/rules", pathParameters),
    }
    return m
}
// NewItemItemRulesRequestBuilder instantiates a new RulesRequestBuilder and sets the default values.
func NewItemItemRulesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemRulesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemItemRulesRequestBuilderInternal(urlParams, requestAdapter)
}
