package repos

import (
    "context"
    i53ac87e8cb3cc9276228f74d38694a208cacb99bb8ceb705eeae99fb88d4d274 "strconv"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i000736ae6dd74f01081193e4f903216bc2bd2954ed818433b986f45d581ed035 "github.com/octokit/go-sdk/github/octokit/models"
    i235e3897503c37ae92ef8cfa89cc51117cff3968f38c100b13fa42113c5c2392 "github.com/octokit/go-sdk/github/octokit/repos/item/item/codescanning/analyses"
)

// ItemItemCodeScanningAnalysesRequestBuilder builds and executes requests for operations under \repos\{repos-id}\{Owner-id}\code-scanning\analyses
type ItemItemCodeScanningAnalysesRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemItemCodeScanningAnalysesRequestBuilderGetQueryParameters lists the details of all code scanning analyses for a repository,starting with the most recent.The response is paginated and you can use the `page` and `per_page` parametersto list the analyses you're interested in.By default 30 analyses are listed per page.The `rules_count` field in the response give the number of rulesthat were run in the analysis.For very old analyses this data is not available,and `0` is returned in this field.You must use an access token with the `security_events` scope to use this endpoint with private repos,the `public_repo` scope also grants permission to read security events on public repos only.**Deprecation notice**:The `tool_name` field is deprecated and will, in future, not be included in the response for this endpoint. The example response reflects this change. The tool name can now be found inside the `tool` field.
type ItemItemCodeScanningAnalysesRequestBuilderGetQueryParameters struct {
    // The direction to sort the results by.
    // Deprecated: This property is deprecated, use directionAsGetDirectionQueryParameterType instead
    Direction *string `uriparametername:"direction"`
    // The direction to sort the results by.
    DirectionAsGetDirectionQueryParameterType *i235e3897503c37ae92ef8cfa89cc51117cff3968f38c100b13fa42113c5c2392.GetDirectionQueryParameterType `uriparametername:"direction"`
    // Page number of the results to fetch.
    Page *int32 `uriparametername:"page"`
    // The number of results per page (max 100).
    Per_page *int32 `uriparametername:"per_page"`
    // The Git reference for the analyses you want to list. The `ref` for a branch can be formatted either as `refs/heads/<branch name>` or simply `<branch name>`. To reference a pull request use `refs/pull/<number>/merge`.
    Ref *string `uriparametername:"ref"`
    // Filter analyses belonging to the same SARIF upload.
    Sarif_id *string `uriparametername:"sarif_id"`
    // The property by which to sort the results.
    // Deprecated: This property is deprecated, use sortAsGetSortQueryParameterType instead
    Sort *string `uriparametername:"sort"`
    // The property by which to sort the results.
    SortAsGetSortQueryParameterType *i235e3897503c37ae92ef8cfa89cc51117cff3968f38c100b13fa42113c5c2392.GetSortQueryParameterType `uriparametername:"sort"`
    // The GUID of a code scanning tool. Only results by this tool will be listed. Note that some code scanning tools may not include a GUID in their analysis data. You can specify the tool by using either `tool_guid` or `tool_name`, but not both.
    Tool_guid *string `uriparametername:"tool_guid"`
    // The name of a code scanning tool. Only results by this tool will be listed. You can specify the tool by using either `tool_name` or `tool_guid`, but not both.
    Tool_name *string `uriparametername:"tool_name"`
}
// ItemItemCodeScanningAnalysesRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemItemCodeScanningAnalysesRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemItemCodeScanningAnalysesRequestBuilderGetQueryParameters
}
// ByAnalysis_id gets an item from the github.com/octokit/go-sdk/github/octokit/.repos.item.item.codeScanning.analyses.item collection
// Deprecated: This indexer is deprecated and will be removed in the next major version. Use the one with the typed parameter instead.
func (m *ItemItemCodeScanningAnalysesRequestBuilder) ByAnalysis_id(analysis_id string)(*ItemItemCodeScanningAnalysesWithAnalysis_ItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if analysis_id != "" {
        urlTplParams["analysis_id"] = analysis_id
    }
    return NewItemItemCodeScanningAnalysesWithAnalysis_ItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// ByAnalysis_idInteger gets an item from the github.com/octokit/go-sdk/github/octokit/.repos.item.item.codeScanning.analyses.item collection
func (m *ItemItemCodeScanningAnalysesRequestBuilder) ByAnalysis_idInteger(analysis_id int32)(*ItemItemCodeScanningAnalysesWithAnalysis_ItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    urlTplParams["analysis_id"] = i53ac87e8cb3cc9276228f74d38694a208cacb99bb8ceb705eeae99fb88d4d274.FormatInt(int64(analysis_id), 10)
    return NewItemItemCodeScanningAnalysesWithAnalysis_ItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewItemItemCodeScanningAnalysesRequestBuilderInternal instantiates a new AnalysesRequestBuilder and sets the default values.
func NewItemItemCodeScanningAnalysesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemCodeScanningAnalysesRequestBuilder) {
    m := &ItemItemCodeScanningAnalysesRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/repos/{repos%2Did}/{Owner%2Did}/code-scanning/analyses{?tool_name*,tool_guid*,page*,per_page*,ref*,sarif_id*,direction*,sort*}", pathParameters),
    }
    return m
}
// NewItemItemCodeScanningAnalysesRequestBuilder instantiates a new AnalysesRequestBuilder and sets the default values.
func NewItemItemCodeScanningAnalysesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemCodeScanningAnalysesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemItemCodeScanningAnalysesRequestBuilderInternal(urlParams, requestAdapter)
}
// Get lists the details of all code scanning analyses for a repository,starting with the most recent.The response is paginated and you can use the `page` and `per_page` parametersto list the analyses you're interested in.By default 30 analyses are listed per page.The `rules_count` field in the response give the number of rulesthat were run in the analysis.For very old analyses this data is not available,and `0` is returned in this field.You must use an access token with the `security_events` scope to use this endpoint with private repos,the `public_repo` scope also grants permission to read security events on public repos only.**Deprecation notice**:The `tool_name` field is deprecated and will, in future, not be included in the response for this endpoint. The example response reflects this change. The tool name can now be found inside the `tool` field.
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/code-scanning/code-scanning#list-code-scanning-analyses-for-a-repository
func (m *ItemItemCodeScanningAnalysesRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemItemCodeScanningAnalysesRequestBuilderGetRequestConfiguration)([]i000736ae6dd74f01081193e4f903216bc2bd2954ed818433b986f45d581ed035.CodeScanningAnalysisable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "403": i000736ae6dd74f01081193e4f903216bc2bd2954ed818433b986f45d581ed035.CreateBasicErrorFromDiscriminatorValue,
        "404": i000736ae6dd74f01081193e4f903216bc2bd2954ed818433b986f45d581ed035.CreateBasicErrorFromDiscriminatorValue,
        "503": i000736ae6dd74f01081193e4f903216bc2bd2954ed818433b986f45d581ed035.CreateAnalyses503ErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.SendCollection(ctx, requestInfo, i000736ae6dd74f01081193e4f903216bc2bd2954ed818433b986f45d581ed035.CreateCodeScanningAnalysisFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    val := make([]i000736ae6dd74f01081193e4f903216bc2bd2954ed818433b986f45d581ed035.CodeScanningAnalysisable, len(res))
    for i, v := range res {
        if v != nil {
            val[i] = v.(i000736ae6dd74f01081193e4f903216bc2bd2954ed818433b986f45d581ed035.CodeScanningAnalysisable)
        }
    }
    return val, nil
}
// ToGetRequestInformation lists the details of all code scanning analyses for a repository,starting with the most recent.The response is paginated and you can use the `page` and `per_page` parametersto list the analyses you're interested in.By default 30 analyses are listed per page.The `rules_count` field in the response give the number of rulesthat were run in the analysis.For very old analyses this data is not available,and `0` is returned in this field.You must use an access token with the `security_events` scope to use this endpoint with private repos,the `public_repo` scope also grants permission to read security events on public repos only.**Deprecation notice**:The `tool_name` field is deprecated and will, in future, not be included in the response for this endpoint. The example response reflects this change. The tool name can now be found inside the `tool` field.
func (m *ItemItemCodeScanningAnalysesRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemItemCodeScanningAnalysesRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *ItemItemCodeScanningAnalysesRequestBuilder) WithUrl(rawUrl string)(*ItemItemCodeScanningAnalysesRequestBuilder) {
    return NewItemItemCodeScanningAnalysesRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
