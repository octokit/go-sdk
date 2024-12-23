package orgs

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type ItemPrivateRegistriesItemWithSecret_namePatchRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The value for your secret, encrypted with [LibSodium](https://libsodium.gitbook.io/doc/bindings_for_other_languages) using the public key retrieved from the [Get private registries public key for an organization](https://docs.github.com/rest/private-registries/organization-configurations#get-private-registries-public-key-for-an-organization) endpoint.
    encrypted_value *string
    // The ID of the key you used to encrypt the secret.
    key_id *string
    // An array of repository IDs that can access the organization private registry. You can only provide a list of repository IDs when `visibility` is set to `selected`. This field should be omitted if `visibility` is set to `all` or `private`.
    selected_repository_ids []int32
    // The username to use when authenticating with the private registry. This field should be omitted if the private registry does not require a username for authentication.
    username *string
}
// NewItemPrivateRegistriesItemWithSecret_namePatchRequestBody instantiates a new ItemPrivateRegistriesItemWithSecret_namePatchRequestBody and sets the default values.
func NewItemPrivateRegistriesItemWithSecret_namePatchRequestBody()(*ItemPrivateRegistriesItemWithSecret_namePatchRequestBody) {
    m := &ItemPrivateRegistriesItemWithSecret_namePatchRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateItemPrivateRegistriesItemWithSecret_namePatchRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateItemPrivateRegistriesItemWithSecret_namePatchRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemPrivateRegistriesItemWithSecret_namePatchRequestBody(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *ItemPrivateRegistriesItemWithSecret_namePatchRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetEncryptedValue gets the encrypted_value property value. The value for your secret, encrypted with [LibSodium](https://libsodium.gitbook.io/doc/bindings_for_other_languages) using the public key retrieved from the [Get private registries public key for an organization](https://docs.github.com/rest/private-registries/organization-configurations#get-private-registries-public-key-for-an-organization) endpoint.
// returns a *string when successful
func (m *ItemPrivateRegistriesItemWithSecret_namePatchRequestBody) GetEncryptedValue()(*string) {
    return m.encrypted_value
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ItemPrivateRegistriesItemWithSecret_namePatchRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["encrypted_value"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEncryptedValue(val)
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
    res["selected_repository_ids"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("int32")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]int32, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = *(v.(*int32))
                }
            }
            m.SetSelectedRepositoryIds(res)
        }
        return nil
    }
    res["username"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUsername(val)
        }
        return nil
    }
    return res
}
// GetKeyId gets the key_id property value. The ID of the key you used to encrypt the secret.
// returns a *string when successful
func (m *ItemPrivateRegistriesItemWithSecret_namePatchRequestBody) GetKeyId()(*string) {
    return m.key_id
}
// GetSelectedRepositoryIds gets the selected_repository_ids property value. An array of repository IDs that can access the organization private registry. You can only provide a list of repository IDs when `visibility` is set to `selected`. This field should be omitted if `visibility` is set to `all` or `private`.
// returns a []int32 when successful
func (m *ItemPrivateRegistriesItemWithSecret_namePatchRequestBody) GetSelectedRepositoryIds()([]int32) {
    return m.selected_repository_ids
}
// GetUsername gets the username property value. The username to use when authenticating with the private registry. This field should be omitted if the private registry does not require a username for authentication.
// returns a *string when successful
func (m *ItemPrivateRegistriesItemWithSecret_namePatchRequestBody) GetUsername()(*string) {
    return m.username
}
// Serialize serializes information the current object
func (m *ItemPrivateRegistriesItemWithSecret_namePatchRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("encrypted_value", m.GetEncryptedValue())
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
    if m.GetSelectedRepositoryIds() != nil {
        err := writer.WriteCollectionOfInt32Values("selected_repository_ids", m.GetSelectedRepositoryIds())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("username", m.GetUsername())
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
func (m *ItemPrivateRegistriesItemWithSecret_namePatchRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetEncryptedValue sets the encrypted_value property value. The value for your secret, encrypted with [LibSodium](https://libsodium.gitbook.io/doc/bindings_for_other_languages) using the public key retrieved from the [Get private registries public key for an organization](https://docs.github.com/rest/private-registries/organization-configurations#get-private-registries-public-key-for-an-organization) endpoint.
func (m *ItemPrivateRegistriesItemWithSecret_namePatchRequestBody) SetEncryptedValue(value *string)() {
    m.encrypted_value = value
}
// SetKeyId sets the key_id property value. The ID of the key you used to encrypt the secret.
func (m *ItemPrivateRegistriesItemWithSecret_namePatchRequestBody) SetKeyId(value *string)() {
    m.key_id = value
}
// SetSelectedRepositoryIds sets the selected_repository_ids property value. An array of repository IDs that can access the organization private registry. You can only provide a list of repository IDs when `visibility` is set to `selected`. This field should be omitted if `visibility` is set to `all` or `private`.
func (m *ItemPrivateRegistriesItemWithSecret_namePatchRequestBody) SetSelectedRepositoryIds(value []int32)() {
    m.selected_repository_ids = value
}
// SetUsername sets the username property value. The username to use when authenticating with the private registry. This field should be omitted if the private registry does not require a username for authentication.
func (m *ItemPrivateRegistriesItemWithSecret_namePatchRequestBody) SetUsername(value *string)() {
    m.username = value
}
type ItemPrivateRegistriesItemWithSecret_namePatchRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetEncryptedValue()(*string)
    GetKeyId()(*string)
    GetSelectedRepositoryIds()([]int32)
    GetUsername()(*string)
    SetEncryptedValue(value *string)()
    SetKeyId(value *string)()
    SetSelectedRepositoryIds(value []int32)()
    SetUsername(value *string)()
}
