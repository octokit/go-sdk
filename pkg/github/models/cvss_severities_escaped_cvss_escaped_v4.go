package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type CvssSeverities_cvss_v4 struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The CVSS 4 score.
    score *float64
    // The CVSS 4 vector string.
    vector_string *string
}
// NewCvssSeverities_cvss_v4 instantiates a new CvssSeverities_cvss_v4 and sets the default values.
func NewCvssSeverities_cvss_v4()(*CvssSeverities_cvss_v4) {
    m := &CvssSeverities_cvss_v4{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateCvssSeverities_cvss_v4FromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateCvssSeverities_cvss_v4FromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCvssSeverities_cvss_v4(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *CvssSeverities_cvss_v4) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *CvssSeverities_cvss_v4) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["score"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetScore(val)
        }
        return nil
    }
    res["vector_string"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetVectorString(val)
        }
        return nil
    }
    return res
}
// GetScore gets the score property value. The CVSS 4 score.
// returns a *float64 when successful
func (m *CvssSeverities_cvss_v4) GetScore()(*float64) {
    return m.score
}
// GetVectorString gets the vector_string property value. The CVSS 4 vector string.
// returns a *string when successful
func (m *CvssSeverities_cvss_v4) GetVectorString()(*string) {
    return m.vector_string
}
// Serialize serializes information the current object
func (m *CvssSeverities_cvss_v4) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("vector_string", m.GetVectorString())
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
func (m *CvssSeverities_cvss_v4) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetScore sets the score property value. The CVSS 4 score.
func (m *CvssSeverities_cvss_v4) SetScore(value *float64)() {
    m.score = value
}
// SetVectorString sets the vector_string property value. The CVSS 4 vector string.
func (m *CvssSeverities_cvss_v4) SetVectorString(value *string)() {
    m.vector_string = value
}
type CvssSeverities_cvss_v4able interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetScore()(*float64)
    GetVectorString()(*string)
    SetScore(value *float64)()
    SetVectorString(value *string)()
}
