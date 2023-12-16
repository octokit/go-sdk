package projects

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemCollaboratorsItemWithUsernamePutRequestBody 
type ItemCollaboratorsItemWithUsernamePutRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
}
// NewItemCollaboratorsItemWithUsernamePutRequestBody instantiates a new ItemCollaboratorsItemWithUsernamePutRequestBody and sets the default values.
func NewItemCollaboratorsItemWithUsernamePutRequestBody()(*ItemCollaboratorsItemWithUsernamePutRequestBody) {
    m := &ItemCollaboratorsItemWithUsernamePutRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateItemCollaboratorsItemWithUsernamePutRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemCollaboratorsItemWithUsernamePutRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemCollaboratorsItemWithUsernamePutRequestBody(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ItemCollaboratorsItemWithUsernamePutRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ItemCollaboratorsItemWithUsernamePutRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    return res
}
// Serialize serializes information the current object
func (m *ItemCollaboratorsItemWithUsernamePutRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ItemCollaboratorsItemWithUsernamePutRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// ItemCollaboratorsItemWithUsernamePutRequestBodyable 
type ItemCollaboratorsItemWithUsernamePutRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}