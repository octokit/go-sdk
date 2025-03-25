package orgs

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemActionsHostedRunnersMachineSizesRequestBuilder builds and executes requests for operations under \orgs\{org}\actions\hosted-runners\machine-sizes
type ItemActionsHostedRunnersMachineSizesRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// NewItemActionsHostedRunnersMachineSizesRequestBuilderInternal instantiates a new ItemActionsHostedRunnersMachineSizesRequestBuilder and sets the default values.
func NewItemActionsHostedRunnersMachineSizesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemActionsHostedRunnersMachineSizesRequestBuilder) {
    m := &ItemActionsHostedRunnersMachineSizesRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/orgs/{org}/actions/hosted-runners/machine-sizes", pathParameters),
    }
    return m
}
// NewItemActionsHostedRunnersMachineSizesRequestBuilder instantiates a new ItemActionsHostedRunnersMachineSizesRequestBuilder and sets the default values.
func NewItemActionsHostedRunnersMachineSizesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemActionsHostedRunnersMachineSizesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemActionsHostedRunnersMachineSizesRequestBuilderInternal(urlParams, requestAdapter)
}
// Get get the list of machine specs available for GitHub-hosted runners for an organization.
// returns a ItemActionsHostedRunnersMachineSizesGetResponseable when successful
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/actions/hosted-runners#get-github-hosted-runners-machine-specs-for-an-organization
func (m *ItemActionsHostedRunnersMachineSizesRequestBuilder) Get(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(ItemActionsHostedRunnersMachineSizesGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateItemActionsHostedRunnersMachineSizesGetResponseFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ItemActionsHostedRunnersMachineSizesGetResponseable), nil
}
// ToGetRequestInformation get the list of machine specs available for GitHub-hosted runners for an organization.
// returns a *RequestInformation when successful
func (m *ItemActionsHostedRunnersMachineSizesRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ConfigureRequestInformation(requestInfo, requestConfiguration)
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemActionsHostedRunnersMachineSizesRequestBuilder when successful
func (m *ItemActionsHostedRunnersMachineSizesRequestBuilder) WithUrl(rawUrl string)(*ItemActionsHostedRunnersMachineSizesRequestBuilder) {
    return NewItemActionsHostedRunnersMachineSizesRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
