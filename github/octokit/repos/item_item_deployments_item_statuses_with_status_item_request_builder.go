package repos

import (
    "context"
    i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e "github.com/octokit/go-sdk/github/octokit/models"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemItemDeploymentsItemStatusesWithStatus_ItemRequestBuilder builds and executes requests for operations under \repos\{owner}\{repo}\deployments\{deployment_id}\statuses\{status_id}
type ItemItemDeploymentsItemStatusesWithStatus_ItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemItemDeploymentsItemStatusesWithStatus_ItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemItemDeploymentsItemStatusesWithStatus_ItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemItemDeploymentsItemStatusesWithStatus_ItemRequestBuilderInternal instantiates a new WithStatus_ItemRequestBuilder and sets the default values.
func NewItemItemDeploymentsItemStatusesWithStatus_ItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemDeploymentsItemStatusesWithStatus_ItemRequestBuilder) {
    m := &ItemItemDeploymentsItemStatusesWithStatus_ItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/repos/{owner}/{repo}/deployments/{deployment_id}/statuses/{status_id}", pathParameters),
    }
    return m
}
// NewItemItemDeploymentsItemStatusesWithStatus_ItemRequestBuilder instantiates a new WithStatus_ItemRequestBuilder and sets the default values.
func NewItemItemDeploymentsItemStatusesWithStatus_ItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemDeploymentsItemStatusesWithStatus_ItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemItemDeploymentsItemStatusesWithStatus_ItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Get users with pull access can view a deployment status for a deployment:
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/deployments/statuses#get-a-deployment-status
func (m *ItemItemDeploymentsItemStatusesWithStatus_ItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemItemDeploymentsItemStatusesWithStatus_ItemRequestBuilderGetRequestConfiguration)(i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.DeploymentStatusable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "404": i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.CreateBasicErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.CreateDeploymentStatusFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.DeploymentStatusable), nil
}
// ToGetRequestInformation users with pull access can view a deployment status for a deployment:
func (m *ItemItemDeploymentsItemStatusesWithStatus_ItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemItemDeploymentsItemStatusesWithStatus_ItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *ItemItemDeploymentsItemStatusesWithStatus_ItemRequestBuilder) WithUrl(rawUrl string)(*ItemItemDeploymentsItemStatusesWithStatus_ItemRequestBuilder) {
    return NewItemItemDeploymentsItemStatusesWithStatus_ItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
