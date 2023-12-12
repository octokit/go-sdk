package repos

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemItemIssuesItemLabelsPostRequestBodyMember2_labels 
type ItemItemIssuesItemLabelsPostRequestBodyMember2_labels struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The name property
    name *string
}
// NewItemItemIssuesItemLabelsPostRequestBodyMember2_labels instantiates a new ItemItemIssuesItemLabelsPostRequestBodyMember2_labels and sets the default values.
func NewItemItemIssuesItemLabelsPostRequestBodyMember2_labels()(*ItemItemIssuesItemLabelsPostRequestBodyMember2_labels) {
    m := &ItemItemIssuesItemLabelsPostRequestBodyMember2_labels{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateItemItemIssuesItemLabelsPostRequestBodyMember2_labelsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemItemIssuesItemLabelsPostRequestBodyMember2_labelsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemItemIssuesItemLabelsPostRequestBodyMember2_labels(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ItemItemIssuesItemLabelsPostRequestBodyMember2_labels) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ItemItemIssuesItemLabelsPostRequestBodyMember2_labels) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["name"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetName(val)
        }
        return nil
    }
    return res
}
// GetName gets the name property value. The name property
func (m *ItemItemIssuesItemLabelsPostRequestBodyMember2_labels) GetName()(*string) {
    return m.name
}
// Serialize serializes information the current object
func (m *ItemItemIssuesItemLabelsPostRequestBodyMember2_labels) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("name", m.GetName())
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
func (m *ItemItemIssuesItemLabelsPostRequestBodyMember2_labels) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetName sets the name property value. The name property
func (m *ItemItemIssuesItemLabelsPostRequestBodyMember2_labels) SetName(value *string)() {
    m.name = value
}
// ItemItemIssuesItemLabelsPostRequestBodyMember2_labelsable 
type ItemItemIssuesItemLabelsPostRequestBodyMember2_labelsable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetName()(*string)
    SetName(value *string)()
}
