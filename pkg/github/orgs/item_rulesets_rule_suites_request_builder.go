package orgs

import (
    "context"
    i53ac87e8cb3cc9276228f74d38694a208cacb99bb8ceb705eeae99fb88d4d274 "strconv"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6 "github.com/octokit/go-sdk/pkg/github/models"
    ice77a9476603b5b5bdcf7933e87c368075e96312df44ed011cd574d7f2e817b9 "github.com/octokit/go-sdk/pkg/github/orgs/item/rulesets/rulesuites"
)

// ItemRulesetsRuleSuitesRequestBuilder builds and executes requests for operations under \orgs\{org}\rulesets\rule-suites
type ItemRulesetsRuleSuitesRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemRulesetsRuleSuitesRequestBuilderGetQueryParameters lists suites of rule evaluations at the organization level.For more information, see "[Managing rulesets for repositories in your organization](https://docs.github.com/organizations/managing-organization-settings/managing-rulesets-for-repositories-in-your-organization#viewing-insights-for-rulesets)."
type ItemRulesetsRuleSuitesRequestBuilderGetQueryParameters struct {
    // The handle for the GitHub user account to filter on. When specified, only rule evaluations triggered by this actor will be returned.
    Actor_name *string `uriparametername:"actor_name"`
    // The page number of the results to fetch. For more information, see "[Using pagination in the REST API](https://docs.github.com/rest/using-the-rest-api/using-pagination-in-the-rest-api)."
    Page *int32 `uriparametername:"page"`
    // The number of results per page (max 100). For more information, see "[Using pagination in the REST API](https://docs.github.com/rest/using-the-rest-api/using-pagination-in-the-rest-api)."
    Per_page *int32 `uriparametername:"per_page"`
    // The name of the repository to filter on. When specified, only rule evaluations from this repository will be returned.
    Repository_name *int32 `uriparametername:"repository_name"`
    // The rule results to filter on. When specified, only suites with this result will be returned.
    // Deprecated: This property is deprecated, use rule_suite_resultAsGetRule_suite_resultQueryParameterType instead
    Rule_suite_result *string `uriparametername:"rule_suite_result"`
    // The rule results to filter on. When specified, only suites with this result will be returned.
    Rule_suite_resultAsGetRule_suite_resultQueryParameterType *ice77a9476603b5b5bdcf7933e87c368075e96312df44ed011cd574d7f2e817b9.GetRule_suite_resultQueryParameterType `uriparametername:"rule_suite_result"`
    // The time period to filter by.For example, `day` will filter for rule suites that occurred in the past 24 hours, and `week` will filter for insights that occurred in the past 7 days (168 hours).
    // Deprecated: This property is deprecated, use time_periodAsGetTime_periodQueryParameterType instead
    Time_period *string `uriparametername:"time_period"`
    // The time period to filter by.For example, `day` will filter for rule suites that occurred in the past 24 hours, and `week` will filter for insights that occurred in the past 7 days (168 hours).
    Time_periodAsGetTime_periodQueryParameterType *ice77a9476603b5b5bdcf7933e87c368075e96312df44ed011cd574d7f2e817b9.GetTime_periodQueryParameterType `uriparametername:"time_period"`
}
// ItemRulesetsRuleSuitesRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemRulesetsRuleSuitesRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemRulesetsRuleSuitesRequestBuilderGetQueryParameters
}
// ByRule_suite_id gets an item from the github.com/octokit/go-sdk/pkg/github/.orgs.item.rulesets.ruleSuites.item collection
// Deprecated: This indexer is deprecated and will be removed in the next major version. Use the one with the typed parameter instead.
func (m *ItemRulesetsRuleSuitesRequestBuilder) ByRule_suite_id(rule_suite_id string)(*ItemRulesetsRuleSuitesWithRule_suite_ItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if rule_suite_id != "" {
        urlTplParams["rule_suite_id"] = rule_suite_id
    }
    return NewItemRulesetsRuleSuitesWithRule_suite_ItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// ByRule_suite_idInteger gets an item from the github.com/octokit/go-sdk/pkg/github/.orgs.item.rulesets.ruleSuites.item collection
func (m *ItemRulesetsRuleSuitesRequestBuilder) ByRule_suite_idInteger(rule_suite_id int32)(*ItemRulesetsRuleSuitesWithRule_suite_ItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    urlTplParams["rule_suite_id"] = i53ac87e8cb3cc9276228f74d38694a208cacb99bb8ceb705eeae99fb88d4d274.FormatInt(int64(rule_suite_id), 10)
    return NewItemRulesetsRuleSuitesWithRule_suite_ItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewItemRulesetsRuleSuitesRequestBuilderInternal instantiates a new RuleSuitesRequestBuilder and sets the default values.
func NewItemRulesetsRuleSuitesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemRulesetsRuleSuitesRequestBuilder) {
    m := &ItemRulesetsRuleSuitesRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/orgs/{org}/rulesets/rule-suites{?repository_name*,time_period*,actor_name*,rule_suite_result*,per_page*,page*}", pathParameters),
    }
    return m
}
// NewItemRulesetsRuleSuitesRequestBuilder instantiates a new RuleSuitesRequestBuilder and sets the default values.
func NewItemRulesetsRuleSuitesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemRulesetsRuleSuitesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemRulesetsRuleSuitesRequestBuilderInternal(urlParams, requestAdapter)
}
// Get lists suites of rule evaluations at the organization level.For more information, see "[Managing rulesets for repositories in your organization](https://docs.github.com/organizations/managing-organization-settings/managing-rulesets-for-repositories-in-your-organization#viewing-insights-for-rulesets)."
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/orgs/rule-suites#list-organization-rule-suites
func (m *ItemRulesetsRuleSuitesRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemRulesetsRuleSuitesRequestBuilderGetRequestConfiguration)([]i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.RuleSuitesable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "404": i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CreateBasicErrorFromDiscriminatorValue,
        "500": i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CreateBasicErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.SendCollection(ctx, requestInfo, i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CreateRuleSuitesFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    val := make([]i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.RuleSuitesable, len(res))
    for i, v := range res {
        if v != nil {
            val[i] = v.(i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.RuleSuitesable)
        }
    }
    return val, nil
}
// ToGetRequestInformation lists suites of rule evaluations at the organization level.For more information, see "[Managing rulesets for repositories in your organization](https://docs.github.com/organizations/managing-organization-settings/managing-rulesets-for-repositories-in-your-organization#viewing-insights-for-rulesets)."
func (m *ItemRulesetsRuleSuitesRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemRulesetsRuleSuitesRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
func (m *ItemRulesetsRuleSuitesRequestBuilder) WithUrl(rawUrl string)(*ItemRulesetsRuleSuitesRequestBuilder) {
    return NewItemRulesetsRuleSuitesRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
