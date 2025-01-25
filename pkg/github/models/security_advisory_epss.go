package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// SecurityAdvisoryEpss the EPSS scores as calculated by the [Exploit Prediction Scoring System](https://www.first.org/epss).
type SecurityAdvisoryEpss struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The percentage property
    percentage *float64
    // The percentile property
    percentile *float64
}
// NewSecurityAdvisoryEpss instantiates a new SecurityAdvisoryEpss and sets the default values.
func NewSecurityAdvisoryEpss()(*SecurityAdvisoryEpss) {
    m := &SecurityAdvisoryEpss{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateSecurityAdvisoryEpssFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateSecurityAdvisoryEpssFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewSecurityAdvisoryEpss(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *SecurityAdvisoryEpss) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *SecurityAdvisoryEpss) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["percentage"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPercentage(val)
        }
        return nil
    }
    res["percentile"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPercentile(val)
        }
        return nil
    }
    return res
}
// GetPercentage gets the percentage property value. The percentage property
// returns a *float64 when successful
func (m *SecurityAdvisoryEpss) GetPercentage()(*float64) {
    return m.percentage
}
// GetPercentile gets the percentile property value. The percentile property
// returns a *float64 when successful
func (m *SecurityAdvisoryEpss) GetPercentile()(*float64) {
    return m.percentile
}
// Serialize serializes information the current object
func (m *SecurityAdvisoryEpss) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteFloat64Value("percentage", m.GetPercentage())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteFloat64Value("percentile", m.GetPercentile())
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
func (m *SecurityAdvisoryEpss) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetPercentage sets the percentage property value. The percentage property
func (m *SecurityAdvisoryEpss) SetPercentage(value *float64)() {
    m.percentage = value
}
// SetPercentile sets the percentile property value. The percentile property
func (m *SecurityAdvisoryEpss) SetPercentile(value *float64)() {
    m.percentile = value
}
type SecurityAdvisoryEpssable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetPercentage()(*float64)
    GetPercentile()(*float64)
    SetPercentage(value *float64)()
    SetPercentile(value *float64)()
}
