package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type ApiOverview_domains_artifact_attestations struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The services property
    services []string
    // The trust_domain property
    trust_domain *string
}
// NewApiOverview_domains_artifact_attestations instantiates a new ApiOverview_domains_artifact_attestations and sets the default values.
func NewApiOverview_domains_artifact_attestations()(*ApiOverview_domains_artifact_attestations) {
    m := &ApiOverview_domains_artifact_attestations{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateApiOverview_domains_artifact_attestationsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateApiOverview_domains_artifact_attestationsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewApiOverview_domains_artifact_attestations(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *ApiOverview_domains_artifact_attestations) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ApiOverview_domains_artifact_attestations) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["services"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
            m.SetServices(res)
        }
        return nil
    }
    res["trust_domain"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTrustDomain(val)
        }
        return nil
    }
    return res
}
// GetServices gets the services property value. The services property
// returns a []string when successful
func (m *ApiOverview_domains_artifact_attestations) GetServices()([]string) {
    return m.services
}
// GetTrustDomain gets the trust_domain property value. The trust_domain property
// returns a *string when successful
func (m *ApiOverview_domains_artifact_attestations) GetTrustDomain()(*string) {
    return m.trust_domain
}
// Serialize serializes information the current object
func (m *ApiOverview_domains_artifact_attestations) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetServices() != nil {
        err := writer.WriteCollectionOfStringValues("services", m.GetServices())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("trust_domain", m.GetTrustDomain())
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
func (m *ApiOverview_domains_artifact_attestations) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetServices sets the services property value. The services property
func (m *ApiOverview_domains_artifact_attestations) SetServices(value []string)() {
    m.services = value
}
// SetTrustDomain sets the trust_domain property value. The trust_domain property
func (m *ApiOverview_domains_artifact_attestations) SetTrustDomain(value *string)() {
    m.trust_domain = value
}
type ApiOverview_domains_artifact_attestationsable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetServices()([]string)
    GetTrustDomain()(*string)
    SetServices(value []string)()
    SetTrustDomain(value *string)()
}
