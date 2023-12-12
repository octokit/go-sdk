package user

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6 "github.com/octokit/go-sdk/pkg/github/models"
)

// InstallationsGetResponse 
type InstallationsGetResponse struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The installations property
    installations []i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.Installationable
    // The total_count property
    total_count *int32
}
// NewInstallationsGetResponse instantiates a new InstallationsGetResponse and sets the default values.
func NewInstallationsGetResponse()(*InstallationsGetResponse) {
    m := &InstallationsGetResponse{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateInstallationsGetResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateInstallationsGetResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewInstallationsGetResponse(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *InstallationsGetResponse) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *InstallationsGetResponse) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["installations"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CreateInstallationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.Installationable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.Installationable)
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
func (m *InstallationsGetResponse) GetInstallations()([]i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.Installationable) {
    return m.installations
}
// GetTotalCount gets the total_count property value. The total_count property
func (m *InstallationsGetResponse) GetTotalCount()(*int32) {
    return m.total_count
}
// Serialize serializes information the current object
func (m *InstallationsGetResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
func (m *InstallationsGetResponse) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetInstallations sets the installations property value. The installations property
func (m *InstallationsGetResponse) SetInstallations(value []i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.Installationable)() {
    m.installations = value
}
// SetTotalCount sets the total_count property value. The total_count property
func (m *InstallationsGetResponse) SetTotalCount(value *int32)() {
    m.total_count = value
}
// InstallationsGetResponseable 
type InstallationsGetResponseable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetInstallations()([]i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.Installationable)
    GetTotalCount()(*int32)
    SetInstallations(value []i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.Installationable)()
    SetTotalCount(value *int32)()
}
