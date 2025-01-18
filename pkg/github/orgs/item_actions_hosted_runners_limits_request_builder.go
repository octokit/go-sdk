package orgs

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6 "github.com/octokit/go-sdk/pkg/github/models"
)

// ItemActionsHostedRunnersLimitsRequestBuilder builds and executes requests for operations under \orgs\{org}\actions\hosted-runners\limits
type ItemActionsHostedRunnersLimitsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// NewItemActionsHostedRunnersLimitsRequestBuilderInternal instantiates a new ItemActionsHostedRunnersLimitsRequestBuilder and sets the default values.
func NewItemActionsHostedRunnersLimitsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemActionsHostedRunnersLimitsRequestBuilder) {
    m := &ItemActionsHostedRunnersLimitsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/orgs/{org}/actions/hosted-runners/limits", pathParameters),
    }
    return m
}
// NewItemActionsHostedRunnersLimitsRequestBuilder instantiates a new ItemActionsHostedRunnersLimitsRequestBuilder and sets the default values.
func NewItemActionsHostedRunnersLimitsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemActionsHostedRunnersLimitsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemActionsHostedRunnersLimitsRequestBuilderInternal(urlParams, requestAdapter)
}
// Get get the GitHub-hosted runners limits for an organization.
// returns a ActionsHostedRunnerLimitsable when successful
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/actions/hosted-runners#get-limits-on-github-hosted-runners-for-an-organization
func (m *ItemActionsHostedRunnersLimitsRequestBuilder) Get(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.ActionsHostedRunnerLimitsable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CreateActionsHostedRunnerLimitsFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.ActionsHostedRunnerLimitsable), nil
}
// ToGetRequestInformation get the GitHub-hosted runners limits for an organization.
// returns a *RequestInformation when successful
func (m *ItemActionsHostedRunnersLimitsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ConfigureRequestInformation(requestInfo, requestConfiguration)
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemActionsHostedRunnersLimitsRequestBuilder when successful
func (m *ItemActionsHostedRunnersLimitsRequestBuilder) WithUrl(rawUrl string)(*ItemActionsHostedRunnersLimitsRequestBuilder) {
    return NewItemActionsHostedRunnersLimitsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
