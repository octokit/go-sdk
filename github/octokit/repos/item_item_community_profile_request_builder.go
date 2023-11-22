package repos

import (
    "context"
    i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e "github.com/octokit/go-sdk/github/octokit/models"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemItemCommunityProfileRequestBuilder builds and executes requests for operations under \repos\{owner}\{repo}\community\profile
type ItemItemCommunityProfileRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemItemCommunityProfileRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemItemCommunityProfileRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemItemCommunityProfileRequestBuilderInternal instantiates a new ProfileRequestBuilder and sets the default values.
func NewItemItemCommunityProfileRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemCommunityProfileRequestBuilder) {
    m := &ItemItemCommunityProfileRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/repos/{owner}/{repo}/community/profile", pathParameters),
    }
    return m
}
// NewItemItemCommunityProfileRequestBuilder instantiates a new ProfileRequestBuilder and sets the default values.
func NewItemItemCommunityProfileRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemCommunityProfileRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemItemCommunityProfileRequestBuilderInternal(urlParams, requestAdapter)
}
// Get returns all community profile metrics for a repository. The repository cannot be a fork.The returned metrics include an overall health score, the repository description, the presence of documentation, thedetected code of conduct, the detected license, and the presence of ISSUE\_TEMPLATE, PULL\_REQUEST\_TEMPLATE,README, and CONTRIBUTING files.The `health_percentage` score is defined as a percentage of how many ofthe recommended community health files are present. For more information, see"[About community profiles for public repositories](https://docs.github.com/communities/setting-up-your-project-for-healthy-contributions/about-community-profiles-for-public-repositories)."`content_reports_enabled` is only returned for organization-owned repositories.
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/metrics/community#get-community-profile-metrics
func (m *ItemItemCommunityProfileRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemItemCommunityProfileRequestBuilderGetRequestConfiguration)(i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.CommunityProfileable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.CreateCommunityProfileFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.CommunityProfileable), nil
}
// ToGetRequestInformation returns all community profile metrics for a repository. The repository cannot be a fork.The returned metrics include an overall health score, the repository description, the presence of documentation, thedetected code of conduct, the detected license, and the presence of ISSUE\_TEMPLATE, PULL\_REQUEST\_TEMPLATE,README, and CONTRIBUTING files.The `health_percentage` score is defined as a percentage of how many ofthe recommended community health files are present. For more information, see"[About community profiles for public repositories](https://docs.github.com/communities/setting-up-your-project-for-healthy-contributions/about-community-profiles-for-public-repositories)."`content_reports_enabled` is only returned for organization-owned repositories.
func (m *ItemItemCommunityProfileRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemItemCommunityProfileRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
func (m *ItemItemCommunityProfileRequestBuilder) WithUrl(rawUrl string)(*ItemItemCommunityProfileRequestBuilder) {
    return NewItemItemCommunityProfileRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
