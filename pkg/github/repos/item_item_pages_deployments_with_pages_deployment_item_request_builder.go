package repos

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6 "github.com/octokit/go-sdk/pkg/github/models"
)

// ItemItemPagesDeploymentsWithPages_deployment_ItemRequestBuilder builds and executes requests for operations under \repos\{repos-id}\{Owner-id}\pages\deployments\{pages_deployment_id}
type ItemItemPagesDeploymentsWithPages_deployment_ItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// Cancel the cancel property
func (m *ItemItemPagesDeploymentsWithPages_deployment_ItemRequestBuilder) Cancel()(*ItemItemPagesDeploymentsItemCancelRequestBuilder) {
    return NewItemItemPagesDeploymentsItemCancelRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewItemItemPagesDeploymentsWithPages_deployment_ItemRequestBuilderInternal instantiates a new WithPages_deployment_ItemRequestBuilder and sets the default values.
func NewItemItemPagesDeploymentsWithPages_deployment_ItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemPagesDeploymentsWithPages_deployment_ItemRequestBuilder) {
    m := &ItemItemPagesDeploymentsWithPages_deployment_ItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/repos/{repos%2Did}/{Owner%2Did}/pages/deployments/{pages_deployment_id}", pathParameters),
    }
    return m
}
// NewItemItemPagesDeploymentsWithPages_deployment_ItemRequestBuilder instantiates a new WithPages_deployment_ItemRequestBuilder and sets the default values.
func NewItemItemPagesDeploymentsWithPages_deployment_ItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemPagesDeploymentsWithPages_deployment_ItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemItemPagesDeploymentsWithPages_deployment_ItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Get gets the current status of a GitHub Pages deployment.Users must have read permission for the GitHub Pages site. GitHub Apps must have the `pages:read` permission.
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/pages/pages#get-the-status-of-a-github-pages-deployment
func (m *ItemItemPagesDeploymentsWithPages_deployment_ItemRequestBuilder) Get(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.PagesDeploymentStatusable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "404": i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CreateBasicErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CreatePagesDeploymentStatusFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.PagesDeploymentStatusable), nil
}
// ToGetRequestInformation gets the current status of a GitHub Pages deployment.Users must have read permission for the GitHub Pages site. GitHub Apps must have the `pages:read` permission.
func (m *ItemItemPagesDeploymentsWithPages_deployment_ItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ConfigureRequestInformation(requestInfo, requestConfiguration)
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
func (m *ItemItemPagesDeploymentsWithPages_deployment_ItemRequestBuilder) WithUrl(rawUrl string)(*ItemItemPagesDeploymentsWithPages_deployment_ItemRequestBuilder) {
    return NewItemItemPagesDeploymentsWithPages_deployment_ItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
