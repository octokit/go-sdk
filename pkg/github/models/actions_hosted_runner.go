package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ActionsHostedRunner a Github-hosted hosted runner.
type ActionsHostedRunner struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The unique identifier of the hosted runner.
    id *int32
    // Provides details of a hosted runner image
    image_details NullableActionsHostedRunnerPoolImageable
    // The time at which the runner was last used, in ISO 8601 format.
    last_active_on *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Provides details of a particular machine spec.
    machine_size_details ActionsHostedRunnerMachineSpecable
    // The maximum amount of hosted runners. Runners will not scale automatically above this number. Use this setting to limit your cost.
    maximum_runners *int32
    // The name of the hosted runner.
    name *string
    // The operating system of the image.
    platform *string
    // Whether public IP is enabled for the hosted runners.
    public_ip_enabled *bool
    // The public IP ranges when public IP is enabled for the hosted runners.
    public_ips []PublicIpable
    // The unique identifier of the group that the hosted runner belongs to.
    runner_group_id *int32
    // The status of the runner.
    status *ActionsHostedRunner_status
}
// NewActionsHostedRunner instantiates a new ActionsHostedRunner and sets the default values.
func NewActionsHostedRunner()(*ActionsHostedRunner) {
    m := &ActionsHostedRunner{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateActionsHostedRunnerFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateActionsHostedRunnerFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewActionsHostedRunner(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *ActionsHostedRunner) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ActionsHostedRunner) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["id"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetId(val)
        }
        return nil
    }
    res["image_details"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateNullableActionsHostedRunnerPoolImageFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetImageDetails(val.(NullableActionsHostedRunnerPoolImageable))
        }
        return nil
    }
    res["last_active_on"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastActiveOn(val)
        }
        return nil
    }
    res["machine_size_details"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateActionsHostedRunnerMachineSpecFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMachineSizeDetails(val.(ActionsHostedRunnerMachineSpecable))
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
    res["platform"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPlatform(val)
        }
        return nil
    }
    res["public_ip_enabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPublicIpEnabled(val)
        }
        return nil
    }
    res["public_ips"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreatePublicIpFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]PublicIpable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(PublicIpable)
                }
            }
            m.SetPublicIps(res)
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
    res["status"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseActionsHostedRunner_status)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStatus(val.(*ActionsHostedRunner_status))
        }
        return nil
    }
    return res
}
// GetId gets the id property value. The unique identifier of the hosted runner.
// returns a *int32 when successful
func (m *ActionsHostedRunner) GetId()(*int32) {
    return m.id
}
// GetImageDetails gets the image_details property value. Provides details of a hosted runner image
// returns a NullableActionsHostedRunnerPoolImageable when successful
func (m *ActionsHostedRunner) GetImageDetails()(NullableActionsHostedRunnerPoolImageable) {
    return m.image_details
}
// GetLastActiveOn gets the last_active_on property value. The time at which the runner was last used, in ISO 8601 format.
// returns a *Time when successful
func (m *ActionsHostedRunner) GetLastActiveOn()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.last_active_on
}
// GetMachineSizeDetails gets the machine_size_details property value. Provides details of a particular machine spec.
// returns a ActionsHostedRunnerMachineSpecable when successful
func (m *ActionsHostedRunner) GetMachineSizeDetails()(ActionsHostedRunnerMachineSpecable) {
    return m.machine_size_details
}
// GetMaximumRunners gets the maximum_runners property value. The maximum amount of hosted runners. Runners will not scale automatically above this number. Use this setting to limit your cost.
// returns a *int32 when successful
func (m *ActionsHostedRunner) GetMaximumRunners()(*int32) {
    return m.maximum_runners
}
// GetName gets the name property value. The name of the hosted runner.
// returns a *string when successful
func (m *ActionsHostedRunner) GetName()(*string) {
    return m.name
}
// GetPlatform gets the platform property value. The operating system of the image.
// returns a *string when successful
func (m *ActionsHostedRunner) GetPlatform()(*string) {
    return m.platform
}
// GetPublicIpEnabled gets the public_ip_enabled property value. Whether public IP is enabled for the hosted runners.
// returns a *bool when successful
func (m *ActionsHostedRunner) GetPublicIpEnabled()(*bool) {
    return m.public_ip_enabled
}
// GetPublicIps gets the public_ips property value. The public IP ranges when public IP is enabled for the hosted runners.
// returns a []PublicIpable when successful
func (m *ActionsHostedRunner) GetPublicIps()([]PublicIpable) {
    return m.public_ips
}
// GetRunnerGroupId gets the runner_group_id property value. The unique identifier of the group that the hosted runner belongs to.
// returns a *int32 when successful
func (m *ActionsHostedRunner) GetRunnerGroupId()(*int32) {
    return m.runner_group_id
}
// GetStatus gets the status property value. The status of the runner.
// returns a *ActionsHostedRunner_status when successful
func (m *ActionsHostedRunner) GetStatus()(*ActionsHostedRunner_status) {
    return m.status
}
// Serialize serializes information the current object
func (m *ActionsHostedRunner) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteInt32Value("id", m.GetId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("image_details", m.GetImageDetails())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteTimeValue("last_active_on", m.GetLastActiveOn())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("machine_size_details", m.GetMachineSizeDetails())
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
        err := writer.WriteStringValue("platform", m.GetPlatform())
        if err != nil {
            return err
        }
    }
    if m.GetPublicIps() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetPublicIps()))
        for i, v := range m.GetPublicIps() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("public_ips", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("public_ip_enabled", m.GetPublicIpEnabled())
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
    if m.GetStatus() != nil {
        cast := (*m.GetStatus()).String()
        err := writer.WriteStringValue("status", &cast)
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
func (m *ActionsHostedRunner) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetId sets the id property value. The unique identifier of the hosted runner.
func (m *ActionsHostedRunner) SetId(value *int32)() {
    m.id = value
}
// SetImageDetails sets the image_details property value. Provides details of a hosted runner image
func (m *ActionsHostedRunner) SetImageDetails(value NullableActionsHostedRunnerPoolImageable)() {
    m.image_details = value
}
// SetLastActiveOn sets the last_active_on property value. The time at which the runner was last used, in ISO 8601 format.
func (m *ActionsHostedRunner) SetLastActiveOn(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.last_active_on = value
}
// SetMachineSizeDetails sets the machine_size_details property value. Provides details of a particular machine spec.
func (m *ActionsHostedRunner) SetMachineSizeDetails(value ActionsHostedRunnerMachineSpecable)() {
    m.machine_size_details = value
}
// SetMaximumRunners sets the maximum_runners property value. The maximum amount of hosted runners. Runners will not scale automatically above this number. Use this setting to limit your cost.
func (m *ActionsHostedRunner) SetMaximumRunners(value *int32)() {
    m.maximum_runners = value
}
// SetName sets the name property value. The name of the hosted runner.
func (m *ActionsHostedRunner) SetName(value *string)() {
    m.name = value
}
// SetPlatform sets the platform property value. The operating system of the image.
func (m *ActionsHostedRunner) SetPlatform(value *string)() {
    m.platform = value
}
// SetPublicIpEnabled sets the public_ip_enabled property value. Whether public IP is enabled for the hosted runners.
func (m *ActionsHostedRunner) SetPublicIpEnabled(value *bool)() {
    m.public_ip_enabled = value
}
// SetPublicIps sets the public_ips property value. The public IP ranges when public IP is enabled for the hosted runners.
func (m *ActionsHostedRunner) SetPublicIps(value []PublicIpable)() {
    m.public_ips = value
}
// SetRunnerGroupId sets the runner_group_id property value. The unique identifier of the group that the hosted runner belongs to.
func (m *ActionsHostedRunner) SetRunnerGroupId(value *int32)() {
    m.runner_group_id = value
}
// SetStatus sets the status property value. The status of the runner.
func (m *ActionsHostedRunner) SetStatus(value *ActionsHostedRunner_status)() {
    m.status = value
}
type ActionsHostedRunnerable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetId()(*int32)
    GetImageDetails()(NullableActionsHostedRunnerPoolImageable)
    GetLastActiveOn()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetMachineSizeDetails()(ActionsHostedRunnerMachineSpecable)
    GetMaximumRunners()(*int32)
    GetName()(*string)
    GetPlatform()(*string)
    GetPublicIpEnabled()(*bool)
    GetPublicIps()([]PublicIpable)
    GetRunnerGroupId()(*int32)
    GetStatus()(*ActionsHostedRunner_status)
    SetId(value *int32)()
    SetImageDetails(value NullableActionsHostedRunnerPoolImageable)()
    SetLastActiveOn(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetMachineSizeDetails(value ActionsHostedRunnerMachineSpecable)()
    SetMaximumRunners(value *int32)()
    SetName(value *string)()
    SetPlatform(value *string)()
    SetPublicIpEnabled(value *bool)()
    SetPublicIps(value []PublicIpable)()
    SetRunnerGroupId(value *int32)()
    SetStatus(value *ActionsHostedRunner_status)()
}
