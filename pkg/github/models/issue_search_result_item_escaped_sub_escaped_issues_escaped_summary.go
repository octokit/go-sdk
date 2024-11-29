package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type IssueSearchResultItem_sub_issues_summary struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The completed property
    completed *int32
    // The percent_completed property
    percent_completed *int32
    // The total property
    total *int32
}
// NewIssueSearchResultItem_sub_issues_summary instantiates a new IssueSearchResultItem_sub_issues_summary and sets the default values.
func NewIssueSearchResultItem_sub_issues_summary()(*IssueSearchResultItem_sub_issues_summary) {
    m := &IssueSearchResultItem_sub_issues_summary{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateIssueSearchResultItem_sub_issues_summaryFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateIssueSearchResultItem_sub_issues_summaryFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewIssueSearchResultItem_sub_issues_summary(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *IssueSearchResultItem_sub_issues_summary) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetCompleted gets the completed property value. The completed property
// returns a *int32 when successful
func (m *IssueSearchResultItem_sub_issues_summary) GetCompleted()(*int32) {
    return m.completed
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *IssueSearchResultItem_sub_issues_summary) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
func (m *IssueSearchResultItem_sub_issues_summary) GetPercentCompleted()(*int32) {
    return m.percent_completed
}
// GetTotal gets the total property value. The total property
// returns a *int32 when successful
func (m *IssueSearchResultItem_sub_issues_summary) GetTotal()(*int32) {
    return m.total
}
// Serialize serializes information the current object
func (m *IssueSearchResultItem_sub_issues_summary) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
func (m *IssueSearchResultItem_sub_issues_summary) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetCompleted sets the completed property value. The completed property
func (m *IssueSearchResultItem_sub_issues_summary) SetCompleted(value *int32)() {
    m.completed = value
}
// SetPercentCompleted sets the percent_completed property value. The percent_completed property
func (m *IssueSearchResultItem_sub_issues_summary) SetPercentCompleted(value *int32)() {
    m.percent_completed = value
}
// SetTotal sets the total property value. The total property
func (m *IssueSearchResultItem_sub_issues_summary) SetTotal(value *int32)() {
    m.total = value
}
type IssueSearchResultItem_sub_issues_summaryable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCompleted()(*int32)
    GetPercentCompleted()(*int32)
    GetTotal()(*int32)
    SetCompleted(value *int32)()
    SetPercentCompleted(value *int32)()
    SetTotal(value *int32)()
}
