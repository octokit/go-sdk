package repos

import (
    i33bcfa37677dfd8b71eaf22ad3bfd140ba1702e7a8418558499c22418bf2fef2 "github.com/octokit/go-sdk/models"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemItemActionsRunsItemArtifactsGetResponse 
type ItemItemActionsRunsItemArtifactsGetResponse struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The artifacts property
    artifacts []i33bcfa37677dfd8b71eaf22ad3bfd140ba1702e7a8418558499c22418bf2fef2.Artifactable
    // The total_count property
    total_count *int32
}
// NewItemItemActionsRunsItemArtifactsGetResponse instantiates a new ItemItemActionsRunsItemArtifactsGetResponse and sets the default values.
func NewItemItemActionsRunsItemArtifactsGetResponse()(*ItemItemActionsRunsItemArtifactsGetResponse) {
    m := &ItemItemActionsRunsItemArtifactsGetResponse{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateItemItemActionsRunsItemArtifactsGetResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemItemActionsRunsItemArtifactsGetResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemItemActionsRunsItemArtifactsGetResponse(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ItemItemActionsRunsItemArtifactsGetResponse) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetArtifacts gets the artifacts property value. The artifacts property
func (m *ItemItemActionsRunsItemArtifactsGetResponse) GetArtifacts()([]i33bcfa37677dfd8b71eaf22ad3bfd140ba1702e7a8418558499c22418bf2fef2.Artifactable) {
    return m.artifacts
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ItemItemActionsRunsItemArtifactsGetResponse) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["artifacts"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(i33bcfa37677dfd8b71eaf22ad3bfd140ba1702e7a8418558499c22418bf2fef2.CreateArtifactFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]i33bcfa37677dfd8b71eaf22ad3bfd140ba1702e7a8418558499c22418bf2fef2.Artifactable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(i33bcfa37677dfd8b71eaf22ad3bfd140ba1702e7a8418558499c22418bf2fef2.Artifactable)
                }
            }
            m.SetArtifacts(res)
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
func (m *ItemItemActionsRunsItemArtifactsGetResponse) GetTotalCount()(*int32) {
    return m.total_count
}
// Serialize serializes information the current object
func (m *ItemItemActionsRunsItemArtifactsGetResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetArtifacts() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetArtifacts()))
        for i, v := range m.GetArtifacts() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("artifacts", cast)
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
func (m *ItemItemActionsRunsItemArtifactsGetResponse) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetArtifacts sets the artifacts property value. The artifacts property
func (m *ItemItemActionsRunsItemArtifactsGetResponse) SetArtifacts(value []i33bcfa37677dfd8b71eaf22ad3bfd140ba1702e7a8418558499c22418bf2fef2.Artifactable)() {
    m.artifacts = value
}
// SetTotalCount sets the total_count property value. The total_count property
func (m *ItemItemActionsRunsItemArtifactsGetResponse) SetTotalCount(value *int32)() {
    m.total_count = value
}
// ItemItemActionsRunsItemArtifactsGetResponseable 
type ItemItemActionsRunsItemArtifactsGetResponseable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetArtifacts()([]i33bcfa37677dfd8b71eaf22ad3bfd140ba1702e7a8418558499c22418bf2fef2.Artifactable)
    GetTotalCount()(*int32)
    SetArtifacts(value []i33bcfa37677dfd8b71eaf22ad3bfd140ba1702e7a8418558499c22418bf2fef2.Artifactable)()
    SetTotalCount(value *int32)()
}
