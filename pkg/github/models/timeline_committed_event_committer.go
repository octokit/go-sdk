package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// TimelineCommittedEvent_committer identifying information for the git-user
type TimelineCommittedEvent_committer struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Timestamp of the commit
    date *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Git email address of the user
    email *string
    // Name of the git user
    name *string
}
// NewTimelineCommittedEvent_committer instantiates a new timelineCommittedEvent_committer and sets the default values.
func NewTimelineCommittedEvent_committer()(*TimelineCommittedEvent_committer) {
    m := &TimelineCommittedEvent_committer{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateTimelineCommittedEvent_committerFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateTimelineCommittedEvent_committerFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewTimelineCommittedEvent_committer(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *TimelineCommittedEvent_committer) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetDate gets the date property value. Timestamp of the commit
func (m *TimelineCommittedEvent_committer) GetDate()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.date
}
// GetEmail gets the email property value. Git email address of the user
func (m *TimelineCommittedEvent_committer) GetEmail()(*string) {
    return m.email
}
// GetFieldDeserializers the deserialization information for the current model
func (m *TimelineCommittedEvent_committer) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["date"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDate(val)
        }
        return nil
    }
    res["email"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEmail(val)
        }
        return nil
    }
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
// GetName gets the name property value. Name of the git user
func (m *TimelineCommittedEvent_committer) GetName()(*string) {
    return m.name
}
// Serialize serializes information the current object
func (m *TimelineCommittedEvent_committer) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteTimeValue("date", m.GetDate())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("email", m.GetEmail())
        if err != nil {
            return err
        }
    }
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
func (m *TimelineCommittedEvent_committer) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetDate sets the date property value. Timestamp of the commit
func (m *TimelineCommittedEvent_committer) SetDate(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.date = value
}
// SetEmail sets the email property value. Git email address of the user
func (m *TimelineCommittedEvent_committer) SetEmail(value *string)() {
    m.email = value
}
// SetName sets the name property value. Name of the git user
func (m *TimelineCommittedEvent_committer) SetName(value *string)() {
    m.name = value
}
// TimelineCommittedEvent_committerable 
type TimelineCommittedEvent_committerable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetDate()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetEmail()(*string)
    GetName()(*string)
    SetDate(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetEmail(value *string)()
    SetName(value *string)()
}
