package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// CopilotSeatDetails information about a Copilot Business seat assignment for a user, team, or organization.
type CopilotSeatDetails struct {
    // A GitHub user.
    assignee SimpleUserable
    // The team through which the assignee is granted access to GitHub Copilot, if applicable.
    assigning_team CopilotSeatDetails_CopilotSeatDetails_assigning_teamable
    // Timestamp of when the assignee was last granted access to GitHub Copilot, in ISO 8601 format.
    created_at *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Timestamp of user's last GitHub Copilot activity, in ISO 8601 format.
    last_activity_at *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Last editor that was used by the user for a GitHub Copilot completion.
    last_activity_editor *string
    // A GitHub organization.
    organization NullableOrganizationSimpleable
    // The pending cancellation date for the seat, in `YYYY-MM-DD` format. This will be null unless the assignee's Copilot access has been canceled during the current billing cycle. If the seat has been cancelled, this corresponds to the start of the organization's next billing cycle.
    pending_cancellation_date *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly
    // The Copilot plan of the organization, or the parent enterprise, when applicable.
    plan_type *CopilotSeatDetails_plan_type
    // **Closing down notice:** This field is no longer relevant and is closing down. Use the `created_at` field to determine when the assignee was last granted access to GitHub Copilot. Timestamp of when the assignee's GitHub Copilot access was last updated, in ISO 8601 format.
    // Deprecated: 
    updated_at *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
}
// CopilotSeatDetails_CopilotSeatDetails_assigning_team composed type wrapper for classes EnterpriseTeamable, Teamable
type CopilotSeatDetails_CopilotSeatDetails_assigning_team struct {
    // Composed type representation for type EnterpriseTeamable
    enterpriseTeam EnterpriseTeamable
    // Composed type representation for type Teamable
    team Teamable
}
// NewCopilotSeatDetails_CopilotSeatDetails_assigning_team instantiates a new CopilotSeatDetails_CopilotSeatDetails_assigning_team and sets the default values.
func NewCopilotSeatDetails_CopilotSeatDetails_assigning_team()(*CopilotSeatDetails_CopilotSeatDetails_assigning_team) {
    m := &CopilotSeatDetails_CopilotSeatDetails_assigning_team{
    }
    return m
}
// CreateCopilotSeatDetails_CopilotSeatDetails_assigning_teamFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateCopilotSeatDetails_CopilotSeatDetails_assigning_teamFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    result := NewCopilotSeatDetails_CopilotSeatDetails_assigning_team()
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
// GetEnterpriseTeam gets the enterpriseTeam property value. Composed type representation for type EnterpriseTeamable
// returns a EnterpriseTeamable when successful
func (m *CopilotSeatDetails_CopilotSeatDetails_assigning_team) GetEnterpriseTeam()(EnterpriseTeamable) {
    return m.enterpriseTeam
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *CopilotSeatDetails_CopilotSeatDetails_assigning_team) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    return make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
}
// GetIsComposedType determines if the current object is a wrapper around a composed type
// returns a bool when successful
func (m *CopilotSeatDetails_CopilotSeatDetails_assigning_team) GetIsComposedType()(bool) {
    return true
}
// GetTeam gets the team property value. Composed type representation for type Teamable
// returns a Teamable when successful
func (m *CopilotSeatDetails_CopilotSeatDetails_assigning_team) GetTeam()(Teamable) {
    return m.team
}
// Serialize serializes information the current object
func (m *CopilotSeatDetails_CopilotSeatDetails_assigning_team) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetEnterpriseTeam() != nil {
        err := writer.WriteObjectValue("", m.GetEnterpriseTeam())
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
// SetEnterpriseTeam sets the enterpriseTeam property value. Composed type representation for type EnterpriseTeamable
func (m *CopilotSeatDetails_CopilotSeatDetails_assigning_team) SetEnterpriseTeam(value EnterpriseTeamable)() {
    m.enterpriseTeam = value
}
// SetTeam sets the team property value. Composed type representation for type Teamable
func (m *CopilotSeatDetails_CopilotSeatDetails_assigning_team) SetTeam(value Teamable)() {
    m.team = value
}
type CopilotSeatDetails_CopilotSeatDetails_assigning_teamable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetEnterpriseTeam()(EnterpriseTeamable)
    GetTeam()(Teamable)
    SetEnterpriseTeam(value EnterpriseTeamable)()
    SetTeam(value Teamable)()
}
// NewCopilotSeatDetails instantiates a new CopilotSeatDetails and sets the default values.
func NewCopilotSeatDetails()(*CopilotSeatDetails) {
    m := &CopilotSeatDetails{
    }
    return m
}
// CreateCopilotSeatDetailsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateCopilotSeatDetailsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCopilotSeatDetails(), nil
}
// GetAssignee gets the assignee property value. A GitHub user.
// returns a SimpleUserable when successful
func (m *CopilotSeatDetails) GetAssignee()(SimpleUserable) {
    return m.assignee
}
// GetAssigningTeam gets the assigning_team property value. The team through which the assignee is granted access to GitHub Copilot, if applicable.
// returns a CopilotSeatDetails_CopilotSeatDetails_assigning_teamable when successful
func (m *CopilotSeatDetails) GetAssigningTeam()(CopilotSeatDetails_CopilotSeatDetails_assigning_teamable) {
    return m.assigning_team
}
// GetCreatedAt gets the created_at property value. Timestamp of when the assignee was last granted access to GitHub Copilot, in ISO 8601 format.
// returns a *Time when successful
func (m *CopilotSeatDetails) GetCreatedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.created_at
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *CopilotSeatDetails) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["assignee"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateSimpleUserFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAssignee(val.(SimpleUserable))
        }
        return nil
    }
    res["assigning_team"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateCopilotSeatDetails_CopilotSeatDetails_assigning_teamFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAssigningTeam(val.(CopilotSeatDetails_CopilotSeatDetails_assigning_teamable))
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
    res["organization"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateNullableOrganizationSimpleFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOrganization(val.(NullableOrganizationSimpleable))
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
    res["plan_type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseCopilotSeatDetails_plan_type)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPlanType(val.(*CopilotSeatDetails_plan_type))
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
// returns a *Time when successful
func (m *CopilotSeatDetails) GetLastActivityAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.last_activity_at
}
// GetLastActivityEditor gets the last_activity_editor property value. Last editor that was used by the user for a GitHub Copilot completion.
// returns a *string when successful
func (m *CopilotSeatDetails) GetLastActivityEditor()(*string) {
    return m.last_activity_editor
}
// GetOrganization gets the organization property value. A GitHub organization.
// returns a NullableOrganizationSimpleable when successful
func (m *CopilotSeatDetails) GetOrganization()(NullableOrganizationSimpleable) {
    return m.organization
}
// GetPendingCancellationDate gets the pending_cancellation_date property value. The pending cancellation date for the seat, in `YYYY-MM-DD` format. This will be null unless the assignee's Copilot access has been canceled during the current billing cycle. If the seat has been cancelled, this corresponds to the start of the organization's next billing cycle.
// returns a *DateOnly when successful
func (m *CopilotSeatDetails) GetPendingCancellationDate()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly) {
    return m.pending_cancellation_date
}
// GetPlanType gets the plan_type property value. The Copilot plan of the organization, or the parent enterprise, when applicable.
// returns a *CopilotSeatDetails_plan_type when successful
func (m *CopilotSeatDetails) GetPlanType()(*CopilotSeatDetails_plan_type) {
    return m.plan_type
}
// GetUpdatedAt gets the updated_at property value. **Closing down notice:** This field is no longer relevant and is closing down. Use the `created_at` field to determine when the assignee was last granted access to GitHub Copilot. Timestamp of when the assignee's GitHub Copilot access was last updated, in ISO 8601 format.
// Deprecated: 
// returns a *Time when successful
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
        err := writer.WriteObjectValue("organization", m.GetOrganization())
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
    if m.GetPlanType() != nil {
        cast := (*m.GetPlanType()).String()
        err := writer.WriteStringValue("plan_type", &cast)
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
// SetAssignee sets the assignee property value. A GitHub user.
func (m *CopilotSeatDetails) SetAssignee(value SimpleUserable)() {
    m.assignee = value
}
// SetAssigningTeam sets the assigning_team property value. The team through which the assignee is granted access to GitHub Copilot, if applicable.
func (m *CopilotSeatDetails) SetAssigningTeam(value CopilotSeatDetails_CopilotSeatDetails_assigning_teamable)() {
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
// SetOrganization sets the organization property value. A GitHub organization.
func (m *CopilotSeatDetails) SetOrganization(value NullableOrganizationSimpleable)() {
    m.organization = value
}
// SetPendingCancellationDate sets the pending_cancellation_date property value. The pending cancellation date for the seat, in `YYYY-MM-DD` format. This will be null unless the assignee's Copilot access has been canceled during the current billing cycle. If the seat has been cancelled, this corresponds to the start of the organization's next billing cycle.
func (m *CopilotSeatDetails) SetPendingCancellationDate(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)() {
    m.pending_cancellation_date = value
}
// SetPlanType sets the plan_type property value. The Copilot plan of the organization, or the parent enterprise, when applicable.
func (m *CopilotSeatDetails) SetPlanType(value *CopilotSeatDetails_plan_type)() {
    m.plan_type = value
}
// SetUpdatedAt sets the updated_at property value. **Closing down notice:** This field is no longer relevant and is closing down. Use the `created_at` field to determine when the assignee was last granted access to GitHub Copilot. Timestamp of when the assignee's GitHub Copilot access was last updated, in ISO 8601 format.
// Deprecated: 
func (m *CopilotSeatDetails) SetUpdatedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.updated_at = value
}
type CopilotSeatDetailsable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAssignee()(SimpleUserable)
    GetAssigningTeam()(CopilotSeatDetails_CopilotSeatDetails_assigning_teamable)
    GetCreatedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetLastActivityAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetLastActivityEditor()(*string)
    GetOrganization()(NullableOrganizationSimpleable)
    GetPendingCancellationDate()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)
    GetPlanType()(*CopilotSeatDetails_plan_type)
    GetUpdatedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    SetAssignee(value SimpleUserable)()
    SetAssigningTeam(value CopilotSeatDetails_CopilotSeatDetails_assigning_teamable)()
    SetCreatedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetLastActivityAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetLastActivityEditor(value *string)()
    SetOrganization(value NullableOrganizationSimpleable)()
    SetPendingCancellationDate(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)()
    SetPlanType(value *CopilotSeatDetails_plan_type)()
    SetUpdatedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
}
