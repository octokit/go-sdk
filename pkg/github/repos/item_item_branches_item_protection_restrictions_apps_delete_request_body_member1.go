package repos

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemItemBranchesItemProtectionRestrictionsAppsDeleteRequestBodyMember1 
type ItemItemBranchesItemProtectionRestrictionsAppsDeleteRequestBodyMember1 struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The GitHub Apps that have push access to this branch. Use the slugified version of the app name. **Note**: The list of users, apps, and teams in total is limited to 100 items.
    apps []string
}
// NewItemItemBranchesItemProtectionRestrictionsAppsDeleteRequestBodyMember1 instantiates a new ItemItemBranchesItemProtectionRestrictionsAppsDeleteRequestBodyMember1 and sets the default values.
func NewItemItemBranchesItemProtectionRestrictionsAppsDeleteRequestBodyMember1()(*ItemItemBranchesItemProtectionRestrictionsAppsDeleteRequestBodyMember1) {
    m := &ItemItemBranchesItemProtectionRestrictionsAppsDeleteRequestBodyMember1{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateItemItemBranchesItemProtectionRestrictionsAppsDeleteRequestBodyMember1FromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemItemBranchesItemProtectionRestrictionsAppsDeleteRequestBodyMember1FromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemItemBranchesItemProtectionRestrictionsAppsDeleteRequestBodyMember1(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ItemItemBranchesItemProtectionRestrictionsAppsDeleteRequestBodyMember1) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetApps gets the apps property value. The GitHub Apps that have push access to this branch. Use the slugified version of the app name. **Note**: The list of users, apps, and teams in total is limited to 100 items.
func (m *ItemItemBranchesItemProtectionRestrictionsAppsDeleteRequestBodyMember1) GetApps()([]string) {
    return m.apps
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ItemItemBranchesItemProtectionRestrictionsAppsDeleteRequestBodyMember1) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["apps"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
            m.SetApps(res)
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *ItemItemBranchesItemProtectionRestrictionsAppsDeleteRequestBodyMember1) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetApps() != nil {
        err := writer.WriteCollectionOfStringValues("apps", m.GetApps())
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
func (m *ItemItemBranchesItemProtectionRestrictionsAppsDeleteRequestBodyMember1) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetApps sets the apps property value. The GitHub Apps that have push access to this branch. Use the slugified version of the app name. **Note**: The list of users, apps, and teams in total is limited to 100 items.
func (m *ItemItemBranchesItemProtectionRestrictionsAppsDeleteRequestBodyMember1) SetApps(value []string)() {
    m.apps = value
}
// ItemItemBranchesItemProtectionRestrictionsAppsDeleteRequestBodyMember1able 
type ItemItemBranchesItemProtectionRestrictionsAppsDeleteRequestBodyMember1able interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetApps()([]string)
    SetApps(value []string)()
}
