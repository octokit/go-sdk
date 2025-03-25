package orgs

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemActionsHostedRunnersImagesRequestBuilder builds and executes requests for operations under \orgs\{org}\actions\hosted-runners\images
type ItemActionsHostedRunnersImagesRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// NewItemActionsHostedRunnersImagesRequestBuilderInternal instantiates a new ItemActionsHostedRunnersImagesRequestBuilder and sets the default values.
func NewItemActionsHostedRunnersImagesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemActionsHostedRunnersImagesRequestBuilder) {
    m := &ItemActionsHostedRunnersImagesRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/orgs/{org}/actions/hosted-runners/images", pathParameters),
    }
    return m
}
// NewItemActionsHostedRunnersImagesRequestBuilder instantiates a new ItemActionsHostedRunnersImagesRequestBuilder and sets the default values.
func NewItemActionsHostedRunnersImagesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemActionsHostedRunnersImagesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemActionsHostedRunnersImagesRequestBuilderInternal(urlParams, requestAdapter)
}
// GithubOwned the githubOwned property
// returns a *ItemActionsHostedRunnersImagesGithubOwnedRequestBuilder when successful
func (m *ItemActionsHostedRunnersImagesRequestBuilder) GithubOwned()(*ItemActionsHostedRunnersImagesGithubOwnedRequestBuilder) {
    return NewItemActionsHostedRunnersImagesGithubOwnedRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Partner the partner property
// returns a *ItemActionsHostedRunnersImagesPartnerRequestBuilder when successful
func (m *ItemActionsHostedRunnersImagesRequestBuilder) Partner()(*ItemActionsHostedRunnersImagesPartnerRequestBuilder) {
    return NewItemActionsHostedRunnersImagesPartnerRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
