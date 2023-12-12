package repos

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemItemBranchesItemProtectionPutRequestBody_required_pull_request_reviews require at least one approving review on a pull request, before merging. Set to `null` to disable.
type ItemItemBranchesItemProtectionPutRequestBody_required_pull_request_reviews struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Allow specific users, teams, or apps to bypass pull request requirements.
    bypass_pull_request_allowances ItemItemBranchesItemProtectionPutRequestBody_required_pull_request_reviews_bypass_pull_request_allowancesable
    // Set to `true` if you want to automatically dismiss approving reviews when someone pushes a new commit.
    dismiss_stale_reviews *bool
    // Specify which users, teams, and apps can dismiss pull request reviews. Pass an empty `dismissal_restrictions` object to disable. User and team `dismissal_restrictions` are only available for organization-owned repositories. Omit this parameter for personal repositories.
    dismissal_restrictions ItemItemBranchesItemProtectionPutRequestBody_required_pull_request_reviews_dismissal_restrictionsable
    // Blocks merging pull requests until [code owners](https://docs.github.com/articles/about-code-owners/) review them.
    require_code_owner_reviews *bool
    // Whether the most recent push must be approved by someone other than the person who pushed it. Default: `false`.
    require_last_push_approval *bool
    // Specify the number of reviewers required to approve pull requests. Use a number between 1 and 6 or 0 to not require reviewers.
    required_approving_review_count *int32
}
// NewItemItemBranchesItemProtectionPutRequestBody_required_pull_request_reviews instantiates a new ItemItemBranchesItemProtectionPutRequestBody_required_pull_request_reviews and sets the default values.
func NewItemItemBranchesItemProtectionPutRequestBody_required_pull_request_reviews()(*ItemItemBranchesItemProtectionPutRequestBody_required_pull_request_reviews) {
    m := &ItemItemBranchesItemProtectionPutRequestBody_required_pull_request_reviews{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateItemItemBranchesItemProtectionPutRequestBody_required_pull_request_reviewsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemItemBranchesItemProtectionPutRequestBody_required_pull_request_reviewsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemItemBranchesItemProtectionPutRequestBody_required_pull_request_reviews(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ItemItemBranchesItemProtectionPutRequestBody_required_pull_request_reviews) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetBypassPullRequestAllowances gets the bypass_pull_request_allowances property value. Allow specific users, teams, or apps to bypass pull request requirements.
func (m *ItemItemBranchesItemProtectionPutRequestBody_required_pull_request_reviews) GetBypassPullRequestAllowances()(ItemItemBranchesItemProtectionPutRequestBody_required_pull_request_reviews_bypass_pull_request_allowancesable) {
    return m.bypass_pull_request_allowances
}
// GetDismissalRestrictions gets the dismissal_restrictions property value. Specify which users, teams, and apps can dismiss pull request reviews. Pass an empty `dismissal_restrictions` object to disable. User and team `dismissal_restrictions` are only available for organization-owned repositories. Omit this parameter for personal repositories.
func (m *ItemItemBranchesItemProtectionPutRequestBody_required_pull_request_reviews) GetDismissalRestrictions()(ItemItemBranchesItemProtectionPutRequestBody_required_pull_request_reviews_dismissal_restrictionsable) {
    return m.dismissal_restrictions
}
// GetDismissStaleReviews gets the dismiss_stale_reviews property value. Set to `true` if you want to automatically dismiss approving reviews when someone pushes a new commit.
func (m *ItemItemBranchesItemProtectionPutRequestBody_required_pull_request_reviews) GetDismissStaleReviews()(*bool) {
    return m.dismiss_stale_reviews
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ItemItemBranchesItemProtectionPutRequestBody_required_pull_request_reviews) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["bypass_pull_request_allowances"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateItemItemBranchesItemProtectionPutRequestBody_required_pull_request_reviews_bypass_pull_request_allowancesFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBypassPullRequestAllowances(val.(ItemItemBranchesItemProtectionPutRequestBody_required_pull_request_reviews_bypass_pull_request_allowancesable))
        }
        return nil
    }
    res["dismiss_stale_reviews"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDismissStaleReviews(val)
        }
        return nil
    }
    res["dismissal_restrictions"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateItemItemBranchesItemProtectionPutRequestBody_required_pull_request_reviews_dismissal_restrictionsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDismissalRestrictions(val.(ItemItemBranchesItemProtectionPutRequestBody_required_pull_request_reviews_dismissal_restrictionsable))
        }
        return nil
    }
    res["require_code_owner_reviews"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRequireCodeOwnerReviews(val)
        }
        return nil
    }
    res["require_last_push_approval"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRequireLastPushApproval(val)
        }
        return nil
    }
    res["required_approving_review_count"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRequiredApprovingReviewCount(val)
        }
        return nil
    }
    return res
}
// GetRequireCodeOwnerReviews gets the require_code_owner_reviews property value. Blocks merging pull requests until [code owners](https://docs.github.com/articles/about-code-owners/) review them.
func (m *ItemItemBranchesItemProtectionPutRequestBody_required_pull_request_reviews) GetRequireCodeOwnerReviews()(*bool) {
    return m.require_code_owner_reviews
}
// GetRequiredApprovingReviewCount gets the required_approving_review_count property value. Specify the number of reviewers required to approve pull requests. Use a number between 1 and 6 or 0 to not require reviewers.
func (m *ItemItemBranchesItemProtectionPutRequestBody_required_pull_request_reviews) GetRequiredApprovingReviewCount()(*int32) {
    return m.required_approving_review_count
}
// GetRequireLastPushApproval gets the require_last_push_approval property value. Whether the most recent push must be approved by someone other than the person who pushed it. Default: `false`.
func (m *ItemItemBranchesItemProtectionPutRequestBody_required_pull_request_reviews) GetRequireLastPushApproval()(*bool) {
    return m.require_last_push_approval
}
// Serialize serializes information the current object
func (m *ItemItemBranchesItemProtectionPutRequestBody_required_pull_request_reviews) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("bypass_pull_request_allowances", m.GetBypassPullRequestAllowances())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("dismissal_restrictions", m.GetDismissalRestrictions())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("dismiss_stale_reviews", m.GetDismissStaleReviews())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("required_approving_review_count", m.GetRequiredApprovingReviewCount())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("require_code_owner_reviews", m.GetRequireCodeOwnerReviews())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("require_last_push_approval", m.GetRequireLastPushApproval())
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
func (m *ItemItemBranchesItemProtectionPutRequestBody_required_pull_request_reviews) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetBypassPullRequestAllowances sets the bypass_pull_request_allowances property value. Allow specific users, teams, or apps to bypass pull request requirements.
func (m *ItemItemBranchesItemProtectionPutRequestBody_required_pull_request_reviews) SetBypassPullRequestAllowances(value ItemItemBranchesItemProtectionPutRequestBody_required_pull_request_reviews_bypass_pull_request_allowancesable)() {
    m.bypass_pull_request_allowances = value
}
// SetDismissalRestrictions sets the dismissal_restrictions property value. Specify which users, teams, and apps can dismiss pull request reviews. Pass an empty `dismissal_restrictions` object to disable. User and team `dismissal_restrictions` are only available for organization-owned repositories. Omit this parameter for personal repositories.
func (m *ItemItemBranchesItemProtectionPutRequestBody_required_pull_request_reviews) SetDismissalRestrictions(value ItemItemBranchesItemProtectionPutRequestBody_required_pull_request_reviews_dismissal_restrictionsable)() {
    m.dismissal_restrictions = value
}
// SetDismissStaleReviews sets the dismiss_stale_reviews property value. Set to `true` if you want to automatically dismiss approving reviews when someone pushes a new commit.
func (m *ItemItemBranchesItemProtectionPutRequestBody_required_pull_request_reviews) SetDismissStaleReviews(value *bool)() {
    m.dismiss_stale_reviews = value
}
// SetRequireCodeOwnerReviews sets the require_code_owner_reviews property value. Blocks merging pull requests until [code owners](https://docs.github.com/articles/about-code-owners/) review them.
func (m *ItemItemBranchesItemProtectionPutRequestBody_required_pull_request_reviews) SetRequireCodeOwnerReviews(value *bool)() {
    m.require_code_owner_reviews = value
}
// SetRequiredApprovingReviewCount sets the required_approving_review_count property value. Specify the number of reviewers required to approve pull requests. Use a number between 1 and 6 or 0 to not require reviewers.
func (m *ItemItemBranchesItemProtectionPutRequestBody_required_pull_request_reviews) SetRequiredApprovingReviewCount(value *int32)() {
    m.required_approving_review_count = value
}
// SetRequireLastPushApproval sets the require_last_push_approval property value. Whether the most recent push must be approved by someone other than the person who pushed it. Default: `false`.
func (m *ItemItemBranchesItemProtectionPutRequestBody_required_pull_request_reviews) SetRequireLastPushApproval(value *bool)() {
    m.require_last_push_approval = value
}
// ItemItemBranchesItemProtectionPutRequestBody_required_pull_request_reviewsable 
type ItemItemBranchesItemProtectionPutRequestBody_required_pull_request_reviewsable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBypassPullRequestAllowances()(ItemItemBranchesItemProtectionPutRequestBody_required_pull_request_reviews_bypass_pull_request_allowancesable)
    GetDismissalRestrictions()(ItemItemBranchesItemProtectionPutRequestBody_required_pull_request_reviews_dismissal_restrictionsable)
    GetDismissStaleReviews()(*bool)
    GetRequireCodeOwnerReviews()(*bool)
    GetRequiredApprovingReviewCount()(*int32)
    GetRequireLastPushApproval()(*bool)
    SetBypassPullRequestAllowances(value ItemItemBranchesItemProtectionPutRequestBody_required_pull_request_reviews_bypass_pull_request_allowancesable)()
    SetDismissalRestrictions(value ItemItemBranchesItemProtectionPutRequestBody_required_pull_request_reviews_dismissal_restrictionsable)()
    SetDismissStaleReviews(value *bool)()
    SetRequireCodeOwnerReviews(value *bool)()
    SetRequiredApprovingReviewCount(value *int32)()
    SetRequireLastPushApproval(value *bool)()
}
