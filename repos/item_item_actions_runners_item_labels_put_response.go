package repos

import (
    i33bcfa37677dfd8b71eaf22ad3bfd140ba1702e7a8418558499c22418bf2fef2 "github.com/octokit/go-sdk/models"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemItemActionsRunnersItemLabelsPutResponse 
type ItemItemActionsRunnersItemLabelsPutResponse struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The labels property
    labels []i33bcfa37677dfd8b71eaf22ad3bfd140ba1702e7a8418558499c22418bf2fef2.RunnerLabelable
    // The total_count property
    total_count *int32
}
// NewItemItemActionsRunnersItemLabelsPutResponse instantiates a new ItemItemActionsRunnersItemLabelsPutResponse and sets the default values.
func NewItemItemActionsRunnersItemLabelsPutResponse()(*ItemItemActionsRunnersItemLabelsPutResponse) {
    m := &ItemItemActionsRunnersItemLabelsPutResponse{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateItemItemActionsRunnersItemLabelsPutResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemItemActionsRunnersItemLabelsPutResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemItemActionsRunnersItemLabelsPutResponse(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ItemItemActionsRunnersItemLabelsPutResponse) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ItemItemActionsRunnersItemLabelsPutResponse) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["labels"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(i33bcfa37677dfd8b71eaf22ad3bfd140ba1702e7a8418558499c22418bf2fef2.CreateRunnerLabelFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]i33bcfa37677dfd8b71eaf22ad3bfd140ba1702e7a8418558499c22418bf2fef2.RunnerLabelable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(i33bcfa37677dfd8b71eaf22ad3bfd140ba1702e7a8418558499c22418bf2fef2.RunnerLabelable)
                }
            }
            m.SetLabels(res)
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
// GetLabels gets the labels property value. The labels property
func (m *ItemItemActionsRunnersItemLabelsPutResponse) GetLabels()([]i33bcfa37677dfd8b71eaf22ad3bfd140ba1702e7a8418558499c22418bf2fef2.RunnerLabelable) {
    return m.labels
}
// GetTotalCount gets the total_count property value. The total_count property
func (m *ItemItemActionsRunnersItemLabelsPutResponse) GetTotalCount()(*int32) {
    return m.total_count
}
// Serialize serializes information the current object
func (m *ItemItemActionsRunnersItemLabelsPutResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetLabels() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetLabels()))
        for i, v := range m.GetLabels() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("labels", cast)
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
func (m *ItemItemActionsRunnersItemLabelsPutResponse) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetLabels sets the labels property value. The labels property
func (m *ItemItemActionsRunnersItemLabelsPutResponse) SetLabels(value []i33bcfa37677dfd8b71eaf22ad3bfd140ba1702e7a8418558499c22418bf2fef2.RunnerLabelable)() {
    m.labels = value
}
// SetTotalCount sets the total_count property value. The total_count property
func (m *ItemItemActionsRunnersItemLabelsPutResponse) SetTotalCount(value *int32)() {
    m.total_count = value
}
// ItemItemActionsRunnersItemLabelsPutResponseable 
type ItemItemActionsRunnersItemLabelsPutResponseable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetLabels()([]i33bcfa37677dfd8b71eaf22ad3bfd140ba1702e7a8418558499c22418bf2fef2.RunnerLabelable)
    GetTotalCount()(*int32)
    SetLabels(value []i33bcfa37677dfd8b71eaf22ad3bfd140ba1702e7a8418558499c22418bf2fef2.RunnerLabelable)()
    SetTotalCount(value *int32)()
}
