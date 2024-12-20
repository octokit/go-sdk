package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type ApiOverview_domains_actions_inbound struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The full_domains property
    full_domains []string
    // The wildcard_domains property
    wildcard_domains []string
}
// NewApiOverview_domains_actions_inbound instantiates a new ApiOverview_domains_actions_inbound and sets the default values.
func NewApiOverview_domains_actions_inbound()(*ApiOverview_domains_actions_inbound) {
    m := &ApiOverview_domains_actions_inbound{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateApiOverview_domains_actions_inboundFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateApiOverview_domains_actions_inboundFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewApiOverview_domains_actions_inbound(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *ApiOverview_domains_actions_inbound) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ApiOverview_domains_actions_inbound) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["full_domains"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
            m.SetFullDomains(res)
        }
        return nil
    }
    res["wildcard_domains"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
            m.SetWildcardDomains(res)
        }
        return nil
    }
    return res
}
// GetFullDomains gets the full_domains property value. The full_domains property
// returns a []string when successful
func (m *ApiOverview_domains_actions_inbound) GetFullDomains()([]string) {
    return m.full_domains
}
// GetWildcardDomains gets the wildcard_domains property value. The wildcard_domains property
// returns a []string when successful
func (m *ApiOverview_domains_actions_inbound) GetWildcardDomains()([]string) {
    return m.wildcard_domains
}
// Serialize serializes information the current object
func (m *ApiOverview_domains_actions_inbound) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetFullDomains() != nil {
        err := writer.WriteCollectionOfStringValues("full_domains", m.GetFullDomains())
        if err != nil {
            return err
        }
    }
    if m.GetWildcardDomains() != nil {
        err := writer.WriteCollectionOfStringValues("wildcard_domains", m.GetWildcardDomains())
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
func (m *ApiOverview_domains_actions_inbound) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetFullDomains sets the full_domains property value. The full_domains property
func (m *ApiOverview_domains_actions_inbound) SetFullDomains(value []string)() {
    m.full_domains = value
}
// SetWildcardDomains sets the wildcard_domains property value. The wildcard_domains property
func (m *ApiOverview_domains_actions_inbound) SetWildcardDomains(value []string)() {
    m.wildcard_domains = value
}
type ApiOverview_domains_actions_inboundable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetFullDomains()([]string)
    GetWildcardDomains()([]string)
    SetFullDomains(value []string)()
    SetWildcardDomains(value []string)()
}
