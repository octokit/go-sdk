package orgs

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type ItemActionsRunnerGroupsItemWithRunner_group_PatchRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Whether the runner group can be used by `public` repositories.
    allows_public_repositories *bool
    // Name of the runner group.
    name *string
    // If `true`, the runner group will be restricted to running only the workflows specified in the `selected_workflows` array.
    restricted_to_workflows *bool
    // List of workflows the runner group should be allowed to run. This setting will be ignored unless `restricted_to_workflows` is set to `true`.
    selected_workflows []string
}
// NewItemActionsRunnerGroupsItemWithRunner_group_PatchRequestBody instantiates a new ItemActionsRunnerGroupsItemWithRunner_group_PatchRequestBody and sets the default values.
func NewItemActionsRunnerGroupsItemWithRunner_group_PatchRequestBody()(*ItemActionsRunnerGroupsItemWithRunner_group_PatchRequestBody) {
    m := &ItemActionsRunnerGroupsItemWithRunner_group_PatchRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateItemActionsRunnerGroupsItemWithRunner_group_PatchRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateItemActionsRunnerGroupsItemWithRunner_group_PatchRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemActionsRunnerGroupsItemWithRunner_group_PatchRequestBody(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *ItemActionsRunnerGroupsItemWithRunner_group_PatchRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetAllowsPublicRepositories gets the allows_public_repositories property value. Whether the runner group can be used by `public` repositories.
// returns a *bool when successful
func (m *ItemActionsRunnerGroupsItemWithRunner_group_PatchRequestBody) GetAllowsPublicRepositories()(*bool) {
    return m.allows_public_repositories
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ItemActionsRunnerGroupsItemWithRunner_group_PatchRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["allows_public_repositories"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAllowsPublicRepositories(val)
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
    res["restricted_to_workflows"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRestrictedToWorkflows(val)
        }
        return nil
    }
    res["selected_workflows"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
            m.SetSelectedWorkflows(res)
        }
        return nil
    }
    return res
}
// GetName gets the name property value. Name of the runner group.
// returns a *string when successful
func (m *ItemActionsRunnerGroupsItemWithRunner_group_PatchRequestBody) GetName()(*string) {
    return m.name
}
// GetRestrictedToWorkflows gets the restricted_to_workflows property value. If `true`, the runner group will be restricted to running only the workflows specified in the `selected_workflows` array.
// returns a *bool when successful
func (m *ItemActionsRunnerGroupsItemWithRunner_group_PatchRequestBody) GetRestrictedToWorkflows()(*bool) {
    return m.restricted_to_workflows
}
// GetSelectedWorkflows gets the selected_workflows property value. List of workflows the runner group should be allowed to run. This setting will be ignored unless `restricted_to_workflows` is set to `true`.
// returns a []string when successful
func (m *ItemActionsRunnerGroupsItemWithRunner_group_PatchRequestBody) GetSelectedWorkflows()([]string) {
    return m.selected_workflows
}
// Serialize serializes information the current object
func (m *ItemActionsRunnerGroupsItemWithRunner_group_PatchRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteBoolValue("allows_public_repositories", m.GetAllowsPublicRepositories())
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
        err := writer.WriteBoolValue("restricted_to_workflows", m.GetRestrictedToWorkflows())
        if err != nil {
            return err
        }
    }
    if m.GetSelectedWorkflows() != nil {
        err := writer.WriteCollectionOfStringValues("selected_workflows", m.GetSelectedWorkflows())
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
func (m *ItemActionsRunnerGroupsItemWithRunner_group_PatchRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetAllowsPublicRepositories sets the allows_public_repositories property value. Whether the runner group can be used by `public` repositories.
func (m *ItemActionsRunnerGroupsItemWithRunner_group_PatchRequestBody) SetAllowsPublicRepositories(value *bool)() {
    m.allows_public_repositories = value
}
// SetName sets the name property value. Name of the runner group.
func (m *ItemActionsRunnerGroupsItemWithRunner_group_PatchRequestBody) SetName(value *string)() {
    m.name = value
}
// SetRestrictedToWorkflows sets the restricted_to_workflows property value. If `true`, the runner group will be restricted to running only the workflows specified in the `selected_workflows` array.
func (m *ItemActionsRunnerGroupsItemWithRunner_group_PatchRequestBody) SetRestrictedToWorkflows(value *bool)() {
    m.restricted_to_workflows = value
}
// SetSelectedWorkflows sets the selected_workflows property value. List of workflows the runner group should be allowed to run. This setting will be ignored unless `restricted_to_workflows` is set to `true`.
func (m *ItemActionsRunnerGroupsItemWithRunner_group_PatchRequestBody) SetSelectedWorkflows(value []string)() {
    m.selected_workflows = value
}
type ItemActionsRunnerGroupsItemWithRunner_group_PatchRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAllowsPublicRepositories()(*bool)
    GetName()(*string)
    GetRestrictedToWorkflows()(*bool)
    GetSelectedWorkflows()([]string)
    SetAllowsPublicRepositories(value *bool)()
    SetName(value *string)()
    SetRestrictedToWorkflows(value *bool)()
    SetSelectedWorkflows(value []string)()
}
