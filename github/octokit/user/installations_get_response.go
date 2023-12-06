package user

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    i000736ae6dd74f01081193e4f903216bc2bd2954ed818433b986f45d581ed035 "github.com/octokit/go-sdk/github/octokit/models"
)

// InstallationsGetResponse 
type InstallationsGetResponse struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The installations property
    installations []i000736ae6dd74f01081193e4f903216bc2bd2954ed818433b986f45d581ed035.Installationable
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
        val, err := n.GetCollectionOfObjectValues(i000736ae6dd74f01081193e4f903216bc2bd2954ed818433b986f45d581ed035.CreateInstallationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]i000736ae6dd74f01081193e4f903216bc2bd2954ed818433b986f45d581ed035.Installationable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(i000736ae6dd74f01081193e4f903216bc2bd2954ed818433b986f45d581ed035.Installationable)
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
func (m *InstallationsGetResponse) GetInstallations()([]i000736ae6dd74f01081193e4f903216bc2bd2954ed818433b986f45d581ed035.Installationable) {
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
func (m *InstallationsGetResponse) SetInstallations(value []i000736ae6dd74f01081193e4f903216bc2bd2954ed818433b986f45d581ed035.Installationable)() {
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
    GetInstallations()([]i000736ae6dd74f01081193e4f903216bc2bd2954ed818433b986f45d581ed035.Installationable)
    GetTotalCount()(*int32)
    SetInstallations(value []i000736ae6dd74f01081193e4f903216bc2bd2954ed818433b986f45d581ed035.Installationable)()
    SetTotalCount(value *int32)()
}
