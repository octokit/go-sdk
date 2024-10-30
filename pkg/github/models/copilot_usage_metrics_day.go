package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// CopilotUsageMetricsDay copilot usage metrics for a given day.
type CopilotUsageMetricsDay struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Usage metrics for Copilot Chat in github.com
    copilot_dotcom_chat CopilotDotcomChatable
    // Usage metrics for Copilot for pull requests.
    copilot_dotcom_pull_requests CopilotDotcomPullRequestsable
    // Usage metrics for Copilot Chat in the IDE.
    copilot_ide_chat CopilotIdeChatable
    // Usage metrics for Copilot editor code completions in the IDE.
    copilot_ide_code_completions CopilotIdeCodeCompletionsable
    // The date for which the usage metrics are aggregated, in `YYYY-MM-DD` format.
    date *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly
    // The total number of Copilot users with activity belonging to any Copilot feature, globally, for the given day. Includes passive activity such as receiving a code suggestion, as well as engagement activity such as accepting a code suggestion or prompting chat. Does not include authentication events. Is not limited to the individual features detailed on the endpoint.
    total_active_users *int32
    // The total number of Copilot users who engaged with any Copilot feature, for the given day. Examples include but are not limited to accepting a code suggestion, prompting Copilot chat, or triggering a PR Summary. Does not include authentication events. Is not limited to the individual features detailed on the endpoint.
    total_engaged_users *int32
}
// NewCopilotUsageMetricsDay instantiates a new CopilotUsageMetricsDay and sets the default values.
func NewCopilotUsageMetricsDay()(*CopilotUsageMetricsDay) {
    m := &CopilotUsageMetricsDay{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateCopilotUsageMetricsDayFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateCopilotUsageMetricsDayFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCopilotUsageMetricsDay(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *CopilotUsageMetricsDay) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetCopilotDotcomChat gets the copilot_dotcom_chat property value. Usage metrics for Copilot Chat in github.com
// returns a CopilotDotcomChatable when successful
func (m *CopilotUsageMetricsDay) GetCopilotDotcomChat()(CopilotDotcomChatable) {
    return m.copilot_dotcom_chat
}
// GetCopilotDotcomPullRequests gets the copilot_dotcom_pull_requests property value. Usage metrics for Copilot for pull requests.
// returns a CopilotDotcomPullRequestsable when successful
func (m *CopilotUsageMetricsDay) GetCopilotDotcomPullRequests()(CopilotDotcomPullRequestsable) {
    return m.copilot_dotcom_pull_requests
}
// GetCopilotIdeChat gets the copilot_ide_chat property value. Usage metrics for Copilot Chat in the IDE.
// returns a CopilotIdeChatable when successful
func (m *CopilotUsageMetricsDay) GetCopilotIdeChat()(CopilotIdeChatable) {
    return m.copilot_ide_chat
}
// GetCopilotIdeCodeCompletions gets the copilot_ide_code_completions property value. Usage metrics for Copilot editor code completions in the IDE.
// returns a CopilotIdeCodeCompletionsable when successful
func (m *CopilotUsageMetricsDay) GetCopilotIdeCodeCompletions()(CopilotIdeCodeCompletionsable) {
    return m.copilot_ide_code_completions
}
// GetDate gets the date property value. The date for which the usage metrics are aggregated, in `YYYY-MM-DD` format.
// returns a *DateOnly when successful
func (m *CopilotUsageMetricsDay) GetDate()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly) {
    return m.date
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *CopilotUsageMetricsDay) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["copilot_dotcom_chat"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateCopilotDotcomChatFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCopilotDotcomChat(val.(CopilotDotcomChatable))
        }
        return nil
    }
    res["copilot_dotcom_pull_requests"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateCopilotDotcomPullRequestsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCopilotDotcomPullRequests(val.(CopilotDotcomPullRequestsable))
        }
        return nil
    }
    res["copilot_ide_chat"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateCopilotIdeChatFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCopilotIdeChat(val.(CopilotIdeChatable))
        }
        return nil
    }
    res["copilot_ide_code_completions"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateCopilotIdeCodeCompletionsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCopilotIdeCodeCompletions(val.(CopilotIdeCodeCompletionsable))
        }
        return nil
    }
    res["date"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetDateOnlyValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDate(val)
        }
        return nil
    }
    res["total_active_users"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTotalActiveUsers(val)
        }
        return nil
    }
    res["total_engaged_users"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTotalEngagedUsers(val)
        }
        return nil
    }
    return res
}
// GetTotalActiveUsers gets the total_active_users property value. The total number of Copilot users with activity belonging to any Copilot feature, globally, for the given day. Includes passive activity such as receiving a code suggestion, as well as engagement activity such as accepting a code suggestion or prompting chat. Does not include authentication events. Is not limited to the individual features detailed on the endpoint.
// returns a *int32 when successful
func (m *CopilotUsageMetricsDay) GetTotalActiveUsers()(*int32) {
    return m.total_active_users
}
// GetTotalEngagedUsers gets the total_engaged_users property value. The total number of Copilot users who engaged with any Copilot feature, for the given day. Examples include but are not limited to accepting a code suggestion, prompting Copilot chat, or triggering a PR Summary. Does not include authentication events. Is not limited to the individual features detailed on the endpoint.
// returns a *int32 when successful
func (m *CopilotUsageMetricsDay) GetTotalEngagedUsers()(*int32) {
    return m.total_engaged_users
}
// Serialize serializes information the current object
func (m *CopilotUsageMetricsDay) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("copilot_dotcom_chat", m.GetCopilotDotcomChat())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("copilot_dotcom_pull_requests", m.GetCopilotDotcomPullRequests())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("copilot_ide_chat", m.GetCopilotIdeChat())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("copilot_ide_code_completions", m.GetCopilotIdeCodeCompletions())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteDateOnlyValue("date", m.GetDate())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("total_active_users", m.GetTotalActiveUsers())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("total_engaged_users", m.GetTotalEngagedUsers())
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
func (m *CopilotUsageMetricsDay) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetCopilotDotcomChat sets the copilot_dotcom_chat property value. Usage metrics for Copilot Chat in github.com
func (m *CopilotUsageMetricsDay) SetCopilotDotcomChat(value CopilotDotcomChatable)() {
    m.copilot_dotcom_chat = value
}
// SetCopilotDotcomPullRequests sets the copilot_dotcom_pull_requests property value. Usage metrics for Copilot for pull requests.
func (m *CopilotUsageMetricsDay) SetCopilotDotcomPullRequests(value CopilotDotcomPullRequestsable)() {
    m.copilot_dotcom_pull_requests = value
}
// SetCopilotIdeChat sets the copilot_ide_chat property value. Usage metrics for Copilot Chat in the IDE.
func (m *CopilotUsageMetricsDay) SetCopilotIdeChat(value CopilotIdeChatable)() {
    m.copilot_ide_chat = value
}
// SetCopilotIdeCodeCompletions sets the copilot_ide_code_completions property value. Usage metrics for Copilot editor code completions in the IDE.
func (m *CopilotUsageMetricsDay) SetCopilotIdeCodeCompletions(value CopilotIdeCodeCompletionsable)() {
    m.copilot_ide_code_completions = value
}
// SetDate sets the date property value. The date for which the usage metrics are aggregated, in `YYYY-MM-DD` format.
func (m *CopilotUsageMetricsDay) SetDate(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)() {
    m.date = value
}
// SetTotalActiveUsers sets the total_active_users property value. The total number of Copilot users with activity belonging to any Copilot feature, globally, for the given day. Includes passive activity such as receiving a code suggestion, as well as engagement activity such as accepting a code suggestion or prompting chat. Does not include authentication events. Is not limited to the individual features detailed on the endpoint.
func (m *CopilotUsageMetricsDay) SetTotalActiveUsers(value *int32)() {
    m.total_active_users = value
}
// SetTotalEngagedUsers sets the total_engaged_users property value. The total number of Copilot users who engaged with any Copilot feature, for the given day. Examples include but are not limited to accepting a code suggestion, prompting Copilot chat, or triggering a PR Summary. Does not include authentication events. Is not limited to the individual features detailed on the endpoint.
func (m *CopilotUsageMetricsDay) SetTotalEngagedUsers(value *int32)() {
    m.total_engaged_users = value
}
type CopilotUsageMetricsDayable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCopilotDotcomChat()(CopilotDotcomChatable)
    GetCopilotDotcomPullRequests()(CopilotDotcomPullRequestsable)
    GetCopilotIdeChat()(CopilotIdeChatable)
    GetCopilotIdeCodeCompletions()(CopilotIdeCodeCompletionsable)
    GetDate()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)
    GetTotalActiveUsers()(*int32)
    GetTotalEngagedUsers()(*int32)
    SetCopilotDotcomChat(value CopilotDotcomChatable)()
    SetCopilotDotcomPullRequests(value CopilotDotcomPullRequestsable)()
    SetCopilotIdeChat(value CopilotIdeChatable)()
    SetCopilotIdeCodeCompletions(value CopilotIdeCodeCompletionsable)()
    SetDate(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)()
    SetTotalActiveUsers(value *int32)()
    SetTotalEngagedUsers(value *int32)()
}
