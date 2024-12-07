package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type BillingUsageReport_usageItems struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Date of the usage line item.
    date *string
    // Discount amount of the usage line item.
    discountAmount *float64
    // Gross amount of the usage line item.
    grossAmount *float64
    // Net amount of the usage line item.
    netAmount *float64
    // Name of the organization.
    organizationName *string
    // Price per unit of the usage line item.
    pricePerUnit *float64
    // Product name.
    product *string
    // Quantity of the usage line item.
    quantity *int32
    // Name of the repository.
    repositoryName *string
    // SKU name.
    sku *string
    // Unit type of the usage line item.
    unitType *string
}
// NewBillingUsageReport_usageItems instantiates a new BillingUsageReport_usageItems and sets the default values.
func NewBillingUsageReport_usageItems()(*BillingUsageReport_usageItems) {
    m := &BillingUsageReport_usageItems{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateBillingUsageReport_usageItemsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateBillingUsageReport_usageItemsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewBillingUsageReport_usageItems(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *BillingUsageReport_usageItems) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetDate gets the date property value. Date of the usage line item.
// returns a *string when successful
func (m *BillingUsageReport_usageItems) GetDate()(*string) {
    return m.date
}
// GetDiscountAmount gets the discountAmount property value. Discount amount of the usage line item.
// returns a *float64 when successful
func (m *BillingUsageReport_usageItems) GetDiscountAmount()(*float64) {
    return m.discountAmount
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *BillingUsageReport_usageItems) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["date"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDate(val)
        }
        return nil
    }
    res["discountAmount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDiscountAmount(val)
        }
        return nil
    }
    res["grossAmount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetGrossAmount(val)
        }
        return nil
    }
    res["netAmount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNetAmount(val)
        }
        return nil
    }
    res["organizationName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOrganizationName(val)
        }
        return nil
    }
    res["pricePerUnit"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPricePerUnit(val)
        }
        return nil
    }
    res["product"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetProduct(val)
        }
        return nil
    }
    res["quantity"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetQuantity(val)
        }
        return nil
    }
    res["repositoryName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRepositoryName(val)
        }
        return nil
    }
    res["sku"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSku(val)
        }
        return nil
    }
    res["unitType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUnitType(val)
        }
        return nil
    }
    return res
}
// GetGrossAmount gets the grossAmount property value. Gross amount of the usage line item.
// returns a *float64 when successful
func (m *BillingUsageReport_usageItems) GetGrossAmount()(*float64) {
    return m.grossAmount
}
// GetNetAmount gets the netAmount property value. Net amount of the usage line item.
// returns a *float64 when successful
func (m *BillingUsageReport_usageItems) GetNetAmount()(*float64) {
    return m.netAmount
}
// GetOrganizationName gets the organizationName property value. Name of the organization.
// returns a *string when successful
func (m *BillingUsageReport_usageItems) GetOrganizationName()(*string) {
    return m.organizationName
}
// GetPricePerUnit gets the pricePerUnit property value. Price per unit of the usage line item.
// returns a *float64 when successful
func (m *BillingUsageReport_usageItems) GetPricePerUnit()(*float64) {
    return m.pricePerUnit
}
// GetProduct gets the product property value. Product name.
// returns a *string when successful
func (m *BillingUsageReport_usageItems) GetProduct()(*string) {
    return m.product
}
// GetQuantity gets the quantity property value. Quantity of the usage line item.
// returns a *int32 when successful
func (m *BillingUsageReport_usageItems) GetQuantity()(*int32) {
    return m.quantity
}
// GetRepositoryName gets the repositoryName property value. Name of the repository.
// returns a *string when successful
func (m *BillingUsageReport_usageItems) GetRepositoryName()(*string) {
    return m.repositoryName
}
// GetSku gets the sku property value. SKU name.
// returns a *string when successful
func (m *BillingUsageReport_usageItems) GetSku()(*string) {
    return m.sku
}
// GetUnitType gets the unitType property value. Unit type of the usage line item.
// returns a *string when successful
func (m *BillingUsageReport_usageItems) GetUnitType()(*string) {
    return m.unitType
}
// Serialize serializes information the current object
func (m *BillingUsageReport_usageItems) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("date", m.GetDate())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteFloat64Value("discountAmount", m.GetDiscountAmount())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteFloat64Value("grossAmount", m.GetGrossAmount())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteFloat64Value("netAmount", m.GetNetAmount())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("organizationName", m.GetOrganizationName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteFloat64Value("pricePerUnit", m.GetPricePerUnit())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("product", m.GetProduct())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("quantity", m.GetQuantity())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("repositoryName", m.GetRepositoryName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("sku", m.GetSku())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("unitType", m.GetUnitType())
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
func (m *BillingUsageReport_usageItems) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetDate sets the date property value. Date of the usage line item.
func (m *BillingUsageReport_usageItems) SetDate(value *string)() {
    m.date = value
}
// SetDiscountAmount sets the discountAmount property value. Discount amount of the usage line item.
func (m *BillingUsageReport_usageItems) SetDiscountAmount(value *float64)() {
    m.discountAmount = value
}
// SetGrossAmount sets the grossAmount property value. Gross amount of the usage line item.
func (m *BillingUsageReport_usageItems) SetGrossAmount(value *float64)() {
    m.grossAmount = value
}
// SetNetAmount sets the netAmount property value. Net amount of the usage line item.
func (m *BillingUsageReport_usageItems) SetNetAmount(value *float64)() {
    m.netAmount = value
}
// SetOrganizationName sets the organizationName property value. Name of the organization.
func (m *BillingUsageReport_usageItems) SetOrganizationName(value *string)() {
    m.organizationName = value
}
// SetPricePerUnit sets the pricePerUnit property value. Price per unit of the usage line item.
func (m *BillingUsageReport_usageItems) SetPricePerUnit(value *float64)() {
    m.pricePerUnit = value
}
// SetProduct sets the product property value. Product name.
func (m *BillingUsageReport_usageItems) SetProduct(value *string)() {
    m.product = value
}
// SetQuantity sets the quantity property value. Quantity of the usage line item.
func (m *BillingUsageReport_usageItems) SetQuantity(value *int32)() {
    m.quantity = value
}
// SetRepositoryName sets the repositoryName property value. Name of the repository.
func (m *BillingUsageReport_usageItems) SetRepositoryName(value *string)() {
    m.repositoryName = value
}
// SetSku sets the sku property value. SKU name.
func (m *BillingUsageReport_usageItems) SetSku(value *string)() {
    m.sku = value
}
// SetUnitType sets the unitType property value. Unit type of the usage line item.
func (m *BillingUsageReport_usageItems) SetUnitType(value *string)() {
    m.unitType = value
}
type BillingUsageReport_usageItemsable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetDate()(*string)
    GetDiscountAmount()(*float64)
    GetGrossAmount()(*float64)
    GetNetAmount()(*float64)
    GetOrganizationName()(*string)
    GetPricePerUnit()(*float64)
    GetProduct()(*string)
    GetQuantity()(*int32)
    GetRepositoryName()(*string)
    GetSku()(*string)
    GetUnitType()(*string)
    SetDate(value *string)()
    SetDiscountAmount(value *float64)()
    SetGrossAmount(value *float64)()
    SetNetAmount(value *float64)()
    SetOrganizationName(value *string)()
    SetPricePerUnit(value *float64)()
    SetProduct(value *string)()
    SetQuantity(value *int32)()
    SetRepositoryName(value *string)()
    SetSku(value *string)()
    SetUnitType(value *string)()
}
