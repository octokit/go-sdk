package repos

import (
    i33bcfa37677dfd8b71eaf22ad3bfd140ba1702e7a8418558499c22418bf2fef2 "github.com/octokit/go-sdk/models"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemItemRulesetsPostRequestBody 
type ItemItemRulesetsPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The actors that can bypass the rules in this ruleset
    bypass_actors []i33bcfa37677dfd8b71eaf22ad3bfd140ba1702e7a8418558499c22418bf2fef2.RepositoryRulesetBypassActorable
    // Parameters for a repository ruleset ref name condition
    conditions i33bcfa37677dfd8b71eaf22ad3bfd140ba1702e7a8418558499c22418bf2fef2.RepositoryRulesetConditionsable
    // The enforcement level of the ruleset. `evaluate` allows admins to test rules before enforcing them. Admins can view insights on the Rule Insights page (`evaluate` is only available with GitHub Enterprise).
    enforcement *i33bcfa37677dfd8b71eaf22ad3bfd140ba1702e7a8418558499c22418bf2fef2.RepositoryRuleEnforcement
    // The name of the ruleset.
    name *string
    // An array of rules within the ruleset.
    rules []i33bcfa37677dfd8b71eaf22ad3bfd140ba1702e7a8418558499c22418bf2fef2.RepositoryRuleable
}
// NewItemItemRulesetsPostRequestBody instantiates a new ItemItemRulesetsPostRequestBody and sets the default values.
func NewItemItemRulesetsPostRequestBody()(*ItemItemRulesetsPostRequestBody) {
    m := &ItemItemRulesetsPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateItemItemRulesetsPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemItemRulesetsPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemItemRulesetsPostRequestBody(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ItemItemRulesetsPostRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetBypassActors gets the bypass_actors property value. The actors that can bypass the rules in this ruleset
func (m *ItemItemRulesetsPostRequestBody) GetBypassActors()([]i33bcfa37677dfd8b71eaf22ad3bfd140ba1702e7a8418558499c22418bf2fef2.RepositoryRulesetBypassActorable) {
    return m.bypass_actors
}
// GetConditions gets the conditions property value. Parameters for a repository ruleset ref name condition
func (m *ItemItemRulesetsPostRequestBody) GetConditions()(i33bcfa37677dfd8b71eaf22ad3bfd140ba1702e7a8418558499c22418bf2fef2.RepositoryRulesetConditionsable) {
    return m.conditions
}
// GetEnforcement gets the enforcement property value. The enforcement level of the ruleset. `evaluate` allows admins to test rules before enforcing them. Admins can view insights on the Rule Insights page (`evaluate` is only available with GitHub Enterprise).
func (m *ItemItemRulesetsPostRequestBody) GetEnforcement()(*i33bcfa37677dfd8b71eaf22ad3bfd140ba1702e7a8418558499c22418bf2fef2.RepositoryRuleEnforcement) {
    return m.enforcement
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ItemItemRulesetsPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["bypass_actors"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(i33bcfa37677dfd8b71eaf22ad3bfd140ba1702e7a8418558499c22418bf2fef2.CreateRepositoryRulesetBypassActorFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]i33bcfa37677dfd8b71eaf22ad3bfd140ba1702e7a8418558499c22418bf2fef2.RepositoryRulesetBypassActorable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(i33bcfa37677dfd8b71eaf22ad3bfd140ba1702e7a8418558499c22418bf2fef2.RepositoryRulesetBypassActorable)
                }
            }
            m.SetBypassActors(res)
        }
        return nil
    }
    res["conditions"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(i33bcfa37677dfd8b71eaf22ad3bfd140ba1702e7a8418558499c22418bf2fef2.CreateRepositoryRulesetConditionsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetConditions(val.(i33bcfa37677dfd8b71eaf22ad3bfd140ba1702e7a8418558499c22418bf2fef2.RepositoryRulesetConditionsable))
        }
        return nil
    }
    res["enforcement"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(i33bcfa37677dfd8b71eaf22ad3bfd140ba1702e7a8418558499c22418bf2fef2.ParseRepositoryRuleEnforcement)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEnforcement(val.(*i33bcfa37677dfd8b71eaf22ad3bfd140ba1702e7a8418558499c22418bf2fef2.RepositoryRuleEnforcement))
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
    res["rules"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(i33bcfa37677dfd8b71eaf22ad3bfd140ba1702e7a8418558499c22418bf2fef2.CreateRepositoryRuleFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]i33bcfa37677dfd8b71eaf22ad3bfd140ba1702e7a8418558499c22418bf2fef2.RepositoryRuleable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(i33bcfa37677dfd8b71eaf22ad3bfd140ba1702e7a8418558499c22418bf2fef2.RepositoryRuleable)
                }
            }
            m.SetRules(res)
        }
        return nil
    }
    return res
}
// GetName gets the name property value. The name of the ruleset.
func (m *ItemItemRulesetsPostRequestBody) GetName()(*string) {
    return m.name
}
// GetRules gets the rules property value. An array of rules within the ruleset.
func (m *ItemItemRulesetsPostRequestBody) GetRules()([]i33bcfa37677dfd8b71eaf22ad3bfd140ba1702e7a8418558499c22418bf2fef2.RepositoryRuleable) {
    return m.rules
}
// Serialize serializes information the current object
func (m *ItemItemRulesetsPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetBypassActors() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetBypassActors()))
        for i, v := range m.GetBypassActors() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("bypass_actors", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("conditions", m.GetConditions())
        if err != nil {
            return err
        }
    }
    if m.GetEnforcement() != nil {
        cast := (*m.GetEnforcement()).String()
        err := writer.WriteStringValue("enforcement", &cast)
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
    if m.GetRules() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetRules()))
        for i, v := range m.GetRules() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("rules", cast)
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
func (m *ItemItemRulesetsPostRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetBypassActors sets the bypass_actors property value. The actors that can bypass the rules in this ruleset
func (m *ItemItemRulesetsPostRequestBody) SetBypassActors(value []i33bcfa37677dfd8b71eaf22ad3bfd140ba1702e7a8418558499c22418bf2fef2.RepositoryRulesetBypassActorable)() {
    m.bypass_actors = value
}
// SetConditions sets the conditions property value. Parameters for a repository ruleset ref name condition
func (m *ItemItemRulesetsPostRequestBody) SetConditions(value i33bcfa37677dfd8b71eaf22ad3bfd140ba1702e7a8418558499c22418bf2fef2.RepositoryRulesetConditionsable)() {
    m.conditions = value
}
// SetEnforcement sets the enforcement property value. The enforcement level of the ruleset. `evaluate` allows admins to test rules before enforcing them. Admins can view insights on the Rule Insights page (`evaluate` is only available with GitHub Enterprise).
func (m *ItemItemRulesetsPostRequestBody) SetEnforcement(value *i33bcfa37677dfd8b71eaf22ad3bfd140ba1702e7a8418558499c22418bf2fef2.RepositoryRuleEnforcement)() {
    m.enforcement = value
}
// SetName sets the name property value. The name of the ruleset.
func (m *ItemItemRulesetsPostRequestBody) SetName(value *string)() {
    m.name = value
}
// SetRules sets the rules property value. An array of rules within the ruleset.
func (m *ItemItemRulesetsPostRequestBody) SetRules(value []i33bcfa37677dfd8b71eaf22ad3bfd140ba1702e7a8418558499c22418bf2fef2.RepositoryRuleable)() {
    m.rules = value
}
// ItemItemRulesetsPostRequestBodyable 
type ItemItemRulesetsPostRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBypassActors()([]i33bcfa37677dfd8b71eaf22ad3bfd140ba1702e7a8418558499c22418bf2fef2.RepositoryRulesetBypassActorable)
    GetConditions()(i33bcfa37677dfd8b71eaf22ad3bfd140ba1702e7a8418558499c22418bf2fef2.RepositoryRulesetConditionsable)
    GetEnforcement()(*i33bcfa37677dfd8b71eaf22ad3bfd140ba1702e7a8418558499c22418bf2fef2.RepositoryRuleEnforcement)
    GetName()(*string)
    GetRules()([]i33bcfa37677dfd8b71eaf22ad3bfd140ba1702e7a8418558499c22418bf2fef2.RepositoryRuleable)
    SetBypassActors(value []i33bcfa37677dfd8b71eaf22ad3bfd140ba1702e7a8418558499c22418bf2fef2.RepositoryRulesetBypassActorable)()
    SetConditions(value i33bcfa37677dfd8b71eaf22ad3bfd140ba1702e7a8418558499c22418bf2fef2.RepositoryRulesetConditionsable)()
    SetEnforcement(value *i33bcfa37677dfd8b71eaf22ad3bfd140ba1702e7a8418558499c22418bf2fef2.RepositoryRuleEnforcement)()
    SetName(value *string)()
    SetRules(value []i33bcfa37677dfd8b71eaf22ad3bfd140ba1702e7a8418558499c22418bf2fef2.RepositoryRuleable)()
}
