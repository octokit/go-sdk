package orgs

import (
    "context"
    i53ac87e8cb3cc9276228f74d38694a208cacb99bb8ceb705eeae99fb88d4d274 "strconv"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6 "github.com/octokit/go-sdk/pkg/github/models"
)

// ItemActionsHostedRunnersRequestBuilder builds and executes requests for operations under \orgs\{org}\actions\hosted-runners
type ItemActionsHostedRunnersRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemActionsHostedRunnersRequestBuilderGetQueryParameters lists all GitHub-hosted runners configured in an organization.OAuth app tokens and personal access tokens (classic) need the `manage_runner:org` scope to use this endpoint.
type ItemActionsHostedRunnersRequestBuilderGetQueryParameters struct {
    // The page number of the results to fetch. For more information, see "[Using pagination in the REST API](https://docs.github.com/rest/using-the-rest-api/using-pagination-in-the-rest-api)."
    Page *int32 `uriparametername:"page"`
    // The number of results per page (max 100). For more information, see "[Using pagination in the REST API](https://docs.github.com/rest/using-the-rest-api/using-pagination-in-the-rest-api)."
    Per_page *int32 `uriparametername:"per_page"`
}
// ByHosted_runner_id gets an item from the github.com/octokit/go-sdk/pkg/github.orgs.item.actions.hostedRunners.item collection
// returns a *ItemActionsHostedRunnersWithHosted_runner_ItemRequestBuilder when successful
func (m *ItemActionsHostedRunnersRequestBuilder) ByHosted_runner_id(hosted_runner_id int32)(*ItemActionsHostedRunnersWithHosted_runner_ItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    urlTplParams["hosted_runner_id"] = i53ac87e8cb3cc9276228f74d38694a208cacb99bb8ceb705eeae99fb88d4d274.FormatInt(int64(hosted_runner_id), 10)
    return NewItemActionsHostedRunnersWithHosted_runner_ItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewItemActionsHostedRunnersRequestBuilderInternal instantiates a new ItemActionsHostedRunnersRequestBuilder and sets the default values.
func NewItemActionsHostedRunnersRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemActionsHostedRunnersRequestBuilder) {
    m := &ItemActionsHostedRunnersRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/orgs/{org}/actions/hosted-runners{?page*,per_page*}", pathParameters),
    }
    return m
}
// NewItemActionsHostedRunnersRequestBuilder instantiates a new ItemActionsHostedRunnersRequestBuilder and sets the default values.
func NewItemActionsHostedRunnersRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemActionsHostedRunnersRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemActionsHostedRunnersRequestBuilderInternal(urlParams, requestAdapter)
}
// Get lists all GitHub-hosted runners configured in an organization.OAuth app tokens and personal access tokens (classic) need the `manage_runner:org` scope to use this endpoint.
// returns a ItemActionsHostedRunnersGetResponseable when successful
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/actions/hosted-runners#list-github-hosted-runners-for-an-organization
func (m *ItemActionsHostedRunnersRequestBuilder) Get(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[ItemActionsHostedRunnersRequestBuilderGetQueryParameters])(ItemActionsHostedRunnersGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateItemActionsHostedRunnersGetResponseFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ItemActionsHostedRunnersGetResponseable), nil
}
// Images the images property
// returns a *ItemActionsHostedRunnersImagesRequestBuilder when successful
func (m *ItemActionsHostedRunnersRequestBuilder) Images()(*ItemActionsHostedRunnersImagesRequestBuilder) {
    return NewItemActionsHostedRunnersImagesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Limits the limits property
// returns a *ItemActionsHostedRunnersLimitsRequestBuilder when successful
func (m *ItemActionsHostedRunnersRequestBuilder) Limits()(*ItemActionsHostedRunnersLimitsRequestBuilder) {
    return NewItemActionsHostedRunnersLimitsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// MachineSizes the machineSizes property
// returns a *ItemActionsHostedRunnersMachineSizesRequestBuilder when successful
func (m *ItemActionsHostedRunnersRequestBuilder) MachineSizes()(*ItemActionsHostedRunnersMachineSizesRequestBuilder) {
    return NewItemActionsHostedRunnersMachineSizesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Platforms the platforms property
// returns a *ItemActionsHostedRunnersPlatformsRequestBuilder when successful
func (m *ItemActionsHostedRunnersRequestBuilder) Platforms()(*ItemActionsHostedRunnersPlatformsRequestBuilder) {
    return NewItemActionsHostedRunnersPlatformsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Post creates a GitHub-hosted runner for an organization.OAuth tokens and personal access tokens (classic) need the `manage_runners:org` scope to use this endpoint.
// returns a ActionsHostedRunnerable when successful
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/actions/hosted-runners#create-a-github-hosted-runner-for-an-organization
func (m *ItemActionsHostedRunnersRequestBuilder) Post(ctx context.Context, body ItemActionsHostedRunnersPostRequestBodyable, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.ActionsHostedRunnerable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CreateActionsHostedRunnerFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.ActionsHostedRunnerable), nil
}
// ToGetRequestInformation lists all GitHub-hosted runners configured in an organization.OAuth app tokens and personal access tokens (classic) need the `manage_runner:org` scope to use this endpoint.
// returns a *RequestInformation when successful
func (m *ItemActionsHostedRunnersRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[ItemActionsHostedRunnersRequestBuilderGetQueryParameters])(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ConfigureRequestInformation(requestInfo, requestConfiguration)
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToPostRequestInformation creates a GitHub-hosted runner for an organization.OAuth tokens and personal access tokens (classic) need the `manage_runners:org` scope to use this endpoint.
// returns a *RequestInformation when successful
func (m *ItemActionsHostedRunnersRequestBuilder) ToPostRequestInformation(ctx context.Context, body ItemActionsHostedRunnersPostRequestBodyable, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ConfigureRequestInformation(requestInfo, requestConfiguration)
    requestInfo.Headers.TryAdd("Accept", "application/json")
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemActionsHostedRunnersRequestBuilder when successful
func (m *ItemActionsHostedRunnersRequestBuilder) WithUrl(rawUrl string)(*ItemActionsHostedRunnersRequestBuilder) {
    return NewItemActionsHostedRunnersRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
