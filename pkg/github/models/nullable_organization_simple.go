package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// NullableOrganizationSimple a GitHub organization.
type NullableOrganizationSimple struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The avatar_url property
    avatar_url *string
    // The description property
    description *string
    // The events_url property
    events_url *string
    // The hooks_url property
    hooks_url *string
    // The id property
    id *int32
    // The issues_url property
    issues_url *string
    // The login property
    login *string
    // The members_url property
    members_url *string
    // The node_id property
    node_id *string
    // The public_members_url property
    public_members_url *string
    // The repos_url property
    repos_url *string
    // The url property
    url *string
}
// NewNullableOrganizationSimple instantiates a new NullableOrganizationSimple and sets the default values.
func NewNullableOrganizationSimple()(*NullableOrganizationSimple) {
    m := &NullableOrganizationSimple{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateNullableOrganizationSimpleFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateNullableOrganizationSimpleFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewNullableOrganizationSimple(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *NullableOrganizationSimple) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetAvatarUrl gets the avatar_url property value. The avatar_url property
// returns a *string when successful
func (m *NullableOrganizationSimple) GetAvatarUrl()(*string) {
    return m.avatar_url
}
// GetDescription gets the description property value. The description property
// returns a *string when successful
func (m *NullableOrganizationSimple) GetDescription()(*string) {
    return m.description
}
// GetEventsUrl gets the events_url property value. The events_url property
// returns a *string when successful
func (m *NullableOrganizationSimple) GetEventsUrl()(*string) {
    return m.events_url
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *NullableOrganizationSimple) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["avatar_url"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAvatarUrl(val)
        }
        return nil
    }
    res["description"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDescription(val)
        }
        return nil
    }
    res["events_url"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEventsUrl(val)
        }
        return nil
    }
    res["hooks_url"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHooksUrl(val)
        }
        return nil
    }
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
    res["issues_url"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIssuesUrl(val)
        }
        return nil
    }
    res["login"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLogin(val)
        }
        return nil
    }
    res["members_url"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMembersUrl(val)
        }
        return nil
    }
    res["node_id"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNodeId(val)
        }
        return nil
    }
    res["public_members_url"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPublicMembersUrl(val)
        }
        return nil
    }
    res["repos_url"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetReposUrl(val)
        }
        return nil
    }
    res["url"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUrl(val)
        }
        return nil
    }
    return res
}
// GetHooksUrl gets the hooks_url property value. The hooks_url property
// returns a *string when successful
func (m *NullableOrganizationSimple) GetHooksUrl()(*string) {
    return m.hooks_url
}
// GetId gets the id property value. The id property
// returns a *int32 when successful
func (m *NullableOrganizationSimple) GetId()(*int32) {
    return m.id
}
// GetIssuesUrl gets the issues_url property value. The issues_url property
// returns a *string when successful
func (m *NullableOrganizationSimple) GetIssuesUrl()(*string) {
    return m.issues_url
}
// GetLogin gets the login property value. The login property
// returns a *string when successful
func (m *NullableOrganizationSimple) GetLogin()(*string) {
    return m.login
}
// GetMembersUrl gets the members_url property value. The members_url property
// returns a *string when successful
func (m *NullableOrganizationSimple) GetMembersUrl()(*string) {
    return m.members_url
}
// GetNodeId gets the node_id property value. The node_id property
// returns a *string when successful
func (m *NullableOrganizationSimple) GetNodeId()(*string) {
    return m.node_id
}
// GetPublicMembersUrl gets the public_members_url property value. The public_members_url property
// returns a *string when successful
func (m *NullableOrganizationSimple) GetPublicMembersUrl()(*string) {
    return m.public_members_url
}
// GetReposUrl gets the repos_url property value. The repos_url property
// returns a *string when successful
func (m *NullableOrganizationSimple) GetReposUrl()(*string) {
    return m.repos_url
}
// GetUrl gets the url property value. The url property
// returns a *string when successful
func (m *NullableOrganizationSimple) GetUrl()(*string) {
    return m.url
}
// Serialize serializes information the current object
func (m *NullableOrganizationSimple) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("avatar_url", m.GetAvatarUrl())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("description", m.GetDescription())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("events_url", m.GetEventsUrl())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("hooks_url", m.GetHooksUrl())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("id", m.GetId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("issues_url", m.GetIssuesUrl())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("login", m.GetLogin())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("members_url", m.GetMembersUrl())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("node_id", m.GetNodeId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("public_members_url", m.GetPublicMembersUrl())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("repos_url", m.GetReposUrl())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("url", m.GetUrl())
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
func (m *NullableOrganizationSimple) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetAvatarUrl sets the avatar_url property value. The avatar_url property
func (m *NullableOrganizationSimple) SetAvatarUrl(value *string)() {
    m.avatar_url = value
}
// SetDescription sets the description property value. The description property
func (m *NullableOrganizationSimple) SetDescription(value *string)() {
    m.description = value
}
// SetEventsUrl sets the events_url property value. The events_url property
func (m *NullableOrganizationSimple) SetEventsUrl(value *string)() {
    m.events_url = value
}
// SetHooksUrl sets the hooks_url property value. The hooks_url property
func (m *NullableOrganizationSimple) SetHooksUrl(value *string)() {
    m.hooks_url = value
}
// SetId sets the id property value. The id property
func (m *NullableOrganizationSimple) SetId(value *int32)() {
    m.id = value
}
// SetIssuesUrl sets the issues_url property value. The issues_url property
func (m *NullableOrganizationSimple) SetIssuesUrl(value *string)() {
    m.issues_url = value
}
// SetLogin sets the login property value. The login property
func (m *NullableOrganizationSimple) SetLogin(value *string)() {
    m.login = value
}
// SetMembersUrl sets the members_url property value. The members_url property
func (m *NullableOrganizationSimple) SetMembersUrl(value *string)() {
    m.members_url = value
}
// SetNodeId sets the node_id property value. The node_id property
func (m *NullableOrganizationSimple) SetNodeId(value *string)() {
    m.node_id = value
}
// SetPublicMembersUrl sets the public_members_url property value. The public_members_url property
func (m *NullableOrganizationSimple) SetPublicMembersUrl(value *string)() {
    m.public_members_url = value
}
// SetReposUrl sets the repos_url property value. The repos_url property
func (m *NullableOrganizationSimple) SetReposUrl(value *string)() {
    m.repos_url = value
}
// SetUrl sets the url property value. The url property
func (m *NullableOrganizationSimple) SetUrl(value *string)() {
    m.url = value
}
type NullableOrganizationSimpleable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAvatarUrl()(*string)
    GetDescription()(*string)
    GetEventsUrl()(*string)
    GetHooksUrl()(*string)
    GetId()(*int32)
    GetIssuesUrl()(*string)
    GetLogin()(*string)
    GetMembersUrl()(*string)
    GetNodeId()(*string)
    GetPublicMembersUrl()(*string)
    GetReposUrl()(*string)
    GetUrl()(*string)
    SetAvatarUrl(value *string)()
    SetDescription(value *string)()
    SetEventsUrl(value *string)()
    SetHooksUrl(value *string)()
    SetId(value *int32)()
    SetIssuesUrl(value *string)()
    SetLogin(value *string)()
    SetMembersUrl(value *string)()
    SetNodeId(value *string)()
    SetPublicMembersUrl(value *string)()
    SetReposUrl(value *string)()
    SetUrl(value *string)()
}
