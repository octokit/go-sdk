package repos

import (
    "context"
    i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e "github.com/octokit/go-sdk/github/octokit/models"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemItemBranchesItemProtectionRestrictionsAppsRequestBuilder builds and executes requests for operations under \repos\{owner}\{repo}\branches\{branch}\protection\restrictions\apps
type ItemItemBranchesItemProtectionRestrictionsAppsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// AppsDeleteRequestBody composed type wrapper for classes ItemItemBranchesItemProtectionRestrictionsAppsDeleteRequestBodyMember1, string
type AppsDeleteRequestBody struct {
    // Composed type representation for type ItemItemBranchesItemProtectionRestrictionsAppsDeleteRequestBodyMember1
    appsDeleteRequestBodyItemItemBranchesItemProtectionRestrictionsAppsDeleteRequestBodyMember1 ItemItemBranchesItemProtectionRestrictionsAppsDeleteRequestBodyMember1able
    // Composed type representation for type string
    appsDeleteRequestBodyString *string
    // Composed type representation for type ItemItemBranchesItemProtectionRestrictionsAppsDeleteRequestBodyMember1
    itemItemBranchesItemProtectionRestrictionsAppsDeleteRequestBodyMember1 ItemItemBranchesItemProtectionRestrictionsAppsDeleteRequestBodyMember1able
    // Composed type representation for type string
    string *string
}
// NewAppsDeleteRequestBody instantiates a new appsDeleteRequestBody and sets the default values.
func NewAppsDeleteRequestBody()(*AppsDeleteRequestBody) {
    m := &AppsDeleteRequestBody{
    }
    return m
}
// CreateAppsDeleteRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAppsDeleteRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    result := NewAppsDeleteRequestBody()
    if parseNode != nil {
        mappingValueNode, err := parseNode.GetChildNode("")
        if err != nil {
            return nil, err
        }
        if mappingValueNode != nil {
            mappingValue, err := mappingValueNode.GetStringValue()
            if err != nil {
                return nil, err
            }
            if mappingValue != nil {
            }
        }
    }
    if val, err := parseNode.GetStringValue(); val != nil {
        if err != nil {
            return nil, err
        }
        result.SetAppsDeleteRequestBodyString(val)
    } else if val, err := parseNode.GetStringValue(); val != nil {
        if err != nil {
            return nil, err
        }
        result.SetString(val)
    }
    return result, nil
}
// GetAppsDeleteRequestBodyItemItemBranchesItemProtectionRestrictionsAppsDeleteRequestBodyMember1 gets the ItemItemBranchesItemProtectionRestrictionsAppsDeleteRequestBodyMember1 property value. Composed type representation for type ItemItemBranchesItemProtectionRestrictionsAppsDeleteRequestBodyMember1
func (m *AppsDeleteRequestBody) GetAppsDeleteRequestBodyItemItemBranchesItemProtectionRestrictionsAppsDeleteRequestBodyMember1()(ItemItemBranchesItemProtectionRestrictionsAppsDeleteRequestBodyMember1able) {
    return m.appsDeleteRequestBodyItemItemBranchesItemProtectionRestrictionsAppsDeleteRequestBodyMember1
}
// GetAppsDeleteRequestBodyString gets the string property value. Composed type representation for type string
func (m *AppsDeleteRequestBody) GetAppsDeleteRequestBodyString()(*string) {
    return m.appsDeleteRequestBodyString
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AppsDeleteRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    return make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
}
// GetIsComposedType determines if the current object is a wrapper around a composed type
func (m *AppsDeleteRequestBody) GetIsComposedType()(bool) {
    return true
}
// GetItemItemBranchesItemProtectionRestrictionsAppsDeleteRequestBodyMember1 gets the ItemItemBranchesItemProtectionRestrictionsAppsDeleteRequestBodyMember1 property value. Composed type representation for type ItemItemBranchesItemProtectionRestrictionsAppsDeleteRequestBodyMember1
func (m *AppsDeleteRequestBody) GetItemItemBranchesItemProtectionRestrictionsAppsDeleteRequestBodyMember1()(ItemItemBranchesItemProtectionRestrictionsAppsDeleteRequestBodyMember1able) {
    return m.itemItemBranchesItemProtectionRestrictionsAppsDeleteRequestBodyMember1
}
// GetString gets the string property value. Composed type representation for type string
func (m *AppsDeleteRequestBody) GetString()(*string) {
    return m.string
}
// Serialize serializes information the current object
func (m *AppsDeleteRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetAppsDeleteRequestBodyItemItemBranchesItemProtectionRestrictionsAppsDeleteRequestBodyMember1() != nil {
        err := writer.WriteObjectValue("", m.GetAppsDeleteRequestBodyItemItemBranchesItemProtectionRestrictionsAppsDeleteRequestBodyMember1())
        if err != nil {
            return err
        }
    } else if m.GetItemItemBranchesItemProtectionRestrictionsAppsDeleteRequestBodyMember1() != nil {
        err := writer.WriteObjectValue("", m.GetItemItemBranchesItemProtectionRestrictionsAppsDeleteRequestBodyMember1())
        if err != nil {
            return err
        }
    } else if m.GetAppsDeleteRequestBodyString() != nil {
        err := writer.WriteStringValue("", m.GetAppsDeleteRequestBodyString())
        if err != nil {
            return err
        }
    } else if m.GetString() != nil {
        err := writer.WriteStringValue("", m.GetString())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAppsDeleteRequestBodyItemItemBranchesItemProtectionRestrictionsAppsDeleteRequestBodyMember1 sets the ItemItemBranchesItemProtectionRestrictionsAppsDeleteRequestBodyMember1 property value. Composed type representation for type ItemItemBranchesItemProtectionRestrictionsAppsDeleteRequestBodyMember1
func (m *AppsDeleteRequestBody) SetAppsDeleteRequestBodyItemItemBranchesItemProtectionRestrictionsAppsDeleteRequestBodyMember1(value ItemItemBranchesItemProtectionRestrictionsAppsDeleteRequestBodyMember1able)() {
    m.appsDeleteRequestBodyItemItemBranchesItemProtectionRestrictionsAppsDeleteRequestBodyMember1 = value
}
// SetAppsDeleteRequestBodyString sets the string property value. Composed type representation for type string
func (m *AppsDeleteRequestBody) SetAppsDeleteRequestBodyString(value *string)() {
    m.appsDeleteRequestBodyString = value
}
// SetItemItemBranchesItemProtectionRestrictionsAppsDeleteRequestBodyMember1 sets the ItemItemBranchesItemProtectionRestrictionsAppsDeleteRequestBodyMember1 property value. Composed type representation for type ItemItemBranchesItemProtectionRestrictionsAppsDeleteRequestBodyMember1
func (m *AppsDeleteRequestBody) SetItemItemBranchesItemProtectionRestrictionsAppsDeleteRequestBodyMember1(value ItemItemBranchesItemProtectionRestrictionsAppsDeleteRequestBodyMember1able)() {
    m.itemItemBranchesItemProtectionRestrictionsAppsDeleteRequestBodyMember1 = value
}
// SetString sets the string property value. Composed type representation for type string
func (m *AppsDeleteRequestBody) SetString(value *string)() {
    m.string = value
}
// AppsPostRequestBody composed type wrapper for classes ItemItemBranchesItemProtectionRestrictionsAppsPostRequestBodyMember1, string
type AppsPostRequestBody struct {
    // Composed type representation for type ItemItemBranchesItemProtectionRestrictionsAppsPostRequestBodyMember1
    appsPostRequestBodyItemItemBranchesItemProtectionRestrictionsAppsPostRequestBodyMember1 ItemItemBranchesItemProtectionRestrictionsAppsPostRequestBodyMember1able
    // Composed type representation for type string
    appsPostRequestBodyString *string
    // Composed type representation for type ItemItemBranchesItemProtectionRestrictionsAppsPostRequestBodyMember1
    itemItemBranchesItemProtectionRestrictionsAppsPostRequestBodyMember1 ItemItemBranchesItemProtectionRestrictionsAppsPostRequestBodyMember1able
    // Composed type representation for type string
    string *string
}
// NewAppsPostRequestBody instantiates a new appsPostRequestBody and sets the default values.
func NewAppsPostRequestBody()(*AppsPostRequestBody) {
    m := &AppsPostRequestBody{
    }
    return m
}
// CreateAppsPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAppsPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    result := NewAppsPostRequestBody()
    if parseNode != nil {
        mappingValueNode, err := parseNode.GetChildNode("")
        if err != nil {
            return nil, err
        }
        if mappingValueNode != nil {
            mappingValue, err := mappingValueNode.GetStringValue()
            if err != nil {
                return nil, err
            }
            if mappingValue != nil {
            }
        }
    }
    if val, err := parseNode.GetStringValue(); val != nil {
        if err != nil {
            return nil, err
        }
        result.SetAppsPostRequestBodyString(val)
    } else if val, err := parseNode.GetStringValue(); val != nil {
        if err != nil {
            return nil, err
        }
        result.SetString(val)
    }
    return result, nil
}
// GetAppsPostRequestBodyItemItemBranchesItemProtectionRestrictionsAppsPostRequestBodyMember1 gets the ItemItemBranchesItemProtectionRestrictionsAppsPostRequestBodyMember1 property value. Composed type representation for type ItemItemBranchesItemProtectionRestrictionsAppsPostRequestBodyMember1
func (m *AppsPostRequestBody) GetAppsPostRequestBodyItemItemBranchesItemProtectionRestrictionsAppsPostRequestBodyMember1()(ItemItemBranchesItemProtectionRestrictionsAppsPostRequestBodyMember1able) {
    return m.appsPostRequestBodyItemItemBranchesItemProtectionRestrictionsAppsPostRequestBodyMember1
}
// GetAppsPostRequestBodyString gets the string property value. Composed type representation for type string
func (m *AppsPostRequestBody) GetAppsPostRequestBodyString()(*string) {
    return m.appsPostRequestBodyString
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AppsPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    return make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
}
// GetIsComposedType determines if the current object is a wrapper around a composed type
func (m *AppsPostRequestBody) GetIsComposedType()(bool) {
    return true
}
// GetItemItemBranchesItemProtectionRestrictionsAppsPostRequestBodyMember1 gets the ItemItemBranchesItemProtectionRestrictionsAppsPostRequestBodyMember1 property value. Composed type representation for type ItemItemBranchesItemProtectionRestrictionsAppsPostRequestBodyMember1
func (m *AppsPostRequestBody) GetItemItemBranchesItemProtectionRestrictionsAppsPostRequestBodyMember1()(ItemItemBranchesItemProtectionRestrictionsAppsPostRequestBodyMember1able) {
    return m.itemItemBranchesItemProtectionRestrictionsAppsPostRequestBodyMember1
}
// GetString gets the string property value. Composed type representation for type string
func (m *AppsPostRequestBody) GetString()(*string) {
    return m.string
}
// Serialize serializes information the current object
func (m *AppsPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetAppsPostRequestBodyItemItemBranchesItemProtectionRestrictionsAppsPostRequestBodyMember1() != nil {
        err := writer.WriteObjectValue("", m.GetAppsPostRequestBodyItemItemBranchesItemProtectionRestrictionsAppsPostRequestBodyMember1())
        if err != nil {
            return err
        }
    } else if m.GetItemItemBranchesItemProtectionRestrictionsAppsPostRequestBodyMember1() != nil {
        err := writer.WriteObjectValue("", m.GetItemItemBranchesItemProtectionRestrictionsAppsPostRequestBodyMember1())
        if err != nil {
            return err
        }
    } else if m.GetAppsPostRequestBodyString() != nil {
        err := writer.WriteStringValue("", m.GetAppsPostRequestBodyString())
        if err != nil {
            return err
        }
    } else if m.GetString() != nil {
        err := writer.WriteStringValue("", m.GetString())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAppsPostRequestBodyItemItemBranchesItemProtectionRestrictionsAppsPostRequestBodyMember1 sets the ItemItemBranchesItemProtectionRestrictionsAppsPostRequestBodyMember1 property value. Composed type representation for type ItemItemBranchesItemProtectionRestrictionsAppsPostRequestBodyMember1
func (m *AppsPostRequestBody) SetAppsPostRequestBodyItemItemBranchesItemProtectionRestrictionsAppsPostRequestBodyMember1(value ItemItemBranchesItemProtectionRestrictionsAppsPostRequestBodyMember1able)() {
    m.appsPostRequestBodyItemItemBranchesItemProtectionRestrictionsAppsPostRequestBodyMember1 = value
}
// SetAppsPostRequestBodyString sets the string property value. Composed type representation for type string
func (m *AppsPostRequestBody) SetAppsPostRequestBodyString(value *string)() {
    m.appsPostRequestBodyString = value
}
// SetItemItemBranchesItemProtectionRestrictionsAppsPostRequestBodyMember1 sets the ItemItemBranchesItemProtectionRestrictionsAppsPostRequestBodyMember1 property value. Composed type representation for type ItemItemBranchesItemProtectionRestrictionsAppsPostRequestBodyMember1
func (m *AppsPostRequestBody) SetItemItemBranchesItemProtectionRestrictionsAppsPostRequestBodyMember1(value ItemItemBranchesItemProtectionRestrictionsAppsPostRequestBodyMember1able)() {
    m.itemItemBranchesItemProtectionRestrictionsAppsPostRequestBodyMember1 = value
}
// SetString sets the string property value. Composed type representation for type string
func (m *AppsPostRequestBody) SetString(value *string)() {
    m.string = value
}
// AppsPutRequestBody composed type wrapper for classes ItemItemBranchesItemProtectionRestrictionsAppsPutRequestBodyMember1, string
type AppsPutRequestBody struct {
    // Composed type representation for type ItemItemBranchesItemProtectionRestrictionsAppsPutRequestBodyMember1
    appsPutRequestBodyItemItemBranchesItemProtectionRestrictionsAppsPutRequestBodyMember1 ItemItemBranchesItemProtectionRestrictionsAppsPutRequestBodyMember1able
    // Composed type representation for type string
    appsPutRequestBodyString *string
    // Composed type representation for type ItemItemBranchesItemProtectionRestrictionsAppsPutRequestBodyMember1
    itemItemBranchesItemProtectionRestrictionsAppsPutRequestBodyMember1 ItemItemBranchesItemProtectionRestrictionsAppsPutRequestBodyMember1able
    // Composed type representation for type string
    string *string
}
// NewAppsPutRequestBody instantiates a new appsPutRequestBody and sets the default values.
func NewAppsPutRequestBody()(*AppsPutRequestBody) {
    m := &AppsPutRequestBody{
    }
    return m
}
// CreateAppsPutRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAppsPutRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    result := NewAppsPutRequestBody()
    if parseNode != nil {
        mappingValueNode, err := parseNode.GetChildNode("")
        if err != nil {
            return nil, err
        }
        if mappingValueNode != nil {
            mappingValue, err := mappingValueNode.GetStringValue()
            if err != nil {
                return nil, err
            }
            if mappingValue != nil {
            }
        }
    }
    if val, err := parseNode.GetStringValue(); val != nil {
        if err != nil {
            return nil, err
        }
        result.SetAppsPutRequestBodyString(val)
    } else if val, err := parseNode.GetStringValue(); val != nil {
        if err != nil {
            return nil, err
        }
        result.SetString(val)
    }
    return result, nil
}
// GetAppsPutRequestBodyItemItemBranchesItemProtectionRestrictionsAppsPutRequestBodyMember1 gets the ItemItemBranchesItemProtectionRestrictionsAppsPutRequestBodyMember1 property value. Composed type representation for type ItemItemBranchesItemProtectionRestrictionsAppsPutRequestBodyMember1
func (m *AppsPutRequestBody) GetAppsPutRequestBodyItemItemBranchesItemProtectionRestrictionsAppsPutRequestBodyMember1()(ItemItemBranchesItemProtectionRestrictionsAppsPutRequestBodyMember1able) {
    return m.appsPutRequestBodyItemItemBranchesItemProtectionRestrictionsAppsPutRequestBodyMember1
}
// GetAppsPutRequestBodyString gets the string property value. Composed type representation for type string
func (m *AppsPutRequestBody) GetAppsPutRequestBodyString()(*string) {
    return m.appsPutRequestBodyString
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AppsPutRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    return make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
}
// GetIsComposedType determines if the current object is a wrapper around a composed type
func (m *AppsPutRequestBody) GetIsComposedType()(bool) {
    return true
}
// GetItemItemBranchesItemProtectionRestrictionsAppsPutRequestBodyMember1 gets the ItemItemBranchesItemProtectionRestrictionsAppsPutRequestBodyMember1 property value. Composed type representation for type ItemItemBranchesItemProtectionRestrictionsAppsPutRequestBodyMember1
func (m *AppsPutRequestBody) GetItemItemBranchesItemProtectionRestrictionsAppsPutRequestBodyMember1()(ItemItemBranchesItemProtectionRestrictionsAppsPutRequestBodyMember1able) {
    return m.itemItemBranchesItemProtectionRestrictionsAppsPutRequestBodyMember1
}
// GetString gets the string property value. Composed type representation for type string
func (m *AppsPutRequestBody) GetString()(*string) {
    return m.string
}
// Serialize serializes information the current object
func (m *AppsPutRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetAppsPutRequestBodyItemItemBranchesItemProtectionRestrictionsAppsPutRequestBodyMember1() != nil {
        err := writer.WriteObjectValue("", m.GetAppsPutRequestBodyItemItemBranchesItemProtectionRestrictionsAppsPutRequestBodyMember1())
        if err != nil {
            return err
        }
    } else if m.GetItemItemBranchesItemProtectionRestrictionsAppsPutRequestBodyMember1() != nil {
        err := writer.WriteObjectValue("", m.GetItemItemBranchesItemProtectionRestrictionsAppsPutRequestBodyMember1())
        if err != nil {
            return err
        }
    } else if m.GetAppsPutRequestBodyString() != nil {
        err := writer.WriteStringValue("", m.GetAppsPutRequestBodyString())
        if err != nil {
            return err
        }
    } else if m.GetString() != nil {
        err := writer.WriteStringValue("", m.GetString())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAppsPutRequestBodyItemItemBranchesItemProtectionRestrictionsAppsPutRequestBodyMember1 sets the ItemItemBranchesItemProtectionRestrictionsAppsPutRequestBodyMember1 property value. Composed type representation for type ItemItemBranchesItemProtectionRestrictionsAppsPutRequestBodyMember1
func (m *AppsPutRequestBody) SetAppsPutRequestBodyItemItemBranchesItemProtectionRestrictionsAppsPutRequestBodyMember1(value ItemItemBranchesItemProtectionRestrictionsAppsPutRequestBodyMember1able)() {
    m.appsPutRequestBodyItemItemBranchesItemProtectionRestrictionsAppsPutRequestBodyMember1 = value
}
// SetAppsPutRequestBodyString sets the string property value. Composed type representation for type string
func (m *AppsPutRequestBody) SetAppsPutRequestBodyString(value *string)() {
    m.appsPutRequestBodyString = value
}
// SetItemItemBranchesItemProtectionRestrictionsAppsPutRequestBodyMember1 sets the ItemItemBranchesItemProtectionRestrictionsAppsPutRequestBodyMember1 property value. Composed type representation for type ItemItemBranchesItemProtectionRestrictionsAppsPutRequestBodyMember1
func (m *AppsPutRequestBody) SetItemItemBranchesItemProtectionRestrictionsAppsPutRequestBodyMember1(value ItemItemBranchesItemProtectionRestrictionsAppsPutRequestBodyMember1able)() {
    m.itemItemBranchesItemProtectionRestrictionsAppsPutRequestBodyMember1 = value
}
// SetString sets the string property value. Composed type representation for type string
func (m *AppsPutRequestBody) SetString(value *string)() {
    m.string = value
}
// ItemItemBranchesItemProtectionRestrictionsAppsRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemItemBranchesItemProtectionRestrictionsAppsRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ItemItemBranchesItemProtectionRestrictionsAppsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemItemBranchesItemProtectionRestrictionsAppsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ItemItemBranchesItemProtectionRestrictionsAppsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemItemBranchesItemProtectionRestrictionsAppsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ItemItemBranchesItemProtectionRestrictionsAppsRequestBuilderPutRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemItemBranchesItemProtectionRestrictionsAppsRequestBuilderPutRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// AppsDeleteRequestBodyable 
type AppsDeleteRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAppsDeleteRequestBodyItemItemBranchesItemProtectionRestrictionsAppsDeleteRequestBodyMember1()(ItemItemBranchesItemProtectionRestrictionsAppsDeleteRequestBodyMember1able)
    GetAppsDeleteRequestBodyString()(*string)
    GetItemItemBranchesItemProtectionRestrictionsAppsDeleteRequestBodyMember1()(ItemItemBranchesItemProtectionRestrictionsAppsDeleteRequestBodyMember1able)
    GetString()(*string)
    SetAppsDeleteRequestBodyItemItemBranchesItemProtectionRestrictionsAppsDeleteRequestBodyMember1(value ItemItemBranchesItemProtectionRestrictionsAppsDeleteRequestBodyMember1able)()
    SetAppsDeleteRequestBodyString(value *string)()
    SetItemItemBranchesItemProtectionRestrictionsAppsDeleteRequestBodyMember1(value ItemItemBranchesItemProtectionRestrictionsAppsDeleteRequestBodyMember1able)()
    SetString(value *string)()
}
// AppsPostRequestBodyable 
type AppsPostRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAppsPostRequestBodyItemItemBranchesItemProtectionRestrictionsAppsPostRequestBodyMember1()(ItemItemBranchesItemProtectionRestrictionsAppsPostRequestBodyMember1able)
    GetAppsPostRequestBodyString()(*string)
    GetItemItemBranchesItemProtectionRestrictionsAppsPostRequestBodyMember1()(ItemItemBranchesItemProtectionRestrictionsAppsPostRequestBodyMember1able)
    GetString()(*string)
    SetAppsPostRequestBodyItemItemBranchesItemProtectionRestrictionsAppsPostRequestBodyMember1(value ItemItemBranchesItemProtectionRestrictionsAppsPostRequestBodyMember1able)()
    SetAppsPostRequestBodyString(value *string)()
    SetItemItemBranchesItemProtectionRestrictionsAppsPostRequestBodyMember1(value ItemItemBranchesItemProtectionRestrictionsAppsPostRequestBodyMember1able)()
    SetString(value *string)()
}
// AppsPutRequestBodyable 
type AppsPutRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAppsPutRequestBodyItemItemBranchesItemProtectionRestrictionsAppsPutRequestBodyMember1()(ItemItemBranchesItemProtectionRestrictionsAppsPutRequestBodyMember1able)
    GetAppsPutRequestBodyString()(*string)
    GetItemItemBranchesItemProtectionRestrictionsAppsPutRequestBodyMember1()(ItemItemBranchesItemProtectionRestrictionsAppsPutRequestBodyMember1able)
    GetString()(*string)
    SetAppsPutRequestBodyItemItemBranchesItemProtectionRestrictionsAppsPutRequestBodyMember1(value ItemItemBranchesItemProtectionRestrictionsAppsPutRequestBodyMember1able)()
    SetAppsPutRequestBodyString(value *string)()
    SetItemItemBranchesItemProtectionRestrictionsAppsPutRequestBodyMember1(value ItemItemBranchesItemProtectionRestrictionsAppsPutRequestBodyMember1able)()
    SetString(value *string)()
}
// NewItemItemBranchesItemProtectionRestrictionsAppsRequestBuilderInternal instantiates a new AppsRequestBuilder and sets the default values.
func NewItemItemBranchesItemProtectionRestrictionsAppsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemBranchesItemProtectionRestrictionsAppsRequestBuilder) {
    m := &ItemItemBranchesItemProtectionRestrictionsAppsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/repos/{owner}/{repo}/branches/{branch}/protection/restrictions/apps", pathParameters),
    }
    return m
}
// NewItemItemBranchesItemProtectionRestrictionsAppsRequestBuilder instantiates a new AppsRequestBuilder and sets the default values.
func NewItemItemBranchesItemProtectionRestrictionsAppsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemBranchesItemProtectionRestrictionsAppsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemItemBranchesItemProtectionRestrictionsAppsRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete protected branches are available in public repositories with GitHub Free and GitHub Free for organizations, and in public and private repositories with GitHub Pro, GitHub Team, GitHub Enterprise Cloud, and GitHub Enterprise Server. For more information, see [GitHub's products](https://docs.github.com/github/getting-started-with-github/githubs-products) in the GitHub Help documentation.Removes the ability of an app to push to this branch. Only installed GitHub Apps with `write` access to the `contents` permission can be added as authorized actors on a protected branch.
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/branches/branch-protection#remove-app-access-restrictions
func (m *ItemItemBranchesItemProtectionRestrictionsAppsRequestBuilder) Delete(ctx context.Context, body AppsDeleteRequestBodyable, requestConfiguration *ItemItemBranchesItemProtectionRestrictionsAppsRequestBuilderDeleteRequestConfiguration)([]i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.Integrationable, error) {
    requestInfo, err := m.ToDeleteRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "422": i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.CreateValidationErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.SendCollection(ctx, requestInfo, i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.CreateIntegrationFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    val := make([]i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.Integrationable, len(res))
    for i, v := range res {
        if v != nil {
            val[i] = v.(i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.Integrationable)
        }
    }
    return val, nil
}
// Get protected branches are available in public repositories with GitHub Free and GitHub Free for organizations, and in public and private repositories with GitHub Pro, GitHub Team, GitHub Enterprise Cloud, and GitHub Enterprise Server. For more information, see [GitHub's products](https://docs.github.com/github/getting-started-with-github/githubs-products) in the GitHub Help documentation.Lists the GitHub Apps that have push access to this branch. Only installed GitHub Apps with `write` access to the `contents` permission can be added as authorized actors on a protected branch.
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/branches/branch-protection#get-apps-with-access-to-the-protected-branch
func (m *ItemItemBranchesItemProtectionRestrictionsAppsRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemItemBranchesItemProtectionRestrictionsAppsRequestBuilderGetRequestConfiguration)([]i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.Integrationable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "404": i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.CreateBasicErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.SendCollection(ctx, requestInfo, i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.CreateIntegrationFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    val := make([]i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.Integrationable, len(res))
    for i, v := range res {
        if v != nil {
            val[i] = v.(i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.Integrationable)
        }
    }
    return val, nil
}
// Post protected branches are available in public repositories with GitHub Free and GitHub Free for organizations, and in public and private repositories with GitHub Pro, GitHub Team, GitHub Enterprise Cloud, and GitHub Enterprise Server. For more information, see [GitHub's products](https://docs.github.com/github/getting-started-with-github/githubs-products) in the GitHub Help documentation.Grants the specified apps push access for this branch. Only installed GitHub Apps with `write` access to the `contents` permission can be added as authorized actors on a protected branch.
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/branches/branch-protection#add-app-access-restrictions
func (m *ItemItemBranchesItemProtectionRestrictionsAppsRequestBuilder) Post(ctx context.Context, body AppsPostRequestBodyable, requestConfiguration *ItemItemBranchesItemProtectionRestrictionsAppsRequestBuilderPostRequestConfiguration)([]i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.Integrationable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "422": i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.CreateValidationErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.SendCollection(ctx, requestInfo, i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.CreateIntegrationFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    val := make([]i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.Integrationable, len(res))
    for i, v := range res {
        if v != nil {
            val[i] = v.(i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.Integrationable)
        }
    }
    return val, nil
}
// Put protected branches are available in public repositories with GitHub Free and GitHub Free for organizations, and in public and private repositories with GitHub Pro, GitHub Team, GitHub Enterprise Cloud, and GitHub Enterprise Server. For more information, see [GitHub's products](https://docs.github.com/github/getting-started-with-github/githubs-products) in the GitHub Help documentation.Replaces the list of apps that have push access to this branch. This removes all apps that previously had push access and grants push access to the new list of apps. Only installed GitHub Apps with `write` access to the `contents` permission can be added as authorized actors on a protected branch.
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/branches/branch-protection#set-app-access-restrictions
func (m *ItemItemBranchesItemProtectionRestrictionsAppsRequestBuilder) Put(ctx context.Context, body AppsPutRequestBodyable, requestConfiguration *ItemItemBranchesItemProtectionRestrictionsAppsRequestBuilderPutRequestConfiguration)([]i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.Integrationable, error) {
    requestInfo, err := m.ToPutRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "422": i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.CreateValidationErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.SendCollection(ctx, requestInfo, i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.CreateIntegrationFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    val := make([]i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.Integrationable, len(res))
    for i, v := range res {
        if v != nil {
            val[i] = v.(i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.Integrationable)
        }
    }
    return val, nil
}
// ToDeleteRequestInformation protected branches are available in public repositories with GitHub Free and GitHub Free for organizations, and in public and private repositories with GitHub Pro, GitHub Team, GitHub Enterprise Cloud, and GitHub Enterprise Server. For more information, see [GitHub's products](https://docs.github.com/github/getting-started-with-github/githubs-products) in the GitHub Help documentation.Removes the ability of an app to push to this branch. Only installed GitHub Apps with `write` access to the `contents` permission can be added as authorized actors on a protected branch.
func (m *ItemItemBranchesItemProtectionRestrictionsAppsRequestBuilder) ToDeleteRequestInformation(ctx context.Context, body AppsDeleteRequestBodyable, requestConfiguration *ItemItemBranchesItemProtectionRestrictionsAppsRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE
    requestInfo.Headers.TryAdd("Accept", "application/json")
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    return requestInfo, nil
}
// ToGetRequestInformation protected branches are available in public repositories with GitHub Free and GitHub Free for organizations, and in public and private repositories with GitHub Pro, GitHub Team, GitHub Enterprise Cloud, and GitHub Enterprise Server. For more information, see [GitHub's products](https://docs.github.com/github/getting-started-with-github/githubs-products) in the GitHub Help documentation.Lists the GitHub Apps that have push access to this branch. Only installed GitHub Apps with `write` access to the `contents` permission can be added as authorized actors on a protected branch.
func (m *ItemItemBranchesItemProtectionRestrictionsAppsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemItemBranchesItemProtectionRestrictionsAppsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation protected branches are available in public repositories with GitHub Free and GitHub Free for organizations, and in public and private repositories with GitHub Pro, GitHub Team, GitHub Enterprise Cloud, and GitHub Enterprise Server. For more information, see [GitHub's products](https://docs.github.com/github/getting-started-with-github/githubs-products) in the GitHub Help documentation.Grants the specified apps push access for this branch. Only installed GitHub Apps with `write` access to the `contents` permission can be added as authorized actors on a protected branch.
func (m *ItemItemBranchesItemProtectionRestrictionsAppsRequestBuilder) ToPostRequestInformation(ctx context.Context, body AppsPostRequestBodyable, requestConfiguration *ItemItemBranchesItemProtectionRestrictionsAppsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
    requestInfo.Headers.TryAdd("Accept", "application/json")
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    return requestInfo, nil
}
// ToPutRequestInformation protected branches are available in public repositories with GitHub Free and GitHub Free for organizations, and in public and private repositories with GitHub Pro, GitHub Team, GitHub Enterprise Cloud, and GitHub Enterprise Server. For more information, see [GitHub's products](https://docs.github.com/github/getting-started-with-github/githubs-products) in the GitHub Help documentation.Replaces the list of apps that have push access to this branch. This removes all apps that previously had push access and grants push access to the new list of apps. Only installed GitHub Apps with `write` access to the `contents` permission can be added as authorized actors on a protected branch.
func (m *ItemItemBranchesItemProtectionRestrictionsAppsRequestBuilder) ToPutRequestInformation(ctx context.Context, body AppsPutRequestBodyable, requestConfiguration *ItemItemBranchesItemProtectionRestrictionsAppsRequestBuilderPutRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PUT
    requestInfo.Headers.TryAdd("Accept", "application/json")
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
func (m *ItemItemBranchesItemProtectionRestrictionsAppsRequestBuilder) WithUrl(rawUrl string)(*ItemItemBranchesItemProtectionRestrictionsAppsRequestBuilder) {
    return NewItemItemBranchesItemProtectionRestrictionsAppsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
