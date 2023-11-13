package orgs

import (
    "context"
    i33bcfa37677dfd8b71eaf22ad3bfd140ba1702e7a8418558499c22418bf2fef2 "github.com/octokit/go-sdk/models"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemActionsPermissionsWorkflowRequestBuilder builds and executes requests for operations under \orgs\{org}\actions\permissions\workflow
type ItemActionsPermissionsWorkflowRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemActionsPermissionsWorkflowRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemActionsPermissionsWorkflowRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ItemActionsPermissionsWorkflowRequestBuilderPutRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemActionsPermissionsWorkflowRequestBuilderPutRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemActionsPermissionsWorkflowRequestBuilderInternal instantiates a new WorkflowRequestBuilder and sets the default values.
func NewItemActionsPermissionsWorkflowRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemActionsPermissionsWorkflowRequestBuilder) {
    m := &ItemActionsPermissionsWorkflowRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/orgs/{org}/actions/permissions/workflow", pathParameters),
    }
    return m
}
// NewItemActionsPermissionsWorkflowRequestBuilder instantiates a new WorkflowRequestBuilder and sets the default values.
func NewItemActionsPermissionsWorkflowRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemActionsPermissionsWorkflowRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemActionsPermissionsWorkflowRequestBuilderInternal(urlParams, requestAdapter)
}
// Get gets the default workflow permissions granted to the `GITHUB_TOKEN` when running workflows in an organization,as well as whether GitHub Actions can submit approving pull request reviews. For more information, see"[Setting the permissions of the GITHUB_TOKEN for your organization](https://docs.github.com/organizations/managing-organization-settings/disabling-or-limiting-github-actions-for-your-organization#setting-the-permissions-of-the-github_token-for-your-organization)."You must authenticate using an access token with the `admin:org` scope to use this endpoint. GitHub Apps must have the `administration` organization permission to use this API.
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/actions/permissions#get-default-workflow-permissions-for-an-organization
func (m *ItemActionsPermissionsWorkflowRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemActionsPermissionsWorkflowRequestBuilderGetRequestConfiguration)(i33bcfa37677dfd8b71eaf22ad3bfd140ba1702e7a8418558499c22418bf2fef2.ActionsGetDefaultWorkflowPermissionsable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i33bcfa37677dfd8b71eaf22ad3bfd140ba1702e7a8418558499c22418bf2fef2.CreateActionsGetDefaultWorkflowPermissionsFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i33bcfa37677dfd8b71eaf22ad3bfd140ba1702e7a8418558499c22418bf2fef2.ActionsGetDefaultWorkflowPermissionsable), nil
}
// Put sets the default workflow permissions granted to the `GITHUB_TOKEN` when running workflows in an organization, and sets if GitHub Actionscan submit approving pull request reviews. For more information, see"[Setting the permissions of the GITHUB_TOKEN for your organization](https://docs.github.com/organizations/managing-organization-settings/disabling-or-limiting-github-actions-for-your-organization#setting-the-permissions-of-the-github_token-for-your-organization)."You must authenticate using an access token with the `admin:org` scope to use this endpoint. GitHub Apps must have the `administration` organization permission to use this API.
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/actions/permissions#set-default-workflow-permissions-for-an-organization
func (m *ItemActionsPermissionsWorkflowRequestBuilder) Put(ctx context.Context, body i33bcfa37677dfd8b71eaf22ad3bfd140ba1702e7a8418558499c22418bf2fef2.ActionsSetDefaultWorkflowPermissionsable, requestConfiguration *ItemActionsPermissionsWorkflowRequestBuilderPutRequestConfiguration)(error) {
    requestInfo, err := m.ToPutRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return err
    }
    err = m.BaseRequestBuilder.RequestAdapter.SendNoContent(ctx, requestInfo, nil)
    if err != nil {
        return err
    }
    return nil
}
// ToGetRequestInformation gets the default workflow permissions granted to the `GITHUB_TOKEN` when running workflows in an organization,as well as whether GitHub Actions can submit approving pull request reviews. For more information, see"[Setting the permissions of the GITHUB_TOKEN for your organization](https://docs.github.com/organizations/managing-organization-settings/disabling-or-limiting-github-actions-for-your-organization#setting-the-permissions-of-the-github_token-for-your-organization)."You must authenticate using an access token with the `admin:org` scope to use this endpoint. GitHub Apps must have the `administration` organization permission to use this API.
func (m *ItemActionsPermissionsWorkflowRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemActionsPermissionsWorkflowRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPutRequestInformation sets the default workflow permissions granted to the `GITHUB_TOKEN` when running workflows in an organization, and sets if GitHub Actionscan submit approving pull request reviews. For more information, see"[Setting the permissions of the GITHUB_TOKEN for your organization](https://docs.github.com/organizations/managing-organization-settings/disabling-or-limiting-github-actions-for-your-organization#setting-the-permissions-of-the-github_token-for-your-organization)."You must authenticate using an access token with the `admin:org` scope to use this endpoint. GitHub Apps must have the `administration` organization permission to use this API.
func (m *ItemActionsPermissionsWorkflowRequestBuilder) ToPutRequestInformation(ctx context.Context, body i33bcfa37677dfd8b71eaf22ad3bfd140ba1702e7a8418558499c22418bf2fef2.ActionsSetDefaultWorkflowPermissionsable, requestConfiguration *ItemActionsPermissionsWorkflowRequestBuilderPutRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PUT
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
func (m *ItemActionsPermissionsWorkflowRequestBuilder) WithUrl(rawUrl string)(*ItemActionsPermissionsWorkflowRequestBuilder) {
    return NewItemActionsPermissionsWorkflowRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
