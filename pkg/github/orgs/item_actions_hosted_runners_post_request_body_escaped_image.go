package orgs

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemActionsHostedRunnersPostRequestBody_image the image of runner. To list all available images, use `GET /actions/hosted-runners/images/github-owned` or `GET /actions/hosted-runners/images/partner`.
type ItemActionsHostedRunnersPostRequestBody_image struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The unique identifier of the runner image.
    id *string
    // The version of the runner image to deploy. This is relevant only for runners using custom images.
    version *string
}
// NewItemActionsHostedRunnersPostRequestBody_image instantiates a new ItemActionsHostedRunnersPostRequestBody_image and sets the default values.
func NewItemActionsHostedRunnersPostRequestBody_image()(*ItemActionsHostedRunnersPostRequestBody_image) {
    m := &ItemActionsHostedRunnersPostRequestBody_image{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateItemActionsHostedRunnersPostRequestBody_imageFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateItemActionsHostedRunnersPostRequestBody_imageFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemActionsHostedRunnersPostRequestBody_image(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *ItemActionsHostedRunnersPostRequestBody_image) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ItemActionsHostedRunnersPostRequestBody_image) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
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
// GetId gets the id property value. The unique identifier of the runner image.
// returns a *string when successful
func (m *ItemActionsHostedRunnersPostRequestBody_image) GetId()(*string) {
    return m.id
}
// GetVersion gets the version property value. The version of the runner image to deploy. This is relevant only for runners using custom images.
// returns a *string when successful
func (m *ItemActionsHostedRunnersPostRequestBody_image) GetVersion()(*string) {
    return m.version
}
// Serialize serializes information the current object
func (m *ItemActionsHostedRunnersPostRequestBody_image) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("id", m.GetId())
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
func (m *ItemActionsHostedRunnersPostRequestBody_image) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetId sets the id property value. The unique identifier of the runner image.
func (m *ItemActionsHostedRunnersPostRequestBody_image) SetId(value *string)() {
    m.id = value
}
// SetVersion sets the version property value. The version of the runner image to deploy. This is relevant only for runners using custom images.
func (m *ItemActionsHostedRunnersPostRequestBody_image) SetVersion(value *string)() {
    m.version = value
}
type ItemActionsHostedRunnersPostRequestBody_imageable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetId()(*string)
    GetVersion()(*string)
    SetId(value *string)()
    SetVersion(value *string)()
}
