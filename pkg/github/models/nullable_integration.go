package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// NullableIntegration gitHub apps are a new way to extend GitHub. They can be installed directly on organizations and user accounts and granted access to specific repositories. They come with granular permissions and built-in webhooks. GitHub apps are first class actors within GitHub.
type NullableIntegration struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The client_id property
    client_id *string
    // The client_secret property
    client_secret *string
    // The created_at property
    created_at *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The description property
    description *string
    // The list of events for the GitHub app
    events []string
    // The external_url property
    external_url *string
    // The html_url property
    html_url *string
    // Unique identifier of the GitHub app
    id *int32
    // The number of installations associated with the GitHub app
    installations_count *int32
    // The name of the GitHub app
    name *string
    // The node_id property
    node_id *string
    // The owner property
    owner NullableIntegration_NullableIntegration_ownerable
    // The pem property
    pem *string
    // The set of permissions for the GitHub app
    permissions NullableIntegration_permissionsable
    // The slug name of the GitHub app
    slug *string
    // The updated_at property
    updated_at *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The webhook_secret property
    webhook_secret *string
}
// NullableIntegration_NullableIntegration_owner composed type wrapper for classes Enterpriseable, SimpleUserable
type NullableIntegration_NullableIntegration_owner struct {
    // Composed type representation for type Enterpriseable
    enterprise Enterpriseable
    // Composed type representation for type SimpleUserable
    simpleUser SimpleUserable
}
// NewNullableIntegration_NullableIntegration_owner instantiates a new NullableIntegration_NullableIntegration_owner and sets the default values.
func NewNullableIntegration_NullableIntegration_owner()(*NullableIntegration_NullableIntegration_owner) {
    m := &NullableIntegration_NullableIntegration_owner{
    }
    return m
}
// CreateNullableIntegration_NullableIntegration_ownerFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateNullableIntegration_NullableIntegration_ownerFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    result := NewNullableIntegration_NullableIntegration_owner()
    if parseNode != nil {
        mappingValueNode, err := parseNode.GetChildNode("")
        if err != nil {
            return nil, err
        }
        if mappingValueNode != nil {
            mappingValue, err := mappingValueNode.GetStringValue()
            if err != nil {
                return nil, err
            }
            if mappingValue != nil {
            }
        }
    }
    return result, nil
}
// GetEnterprise gets the enterprise property value. Composed type representation for type Enterpriseable
// returns a Enterpriseable when successful
func (m *NullableIntegration_NullableIntegration_owner) GetEnterprise()(Enterpriseable) {
    return m.enterprise
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *NullableIntegration_NullableIntegration_owner) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    return make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
}
// GetIsComposedType determines if the current object is a wrapper around a composed type
// returns a bool when successful
func (m *NullableIntegration_NullableIntegration_owner) GetIsComposedType()(bool) {
    return true
}
// GetSimpleUser gets the simpleUser property value. Composed type representation for type SimpleUserable
// returns a SimpleUserable when successful
func (m *NullableIntegration_NullableIntegration_owner) GetSimpleUser()(SimpleUserable) {
    return m.simpleUser
}
// Serialize serializes information the current object
func (m *NullableIntegration_NullableIntegration_owner) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetEnterprise() != nil {
        err := writer.WriteObjectValue("", m.GetEnterprise())
        if err != nil {
            return err
        }
    } else if m.GetSimpleUser() != nil {
        err := writer.WriteObjectValue("", m.GetSimpleUser())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetEnterprise sets the enterprise property value. Composed type representation for type Enterpriseable
func (m *NullableIntegration_NullableIntegration_owner) SetEnterprise(value Enterpriseable)() {
    m.enterprise = value
}
// SetSimpleUser sets the simpleUser property value. Composed type representation for type SimpleUserable
func (m *NullableIntegration_NullableIntegration_owner) SetSimpleUser(value SimpleUserable)() {
    m.simpleUser = value
}
type NullableIntegration_NullableIntegration_ownerable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetEnterprise()(Enterpriseable)
    GetSimpleUser()(SimpleUserable)
    SetEnterprise(value Enterpriseable)()
    SetSimpleUser(value SimpleUserable)()
}
// NewNullableIntegration instantiates a new NullableIntegration and sets the default values.
func NewNullableIntegration()(*NullableIntegration) {
    m := &NullableIntegration{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateNullableIntegrationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateNullableIntegrationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewNullableIntegration(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *NullableIntegration) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetClientId gets the client_id property value. The client_id property
// returns a *string when successful
func (m *NullableIntegration) GetClientId()(*string) {
    return m.client_id
}
// GetClientSecret gets the client_secret property value. The client_secret property
// returns a *string when successful
func (m *NullableIntegration) GetClientSecret()(*string) {
    return m.client_secret
}
// GetCreatedAt gets the created_at property value. The created_at property
// returns a *Time when successful
func (m *NullableIntegration) GetCreatedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.created_at
}
// GetDescription gets the description property value. The description property
// returns a *string when successful
func (m *NullableIntegration) GetDescription()(*string) {
    return m.description
}
// GetEvents gets the events property value. The list of events for the GitHub app
// returns a []string when successful
func (m *NullableIntegration) GetEvents()([]string) {
    return m.events
}
// GetExternalUrl gets the external_url property value. The external_url property
// returns a *string when successful
func (m *NullableIntegration) GetExternalUrl()(*string) {
    return m.external_url
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *NullableIntegration) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["client_id"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetClientId(val)
        }
        return nil
    }
    res["client_secret"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetClientSecret(val)
        }
        return nil
    }
    res["created_at"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCreatedAt(val)
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
    res["events"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
            m.SetEvents(res)
        }
        return nil
    }
    res["external_url"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetExternalUrl(val)
        }
        return nil
    }
    res["html_url"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHtmlUrl(val)
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
    res["installations_count"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetInstallationsCount(val)
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
    res["owner"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateNullableIntegration_NullableIntegration_ownerFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOwner(val.(NullableIntegration_NullableIntegration_ownerable))
        }
        return nil
    }
    res["pem"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPem(val)
        }
        return nil
    }
    res["permissions"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateNullableIntegration_permissionsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPermissions(val.(NullableIntegration_permissionsable))
        }
        return nil
    }
    res["slug"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSlug(val)
        }
        return nil
    }
    res["updated_at"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUpdatedAt(val)
        }
        return nil
    }
    res["webhook_secret"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWebhookSecret(val)
        }
        return nil
    }
    return res
}
// GetHtmlUrl gets the html_url property value. The html_url property
// returns a *string when successful
func (m *NullableIntegration) GetHtmlUrl()(*string) {
    return m.html_url
}
// GetId gets the id property value. Unique identifier of the GitHub app
// returns a *int32 when successful
func (m *NullableIntegration) GetId()(*int32) {
    return m.id
}
// GetInstallationsCount gets the installations_count property value. The number of installations associated with the GitHub app
// returns a *int32 when successful
func (m *NullableIntegration) GetInstallationsCount()(*int32) {
    return m.installations_count
}
// GetName gets the name property value. The name of the GitHub app
// returns a *string when successful
func (m *NullableIntegration) GetName()(*string) {
    return m.name
}
// GetNodeId gets the node_id property value. The node_id property
// returns a *string when successful
func (m *NullableIntegration) GetNodeId()(*string) {
    return m.node_id
}
// GetOwner gets the owner property value. The owner property
// returns a NullableIntegration_NullableIntegration_ownerable when successful
func (m *NullableIntegration) GetOwner()(NullableIntegration_NullableIntegration_ownerable) {
    return m.owner
}
// GetPem gets the pem property value. The pem property
// returns a *string when successful
func (m *NullableIntegration) GetPem()(*string) {
    return m.pem
}
// GetPermissions gets the permissions property value. The set of permissions for the GitHub app
// returns a NullableIntegration_permissionsable when successful
func (m *NullableIntegration) GetPermissions()(NullableIntegration_permissionsable) {
    return m.permissions
}
// GetSlug gets the slug property value. The slug name of the GitHub app
// returns a *string when successful
func (m *NullableIntegration) GetSlug()(*string) {
    return m.slug
}
// GetUpdatedAt gets the updated_at property value. The updated_at property
// returns a *Time when successful
func (m *NullableIntegration) GetUpdatedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.updated_at
}
// GetWebhookSecret gets the webhook_secret property value. The webhook_secret property
// returns a *string when successful
func (m *NullableIntegration) GetWebhookSecret()(*string) {
    return m.webhook_secret
}
// Serialize serializes information the current object
func (m *NullableIntegration) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("client_id", m.GetClientId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("client_secret", m.GetClientSecret())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteTimeValue("created_at", m.GetCreatedAt())
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
    if m.GetEvents() != nil {
        err := writer.WriteCollectionOfStringValues("events", m.GetEvents())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("external_url", m.GetExternalUrl())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("html_url", m.GetHtmlUrl())
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
        err := writer.WriteInt32Value("installations_count", m.GetInstallationsCount())
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
        err := writer.WriteStringValue("node_id", m.GetNodeId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("owner", m.GetOwner())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("pem", m.GetPem())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("permissions", m.GetPermissions())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("slug", m.GetSlug())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteTimeValue("updated_at", m.GetUpdatedAt())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("webhook_secret", m.GetWebhookSecret())
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
func (m *NullableIntegration) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetClientId sets the client_id property value. The client_id property
func (m *NullableIntegration) SetClientId(value *string)() {
    m.client_id = value
}
// SetClientSecret sets the client_secret property value. The client_secret property
func (m *NullableIntegration) SetClientSecret(value *string)() {
    m.client_secret = value
}
// SetCreatedAt sets the created_at property value. The created_at property
func (m *NullableIntegration) SetCreatedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.created_at = value
}
// SetDescription sets the description property value. The description property
func (m *NullableIntegration) SetDescription(value *string)() {
    m.description = value
}
// SetEvents sets the events property value. The list of events for the GitHub app
func (m *NullableIntegration) SetEvents(value []string)() {
    m.events = value
}
// SetExternalUrl sets the external_url property value. The external_url property
func (m *NullableIntegration) SetExternalUrl(value *string)() {
    m.external_url = value
}
// SetHtmlUrl sets the html_url property value. The html_url property
func (m *NullableIntegration) SetHtmlUrl(value *string)() {
    m.html_url = value
}
// SetId sets the id property value. Unique identifier of the GitHub app
func (m *NullableIntegration) SetId(value *int32)() {
    m.id = value
}
// SetInstallationsCount sets the installations_count property value. The number of installations associated with the GitHub app
func (m *NullableIntegration) SetInstallationsCount(value *int32)() {
    m.installations_count = value
}
// SetName sets the name property value. The name of the GitHub app
func (m *NullableIntegration) SetName(value *string)() {
    m.name = value
}
// SetNodeId sets the node_id property value. The node_id property
func (m *NullableIntegration) SetNodeId(value *string)() {
    m.node_id = value
}
// SetOwner sets the owner property value. The owner property
func (m *NullableIntegration) SetOwner(value NullableIntegration_NullableIntegration_ownerable)() {
    m.owner = value
}
// SetPem sets the pem property value. The pem property
func (m *NullableIntegration) SetPem(value *string)() {
    m.pem = value
}
// SetPermissions sets the permissions property value. The set of permissions for the GitHub app
func (m *NullableIntegration) SetPermissions(value NullableIntegration_permissionsable)() {
    m.permissions = value
}
// SetSlug sets the slug property value. The slug name of the GitHub app
func (m *NullableIntegration) SetSlug(value *string)() {
    m.slug = value
}
// SetUpdatedAt sets the updated_at property value. The updated_at property
func (m *NullableIntegration) SetUpdatedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.updated_at = value
}
// SetWebhookSecret sets the webhook_secret property value. The webhook_secret property
func (m *NullableIntegration) SetWebhookSecret(value *string)() {
    m.webhook_secret = value
}
type NullableIntegrationable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetClientId()(*string)
    GetClientSecret()(*string)
    GetCreatedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetDescription()(*string)
    GetEvents()([]string)
    GetExternalUrl()(*string)
    GetHtmlUrl()(*string)
    GetId()(*int32)
    GetInstallationsCount()(*int32)
    GetName()(*string)
    GetNodeId()(*string)
    GetOwner()(NullableIntegration_NullableIntegration_ownerable)
    GetPem()(*string)
    GetPermissions()(NullableIntegration_permissionsable)
    GetSlug()(*string)
    GetUpdatedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetWebhookSecret()(*string)
    SetClientId(value *string)()
    SetClientSecret(value *string)()
    SetCreatedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetDescription(value *string)()
    SetEvents(value []string)()
    SetExternalUrl(value *string)()
    SetHtmlUrl(value *string)()
    SetId(value *int32)()
    SetInstallationsCount(value *int32)()
    SetName(value *string)()
    SetNodeId(value *string)()
    SetOwner(value NullableIntegration_NullableIntegration_ownerable)()
    SetPem(value *string)()
    SetPermissions(value NullableIntegration_permissionsable)()
    SetSlug(value *string)()
    SetUpdatedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetWebhookSecret(value *string)()
}
