package orgs

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    i000736ae6dd74f01081193e4f903216bc2bd2954ed818433b986f45d581ed035 "github.com/octokit/go-sdk/github/octokit/models"
)

// ItemActionsCacheUsageByRepositoryGetResponse 
type ItemActionsCacheUsageByRepositoryGetResponse struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The repository_cache_usages property
    repository_cache_usages []i000736ae6dd74f01081193e4f903216bc2bd2954ed818433b986f45d581ed035.ActionsCacheUsageByRepositoryable
    // The total_count property
    total_count *int32
}
// NewItemActionsCacheUsageByRepositoryGetResponse instantiates a new ItemActionsCacheUsageByRepositoryGetResponse and sets the default values.
func NewItemActionsCacheUsageByRepositoryGetResponse()(*ItemActionsCacheUsageByRepositoryGetResponse) {
    m := &ItemActionsCacheUsageByRepositoryGetResponse{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateItemActionsCacheUsageByRepositoryGetResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemActionsCacheUsageByRepositoryGetResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemActionsCacheUsageByRepositoryGetResponse(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ItemActionsCacheUsageByRepositoryGetResponse) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ItemActionsCacheUsageByRepositoryGetResponse) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["repository_cache_usages"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(i000736ae6dd74f01081193e4f903216bc2bd2954ed818433b986f45d581ed035.CreateActionsCacheUsageByRepositoryFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]i000736ae6dd74f01081193e4f903216bc2bd2954ed818433b986f45d581ed035.ActionsCacheUsageByRepositoryable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(i000736ae6dd74f01081193e4f903216bc2bd2954ed818433b986f45d581ed035.ActionsCacheUsageByRepositoryable)
                }
            }
            m.SetRepositoryCacheUsages(res)
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
// GetRepositoryCacheUsages gets the repository_cache_usages property value. The repository_cache_usages property
func (m *ItemActionsCacheUsageByRepositoryGetResponse) GetRepositoryCacheUsages()([]i000736ae6dd74f01081193e4f903216bc2bd2954ed818433b986f45d581ed035.ActionsCacheUsageByRepositoryable) {
    return m.repository_cache_usages
}
// GetTotalCount gets the total_count property value. The total_count property
func (m *ItemActionsCacheUsageByRepositoryGetResponse) GetTotalCount()(*int32) {
    return m.total_count
}
// Serialize serializes information the current object
func (m *ItemActionsCacheUsageByRepositoryGetResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetRepositoryCacheUsages() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetRepositoryCacheUsages()))
        for i, v := range m.GetRepositoryCacheUsages() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("repository_cache_usages", cast)
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
func (m *ItemActionsCacheUsageByRepositoryGetResponse) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetRepositoryCacheUsages sets the repository_cache_usages property value. The repository_cache_usages property
func (m *ItemActionsCacheUsageByRepositoryGetResponse) SetRepositoryCacheUsages(value []i000736ae6dd74f01081193e4f903216bc2bd2954ed818433b986f45d581ed035.ActionsCacheUsageByRepositoryable)() {
    m.repository_cache_usages = value
}
// SetTotalCount sets the total_count property value. The total_count property
func (m *ItemActionsCacheUsageByRepositoryGetResponse) SetTotalCount(value *int32)() {
    m.total_count = value
}
// ItemActionsCacheUsageByRepositoryGetResponseable 
type ItemActionsCacheUsageByRepositoryGetResponseable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetRepositoryCacheUsages()([]i000736ae6dd74f01081193e4f903216bc2bd2954ed818433b986f45d581ed035.ActionsCacheUsageByRepositoryable)
    GetTotalCount()(*int32)
    SetRepositoryCacheUsages(value []i000736ae6dd74f01081193e4f903216bc2bd2954ed818433b986f45d581ed035.ActionsCacheUsageByRepositoryable)()
    SetTotalCount(value *int32)()
}
