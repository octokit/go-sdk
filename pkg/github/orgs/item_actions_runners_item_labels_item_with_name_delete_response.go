package orgs

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6 "github.com/octokit/go-sdk/pkg/github/models"
)

// ItemActionsRunnersItemLabelsItemWithNameDeleteResponse 
type ItemActionsRunnersItemLabelsItemWithNameDeleteResponse struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The labels property
    labels []i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.RunnerLabelable
    // The total_count property
    total_count *int32
}
// NewItemActionsRunnersItemLabelsItemWithNameDeleteResponse instantiates a new ItemActionsRunnersItemLabelsItemWithNameDeleteResponse and sets the default values.
func NewItemActionsRunnersItemLabelsItemWithNameDeleteResponse()(*ItemActionsRunnersItemLabelsItemWithNameDeleteResponse) {
    m := &ItemActionsRunnersItemLabelsItemWithNameDeleteResponse{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateItemActionsRunnersItemLabelsItemWithNameDeleteResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemActionsRunnersItemLabelsItemWithNameDeleteResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemActionsRunnersItemLabelsItemWithNameDeleteResponse(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ItemActionsRunnersItemLabelsItemWithNameDeleteResponse) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ItemActionsRunnersItemLabelsItemWithNameDeleteResponse) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["labels"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CreateRunnerLabelFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.RunnerLabelable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.RunnerLabelable)
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
func (m *ItemActionsRunnersItemLabelsItemWithNameDeleteResponse) GetLabels()([]i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.RunnerLabelable) {
    return m.labels
}
// GetTotalCount gets the total_count property value. The total_count property
func (m *ItemActionsRunnersItemLabelsItemWithNameDeleteResponse) GetTotalCount()(*int32) {
    return m.total_count
}
// Serialize serializes information the current object
func (m *ItemActionsRunnersItemLabelsItemWithNameDeleteResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
func (m *ItemActionsRunnersItemLabelsItemWithNameDeleteResponse) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetLabels sets the labels property value. The labels property
func (m *ItemActionsRunnersItemLabelsItemWithNameDeleteResponse) SetLabels(value []i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.RunnerLabelable)() {
    m.labels = value
}
// SetTotalCount sets the total_count property value. The total_count property
func (m *ItemActionsRunnersItemLabelsItemWithNameDeleteResponse) SetTotalCount(value *int32)() {
    m.total_count = value
}
// ItemActionsRunnersItemLabelsItemWithNameDeleteResponseable 
type ItemActionsRunnersItemLabelsItemWithNameDeleteResponseable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetLabels()([]i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.RunnerLabelable)
    GetTotalCount()(*int32)
    SetLabels(value []i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.RunnerLabelable)()
    SetTotalCount(value *int32)()
}
