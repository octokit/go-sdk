package repos

import (
    "context"
    i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e "github.com/octokit/go-sdk/github/octokit/models"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemItemEnvironmentsWithEnvironment_nameItemRequestBuilder builds and executes requests for operations under \repos\{owner}\{repo}\environments\{environment_name}
type ItemItemEnvironmentsWithEnvironment_nameItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemItemEnvironmentsWithEnvironment_nameItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemItemEnvironmentsWithEnvironment_nameItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ItemItemEnvironmentsWithEnvironment_nameItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemItemEnvironmentsWithEnvironment_nameItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ItemItemEnvironmentsWithEnvironment_nameItemRequestBuilderPutRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemItemEnvironmentsWithEnvironment_nameItemRequestBuilderPutRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemItemEnvironmentsWithEnvironment_nameItemRequestBuilderInternal instantiates a new WithEnvironment_nameItemRequestBuilder and sets the default values.
func NewItemItemEnvironmentsWithEnvironment_nameItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemEnvironmentsWithEnvironment_nameItemRequestBuilder) {
    m := &ItemItemEnvironmentsWithEnvironment_nameItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/repos/{owner}/{repo}/environments/{environment_name}", pathParameters),
    }
    return m
}
// NewItemItemEnvironmentsWithEnvironment_nameItemRequestBuilder instantiates a new WithEnvironment_nameItemRequestBuilder and sets the default values.
func NewItemItemEnvironmentsWithEnvironment_nameItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemEnvironmentsWithEnvironment_nameItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemItemEnvironmentsWithEnvironment_nameItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete you must authenticate using an access token with the repo scope to use this endpoint.
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/deployments/environments#delete-an-environment
func (m *ItemItemEnvironmentsWithEnvironment_nameItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *ItemItemEnvironmentsWithEnvironment_nameItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Deployment_protection_rules the deployment_protection_rules property
func (m *ItemItemEnvironmentsWithEnvironment_nameItemRequestBuilder) Deployment_protection_rules()(*ItemItemEnvironmentsItemDeployment_protection_rulesRequestBuilder) {
    return NewItemItemEnvironmentsItemDeployment_protection_rulesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// DeploymentBranchPolicies the deploymentBranchPolicies property
func (m *ItemItemEnvironmentsWithEnvironment_nameItemRequestBuilder) DeploymentBranchPolicies()(*ItemItemEnvironmentsItemDeploymentBranchPoliciesRequestBuilder) {
    return NewItemItemEnvironmentsItemDeploymentBranchPoliciesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get **Note:** To get information about name patterns that branches must match in order to deploy to this environment, see "[Get a deployment branch policy](/rest/deployments/branch-policies#get-a-deployment-branch-policy)."Anyone with read access to the repository can use this endpoint. If therepository is private, you must use an access token with the `repo` scope. GitHubApps must have the `actions:read` permission to use this endpoint.
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/deployments/environments#get-an-environment
func (m *ItemItemEnvironmentsWithEnvironment_nameItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemItemEnvironmentsWithEnvironment_nameItemRequestBuilderGetRequestConfiguration)(i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.Environmentable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.CreateEnvironmentFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.Environmentable), nil
}
// Put create or update an environment with protection rules, such as required reviewers. For more information about environment protection rules, see "[Environments](/actions/reference/environments#environment-protection-rules)."**Note:** To create or update name patterns that branches must match in order to deploy to this environment, see "[Deployment branch policies](/rest/deployments/branch-policies)."**Note:** To create or update secrets for an environment, see "[GitHub Actions secrets](/rest/actions/secrets)."You must authenticate using an access token with the `repo` scope to use this endpoint. GitHub Apps must have the `administration:write` permission for the repository to use this endpoint.
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/deployments/environments#create-or-update-an-environment
func (m *ItemItemEnvironmentsWithEnvironment_nameItemRequestBuilder) Put(ctx context.Context, body ItemItemEnvironmentsItemWithEnvironment_namePutRequestBodyable, requestConfiguration *ItemItemEnvironmentsWithEnvironment_nameItemRequestBuilderPutRequestConfiguration)(i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.Environmentable, error) {
    requestInfo, err := m.ToPutRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "422": i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.CreateBasicErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.CreateEnvironmentFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.Environmentable), nil
}
// ToDeleteRequestInformation you must authenticate using an access token with the repo scope to use this endpoint.
func (m *ItemItemEnvironmentsWithEnvironment_nameItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *ItemItemEnvironmentsWithEnvironment_nameItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToGetRequestInformation **Note:** To get information about name patterns that branches must match in order to deploy to this environment, see "[Get a deployment branch policy](/rest/deployments/branch-policies#get-a-deployment-branch-policy)."Anyone with read access to the repository can use this endpoint. If therepository is private, you must use an access token with the `repo` scope. GitHubApps must have the `actions:read` permission to use this endpoint.
func (m *ItemItemEnvironmentsWithEnvironment_nameItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemItemEnvironmentsWithEnvironment_nameItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPutRequestInformation create or update an environment with protection rules, such as required reviewers. For more information about environment protection rules, see "[Environments](/actions/reference/environments#environment-protection-rules)."**Note:** To create or update name patterns that branches must match in order to deploy to this environment, see "[Deployment branch policies](/rest/deployments/branch-policies)."**Note:** To create or update secrets for an environment, see "[GitHub Actions secrets](/rest/actions/secrets)."You must authenticate using an access token with the `repo` scope to use this endpoint. GitHub Apps must have the `administration:write` permission for the repository to use this endpoint.
func (m *ItemItemEnvironmentsWithEnvironment_nameItemRequestBuilder) ToPutRequestInformation(ctx context.Context, body ItemItemEnvironmentsItemWithEnvironment_namePutRequestBodyable, requestConfiguration *ItemItemEnvironmentsWithEnvironment_nameItemRequestBuilderPutRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PUT
    requestInfo.Headers.TryAdd("Accept", "application/json")
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
func (m *ItemItemEnvironmentsWithEnvironment_nameItemRequestBuilder) WithUrl(rawUrl string)(*ItemItemEnvironmentsWithEnvironment_nameItemRequestBuilder) {
    return NewItemItemEnvironmentsWithEnvironment_nameItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
