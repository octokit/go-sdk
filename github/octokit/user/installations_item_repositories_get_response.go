package user

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    i000736ae6dd74f01081193e4f903216bc2bd2954ed818433b986f45d581ed035 "github.com/octokit/go-sdk/github/octokit/models"
)

// InstallationsItemRepositoriesGetResponse 
type InstallationsItemRepositoriesGetResponse struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The repositories property
    repositories []i000736ae6dd74f01081193e4f903216bc2bd2954ed818433b986f45d581ed035.Repositoryable
    // The repository_selection property
    repository_selection *string
    // The total_count property
    total_count *int32
}
// NewInstallationsItemRepositoriesGetResponse instantiates a new InstallationsItemRepositoriesGetResponse and sets the default values.
func NewInstallationsItemRepositoriesGetResponse()(*InstallationsItemRepositoriesGetResponse) {
    m := &InstallationsItemRepositoriesGetResponse{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateInstallationsItemRepositoriesGetResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateInstallationsItemRepositoriesGetResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewInstallationsItemRepositoriesGetResponse(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *InstallationsItemRepositoriesGetResponse) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *InstallationsItemRepositoriesGetResponse) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["repositories"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(i000736ae6dd74f01081193e4f903216bc2bd2954ed818433b986f45d581ed035.CreateRepositoryFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]i000736ae6dd74f01081193e4f903216bc2bd2954ed818433b986f45d581ed035.Repositoryable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(i000736ae6dd74f01081193e4f903216bc2bd2954ed818433b986f45d581ed035.Repositoryable)
                }
            }
            m.SetRepositories(res)
        }
        return nil
    }
    res["repository_selection"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRepositorySelection(val)
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
func (m *InstallationsItemRepositoriesGetResponse) GetRepositories()([]i000736ae6dd74f01081193e4f903216bc2bd2954ed818433b986f45d581ed035.Repositoryable) {
    return m.repositories
}
// GetRepositorySelection gets the repository_selection property value. The repository_selection property
func (m *InstallationsItemRepositoriesGetResponse) GetRepositorySelection()(*string) {
    return m.repository_selection
}
// GetTotalCount gets the total_count property value. The total_count property
func (m *InstallationsItemRepositoriesGetResponse) GetTotalCount()(*int32) {
    return m.total_count
}
// Serialize serializes information the current object
func (m *InstallationsItemRepositoriesGetResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
        err := writer.WriteStringValue("repository_selection", m.GetRepositorySelection())
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
func (m *InstallationsItemRepositoriesGetResponse) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetRepositories sets the repositories property value. The repositories property
func (m *InstallationsItemRepositoriesGetResponse) SetRepositories(value []i000736ae6dd74f01081193e4f903216bc2bd2954ed818433b986f45d581ed035.Repositoryable)() {
    m.repositories = value
}
// SetRepositorySelection sets the repository_selection property value. The repository_selection property
func (m *InstallationsItemRepositoriesGetResponse) SetRepositorySelection(value *string)() {
    m.repository_selection = value
}
// SetTotalCount sets the total_count property value. The total_count property
func (m *InstallationsItemRepositoriesGetResponse) SetTotalCount(value *int32)() {
    m.total_count = value
}
// InstallationsItemRepositoriesGetResponseable 
type InstallationsItemRepositoriesGetResponseable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetRepositories()([]i000736ae6dd74f01081193e4f903216bc2bd2954ed818433b986f45d581ed035.Repositoryable)
    GetRepositorySelection()(*string)
    GetTotalCount()(*int32)
    SetRepositories(value []i000736ae6dd74f01081193e4f903216bc2bd2954ed818433b986f45d581ed035.Repositoryable)()
    SetRepositorySelection(value *string)()
    SetTotalCount(value *int32)()
}
