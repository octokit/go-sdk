package enterprises

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemCodeSecurityConfigurationsItemWithConfiguration_PatchRequestBody_dependency_graph_autosubmit_action_options feature options for Automatic dependency submission
type ItemCodeSecurityConfigurationsItemWithConfiguration_PatchRequestBody_dependency_graph_autosubmit_action_options struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Whether to use runners labeled with 'dependency-submission' or standard GitHub runners.
    labeled_runners *bool
}
// NewItemCodeSecurityConfigurationsItemWithConfiguration_PatchRequestBody_dependency_graph_autosubmit_action_options instantiates a new ItemCodeSecurityConfigurationsItemWithConfiguration_PatchRequestBody_dependency_graph_autosubmit_action_options and sets the default values.
func NewItemCodeSecurityConfigurationsItemWithConfiguration_PatchRequestBody_dependency_graph_autosubmit_action_options()(*ItemCodeSecurityConfigurationsItemWithConfiguration_PatchRequestBody_dependency_graph_autosubmit_action_options) {
    m := &ItemCodeSecurityConfigurationsItemWithConfiguration_PatchRequestBody_dependency_graph_autosubmit_action_options{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateItemCodeSecurityConfigurationsItemWithConfiguration_PatchRequestBody_dependency_graph_autosubmit_action_optionsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateItemCodeSecurityConfigurationsItemWithConfiguration_PatchRequestBody_dependency_graph_autosubmit_action_optionsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemCodeSecurityConfigurationsItemWithConfiguration_PatchRequestBody_dependency_graph_autosubmit_action_options(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *ItemCodeSecurityConfigurationsItemWithConfiguration_PatchRequestBody_dependency_graph_autosubmit_action_options) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ItemCodeSecurityConfigurationsItemWithConfiguration_PatchRequestBody_dependency_graph_autosubmit_action_options) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["labeled_runners"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLabeledRunners(val)
        }
        return nil
    }
    return res
}
// GetLabeledRunners gets the labeled_runners property value. Whether to use runners labeled with 'dependency-submission' or standard GitHub runners.
// returns a *bool when successful
func (m *ItemCodeSecurityConfigurationsItemWithConfiguration_PatchRequestBody_dependency_graph_autosubmit_action_options) GetLabeledRunners()(*bool) {
    return m.labeled_runners
}
// Serialize serializes information the current object
func (m *ItemCodeSecurityConfigurationsItemWithConfiguration_PatchRequestBody_dependency_graph_autosubmit_action_options) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteBoolValue("labeled_runners", m.GetLabeledRunners())
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
func (m *ItemCodeSecurityConfigurationsItemWithConfiguration_PatchRequestBody_dependency_graph_autosubmit_action_options) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetLabeledRunners sets the labeled_runners property value. Whether to use runners labeled with 'dependency-submission' or standard GitHub runners.
func (m *ItemCodeSecurityConfigurationsItemWithConfiguration_PatchRequestBody_dependency_graph_autosubmit_action_options) SetLabeledRunners(value *bool)() {
    m.labeled_runners = value
}
type ItemCodeSecurityConfigurationsItemWithConfiguration_PatchRequestBody_dependency_graph_autosubmit_action_optionsable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetLabeledRunners()(*bool)
    SetLabeledRunners(value *bool)()
}
