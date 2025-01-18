package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ActionsHostedRunnerLimits_public_ips provides details of static public IP limits for GitHub-hosted Hosted Runners
type ActionsHostedRunnerLimits_public_ips struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The current number of static public IP addresses in use by Hosted Runners.
    current_usage *int32
    // The maximum number of static public IP addresses that can be used for Hosted Runners.
    maximum *int32
}
// NewActionsHostedRunnerLimits_public_ips instantiates a new ActionsHostedRunnerLimits_public_ips and sets the default values.
func NewActionsHostedRunnerLimits_public_ips()(*ActionsHostedRunnerLimits_public_ips) {
    m := &ActionsHostedRunnerLimits_public_ips{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateActionsHostedRunnerLimits_public_ipsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateActionsHostedRunnerLimits_public_ipsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewActionsHostedRunnerLimits_public_ips(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *ActionsHostedRunnerLimits_public_ips) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetCurrentUsage gets the current_usage property value. The current number of static public IP addresses in use by Hosted Runners.
// returns a *int32 when successful
func (m *ActionsHostedRunnerLimits_public_ips) GetCurrentUsage()(*int32) {
    return m.current_usage
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ActionsHostedRunnerLimits_public_ips) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["current_usage"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCurrentUsage(val)
        }
        return nil
    }
    res["maximum"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMaximum(val)
        }
        return nil
    }
    return res
}
// GetMaximum gets the maximum property value. The maximum number of static public IP addresses that can be used for Hosted Runners.
// returns a *int32 when successful
func (m *ActionsHostedRunnerLimits_public_ips) GetMaximum()(*int32) {
    return m.maximum
}
// Serialize serializes information the current object
func (m *ActionsHostedRunnerLimits_public_ips) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteInt32Value("current_usage", m.GetCurrentUsage())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("maximum", m.GetMaximum())
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
func (m *ActionsHostedRunnerLimits_public_ips) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetCurrentUsage sets the current_usage property value. The current number of static public IP addresses in use by Hosted Runners.
func (m *ActionsHostedRunnerLimits_public_ips) SetCurrentUsage(value *int32)() {
    m.current_usage = value
}
// SetMaximum sets the maximum property value. The maximum number of static public IP addresses that can be used for Hosted Runners.
func (m *ActionsHostedRunnerLimits_public_ips) SetMaximum(value *int32)() {
    m.maximum = value
}
type ActionsHostedRunnerLimits_public_ipsable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCurrentUsage()(*int32)
    GetMaximum()(*int32)
    SetCurrentUsage(value *int32)()
    SetMaximum(value *int32)()
}
