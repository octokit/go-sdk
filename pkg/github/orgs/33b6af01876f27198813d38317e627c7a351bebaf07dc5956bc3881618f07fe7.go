package orgs

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type ItemSettingsNetworkConfigurationsItemWithNetwork_configuration_PatchRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Name of the network configuration. Must be between 1 and 100 characters and may only contain upper and lowercase letters a-z, numbers 0-9, '.', '-', and '_'.
    name *string
    // The identifier of the network settings to use for the network configuration. Exactly one network settings must be specified.
    network_settings_ids []string
}
// NewItemSettingsNetworkConfigurationsItemWithNetwork_configuration_PatchRequestBody instantiates a new ItemSettingsNetworkConfigurationsItemWithNetwork_configuration_PatchRequestBody and sets the default values.
func NewItemSettingsNetworkConfigurationsItemWithNetwork_configuration_PatchRequestBody()(*ItemSettingsNetworkConfigurationsItemWithNetwork_configuration_PatchRequestBody) {
    m := &ItemSettingsNetworkConfigurationsItemWithNetwork_configuration_PatchRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateItemSettingsNetworkConfigurationsItemWithNetwork_configuration_PatchRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateItemSettingsNetworkConfigurationsItemWithNetwork_configuration_PatchRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemSettingsNetworkConfigurationsItemWithNetwork_configuration_PatchRequestBody(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *ItemSettingsNetworkConfigurationsItemWithNetwork_configuration_PatchRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ItemSettingsNetworkConfigurationsItemWithNetwork_configuration_PatchRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
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
    res["network_settings_ids"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = *(v.(*string))
                }
            }
            m.SetNetworkSettingsIds(res)
        }
        return nil
    }
    return res
}
// GetName gets the name property value. Name of the network configuration. Must be between 1 and 100 characters and may only contain upper and lowercase letters a-z, numbers 0-9, '.', '-', and '_'.
// returns a *string when successful
func (m *ItemSettingsNetworkConfigurationsItemWithNetwork_configuration_PatchRequestBody) GetName()(*string) {
    return m.name
}
// GetNetworkSettingsIds gets the network_settings_ids property value. The identifier of the network settings to use for the network configuration. Exactly one network settings must be specified.
// returns a []string when successful
func (m *ItemSettingsNetworkConfigurationsItemWithNetwork_configuration_PatchRequestBody) GetNetworkSettingsIds()([]string) {
    return m.network_settings_ids
}
// Serialize serializes information the current object
func (m *ItemSettingsNetworkConfigurationsItemWithNetwork_configuration_PatchRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("name", m.GetName())
        if err != nil {
            return err
        }
    }
    if m.GetNetworkSettingsIds() != nil {
        err := writer.WriteCollectionOfStringValues("network_settings_ids", m.GetNetworkSettingsIds())
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
func (m *ItemSettingsNetworkConfigurationsItemWithNetwork_configuration_PatchRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetName sets the name property value. Name of the network configuration. Must be between 1 and 100 characters and may only contain upper and lowercase letters a-z, numbers 0-9, '.', '-', and '_'.
func (m *ItemSettingsNetworkConfigurationsItemWithNetwork_configuration_PatchRequestBody) SetName(value *string)() {
    m.name = value
}
// SetNetworkSettingsIds sets the network_settings_ids property value. The identifier of the network settings to use for the network configuration. Exactly one network settings must be specified.
func (m *ItemSettingsNetworkConfigurationsItemWithNetwork_configuration_PatchRequestBody) SetNetworkSettingsIds(value []string)() {
    m.network_settings_ids = value
}
type ItemSettingsNetworkConfigurationsItemWithNetwork_configuration_PatchRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetName()(*string)
    GetNetworkSettingsIds()([]string)
    SetName(value *string)()
    SetNetworkSettingsIds(value []string)()
}
