package orgs

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type ItemPrivateRegistriesPublicKeyGetResponse struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The Base64 encoded public key.
    key *string
    // The identifier for the key.
    key_id *string
}
// NewItemPrivateRegistriesPublicKeyGetResponse instantiates a new ItemPrivateRegistriesPublicKeyGetResponse and sets the default values.
func NewItemPrivateRegistriesPublicKeyGetResponse()(*ItemPrivateRegistriesPublicKeyGetResponse) {
    m := &ItemPrivateRegistriesPublicKeyGetResponse{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateItemPrivateRegistriesPublicKeyGetResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateItemPrivateRegistriesPublicKeyGetResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemPrivateRegistriesPublicKeyGetResponse(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *ItemPrivateRegistriesPublicKeyGetResponse) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ItemPrivateRegistriesPublicKeyGetResponse) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
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
    return res
}
// GetKey gets the key property value. The Base64 encoded public key.
// returns a *string when successful
func (m *ItemPrivateRegistriesPublicKeyGetResponse) GetKey()(*string) {
    return m.key
}
// GetKeyId gets the key_id property value. The identifier for the key.
// returns a *string when successful
func (m *ItemPrivateRegistriesPublicKeyGetResponse) GetKeyId()(*string) {
    return m.key_id
}
// Serialize serializes information the current object
func (m *ItemPrivateRegistriesPublicKeyGetResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ItemPrivateRegistriesPublicKeyGetResponse) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetKey sets the key property value. The Base64 encoded public key.
func (m *ItemPrivateRegistriesPublicKeyGetResponse) SetKey(value *string)() {
    m.key = value
}
// SetKeyId sets the key_id property value. The identifier for the key.
func (m *ItemPrivateRegistriesPublicKeyGetResponse) SetKeyId(value *string)() {
    m.key_id = value
}
type ItemPrivateRegistriesPublicKeyGetResponseable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetKey()(*string)
    GetKeyId()(*string)
    SetKey(value *string)()
    SetKeyId(value *string)()
}
