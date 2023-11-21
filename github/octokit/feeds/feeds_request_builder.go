package feeds

import (
    "context"
    i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e "github.com/octokit/go-sdk/github/octokit/models"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// FeedsRequestBuilder builds and executes requests for operations under \feeds
type FeedsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// FeedsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type FeedsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewFeedsRequestBuilderInternal instantiates a new FeedsRequestBuilder and sets the default values.
func NewFeedsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*FeedsRequestBuilder) {
    m := &FeedsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/feeds", pathParameters),
    }
    return m
}
// NewFeedsRequestBuilder instantiates a new FeedsRequestBuilder and sets the default values.
func NewFeedsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*FeedsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewFeedsRequestBuilderInternal(urlParams, requestAdapter)
}
// Get gitHub provides several timeline resources in [Atom](http://en.wikipedia.org/wiki/Atom_(standard)) format. The Feeds API lists all the feeds available to the authenticated user:*   **Timeline**: The GitHub global public timeline*   **User**: The public timeline for any user, using [URI template](https://docs.github.com/rest/overview/resources-in-the-rest-api#hypermedia)*   **Current user public**: The public timeline for the authenticated user*   **Current user**: The private timeline for the authenticated user*   **Current user actor**: The private timeline for activity created by the authenticated user*   **Current user organizations**: The private timeline for the organizations the authenticated user is a member of.*   **Security advisories**: A collection of public announcements that provide information about security-related vulnerabilities in software on GitHub.**Note**: Private feeds are only returned when [authenticating via Basic Auth](https://docs.github.com/rest/overview/other-authentication-methods#basic-authentication) since current feed URIs use the older, non revocable auth tokens.
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/activity/feeds#get-feeds
func (m *FeedsRequestBuilder) Get(ctx context.Context, requestConfiguration *FeedsRequestBuilderGetRequestConfiguration)(i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.Feedable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.CreateFeedFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.Feedable), nil
}
// ToGetRequestInformation gitHub provides several timeline resources in [Atom](http://en.wikipedia.org/wiki/Atom_(standard)) format. The Feeds API lists all the feeds available to the authenticated user:*   **Timeline**: The GitHub global public timeline*   **User**: The public timeline for any user, using [URI template](https://docs.github.com/rest/overview/resources-in-the-rest-api#hypermedia)*   **Current user public**: The public timeline for the authenticated user*   **Current user**: The private timeline for the authenticated user*   **Current user actor**: The private timeline for activity created by the authenticated user*   **Current user organizations**: The private timeline for the organizations the authenticated user is a member of.*   **Security advisories**: A collection of public announcements that provide information about security-related vulnerabilities in software on GitHub.**Note**: Private feeds are only returned when [authenticating via Basic Auth](https://docs.github.com/rest/overview/other-authentication-methods#basic-authentication) since current feed URIs use the older, non revocable auth tokens.
func (m *FeedsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *FeedsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *FeedsRequestBuilder) WithUrl(rawUrl string)(*FeedsRequestBuilder) {
    return NewFeedsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
