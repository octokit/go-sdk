package repos

import (
    "context"
    i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e "github.com/octokit/go-sdk/github/octokit/models"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadfcdfeac8f91787149ba61da7f66fde5e776c351f03fcb27f040f81408d5f9f "github.com/octokit/go-sdk/github/octokit/repos/item/item/activity"
)

// ItemItemActivityRequestBuilder builds and executes requests for operations under \repos\{owner}\{repo}\activity
type ItemItemActivityRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemItemActivityRequestBuilderGetQueryParameters lists a detailed history of changes to a repository, such as pushes, merges, force pushes, and branch changes, and associates these changes with commits and users.For more information about viewing repository activity,see "[Viewing activity and data for your repository](https://docs.github.com/repositories/viewing-activity-and-data-for-your-repository)."
type ItemItemActivityRequestBuilderGetQueryParameters struct {
    // The activity type to filter by.For example, you can choose to filter by "force_push", to see all force pushes to the repository.
    // Deprecated: This property is deprecated, use activity_typeAsGetActivity_typeQueryParameterType instead
    Activity_type *string `uriparametername:"activity_type"`
    // The activity type to filter by.For example, you can choose to filter by "force_push", to see all force pushes to the repository.
    Activity_typeAsGetActivity_typeQueryParameterType *iadfcdfeac8f91787149ba61da7f66fde5e776c351f03fcb27f040f81408d5f9f.GetActivity_typeQueryParameterType `uriparametername:"activity_type"`
    // The GitHub username to use to filter by the actor who performed the activity.
    Actor *string `uriparametername:"actor"`
    // A cursor, as given in the [Link header](https://docs.github.com/rest/guides/using-pagination-in-the-rest-api#using-link-headers). If specified, the query only searches for results after this cursor.
    After *string `uriparametername:"after"`
    // A cursor, as given in the [Link header](https://docs.github.com/rest/guides/using-pagination-in-the-rest-api#using-link-headers). If specified, the query only searches for results before this cursor.
    Before *string `uriparametername:"before"`
    // The direction to sort the results by.
    // Deprecated: This property is deprecated, use directionAsGetDirectionQueryParameterType instead
    Direction *string `uriparametername:"direction"`
    // The direction to sort the results by.
    DirectionAsGetDirectionQueryParameterType *iadfcdfeac8f91787149ba61da7f66fde5e776c351f03fcb27f040f81408d5f9f.GetDirectionQueryParameterType `uriparametername:"direction"`
    // The number of results per page (max 100).
    Per_page *int32 `uriparametername:"per_page"`
    // The Git reference for the activities you want to list.The `ref` for a branch can be formatted either as `refs/heads/BRANCH_NAME` or `BRANCH_NAME`, where `BRANCH_NAME` is the name of your branch.
    Ref *string `uriparametername:"ref"`
    // The time period to filter by.For example, `day` will filter for activity that occurred in the past 24 hours, and `week` will filter for activity that occurred in the past 7 days (168 hours).
    // Deprecated: This property is deprecated, use time_periodAsGetTime_periodQueryParameterType instead
    Time_period *string `uriparametername:"time_period"`
    // The time period to filter by.For example, `day` will filter for activity that occurred in the past 24 hours, and `week` will filter for activity that occurred in the past 7 days (168 hours).
    Time_periodAsGetTime_periodQueryParameterType *iadfcdfeac8f91787149ba61da7f66fde5e776c351f03fcb27f040f81408d5f9f.GetTime_periodQueryParameterType `uriparametername:"time_period"`
}
// ItemItemActivityRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemItemActivityRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemItemActivityRequestBuilderGetQueryParameters
}
// NewItemItemActivityRequestBuilderInternal instantiates a new ActivityRequestBuilder and sets the default values.
func NewItemItemActivityRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemActivityRequestBuilder) {
    m := &ItemItemActivityRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/repos/{owner}/{repo}/activity{?direction*,per_page*,before*,after*,ref*,actor*,time_period*,activity_type*}", pathParameters),
    }
    return m
}
// NewItemItemActivityRequestBuilder instantiates a new ActivityRequestBuilder and sets the default values.
func NewItemItemActivityRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemActivityRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemItemActivityRequestBuilderInternal(urlParams, requestAdapter)
}
// Get lists a detailed history of changes to a repository, such as pushes, merges, force pushes, and branch changes, and associates these changes with commits and users.For more information about viewing repository activity,see "[Viewing activity and data for your repository](https://docs.github.com/repositories/viewing-activity-and-data-for-your-repository)."
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/repos/repos#list-repository-activities
func (m *ItemItemActivityRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemItemActivityRequestBuilderGetRequestConfiguration)([]i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.Activityable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "422": i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.CreateValidationErrorSimpleFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.SendCollection(ctx, requestInfo, i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.CreateActivityFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    val := make([]i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.Activityable, len(res))
    for i, v := range res {
        if v != nil {
            val[i] = v.(i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.Activityable)
        }
    }
    return val, nil
}
// ToGetRequestInformation lists a detailed history of changes to a repository, such as pushes, merges, force pushes, and branch changes, and associates these changes with commits and users.For more information about viewing repository activity,see "[Viewing activity and data for your repository](https://docs.github.com/repositories/viewing-activity-and-data-for-your-repository)."
func (m *ItemItemActivityRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemItemActivityRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *ItemItemActivityRequestBuilder) WithUrl(rawUrl string)(*ItemItemActivityRequestBuilder) {
    return NewItemItemActivityRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
