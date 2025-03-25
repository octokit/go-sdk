package orgs

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemActionsHostedRunnersImagesPartnerRequestBuilder builds and executes requests for operations under \orgs\{org}\actions\hosted-runners\images\partner
type ItemActionsHostedRunnersImagesPartnerRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// NewItemActionsHostedRunnersImagesPartnerRequestBuilderInternal instantiates a new ItemActionsHostedRunnersImagesPartnerRequestBuilder and sets the default values.
func NewItemActionsHostedRunnersImagesPartnerRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemActionsHostedRunnersImagesPartnerRequestBuilder) {
    m := &ItemActionsHostedRunnersImagesPartnerRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/orgs/{org}/actions/hosted-runners/images/partner", pathParameters),
    }
    return m
}
// NewItemActionsHostedRunnersImagesPartnerRequestBuilder instantiates a new ItemActionsHostedRunnersImagesPartnerRequestBuilder and sets the default values.
func NewItemActionsHostedRunnersImagesPartnerRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemActionsHostedRunnersImagesPartnerRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemActionsHostedRunnersImagesPartnerRequestBuilderInternal(urlParams, requestAdapter)
}
// Get get the list of partner images available for GitHub-hosted runners for an organization.
// returns a ItemActionsHostedRunnersImagesPartnerGetResponseable when successful
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/actions/hosted-runners#get-partner-images-for-github-hosted-runners-in-an-organization
func (m *ItemActionsHostedRunnersImagesPartnerRequestBuilder) Get(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(ItemActionsHostedRunnersImagesPartnerGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateItemActionsHostedRunnersImagesPartnerGetResponseFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ItemActionsHostedRunnersImagesPartnerGetResponseable), nil
}
// ToGetRequestInformation get the list of partner images available for GitHub-hosted runners for an organization.
// returns a *RequestInformation when successful
func (m *ItemActionsHostedRunnersImagesPartnerRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ConfigureRequestInformation(requestInfo, requestConfiguration)
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemActionsHostedRunnersImagesPartnerRequestBuilder when successful
func (m *ItemActionsHostedRunnersImagesPartnerRequestBuilder) WithUrl(rawUrl string)(*ItemActionsHostedRunnersImagesPartnerRequestBuilder) {
    return NewItemActionsHostedRunnersImagesPartnerRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
