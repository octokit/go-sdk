package orgs

import (
    "context"
    i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e "github.com/octokit/go-sdk/github/octokit/models"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemPropertiesValuesRequestBuilder builds and executes requests for operations under \orgs\{org}\properties\values
type ItemPropertiesValuesRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemPropertiesValuesRequestBuilderGetQueryParameters lists organization repositories with all of their custom property values.Organization members can read these properties.
type ItemPropertiesValuesRequestBuilderGetQueryParameters struct {
    // Page number of the results to fetch.
    Page *int32 `uriparametername:"page"`
    // The number of results per page (max 100).
    Per_page *int32 `uriparametername:"per_page"`
    // Finds repositories in the organization with a query containing one or more search keywords and qualifiers. Qualifiers allow you to limit your search to specific areas of GitHub. The REST API supports the same qualifiers as the web interface for GitHub. To learn more about the format of the query, see [Constructing a search query](https://docs.github.com/rest/search/search#constructing-a-search-query). See "[Searching for repositories](https://docs.github.com/articles/searching-for-repositories/)" for a detailed list of qualifiers.
    Repository_query *string `uriparametername:"repository_query"`
}
// ItemPropertiesValuesRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemPropertiesValuesRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemPropertiesValuesRequestBuilderGetQueryParameters
}
// ItemPropertiesValuesRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemPropertiesValuesRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemPropertiesValuesRequestBuilderInternal instantiates a new ValuesRequestBuilder and sets the default values.
func NewItemPropertiesValuesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemPropertiesValuesRequestBuilder) {
    m := &ItemPropertiesValuesRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/orgs/{org}/properties/values{?per_page*,page*,repository_query*}", pathParameters),
    }
    return m
}
// NewItemPropertiesValuesRequestBuilder instantiates a new ValuesRequestBuilder and sets the default values.
func NewItemPropertiesValuesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemPropertiesValuesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemPropertiesValuesRequestBuilderInternal(urlParams, requestAdapter)
}
// Get lists organization repositories with all of their custom property values.Organization members can read these properties.
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/orgs/custom-properties#list-custom-property-values-for-organization-repositories
func (m *ItemPropertiesValuesRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemPropertiesValuesRequestBuilderGetRequestConfiguration)([]i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.OrgRepoCustomPropertyValuesable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "403": i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.CreateBasicErrorFromDiscriminatorValue,
        "404": i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.CreateBasicErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.SendCollection(ctx, requestInfo, i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.CreateOrgRepoCustomPropertyValuesFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    val := make([]i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.OrgRepoCustomPropertyValuesable, len(res))
    for i, v := range res {
        if v != nil {
            val[i] = v.(i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.OrgRepoCustomPropertyValuesable)
        }
    }
    return val, nil
}
// Patch create new or update existing custom property values for repositories in a batch that belong to an organization.Each target repository will have its custom property values updated to match the values provided in the request.A maximum of 30 repositories can be updated in a single request.Using a value of `null` for a custom property will remove or 'unset' the property value from the repository.Only organization owners (or users with the proper permissions granted by them) can update these properties
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/orgs/custom-properties#create-or-update-custom-property-values-for-organization-repositories
func (m *ItemPropertiesValuesRequestBuilder) Patch(ctx context.Context, body ItemPropertiesValuesPatchRequestBodyable, requestConfiguration *ItemPropertiesValuesRequestBuilderPatchRequestConfiguration)(error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "403": i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.CreateBasicErrorFromDiscriminatorValue,
        "404": i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.CreateBasicErrorFromDiscriminatorValue,
    }
    err = m.BaseRequestBuilder.RequestAdapter.SendNoContent(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// ToGetRequestInformation lists organization repositories with all of their custom property values.Organization members can read these properties.
func (m *ItemPropertiesValuesRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemPropertiesValuesRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation create new or update existing custom property values for repositories in a batch that belong to an organization.Each target repository will have its custom property values updated to match the values provided in the request.A maximum of 30 repositories can be updated in a single request.Using a value of `null` for a custom property will remove or 'unset' the property value from the repository.Only organization owners (or users with the proper permissions granted by them) can update these properties
func (m *ItemPropertiesValuesRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ItemPropertiesValuesPatchRequestBodyable, requestConfiguration *ItemPropertiesValuesRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
    requestInfo.Headers.TryAdd("Accept", "application/json")
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
func (m *ItemPropertiesValuesRequestBuilder) WithUrl(rawUrl string)(*ItemPropertiesValuesRequestBuilder) {
    return NewItemPropertiesValuesRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
