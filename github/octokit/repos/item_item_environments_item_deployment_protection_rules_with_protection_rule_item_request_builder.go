package repos

import (
    "context"
    i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e "github.com/octokit/go-sdk/github/octokit/models"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemItemEnvironmentsItemDeployment_protection_rulesWithProtection_rule_ItemRequestBuilder builds and executes requests for operations under \repos\{owner}\{repo}\environments\{environment_name}\deployment_protection_rules\{protection_rule_id}
type ItemItemEnvironmentsItemDeployment_protection_rulesWithProtection_rule_ItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemItemEnvironmentsItemDeployment_protection_rulesWithProtection_rule_ItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemItemEnvironmentsItemDeployment_protection_rulesWithProtection_rule_ItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ItemItemEnvironmentsItemDeployment_protection_rulesWithProtection_rule_ItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemItemEnvironmentsItemDeployment_protection_rulesWithProtection_rule_ItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemItemEnvironmentsItemDeployment_protection_rulesWithProtection_rule_ItemRequestBuilderInternal instantiates a new WithProtection_rule_ItemRequestBuilder and sets the default values.
func NewItemItemEnvironmentsItemDeployment_protection_rulesWithProtection_rule_ItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemEnvironmentsItemDeployment_protection_rulesWithProtection_rule_ItemRequestBuilder) {
    m := &ItemItemEnvironmentsItemDeployment_protection_rulesWithProtection_rule_ItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/repos/{owner}/{repo}/environments/{environment_name}/deployment_protection_rules/{protection_rule_id}", pathParameters),
    }
    return m
}
// NewItemItemEnvironmentsItemDeployment_protection_rulesWithProtection_rule_ItemRequestBuilder instantiates a new WithProtection_rule_ItemRequestBuilder and sets the default values.
func NewItemItemEnvironmentsItemDeployment_protection_rulesWithProtection_rule_ItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemEnvironmentsItemDeployment_protection_rulesWithProtection_rule_ItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemItemEnvironmentsItemDeployment_protection_rulesWithProtection_rule_ItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete disables a custom deployment protection rule for an environment.You must authenticate using an access token with the `repo` scope to use this endpoint. Removing a custom protection rule requires admin or owner permissions to the repository. GitHub Apps must have the `actions:write` permission to use this endpoint. For more information, see "[Get an app](https://docs.github.com/rest/apps/apps#get-an-app)".
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/deployments/protection-rules#disable-a-custom-protection-rule-for-an-environment
func (m *ItemItemEnvironmentsItemDeployment_protection_rulesWithProtection_rule_ItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *ItemItemEnvironmentsItemDeployment_protection_rulesWithProtection_rule_ItemRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.ToDeleteRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    err = m.BaseRequestBuilder.RequestAdapter.SendNoContent(ctx, requestInfo, nil)
    if err != nil {
        return err
    }
    return nil
}
// Get gets an enabled custom deployment protection rule for an environment. Anyone with read access to the repository can use this endpoint. If the repository is private and you want to use a personal access token (classic), you must use an access token with the `repo` scope. GitHub Apps and fine-grained personal access tokens must have the `actions:read` permission to use this endpoint. For more information about environments, see "[Using environments for deployment](https://docs.github.com/en/actions/deployment/targeting-different-environments/using-environments-for-deployment)."For more information about the app that is providing this custom deployment rule, see [`GET /apps/{app_slug}`](https://docs.github.com/rest/apps/apps#get-an-app).
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/deployments/protection-rules#get-a-custom-deployment-protection-rule
func (m *ItemItemEnvironmentsItemDeployment_protection_rulesWithProtection_rule_ItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemItemEnvironmentsItemDeployment_protection_rulesWithProtection_rule_ItemRequestBuilderGetRequestConfiguration)(i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.DeploymentProtectionRuleable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
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
// ToDeleteRequestInformation disables a custom deployment protection rule for an environment.You must authenticate using an access token with the `repo` scope to use this endpoint. Removing a custom protection rule requires admin or owner permissions to the repository. GitHub Apps must have the `actions:write` permission to use this endpoint. For more information, see "[Get an app](https://docs.github.com/rest/apps/apps#get-an-app)".
func (m *ItemItemEnvironmentsItemDeployment_protection_rulesWithProtection_rule_ItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *ItemItemEnvironmentsItemDeployment_protection_rulesWithProtection_rule_ItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE
    return requestInfo, nil
}
// ToGetRequestInformation gets an enabled custom deployment protection rule for an environment. Anyone with read access to the repository can use this endpoint. If the repository is private and you want to use a personal access token (classic), you must use an access token with the `repo` scope. GitHub Apps and fine-grained personal access tokens must have the `actions:read` permission to use this endpoint. For more information about environments, see "[Using environments for deployment](https://docs.github.com/en/actions/deployment/targeting-different-environments/using-environments-for-deployment)."For more information about the app that is providing this custom deployment rule, see [`GET /apps/{app_slug}`](https://docs.github.com/rest/apps/apps#get-an-app).
func (m *ItemItemEnvironmentsItemDeployment_protection_rulesWithProtection_rule_ItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemItemEnvironmentsItemDeployment_protection_rulesWithProtection_rule_ItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *ItemItemEnvironmentsItemDeployment_protection_rulesWithProtection_rule_ItemRequestBuilder) WithUrl(rawUrl string)(*ItemItemEnvironmentsItemDeployment_protection_rulesWithProtection_rule_ItemRequestBuilder) {
    return NewItemItemEnvironmentsItemDeployment_protection_rulesWithProtection_rule_ItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
