package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type SubIssuesSummary struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The completed property
    completed *int32
    // The percent_completed property
    percent_completed *int32
    // The total property
    total *int32
}
// NewSubIssuesSummary instantiates a new SubIssuesSummary and sets the default values.
func NewSubIssuesSummary()(*SubIssuesSummary) {
    m := &SubIssuesSummary{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateSubIssuesSummaryFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateSubIssuesSummaryFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewSubIssuesSummary(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *SubIssuesSummary) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetCompleted gets the completed property value. The completed property
// returns a *int32 when successful
func (m *SubIssuesSummary) GetCompleted()(*int32) {
    return m.completed
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *SubIssuesSummary) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["completed"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCompleted(val)
        }
        return nil
    }
    res["percent_completed"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPercentCompleted(val)
        }
        return nil
    }
    res["total"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTotal(val)
        }
        return nil
    }
    return res
}
// GetPercentCompleted gets the percent_completed property value. The percent_completed property
// returns a *int32 when successful
func (m *SubIssuesSummary) GetPercentCompleted()(*int32) {
    return m.percent_completed
}
// GetTotal gets the total property value. The total property
// returns a *int32 when successful
func (m *SubIssuesSummary) GetTotal()(*int32) {
    return m.total
}
// Serialize serializes information the current object
func (m *SubIssuesSummary) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteInt32Value("completed", m.GetCompleted())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("percent_completed", m.GetPercentCompleted())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("total", m.GetTotal())
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
func (m *SubIssuesSummary) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetCompleted sets the completed property value. The completed property
func (m *SubIssuesSummary) SetCompleted(value *int32)() {
    m.completed = value
}
// SetPercentCompleted sets the percent_completed property value. The percent_completed property
func (m *SubIssuesSummary) SetPercentCompleted(value *int32)() {
    m.percent_completed = value
}
// SetTotal sets the total property value. The total property
func (m *SubIssuesSummary) SetTotal(value *int32)() {
    m.total = value
}
type SubIssuesSummaryable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCompleted()(*int32)
    GetPercentCompleted()(*int32)
    GetTotal()(*int32)
    SetCompleted(value *int32)()
    SetPercentCompleted(value *int32)()
    SetTotal(value *int32)()
}
