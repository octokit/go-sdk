package orgs

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemActionsRunnersGenerateJitconfigPostRequestBody 
type ItemActionsRunnersGenerateJitconfigPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The names of the custom labels to add to the runner. **Minimum items**: 1. **Maximum items**: 100.
    labels []string
    // The name of the new runner.
    name *string
    // The ID of the runner group to register the runner to.
    runner_group_id *int32
    // The working directory to be used for job execution, relative to the runner install directory.
    work_folder *string
}
// NewItemActionsRunnersGenerateJitconfigPostRequestBody instantiates a new ItemActionsRunnersGenerateJitconfigPostRequestBody and sets the default values.
func NewItemActionsRunnersGenerateJitconfigPostRequestBody()(*ItemActionsRunnersGenerateJitconfigPostRequestBody) {
    m := &ItemActionsRunnersGenerateJitconfigPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    work_folderValue := "_work"
    m.SetWorkFolder(&work_folderValue)
    return m
}
// CreateItemActionsRunnersGenerateJitconfigPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemActionsRunnersGenerateJitconfigPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemActionsRunnersGenerateJitconfigPostRequestBody(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ItemActionsRunnersGenerateJitconfigPostRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ItemActionsRunnersGenerateJitconfigPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["labels"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
            m.SetLabels(res)
        }
        return nil
    }
    res["name"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetName(val)
        }
        return nil
    }
    res["runner_group_id"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRunnerGroupId(val)
        }
        return nil
    }
    res["work_folder"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWorkFolder(val)
        }
        return nil
    }
    return res
}
// GetLabels gets the labels property value. The names of the custom labels to add to the runner. **Minimum items**: 1. **Maximum items**: 100.
func (m *ItemActionsRunnersGenerateJitconfigPostRequestBody) GetLabels()([]string) {
    return m.labels
}
// GetName gets the name property value. The name of the new runner.
func (m *ItemActionsRunnersGenerateJitconfigPostRequestBody) GetName()(*string) {
    return m.name
}
// GetRunnerGroupId gets the runner_group_id property value. The ID of the runner group to register the runner to.
func (m *ItemActionsRunnersGenerateJitconfigPostRequestBody) GetRunnerGroupId()(*int32) {
    return m.runner_group_id
}
// GetWorkFolder gets the work_folder property value. The working directory to be used for job execution, relative to the runner install directory.
func (m *ItemActionsRunnersGenerateJitconfigPostRequestBody) GetWorkFolder()(*string) {
    return m.work_folder
}
// Serialize serializes information the current object
func (m *ItemActionsRunnersGenerateJitconfigPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetLabels() != nil {
        err := writer.WriteCollectionOfStringValues("labels", m.GetLabels())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("name", m.GetName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("runner_group_id", m.GetRunnerGroupId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("work_folder", m.GetWorkFolder())
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
func (m *ItemActionsRunnersGenerateJitconfigPostRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetLabels sets the labels property value. The names of the custom labels to add to the runner. **Minimum items**: 1. **Maximum items**: 100.
func (m *ItemActionsRunnersGenerateJitconfigPostRequestBody) SetLabels(value []string)() {
    m.labels = value
}
// SetName sets the name property value. The name of the new runner.
func (m *ItemActionsRunnersGenerateJitconfigPostRequestBody) SetName(value *string)() {
    m.name = value
}
// SetRunnerGroupId sets the runner_group_id property value. The ID of the runner group to register the runner to.
func (m *ItemActionsRunnersGenerateJitconfigPostRequestBody) SetRunnerGroupId(value *int32)() {
    m.runner_group_id = value
}
// SetWorkFolder sets the work_folder property value. The working directory to be used for job execution, relative to the runner install directory.
func (m *ItemActionsRunnersGenerateJitconfigPostRequestBody) SetWorkFolder(value *string)() {
    m.work_folder = value
}
// ItemActionsRunnersGenerateJitconfigPostRequestBodyable 
type ItemActionsRunnersGenerateJitconfigPostRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetLabels()([]string)
    GetName()(*string)
    GetRunnerGroupId()(*int32)
    GetWorkFolder()(*string)
    SetLabels(value []string)()
    SetName(value *string)()
    SetRunnerGroupId(value *int32)()
    SetWorkFolder(value *string)()
}
