package rate_limit

import (
    "context"
    i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e "github.com/octokit/go-sdk/github/octokit/models"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// Rate_limitRequestBuilder builds and executes requests for operations under \rate_limit
type Rate_limitRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// Rate_limitRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type Rate_limitRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewRate_limitRequestBuilderInternal instantiates a new Rate_limitRequestBuilder and sets the default values.
func NewRate_limitRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*Rate_limitRequestBuilder) {
    m := &Rate_limitRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/rate_limit", pathParameters),
    }
    return m
}
// NewRate_limitRequestBuilder instantiates a new Rate_limitRequestBuilder and sets the default values.
func NewRate_limitRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*Rate_limitRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewRate_limitRequestBuilderInternal(urlParams, requestAdapter)
}
// Get **Note:** Accessing this endpoint does not count against your REST API rate limit.Some categories of endpoints have custom rate limits that are separate from the rate limit governing the other REST API endpoints. For this reason, the API response categorizes your rate limit. Under `resources`, you'll see objects relating to different categories:* The `core` object provides your rate limit status for all non-search-related resources in the REST API.* The `search` object provides your rate limit status for the REST API for searching (excluding code searches). For more information, see "[Search](https://docs.github.com/rest/search)."* The `code_search` object provides your rate limit status for the REST API for searching code. For more information, see "[Search code](https://docs.github.com/rest/search/search#search-code)."* The `graphql` object provides your rate limit status for the GraphQL API. For more information, see "[Resource limitations](https://docs.github.com/graphql/overview/resource-limitations#rate-limit)."* The `integration_manifest` object provides your rate limit status for the `POST /app-manifests/{code}/conversions` operation. For more information, see "[Creating a GitHub App from a manifest](https://docs.github.com/apps/creating-github-apps/setting-up-a-github-app/creating-a-github-app-from-a-manifest#3-you-exchange-the-temporary-code-to-retrieve-the-app-configuration)."* The `dependency_snapshots` object provides your rate limit status for submitting snapshots to the dependency graph. For more information, see "[Dependency graph](https://docs.github.com/rest/dependency-graph)."* The `code_scanning_upload` object provides your rate limit status for uploading SARIF results to code scanning. For more information, see "[Uploading a SARIF file to GitHub](https://docs.github.com/code-security/code-scanning/integrating-with-code-scanning/uploading-a-sarif-file-to-github)."* The `actions_runner_registration` object provides your rate limit status for registering self-hosted runners in GitHub Actions. For more information, see "[Self-hosted runners](https://docs.github.com/rest/actions/self-hosted-runners)."* The `source_import` object is no longer in use for any API endpoints, and it will be removed in the next API version. For more information about API versions, see "[API Versions](https://docs.github.com/rest/overview/api-versions)."**Note:** The `rate` object is deprecated. If you're writing new API client code or updating existing code, you should use the `core` object instead of the `rate` object. The `core` object contains the same information that is present in the `rate` object.
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/rate-limit/rate-limit#get-rate-limit-status-for-the-authenticated-user
func (m *Rate_limitRequestBuilder) Get(ctx context.Context, requestConfiguration *Rate_limitRequestBuilderGetRequestConfiguration)(i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.RateLimitOverviewable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "404": i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.CreateBasicErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.CreateRateLimitOverviewFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.RateLimitOverviewable), nil
}
// ToGetRequestInformation **Note:** Accessing this endpoint does not count against your REST API rate limit.Some categories of endpoints have custom rate limits that are separate from the rate limit governing the other REST API endpoints. For this reason, the API response categorizes your rate limit. Under `resources`, you'll see objects relating to different categories:* The `core` object provides your rate limit status for all non-search-related resources in the REST API.* The `search` object provides your rate limit status for the REST API for searching (excluding code searches). For more information, see "[Search](https://docs.github.com/rest/search)."* The `code_search` object provides your rate limit status for the REST API for searching code. For more information, see "[Search code](https://docs.github.com/rest/search/search#search-code)."* The `graphql` object provides your rate limit status for the GraphQL API. For more information, see "[Resource limitations](https://docs.github.com/graphql/overview/resource-limitations#rate-limit)."* The `integration_manifest` object provides your rate limit status for the `POST /app-manifests/{code}/conversions` operation. For more information, see "[Creating a GitHub App from a manifest](https://docs.github.com/apps/creating-github-apps/setting-up-a-github-app/creating-a-github-app-from-a-manifest#3-you-exchange-the-temporary-code-to-retrieve-the-app-configuration)."* The `dependency_snapshots` object provides your rate limit status for submitting snapshots to the dependency graph. For more information, see "[Dependency graph](https://docs.github.com/rest/dependency-graph)."* The `code_scanning_upload` object provides your rate limit status for uploading SARIF results to code scanning. For more information, see "[Uploading a SARIF file to GitHub](https://docs.github.com/code-security/code-scanning/integrating-with-code-scanning/uploading-a-sarif-file-to-github)."* The `actions_runner_registration` object provides your rate limit status for registering self-hosted runners in GitHub Actions. For more information, see "[Self-hosted runners](https://docs.github.com/rest/actions/self-hosted-runners)."* The `source_import` object is no longer in use for any API endpoints, and it will be removed in the next API version. For more information about API versions, see "[API Versions](https://docs.github.com/rest/overview/api-versions)."**Note:** The `rate` object is deprecated. If you're writing new API client code or updating existing code, you should use the `core` object instead of the `rate` object. The `core` object contains the same information that is present in the `rate` object.
func (m *Rate_limitRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *Rate_limitRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
func (m *Rate_limitRequestBuilder) WithUrl(rawUrl string)(*Rate_limitRequestBuilder) {
    return NewRate_limitRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
