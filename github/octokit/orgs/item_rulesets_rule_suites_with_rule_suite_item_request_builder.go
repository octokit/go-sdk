package orgs

import (
    "context"
    i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e "github.com/octokit/go-sdk/github/octokit/models"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemRulesetsRuleSuitesWithRule_suite_ItemRequestBuilder builds and executes requests for operations under \orgs\{org}\rulesets\rule-suites\{rule_suite_id}
type ItemRulesetsRuleSuitesWithRule_suite_ItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemRulesetsRuleSuitesWithRule_suite_ItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemRulesetsRuleSuitesWithRule_suite_ItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemRulesetsRuleSuitesWithRule_suite_ItemRequestBuilderInternal instantiates a new WithRule_suite_ItemRequestBuilder and sets the default values.
func NewItemRulesetsRuleSuitesWithRule_suite_ItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemRulesetsRuleSuitesWithRule_suite_ItemRequestBuilder) {
    m := &ItemRulesetsRuleSuitesWithRule_suite_ItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/orgs/{org}/rulesets/rule-suites/{rule_suite_id}", pathParameters),
    }
    return m
}
// NewItemRulesetsRuleSuitesWithRule_suite_ItemRequestBuilder instantiates a new WithRule_suite_ItemRequestBuilder and sets the default values.
func NewItemRulesetsRuleSuitesWithRule_suite_ItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemRulesetsRuleSuitesWithRule_suite_ItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemRulesetsRuleSuitesWithRule_suite_ItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Get gets information about a suite of rule evaluations from within an organization.For more information, see "[Managing rulesets for repositories in your organization](https://docs.github.com/organizations/managing-organization-settings/managing-rulesets-for-repositories-in-your-organization#viewing-insights-for-rulesets)."
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/orgs/rule-suites#get-an-organization-rule-suite
func (m *ItemRulesetsRuleSuitesWithRule_suite_ItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemRulesetsRuleSuitesWithRule_suite_ItemRequestBuilderGetRequestConfiguration)(i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.RuleSuiteable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "404": i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.CreateBasicErrorFromDiscriminatorValue,
        "500": i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.CreateBasicErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.CreateRuleSuiteFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.RuleSuiteable), nil
}
// ToGetRequestInformation gets information about a suite of rule evaluations from within an organization.For more information, see "[Managing rulesets for repositories in your organization](https://docs.github.com/organizations/managing-organization-settings/managing-rulesets-for-repositories-in-your-organization#viewing-insights-for-rulesets)."
func (m *ItemRulesetsRuleSuitesWithRule_suite_ItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemRulesetsRuleSuitesWithRule_suite_ItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers.TryAdd("Accept", "application/json;q=1")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
func (m *ItemRulesetsRuleSuitesWithRule_suite_ItemRequestBuilder) WithUrl(rawUrl string)(*ItemRulesetsRuleSuitesWithRule_suite_ItemRequestBuilder) {
    return NewItemRulesetsRuleSuitesWithRule_suite_ItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
