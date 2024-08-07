package orgs

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type ItemCodeSecurityConfigurationsDetachDeleteRequestBody struct {
    // An array of repository IDs to detach from configurations.
    selected_repository_ids []int32
}
// NewItemCodeSecurityConfigurationsDetachDeleteRequestBody instantiates a new ItemCodeSecurityConfigurationsDetachDeleteRequestBody and sets the default values.
func NewItemCodeSecurityConfigurationsDetachDeleteRequestBody()(*ItemCodeSecurityConfigurationsDetachDeleteRequestBody) {
    m := &ItemCodeSecurityConfigurationsDetachDeleteRequestBody{
    }
    return m
}
// CreateItemCodeSecurityConfigurationsDetachDeleteRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateItemCodeSecurityConfigurationsDetachDeleteRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemCodeSecurityConfigurationsDetachDeleteRequestBody(), nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ItemCodeSecurityConfigurationsDetachDeleteRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
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
    return res
}
// GetSelectedRepositoryIds gets the selected_repository_ids property value. An array of repository IDs to detach from configurations.
// returns a []int32 when successful
func (m *ItemCodeSecurityConfigurationsDetachDeleteRequestBody) GetSelectedRepositoryIds()([]int32) {
    return m.selected_repository_ids
}
// Serialize serializes information the current object
func (m *ItemCodeSecurityConfigurationsDetachDeleteRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetSelectedRepositoryIds() != nil {
        err := writer.WriteCollectionOfInt32Values("selected_repository_ids", m.GetSelectedRepositoryIds())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetSelectedRepositoryIds sets the selected_repository_ids property value. An array of repository IDs to detach from configurations.
func (m *ItemCodeSecurityConfigurationsDetachDeleteRequestBody) SetSelectedRepositoryIds(value []int32)() {
    m.selected_repository_ids = value
}
type ItemCodeSecurityConfigurationsDetachDeleteRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetSelectedRepositoryIds()([]int32)
    SetSelectedRepositoryIds(value []int32)()
}
