package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type ActionsHostedRunnerLimits struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Provides details of static public IP limits for GitHub-hosted Hosted Runners
    public_ips ActionsHostedRunnerLimits_public_ipsable
}
// NewActionsHostedRunnerLimits instantiates a new ActionsHostedRunnerLimits and sets the default values.
func NewActionsHostedRunnerLimits()(*ActionsHostedRunnerLimits) {
    m := &ActionsHostedRunnerLimits{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateActionsHostedRunnerLimitsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateActionsHostedRunnerLimitsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewActionsHostedRunnerLimits(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *ActionsHostedRunnerLimits) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ActionsHostedRunnerLimits) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["public_ips"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateActionsHostedRunnerLimits_public_ipsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPublicIps(val.(ActionsHostedRunnerLimits_public_ipsable))
        }
        return nil
    }
    return res
}
// GetPublicIps gets the public_ips property value. Provides details of static public IP limits for GitHub-hosted Hosted Runners
// returns a ActionsHostedRunnerLimits_public_ipsable when successful
func (m *ActionsHostedRunnerLimits) GetPublicIps()(ActionsHostedRunnerLimits_public_ipsable) {
    return m.public_ips
}
// Serialize serializes information the current object
func (m *ActionsHostedRunnerLimits) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("public_ips", m.GetPublicIps())
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
func (m *ActionsHostedRunnerLimits) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetPublicIps sets the public_ips property value. Provides details of static public IP limits for GitHub-hosted Hosted Runners
func (m *ActionsHostedRunnerLimits) SetPublicIps(value ActionsHostedRunnerLimits_public_ipsable)() {
    m.public_ips = value
}
type ActionsHostedRunnerLimitsable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetPublicIps()(ActionsHostedRunnerLimits_public_ipsable)
    SetPublicIps(value ActionsHostedRunnerLimits_public_ipsable)()
}
