package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type CvssSeverities struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The cvss_v3 property
    cvss_v3 CvssSeverities_cvss_v3able
    // The cvss_v4 property
    cvss_v4 CvssSeverities_cvss_v4able
}
// NewCvssSeverities instantiates a new CvssSeverities and sets the default values.
func NewCvssSeverities()(*CvssSeverities) {
    m := &CvssSeverities{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateCvssSeveritiesFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateCvssSeveritiesFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCvssSeverities(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *CvssSeverities) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetCvssV3 gets the cvss_v3 property value. The cvss_v3 property
// returns a CvssSeverities_cvss_v3able when successful
func (m *CvssSeverities) GetCvssV3()(CvssSeverities_cvss_v3able) {
    return m.cvss_v3
}
// GetCvssV4 gets the cvss_v4 property value. The cvss_v4 property
// returns a CvssSeverities_cvss_v4able when successful
func (m *CvssSeverities) GetCvssV4()(CvssSeverities_cvss_v4able) {
    return m.cvss_v4
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *CvssSeverities) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["cvss_v3"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateCvssSeverities_cvss_v3FromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCvssV3(val.(CvssSeverities_cvss_v3able))
        }
        return nil
    }
    res["cvss_v4"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateCvssSeverities_cvss_v4FromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCvssV4(val.(CvssSeverities_cvss_v4able))
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *CvssSeverities) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("cvss_v3", m.GetCvssV3())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("cvss_v4", m.GetCvssV4())
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
func (m *CvssSeverities) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetCvssV3 sets the cvss_v3 property value. The cvss_v3 property
func (m *CvssSeverities) SetCvssV3(value CvssSeverities_cvss_v3able)() {
    m.cvss_v3 = value
}
// SetCvssV4 sets the cvss_v4 property value. The cvss_v4 property
func (m *CvssSeverities) SetCvssV4(value CvssSeverities_cvss_v4able)() {
    m.cvss_v4 = value
}
type CvssSeveritiesable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCvssV3()(CvssSeverities_cvss_v3able)
    GetCvssV4()(CvssSeverities_cvss_v4able)
    SetCvssV3(value CvssSeverities_cvss_v3able)()
    SetCvssV4(value CvssSeverities_cvss_v4able)()
}
