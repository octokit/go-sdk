package app

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    i000736ae6dd74f01081193e4f903216bc2bd2954ed818433b986f45d581ed035 "github.com/octokit/go-sdk/github/octokit/models"
)

// InstallationsItemAccess_tokensPostRequestBody 
type InstallationsItemAccess_tokensPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The permissions granted to the user access token.
    permissions i000736ae6dd74f01081193e4f903216bc2bd2954ed818433b986f45d581ed035.AppPermissionsable
    // List of repository names that the token should have access to
    repositories []string
    // List of repository IDs that the token should have access to
    repository_ids []int32
}
// NewInstallationsItemAccess_tokensPostRequestBody instantiates a new InstallationsItemAccess_tokensPostRequestBody and sets the default values.
func NewInstallationsItemAccess_tokensPostRequestBody()(*InstallationsItemAccess_tokensPostRequestBody) {
    m := &InstallationsItemAccess_tokensPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateInstallationsItemAccess_tokensPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateInstallationsItemAccess_tokensPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewInstallationsItemAccess_tokensPostRequestBody(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *InstallationsItemAccess_tokensPostRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *InstallationsItemAccess_tokensPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["permissions"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(i000736ae6dd74f01081193e4f903216bc2bd2954ed818433b986f45d581ed035.CreateAppPermissionsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPermissions(val.(i000736ae6dd74f01081193e4f903216bc2bd2954ed818433b986f45d581ed035.AppPermissionsable))
        }
        return nil
    }
    res["repositories"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
            m.SetRepositories(res)
        }
        return nil
    }
    res["repository_ids"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("int32")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]int32, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = *(v.(*int32))
                }
            }
            m.SetRepositoryIds(res)
        }
        return nil
    }
    return res
}
// GetPermissions gets the permissions property value. The permissions granted to the user access token.
func (m *InstallationsItemAccess_tokensPostRequestBody) GetPermissions()(i000736ae6dd74f01081193e4f903216bc2bd2954ed818433b986f45d581ed035.AppPermissionsable) {
    return m.permissions
}
// GetRepositories gets the repositories property value. List of repository names that the token should have access to
func (m *InstallationsItemAccess_tokensPostRequestBody) GetRepositories()([]string) {
    return m.repositories
}
// GetRepositoryIds gets the repository_ids property value. List of repository IDs that the token should have access to
func (m *InstallationsItemAccess_tokensPostRequestBody) GetRepositoryIds()([]int32) {
    return m.repository_ids
}
// Serialize serializes information the current object
func (m *InstallationsItemAccess_tokensPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("permissions", m.GetPermissions())
        if err != nil {
            return err
        }
    }
    if m.GetRepositories() != nil {
        err := writer.WriteCollectionOfStringValues("repositories", m.GetRepositories())
        if err != nil {
            return err
        }
    }
    if m.GetRepositoryIds() != nil {
        err := writer.WriteCollectionOfInt32Values("repository_ids", m.GetRepositoryIds())
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
func (m *InstallationsItemAccess_tokensPostRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetPermissions sets the permissions property value. The permissions granted to the user access token.
func (m *InstallationsItemAccess_tokensPostRequestBody) SetPermissions(value i000736ae6dd74f01081193e4f903216bc2bd2954ed818433b986f45d581ed035.AppPermissionsable)() {
    m.permissions = value
}
// SetRepositories sets the repositories property value. List of repository names that the token should have access to
func (m *InstallationsItemAccess_tokensPostRequestBody) SetRepositories(value []string)() {
    m.repositories = value
}
// SetRepositoryIds sets the repository_ids property value. List of repository IDs that the token should have access to
func (m *InstallationsItemAccess_tokensPostRequestBody) SetRepositoryIds(value []int32)() {
    m.repository_ids = value
}
// InstallationsItemAccess_tokensPostRequestBodyable 
type InstallationsItemAccess_tokensPostRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetPermissions()(i000736ae6dd74f01081193e4f903216bc2bd2954ed818433b986f45d581ed035.AppPermissionsable)
    GetRepositories()([]string)
    GetRepositoryIds()([]int32)
    SetPermissions(value i000736ae6dd74f01081193e4f903216bc2bd2954ed818433b986f45d581ed035.AppPermissionsable)()
    SetRepositories(value []string)()
    SetRepositoryIds(value []int32)()
}
