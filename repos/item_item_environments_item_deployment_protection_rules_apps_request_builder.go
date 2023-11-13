package repos

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemItemEnvironmentsItemDeployment_protection_rulesAppsRequestBuilder builds and executes requests for operations under \repos\{owner}\{repo}\environments\{environment_name}\deployment_protection_rules\apps
type ItemItemEnvironmentsItemDeployment_protection_rulesAppsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemItemEnvironmentsItemDeployment_protection_rulesAppsRequestBuilderGetQueryParameters gets all custom deployment protection rule integrations that are available for an environment. Anyone with read access to the repository can use this endpoint. If the repository is private and you want to use a personal access token (classic), you must use an access token with the `repo` scope. GitHub Apps and fine-grained personal access tokens must have the `actions:read` permission to use this endpoint.For more information about environments, see "[Using environments for deployment](https://docs.github.com/en/actions/deployment/targeting-different-environments/using-environments-for-deployment)."For more information about the app that is providing this custom deployment rule, see "[GET an app](https://docs.github.com/rest/apps/apps#get-an-app)".
type ItemItemEnvironmentsItemDeployment_protection_rulesAppsRequestBuilderGetQueryParameters struct {
    // Page number of the results to fetch.
    Page *int32 `uriparametername:"page"`
    // The number of results per page (max 100).
    Per_page *int32 `uriparametername:"per_page"`
}
// ItemItemEnvironmentsItemDeployment_protection_rulesAppsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemItemEnvironmentsItemDeployment_protection_rulesAppsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemItemEnvironmentsItemDeployment_protection_rulesAppsRequestBuilderGetQueryParameters
}
// NewItemItemEnvironmentsItemDeployment_protection_rulesAppsRequestBuilderInternal instantiates a new AppsRequestBuilder and sets the default values.
func NewItemItemEnvironmentsItemDeployment_protection_rulesAppsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemEnvironmentsItemDeployment_protection_rulesAppsRequestBuilder) {
    m := &ItemItemEnvironmentsItemDeployment_protection_rulesAppsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/repos/{owner}/{repo}/environments/{environment_name}/deployment_protection_rules/apps{?page*,per_page*}", pathParameters),
    }
    return m
}
// NewItemItemEnvironmentsItemDeployment_protection_rulesAppsRequestBuilder instantiates a new AppsRequestBuilder and sets the default values.
func NewItemItemEnvironmentsItemDeployment_protection_rulesAppsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemEnvironmentsItemDeployment_protection_rulesAppsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemItemEnvironmentsItemDeployment_protection_rulesAppsRequestBuilderInternal(urlParams, requestAdapter)
}
// Get gets all custom deployment protection rule integrations that are available for an environment. Anyone with read access to the repository can use this endpoint. If the repository is private and you want to use a personal access token (classic), you must use an access token with the `repo` scope. GitHub Apps and fine-grained personal access tokens must have the `actions:read` permission to use this endpoint.For more information about environments, see "[Using environments for deployment](https://docs.github.com/en/actions/deployment/targeting-different-environments/using-environments-for-deployment)."For more information about the app that is providing this custom deployment rule, see "[GET an app](https://docs.github.com/rest/apps/apps#get-an-app)".
// Deprecated: This method is obsolete. Use GetAsAppsGetResponse instead.
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/deployments/protection-rules#list-custom-deployment-rule-integrations-available-for-an-environment
func (m *ItemItemEnvironmentsItemDeployment_protection_rulesAppsRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemItemEnvironmentsItemDeployment_protection_rulesAppsRequestBuilderGetRequestConfiguration)(ItemItemEnvironmentsItemDeployment_protection_rulesAppsResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateItemItemEnvironmentsItemDeployment_protection_rulesAppsResponseFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ItemItemEnvironmentsItemDeployment_protection_rulesAppsResponseable), nil
}
// GetAsAppsGetResponse gets all custom deployment protection rule integrations that are available for an environment. Anyone with read access to the repository can use this endpoint. If the repository is private and you want to use a personal access token (classic), you must use an access token with the `repo` scope. GitHub Apps and fine-grained personal access tokens must have the `actions:read` permission to use this endpoint.For more information about environments, see "[Using environments for deployment](https://docs.github.com/en/actions/deployment/targeting-different-environments/using-environments-for-deployment)."For more information about the app that is providing this custom deployment rule, see "[GET an app](https://docs.github.com/rest/apps/apps#get-an-app)".
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/deployments/protection-rules#list-custom-deployment-rule-integrations-available-for-an-environment
func (m *ItemItemEnvironmentsItemDeployment_protection_rulesAppsRequestBuilder) GetAsAppsGetResponse(ctx context.Context, requestConfiguration *ItemItemEnvironmentsItemDeployment_protection_rulesAppsRequestBuilderGetRequestConfiguration)(ItemItemEnvironmentsItemDeployment_protection_rulesAppsGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateItemItemEnvironmentsItemDeployment_protection_rulesAppsGetResponseFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ItemItemEnvironmentsItemDeployment_protection_rulesAppsGetResponseable), nil
}
// ToGetRequestInformation gets all custom deployment protection rule integrations that are available for an environment. Anyone with read access to the repository can use this endpoint. If the repository is private and you want to use a personal access token (classic), you must use an access token with the `repo` scope. GitHub Apps and fine-grained personal access tokens must have the `actions:read` permission to use this endpoint.For more information about environments, see "[Using environments for deployment](https://docs.github.com/en/actions/deployment/targeting-different-environments/using-environments-for-deployment)."For more information about the app that is providing this custom deployment rule, see "[GET an app](https://docs.github.com/rest/apps/apps#get-an-app)".
func (m *ItemItemEnvironmentsItemDeployment_protection_rulesAppsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemItemEnvironmentsItemDeployment_protection_rulesAppsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
    requestInfo.Headers.TryAdd("Accept", "application/json;q=1")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
func (m *ItemItemEnvironmentsItemDeployment_protection_rulesAppsRequestBuilder) WithUrl(rawUrl string)(*ItemItemEnvironmentsItemDeployment_protection_rulesAppsRequestBuilder) {
    return NewItemItemEnvironmentsItemDeployment_protection_rulesAppsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
