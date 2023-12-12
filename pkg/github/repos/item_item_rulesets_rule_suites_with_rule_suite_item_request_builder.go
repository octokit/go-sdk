package repos

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6 "github.com/octokit/go-sdk/pkg/github/models"
)

// ItemItemRulesetsRuleSuitesWithRule_suite_ItemRequestBuilder builds and executes requests for operations under \repos\{repos-id}\{Owner-id}\rulesets\rule-suites\{rule_suite_id}
type ItemItemRulesetsRuleSuitesWithRule_suite_ItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemItemRulesetsRuleSuitesWithRule_suite_ItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemItemRulesetsRuleSuitesWithRule_suite_ItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemItemRulesetsRuleSuitesWithRule_suite_ItemRequestBuilderInternal instantiates a new WithRule_suite_ItemRequestBuilder and sets the default values.
func NewItemItemRulesetsRuleSuitesWithRule_suite_ItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemRulesetsRuleSuitesWithRule_suite_ItemRequestBuilder) {
    m := &ItemItemRulesetsRuleSuitesWithRule_suite_ItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/repos/{repos%2Did}/{Owner%2Did}/rulesets/rule-suites/{rule_suite_id}", pathParameters),
    }
    return m
}
// NewItemItemRulesetsRuleSuitesWithRule_suite_ItemRequestBuilder instantiates a new WithRule_suite_ItemRequestBuilder and sets the default values.
func NewItemItemRulesetsRuleSuitesWithRule_suite_ItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemRulesetsRuleSuitesWithRule_suite_ItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemItemRulesetsRuleSuitesWithRule_suite_ItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Get gets information about a suite of rule evaluations from within a repository.For more information, see "[Managing rulesets for a repository](https://docs.github.com/repositories/configuring-branches-and-merges-in-your-repository/managing-rulesets/managing-rulesets-for-a-repository#viewing-insights-for-rulesets)."
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/repos/rule-suites#get-a-repository-rule-suite
func (m *ItemItemRulesetsRuleSuitesWithRule_suite_ItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemItemRulesetsRuleSuitesWithRule_suite_ItemRequestBuilderGetRequestConfiguration)(i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.RuleSuiteable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "404": i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CreateBasicErrorFromDiscriminatorValue,
        "500": i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CreateBasicErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CreateRuleSuiteFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.RuleSuiteable), nil
}
// ToGetRequestInformation gets information about a suite of rule evaluations from within a repository.For more information, see "[Managing rulesets for a repository](https://docs.github.com/repositories/configuring-branches-and-merges-in-your-repository/managing-rulesets/managing-rulesets-for-a-repository#viewing-insights-for-rulesets)."
func (m *ItemItemRulesetsRuleSuitesWithRule_suite_ItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemItemRulesetsRuleSuitesWithRule_suite_ItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
func (m *ItemItemRulesetsRuleSuitesWithRule_suite_ItemRequestBuilder) WithUrl(rawUrl string)(*ItemItemRulesetsRuleSuitesWithRule_suite_ItemRequestBuilder) {
    return NewItemItemRulesetsRuleSuitesWithRule_suite_ItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
