package repos

import (
    i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e "github.com/octokit/go-sdk/github/octokit/models"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemItemActionsVariablesGetResponse 
type ItemItemActionsVariablesGetResponse struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The total_count property
    total_count *int32
    // The variables property
    variables []i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.ActionsVariableable
}
// NewItemItemActionsVariablesGetResponse instantiates a new ItemItemActionsVariablesGetResponse and sets the default values.
func NewItemItemActionsVariablesGetResponse()(*ItemItemActionsVariablesGetResponse) {
    m := &ItemItemActionsVariablesGetResponse{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateItemItemActionsVariablesGetResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemItemActionsVariablesGetResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemItemActionsVariablesGetResponse(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ItemItemActionsVariablesGetResponse) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ItemItemActionsVariablesGetResponse) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
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
    res["variables"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.CreateActionsVariableFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.ActionsVariableable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.ActionsVariableable)
                }
            }
            m.SetVariables(res)
        }
        return nil
    }
    return res
}
// GetTotalCount gets the total_count property value. The total_count property
func (m *ItemItemActionsVariablesGetResponse) GetTotalCount()(*int32) {
    return m.total_count
}
// GetVariables gets the variables property value. The variables property
func (m *ItemItemActionsVariablesGetResponse) GetVariables()([]i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.ActionsVariableable) {
    return m.variables
}
// Serialize serializes information the current object
func (m *ItemItemActionsVariablesGetResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteInt32Value("total_count", m.GetTotalCount())
        if err != nil {
            return err
        }
    }
    if m.GetVariables() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetVariables()))
        for i, v := range m.GetVariables() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("variables", cast)
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
func (m *ItemItemActionsVariablesGetResponse) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetTotalCount sets the total_count property value. The total_count property
func (m *ItemItemActionsVariablesGetResponse) SetTotalCount(value *int32)() {
    m.total_count = value
}
// SetVariables sets the variables property value. The variables property
func (m *ItemItemActionsVariablesGetResponse) SetVariables(value []i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.ActionsVariableable)() {
    m.variables = value
}
// ItemItemActionsVariablesGetResponseable 
type ItemItemActionsVariablesGetResponseable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetTotalCount()(*int32)
    GetVariables()([]i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.ActionsVariableable)
    SetTotalCount(value *int32)()
    SetVariables(value []i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.ActionsVariableable)()
}
