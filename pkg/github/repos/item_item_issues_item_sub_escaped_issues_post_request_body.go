package repos

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type ItemItemIssuesItemSub_issuesPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Option that, when true, instructs the operation to replace the sub-issues current parent issue
    replace_parent *bool
    // The sub-issue to add
    sub_issue_id *int32
}
// NewItemItemIssuesItemSub_issuesPostRequestBody instantiates a new ItemItemIssuesItemSub_issuesPostRequestBody and sets the default values.
func NewItemItemIssuesItemSub_issuesPostRequestBody()(*ItemItemIssuesItemSub_issuesPostRequestBody) {
    m := &ItemItemIssuesItemSub_issuesPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateItemItemIssuesItemSub_issuesPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateItemItemIssuesItemSub_issuesPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemItemIssuesItemSub_issuesPostRequestBody(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *ItemItemIssuesItemSub_issuesPostRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ItemItemIssuesItemSub_issuesPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["replace_parent"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetReplaceParent(val)
        }
        return nil
    }
    res["sub_issue_id"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSubIssueId(val)
        }
        return nil
    }
    return res
}
// GetReplaceParent gets the replace_parent property value. Option that, when true, instructs the operation to replace the sub-issues current parent issue
// returns a *bool when successful
func (m *ItemItemIssuesItemSub_issuesPostRequestBody) GetReplaceParent()(*bool) {
    return m.replace_parent
}
// GetSubIssueId gets the sub_issue_id property value. The sub-issue to add
// returns a *int32 when successful
func (m *ItemItemIssuesItemSub_issuesPostRequestBody) GetSubIssueId()(*int32) {
    return m.sub_issue_id
}
// Serialize serializes information the current object
func (m *ItemItemIssuesItemSub_issuesPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteBoolValue("replace_parent", m.GetReplaceParent())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("sub_issue_id", m.GetSubIssueId())
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
func (m *ItemItemIssuesItemSub_issuesPostRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetReplaceParent sets the replace_parent property value. Option that, when true, instructs the operation to replace the sub-issues current parent issue
func (m *ItemItemIssuesItemSub_issuesPostRequestBody) SetReplaceParent(value *bool)() {
    m.replace_parent = value
}
// SetSubIssueId sets the sub_issue_id property value. The sub-issue to add
func (m *ItemItemIssuesItemSub_issuesPostRequestBody) SetSubIssueId(value *int32)() {
    m.sub_issue_id = value
}
type ItemItemIssuesItemSub_issuesPostRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetReplaceParent()(*bool)
    GetSubIssueId()(*int32)
    SetReplaceParent(value *bool)()
    SetSubIssueId(value *int32)()
}
