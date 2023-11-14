package repos

import (
    i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e "github.com/octokit/go-sdk/github/octokit/models"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemItemCommitsItemCheckSuitesGetResponse 
type ItemItemCommitsItemCheckSuitesGetResponse struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The check_suites property
    check_suites []i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.CheckSuiteable
    // The total_count property
    total_count *int32
}
// NewItemItemCommitsItemCheckSuitesGetResponse instantiates a new ItemItemCommitsItemCheckSuitesGetResponse and sets the default values.
func NewItemItemCommitsItemCheckSuitesGetResponse()(*ItemItemCommitsItemCheckSuitesGetResponse) {
    m := &ItemItemCommitsItemCheckSuitesGetResponse{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateItemItemCommitsItemCheckSuitesGetResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemItemCommitsItemCheckSuitesGetResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemItemCommitsItemCheckSuitesGetResponse(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ItemItemCommitsItemCheckSuitesGetResponse) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetCheckSuites gets the check_suites property value. The check_suites property
func (m *ItemItemCommitsItemCheckSuitesGetResponse) GetCheckSuites()([]i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.CheckSuiteable) {
    return m.check_suites
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ItemItemCommitsItemCheckSuitesGetResponse) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["check_suites"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.CreateCheckSuiteFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.CheckSuiteable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.CheckSuiteable)
                }
            }
            m.SetCheckSuites(res)
        }
        return nil
    }
    res["total_count"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTotalCount(val)
        }
        return nil
    }
    return res
}
// GetTotalCount gets the total_count property value. The total_count property
func (m *ItemItemCommitsItemCheckSuitesGetResponse) GetTotalCount()(*int32) {
    return m.total_count
}
// Serialize serializes information the current object
func (m *ItemItemCommitsItemCheckSuitesGetResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetCheckSuites() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetCheckSuites()))
        for i, v := range m.GetCheckSuites() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("check_suites", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("total_count", m.GetTotalCount())
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
func (m *ItemItemCommitsItemCheckSuitesGetResponse) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetCheckSuites sets the check_suites property value. The check_suites property
func (m *ItemItemCommitsItemCheckSuitesGetResponse) SetCheckSuites(value []i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.CheckSuiteable)() {
    m.check_suites = value
}
// SetTotalCount sets the total_count property value. The total_count property
func (m *ItemItemCommitsItemCheckSuitesGetResponse) SetTotalCount(value *int32)() {
    m.total_count = value
}
// ItemItemCommitsItemCheckSuitesGetResponseable 
type ItemItemCommitsItemCheckSuitesGetResponseable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCheckSuites()([]i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.CheckSuiteable)
    GetTotalCount()(*int32)
    SetCheckSuites(value []i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.CheckSuiteable)()
    SetTotalCount(value *int32)()
}
