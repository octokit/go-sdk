package orgs

import (
    i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e "github.com/octokit/go-sdk/github/octokit/models"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemInstallationsGetResponse 
type ItemInstallationsGetResponse struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The installations property
    installations []i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.Installationable
    // The total_count property
    total_count *int32
}
// NewItemInstallationsGetResponse instantiates a new ItemInstallationsGetResponse and sets the default values.
func NewItemInstallationsGetResponse()(*ItemInstallationsGetResponse) {
    m := &ItemInstallationsGetResponse{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateItemInstallationsGetResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemInstallationsGetResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemInstallationsGetResponse(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ItemInstallationsGetResponse) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ItemInstallationsGetResponse) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["installations"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.CreateInstallationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.Installationable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.Installationable)
                }
            }
            m.SetInstallations(res)
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
// GetInstallations gets the installations property value. The installations property
func (m *ItemInstallationsGetResponse) GetInstallations()([]i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.Installationable) {
    return m.installations
}
// GetTotalCount gets the total_count property value. The total_count property
func (m *ItemInstallationsGetResponse) GetTotalCount()(*int32) {
    return m.total_count
}
// Serialize serializes information the current object
func (m *ItemInstallationsGetResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetInstallations() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetInstallations()))
        for i, v := range m.GetInstallations() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("installations", cast)
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
func (m *ItemInstallationsGetResponse) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetInstallations sets the installations property value. The installations property
func (m *ItemInstallationsGetResponse) SetInstallations(value []i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.Installationable)() {
    m.installations = value
}
// SetTotalCount sets the total_count property value. The total_count property
func (m *ItemInstallationsGetResponse) SetTotalCount(value *int32)() {
    m.total_count = value
}
// ItemInstallationsGetResponseable 
type ItemInstallationsGetResponseable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetInstallations()([]i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.Installationable)
    GetTotalCount()(*int32)
    SetInstallations(value []i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.Installationable)()
    SetTotalCount(value *int32)()
}
