package repos

import (
    "context"
    i53ac87e8cb3cc9276228f74d38694a208cacb99bb8ceb705eeae99fb88d4d274 "strconv"
    i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e "github.com/octokit/go-sdk/github/octokit/models"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemItemEnvironmentsItemDeploymentBranchPoliciesRequestBuilder builds and executes requests for operations under \repos\{owner}\{repo}\environments\{environment_name}\deployment-branch-policies
type ItemItemEnvironmentsItemDeploymentBranchPoliciesRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemItemEnvironmentsItemDeploymentBranchPoliciesRequestBuilderGetQueryParameters lists the deployment branch policies for an environment.Anyone with read access to the repository can use this endpoint. If the repository is private, you must use an access token with the `repo` scope. GitHub Apps must have the `actions:read` permission to use this endpoint.
type ItemItemEnvironmentsItemDeploymentBranchPoliciesRequestBuilderGetQueryParameters struct {
    // Page number of the results to fetch.
    Page *int32 `uriparametername:"page"`
    // The number of results per page (max 100).
    Per_page *int32 `uriparametername:"per_page"`
}
// ItemItemEnvironmentsItemDeploymentBranchPoliciesRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemItemEnvironmentsItemDeploymentBranchPoliciesRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemItemEnvironmentsItemDeploymentBranchPoliciesRequestBuilderGetQueryParameters
}
// ItemItemEnvironmentsItemDeploymentBranchPoliciesRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemItemEnvironmentsItemDeploymentBranchPoliciesRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ByBranch_policy_id gets an item from the octokit.repos.item.item.environments.item.deploymentBranchPolicies.item collection
// Deprecated: This indexer is deprecated and will be removed in the next major version. Use the one with the typed parameter instead.
func (m *ItemItemEnvironmentsItemDeploymentBranchPoliciesRequestBuilder) ByBranch_policy_id(branch_policy_id string)(*ItemItemEnvironmentsItemDeploymentBranchPoliciesWithBranch_policy_ItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if branch_policy_id != "" {
        urlTplParams["branch_policy_id"] = branch_policy_id
    }
    return NewItemItemEnvironmentsItemDeploymentBranchPoliciesWithBranch_policy_ItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// ByBranch_policy_idInteger gets an item from the octokit.repos.item.item.environments.item.deploymentBranchPolicies.item collection
func (m *ItemItemEnvironmentsItemDeploymentBranchPoliciesRequestBuilder) ByBranch_policy_idInteger(branch_policy_id int32)(*ItemItemEnvironmentsItemDeploymentBranchPoliciesWithBranch_policy_ItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    urlTplParams["branch_policy_id"] = i53ac87e8cb3cc9276228f74d38694a208cacb99bb8ceb705eeae99fb88d4d274.FormatInt(int64(branch_policy_id), 10)
    return NewItemItemEnvironmentsItemDeploymentBranchPoliciesWithBranch_policy_ItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewItemItemEnvironmentsItemDeploymentBranchPoliciesRequestBuilderInternal instantiates a new DeploymentBranchPoliciesRequestBuilder and sets the default values.
func NewItemItemEnvironmentsItemDeploymentBranchPoliciesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemEnvironmentsItemDeploymentBranchPoliciesRequestBuilder) {
    m := &ItemItemEnvironmentsItemDeploymentBranchPoliciesRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/repos/{owner}/{repo}/environments/{environment_name}/deployment-branch-policies{?per_page*,page*}", pathParameters),
    }
    return m
}
// NewItemItemEnvironmentsItemDeploymentBranchPoliciesRequestBuilder instantiates a new DeploymentBranchPoliciesRequestBuilder and sets the default values.
func NewItemItemEnvironmentsItemDeploymentBranchPoliciesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemEnvironmentsItemDeploymentBranchPoliciesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemItemEnvironmentsItemDeploymentBranchPoliciesRequestBuilderInternal(urlParams, requestAdapter)
}
// Get lists the deployment branch policies for an environment.Anyone with read access to the repository can use this endpoint. If the repository is private, you must use an access token with the `repo` scope. GitHub Apps must have the `actions:read` permission to use this endpoint.
// Deprecated: This method is obsolete. Use GetAsDeploymentBranchPoliciesGetResponse instead.
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/deployments/branch-policies#list-deployment-branch-policies
func (m *ItemItemEnvironmentsItemDeploymentBranchPoliciesRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemItemEnvironmentsItemDeploymentBranchPoliciesRequestBuilderGetRequestConfiguration)(ItemItemEnvironmentsItemDeploymentBranchPoliciesResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateItemItemEnvironmentsItemDeploymentBranchPoliciesResponseFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ItemItemEnvironmentsItemDeploymentBranchPoliciesResponseable), nil
}
// GetAsDeploymentBranchPoliciesGetResponse lists the deployment branch policies for an environment.Anyone with read access to the repository can use this endpoint. If the repository is private, you must use an access token with the `repo` scope. GitHub Apps must have the `actions:read` permission to use this endpoint.
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/deployments/branch-policies#list-deployment-branch-policies
func (m *ItemItemEnvironmentsItemDeploymentBranchPoliciesRequestBuilder) GetAsDeploymentBranchPoliciesGetResponse(ctx context.Context, requestConfiguration *ItemItemEnvironmentsItemDeploymentBranchPoliciesRequestBuilderGetRequestConfiguration)(ItemItemEnvironmentsItemDeploymentBranchPoliciesGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateItemItemEnvironmentsItemDeploymentBranchPoliciesGetResponseFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ItemItemEnvironmentsItemDeploymentBranchPoliciesGetResponseable), nil
}
// Post creates a deployment branch or tag policy for an environment.You must authenticate using an access token with the `repo` scope to use this endpoint. GitHub Apps must have the `administration:write` permission for the repository to use this endpoint.
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/deployments/branch-policies#create-a-deployment-branch-policy
func (m *ItemItemEnvironmentsItemDeploymentBranchPoliciesRequestBuilder) Post(ctx context.Context, body i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.DeploymentBranchPolicyNamePatternWithTypeable, requestConfiguration *ItemItemEnvironmentsItemDeploymentBranchPoliciesRequestBuilderPostRequestConfiguration)(i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.DeploymentBranchPolicyable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.CreateDeploymentBranchPolicyFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.DeploymentBranchPolicyable), nil
}
// ToGetRequestInformation lists the deployment branch policies for an environment.Anyone with read access to the repository can use this endpoint. If the repository is private, you must use an access token with the `repo` scope. GitHub Apps must have the `actions:read` permission to use this endpoint.
func (m *ItemItemEnvironmentsItemDeploymentBranchPoliciesRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemItemEnvironmentsItemDeploymentBranchPoliciesRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation creates a deployment branch or tag policy for an environment.You must authenticate using an access token with the `repo` scope to use this endpoint. GitHub Apps must have the `administration:write` permission for the repository to use this endpoint.
func (m *ItemItemEnvironmentsItemDeploymentBranchPoliciesRequestBuilder) ToPostRequestInformation(ctx context.Context, body i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.DeploymentBranchPolicyNamePatternWithTypeable, requestConfiguration *ItemItemEnvironmentsItemDeploymentBranchPoliciesRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *ItemItemEnvironmentsItemDeploymentBranchPoliciesRequestBuilder) WithUrl(rawUrl string)(*ItemItemEnvironmentsItemDeploymentBranchPoliciesRequestBuilder) {
    return NewItemItemEnvironmentsItemDeploymentBranchPoliciesRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
