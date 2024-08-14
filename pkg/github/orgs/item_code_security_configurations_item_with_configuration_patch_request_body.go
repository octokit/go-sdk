package orgs

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type ItemCodeSecurityConfigurationsItemWithConfiguration_PatchRequestBody struct {
    // Feature options for Automatic dependency submission
    dependency_graph_autosubmit_action_options ItemCodeSecurityConfigurationsItemWithConfiguration_PatchRequestBody_dependency_graph_autosubmit_action_optionsable
    // A description of the code security configuration
    description *string
    // The name of the code security configuration. Must be unique within the organization.
    name *string
}
// NewItemCodeSecurityConfigurationsItemWithConfiguration_PatchRequestBody instantiates a new ItemCodeSecurityConfigurationsItemWithConfiguration_PatchRequestBody and sets the default values.
func NewItemCodeSecurityConfigurationsItemWithConfiguration_PatchRequestBody()(*ItemCodeSecurityConfigurationsItemWithConfiguration_PatchRequestBody) {
    m := &ItemCodeSecurityConfigurationsItemWithConfiguration_PatchRequestBody{
    }
    return m
}
// CreateItemCodeSecurityConfigurationsItemWithConfiguration_PatchRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateItemCodeSecurityConfigurationsItemWithConfiguration_PatchRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemCodeSecurityConfigurationsItemWithConfiguration_PatchRequestBody(), nil
}
// GetDependencyGraphAutosubmitActionOptions gets the dependency_graph_autosubmit_action_options property value. Feature options for Automatic dependency submission
// returns a ItemCodeSecurityConfigurationsItemWithConfiguration_PatchRequestBody_dependency_graph_autosubmit_action_optionsable when successful
func (m *ItemCodeSecurityConfigurationsItemWithConfiguration_PatchRequestBody) GetDependencyGraphAutosubmitActionOptions()(ItemCodeSecurityConfigurationsItemWithConfiguration_PatchRequestBody_dependency_graph_autosubmit_action_optionsable) {
    return m.dependency_graph_autosubmit_action_options
}
// GetDescription gets the description property value. A description of the code security configuration
// returns a *string when successful
func (m *ItemCodeSecurityConfigurationsItemWithConfiguration_PatchRequestBody) GetDescription()(*string) {
    return m.description
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ItemCodeSecurityConfigurationsItemWithConfiguration_PatchRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["dependency_graph_autosubmit_action_options"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateItemCodeSecurityConfigurationsItemWithConfiguration_PatchRequestBody_dependency_graph_autosubmit_action_optionsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDependencyGraphAutosubmitActionOptions(val.(ItemCodeSecurityConfigurationsItemWithConfiguration_PatchRequestBody_dependency_graph_autosubmit_action_optionsable))
        }
        return nil
    }
    res["description"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDescription(val)
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
// GetName gets the name property value. The name of the code security configuration. Must be unique within the organization.
// returns a *string when successful
func (m *ItemCodeSecurityConfigurationsItemWithConfiguration_PatchRequestBody) GetName()(*string) {
    return m.name
}
// Serialize serializes information the current object
func (m *ItemCodeSecurityConfigurationsItemWithConfiguration_PatchRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("dependency_graph_autosubmit_action_options", m.GetDependencyGraphAutosubmitActionOptions())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("description", m.GetDescription())
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
    return nil
}
// SetDependencyGraphAutosubmitActionOptions sets the dependency_graph_autosubmit_action_options property value. Feature options for Automatic dependency submission
func (m *ItemCodeSecurityConfigurationsItemWithConfiguration_PatchRequestBody) SetDependencyGraphAutosubmitActionOptions(value ItemCodeSecurityConfigurationsItemWithConfiguration_PatchRequestBody_dependency_graph_autosubmit_action_optionsable)() {
    m.dependency_graph_autosubmit_action_options = value
}
// SetDescription sets the description property value. A description of the code security configuration
func (m *ItemCodeSecurityConfigurationsItemWithConfiguration_PatchRequestBody) SetDescription(value *string)() {
    m.description = value
}
// SetName sets the name property value. The name of the code security configuration. Must be unique within the organization.
func (m *ItemCodeSecurityConfigurationsItemWithConfiguration_PatchRequestBody) SetName(value *string)() {
    m.name = value
}
type ItemCodeSecurityConfigurationsItemWithConfiguration_PatchRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetDependencyGraphAutosubmitActionOptions()(ItemCodeSecurityConfigurationsItemWithConfiguration_PatchRequestBody_dependency_graph_autosubmit_action_optionsable)
    GetDescription()(*string)
    GetName()(*string)
    SetDependencyGraphAutosubmitActionOptions(value ItemCodeSecurityConfigurationsItemWithConfiguration_PatchRequestBody_dependency_graph_autosubmit_action_optionsable)()
    SetDescription(value *string)()
    SetName(value *string)()
}
