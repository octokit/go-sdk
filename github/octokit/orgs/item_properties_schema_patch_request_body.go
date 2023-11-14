package orgs

import (
    i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e "github.com/octokit/go-sdk/github/octokit/models"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemPropertiesSchemaPatchRequestBody 
type ItemPropertiesSchemaPatchRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The array of custom properties to create or update.
    properties []i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.OrgCustomPropertyable
}
// NewItemPropertiesSchemaPatchRequestBody instantiates a new ItemPropertiesSchemaPatchRequestBody and sets the default values.
func NewItemPropertiesSchemaPatchRequestBody()(*ItemPropertiesSchemaPatchRequestBody) {
    m := &ItemPropertiesSchemaPatchRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateItemPropertiesSchemaPatchRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemPropertiesSchemaPatchRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemPropertiesSchemaPatchRequestBody(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ItemPropertiesSchemaPatchRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ItemPropertiesSchemaPatchRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["properties"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.CreateOrgCustomPropertyFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.OrgCustomPropertyable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.OrgCustomPropertyable)
                }
            }
            m.SetProperties(res)
        }
        return nil
    }
    return res
}
// GetProperties gets the properties property value. The array of custom properties to create or update.
func (m *ItemPropertiesSchemaPatchRequestBody) GetProperties()([]i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.OrgCustomPropertyable) {
    return m.properties
}
// Serialize serializes information the current object
func (m *ItemPropertiesSchemaPatchRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetProperties() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetProperties()))
        for i, v := range m.GetProperties() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("properties", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ItemPropertiesSchemaPatchRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetProperties sets the properties property value. The array of custom properties to create or update.
func (m *ItemPropertiesSchemaPatchRequestBody) SetProperties(value []i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.OrgCustomPropertyable)() {
    m.properties = value
}
// ItemPropertiesSchemaPatchRequestBodyable 
type ItemPropertiesSchemaPatchRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetProperties()([]i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.OrgCustomPropertyable)
    SetProperties(value []i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.OrgCustomPropertyable)()
}
