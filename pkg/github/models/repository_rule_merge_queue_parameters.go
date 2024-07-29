package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type RepositoryRuleMergeQueue_parameters struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Maximum time for a required status check to report a conclusion. After this much time has elapsed, checks that have not reported a conclusion will be assumed to have failed
    check_response_timeout_minutes *int32
    // When set to ALLGREEN, the merge commit created by merge queue for each PR in the group must pass all required checks to merge. When set to HEADGREEN, only the commit at the head of the merge group, i.e. the commit containing changes from all of the PRs in the group, must pass its required checks to merge.
    grouping_strategy *RepositoryRuleMergeQueue_parameters_grouping_strategy
    // Limit the number of queued pull requests requesting checks and workflow runs at the same time.
    max_entries_to_build *int32
    // The maximum number of PRs that will be merged together in a group.
    max_entries_to_merge *int32
    // Method to use when merging changes from queued pull requests.
    merge_method *RepositoryRuleMergeQueue_parameters_merge_method
    // The minimum number of PRs that will be merged together in a group.
    min_entries_to_merge *int32
    // The time merge queue should wait after the first PR is added to the queue for the minimum group size to be met. After this time has elapsed, the minimum group size will be ignored and a smaller group will be merged.
    min_entries_to_merge_wait_minutes *int32
}
// NewRepositoryRuleMergeQueue_parameters instantiates a new RepositoryRuleMergeQueue_parameters and sets the default values.
func NewRepositoryRuleMergeQueue_parameters()(*RepositoryRuleMergeQueue_parameters) {
    m := &RepositoryRuleMergeQueue_parameters{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateRepositoryRuleMergeQueue_parametersFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateRepositoryRuleMergeQueue_parametersFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewRepositoryRuleMergeQueue_parameters(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *RepositoryRuleMergeQueue_parameters) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetCheckResponseTimeoutMinutes gets the check_response_timeout_minutes property value. Maximum time for a required status check to report a conclusion. After this much time has elapsed, checks that have not reported a conclusion will be assumed to have failed
// returns a *int32 when successful
func (m *RepositoryRuleMergeQueue_parameters) GetCheckResponseTimeoutMinutes()(*int32) {
    return m.check_response_timeout_minutes
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *RepositoryRuleMergeQueue_parameters) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["check_response_timeout_minutes"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCheckResponseTimeoutMinutes(val)
        }
        return nil
    }
    res["grouping_strategy"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseRepositoryRuleMergeQueue_parameters_grouping_strategy)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetGroupingStrategy(val.(*RepositoryRuleMergeQueue_parameters_grouping_strategy))
        }
        return nil
    }
    res["max_entries_to_build"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMaxEntriesToBuild(val)
        }
        return nil
    }
    res["max_entries_to_merge"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMaxEntriesToMerge(val)
        }
        return nil
    }
    res["merge_method"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseRepositoryRuleMergeQueue_parameters_merge_method)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMergeMethod(val.(*RepositoryRuleMergeQueue_parameters_merge_method))
        }
        return nil
    }
    res["min_entries_to_merge"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMinEntriesToMerge(val)
        }
        return nil
    }
    res["min_entries_to_merge_wait_minutes"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMinEntriesToMergeWaitMinutes(val)
        }
        return nil
    }
    return res
}
// GetGroupingStrategy gets the grouping_strategy property value. When set to ALLGREEN, the merge commit created by merge queue for each PR in the group must pass all required checks to merge. When set to HEADGREEN, only the commit at the head of the merge group, i.e. the commit containing changes from all of the PRs in the group, must pass its required checks to merge.
// returns a *RepositoryRuleMergeQueue_parameters_grouping_strategy when successful
func (m *RepositoryRuleMergeQueue_parameters) GetGroupingStrategy()(*RepositoryRuleMergeQueue_parameters_grouping_strategy) {
    return m.grouping_strategy
}
// GetMaxEntriesToBuild gets the max_entries_to_build property value. Limit the number of queued pull requests requesting checks and workflow runs at the same time.
// returns a *int32 when successful
func (m *RepositoryRuleMergeQueue_parameters) GetMaxEntriesToBuild()(*int32) {
    return m.max_entries_to_build
}
// GetMaxEntriesToMerge gets the max_entries_to_merge property value. The maximum number of PRs that will be merged together in a group.
// returns a *int32 when successful
func (m *RepositoryRuleMergeQueue_parameters) GetMaxEntriesToMerge()(*int32) {
    return m.max_entries_to_merge
}
// GetMergeMethod gets the merge_method property value. Method to use when merging changes from queued pull requests.
// returns a *RepositoryRuleMergeQueue_parameters_merge_method when successful
func (m *RepositoryRuleMergeQueue_parameters) GetMergeMethod()(*RepositoryRuleMergeQueue_parameters_merge_method) {
    return m.merge_method
}
// GetMinEntriesToMerge gets the min_entries_to_merge property value. The minimum number of PRs that will be merged together in a group.
// returns a *int32 when successful
func (m *RepositoryRuleMergeQueue_parameters) GetMinEntriesToMerge()(*int32) {
    return m.min_entries_to_merge
}
// GetMinEntriesToMergeWaitMinutes gets the min_entries_to_merge_wait_minutes property value. The time merge queue should wait after the first PR is added to the queue for the minimum group size to be met. After this time has elapsed, the minimum group size will be ignored and a smaller group will be merged.
// returns a *int32 when successful
func (m *RepositoryRuleMergeQueue_parameters) GetMinEntriesToMergeWaitMinutes()(*int32) {
    return m.min_entries_to_merge_wait_minutes
}
// Serialize serializes information the current object
func (m *RepositoryRuleMergeQueue_parameters) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteInt32Value("check_response_timeout_minutes", m.GetCheckResponseTimeoutMinutes())
        if err != nil {
            return err
        }
    }
    if m.GetGroupingStrategy() != nil {
        cast := (*m.GetGroupingStrategy()).String()
        err := writer.WriteStringValue("grouping_strategy", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("max_entries_to_build", m.GetMaxEntriesToBuild())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("max_entries_to_merge", m.GetMaxEntriesToMerge())
        if err != nil {
            return err
        }
    }
    if m.GetMergeMethod() != nil {
        cast := (*m.GetMergeMethod()).String()
        err := writer.WriteStringValue("merge_method", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("min_entries_to_merge", m.GetMinEntriesToMerge())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("min_entries_to_merge_wait_minutes", m.GetMinEntriesToMergeWaitMinutes())
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
func (m *RepositoryRuleMergeQueue_parameters) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetCheckResponseTimeoutMinutes sets the check_response_timeout_minutes property value. Maximum time for a required status check to report a conclusion. After this much time has elapsed, checks that have not reported a conclusion will be assumed to have failed
func (m *RepositoryRuleMergeQueue_parameters) SetCheckResponseTimeoutMinutes(value *int32)() {
    m.check_response_timeout_minutes = value
}
// SetGroupingStrategy sets the grouping_strategy property value. When set to ALLGREEN, the merge commit created by merge queue for each PR in the group must pass all required checks to merge. When set to HEADGREEN, only the commit at the head of the merge group, i.e. the commit containing changes from all of the PRs in the group, must pass its required checks to merge.
func (m *RepositoryRuleMergeQueue_parameters) SetGroupingStrategy(value *RepositoryRuleMergeQueue_parameters_grouping_strategy)() {
    m.grouping_strategy = value
}
// SetMaxEntriesToBuild sets the max_entries_to_build property value. Limit the number of queued pull requests requesting checks and workflow runs at the same time.
func (m *RepositoryRuleMergeQueue_parameters) SetMaxEntriesToBuild(value *int32)() {
    m.max_entries_to_build = value
}
// SetMaxEntriesToMerge sets the max_entries_to_merge property value. The maximum number of PRs that will be merged together in a group.
func (m *RepositoryRuleMergeQueue_parameters) SetMaxEntriesToMerge(value *int32)() {
    m.max_entries_to_merge = value
}
// SetMergeMethod sets the merge_method property value. Method to use when merging changes from queued pull requests.
func (m *RepositoryRuleMergeQueue_parameters) SetMergeMethod(value *RepositoryRuleMergeQueue_parameters_merge_method)() {
    m.merge_method = value
}
// SetMinEntriesToMerge sets the min_entries_to_merge property value. The minimum number of PRs that will be merged together in a group.
func (m *RepositoryRuleMergeQueue_parameters) SetMinEntriesToMerge(value *int32)() {
    m.min_entries_to_merge = value
}
// SetMinEntriesToMergeWaitMinutes sets the min_entries_to_merge_wait_minutes property value. The time merge queue should wait after the first PR is added to the queue for the minimum group size to be met. After this time has elapsed, the minimum group size will be ignored and a smaller group will be merged.
func (m *RepositoryRuleMergeQueue_parameters) SetMinEntriesToMergeWaitMinutes(value *int32)() {
    m.min_entries_to_merge_wait_minutes = value
}
type RepositoryRuleMergeQueue_parametersable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCheckResponseTimeoutMinutes()(*int32)
    GetGroupingStrategy()(*RepositoryRuleMergeQueue_parameters_grouping_strategy)
    GetMaxEntriesToBuild()(*int32)
    GetMaxEntriesToMerge()(*int32)
    GetMergeMethod()(*RepositoryRuleMergeQueue_parameters_merge_method)
    GetMinEntriesToMerge()(*int32)
    GetMinEntriesToMergeWaitMinutes()(*int32)
    SetCheckResponseTimeoutMinutes(value *int32)()
    SetGroupingStrategy(value *RepositoryRuleMergeQueue_parameters_grouping_strategy)()
    SetMaxEntriesToBuild(value *int32)()
    SetMaxEntriesToMerge(value *int32)()
    SetMergeMethod(value *RepositoryRuleMergeQueue_parameters_merge_method)()
    SetMinEntriesToMerge(value *int32)()
    SetMinEntriesToMergeWaitMinutes(value *int32)()
}
