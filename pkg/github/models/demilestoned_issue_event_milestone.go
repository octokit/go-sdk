package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DemilestonedIssueEvent_milestone 
type DemilestonedIssueEvent_milestone struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The title property
    title *string
}
// NewDemilestonedIssueEvent_milestone instantiates a new demilestonedIssueEvent_milestone and sets the default values.
func NewDemilestonedIssueEvent_milestone()(*DemilestonedIssueEvent_milestone) {
    m := &DemilestonedIssueEvent_milestone{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateDemilestonedIssueEvent_milestoneFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDemilestonedIssueEvent_milestoneFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDemilestonedIssueEvent_milestone(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *DemilestonedIssueEvent_milestone) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DemilestonedIssueEvent_milestone) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["title"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTitle(val)
        }
        return nil
    }
    return res
}
// GetTitle gets the title property value. The title property
func (m *DemilestonedIssueEvent_milestone) GetTitle()(*string) {
    return m.title
}
// Serialize serializes information the current object
func (m *DemilestonedIssueEvent_milestone) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("title", m.GetTitle())
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
func (m *DemilestonedIssueEvent_milestone) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetTitle sets the title property value. The title property
func (m *DemilestonedIssueEvent_milestone) SetTitle(value *string)() {
    m.title = value
}
// DemilestonedIssueEvent_milestoneable 
type DemilestonedIssueEvent_milestoneable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetTitle()(*string)
    SetTitle(value *string)()
}
