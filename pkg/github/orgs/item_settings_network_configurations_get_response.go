package orgs

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6 "github.com/octokit/go-sdk/pkg/github/models"
)

type ItemSettingsNetworkConfigurationsGetResponse struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The network_configurations property
    network_configurations []i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.NetworkConfigurationable
    // The total_count property
    total_count *int32
}
// NewItemSettingsNetworkConfigurationsGetResponse instantiates a new ItemSettingsNetworkConfigurationsGetResponse and sets the default values.
func NewItemSettingsNetworkConfigurationsGetResponse()(*ItemSettingsNetworkConfigurationsGetResponse) {
    m := &ItemSettingsNetworkConfigurationsGetResponse{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateItemSettingsNetworkConfigurationsGetResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateItemSettingsNetworkConfigurationsGetResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemSettingsNetworkConfigurationsGetResponse(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *ItemSettingsNetworkConfigurationsGetResponse) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ItemSettingsNetworkConfigurationsGetResponse) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["network_configurations"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CreateNetworkConfigurationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.NetworkConfigurationable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.NetworkConfigurationable)
                }
            }
            m.SetNetworkConfigurations(res)
        }
        return nil
    }
    res["total_count"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTotalCount(val)
        }
        return nil
    }
    return res
}
// GetNetworkConfigurations gets the network_configurations property value. The network_configurations property
// returns a []NetworkConfigurationable when successful
func (m *ItemSettingsNetworkConfigurationsGetResponse) GetNetworkConfigurations()([]i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.NetworkConfigurationable) {
    return m.network_configurations
}
// GetTotalCount gets the total_count property value. The total_count property
// returns a *int32 when successful
func (m *ItemSettingsNetworkConfigurationsGetResponse) GetTotalCount()(*int32) {
    return m.total_count
}
// Serialize serializes information the current object
func (m *ItemSettingsNetworkConfigurationsGetResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetNetworkConfigurations() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetNetworkConfigurations()))
        for i, v := range m.GetNetworkConfigurations() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("network_configurations", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("total_count", m.GetTotalCount())
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
func (m *ItemSettingsNetworkConfigurationsGetResponse) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetNetworkConfigurations sets the network_configurations property value. The network_configurations property
func (m *ItemSettingsNetworkConfigurationsGetResponse) SetNetworkConfigurations(value []i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.NetworkConfigurationable)() {
    m.network_configurations = value
}
// SetTotalCount sets the total_count property value. The total_count property
func (m *ItemSettingsNetworkConfigurationsGetResponse) SetTotalCount(value *int32)() {
    m.total_count = value
}
type ItemSettingsNetworkConfigurationsGetResponseable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetNetworkConfigurations()([]i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.NetworkConfigurationable)
    GetTotalCount()(*int32)
    SetNetworkConfigurations(value []i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.NetworkConfigurationable)()
    SetTotalCount(value *int32)()
}
