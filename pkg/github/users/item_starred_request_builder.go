package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    i79bee196f908e2ec3c0a95608a335c04823e48e52984cbf48aa57b354fe8088c "github.com/octokit/go-sdk/pkg/github/users/item/starred"
)

// ItemStarredRequestBuilder builds and executes requests for operations under \users\{username}\starred
type ItemStarredRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemStarredRequestBuilderGetQueryParameters lists repositories a user has starred.You can also find out _when_ stars were created by passing the following custom [media type](https://docs.github.com/rest/overview/media-types/) via the `Accept` header: `application/vnd.github.star+json`.
type ItemStarredRequestBuilderGetQueryParameters struct {
    // The direction to sort the results by.
    // Deprecated: This property is deprecated, use directionAsGetDirectionQueryParameterType instead
    Direction *string `uriparametername:"direction"`
    // The direction to sort the results by.
    DirectionAsGetDirectionQueryParameterType *i79bee196f908e2ec3c0a95608a335c04823e48e52984cbf48aa57b354fe8088c.GetDirectionQueryParameterType `uriparametername:"direction"`
    // Page number of the results to fetch.
    Page *int32 `uriparametername:"page"`
    // The number of results per page (max 100).
    Per_page *int32 `uriparametername:"per_page"`
    // The property to sort the results by. `created` means when the repository was starred. `updated` means when the repository was last pushed to.
    // Deprecated: This property is deprecated, use sortAsGetSortQueryParameterType instead
    Sort *string `uriparametername:"sort"`
    // The property to sort the results by. `created` means when the repository was starred. `updated` means when the repository was last pushed to.
    SortAsGetSortQueryParameterType *i79bee196f908e2ec3c0a95608a335c04823e48e52984cbf48aa57b354fe8088c.GetSortQueryParameterType `uriparametername:"sort"`
}
// ItemStarredRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemStarredRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemStarredRequestBuilderGetQueryParameters
}
// StarredGetResponse composed type wrapper for classes ItemStarredRepository, ItemStarredRepository
type StarredGetResponse struct {
    // Composed type representation for type ItemStarredRepository
    itemStarredRepository []ItemStarredRepositoryable
    // Composed type representation for type ItemStarredRepository
    starredGetResponseItemStarredRepository []ItemStarredRepositoryable
}
// NewStarredGetResponse instantiates a new starredGetResponse and sets the default values.
func NewStarredGetResponse()(*StarredGetResponse) {
    m := &StarredGetResponse{
    }
    return m
}
// CreateStarredGetResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateStarredGetResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    result := NewStarredGetResponse()
    if parseNode != nil {
        if val, err := parseNode.GetCollectionOfObjectValues(CreateItemStarredRepositoryFromDiscriminatorValue); val != nil {
            if err != nil {
                return nil, err
            }
            cast := make([]ItemStarredRepositoryable, len(val))
            for i, v := range val {
                if v != nil {
                    cast[i] = v.(ItemStarredRepositoryable)
                }
            }
            result.SetItemStarredRepository(cast)
        } else if val, err := parseNode.GetCollectionOfObjectValues(CreateItemStarredRepositoryFromDiscriminatorValue); val != nil {
            if err != nil {
                return nil, err
            }
            cast := make([]ItemStarredRepositoryable, len(val))
            for i, v := range val {
                if v != nil {
                    cast[i] = v.(ItemStarredRepositoryable)
                }
            }
            result.SetStarredGetResponseItemStarredRepository(cast)
        }
    }
    return result, nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *StarredGetResponse) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    return make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
}
// GetIsComposedType determines if the current object is a wrapper around a composed type
func (m *StarredGetResponse) GetIsComposedType()(bool) {
    return true
}
// GetItemStarredRepository gets the ItemStarredRepository property value. Composed type representation for type ItemStarredRepository
func (m *StarredGetResponse) GetItemStarredRepository()([]ItemStarredRepositoryable) {
    return m.itemStarredRepository
}
// GetStarredGetResponseItemStarredRepository gets the ItemStarredRepository property value. Composed type representation for type ItemStarredRepository
func (m *StarredGetResponse) GetStarredGetResponseItemStarredRepository()([]ItemStarredRepositoryable) {
    return m.starredGetResponseItemStarredRepository
}
// Serialize serializes information the current object
func (m *StarredGetResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetItemStarredRepository() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetItemStarredRepository()))
        for i, v := range m.GetItemStarredRepository() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("", cast)
        if err != nil {
            return err
        }
    } else if m.GetStarredGetResponseItemStarredRepository() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetStarredGetResponseItemStarredRepository()))
        for i, v := range m.GetStarredGetResponseItemStarredRepository() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetItemStarredRepository sets the ItemStarredRepository property value. Composed type representation for type ItemStarredRepository
