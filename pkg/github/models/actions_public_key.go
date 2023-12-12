package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ActionsPublicKey the public key used for setting Actions Secrets.
type ActionsPublicKey struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The created_at property
    created_at *string
    // The id property
    id *int32
    // The Base64 encoded public key.
    key *string
    // The identifier for the key.
    key_id *string
    // The title property
    title *string
    // The url property
    url *string
}
// NewActionsPublicKey instantiates a new actionsPublicKey and sets the default values.
func NewActionsPublicKey()(*ActionsPublicKey) {
    m := &ActionsPublicKey{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateActionsPublicKeyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateActionsPublicKeyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewActionsPublicKey(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ActionsPublicKey) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetCreatedAt gets the created_at property value. The created_at property
func (m *ActionsPublicKey) GetCreatedAt()(*string) {
    return m.created_at
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ActionsPublicKey) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
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
    res["key"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetKey(val)
        }
        return nil
    }
    res["key_id"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetKeyId(val)
        }
        return nil
    }
    res["title"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTitle(val)
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
func (m *ActionsPublicKey) GetId()(*int32) {
    return m.id
}
// GetKey gets the key property value. The Base64 encoded public key.
func (m *ActionsPublicKey) GetKey()(*string) {
    return m.key
}
// GetKeyId gets the key_id property value. The identifier for the key.
func (m *ActionsPublicKey) GetKeyId()(*string) {
    return m.key_id
}
// GetTitle gets the title property value. The title property
func (m *ActionsPublicKey) GetTitle()(*string) {
    return m.title
}
// GetUrl gets the url property value. The url property
func (m *ActionsPublicKey) GetUrl()(*string) {
    return m.url
}
// Serialize serializes information the current object
func (m *ActionsPublicKey) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("created_at", m.GetCreatedAt())
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
        err := writer.WriteStringValue("key", m.GetKey())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("key_id", m.GetKeyId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("title", m.GetTitle())
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
func (m *ActionsPublicKey) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetCreatedAt sets the created_at property value. The created_at property
func (m *ActionsPublicKey) SetCreatedAt(value *string)() {
    m.created_at = value
}
// SetId sets the id property value. The id property
func (m *ActionsPublicKey) SetId(value *int32)() {
    m.id = value
}
// SetKey sets the key property value. The Base64 encoded public key.
func (m *ActionsPublicKey) SetKey(value *string)() {
    m.key = value
}
// SetKeyId sets the key_id property value. The identifier for the key.
func (m *ActionsPublicKey) SetKeyId(value *string)() {
    m.key_id = value
}
// SetTitle sets the title property value. The title property
func (m *ActionsPublicKey) SetTitle(value *string)() {
    m.title = value
}
// SetUrl sets the url property value. The url property
func (m *ActionsPublicKey) SetUrl(value *string)() {
    m.url = value
}
// ActionsPublicKeyable 
type ActionsPublicKeyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCreatedAt()(*string)
    GetId()(*int32)
    GetKey()(*string)
    GetKeyId()(*string)
    GetTitle()(*string)
    GetUrl()(*string)
    SetCreatedAt(value *string)()
    SetId(value *int32)()
    SetKey(value *string)()
    SetKeyId(value *string)()
    SetTitle(value *string)()
    SetUrl(value *string)()
}
