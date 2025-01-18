package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// NullableActionsHostedRunnerPoolImage provides details of a hosted runner image
type NullableActionsHostedRunnerPoolImage struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Display name for this image.
    display_name *string
    // The ID of the image. Use this ID for the `image` parameter when creating a new larger runner.
    id *string
    // Image size in GB.
    size_gb *int32
    // The image provider.
    source *NullableActionsHostedRunnerPoolImage_source
    // The image version of the hosted runner pool.
    version *string
}
// NewNullableActionsHostedRunnerPoolImage instantiates a new NullableActionsHostedRunnerPoolImage and sets the default values.
func NewNullableActionsHostedRunnerPoolImage()(*NullableActionsHostedRunnerPoolImage) {
    m := &NullableActionsHostedRunnerPoolImage{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateNullableActionsHostedRunnerPoolImageFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateNullableActionsHostedRunnerPoolImageFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewNullableActionsHostedRunnerPoolImage(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *NullableActionsHostedRunnerPoolImage) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetDisplayName gets the display_name property value. Display name for this image.
// returns a *string when successful
func (m *NullableActionsHostedRunnerPoolImage) GetDisplayName()(*string) {
    return m.display_name
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *NullableActionsHostedRunnerPoolImage) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["display_name"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisplayName(val)
        }
        return nil
    }
    res["id"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetId(val)
        }
        return nil
    }
    res["size_gb"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSizeGb(val)
        }
        return nil
    }
    res["source"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseNullableActionsHostedRunnerPoolImage_source)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSource(val.(*NullableActionsHostedRunnerPoolImage_source))
        }
        return nil
    }
    res["version"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetVersion(val)
        }
        return nil
    }
    return res
}
// GetId gets the id property value. The ID of the image. Use this ID for the `image` parameter when creating a new larger runner.
// returns a *string when successful
func (m *NullableActionsHostedRunnerPoolImage) GetId()(*string) {
    return m.id
}
// GetSizeGb gets the size_gb property value. Image size in GB.
// returns a *int32 when successful
func (m *NullableActionsHostedRunnerPoolImage) GetSizeGb()(*int32) {
    return m.size_gb
}
// GetSource gets the source property value. The image provider.
// returns a *NullableActionsHostedRunnerPoolImage_source when successful
func (m *NullableActionsHostedRunnerPoolImage) GetSource()(*NullableActionsHostedRunnerPoolImage_source) {
    return m.source
}
// GetVersion gets the version property value. The image version of the hosted runner pool.
// returns a *string when successful
func (m *NullableActionsHostedRunnerPoolImage) GetVersion()(*string) {
    return m.version
}
// Serialize serializes information the current object
func (m *NullableActionsHostedRunnerPoolImage) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("display_name", m.GetDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("id", m.GetId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("size_gb", m.GetSizeGb())
        if err != nil {
            return err
        }
    }
    if m.GetSource() != nil {
        cast := (*m.GetSource()).String()
        err := writer.WriteStringValue("source", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("version", m.GetVersion())
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
func (m *NullableActionsHostedRunnerPoolImage) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetDisplayName sets the display_name property value. Display name for this image.
func (m *NullableActionsHostedRunnerPoolImage) SetDisplayName(value *string)() {
    m.display_name = value
}
// SetId sets the id property value. The ID of the image. Use this ID for the `image` parameter when creating a new larger runner.
func (m *NullableActionsHostedRunnerPoolImage) SetId(value *string)() {
    m.id = value
}
// SetSizeGb sets the size_gb property value. Image size in GB.
func (m *NullableActionsHostedRunnerPoolImage) SetSizeGb(value *int32)() {
    m.size_gb = value
}
// SetSource sets the source property value. The image provider.
func (m *NullableActionsHostedRunnerPoolImage) SetSource(value *NullableActionsHostedRunnerPoolImage_source)() {
    m.source = value
}
// SetVersion sets the version property value. The image version of the hosted runner pool.
func (m *NullableActionsHostedRunnerPoolImage) SetVersion(value *string)() {
    m.version = value
}
type NullableActionsHostedRunnerPoolImageable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetDisplayName()(*string)
    GetId()(*string)
    GetSizeGb()(*int32)
    GetSource()(*NullableActionsHostedRunnerPoolImage_source)
    GetVersion()(*string)
    SetDisplayName(value *string)()
    SetId(value *string)()
    SetSizeGb(value *int32)()
    SetSource(value *NullableActionsHostedRunnerPoolImage_source)()
    SetVersion(value *string)()
}
