package orgs

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemActionsHostedRunnersImagesGithubOwnedRequestBuilder builds and executes requests for operations under \orgs\{org}\actions\hosted-runners\images\github-owned
type ItemActionsHostedRunnersImagesGithubOwnedRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// NewItemActionsHostedRunnersImagesGithubOwnedRequestBuilderInternal instantiates a new ItemActionsHostedRunnersImagesGithubOwnedRequestBuilder and sets the default values.
func NewItemActionsHostedRunnersImagesGithubOwnedRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemActionsHostedRunnersImagesGithubOwnedRequestBuilder) {
    m := &ItemActionsHostedRunnersImagesGithubOwnedRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/orgs/{org}/actions/hosted-runners/images/github-owned", pathParameters),
    }
    return m
}
// NewItemActionsHostedRunnersImagesGithubOwnedRequestBuilder instantiates a new ItemActionsHostedRunnersImagesGithubOwnedRequestBuilder and sets the default values.
func NewItemActionsHostedRunnersImagesGithubOwnedRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemActionsHostedRunnersImagesGithubOwnedRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemActionsHostedRunnersImagesGithubOwnedRequestBuilderInternal(urlParams, requestAdapter)
}
// Get get the list of GitHub-owned images available for GitHub-hosted runners for an organization.
// returns a ItemActionsHostedRunnersImagesGithubOwnedGetResponseable when successful
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/actions/hosted-runners#get-github-owned-images-for-github-hosted-runners-in-an-organization
func (m *ItemActionsHostedRunnersImagesGithubOwnedRequestBuilder) Get(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(ItemActionsHostedRunnersImagesGithubOwnedGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateItemActionsHostedRunnersImagesGithubOwnedGetResponseFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ItemActionsHostedRunnersImagesGithubOwnedGetResponseable), nil
}
// ToGetRequestInformation get the list of GitHub-owned images available for GitHub-hosted runners for an organization.
// returns a *RequestInformation when successful
func (m *ItemActionsHostedRunnersImagesGithubOwnedRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ConfigureRequestInformation(requestInfo, requestConfiguration)
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemActionsHostedRunnersImagesGithubOwnedRequestBuilder when successful
func (m *ItemActionsHostedRunnersImagesGithubOwnedRequestBuilder) WithUrl(rawUrl string)(*ItemActionsHostedRunnersImagesGithubOwnedRequestBuilder) {
    return NewItemActionsHostedRunnersImagesGithubOwnedRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
