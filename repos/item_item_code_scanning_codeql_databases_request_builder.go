package repos

import (
    "context"
    i33bcfa37677dfd8b71eaf22ad3bfd140ba1702e7a8418558499c22418bf2fef2 "github.com/octokit/go-sdk/models"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemItemCodeScanningCodeqlDatabasesRequestBuilder builds and executes requests for operations under \repos\{owner}\{repo}\code-scanning\codeql\databases
type ItemItemCodeScanningCodeqlDatabasesRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemItemCodeScanningCodeqlDatabasesRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemItemCodeScanningCodeqlDatabasesRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ByLanguage gets an item from the go-sdk.repos.item.item.codeScanning.codeql.databases.item collection
func (m *ItemItemCodeScanningCodeqlDatabasesRequestBuilder) ByLanguage(language string)(*ItemItemCodeScanningCodeqlDatabasesWithLanguageItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if language != "" {
        urlTplParams["language"] = language
    }
    return NewItemItemCodeScanningCodeqlDatabasesWithLanguageItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewItemItemCodeScanningCodeqlDatabasesRequestBuilderInternal instantiates a new DatabasesRequestBuilder and sets the default values.
func NewItemItemCodeScanningCodeqlDatabasesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemCodeScanningCodeqlDatabasesRequestBuilder) {
    m := &ItemItemCodeScanningCodeqlDatabasesRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/repos/{owner}/{repo}/code-scanning/codeql/databases", pathParameters),
    }
    return m
}
// NewItemItemCodeScanningCodeqlDatabasesRequestBuilder instantiates a new DatabasesRequestBuilder and sets the default values.
func NewItemItemCodeScanningCodeqlDatabasesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemCodeScanningCodeqlDatabasesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemItemCodeScanningCodeqlDatabasesRequestBuilderInternal(urlParams, requestAdapter)
}
// Get lists the CodeQL databases that are available in a repository.For private repositories, you must use an access token with the `security_events` scope.For public repositories, you can use tokens with the `security_events` or `public_repo` scope.
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/code-scanning/code-scanning#list-codeql-databases-for-a-repository
func (m *ItemItemCodeScanningCodeqlDatabasesRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemItemCodeScanningCodeqlDatabasesRequestBuilderGetRequestConfiguration)([]i33bcfa37677dfd8b71eaf22ad3bfd140ba1702e7a8418558499c22418bf2fef2.CodeScanningCodeqlDatabaseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "403": i33bcfa37677dfd8b71eaf22ad3bfd140ba1702e7a8418558499c22418bf2fef2.CreateBasicErrorFromDiscriminatorValue,
        "404": i33bcfa37677dfd8b71eaf22ad3bfd140ba1702e7a8418558499c22418bf2fef2.CreateBasicErrorFromDiscriminatorValue,
        "503": i33bcfa37677dfd8b71eaf22ad3bfd140ba1702e7a8418558499c22418bf2fef2.CreateDatabases503ErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.SendCollection(ctx, requestInfo, i33bcfa37677dfd8b71eaf22ad3bfd140ba1702e7a8418558499c22418bf2fef2.CreateCodeScanningCodeqlDatabaseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    val := make([]i33bcfa37677dfd8b71eaf22ad3bfd140ba1702e7a8418558499c22418bf2fef2.CodeScanningCodeqlDatabaseable, len(res))
    for i, v := range res {
        if v != nil {
            val[i] = v.(i33bcfa37677dfd8b71eaf22ad3bfd140ba1702e7a8418558499c22418bf2fef2.CodeScanningCodeqlDatabaseable)
        }
    }
    return val, nil
}
// ToGetRequestInformation lists the CodeQL databases that are available in a repository.For private repositories, you must use an access token with the `security_events` scope.For public repositories, you can use tokens with the `security_events` or `public_repo` scope.
func (m *ItemItemCodeScanningCodeqlDatabasesRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemItemCodeScanningCodeqlDatabasesRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *ItemItemCodeScanningCodeqlDatabasesRequestBuilder) WithUrl(rawUrl string)(*ItemItemCodeScanningCodeqlDatabasesRequestBuilder) {
    return NewItemItemCodeScanningCodeqlDatabasesRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
