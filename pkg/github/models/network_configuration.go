package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// NetworkConfiguration a hosted compute network configuration.
type NetworkConfiguration struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The hosted compute service the network configuration supports.
    compute_service *NetworkConfiguration_compute_service
    // The time at which the network configuration was created, in ISO 8601 format.
    created_on *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The unique identifier of the network configuration.
    id *string
    // The name of the network configuration.
    name *string
    // The unique identifier of each network settings in the configuration.
    network_settings_ids []string
}
// NewNetworkConfiguration instantiates a new NetworkConfiguration and sets the default values.
func NewNetworkConfiguration()(*NetworkConfiguration) {
    m := &NetworkConfiguration{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateNetworkConfigurationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateNetworkConfigurationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewNetworkConfiguration(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *NetworkConfiguration) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetComputeService gets the compute_service property value. The hosted compute service the network configuration supports.
// returns a *NetworkConfiguration_compute_service when successful
func (m *NetworkConfiguration) GetComputeService()(*NetworkConfiguration_compute_service) {
    return m.compute_service
}
// GetCreatedOn gets the created_on property value. The time at which the network configuration was created, in ISO 8601 format.
// returns a *Time when successful
func (m *NetworkConfiguration) GetCreatedOn()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.created_on
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *NetworkConfiguration) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["compute_service"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseNetworkConfiguration_compute_service)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetComputeService(val.(*NetworkConfiguration_compute_service))
        }
        return nil
    }
    res["created_on"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCreatedOn(val)
        }
        return nil
    }
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
// GetId gets the id property value. The unique identifier of the network configuration.
// returns a *string when successful
func (m *NetworkConfiguration) GetId()(*string) {
    return m.id
}
// GetName gets the name property value. The name of the network configuration.
// returns a *string when successful
func (m *NetworkConfiguration) GetName()(*string) {
    return m.name
}
// GetNetworkSettingsIds gets the network_settings_ids property value. The unique identifier of each network settings in the configuration.
// returns a []string when successful
func (m *NetworkConfiguration) GetNetworkSettingsIds()([]string) {
    return m.network_settings_ids
}
// Serialize serializes information the current object
func (m *NetworkConfiguration) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetComputeService() != nil {
        cast := (*m.GetComputeService()).String()
        err := writer.WriteStringValue("compute_service", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteTimeValue("created_on", m.GetCreatedOn())
        if err != nil {
            return err
        }
    }
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
func (m *NetworkConfiguration) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetComputeService sets the compute_service property value. The hosted compute service the network configuration supports.
func (m *NetworkConfiguration) SetComputeService(value *NetworkConfiguration_compute_service)() {
    m.compute_service = value
}
// SetCreatedOn sets the created_on property value. The time at which the network configuration was created, in ISO 8601 format.
func (m *NetworkConfiguration) SetCreatedOn(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.created_on = value
}
// SetId sets the id property value. The unique identifier of the network configuration.
func (m *NetworkConfiguration) SetId(value *string)() {
    m.id = value
}
// SetName sets the name property value. The name of the network configuration.
func (m *NetworkConfiguration) SetName(value *string)() {
    m.name = value
}
// SetNetworkSettingsIds sets the network_settings_ids property value. The unique identifier of each network settings in the configuration.
func (m *NetworkConfiguration) SetNetworkSettingsIds(value []string)() {
    m.network_settings_ids = value
}
type NetworkConfigurationable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetComputeService()(*NetworkConfiguration_compute_service)
    GetCreatedOn()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetId()(*string)
    GetName()(*string)
    GetNetworkSettingsIds()([]string)
    SetComputeService(value *NetworkConfiguration_compute_service)()
    SetCreatedOn(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetId(value *string)()
    SetName(value *string)()
    SetNetworkSettingsIds(value []string)()
}
