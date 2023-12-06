package orgs

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemActionsCacheUsageByRepositoryRequestBuilder builds and executes requests for operations under \orgs\{org}\actions\cache\usage-by-repository
type ItemActionsCacheUsageByRepositoryRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemActionsCacheUsageByRepositoryRequestBuilderGetQueryParameters lists repositories and their GitHub Actions cache usage for an organization.The data fetched using this API is refreshed approximately every 5 minutes, so values returned from this endpoint may take at least 5 minutes to get updated.You must authenticate using an access token with the `read:org` scope to use this endpoint. GitHub Apps must have the `organization_admistration:read` permission to use this endpoint.
type ItemActionsCacheUsageByRepositoryRequestBuilderGetQueryParameters struct {
    // Page number of the results to fetch.
    Page *int32 `uriparametername:"page"`
    // The number of results per page (max 100).
    Per_page *int32 `uriparametername:"per_page"`
}
// ItemActionsCacheUsageByRepositoryRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemActionsCacheUsageByRepositoryRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemActionsCacheUsageByRepositoryRequestBuilderGetQueryParameters
}
// NewItemActionsCacheUsageByRepositoryRequestBuilderInternal instantiates a new UsageByRepositoryRequestBuilder and sets the default values.
func NewItemActionsCacheUsageByRepositoryRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemActionsCacheUsageByRepositoryRequestBuilder) {
    m := &ItemActionsCacheUsageByRepositoryRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/orgs/{org}/actions/cache/usage-by-repository{?per_page*,page*}", pathParameters),
    }
    return m
}
// NewItemActionsCacheUsageByRepositoryRequestBuilder instantiates a new UsageByRepositoryRequestBuilder and sets the default values.
func NewItemActionsCacheUsageByRepositoryRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemActionsCacheUsageByRepositoryRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemActionsCacheUsageByRepositoryRequestBuilderInternal(urlParams, requestAdapter)
}
// Get lists repositories and their GitHub Actions cache usage for an organization.The data fetched using this API is refreshed approximately every 5 minutes, so values returned from this endpoint may take at least 5 minutes to get updated.You must authenticate using an access token with the `read:org` scope to use this endpoint. GitHub Apps must have the `organization_admistration:read` permission to use this endpoint.
// Deprecated: This method is obsolete. Use GetAsUsageByRepositoryGetResponse instead.
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/actions/cache#list-repositories-with-github-actions-cache-usage-for-an-organization
func (m *ItemActionsCacheUsageByRepositoryRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemActionsCacheUsageByRepositoryRequestBuilderGetRequestConfiguration)(ItemActionsCacheUsageByRepositoryResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateItemActionsCacheUsageByRepositoryResponseFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ItemActionsCacheUsageByRepositoryResponseable), nil
}
// GetAsUsageByRepositoryGetResponse lists repositories and their GitHub Actions cache usage for an organization.The data fetched using this API is refreshed approximately every 5 minutes, so values returned from this endpoint may take at least 5 minutes to get updated.You must authenticate using an access token with the `read:org` scope to use this endpoint. GitHub Apps must have the `organization_admistration:read` permission to use this endpoint.
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/actions/cache#list-repositories-with-github-actions-cache-usage-for-an-organization
func (m *ItemActionsCacheUsageByRepositoryRequestBuilder) GetAsUsageByRepositoryGetResponse(ctx context.Context, requestConfiguration *ItemActionsCacheUsageByRepositoryRequestBuilderGetRequestConfiguration)(ItemActionsCacheUsageByRepositoryGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateItemActionsCacheUsageByRepositoryGetResponseFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ItemActionsCacheUsageByRepositoryGetResponseable), nil
}
// ToGetRequestInformation lists repositories and their GitHub Actions cache usage for an organization.The data fetched using this API is refreshed approximately every 5 minutes, so values returned from this endpoint may take at least 5 minutes to get updated.You must authenticate using an access token with the `read:org` scope to use this endpoint. GitHub Apps must have the `organization_admistration:read` permission to use this endpoint.
func (m *ItemActionsCacheUsageByRepositoryRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemActionsCacheUsageByRepositoryRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *ItemActionsCacheUsageByRepositoryRequestBuilder) WithUrl(rawUrl string)(*ItemActionsCacheUsageByRepositoryRequestBuilder) {
    return NewItemActionsCacheUsageByRepositoryRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
