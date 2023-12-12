package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// UnassignedIssueEvent unassigned Issue Event
type UnassignedIssueEvent struct {
    // A GitHub user.
    actor SimpleUserable
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // A GitHub user.
    assignee SimpleUserable
    // A GitHub user.
    assigner SimpleUserable
    // The commit_id property
    commit_id *string
    // The commit_url property
    commit_url *string
    // The created_at property
    created_at *string
    // The event property
    event *string
    // The id property
    id *int32
    // The node_id property
    node_id *string
    // GitHub apps are a new way to extend GitHub. They can be installed directly on organizations and user accounts and granted access to specific repositories. They come with granular permissions and built-in webhooks. GitHub apps are first class actors within GitHub.
    performed_via_github_app NullableIntegrationable
    // The url property
    url *string
}
// NewUnassignedIssueEvent instantiates a new unassignedIssueEvent and sets the default values.
func NewUnassignedIssueEvent()(*UnassignedIssueEvent) {
    m := &UnassignedIssueEvent{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateUnassignedIssueEventFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateUnassignedIssueEventFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewUnassignedIssueEvent(), nil
}
// GetActor gets the actor property value. A GitHub user.
func (m *UnassignedIssueEvent) GetActor()(SimpleUserable) {
    return m.actor
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *UnassignedIssueEvent) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetAssignee gets the assignee property value. A GitHub user.
func (m *UnassignedIssueEvent) GetAssignee()(SimpleUserable) {
    return m.assignee
}
// GetAssigner gets the assigner property value. A GitHub user.
func (m *UnassignedIssueEvent) GetAssigner()(SimpleUserable) {
    return m.assigner
}
// GetCommitId gets the commit_id property value. The commit_id property
func (m *UnassignedIssueEvent) GetCommitId()(*string) {
    return m.commit_id
}
// GetCommitUrl gets the commit_url property value. The commit_url property
func (m *UnassignedIssueEvent) GetCommitUrl()(*string) {
    return m.commit_url
}
// GetCreatedAt gets the created_at property value. The created_at property
func (m *UnassignedIssueEvent) GetCreatedAt()(*string) {
    return m.created_at
}
// GetEvent gets the event property value. The event property
func (m *UnassignedIssueEvent) GetEvent()(*string) {
    return m.event
}
// GetFieldDeserializers the deserialization information for the current model
func (m *UnassignedIssueEvent) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["actor"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateSimpleUserFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetActor(val.(SimpleUserable))
        }
        return nil
    }
    res["assignee"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateSimpleUserFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAssignee(val.(SimpleUserable))
        }
        return nil
    }
    res["assigner"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateSimpleUserFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAssigner(val.(SimpleUserable))
        }
        return nil
    }
    res["commit_id"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCommitId(val)
        }
        return nil
    }
    res["commit_url"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCommitUrl(val)
        }
        return nil
    }
    res["created_at"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCreatedAt(val)
        }
        return nil
    }
    res["event"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEvent(val)
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
    res["performed_via_github_app"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateNullableIntegrationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPerformedViaGithubApp(val.(NullableIntegrationable))
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
// GetId gets the id property value. The id property
func (m *UnassignedIssueEvent) GetId()(*int32) {
    return m.id
}
// GetNodeId gets the node_id property value. The node_id property
func (m *UnassignedIssueEvent) GetNodeId()(*string) {
    return m.node_id
}
// GetPerformedViaGithubApp gets the performed_via_github_app property value. GitHub apps are a new way to extend GitHub. They can be installed directly on organizations and user accounts and granted access to specific repositories. They come with granular permissions and built-in webhooks. GitHub apps are first class actors within GitHub.
func (m *UnassignedIssueEvent) GetPerformedViaGithubApp()(NullableIntegrationable) {
    return m.performed_via_github_app
}
// GetUrl gets the url property value. The url property
func (m *UnassignedIssueEvent) GetUrl()(*string) {
    return m.url
}
// Serialize serializes information the current object
func (m *UnassignedIssueEvent) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("actor", m.GetActor())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("assignee", m.GetAssignee())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("assigner", m.GetAssigner())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("commit_id", m.GetCommitId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("commit_url", m.GetCommitUrl())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("created_at", m.GetCreatedAt())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("event", m.GetEvent())
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
        err := writer.WriteStringValue("node_id", m.GetNodeId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("performed_via_github_app", m.GetPerformedViaGithubApp())
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
// SetActor sets the actor property value. A GitHub user.
func (m *UnassignedIssueEvent) SetActor(value SimpleUserable)() {
    m.actor = value
}
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *UnassignedIssueEvent) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetAssignee sets the assignee property value. A GitHub user.
func (m *UnassignedIssueEvent) SetAssignee(value SimpleUserable)() {
    m.assignee = value
}
// SetAssigner sets the assigner property value. A GitHub user.
func (m *UnassignedIssueEvent) SetAssigner(value SimpleUserable)() {
    m.assigner = value
}
// SetCommitId sets the commit_id property value. The commit_id property
func (m *UnassignedIssueEvent) SetCommitId(value *string)() {
    m.commit_id = value
}
// SetCommitUrl sets the commit_url property value. The commit_url property
func (m *UnassignedIssueEvent) SetCommitUrl(value *string)() {
    m.commit_url = value
}
// SetCreatedAt sets the created_at property value. The created_at property
func (m *UnassignedIssueEvent) SetCreatedAt(value *string)() {
    m.created_at = value
}
// SetEvent sets the event property value. The event property
func (m *UnassignedIssueEvent) SetEvent(value *string)() {
    m.event = value
}
// SetId sets the id property value. The id property
func (m *UnassignedIssueEvent) SetId(value *int32)() {
    m.id = value
}
// SetNodeId sets the node_id property value. The node_id property
func (m *UnassignedIssueEvent) SetNodeId(value *string)() {
    m.node_id = value
}
// SetPerformedViaGithubApp sets the performed_via_github_app property value. GitHub apps are a new way to extend GitHub. They can be installed directly on organizations and user accounts and granted access to specific repositories. They come with granular permissions and built-in webhooks. GitHub apps are first class actors within GitHub.
func (m *UnassignedIssueEvent) SetPerformedViaGithubApp(value NullableIntegrationable)() {
    m.performed_via_github_app = value
}
// SetUrl sets the url property value. The url property
func (m *UnassignedIssueEvent) SetUrl(value *string)() {
    m.url = value
}
// UnassignedIssueEventable 
type UnassignedIssueEventable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetActor()(SimpleUserable)
    GetAssignee()(SimpleUserable)
    GetAssigner()(SimpleUserable)
    GetCommitId()(*string)
    GetCommitUrl()(*string)
    GetCreatedAt()(*string)
    GetEvent()(*string)
    GetId()(*int32)
    GetNodeId()(*string)
    GetPerformedViaGithubApp()(NullableIntegrationable)
    GetUrl()(*string)
    SetActor(value SimpleUserable)()
    SetAssignee(value SimpleUserable)()
    SetAssigner(value SimpleUserable)()
    SetCommitId(value *string)()
    SetCommitUrl(value *string)()
    SetCreatedAt(value *string)()
    SetEvent(value *string)()
    SetId(value *int32)()
    SetNodeId(value *string)()
    SetPerformedViaGithubApp(value NullableIntegrationable)()
    SetUrl(value *string)()
}
