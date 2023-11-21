package classrooms

import (
    "context"
    i53ac87e8cb3cc9276228f74d38694a208cacb99bb8ceb705eeae99fb88d4d274 "strconv"
    i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e "github.com/octokit/go-sdk/github/octokit/models"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ClassroomsRequestBuilder builds and executes requests for operations under \classrooms
type ClassroomsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ClassroomsRequestBuilderGetQueryParameters lists GitHub Classroom classrooms for the current user. Classrooms will only be returned if the current user is an administrator of one or more GitHub Classrooms.
type ClassroomsRequestBuilderGetQueryParameters struct {
    // Page number of the results to fetch.
    Page *int32 `uriparametername:"page"`
    // The number of results per page (max 100).
    Per_page *int32 `uriparametername:"per_page"`
}
// ClassroomsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ClassroomsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ClassroomsRequestBuilderGetQueryParameters
}
// ByClassroom_id gets an item from the octokit.classrooms.item collection
// Deprecated: This indexer is deprecated and will be removed in the next major version. Use the one with the typed parameter instead.
func (m *ClassroomsRequestBuilder) ByClassroom_id(classroom_id string)(*WithClassroom_ItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if classroom_id != "" {
        urlTplParams["classroom_id"] = classroom_id
    }
    return NewWithClassroom_ItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// ByClassroom_idInteger gets an item from the octokit.classrooms.item collection
func (m *ClassroomsRequestBuilder) ByClassroom_idInteger(classroom_id int32)(*WithClassroom_ItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    urlTplParams["classroom_id"] = i53ac87e8cb3cc9276228f74d38694a208cacb99bb8ceb705eeae99fb88d4d274.FormatInt(int64(classroom_id), 10)
    return NewWithClassroom_ItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewClassroomsRequestBuilderInternal instantiates a new ClassroomsRequestBuilder and sets the default values.
func NewClassroomsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ClassroomsRequestBuilder) {
    m := &ClassroomsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/classrooms{?page*,per_page*}", pathParameters),
    }
    return m
}
// NewClassroomsRequestBuilder instantiates a new ClassroomsRequestBuilder and sets the default values.
func NewClassroomsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ClassroomsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewClassroomsRequestBuilderInternal(urlParams, requestAdapter)
}
// Get lists GitHub Classroom classrooms for the current user. Classrooms will only be returned if the current user is an administrator of one or more GitHub Classrooms.
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/classroom/classroom#list-classrooms
func (m *ClassroomsRequestBuilder) Get(ctx context.Context, requestConfiguration *ClassroomsRequestBuilderGetRequestConfiguration)([]i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.SimpleClassroomable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.SendCollection(ctx, requestInfo, i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.CreateSimpleClassroomFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    val := make([]i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.SimpleClassroomable, len(res))
    for i, v := range res {
        if v != nil {
            val[i] = v.(i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.SimpleClassroomable)
        }
    }
    return val, nil
}
// ToGetRequestInformation lists GitHub Classroom classrooms for the current user. Classrooms will only be returned if the current user is an administrator of one or more GitHub Classrooms.
func (m *ClassroomsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ClassroomsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
func (m *ClassroomsRequestBuilder) WithUrl(rawUrl string)(*ClassroomsRequestBuilder) {
    return NewClassroomsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
