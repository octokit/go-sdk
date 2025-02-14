package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// CopilotOrganizationSeatBreakdown the breakdown of Copilot Business seats for the organization.
type CopilotOrganizationSeatBreakdown struct {
    // The number of seats that have used Copilot during the current billing cycle.
    active_this_cycle *int32
    // Seats added during the current billing cycle.
    added_this_cycle *int32
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The number of seats that have not used Copilot during the current billing cycle.
    inactive_this_cycle *int32
    // The number of seats that are pending cancellation at the end of the current billing cycle.
    pending_cancellation *int32
    // The number of users who have been invited to receive a Copilot seat through this organization.
    pending_invitation *int32
    // The total number of seats being billed for the organization as of the current billing cycle.
    total *int32
}
// NewCopilotOrganizationSeatBreakdown instantiates a new CopilotOrganizationSeatBreakdown and sets the default values.
func NewCopilotOrganizationSeatBreakdown()(*CopilotOrganizationSeatBreakdown) {
    m := &CopilotOrganizationSeatBreakdown{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateCopilotOrganizationSeatBreakdownFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateCopilotOrganizationSeatBreakdownFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCopilotOrganizationSeatBreakdown(), nil
}
// GetActiveThisCycle gets the active_this_cycle property value. The number of seats that have used Copilot during the current billing cycle.
// returns a *int32 when successful
func (m *CopilotOrganizationSeatBreakdown) GetActiveThisCycle()(*int32) {
    return m.active_this_cycle
}
// GetAddedThisCycle gets the added_this_cycle property value. Seats added during the current billing cycle.
// returns a *int32 when successful
func (m *CopilotOrganizationSeatBreakdown) GetAddedThisCycle()(*int32) {
    return m.added_this_cycle
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *CopilotOrganizationSeatBreakdown) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *CopilotOrganizationSeatBreakdown) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["active_this_cycle"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetActiveThisCycle(val)
        }
        return nil
    }
    res["added_this_cycle"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAddedThisCycle(val)
        }
        return nil
    }
    res["inactive_this_cycle"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetInactiveThisCycle(val)
        }
        return nil
    }
    res["pending_cancellation"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPendingCancellation(val)
        }
        return nil
    }
    res["pending_invitation"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPendingInvitation(val)
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
// GetInactiveThisCycle gets the inactive_this_cycle property value. The number of seats that have not used Copilot during the current billing cycle.
// returns a *int32 when successful
func (m *CopilotOrganizationSeatBreakdown) GetInactiveThisCycle()(*int32) {
    return m.inactive_this_cycle
}
// GetPendingCancellation gets the pending_cancellation property value. The number of seats that are pending cancellation at the end of the current billing cycle.
// returns a *int32 when successful
func (m *CopilotOrganizationSeatBreakdown) GetPendingCancellation()(*int32) {
    return m.pending_cancellation
}
// GetPendingInvitation gets the pending_invitation property value. The number of users who have been invited to receive a Copilot seat through this organization.
// returns a *int32 when successful
func (m *CopilotOrganizationSeatBreakdown) GetPendingInvitation()(*int32) {
    return m.pending_invitation
}
// GetTotal gets the total property value. The total number of seats being billed for the organization as of the current billing cycle.
// returns a *int32 when successful
func (m *CopilotOrganizationSeatBreakdown) GetTotal()(*int32) {
    return m.total
}
// Serialize serializes information the current object
func (m *CopilotOrganizationSeatBreakdown) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteInt32Value("active_this_cycle", m.GetActiveThisCycle())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("added_this_cycle", m.GetAddedThisCycle())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("inactive_this_cycle", m.GetInactiveThisCycle())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("pending_cancellation", m.GetPendingCancellation())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("pending_invitation", m.GetPendingInvitation())
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
// SetActiveThisCycle sets the active_this_cycle property value. The number of seats that have used Copilot during the current billing cycle.
func (m *CopilotOrganizationSeatBreakdown) SetActiveThisCycle(value *int32)() {
    m.active_this_cycle = value
}
// SetAddedThisCycle sets the added_this_cycle property value. Seats added during the current billing cycle.
func (m *CopilotOrganizationSeatBreakdown) SetAddedThisCycle(value *int32)() {
    m.added_this_cycle = value
}
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *CopilotOrganizationSeatBreakdown) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetInactiveThisCycle sets the inactive_this_cycle property value. The number of seats that have not used Copilot during the current billing cycle.
func (m *CopilotOrganizationSeatBreakdown) SetInactiveThisCycle(value *int32)() {
    m.inactive_this_cycle = value
}
// SetPendingCancellation sets the pending_cancellation property value. The number of seats that are pending cancellation at the end of the current billing cycle.
func (m *CopilotOrganizationSeatBreakdown) SetPendingCancellation(value *int32)() {
    m.pending_cancellation = value
}
// SetPendingInvitation sets the pending_invitation property value. The number of users who have been invited to receive a Copilot seat through this organization.
func (m *CopilotOrganizationSeatBreakdown) SetPendingInvitation(value *int32)() {
    m.pending_invitation = value
}
// SetTotal sets the total property value. The total number of seats being billed for the organization as of the current billing cycle.
func (m *CopilotOrganizationSeatBreakdown) SetTotal(value *int32)() {
    m.total = value
}
type CopilotOrganizationSeatBreakdownable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetActiveThisCycle()(*int32)
    GetAddedThisCycle()(*int32)
    GetInactiveThisCycle()(*int32)
    GetPendingCancellation()(*int32)
    GetPendingInvitation()(*int32)
    GetTotal()(*int32)
    SetActiveThisCycle(value *int32)()
    SetAddedThisCycle(value *int32)()
    SetInactiveThisCycle(value *int32)()
    SetPendingCancellation(value *int32)()
    SetPendingInvitation(value *int32)()
    SetTotal(value *int32)()
}
