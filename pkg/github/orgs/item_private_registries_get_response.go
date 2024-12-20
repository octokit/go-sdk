package orgs

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6 "github.com/octokit/go-sdk/pkg/github/models"
)

type ItemPrivateRegistriesGetResponse struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The configurations property
    configurations []i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.OrgPrivateRegistryConfigurationable
    // The total_count property
    total_count *int32
}
// NewItemPrivateRegistriesGetResponse instantiates a new ItemPrivateRegistriesGetResponse and sets the default values.
func NewItemPrivateRegistriesGetResponse()(*ItemPrivateRegistriesGetResponse) {
    m := &ItemPrivateRegistriesGetResponse{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateItemPrivateRegistriesGetResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateItemPrivateRegistriesGetResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemPrivateRegistriesGetResponse(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *ItemPrivateRegistriesGetResponse) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetConfigurations gets the configurations property value. The configurations property
// returns a []OrgPrivateRegistryConfigurationable when successful
func (m *ItemPrivateRegistriesGetResponse) GetConfigurations()([]i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.OrgPrivateRegistryConfigurationable) {
    return m.configurations
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ItemPrivateRegistriesGetResponse) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["configurations"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CreateOrgPrivateRegistryConfigurationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.OrgPrivateRegistryConfigurationable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.OrgPrivateRegistryConfigurationable)
                }
            }
            m.SetConfigurations(res)
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
// GetTotalCount gets the total_count property value. The total_count property
// returns a *int32 when successful
func (m *ItemPrivateRegistriesGetResponse) GetTotalCount()(*int32) {
    return m.total_count
}
// Serialize serializes information the current object
func (m *ItemPrivateRegistriesGetResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetConfigurations() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetConfigurations()))
        for i, v := range m.GetConfigurations() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("configurations", cast)
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
func (m *ItemPrivateRegistriesGetResponse) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetConfigurations sets the configurations property value. The configurations property
func (m *ItemPrivateRegistriesGetResponse) SetConfigurations(value []i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.OrgPrivateRegistryConfigurationable)() {
    m.configurations = value
}
// SetTotalCount sets the total_count property value. The total_count property
func (m *ItemPrivateRegistriesGetResponse) SetTotalCount(value *int32)() {
    m.total_count = value
}
type ItemPrivateRegistriesGetResponseable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetConfigurations()([]i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.OrgPrivateRegistryConfigurationable)
    GetTotalCount()(*int32)
    SetConfigurations(value []i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.OrgPrivateRegistryConfigurationable)()
    SetTotalCount(value *int32)()
}
