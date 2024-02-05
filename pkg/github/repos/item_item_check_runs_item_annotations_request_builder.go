package repos

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6 "github.com/octokit/go-sdk/pkg/github/models"
)

// ItemItemCheckRunsItemAnnotationsRequestBuilder builds and executes requests for operations under \repos\{repos-id}\{Owner-id}\check-runs\{check_run_id}\annotations
type ItemItemCheckRunsItemAnnotationsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemItemCheckRunsItemAnnotationsRequestBuilderGetQueryParameters lists annotations for a check run using the annotation `id`.GitHub Appsmust have the `checks:read` permission on a private repository or pull access toa public repository to get annotations for a check run. OAuth apps and authenticatedusers must have the `repo` scope to get annotations for a check run in a privaterepository.
type ItemItemCheckRunsItemAnnotationsRequestBuilderGetQueryParameters struct {
    // The page number of the results to fetch. For more information, see "[Using pagination in the REST API](https://docs.github.com/rest/using-the-rest-api/using-pagination-in-the-rest-api)."
    Page *int32 `uriparametername:"page"`
    // The number of results per page (max 100). For more information, see "[Using pagination in the REST API](https://docs.github.com/rest/using-the-rest-api/using-pagination-in-the-rest-api)."
    Per_page *int32 `uriparametername:"per_page"`
}
// NewItemItemCheckRunsItemAnnotationsRequestBuilderInternal instantiates a new AnnotationsRequestBuilder and sets the default values.
func NewItemItemCheckRunsItemAnnotationsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemCheckRunsItemAnnotationsRequestBuilder) {
    m := &ItemItemCheckRunsItemAnnotationsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/repos/{repos%2Did}/{Owner%2Did}/check-runs/{check_run_id}/annotations{?page*,per_page*}", pathParameters),
    }
    return m
}
// NewItemItemCheckRunsItemAnnotationsRequestBuilder instantiates a new AnnotationsRequestBuilder and sets the default values.
func NewItemItemCheckRunsItemAnnotationsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemCheckRunsItemAnnotationsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemItemCheckRunsItemAnnotationsRequestBuilderInternal(urlParams, requestAdapter)
}
// Get lists annotations for a check run using the annotation `id`.GitHub Appsmust have the `checks:read` permission on a private repository or pull access toa public repository to get annotations for a check run. OAuth apps and authenticatedusers must have the `repo` scope to get annotations for a check run in a privaterepository.
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/checks/runs#list-check-run-annotations
func (m *ItemItemCheckRunsItemAnnotationsRequestBuilder) Get(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[ItemItemCheckRunsItemAnnotationsRequestBuilderGetQueryParameters])([]i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CheckAnnotationable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.SendCollection(ctx, requestInfo, i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CreateCheckAnnotationFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    val := make([]i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CheckAnnotationable, len(res))
    for i, v := range res {
        if v != nil {
            val[i] = v.(i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CheckAnnotationable)
        }
    }
    return val, nil
}
// ToGetRequestInformation lists annotations for a check run using the annotation `id`.GitHub Appsmust have the `checks:read` permission on a private repository or pull access toa public repository to get annotations for a check run. OAuth apps and authenticatedusers must have the `repo` scope to get annotations for a check run in a privaterepository.
func (m *ItemItemCheckRunsItemAnnotationsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[ItemItemCheckRunsItemAnnotationsRequestBuilderGetQueryParameters])(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ConfigureRequestInformation(requestInfo, requestConfiguration)
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
func (m *ItemItemCheckRunsItemAnnotationsRequestBuilder) WithUrl(rawUrl string)(*ItemItemCheckRunsItemAnnotationsRequestBuilder) {
    return NewItemItemCheckRunsItemAnnotationsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
