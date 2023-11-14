package repos

import (
    "context"
    i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e "github.com/octokit/go-sdk/github/octokit/models"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemItemStargazersRequestBuilder builds and executes requests for operations under \repos\{owner}\{repo}\stargazers
type ItemItemStargazersRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemItemStargazersRequestBuilderGetQueryParameters lists the people that have starred the repository.You can also find out _when_ stars were created by passing the following custom [media type](https://docs.github.com/rest/overview/media-types/) via the `Accept` header: `application/vnd.github.star+json`.
type ItemItemStargazersRequestBuilderGetQueryParameters struct {
    // Page number of the results to fetch.
    Page *int32 `uriparametername:"page"`
    // The number of results per page (max 100).
    Per_page *int32 `uriparametername:"per_page"`
}
// ItemItemStargazersRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemItemStargazersRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemItemStargazersRequestBuilderGetQueryParameters
}
// StargazersGetResponse composed type wrapper for classes ItemItemStargazersSimpleUser, ItemItemStargazersStargazer
type StargazersGetResponse struct {
    // Composed type representation for type ItemItemStargazersSimpleUser
    itemItemStargazersSimpleUser ItemItemStargazersSimpleUserable
    // Composed type representation for type ItemItemStargazersStargazer
    itemItemStargazersStargazer ItemItemStargazersStargazerable
}
// NewStargazersGetResponse instantiates a new stargazersGetResponse and sets the default values.
func NewStargazersGetResponse()(*StargazersGetResponse) {
    m := &StargazersGetResponse{
    }
    return m
}
// CreateStargazersGetResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateStargazersGetResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    result := NewStargazersGetResponse()
    if parseNode != nil {
        if val, err := parseNode.GetObjectValue(CreateItemItemStargazersSimpleUserFromDiscriminatorValue); val != nil {
            if err != nil {
                return nil, err
            }
            if cast, ok := val.(ItemItemStargazersSimpleUserable); ok {
                result.SetItemItemStargazersSimpleUser(cast)
            }
        } else if val, err := parseNode.GetObjectValue(CreateItemItemStargazersStargazerFromDiscriminatorValue); val != nil {
            if err != nil {
                return nil, err
            }
            if cast, ok := val.(ItemItemStargazersStargazerable); ok {
                result.SetItemItemStargazersStargazer(cast)
            }
        }
    }
    return result, nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *StargazersGetResponse) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    return make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
}
// GetIsComposedType determines if the current object is a wrapper around a composed type
func (m *StargazersGetResponse) GetIsComposedType()(bool) {
    return true
}
// GetItemItemStargazersSimpleUser gets the ItemItemStargazersSimpleUser property value. Composed type representation for type ItemItemStargazersSimpleUser
func (m *StargazersGetResponse) GetItemItemStargazersSimpleUser()(ItemItemStargazersSimpleUserable) {
    return m.itemItemStargazersSimpleUser
}
// GetItemItemStargazersStargazer gets the ItemItemStargazersStargazer property value. Composed type representation for type ItemItemStargazersStargazer
func (m *StargazersGetResponse) GetItemItemStargazersStargazer()(ItemItemStargazersStargazerable) {
    return m.itemItemStargazersStargazer
}
// Serialize serializes information the current object
func (m *StargazersGetResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetItemItemStargazersSimpleUser() != nil {
        err := writer.WriteObjectValue("", m.GetItemItemStargazersSimpleUser())
        if err != nil {
            return err
        }
    } else if m.GetItemItemStargazersStargazer() != nil {
        err := writer.WriteObjectValue("", m.GetItemItemStargazersStargazer())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetItemItemStargazersSimpleUser sets the ItemItemStargazersSimpleUser property value. Composed type representation for type ItemItemStargazersSimpleUser
func (m *StargazersGetResponse) SetItemItemStargazersSimpleUser(value ItemItemStargazersSimpleUserable)() {
    m.itemItemStargazersSimpleUser = value
}
// SetItemItemStargazersStargazer sets the ItemItemStargazersStargazer property value. Composed type representation for type ItemItemStargazersStargazer
func (m *StargazersGetResponse) SetItemItemStargazersStargazer(value ItemItemStargazersStargazerable)() {
    m.itemItemStargazersStargazer = value
}
// StargazersResponse composed type wrapper for classes ItemItemStargazersSimpleUser, ItemItemStargazersStargazer
type StargazersResponse struct {
    // Composed type representation for type ItemItemStargazersSimpleUser
    itemItemStargazersSimpleUser ItemItemStargazersSimpleUserable
    // Composed type representation for type ItemItemStargazersStargazer
    itemItemStargazersStargazer ItemItemStargazersStargazerable
}
// NewStargazersResponse instantiates a new stargazersResponse and sets the default values.
func NewStargazersResponse()(*StargazersResponse) {
    m := &StargazersResponse{
    }
    return m
}
// CreateStargazersResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateStargazersResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    result := NewStargazersResponse()
    if parseNode != nil {
        if val, err := parseNode.GetObjectValue(CreateItemItemStargazersSimpleUserFromDiscriminatorValue); val != nil {
            if err != nil {
                return nil, err
            }
            if cast, ok := val.(ItemItemStargazersSimpleUserable); ok {
                result.SetItemItemStargazersSimpleUser(cast)
            }
        } else if val, err := parseNode.GetObjectValue(CreateItemItemStargazersStargazerFromDiscriminatorValue); val != nil {
            if err != nil {
                return nil, err
            }
            if cast, ok := val.(ItemItemStargazersStargazerable); ok {
                result.SetItemItemStargazersStargazer(cast)
            }
        }
    }
    return result, nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *StargazersResponse) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    return make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
}
// GetIsComposedType determines if the current object is a wrapper around a composed type
func (m *StargazersResponse) GetIsComposedType()(bool) {
    return true
}
// GetItemItemStargazersSimpleUser gets the ItemItemStargazersSimpleUser property value. Composed type representation for type ItemItemStargazersSimpleUser
func (m *StargazersResponse) GetItemItemStargazersSimpleUser()(ItemItemStargazersSimpleUserable) {
    return m.itemItemStargazersSimpleUser
}
// GetItemItemStargazersStargazer gets the ItemItemStargazersStargazer property value. Composed type representation for type ItemItemStargazersStargazer
func (m *StargazersResponse) GetItemItemStargazersStargazer()(ItemItemStargazersStargazerable) {
    return m.itemItemStargazersStargazer
}
// Serialize serializes information the current object
func (m *StargazersResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetItemItemStargazersSimpleUser() != nil {
        err := writer.WriteObjectValue("", m.GetItemItemStargazersSimpleUser())
        if err != nil {
            return err
        }
    } else if m.GetItemItemStargazersStargazer() != nil {
        err := writer.WriteObjectValue("", m.GetItemItemStargazersStargazer())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetItemItemStargazersSimpleUser sets the ItemItemStargazersSimpleUser property value. Composed type representation for type ItemItemStargazersSimpleUser
func (m *StargazersResponse) SetItemItemStargazersSimpleUser(value ItemItemStargazersSimpleUserable)() {
    m.itemItemStargazersSimpleUser = value
}
// SetItemItemStargazersStargazer sets the ItemItemStargazersStargazer property value. Composed type representation for type ItemItemStargazersStargazer
func (m *StargazersResponse) SetItemItemStargazersStargazer(value ItemItemStargazersStargazerable)() {
    m.itemItemStargazersStargazer = value
}
// StargazersGetResponseable 
type StargazersGetResponseable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetItemItemStargazersSimpleUser()(ItemItemStargazersSimpleUserable)
    GetItemItemStargazersStargazer()(ItemItemStargazersStargazerable)
    SetItemItemStargazersSimpleUser(value ItemItemStargazersSimpleUserable)()
    SetItemItemStargazersStargazer(value ItemItemStargazersStargazerable)()
}
// StargazersResponseable 
type StargazersResponseable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetItemItemStargazersSimpleUser()(ItemItemStargazersSimpleUserable)
    GetItemItemStargazersStargazer()(ItemItemStargazersStargazerable)
    SetItemItemStargazersSimpleUser(value ItemItemStargazersSimpleUserable)()
    SetItemItemStargazersStargazer(value ItemItemStargazersStargazerable)()
}
// NewItemItemStargazersRequestBuilderInternal instantiates a new StargazersRequestBuilder and sets the default values.
func NewItemItemStargazersRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemStargazersRequestBuilder) {
    m := &ItemItemStargazersRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/repos/{owner}/{repo}/stargazers{?per_page*,page*}", pathParameters),
    }
    return m
}
// NewItemItemStargazersRequestBuilder instantiates a new StargazersRequestBuilder and sets the default values.
func NewItemItemStargazersRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemStargazersRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemItemStargazersRequestBuilderInternal(urlParams, requestAdapter)
}
// Get lists the people that have starred the repository.You can also find out _when_ stars were created by passing the following custom [media type](https://docs.github.com/rest/overview/media-types/) via the `Accept` header: `application/vnd.github.star+json`.
// Deprecated: This method is obsolete. Use GetAsStargazersGetResponse instead.
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/activity/starring#list-stargazers
func (m *ItemItemStargazersRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemItemStargazersRequestBuilderGetRequestConfiguration)(StargazersResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "422": i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.CreateValidationErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateStargazersResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(StargazersResponseable), nil
}
// GetAsStargazersGetResponse lists the people that have starred the repository.You can also find out _when_ stars were created by passing the following custom [media type](https://docs.github.com/rest/overview/media-types/) via the `Accept` header: `application/vnd.github.star+json`.
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/activity/starring#list-stargazers
func (m *ItemItemStargazersRequestBuilder) GetAsStargazersGetResponse(ctx context.Context, requestConfiguration *ItemItemStargazersRequestBuilderGetRequestConfiguration)(StargazersGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "422": i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.CreateValidationErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateStargazersGetResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(StargazersGetResponseable), nil
}
// ToGetRequestInformation lists the people that have starred the repository.You can also find out _when_ stars were created by passing the following custom [media type](https://docs.github.com/rest/overview/media-types/) via the `Accept` header: `application/vnd.github.star+json`.
func (m *ItemItemStargazersRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemItemStargazersRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *ItemItemStargazersRequestBuilder) WithUrl(rawUrl string)(*ItemItemStargazersRequestBuilder) {
    return NewItemItemStargazersRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
