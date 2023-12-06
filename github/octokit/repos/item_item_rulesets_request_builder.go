package repos

import (
    "context"
    i53ac87e8cb3cc9276228f74d38694a208cacb99bb8ceb705eeae99fb88d4d274 "strconv"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i000736ae6dd74f01081193e4f903216bc2bd2954ed818433b986f45d581ed035 "github.com/octokit/go-sdk/github/octokit/models"
)

// ItemItemRulesetsRequestBuilder builds and executes requests for operations under \repos\{repos-id}\{Owner-id}\rulesets
type ItemItemRulesetsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemItemRulesetsRequestBuilderGetQueryParameters get all the rulesets for a repository.
type ItemItemRulesetsRequestBuilderGetQueryParameters struct {
    // Include rulesets configured at higher levels that apply to this repository
    Includes_parents *bool `uriparametername:"includes_parents"`
    // Page number of the results to fetch.
    Page *int32 `uriparametername:"page"`
    // The number of results per page (max 100).
    Per_page *int32 `uriparametername:"per_page"`
}
// ItemItemRulesetsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemItemRulesetsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemItemRulesetsRequestBuilderGetQueryParameters
}
// ItemItemRulesetsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemItemRulesetsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ByRuleset_id gets an item from the github.com/octokit/go-sdk/github/octokit/.repos.item.item.rulesets.item collection
// Deprecated: This indexer is deprecated and will be removed in the next major version. Use the one with the typed parameter instead.
func (m *ItemItemRulesetsRequestBuilder) ByRuleset_id(ruleset_id string)(*ItemItemRulesetsWithRuleset_ItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if ruleset_id != "" {
        urlTplParams["ruleset_id"] = ruleset_id
    }
    return NewItemItemRulesetsWithRuleset_ItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// ByRuleset_idInteger gets an item from the github.com/octokit/go-sdk/github/octokit/.repos.item.item.rulesets.item collection
func (m *ItemItemRulesetsRequestBuilder) ByRuleset_idInteger(ruleset_id int32)(*ItemItemRulesetsWithRuleset_ItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    urlTplParams["ruleset_id"] = i53ac87e8cb3cc9276228f74d38694a208cacb99bb8ceb705eeae99fb88d4d274.FormatInt(int64(ruleset_id), 10)
    return NewItemItemRulesetsWithRuleset_ItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewItemItemRulesetsRequestBuilderInternal instantiates a new RulesetsRequestBuilder and sets the default values.
func NewItemItemRulesetsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemRulesetsRequestBuilder) {
    m := &ItemItemRulesetsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/repos/{repos%2Did}/{Owner%2Did}/rulesets{?per_page*,page*,includes_parents*}", pathParameters),
    }
    return m
}
// NewItemItemRulesetsRequestBuilder instantiates a new RulesetsRequestBuilder and sets the default values.
func NewItemItemRulesetsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemRulesetsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemItemRulesetsRequestBuilderInternal(urlParams, requestAdapter)
}
// Get get all the rulesets for a repository.
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/repos/rules#get-all-repository-rulesets
func (m *ItemItemRulesetsRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemItemRulesetsRequestBuilderGetRequestConfiguration)([]i000736ae6dd74f01081193e4f903216bc2bd2954ed818433b986f45d581ed035.RepositoryRulesetable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "404": i000736ae6dd74f01081193e4f903216bc2bd2954ed818433b986f45d581ed035.CreateBasicErrorFromDiscriminatorValue,
        "500": i000736ae6dd74f01081193e4f903216bc2bd2954ed818433b986f45d581ed035.CreateBasicErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.SendCollection(ctx, requestInfo, i000736ae6dd74f01081193e4f903216bc2bd2954ed818433b986f45d581ed035.CreateRepositoryRulesetFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    val := make([]i000736ae6dd74f01081193e4f903216bc2bd2954ed818433b986f45d581ed035.RepositoryRulesetable, len(res))
    for i, v := range res {
        if v != nil {
            val[i] = v.(i000736ae6dd74f01081193e4f903216bc2bd2954ed818433b986f45d581ed035.RepositoryRulesetable)
        }
    }
    return val, nil
}
// Post create a ruleset for a repository.
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/repos/rules#create-a-repository-ruleset
func (m *ItemItemRulesetsRequestBuilder) Post(ctx context.Context, body ItemItemRulesetsPostRequestBodyable, requestConfiguration *ItemItemRulesetsRequestBuilderPostRequestConfiguration)(i000736ae6dd74f01081193e4f903216bc2bd2954ed818433b986f45d581ed035.RepositoryRulesetable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "404": i000736ae6dd74f01081193e4f903216bc2bd2954ed818433b986f45d581ed035.CreateBasicErrorFromDiscriminatorValue,
        "500": i000736ae6dd74f01081193e4f903216bc2bd2954ed818433b986f45d581ed035.CreateBasicErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i000736ae6dd74f01081193e4f903216bc2bd2954ed818433b986f45d581ed035.CreateRepositoryRulesetFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i000736ae6dd74f01081193e4f903216bc2bd2954ed818433b986f45d581ed035.RepositoryRulesetable), nil
}
// RuleSuites the ruleSuites property
func (m *ItemItemRulesetsRequestBuilder) RuleSuites()(*ItemItemRulesetsRuleSuitesRequestBuilder) {
    return NewItemItemRulesetsRuleSuitesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToGetRequestInformation get all the rulesets for a repository.
func (m *ItemItemRulesetsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemItemRulesetsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation create a ruleset for a repository.
func (m *ItemItemRulesetsRequestBuilder) ToPostRequestInformation(ctx context.Context, body ItemItemRulesetsPostRequestBodyable, requestConfiguration *ItemItemRulesetsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
func (m *ItemItemRulesetsRequestBuilder) WithUrl(rawUrl string)(*ItemItemRulesetsRequestBuilder) {
    return NewItemItemRulesetsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