func (m *StarredGetResponse) SetItemStarredRepository(value []ItemStarredRepositoryable)() {
    m.itemStarredRepository = value
}
// SetStarredGetResponseItemStarredRepository sets the ItemStarredRepository property value. Composed type representation for type ItemStarredRepository
func (m *StarredGetResponse) SetStarredGetResponseItemStarredRepository(value []ItemStarredRepositoryable)() {
    m.starredGetResponseItemStarredRepository = value
}
// StarredResponse composed type wrapper for classes ItemStarredRepository, ItemStarredRepository
type StarredResponse struct {
    // Composed type representation for type ItemStarredRepository
    itemStarredRepository []ItemStarredRepositoryable
    // Composed type representation for type ItemStarredRepository
    starredResponseItemStarredRepository []ItemStarredRepositoryable
}
// NewStarredResponse instantiates a new starredResponse and sets the default values.
func NewStarredResponse()(*StarredResponse) {
    m := &StarredResponse{
    }
    return m
}
// CreateStarredResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateStarredResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    result := NewStarredResponse()
    if parseNode != nil {
        if val, err := parseNode.GetCollectionOfObjectValues(CreateItemStarredRepositoryFromDiscriminatorValue); val != nil {
            if err != nil {
                return nil, err
            }
            cast := make([]ItemStarredRepositoryable, len(val))
            for i, v := range val {
                if v != nil {
                    cast[i] = v.(ItemStarredRepositoryable)
                }
            }
            result.SetItemStarredRepository(cast)
        } else if val, err := parseNode.GetCollectionOfObjectValues(CreateItemStarredRepositoryFromDiscriminatorValue); val != nil {
            if err != nil {
                return nil, err
            }
            cast := make([]ItemStarredRepositoryable, len(val))
            for i, v := range val {
                if v != nil {
                    cast[i] = v.(ItemStarredRepositoryable)
                }
            }
            result.SetStarredResponseItemStarredRepository(cast)
        }
    }
    return result, nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *StarredResponse) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    return make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
}
// GetIsComposedType determines if the current object is a wrapper around a composed type
func (m *StarredResponse) GetIsComposedType()(bool) {
    return true
}
// GetItemStarredRepository gets the ItemStarredRepository property value. Composed type representation for type ItemStarredRepository
func (m *StarredResponse) GetItemStarredRepository()([]ItemStarredRepositoryable) {
    return m.itemStarredRepository
}
// GetStarredResponseItemStarredRepository gets the ItemStarredRepository property value. Composed type representation for type ItemStarredRepository
func (m *StarredResponse) GetStarredResponseItemStarredRepository()([]ItemStarredRepositoryable) {
    return m.starredResponseItemStarredRepository
}
// Serialize serializes information the current object
func (m *StarredResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetItemStarredRepository() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetItemStarredRepository()))
        for i, v := range m.GetItemStarredRepository() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("", cast)
        if err != nil {
            return err
        }
    } else if m.GetStarredResponseItemStarredRepository() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetStarredResponseItemStarredRepository()))
        for i, v := range m.GetStarredResponseItemStarredRepository() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetItemStarredRepository sets the ItemStarredRepository property value. Composed type representation for type ItemStarredRepository
func (m *StarredResponse) SetItemStarredRepository(value []ItemStarredRepositoryable)() {
    m.itemStarredRepository = value
}
// SetStarredResponseItemStarredRepository sets the ItemStarredRepository property value. Composed type representation for type ItemStarredRepository
func (m *StarredResponse) SetStarredResponseItemStarredRepository(value []ItemStarredRepositoryable)() {
    m.starredResponseItemStarredRepository = value
}
// StarredGetResponseable 
type StarredGetResponseable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetItemStarredRepository()([]ItemStarredRepositoryable)
    GetStarredGetResponseItemStarredRepository()([]ItemStarredRepositoryable)
    SetItemStarredRepository(value []ItemStarredRepositoryable)()
    SetStarredGetResponseItemStarredRepository(value []ItemStarredRepositoryable)()
}
// StarredResponseable 
type StarredResponseable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetItemStarredRepository()([]ItemStarredRepositoryable)
    GetStarredResponseItemStarredRepository()([]ItemStarredRepositoryable)
    SetItemStarredRepository(value []ItemStarredRepositoryable)()
    SetStarredResponseItemStarredRepository(value []ItemStarredRepositoryable)()
}
// NewItemStarredRequestBuilderInternal instantiates a new StarredRequestBuilder and sets the default values.
func NewItemStarredRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemStarredRequestBuilder) {
    m := &ItemStarredRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/users/{username}/starred{?sort*,direction*,per_page*,page*}", pathParameters),
    }
    return m
}
// NewItemStarredRequestBuilder instantiates a new StarredRequestBuilder and sets the default values.
func NewItemStarredRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemStarredRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemStarredRequestBuilderInternal(urlParams, requestAdapter)
}
// Get lists repositories a user has starred.You can also find out _when_ stars were created by passing the following custom [media type](https://docs.github.com/rest/overview/media-types/) via the `Accept` header: `application/vnd.github.star+json`.
// Deprecated: This method is obsolete. Use GetAsStarredGetResponse instead.
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/activity/starring#list-repositories-starred-by-a-user
func (m *ItemStarredRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemStarredRequestBuilderGetRequestConfiguration)(StarredResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateStarredResponseFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(StarredResponseable), nil
}
// GetAsStarredGetResponse lists repositories a user has starred.You can also find out _when_ stars were created by passing the following custom [media type](https://docs.github.com/rest/overview/media-types/) via the `Accept` header: `application/vnd.github.star+json`.
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/activity/starring#list-repositories-starred-by-a-user
func (m *ItemStarredRequestBuilder) GetAsStarredGetResponse(ctx context.Context, requestConfiguration *ItemStarredRequestBuilderGetRequestConfiguration)(StarredGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateStarredGetResponseFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(StarredGetResponseable), nil
}
// ToGetRequestInformation lists repositories a user has starred.You can also find out _when_ stars were created by passing the following custom [media type](https://docs.github.com/rest/overview/media-types/) via the `Accept` header: `application/vnd.github.star+json`.
func (m *ItemStarredRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemStarredRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
func (m *ItemStarredRequestBuilder) WithUrl(rawUrl string)(*ItemStarredRequestBuilder) {
    return NewItemStarredRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
