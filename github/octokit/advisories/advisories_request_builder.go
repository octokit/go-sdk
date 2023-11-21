package advisories

import (
    "context"
    i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e "github.com/octokit/go-sdk/github/octokit/models"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// AdvisoriesRequestBuilder builds and executes requests for operations under \advisories
type AdvisoriesRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// AdvisoriesRequestBuilderGetQueryParameters lists all global security advisories that match the specified parameters. If no other parameters are defined, the request will return only GitHub-reviewed advisories that are not malware.By default, all responses will exclude advisories for malware, because malware are not standard vulnerabilities. To list advisories for malware, you must include the `type` parameter in your request, with the value `malware`. For more information about the different types of security advisories, see "[About the GitHub Advisory database](https://docs.github.com/code-security/security-advisories/global-security-advisories/about-the-github-advisory-database#about-types-of-security-advisories)."
type AdvisoriesRequestBuilderGetQueryParameters struct {
    // If specified, only return advisories that affect any of `package` or `package@version`. A maximum of 1000 packages can be specified.If the query parameter causes the URL to exceed the maximum URL length supported by your client, you must specify fewer packages.Example: `affects=package1,package2@1.0.0,package3@^2.0.0` or `affects[]=package1&affects[]=package2@1.0.0`
    Affects *string `uriparametername:"affects"`
    // A cursor, as given in the [Link header](https://docs.github.com/rest/guides/using-pagination-in-the-rest-api#using-link-headers). If specified, the query only searches for results after this cursor.
    After *string `uriparametername:"after"`
    // A cursor, as given in the [Link header](https://docs.github.com/rest/guides/using-pagination-in-the-rest-api#using-link-headers). If specified, the query only searches for results before this cursor.
    Before *string `uriparametername:"before"`
    // If specified, only advisories with this CVE (Common Vulnerabilities and Exposures) identifier will be returned.
    Cve_id *string `uriparametername:"cve_id"`
    // If specified, only advisories with these Common Weakness Enumerations (CWEs) will be returned.Example: `cwes=79,284,22` or `cwes[]=79&cwes[]=284&cwes[]=22`
    Cwes *string `uriparametername:"cwes"`
    // The direction to sort the results by.
    // Deprecated: This property is deprecated, use directionAsGetDirectionQueryParameterType instead
    Direction *string `uriparametername:"direction"`
    // The direction to sort the results by.
    DirectionAsGetDirectionQueryParameterType *GetDirectionQueryParameterType `uriparametername:"direction"`
    // If specified, only advisories for these ecosystems will be returned.
    // Deprecated: This property is deprecated, use ecosystemAsGetEcosystemQueryParameterType instead
    Ecosystem *string `uriparametername:"ecosystem"`
    // If specified, only advisories for these ecosystems will be returned.
    EcosystemAsGetEcosystemQueryParameterType *GetEcosystemQueryParameterType `uriparametername:"ecosystem"`
    // If specified, only advisories with this GHSA (GitHub Security Advisory) identifier will be returned.
    Ghsa_id *string `uriparametername:"ghsa_id"`
    // Whether to only return advisories that have been withdrawn.
    Is_withdrawn *bool `uriparametername:"is_withdrawn"`
    // If specified, only show advisories that were updated or published on a date or date range.For more information on the syntax of the date range, see "[Understanding the search syntax](https://docs.github.com/search-github/getting-started-with-searching-on-github/understanding-the-search-syntax#query-for-dates)."
    Modified *string `uriparametername:"modified"`
    // The number of results per page (max 100).
    Per_page *int32 `uriparametername:"per_page"`
    // If specified, only return advisories that were published on a date or date range.For more information on the syntax of the date range, see "[Understanding the search syntax](https://docs.github.com/search-github/getting-started-with-searching-on-github/understanding-the-search-syntax#query-for-dates)."
    Published *string `uriparametername:"published"`
    // If specified, only advisories with these severities will be returned.
    // Deprecated: This property is deprecated, use severityAsGetSeverityQueryParameterType instead
    Severity *string `uriparametername:"severity"`
    // If specified, only advisories with these severities will be returned.
    SeverityAsGetSeverityQueryParameterType *GetSeverityQueryParameterType `uriparametername:"severity"`
    // The property to sort the results by.
    // Deprecated: This property is deprecated, use sortAsGetSortQueryParameterType instead
    Sort *string `uriparametername:"sort"`
    // The property to sort the results by.
    SortAsGetSortQueryParameterType *GetSortQueryParameterType `uriparametername:"sort"`
    // If specified, only advisories of this type will be returned. By default, a request with no other parameters defined will only return reviewed advisories that are not malware.
    // Deprecated: This property is deprecated, use typeAsGetTypeQueryParameterType instead
    Type *string `uriparametername:"type"`
    // If specified, only advisories of this type will be returned. By default, a request with no other parameters defined will only return reviewed advisories that are not malware.
    TypeAsGetTypeQueryParameterType *GetTypeQueryParameterType `uriparametername:"type"`
    // If specified, only return advisories that were updated on a date or date range.For more information on the syntax of the date range, see "[Understanding the search syntax](https://docs.github.com/search-github/getting-started-with-searching-on-github/understanding-the-search-syntax#query-for-dates)."
    Updated *string `uriparametername:"updated"`
}
// AdvisoriesRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AdvisoriesRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *AdvisoriesRequestBuilderGetQueryParameters
}
// ByGhsa_id gets an item from the octokit.advisories.item collection
func (m *AdvisoriesRequestBuilder) ByGhsa_id(ghsa_id string)(*WithGhsa_ItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if ghsa_id != "" {
        urlTplParams["ghsa_id"] = ghsa_id
    }
    return NewWithGhsa_ItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewAdvisoriesRequestBuilderInternal instantiates a new AdvisoriesRequestBuilder and sets the default values.
func NewAdvisoriesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AdvisoriesRequestBuilder) {
    m := &AdvisoriesRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/advisories{?ghsa_id*,type*,cve_id*,ecosystem*,severity*,cwes*,is_withdrawn*,affects*,published*,updated*,modified*,before*,after*,direction*,per_page*,sort*}", pathParameters),
    }
    return m
}
// NewAdvisoriesRequestBuilder instantiates a new AdvisoriesRequestBuilder and sets the default values.
func NewAdvisoriesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AdvisoriesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewAdvisoriesRequestBuilderInternal(urlParams, requestAdapter)
}
// Get lists all global security advisories that match the specified parameters. If no other parameters are defined, the request will return only GitHub-reviewed advisories that are not malware.By default, all responses will exclude advisories for malware, because malware are not standard vulnerabilities. To list advisories for malware, you must include the `type` parameter in your request, with the value `malware`. For more information about the different types of security advisories, see "[About the GitHub Advisory database](https://docs.github.com/code-security/security-advisories/global-security-advisories/about-the-github-advisory-database#about-types-of-security-advisories)."
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/security-advisories/global-advisories#list-global-security-advisories
func (m *AdvisoriesRequestBuilder) Get(ctx context.Context, requestConfiguration *AdvisoriesRequestBuilderGetRequestConfiguration)([]i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.GlobalAdvisoryable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "422": i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.CreateValidationErrorSimpleFromDiscriminatorValue,
        "429": i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.CreateBasicErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.SendCollection(ctx, requestInfo, i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.CreateGlobalAdvisoryFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    val := make([]i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.GlobalAdvisoryable, len(res))
    for i, v := range res {
        if v != nil {
            val[i] = v.(i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.GlobalAdvisoryable)
        }
    }
    return val, nil
}
// ToGetRequestInformation lists all global security advisories that match the specified parameters. If no other parameters are defined, the request will return only GitHub-reviewed advisories that are not malware.By default, all responses will exclude advisories for malware, because malware are not standard vulnerabilities. To list advisories for malware, you must include the `type` parameter in your request, with the value `malware`. For more information about the different types of security advisories, see "[About the GitHub Advisory database](https://docs.github.com/code-security/security-advisories/global-security-advisories/about-the-github-advisory-database#about-types-of-security-advisories)."
func (m *AdvisoriesRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *AdvisoriesRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
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
func (m *AdvisoriesRequestBuilder) WithUrl(rawUrl string)(*AdvisoriesRequestBuilder) {
    return NewAdvisoriesRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
