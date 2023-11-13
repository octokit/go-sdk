package repos

import (
    i33bcfa37677dfd8b71eaf22ad3bfd140ba1702e7a8418558499c22418bf2fef2 "github.com/octokit/go-sdk/models"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemItemActionsPermissionsPutRequestBody 
type ItemItemActionsPermissionsPutRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The permissions policy that controls the actions and reusable workflows that are allowed to run.
    allowed_actions *i33bcfa37677dfd8b71eaf22ad3bfd140ba1702e7a8418558499c22418bf2fef2.AllowedActions
    // Whether GitHub Actions is enabled on the repository.
    enabled *bool
}
// NewItemItemActionsPermissionsPutRequestBody instantiates a new ItemItemActionsPermissionsPutRequestBody and sets the default values.
func NewItemItemActionsPermissionsPutRequestBody()(*ItemItemActionsPermissionsPutRequestBody) {
    m := &ItemItemActionsPermissionsPutRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateItemItemActionsPermissionsPutRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemItemActionsPermissionsPutRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemItemActionsPermissionsPutRequestBody(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ItemItemActionsPermissionsPutRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetAllowedActions gets the allowed_actions property value. The permissions policy that controls the actions and reusable workflows that are allowed to run.
func (m *ItemItemActionsPermissionsPutRequestBody) GetAllowedActions()(*i33bcfa37677dfd8b71eaf22ad3bfd140ba1702e7a8418558499c22418bf2fef2.AllowedActions) {
    return m.allowed_actions
}
// GetEnabled gets the enabled property value. Whether GitHub Actions is enabled on the repository.
func (m *ItemItemActionsPermissionsPutRequestBody) GetEnabled()(*bool) {
    return m.enabled
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ItemItemActionsPermissionsPutRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["allowed_actions"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(i33bcfa37677dfd8b71eaf22ad3bfd140ba1702e7a8418558499c22418bf2fef2.ParseAllowedActions)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAllowedActions(val.(*i33bcfa37677dfd8b71eaf22ad3bfd140ba1702e7a8418558499c22418bf2fef2.AllowedActions))
        }
        return nil
    }
    res["enabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEnabled(val)
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *ItemItemActionsPermissionsPutRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetAllowedActions() != nil {
        cast := (*m.GetAllowedActions()).String()
        err := writer.WriteStringValue("allowed_actions", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("enabled", m.GetEnabled())
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
func (m *ItemItemActionsPermissionsPutRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetAllowedActions sets the allowed_actions property value. The permissions policy that controls the actions and reusable workflows that are allowed to run.
func (m *ItemItemActionsPermissionsPutRequestBody) SetAllowedActions(value *i33bcfa37677dfd8b71eaf22ad3bfd140ba1702e7a8418558499c22418bf2fef2.AllowedActions)() {
    m.allowed_actions = value
}
// SetEnabled sets the enabled property value. Whether GitHub Actions is enabled on the repository.
func (m *ItemItemActionsPermissionsPutRequestBody) SetEnabled(value *bool)() {
    m.enabled = value
}
// ItemItemActionsPermissionsPutRequestBodyable 
type ItemItemActionsPermissionsPutRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAllowedActions()(*i33bcfa37677dfd8b71eaf22ad3bfd140ba1702e7a8418558499c22418bf2fef2.AllowedActions)
    GetEnabled()(*bool)
    SetAllowedActions(value *i33bcfa37677dfd8b71eaf22ad3bfd140ba1702e7a8418558499c22418bf2fef2.AllowedActions)()
    SetEnabled(value *bool)()
}
