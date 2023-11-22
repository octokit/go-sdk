package repos

import (
    "context"
    i53ac87e8cb3cc9276228f74d38694a208cacb99bb8ceb705eeae99fb88d4d274 "strconv"
    i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e "github.com/octokit/go-sdk/github/octokit/models"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemItemEnvironmentsItemDeployment_protection_rulesRequestBuilder builds and executes requests for operations under \repos\{owner}\{repo}\environments\{environment_name}\deployment_protection_rules
type ItemItemEnvironmentsItemDeployment_protection_rulesRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemItemEnvironmentsItemDeployment_protection_rulesRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemItemEnvironmentsItemDeployment_protection_rulesRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ItemItemEnvironmentsItemDeployment_protection_rulesRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemItemEnvironmentsItemDeployment_protection_rulesRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// Apps the apps property
func (m *ItemItemEnvironmentsItemDeployment_protection_rulesRequestBuilder) Apps()(*ItemItemEnvironmentsItemDeployment_protection_rulesAppsRequestBuilder) {
    return NewItemItemEnvironmentsItemDeployment_protection_rulesAppsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ByProtection_rule_id gets an item from the octokit.repos.item.item.environments.item.deployment_protection_rules.item collection
// Deprecated: This indexer is deprecated and will be removed in the next major version. Use the one with the typed parameter instead.
func (m *ItemItemEnvironmentsItemDeployment_protection_rulesRequestBuilder) ByProtection_rule_id(protection_rule_id string)(*ItemItemEnvironmentsItemDeployment_protection_rulesWithProtection_rule_ItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if protection_rule_id != "" {
        urlTplParams["protection_rule_id"] = protection_rule_id
    }
    return NewItemItemEnvironmentsItemDeployment_protection_rulesWithProtection_rule_ItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// ByProtection_rule_idInteger gets an item from the octokit.repos.item.item.environments.item.deployment_protection_rules.item collection
func (m *ItemItemEnvironmentsItemDeployment_protection_rulesRequestBuilder) ByProtection_rule_idInteger(protection_rule_id int32)(*ItemItemEnvironmentsItemDeployment_protection_rulesWithProtection_rule_ItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    urlTplParams["protection_rule_id"] = i53ac87e8cb3cc9276228f74d38694a208cacb99bb8ceb705eeae99fb88d4d274.FormatInt(int64(protection_rule_id), 10)
    return NewItemItemEnvironmentsItemDeployment_protection_rulesWithProtection_rule_ItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewItemItemEnvironmentsItemDeployment_protection_rulesRequestBuilderInternal instantiates a new Deployment_protection_rulesRequestBuilder and sets the default values.
func NewItemItemEnvironmentsItemDeployment_protection_rulesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemEnvironmentsItemDeployment_protection_rulesRequestBuilder) {
    m := &ItemItemEnvironmentsItemDeployment_protection_rulesRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/repos/{owner}/{repo}/environments/{environment_name}/deployment_protection_rules", pathParameters),
    }
    return m
}
// NewItemItemEnvironmentsItemDeployment_protection_rulesRequestBuilder instantiates a new Deployment_protection_rulesRequestBuilder and sets the default values.
func NewItemItemEnvironmentsItemDeployment_protection_rulesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemEnvironmentsItemDeployment_protection_rulesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemItemEnvironmentsItemDeployment_protection_rulesRequestBuilderInternal(urlParams, requestAdapter)
}
// Get gets all custom deployment protection rules that are enabled for an environment. Anyone with read access to the repository can use this endpoint. If the repository is private and you want to use a personal access token (classic), you must use an access token with the `repo` scope. GitHub Apps and fine-grained personal access tokens must have the `actions:read` permission to use this endpoint. For more information about environments, see "[Using environments for deployment](https://docs.github.com/en/actions/deployment/targeting-different-environments/using-environments-for-deployment)."For more information about the app that is providing this custom deployment rule, see the [documentation for the `GET /apps/{app_slug}` endpoint](https://docs.github.com/rest/apps/apps#get-an-app).
// Deprecated: This method is obsolete. Use GetAsDeployment_protection_rulesGetResponse instead.
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/deployments/protection-rules#get-all-deployment-protection-rules-for-an-environment
func (m *ItemItemEnvironmentsItemDeployment_protection_rulesRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemItemEnvironmentsItemDeployment_protection_rulesRequestBuilderGetRequestConfiguration)(ItemItemEnvironmentsItemDeployment_protection_rulesResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateItemItemEnvironmentsItemDeployment_protection_rulesResponseFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ItemItemEnvironmentsItemDeployment_protection_rulesResponseable), nil
}
// GetAsDeployment_protection_rulesGetResponse gets all custom deployment protection rules that are enabled for an environment. Anyone with read access to the repository can use this endpoint. If the repository is private and you want to use a personal access token (classic), you must use an access token with the `repo` scope. GitHub Apps and fine-grained personal access tokens must have the `actions:read` permission to use this endpoint. For more information about environments, see "[Using environments for deployment](https://docs.github.com/en/actions/deployment/targeting-different-environments/using-environments-for-deployment)."For more information about the app that is providing this custom deployment rule, see the [documentation for the `GET /apps/{app_slug}` endpoint](https://docs.github.com/rest/apps/apps#get-an-app).
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/deployments/protection-rules#get-all-deployment-protection-rules-for-an-environment
func (m *ItemItemEnvironmentsItemDeployment_protection_rulesRequestBuilder) GetAsDeployment_protection_rulesGetResponse(ctx context.Context, requestConfiguration *ItemItemEnvironmentsItemDeployment_protection_rulesRequestBuilderGetRequestConfiguration)(ItemItemEnvironmentsItemDeployment_protection_rulesGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateItemItemEnvironmentsItemDeployment_protection_rulesGetResponseFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ItemItemEnvironmentsItemDeployment_protection_rulesGetResponseable), nil
}
// Post enable a custom deployment protection rule for an environment.You must authenticate using an access token with the `repo` scope to use this endpoint. Enabling a custom protection rule requires admin or owner permissions to the repository. GitHub Apps must have the `actions:write` permission to use this endpoint.For more information about the app that is providing this custom deployment rule, see the [documentation for the `GET /apps/{app_slug}` endpoint](https://docs.github.com/rest/apps/apps#get-an-app).
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/deployments/protection-rules#create-a-custom-deployment-protection-rule-on-an-environment
func (m *ItemItemEnvironmentsItemDeployment_protection_rulesRequestBuilder) Post(ctx context.Context, body ItemItemEnvironmentsItemDeployment_protection_rulesPostRequestBodyable, requestConfiguration *ItemItemEnvironmentsItemDeployment_protection_rulesRequestBuilderPostRequestConfiguration)(i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.DeploymentProtectionRuleable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.CreateDeploymentProtectionRuleFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.DeploymentProtectionRuleable), nil
}
// ToGetRequestInformation gets all custom deployment protection rules that are enabled for an environment. Anyone with read access to the repository can use this endpoint. If the repository is private and you want to use a personal access token (classic), you must use an access token with the `repo` scope. GitHub Apps and fine-grained personal access tokens must have the `actions:read` permission to use this endpoint. For more information about environments, see "[Using environments for deployment](https://docs.github.com/en/actions/deployment/targeting-different-environments/using-environments-for-deployment)."For more information about the app that is providing this custom deployment rule, see the [documentation for the `GET /apps/{app_slug}` endpoint](https://docs.github.com/rest/apps/apps#get-an-app).
func (m *ItemItemEnvironmentsItemDeployment_protection_rulesRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemItemEnvironmentsItemDeployment_protection_rulesRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation enable a custom deployment protection rule for an environment.You must authenticate using an access token with the `repo` scope to use this endpoint. Enabling a custom protection rule requires admin or owner permissions to the repository. GitHub Apps must have the `actions:write` permission to use this endpoint.For more information about the app that is providing this custom deployment rule, see the [documentation for the `GET /apps/{app_slug}` endpoint](https://docs.github.com/rest/apps/apps#get-an-app).
func (m *ItemItemEnvironmentsItemDeployment_protection_rulesRequestBuilder) ToPostRequestInformation(ctx context.Context, body ItemItemEnvironmentsItemDeployment_protection_rulesPostRequestBodyable, requestConfiguration *ItemItemEnvironmentsItemDeployment_protection_rulesRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
    requestInfo.Headers.TryAdd("Accept", "application/json")
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
func (m *ItemItemEnvironmentsItemDeployment_protection_rulesRequestBuilder) WithUrl(rawUrl string)(*ItemItemEnvironmentsItemDeployment_protection_rulesRequestBuilder) {
    return NewItemItemEnvironmentsItemDeployment_protection_rulesRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
