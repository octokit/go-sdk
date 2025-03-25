package orgs

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type ItemActionsHostedRunnersPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Whether this runner should be created with a static public IP. Note limit on account. To list limits on account, use `GET actions/hosted-runners/limits`
    enable_static_ip *bool
    // The image of runner. To list all available images, use `GET /actions/hosted-runners/images/github-owned` or `GET /actions/hosted-runners/images/partner`.
    image ItemActionsHostedRunnersPostRequestBody_imageable
    // The maximum amount of runners to scale up to. Runners will not auto-scale above this number. Use this setting to limit your cost.
    maximum_runners *int32
    // Name of the runner. Must be between 1 and 64 characters and may only contain upper and lowercase letters a-z, numbers 0-9, '.', '-', and '_'.
    name *string
    // The existing runner group to add this runner to.
    runner_group_id *int32
    // The machine size of the runner. To list available sizes, use `GET actions/hosted-runners/machine-sizes`
    size *string
}
// NewItemActionsHostedRunnersPostRequestBody instantiates a new ItemActionsHostedRunnersPostRequestBody and sets the default values.
func NewItemActionsHostedRunnersPostRequestBody()(*ItemActionsHostedRunnersPostRequestBody) {
    m := &ItemActionsHostedRunnersPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateItemActionsHostedRunnersPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateItemActionsHostedRunnersPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemActionsHostedRunnersPostRequestBody(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *ItemActionsHostedRunnersPostRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetEnableStaticIp gets the enable_static_ip property value. Whether this runner should be created with a static public IP. Note limit on account. To list limits on account, use `GET actions/hosted-runners/limits`
// returns a *bool when successful
func (m *ItemActionsHostedRunnersPostRequestBody) GetEnableStaticIp()(*bool) {
    return m.enable_static_ip
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ItemActionsHostedRunnersPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
    res["image"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateItemActionsHostedRunnersPostRequestBody_imageFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetImage(val.(ItemActionsHostedRunnersPostRequestBody_imageable))
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
    res["size"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSize(val)
        }
        return nil
    }
    return res
}
// GetImage gets the image property value. The image of runner. To list all available images, use `GET /actions/hosted-runners/images/github-owned` or `GET /actions/hosted-runners/images/partner`.
// returns a ItemActionsHostedRunnersPostRequestBody_imageable when successful
func (m *ItemActionsHostedRunnersPostRequestBody) GetImage()(ItemActionsHostedRunnersPostRequestBody_imageable) {
    return m.image
}
// GetMaximumRunners gets the maximum_runners property value. The maximum amount of runners to scale up to. Runners will not auto-scale above this number. Use this setting to limit your cost.
// returns a *int32 when successful
func (m *ItemActionsHostedRunnersPostRequestBody) GetMaximumRunners()(*int32) {
    return m.maximum_runners
}
// GetName gets the name property value. Name of the runner. Must be between 1 and 64 characters and may only contain upper and lowercase letters a-z, numbers 0-9, '.', '-', and '_'.
// returns a *string when successful
func (m *ItemActionsHostedRunnersPostRequestBody) GetName()(*string) {
    return m.name
}
// GetRunnerGroupId gets the runner_group_id property value. The existing runner group to add this runner to.
// returns a *int32 when successful
func (m *ItemActionsHostedRunnersPostRequestBody) GetRunnerGroupId()(*int32) {
    return m.runner_group_id
}
// GetSize gets the size property value. The machine size of the runner. To list available sizes, use `GET actions/hosted-runners/machine-sizes`
// returns a *string when successful
func (m *ItemActionsHostedRunnersPostRequestBody) GetSize()(*string) {
    return m.size
}
// Serialize serializes information the current object
func (m *ItemActionsHostedRunnersPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteBoolValue("enable_static_ip", m.GetEnableStaticIp())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("image", m.GetImage())
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
        err := writer.WriteStringValue("size", m.GetSize())
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
func (m *ItemActionsHostedRunnersPostRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetEnableStaticIp sets the enable_static_ip property value. Whether this runner should be created with a static public IP. Note limit on account. To list limits on account, use `GET actions/hosted-runners/limits`
func (m *ItemActionsHostedRunnersPostRequestBody) SetEnableStaticIp(value *bool)() {
    m.enable_static_ip = value
}
// SetImage sets the image property value. The image of runner. To list all available images, use `GET /actions/hosted-runners/images/github-owned` or `GET /actions/hosted-runners/images/partner`.
func (m *ItemActionsHostedRunnersPostRequestBody) SetImage(value ItemActionsHostedRunnersPostRequestBody_imageable)() {
    m.image = value
}
// SetMaximumRunners sets the maximum_runners property value. The maximum amount of runners to scale up to. Runners will not auto-scale above this number. Use this setting to limit your cost.
func (m *ItemActionsHostedRunnersPostRequestBody) SetMaximumRunners(value *int32)() {
    m.maximum_runners = value
}
// SetName sets the name property value. Name of the runner. Must be between 1 and 64 characters and may only contain upper and lowercase letters a-z, numbers 0-9, '.', '-', and '_'.
func (m *ItemActionsHostedRunnersPostRequestBody) SetName(value *string)() {
    m.name = value
}
// SetRunnerGroupId sets the runner_group_id property value. The existing runner group to add this runner to.
func (m *ItemActionsHostedRunnersPostRequestBody) SetRunnerGroupId(value *int32)() {
    m.runner_group_id = value
}
// SetSize sets the size property value. The machine size of the runner. To list available sizes, use `GET actions/hosted-runners/machine-sizes`
func (m *ItemActionsHostedRunnersPostRequestBody) SetSize(value *string)() {
    m.size = value
}
type ItemActionsHostedRunnersPostRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetEnableStaticIp()(*bool)
    GetImage()(ItemActionsHostedRunnersPostRequestBody_imageable)
    GetMaximumRunners()(*int32)
    GetName()(*string)
    GetRunnerGroupId()(*int32)
    GetSize()(*string)
    SetEnableStaticIp(value *bool)()
    SetImage(value ItemActionsHostedRunnersPostRequestBody_imageable)()
    SetMaximumRunners(value *int32)()
    SetName(value *string)()
    SetRunnerGroupId(value *int32)()
    SetSize(value *string)()
}
