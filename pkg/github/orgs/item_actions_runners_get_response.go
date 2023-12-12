package orgs

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6 "github.com/octokit/go-sdk/pkg/github/models"
)

// ItemActionsRunnersGetResponse 
type ItemActionsRunnersGetResponse struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The runners property
    runners []i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.Runnerable
    // The total_count property
    total_count *int32
}
// NewItemActionsRunnersGetResponse instantiates a new ItemActionsRunnersGetResponse and sets the default values.
func NewItemActionsRunnersGetResponse()(*ItemActionsRunnersGetResponse) {
    m := &ItemActionsRunnersGetResponse{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateItemActionsRunnersGetResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemActionsRunnersGetResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemActionsRunnersGetResponse(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ItemActionsRunnersGetResponse) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ItemActionsRunnersGetResponse) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["runners"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CreateRunnerFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.Runnerable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.Runnerable)
                }
            }
            m.SetRunners(res)
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
// GetRunners gets the runners property value. The runners property
func (m *ItemActionsRunnersGetResponse) GetRunners()([]i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.Runnerable) {
    return m.runners
}
// GetTotalCount gets the total_count property value. The total_count property
func (m *ItemActionsRunnersGetResponse) GetTotalCount()(*int32) {
    return m.total_count
}
// Serialize serializes information the current object
func (m *ItemActionsRunnersGetResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetRunners() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetRunners()))
        for i, v := range m.GetRunners() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("runners", cast)
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
func (m *ItemActionsRunnersGetResponse) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetRunners sets the runners property value. The runners property
func (m *ItemActionsRunnersGetResponse) SetRunners(value []i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.Runnerable)() {
    m.runners = value
}
// SetTotalCount sets the total_count property value. The total_count property
func (m *ItemActionsRunnersGetResponse) SetTotalCount(value *int32)() {
    m.total_count = value
}
// ItemActionsRunnersGetResponseable 
type ItemActionsRunnersGetResponseable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetRunners()([]i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.Runnerable)
    GetTotalCount()(*int32)
    SetRunners(value []i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.Runnerable)()
    SetTotalCount(value *int32)()
}
