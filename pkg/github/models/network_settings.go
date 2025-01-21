package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// NetworkSettings a hosted compute network settings resource.
type NetworkSettings struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The unique identifier of the network settings resource.
    id *string
    // The name of the network settings resource.
    name *string
    // The identifier of the network configuration that is using this settings resource.
    network_configuration_id *string
    // The location of the subnet this network settings resource is configured for.
    region *string
    // The subnet this network settings resource is configured for.
    subnet_id *string
}
// NewNetworkSettings instantiates a new NetworkSettings and sets the default values.
func NewNetworkSettings()(*NetworkSettings) {
    m := &NetworkSettings{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateNetworkSettingsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateNetworkSettingsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewNetworkSettings(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *NetworkSettings) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *NetworkSettings) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["id"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetId(val)
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
    res["network_configuration_id"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNetworkConfigurationId(val)
        }
        return nil
    }
    res["region"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRegion(val)
        }
        return nil
    }
    res["subnet_id"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSubnetId(val)
        }
        return nil
    }
    return res
}
// GetId gets the id property value. The unique identifier of the network settings resource.
// returns a *string when successful
func (m *NetworkSettings) GetId()(*string) {
    return m.id
}
// GetName gets the name property value. The name of the network settings resource.
// returns a *string when successful
func (m *NetworkSettings) GetName()(*string) {
    return m.name
}
// GetNetworkConfigurationId gets the network_configuration_id property value. The identifier of the network configuration that is using this settings resource.
// returns a *string when successful
func (m *NetworkSettings) GetNetworkConfigurationId()(*string) {
    return m.network_configuration_id
}
// GetRegion gets the region property value. The location of the subnet this network settings resource is configured for.
// returns a *string when successful
func (m *NetworkSettings) GetRegion()(*string) {
    return m.region
}
// GetSubnetId gets the subnet_id property value. The subnet this network settings resource is configured for.
// returns a *string when successful
func (m *NetworkSettings) GetSubnetId()(*string) {
    return m.subnet_id
}
// Serialize serializes information the current object
func (m *NetworkSettings) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("id", m.GetId())
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
        err := writer.WriteStringValue("network_configuration_id", m.GetNetworkConfigurationId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("region", m.GetRegion())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("subnet_id", m.GetSubnetId())
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
func (m *NetworkSettings) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetId sets the id property value. The unique identifier of the network settings resource.
func (m *NetworkSettings) SetId(value *string)() {
    m.id = value
}
// SetName sets the name property value. The name of the network settings resource.
func (m *NetworkSettings) SetName(value *string)() {
    m.name = value
}
// SetNetworkConfigurationId sets the network_configuration_id property value. The identifier of the network configuration that is using this settings resource.
func (m *NetworkSettings) SetNetworkConfigurationId(value *string)() {
    m.network_configuration_id = value
}
// SetRegion sets the region property value. The location of the subnet this network settings resource is configured for.
func (m *NetworkSettings) SetRegion(value *string)() {
    m.region = value
}
// SetSubnetId sets the subnet_id property value. The subnet this network settings resource is configured for.
func (m *NetworkSettings) SetSubnetId(value *string)() {
    m.subnet_id = value
}
type NetworkSettingsable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetId()(*string)
    GetName()(*string)
    GetNetworkConfigurationId()(*string)
    GetRegion()(*string)
    GetSubnetId()(*string)
    SetId(value *string)()
    SetName(value *string)()
    SetNetworkConfigurationId(value *string)()
    SetRegion(value *string)()
    SetSubnetId(value *string)()
}
