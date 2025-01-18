package orgs

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type ItemActionsHostedRunnersItemWithHosted_runner_PatchRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Whether this runner should be updated with a static public IP. Note limit on account. To list limits on account, use `GET actions/hosted-runners/limits`
    enable_static_ip *bool
    // The version of the runner image to deploy. This is relevant only for runners using custom images.
    image_version *string
    // The maximum amount of runners to scale up to. Runners will not auto-scale above this number. Use this setting to limit your cost.
    maximum_runners *int32
    // Name of the runner. Must be between 1 and 64 characters and may only contain upper and lowercase letters a-z, numbers 0-9, '.', '-', and '_'.
    name *string
    // The existing runner group to add this runner to.
    runner_group_id *int32
}
// NewItemActionsHostedRunnersItemWithHosted_runner_PatchRequestBody instantiates a new ItemActionsHostedRunnersItemWithHosted_runner_PatchRequestBody and sets the default values.
func NewItemActionsHostedRunnersItemWithHosted_runner_PatchRequestBody()(*ItemActionsHostedRunnersItemWithHosted_runner_PatchRequestBody) {
    m := &ItemActionsHostedRunnersItemWithHosted_runner_PatchRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateItemActionsHostedRunnersItemWithHosted_runner_PatchRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateItemActionsHostedRunnersItemWithHosted_runner_PatchRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemActionsHostedRunnersItemWithHosted_runner_PatchRequestBody(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *ItemActionsHostedRunnersItemWithHosted_runner_PatchRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetEnableStaticIp gets the enable_static_ip property value. Whether this runner should be updated with a static public IP. Note limit on account. To list limits on account, use `GET actions/hosted-runners/limits`
// returns a *bool when successful
func (m *ItemActionsHostedRunnersItemWithHosted_runner_PatchRequestBody) GetEnableStaticIp()(*bool) {
    return m.enable_static_ip
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ItemActionsHostedRunnersItemWithHosted_runner_PatchRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["enable_static_ip"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEnableStaticIp(val)
        }
        return nil
    }
    res["image_version"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetImageVersion(val)
        }
        return nil
    }
    res["maximum_runners"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMaximumRunners(val)
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
    return res
}
// GetImageVersion gets the image_version property value. The version of the runner image to deploy. This is relevant only for runners using custom images.
// returns a *string when successful
func (m *ItemActionsHostedRunnersItemWithHosted_runner_PatchRequestBody) GetImageVersion()(*string) {
    return m.image_version
}
// GetMaximumRunners gets the maximum_runners property value. The maximum amount of runners to scale up to. Runners will not auto-scale above this number. Use this setting to limit your cost.
// returns a *int32 when successful
func (m *ItemActionsHostedRunnersItemWithHosted_runner_PatchRequestBody) GetMaximumRunners()(*int32) {
    return m.maximum_runners
}
// GetName gets the name property value. Name of the runner. Must be between 1 and 64 characters and may only contain upper and lowercase letters a-z, numbers 0-9, '.', '-', and '_'.
// returns a *string when successful
func (m *ItemActionsHostedRunnersItemWithHosted_runner_PatchRequestBody) GetName()(*string) {
    return m.name
}
// GetRunnerGroupId gets the runner_group_id property value. The existing runner group to add this runner to.
// returns a *int32 when successful
func (m *ItemActionsHostedRunnersItemWithHosted_runner_PatchRequestBody) GetRunnerGroupId()(*int32) {
    return m.runner_group_id
}
// Serialize serializes information the current object
func (m *ItemActionsHostedRunnersItemWithHosted_runner_PatchRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteBoolValue("enable_static_ip", m.GetEnableStaticIp())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("image_version", m.GetImageVersion())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("maximum_runners", m.GetMaximumRunners())
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
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ItemActionsHostedRunnersItemWithHosted_runner_PatchRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetEnableStaticIp sets the enable_static_ip property value. Whether this runner should be updated with a static public IP. Note limit on account. To list limits on account, use `GET actions/hosted-runners/limits`
func (m *ItemActionsHostedRunnersItemWithHosted_runner_PatchRequestBody) SetEnableStaticIp(value *bool)() {
    m.enable_static_ip = value
}
// SetImageVersion sets the image_version property value. The version of the runner image to deploy. This is relevant only for runners using custom images.
func (m *ItemActionsHostedRunnersItemWithHosted_runner_PatchRequestBody) SetImageVersion(value *string)() {
    m.image_version = value
}
// SetMaximumRunners sets the maximum_runners property value. The maximum amount of runners to scale up to. Runners will not auto-scale above this number. Use this setting to limit your cost.
func (m *ItemActionsHostedRunnersItemWithHosted_runner_PatchRequestBody) SetMaximumRunners(value *int32)() {
    m.maximum_runners = value
}
// SetName sets the name property value. Name of the runner. Must be between 1 and 64 characters and may only contain upper and lowercase letters a-z, numbers 0-9, '.', '-', and '_'.
func (m *ItemActionsHostedRunnersItemWithHosted_runner_PatchRequestBody) SetName(value *string)() {
    m.name = value
}
// SetRunnerGroupId sets the runner_group_id property value. The existing runner group to add this runner to.
func (m *ItemActionsHostedRunnersItemWithHosted_runner_PatchRequestBody) SetRunnerGroupId(value *int32)() {
    m.runner_group_id = value
}
type ItemActionsHostedRunnersItemWithHosted_runner_PatchRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetEnableStaticIp()(*bool)
    GetImageVersion()(*string)
    GetMaximumRunners()(*int32)
    GetName()(*string)
    GetRunnerGroupId()(*int32)
    SetEnableStaticIp(value *bool)()
    SetImageVersion(value *string)()
    SetMaximumRunners(value *int32)()
    SetName(value *string)()
    SetRunnerGroupId(value *int32)()
}
