package orgs

import (
    i33bcfa37677dfd8b71eaf22ad3bfd140ba1702e7a8418558499c22418bf2fef2 "github.com/octokit/go-sdk/models"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemCodespacesSecretsItemRepositoriesGetResponse 
type ItemCodespacesSecretsItemRepositoriesGetResponse struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The repositories property
    repositories []i33bcfa37677dfd8b71eaf22ad3bfd140ba1702e7a8418558499c22418bf2fef2.MinimalRepositoryable
    // The total_count property
    total_count *int32
}
// NewItemCodespacesSecretsItemRepositoriesGetResponse instantiates a new ItemCodespacesSecretsItemRepositoriesGetResponse and sets the default values.
func NewItemCodespacesSecretsItemRepositoriesGetResponse()(*ItemCodespacesSecretsItemRepositoriesGetResponse) {
    m := &ItemCodespacesSecretsItemRepositoriesGetResponse{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateItemCodespacesSecretsItemRepositoriesGetResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemCodespacesSecretsItemRepositoriesGetResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemCodespacesSecretsItemRepositoriesGetResponse(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ItemCodespacesSecretsItemRepositoriesGetResponse) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ItemCodespacesSecretsItemRepositoriesGetResponse) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["repositories"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(i33bcfa37677dfd8b71eaf22ad3bfd140ba1702e7a8418558499c22418bf2fef2.CreateMinimalRepositoryFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]i33bcfa37677dfd8b71eaf22ad3bfd140ba1702e7a8418558499c22418bf2fef2.MinimalRepositoryable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(i33bcfa37677dfd8b71eaf22ad3bfd140ba1702e7a8418558499c22418bf2fef2.MinimalRepositoryable)
                }
            }
            m.SetRepositories(res)
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
// GetRepositories gets the repositories property value. The repositories property
func (m *ItemCodespacesSecretsItemRepositoriesGetResponse) GetRepositories()([]i33bcfa37677dfd8b71eaf22ad3bfd140ba1702e7a8418558499c22418bf2fef2.MinimalRepositoryable) {
    return m.repositories
}
// GetTotalCount gets the total_count property value. The total_count property
func (m *ItemCodespacesSecretsItemRepositoriesGetResponse) GetTotalCount()(*int32) {
    return m.total_count
}
// Serialize serializes information the current object
func (m *ItemCodespacesSecretsItemRepositoriesGetResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetRepositories() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetRepositories()))
        for i, v := range m.GetRepositories() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("repositories", cast)
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
func (m *ItemCodespacesSecretsItemRepositoriesGetResponse) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetRepositories sets the repositories property value. The repositories property
func (m *ItemCodespacesSecretsItemRepositoriesGetResponse) SetRepositories(value []i33bcfa37677dfd8b71eaf22ad3bfd140ba1702e7a8418558499c22418bf2fef2.MinimalRepositoryable)() {
    m.repositories = value
}
// SetTotalCount sets the total_count property value. The total_count property
func (m *ItemCodespacesSecretsItemRepositoriesGetResponse) SetTotalCount(value *int32)() {
    m.total_count = value
}
// ItemCodespacesSecretsItemRepositoriesGetResponseable 
type ItemCodespacesSecretsItemRepositoriesGetResponseable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetRepositories()([]i33bcfa37677dfd8b71eaf22ad3bfd140ba1702e7a8418558499c22418bf2fef2.MinimalRepositoryable)
    GetTotalCount()(*int32)
    SetRepositories(value []i33bcfa37677dfd8b71eaf22ad3bfd140ba1702e7a8418558499c22418bf2fef2.MinimalRepositoryable)()
    SetTotalCount(value *int32)()
}
