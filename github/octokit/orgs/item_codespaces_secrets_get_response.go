package orgs

import (
    i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e "github.com/octokit/go-sdk/github/octokit/models"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemCodespacesSecretsGetResponse 
type ItemCodespacesSecretsGetResponse struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The secrets property
    secrets []i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.CodespacesOrgSecretable
    // The total_count property
    total_count *int32
}
// NewItemCodespacesSecretsGetResponse instantiates a new ItemCodespacesSecretsGetResponse and sets the default values.
func NewItemCodespacesSecretsGetResponse()(*ItemCodespacesSecretsGetResponse) {
    m := &ItemCodespacesSecretsGetResponse{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateItemCodespacesSecretsGetResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemCodespacesSecretsGetResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemCodespacesSecretsGetResponse(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ItemCodespacesSecretsGetResponse) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ItemCodespacesSecretsGetResponse) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["secrets"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.CreateCodespacesOrgSecretFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.CodespacesOrgSecretable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.CodespacesOrgSecretable)
                }
            }
            m.SetSecrets(res)
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
// GetSecrets gets the secrets property value. The secrets property
func (m *ItemCodespacesSecretsGetResponse) GetSecrets()([]i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.CodespacesOrgSecretable) {
    return m.secrets
}
// GetTotalCount gets the total_count property value. The total_count property
func (m *ItemCodespacesSecretsGetResponse) GetTotalCount()(*int32) {
    return m.total_count
}
// Serialize serializes information the current object
func (m *ItemCodespacesSecretsGetResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetSecrets() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetSecrets()))
        for i, v := range m.GetSecrets() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("secrets", cast)
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
func (m *ItemCodespacesSecretsGetResponse) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetSecrets sets the secrets property value. The secrets property
func (m *ItemCodespacesSecretsGetResponse) SetSecrets(value []i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.CodespacesOrgSecretable)() {
    m.secrets = value
}
// SetTotalCount sets the total_count property value. The total_count property
func (m *ItemCodespacesSecretsGetResponse) SetTotalCount(value *int32)() {
    m.total_count = value
}
// ItemCodespacesSecretsGetResponseable 
type ItemCodespacesSecretsGetResponseable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetSecrets()([]i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.CodespacesOrgSecretable)
    GetTotalCount()(*int32)
    SetSecrets(value []i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.CodespacesOrgSecretable)()
    SetTotalCount(value *int32)()
}
