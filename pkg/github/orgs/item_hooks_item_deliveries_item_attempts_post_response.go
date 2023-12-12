package orgs

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemHooksItemDeliveriesItemAttemptsPostResponse 
type ItemHooksItemDeliveriesItemAttemptsPostResponse struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
}
// NewItemHooksItemDeliveriesItemAttemptsPostResponse instantiates a new ItemHooksItemDeliveriesItemAttemptsPostResponse and sets the default values.
func NewItemHooksItemDeliveriesItemAttemptsPostResponse()(*ItemHooksItemDeliveriesItemAttemptsPostResponse) {
    m := &ItemHooksItemDeliveriesItemAttemptsPostResponse{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateItemHooksItemDeliveriesItemAttemptsPostResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemHooksItemDeliveriesItemAttemptsPostResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemHooksItemDeliveriesItemAttemptsPostResponse(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ItemHooksItemDeliveriesItemAttemptsPostResponse) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ItemHooksItemDeliveriesItemAttemptsPostResponse) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    return res
}
// Serialize serializes information the current object
func (m *ItemHooksItemDeliveriesItemAttemptsPostResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ItemHooksItemDeliveriesItemAttemptsPostResponse) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// ItemHooksItemDeliveriesItemAttemptsPostResponseable 
type ItemHooksItemDeliveriesItemAttemptsPostResponseable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
