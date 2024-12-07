package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type BillingUsageReport struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The usageItems property
    usageItems []BillingUsageReport_usageItemsable
}
// NewBillingUsageReport instantiates a new BillingUsageReport and sets the default values.
func NewBillingUsageReport()(*BillingUsageReport) {
    m := &BillingUsageReport{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateBillingUsageReportFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateBillingUsageReportFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewBillingUsageReport(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *BillingUsageReport) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *BillingUsageReport) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["usageItems"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateBillingUsageReport_usageItemsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]BillingUsageReport_usageItemsable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(BillingUsageReport_usageItemsable)
                }
            }
            m.SetUsageItems(res)
        }
        return nil
    }
    return res
}
// GetUsageItems gets the usageItems property value. The usageItems property
// returns a []BillingUsageReport_usageItemsable when successful
func (m *BillingUsageReport) GetUsageItems()([]BillingUsageReport_usageItemsable) {
    return m.usageItems
}
// Serialize serializes information the current object
func (m *BillingUsageReport) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetUsageItems() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetUsageItems()))
        for i, v := range m.GetUsageItems() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("usageItems", cast)
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
func (m *BillingUsageReport) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetUsageItems sets the usageItems property value. The usageItems property
func (m *BillingUsageReport) SetUsageItems(value []BillingUsageReport_usageItemsable)() {
    m.usageItems = value
}
type BillingUsageReportable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetUsageItems()([]BillingUsageReport_usageItemsable)
    SetUsageItems(value []BillingUsageReport_usageItemsable)()
}
