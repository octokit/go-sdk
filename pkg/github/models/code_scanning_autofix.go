package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type CodeScanningAutofix struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The description of an autofix.
    description *string
    // The start time of an autofix in ISO 8601 format: `YYYY-MM-DDTHH:MM:SSZ`.
    started_at *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The status of an autofix.
    status *CodeScanningAutofixStatus
}
// NewCodeScanningAutofix instantiates a new CodeScanningAutofix and sets the default values.
func NewCodeScanningAutofix()(*CodeScanningAutofix) {
    m := &CodeScanningAutofix{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateCodeScanningAutofixFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateCodeScanningAutofixFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCodeScanningAutofix(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *CodeScanningAutofix) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetDescription gets the description property value. The description of an autofix.
// returns a *string when successful
func (m *CodeScanningAutofix) GetDescription()(*string) {
    return m.description
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *CodeScanningAutofix) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["description"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDescription(val)
        }
        return nil
    }
    res["started_at"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStartedAt(val)
        }
        return nil
    }
    res["status"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseCodeScanningAutofixStatus)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStatus(val.(*CodeScanningAutofixStatus))
        }
        return nil
    }
    return res
}
// GetStartedAt gets the started_at property value. The start time of an autofix in ISO 8601 format: `YYYY-MM-DDTHH:MM:SSZ`.
// returns a *Time when successful
func (m *CodeScanningAutofix) GetStartedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.started_at
}
// GetStatus gets the status property value. The status of an autofix.
// returns a *CodeScanningAutofixStatus when successful
func (m *CodeScanningAutofix) GetStatus()(*CodeScanningAutofixStatus) {
    return m.status
}
// Serialize serializes information the current object
func (m *CodeScanningAutofix) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("description", m.GetDescription())
        if err != nil {
            return err
        }
    }
    if m.GetStatus() != nil {
        cast := (*m.GetStatus()).String()
        err := writer.WriteStringValue("status", &cast)
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
func (m *CodeScanningAutofix) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetDescription sets the description property value. The description of an autofix.
func (m *CodeScanningAutofix) SetDescription(value *string)() {
    m.description = value
}
// SetStartedAt sets the started_at property value. The start time of an autofix in ISO 8601 format: `YYYY-MM-DDTHH:MM:SSZ`.
func (m *CodeScanningAutofix) SetStartedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.started_at = value
}
// SetStatus sets the status property value. The status of an autofix.
func (m *CodeScanningAutofix) SetStatus(value *CodeScanningAutofixStatus)() {
    m.status = value
}
type CodeScanningAutofixable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetDescription()(*string)
    GetStartedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetStatus()(*CodeScanningAutofixStatus)
    SetDescription(value *string)()
    SetStartedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetStatus(value *CodeScanningAutofixStatus)()
}
