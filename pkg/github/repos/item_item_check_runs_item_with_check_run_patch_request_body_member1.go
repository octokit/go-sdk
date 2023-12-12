package repos

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemItemCheckRunsItemWithCheck_run_PatchRequestBodyMember1 
type ItemItemCheckRunsItemWithCheck_run_PatchRequestBodyMember1 struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
}
// NewItemItemCheckRunsItemWithCheck_run_PatchRequestBodyMember1 instantiates a new ItemItemCheckRunsItemWithCheck_run_PatchRequestBodyMember1 and sets the default values.
func NewItemItemCheckRunsItemWithCheck_run_PatchRequestBodyMember1()(*ItemItemCheckRunsItemWithCheck_run_PatchRequestBodyMember1) {
    m := &ItemItemCheckRunsItemWithCheck_run_PatchRequestBodyMember1{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateItemItemCheckRunsItemWithCheck_run_PatchRequestBodyMember1FromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemItemCheckRunsItemWithCheck_run_PatchRequestBodyMember1FromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemItemCheckRunsItemWithCheck_run_PatchRequestBodyMember1(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ItemItemCheckRunsItemWithCheck_run_PatchRequestBodyMember1) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ItemItemCheckRunsItemWithCheck_run_PatchRequestBodyMember1) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    return res
}
// Serialize serializes information the current object
func (m *ItemItemCheckRunsItemWithCheck_run_PatchRequestBodyMember1) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ItemItemCheckRunsItemWithCheck_run_PatchRequestBodyMember1) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// ItemItemCheckRunsItemWithCheck_run_PatchRequestBodyMember1able 
type ItemItemCheckRunsItemWithCheck_run_PatchRequestBodyMember1able interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
