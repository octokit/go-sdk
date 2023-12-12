package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// NullableLicenseSimple license Simple
type NullableLicenseSimple struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The html_url property
    html_url *string
    // The key property
    key *string
    // The name property
    name *string
    // The node_id property
    node_id *string
    // The spdx_id property
    spdx_id *string
    // The url property
    url *string
}
// NewNullableLicenseSimple instantiates a new nullableLicenseSimple and sets the default values.
func NewNullableLicenseSimple()(*NullableLicenseSimple) {
    m := &NullableLicenseSimple{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateNullableLicenseSimpleFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateNullableLicenseSimpleFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewNullableLicenseSimple(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *NullableLicenseSimple) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *NullableLicenseSimple) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
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
    res["spdx_id"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSpdxId(val)
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
// GetHtmlUrl gets the html_url property value. The html_url property
func (m *NullableLicenseSimple) GetHtmlUrl()(*string) {
    return m.html_url
}
// GetKey gets the key property value. The key property
func (m *NullableLicenseSimple) GetKey()(*string) {
    return m.key
}
// GetName gets the name property value. The name property
func (m *NullableLicenseSimple) GetName()(*string) {
    return m.name
}
// GetNodeId gets the node_id property value. The node_id property
func (m *NullableLicenseSimple) GetNodeId()(*string) {
    return m.node_id
}
// GetSpdxId gets the spdx_id property value. The spdx_id property
func (m *NullableLicenseSimple) GetSpdxId()(*string) {
    return m.spdx_id
}
// GetUrl gets the url property value. The url property
func (m *NullableLicenseSimple) GetUrl()(*string) {
    return m.url
}
// Serialize serializes information the current object
func (m *NullableLicenseSimple) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("html_url", m.GetHtmlUrl())
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
        err := writer.WriteStringValue("spdx_id", m.GetSpdxId())
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
func (m *NullableLicenseSimple) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetHtmlUrl sets the html_url property value. The html_url property
func (m *NullableLicenseSimple) SetHtmlUrl(value *string)() {
    m.html_url = value
}
// SetKey sets the key property value. The key property
func (m *NullableLicenseSimple) SetKey(value *string)() {
    m.key = value
}
// SetName sets the name property value. The name property
func (m *NullableLicenseSimple) SetName(value *string)() {
    m.name = value
}
// SetNodeId sets the node_id property value. The node_id property
func (m *NullableLicenseSimple) SetNodeId(value *string)() {
    m.node_id = value
}
// SetSpdxId sets the spdx_id property value. The spdx_id property
func (m *NullableLicenseSimple) SetSpdxId(value *string)() {
    m.spdx_id = value
}
// SetUrl sets the url property value. The url property
func (m *NullableLicenseSimple) SetUrl(value *string)() {
    m.url = value
}
// NullableLicenseSimpleable 
type NullableLicenseSimpleable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetHtmlUrl()(*string)
    GetKey()(*string)
    GetName()(*string)
    GetNodeId()(*string)
    GetSpdxId()(*string)
    GetUrl()(*string)
    SetHtmlUrl(value *string)()
    SetKey(value *string)()
    SetName(value *string)()
    SetNodeId(value *string)()
    SetSpdxId(value *string)()
    SetUrl(value *string)()
}
