package user

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    i000736ae6dd74f01081193e4f903216bc2bd2954ed818433b986f45d581ed035 "github.com/octokit/go-sdk/github/octokit/models"
)

// CodespacesSecretsGetResponse 
type CodespacesSecretsGetResponse struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The secrets property
    secrets []i000736ae6dd74f01081193e4f903216bc2bd2954ed818433b986f45d581ed035.CodespacesSecretable
    // The total_count property
    total_count *int32
}
// NewCodespacesSecretsGetResponse instantiates a new CodespacesSecretsGetResponse and sets the default values.
func NewCodespacesSecretsGetResponse()(*CodespacesSecretsGetResponse) {
    m := &CodespacesSecretsGetResponse{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateCodespacesSecretsGetResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateCodespacesSecretsGetResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCodespacesSecretsGetResponse(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *CodespacesSecretsGetResponse) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *CodespacesSecretsGetResponse) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["secrets"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(i000736ae6dd74f01081193e4f903216bc2bd2954ed818433b986f45d581ed035.CreateCodespacesSecretFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]i000736ae6dd74f01081193e4f903216bc2bd2954ed818433b986f45d581ed035.CodespacesSecretable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(i000736ae6dd74f01081193e4f903216bc2bd2954ed818433b986f45d581ed035.CodespacesSecretable)
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
func (m *CodespacesSecretsGetResponse) GetSecrets()([]i000736ae6dd74f01081193e4f903216bc2bd2954ed818433b986f45d581ed035.CodespacesSecretable) {
    return m.secrets
}
// GetTotalCount gets the total_count property value. The total_count property
func (m *CodespacesSecretsGetResponse) GetTotalCount()(*int32) {
    return m.total_count
}
// Serialize serializes information the current object
func (m *CodespacesSecretsGetResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
func (m *CodespacesSecretsGetResponse) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetSecrets sets the secrets property value. The secrets property
func (m *CodespacesSecretsGetResponse) SetSecrets(value []i000736ae6dd74f01081193e4f903216bc2bd2954ed818433b986f45d581ed035.CodespacesSecretable)() {
    m.secrets = value
}
// SetTotalCount sets the total_count property value. The total_count property
func (m *CodespacesSecretsGetResponse) SetTotalCount(value *int32)() {
    m.total_count = value
}
// CodespacesSecretsGetResponseable 
type CodespacesSecretsGetResponseable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetSecrets()([]i000736ae6dd74f01081193e4f903216bc2bd2954ed818433b986f45d581ed035.CodespacesSecretable)
    GetTotalCount()(*int32)
    SetSecrets(value []i000736ae6dd74f01081193e4f903216bc2bd2954ed818433b986f45d581ed035.CodespacesSecretable)()
    SetTotalCount(value *int32)()
}
