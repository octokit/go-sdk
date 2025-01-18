package orgs

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemActionsHostedRunnersPlatformsRequestBuilder builds and executes requests for operations under \orgs\{org}\actions\hosted-runners\platforms
type ItemActionsHostedRunnersPlatformsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// NewItemActionsHostedRunnersPlatformsRequestBuilderInternal instantiates a new ItemActionsHostedRunnersPlatformsRequestBuilder and sets the default values.
func NewItemActionsHostedRunnersPlatformsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemActionsHostedRunnersPlatformsRequestBuilder) {
    m := &ItemActionsHostedRunnersPlatformsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/orgs/{org}/actions/hosted-runners/platforms", pathParameters),
    }
    return m
}
// NewItemActionsHostedRunnersPlatformsRequestBuilder instantiates a new ItemActionsHostedRunnersPlatformsRequestBuilder and sets the default values.
func NewItemActionsHostedRunnersPlatformsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemActionsHostedRunnersPlatformsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemActionsHostedRunnersPlatformsRequestBuilderInternal(urlParams, requestAdapter)
}
// Get get the list of platforms available for GitHub-hosted runners for an organization.
// returns a ItemActionsHostedRunnersPlatformsGetResponseable when successful
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/actions/hosted-runners#get-platforms-for-github-hosted-runners-in-an-organization
func (m *ItemActionsHostedRunnersPlatformsRequestBuilder) Get(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(ItemActionsHostedRunnersPlatformsGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateItemActionsHostedRunnersPlatformsGetResponseFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ItemActionsHostedRunnersPlatformsGetResponseable), nil
}
// ToGetRequestInformation get the list of platforms available for GitHub-hosted runners for an organization.
// returns a *RequestInformation when successful
func (m *ItemActionsHostedRunnersPlatformsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ConfigureRequestInformation(requestInfo, requestConfiguration)
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemActionsHostedRunnersPlatformsRequestBuilder when successful
func (m *ItemActionsHostedRunnersPlatformsRequestBuilder) WithUrl(rawUrl string)(*ItemActionsHostedRunnersPlatformsRequestBuilder) {
    return NewItemActionsHostedRunnersPlatformsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
