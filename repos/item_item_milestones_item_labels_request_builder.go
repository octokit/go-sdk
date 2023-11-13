package repos

import (
    "context"
    i33bcfa37677dfd8b71eaf22ad3bfd140ba1702e7a8418558499c22418bf2fef2 "github.com/octokit/go-sdk/models"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemItemMilestonesItemLabelsRequestBuilder builds and executes requests for operations under \repos\{owner}\{repo}\milestones\{milestone_number}\labels
type ItemItemMilestonesItemLabelsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemItemMilestonesItemLabelsRequestBuilderGetQueryParameters lists labels for issues in a milestone.
type ItemItemMilestonesItemLabelsRequestBuilderGetQueryParameters struct {
    // Page number of the results to fetch.
    Page *int32 `uriparametername:"page"`
    // The number of results per page (max 100).
    Per_page *int32 `uriparametername:"per_page"`
}
// ItemItemMilestonesItemLabelsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemItemMilestonesItemLabelsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemItemMilestonesItemLabelsRequestBuilderGetQueryParameters
}
// NewItemItemMilestonesItemLabelsRequestBuilderInternal instantiates a new LabelsRequestBuilder and sets the default values.
func NewItemItemMilestonesItemLabelsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemMilestonesItemLabelsRequestBuilder) {
    m := &ItemItemMilestonesItemLabelsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/repos/{owner}/{repo}/milestones/{milestone_number}/labels{?per_page*,page*}", pathParameters),
    }
    return m
}
// NewItemItemMilestonesItemLabelsRequestBuilder instantiates a new LabelsRequestBuilder and sets the default values.
func NewItemItemMilestonesItemLabelsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemMilestonesItemLabelsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemItemMilestonesItemLabelsRequestBuilderInternal(urlParams, requestAdapter)
}
// Get lists labels for issues in a milestone.
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/issues/labels#list-labels-for-issues-in-a-milestone
func (m *ItemItemMilestonesItemLabelsRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemItemMilestonesItemLabelsRequestBuilderGetRequestConfiguration)([]i33bcfa37677dfd8b71eaf22ad3bfd140ba1702e7a8418558499c22418bf2fef2.Labelable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.SendCollection(ctx, requestInfo, i33bcfa37677dfd8b71eaf22ad3bfd140ba1702e7a8418558499c22418bf2fef2.CreateLabelFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    val := make([]i33bcfa37677dfd8b71eaf22ad3bfd140ba1702e7a8418558499c22418bf2fef2.Labelable, len(res))
    for i, v := range res {
        if v != nil {
            val[i] = v.(i33bcfa37677dfd8b71eaf22ad3bfd140ba1702e7a8418558499c22418bf2fef2.Labelable)
        }
    }
    return val, nil
}
// ToGetRequestInformation lists labels for issues in a milestone.
func (m *ItemItemMilestonesItemLabelsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemItemMilestonesItemLabelsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers.TryAdd("Accept", "application/json;q=1")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
func (m *ItemItemMilestonesItemLabelsRequestBuilder) WithUrl(rawUrl string)(*ItemItemMilestonesItemLabelsRequestBuilder) {
    return NewItemItemMilestonesItemLabelsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
