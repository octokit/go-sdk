package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// CopilotSeatDetails information about a Copilot Business seat assignment for a user, team, or organization.
type CopilotSeatDetails struct {
    // The assignee that has been granted access to GitHub Copilot.
    assignee CopilotSeatDetails_CopilotSeatDetails_assigneeable
    // The team that granted access to GitHub Copilot to the assignee. This will be null if the user was assigned a seat individually.
    assigning_team Teamable
    // Timestamp of when the assignee was last granted access to GitHub Copilot, in ISO 8601 format.
    created_at *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Timestamp of user's last GitHub Copilot activity, in ISO 8601 format.
    last_activity_at *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Last editor that was used by the user for a GitHub Copilot completion.
    last_activity_editor *string
    // The pending cancellation date for the seat, in `YYYY-MM-DD` format. This will be null unless the assignee's Copilot access has been canceled during the current billing cycle. If the seat has been cancelled, this corresponds to the start of the organization's next billing cycle.
    pending_cancellation_date *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly
    // Timestamp of when the assignee's GitHub Copilot access was last updated, in ISO 8601 format.
    updated_at *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
}
// CopilotSeatDetails_CopilotSeatDetails_assignee composed type wrapper for classes organization, simpleUser, team
type CopilotSeatDetails_CopilotSeatDetails_assignee struct {
    // Composed type representation for type organization
    organization Organizationable
    // Composed type representation for type simpleUser
    simpleUser SimpleUserable
    // Composed type representation for type team
    team Teamable
}
// NewCopilotSeatDetails_CopilotSeatDetails_assignee instantiates a new copilotSeatDetails_assignee and sets the default values.
func NewCopilotSeatDetails_CopilotSeatDetails_assignee()(*CopilotSeatDetails_CopilotSeatDetails_assignee) {
    m := &CopilotSeatDetails_CopilotSeatDetails_assignee{
    }
    return m
}
// CreateCopilotSeatDetails_CopilotSeatDetails_assigneeFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateCopilotSeatDetails_CopilotSeatDetails_assigneeFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    result := NewCopilotSeatDetails_CopilotSeatDetails_assignee()
    if parseNode != nil {
        mappingValueNode, err := parseNode.GetChildNode("")
        if err != nil {
            return nil, err
        }
        if mappingValueNode != nil {
            mappingValue, err := mappingValueNode.GetStringValue()
            if err != nil {
                return nil, err
            }
            if mappingValue != nil {
            }
        }
    }
    return result, nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *CopilotSeatDetails_CopilotSeatDetails_assignee) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    return make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
}
// GetIsComposedType determines if the current object is a wrapper around a composed type
func (m *CopilotSeatDetails_CopilotSeatDetails_assignee) GetIsComposedType()(bool) {
    return true
}
// GetOrganization gets the organization property value. Composed type representation for type organization
func (m *CopilotSeatDetails_CopilotSeatDetails_assignee) GetOrganization()(Organizationable) {
    return m.organization
}
// GetSimpleUser gets the simpleUser property value. Composed type representation for type simpleUser
func (m *CopilotSeatDetails_CopilotSeatDetails_assignee) GetSimpleUser()(SimpleUserable) {
    return m.simpleUser
}
// GetTeam gets the team property value. Composed type representation for type team
func (m *CopilotSeatDetails_CopilotSeatDetails_assignee) GetTeam()(Teamable) {
    return m.team
}
// Serialize serializes information the current object
func (m *CopilotSeatDetails_CopilotSeatDetails_assignee) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetOrganization() != nil {
        err := writer.WriteObjectValue("", m.GetOrganization())
        if err != nil {
            return err
        }
    } else if m.GetSimpleUser() != nil {
        err := writer.WriteObjectValue("", m.GetSimpleUser())
        if err != nil {
            return err
        }
    } else if m.GetTeam() != nil {
        err := writer.WriteObjectValue("", m.GetTeam())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetOrganization sets the organization property value. Composed type representation for type organization
func (m *CopilotSeatDetails_CopilotSeatDetails_assignee) SetOrganization(value Organizationable)() {
    m.organization = value
}
// SetSimpleUser sets the simpleUser property value. Composed type representation for type simpleUser
func (m *CopilotSeatDetails_CopilotSeatDetails_assignee) SetSimpleUser(value SimpleUserable)() {
    m.simpleUser = value
}
// SetTeam sets the team property value. Composed type representation for type team
func (m *CopilotSeatDetails_CopilotSeatDetails_assignee) SetTeam(value Teamable)() {
    m.team = value
}
// CopilotSeatDetails_CopilotSeatDetails_assigneeable 
type CopilotSeatDetails_CopilotSeatDetails_assigneeable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetOrganization()(Organizationable)
    GetSimpleUser()(SimpleUserable)
    GetTeam()(Teamable)
    SetOrganization(value Organizationable)()
    SetSimpleUser(value SimpleUserable)()
    SetTeam(value Teamable)()
}
// NewCopilotSeatDetails instantiates a new copilotSeatDetails and sets the default values.
func NewCopilotSeatDetails()(*CopilotSeatDetails) {
    m := &CopilotSeatDetails{
    }
    return m
}
// CreateCopilotSeatDetailsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateCopilotSeatDetailsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCopilotSeatDetails(), nil
}
// GetAssignee gets the assignee property value. The assignee that has been granted access to GitHub Copilot.
func (m *CopilotSeatDetails) GetAssignee()(CopilotSeatDetails_CopilotSeatDetails_assigneeable) {
    return m.assignee
}
// GetAssigningTeam gets the assigning_team property value. The team that granted access to GitHub Copilot to the assignee. This will be null if the user was assigned a seat individually.
func (m *CopilotSeatDetails) GetAssigningTeam()(Teamable) {
    return m.assigning_team
}
// GetCreatedAt gets the created_at property value. Timestamp of when the assignee was last granted access to GitHub Copilot, in ISO 8601 format.
func (m *CopilotSeatDetails) GetCreatedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.created_at
}
// GetFieldDeserializers the deserialization information for the current model
func (m *CopilotSeatDetails) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["assignee"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateCopilotSeatDetails_CopilotSeatDetails_assigneeFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAssignee(val.(CopilotSeatDetails_CopilotSeatDetails_assigneeable))
        }
        return nil
    }
    res["assigning_team"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateTeamFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAssigningTeam(val.(Teamable))
        }
        return nil
    }
    res["created_at"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCreatedAt(val)
        }
        return nil
    }
    res["last_activity_at"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastActivityAt(val)
        }
        return nil
    }
    res["last_activity_editor"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastActivityEditor(val)
        }
        return nil
    }
    res["pending_cancellation_date"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetDateOnlyValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPendingCancellationDate(val)
        }
        return nil
    }
    res["updated_at"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUpdatedAt(val)
        }
        return nil
    }
    return res
}
// GetLastActivityAt gets the last_activity_at property value. Timestamp of user's last GitHub Copilot activity, in ISO 8601 format.
func (m *CopilotSeatDetails) GetLastActivityAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.last_activity_at
}
// GetLastActivityEditor gets the last_activity_editor property value. Last editor that was used by the user for a GitHub Copilot completion.
func (m *CopilotSeatDetails) GetLastActivityEditor()(*string) {
    return m.last_activity_editor
}
// GetPendingCancellationDate gets the pending_cancellation_date property value. The pending cancellation date for the seat, in `YYYY-MM-DD` format. This will be null unless the assignee's Copilot access has been canceled during the current billing cycle. If the seat has been cancelled, this corresponds to the start of the organization's next billing cycle.
func (m *CopilotSeatDetails) GetPendingCancellationDate()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly) {
    return m.pending_cancellation_date
}
// GetUpdatedAt gets the updated_at property value. Timestamp of when the assignee's GitHub Copilot access was last updated, in ISO 8601 format.
func (m *CopilotSeatDetails) GetUpdatedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.updated_at
}
// Serialize serializes information the current object
func (m *CopilotSeatDetails) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("assignee", m.GetAssignee())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("assigning_team", m.GetAssigningTeam())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteTimeValue("created_at", m.GetCreatedAt())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteTimeValue("last_activity_at", m.GetLastActivityAt())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("last_activity_editor", m.GetLastActivityEditor())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteDateOnlyValue("pending_cancellation_date", m.GetPendingCancellationDate())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteTimeValue("updated_at", m.GetUpdatedAt())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAssignee sets the assignee property value. The assignee that has been granted access to GitHub Copilot.
func (m *CopilotSeatDetails) SetAssignee(value CopilotSeatDetails_CopilotSeatDetails_assigneeable)() {
    m.assignee = value
}
// SetAssigningTeam sets the assigning_team property value. The team that granted access to GitHub Copilot to the assignee. This will be null if the user was assigned a seat individually.
func (m *CopilotSeatDetails) SetAssigningTeam(value Teamable)() {
    m.assigning_team = value
}
// SetCreatedAt sets the created_at property value. Timestamp of when the assignee was last granted access to GitHub Copilot, in ISO 8601 format.
func (m *CopilotSeatDetails) SetCreatedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.created_at = value
}
// SetLastActivityAt sets the last_activity_at property value. Timestamp of user's last GitHub Copilot activity, in ISO 8601 format.
func (m *CopilotSeatDetails) SetLastActivityAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.last_activity_at = value
}
// SetLastActivityEditor sets the last_activity_editor property value. Last editor that was used by the user for a GitHub Copilot completion.
func (m *CopilotSeatDetails) SetLastActivityEditor(value *string)() {
    m.last_activity_editor = value
}
// SetPendingCancellationDate sets the pending_cancellation_date property value. The pending cancellation date for the seat, in `YYYY-MM-DD` format. This will be null unless the assignee's Copilot access has been canceled during the current billing cycle. If the seat has been cancelled, this corresponds to the start of the organization's next billing cycle.
func (m *CopilotSeatDetails) SetPendingCancellationDate(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)() {
    m.pending_cancellation_date = value
}
// SetUpdatedAt sets the updated_at property value. Timestamp of when the assignee's GitHub Copilot access was last updated, in ISO 8601 format.
func (m *CopilotSeatDetails) SetUpdatedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.updated_at = value
}
// CopilotSeatDetailsable 
type CopilotSeatDetailsable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAssignee()(CopilotSeatDetails_CopilotSeatDetails_assigneeable)
    GetAssigningTeam()(Teamable)
    GetCreatedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetLastActivityAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetLastActivityEditor()(*string)
    GetPendingCancellationDate()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)
    GetUpdatedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    SetAssignee(value CopilotSeatDetails_CopilotSeatDetails_assigneeable)()
    SetAssigningTeam(value Teamable)()
    SetCreatedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetLastActivityAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetLastActivityEditor(value *string)()
    SetPendingCancellationDate(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)()
    SetUpdatedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
}
