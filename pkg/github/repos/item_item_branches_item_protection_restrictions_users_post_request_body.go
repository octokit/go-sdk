package repos

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type ItemItemBranchesItemProtectionRestrictionsUsersPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The username for users
    users []string
}
// NewItemItemBranchesItemProtectionRestrictionsUsersPostRequestBody instantiates a new ItemItemBranchesItemProtectionRestrictionsUsersPostRequestBody and sets the default values.
func NewItemItemBranchesItemProtectionRestrictionsUsersPostRequestBody()(*ItemItemBranchesItemProtectionRestrictionsUsersPostRequestBody) {
    m := &ItemItemBranchesItemProtectionRestrictionsUsersPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateItemItemBranchesItemProtectionRestrictionsUsersPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateItemItemBranchesItemProtectionRestrictionsUsersPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemItemBranchesItemProtectionRestrictionsUsersPostRequestBody(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *ItemItemBranchesItemProtectionRestrictionsUsersPostRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ItemItemBranchesItemProtectionRestrictionsUsersPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["users"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
            m.SetUsers(res)
        }
        return nil
    }
    return res
}
// GetUsers gets the users property value. The username for users
// returns a []string when successful
func (m *ItemItemBranchesItemProtectionRestrictionsUsersPostRequestBody) GetUsers()([]string) {
    return m.users
}
// Serialize serializes information the current object
func (m *ItemItemBranchesItemProtectionRestrictionsUsersPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetUsers() != nil {
        err := writer.WriteCollectionOfStringValues("users", m.GetUsers())
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
func (m *ItemItemBranchesItemProtectionRestrictionsUsersPostRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetUsers sets the users property value. The username for users
func (m *ItemItemBranchesItemProtectionRestrictionsUsersPostRequestBody) SetUsers(value []string)() {
    m.users = value
}
type ItemItemBranchesItemProtectionRestrictionsUsersPostRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetUsers()([]string)
    SetUsers(value []string)()
}
