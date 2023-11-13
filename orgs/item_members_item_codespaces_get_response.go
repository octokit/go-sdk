package orgs

import (
    i33bcfa37677dfd8b71eaf22ad3bfd140ba1702e7a8418558499c22418bf2fef2 "github.com/octokit/go-sdk/models"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemMembersItemCodespacesGetResponse 
type ItemMembersItemCodespacesGetResponse struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The codespaces property
    codespaces []i33bcfa37677dfd8b71eaf22ad3bfd140ba1702e7a8418558499c22418bf2fef2.Codespaceable
    // The total_count property
    total_count *int32
}
// NewItemMembersItemCodespacesGetResponse instantiates a new ItemMembersItemCodespacesGetResponse and sets the default values.
func NewItemMembersItemCodespacesGetResponse()(*ItemMembersItemCodespacesGetResponse) {
    m := &ItemMembersItemCodespacesGetResponse{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateItemMembersItemCodespacesGetResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemMembersItemCodespacesGetResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemMembersItemCodespacesGetResponse(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ItemMembersItemCodespacesGetResponse) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetCodespaces gets the codespaces property value. The codespaces property
func (m *ItemMembersItemCodespacesGetResponse) GetCodespaces()([]i33bcfa37677dfd8b71eaf22ad3bfd140ba1702e7a8418558499c22418bf2fef2.Codespaceable) {
    return m.codespaces
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ItemMembersItemCodespacesGetResponse) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["codespaces"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(i33bcfa37677dfd8b71eaf22ad3bfd140ba1702e7a8418558499c22418bf2fef2.CreateCodespaceFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]i33bcfa37677dfd8b71eaf22ad3bfd140ba1702e7a8418558499c22418bf2fef2.Codespaceable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(i33bcfa37677dfd8b71eaf22ad3bfd140ba1702e7a8418558499c22418bf2fef2.Codespaceable)
                }
            }
            m.SetCodespaces(res)
        }
        return nil
    }
    res["total_count"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTotalCount(val)
        }
        return nil
    }
    return res
}
// GetTotalCount gets the total_count property value. The total_count property
func (m *ItemMembersItemCodespacesGetResponse) GetTotalCount()(*int32) {
    return m.total_count
}
// Serialize serializes information the current object
func (m *ItemMembersItemCodespacesGetResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetCodespaces() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetCodespaces()))
        for i, v := range m.GetCodespaces() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("codespaces", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("total_count", m.GetTotalCount())
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
func (m *ItemMembersItemCodespacesGetResponse) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetCodespaces sets the codespaces property value. The codespaces property
func (m *ItemMembersItemCodespacesGetResponse) SetCodespaces(value []i33bcfa37677dfd8b71eaf22ad3bfd140ba1702e7a8418558499c22418bf2fef2.Codespaceable)() {
    m.codespaces = value
}
// SetTotalCount sets the total_count property value. The total_count property
func (m *ItemMembersItemCodespacesGetResponse) SetTotalCount(value *int32)() {
    m.total_count = value
}
// ItemMembersItemCodespacesGetResponseable 
type ItemMembersItemCodespacesGetResponseable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCodespaces()([]i33bcfa37677dfd8b71eaf22ad3bfd140ba1702e7a8418558499c22418bf2fef2.Codespaceable)
    GetTotalCount()(*int32)
    SetCodespaces(value []i33bcfa37677dfd8b71eaf22ad3bfd140ba1702e7a8418558499c22418bf2fef2.Codespaceable)()
    SetTotalCount(value *int32)()
}
