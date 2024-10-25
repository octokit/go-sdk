package orgs

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type ItemCodeSecurityConfigurationsPostRequestBody struct {
    // Feature options for Automatic dependency submission
    dependency_graph_autosubmit_action_options ItemCodeSecurityConfigurationsPostRequestBody_dependency_graph_autosubmit_action_optionsable
    // A description of the code security configuration
    description *string
    // The name of the code security configuration. Must be unique within the organization.
    name *string
    // Feature options for secret scanning delegated bypass
    secret_scanning_delegated_bypass_options ItemCodeSecurityConfigurationsPostRequestBody_secret_scanning_delegated_bypass_optionsable
}
// NewItemCodeSecurityConfigurationsPostRequestBody instantiates a new ItemCodeSecurityConfigurationsPostRequestBody and sets the default values.
func NewItemCodeSecurityConfigurationsPostRequestBody()(*ItemCodeSecurityConfigurationsPostRequestBody) {
    m := &ItemCodeSecurityConfigurationsPostRequestBody{
    }
    return m
}
// CreateItemCodeSecurityConfigurationsPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateItemCodeSecurityConfigurationsPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemCodeSecurityConfigurationsPostRequestBody(), nil
}
// GetDependencyGraphAutosubmitActionOptions gets the dependency_graph_autosubmit_action_options property value. Feature options for Automatic dependency submission
// returns a ItemCodeSecurityConfigurationsPostRequestBody_dependency_graph_autosubmit_action_optionsable when successful
func (m *ItemCodeSecurityConfigurationsPostRequestBody) GetDependencyGraphAutosubmitActionOptions()(ItemCodeSecurityConfigurationsPostRequestBody_dependency_graph_autosubmit_action_optionsable) {
    return m.dependency_graph_autosubmit_action_options
}
// GetDescription gets the description property value. A description of the code security configuration
// returns a *string when successful
func (m *ItemCodeSecurityConfigurationsPostRequestBody) GetDescription()(*string) {
    return m.description
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ItemCodeSecurityConfigurationsPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["dependency_graph_autosubmit_action_options"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateItemCodeSecurityConfigurationsPostRequestBody_dependency_graph_autosubmit_action_optionsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDependencyGraphAutosubmitActionOptions(val.(ItemCodeSecurityConfigurationsPostRequestBody_dependency_graph_autosubmit_action_optionsable))
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
    res["secret_scanning_delegated_bypass_options"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateItemCodeSecurityConfigurationsPostRequestBody_secret_scanning_delegated_bypass_optionsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSecretScanningDelegatedBypassOptions(val.(ItemCodeSecurityConfigurationsPostRequestBody_secret_scanning_delegated_bypass_optionsable))
        }
        return nil
    }
    return res
}
// GetName gets the name property value. The name of the code security configuration. Must be unique within the organization.
// returns a *string when successful
func (m *ItemCodeSecurityConfigurationsPostRequestBody) GetName()(*string) {
    return m.name
}
// GetSecretScanningDelegatedBypassOptions gets the secret_scanning_delegated_bypass_options property value. Feature options for secret scanning delegated bypass
// returns a ItemCodeSecurityConfigurationsPostRequestBody_secret_scanning_delegated_bypass_optionsable when successful
func (m *ItemCodeSecurityConfigurationsPostRequestBody) GetSecretScanningDelegatedBypassOptions()(ItemCodeSecurityConfigurationsPostRequestBody_secret_scanning_delegated_bypass_optionsable) {
    return m.secret_scanning_delegated_bypass_options
}
// Serialize serializes information the current object
func (m *ItemCodeSecurityConfigurationsPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
    {
        err := writer.WriteObjectValue("secret_scanning_delegated_bypass_options", m.GetSecretScanningDelegatedBypassOptions())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetDependencyGraphAutosubmitActionOptions sets the dependency_graph_autosubmit_action_options property value. Feature options for Automatic dependency submission
func (m *ItemCodeSecurityConfigurationsPostRequestBody) SetDependencyGraphAutosubmitActionOptions(value ItemCodeSecurityConfigurationsPostRequestBody_dependency_graph_autosubmit_action_optionsable)() {
    m.dependency_graph_autosubmit_action_options = value
}
// SetDescription sets the description property value. A description of the code security configuration
func (m *ItemCodeSecurityConfigurationsPostRequestBody) SetDescription(value *string)() {
    m.description = value
}
// SetName sets the name property value. The name of the code security configuration. Must be unique within the organization.
func (m *ItemCodeSecurityConfigurationsPostRequestBody) SetName(value *string)() {
    m.name = value
}
// SetSecretScanningDelegatedBypassOptions sets the secret_scanning_delegated_bypass_options property value. Feature options for secret scanning delegated bypass
func (m *ItemCodeSecurityConfigurationsPostRequestBody) SetSecretScanningDelegatedBypassOptions(value ItemCodeSecurityConfigurationsPostRequestBody_secret_scanning_delegated_bypass_optionsable)() {
    m.secret_scanning_delegated_bypass_options = value
}
type ItemCodeSecurityConfigurationsPostRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetDependencyGraphAutosubmitActionOptions()(ItemCodeSecurityConfigurationsPostRequestBody_dependency_graph_autosubmit_action_optionsable)
    GetDescription()(*string)
    GetName()(*string)
    GetSecretScanningDelegatedBypassOptions()(ItemCodeSecurityConfigurationsPostRequestBody_secret_scanning_delegated_bypass_optionsable)
    SetDependencyGraphAutosubmitActionOptions(value ItemCodeSecurityConfigurationsPostRequestBody_dependency_graph_autosubmit_action_optionsable)()
    SetDescription(value *string)()
    SetName(value *string)()
    SetSecretScanningDelegatedBypassOptions(value ItemCodeSecurityConfigurationsPostRequestBody_secret_scanning_delegated_bypass_optionsable)()
}
