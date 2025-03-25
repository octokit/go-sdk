package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type RulesetVersionWithState struct {
    RulesetVersion
    // The state of the ruleset version
    state RulesetVersionWithState_stateable
}
// NewRulesetVersionWithState instantiates a new RulesetVersionWithState and sets the default values.
func NewRulesetVersionWithState()(*RulesetVersionWithState) {
    m := &RulesetVersionWithState{
        RulesetVersion: *NewRulesetVersion(),
    }
    return m
}
// CreateRulesetVersionWithStateFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateRulesetVersionWithStateFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewRulesetVersionWithState(), nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *RulesetVersionWithState) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.RulesetVersion.GetFieldDeserializers()
    res["state"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateRulesetVersionWithState_stateFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetState(val.(RulesetVersionWithState_stateable))
        }
        return nil
    }
    return res
}
// GetState gets the state property value. The state of the ruleset version
// returns a RulesetVersionWithState_stateable when successful
func (m *RulesetVersionWithState) GetState()(RulesetVersionWithState_stateable) {
    return m.state
}
// Serialize serializes information the current object
func (m *RulesetVersionWithState) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.RulesetVersion.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("state", m.GetState())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetState sets the state property value. The state of the ruleset version
func (m *RulesetVersionWithState) SetState(value RulesetVersionWithState_stateable)() {
    m.state = value
}
type RulesetVersionWithStateable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    RulesetVersionable
    GetState()(RulesetVersionWithState_stateable)
    SetState(value RulesetVersionWithState_stateable)()
}
