package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type SecretScanningScanHistory struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The backfill_scans property
    backfill_scans []SecretScanningScanable
    // The custom_pattern_backfill_scans property
    custom_pattern_backfill_scans []SecretScanningScanable
    // The incremental_scans property
    incremental_scans []SecretScanningScanable
    // The pattern_update_scans property
    pattern_update_scans []SecretScanningScanable
}
// NewSecretScanningScanHistory instantiates a new SecretScanningScanHistory and sets the default values.
func NewSecretScanningScanHistory()(*SecretScanningScanHistory) {
    m := &SecretScanningScanHistory{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateSecretScanningScanHistoryFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateSecretScanningScanHistoryFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewSecretScanningScanHistory(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *SecretScanningScanHistory) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetBackfillScans gets the backfill_scans property value. The backfill_scans property
// returns a []SecretScanningScanable when successful
func (m *SecretScanningScanHistory) GetBackfillScans()([]SecretScanningScanable) {
    return m.backfill_scans
}
// GetCustomPatternBackfillScans gets the custom_pattern_backfill_scans property value. The custom_pattern_backfill_scans property
// returns a []SecretScanningScanable when successful
func (m *SecretScanningScanHistory) GetCustomPatternBackfillScans()([]SecretScanningScanable) {
    return m.custom_pattern_backfill_scans
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *SecretScanningScanHistory) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["backfill_scans"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateSecretScanningScanFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]SecretScanningScanable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(SecretScanningScanable)
                }
            }
            m.SetBackfillScans(res)
        }
        return nil
    }
    res["custom_pattern_backfill_scans"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateSecretScanningScanFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]SecretScanningScanable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(SecretScanningScanable)
                }
            }
            m.SetCustomPatternBackfillScans(res)
        }
        return nil
    }
    res["incremental_scans"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateSecretScanningScanFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]SecretScanningScanable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(SecretScanningScanable)
                }
            }
            m.SetIncrementalScans(res)
        }
        return nil
    }
    res["pattern_update_scans"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateSecretScanningScanFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]SecretScanningScanable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(SecretScanningScanable)
                }
            }
            m.SetPatternUpdateScans(res)
        }
        return nil
    }
    return res
}
// GetIncrementalScans gets the incremental_scans property value. The incremental_scans property
// returns a []SecretScanningScanable when successful
func (m *SecretScanningScanHistory) GetIncrementalScans()([]SecretScanningScanable) {
    return m.incremental_scans
}
// GetPatternUpdateScans gets the pattern_update_scans property value. The pattern_update_scans property
// returns a []SecretScanningScanable when successful
func (m *SecretScanningScanHistory) GetPatternUpdateScans()([]SecretScanningScanable) {
    return m.pattern_update_scans
}
// Serialize serializes information the current object
func (m *SecretScanningScanHistory) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetBackfillScans() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetBackfillScans()))
        for i, v := range m.GetBackfillScans() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("backfill_scans", cast)
        if err != nil {
            return err
        }
    }
    if m.GetCustomPatternBackfillScans() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetCustomPatternBackfillScans()))
        for i, v := range m.GetCustomPatternBackfillScans() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("custom_pattern_backfill_scans", cast)
        if err != nil {
            return err
        }
    }
    if m.GetIncrementalScans() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetIncrementalScans()))
        for i, v := range m.GetIncrementalScans() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("incremental_scans", cast)
        if err != nil {
            return err
        }
    }
    if m.GetPatternUpdateScans() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetPatternUpdateScans()))
        for i, v := range m.GetPatternUpdateScans() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("pattern_update_scans", cast)
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
func (m *SecretScanningScanHistory) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetBackfillScans sets the backfill_scans property value. The backfill_scans property
func (m *SecretScanningScanHistory) SetBackfillScans(value []SecretScanningScanable)() {
    m.backfill_scans = value
}
// SetCustomPatternBackfillScans sets the custom_pattern_backfill_scans property value. The custom_pattern_backfill_scans property
func (m *SecretScanningScanHistory) SetCustomPatternBackfillScans(value []SecretScanningScanable)() {
    m.custom_pattern_backfill_scans = value
}
// SetIncrementalScans sets the incremental_scans property value. The incremental_scans property
func (m *SecretScanningScanHistory) SetIncrementalScans(value []SecretScanningScanable)() {
    m.incremental_scans = value
}
// SetPatternUpdateScans sets the pattern_update_scans property value. The pattern_update_scans property
func (m *SecretScanningScanHistory) SetPatternUpdateScans(value []SecretScanningScanable)() {
    m.pattern_update_scans = value
}
type SecretScanningScanHistoryable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBackfillScans()([]SecretScanningScanable)
    GetCustomPatternBackfillScans()([]SecretScanningScanable)
    GetIncrementalScans()([]SecretScanningScanable)
    GetPatternUpdateScans()([]SecretScanningScanable)
    SetBackfillScans(value []SecretScanningScanable)()
    SetCustomPatternBackfillScans(value []SecretScanningScanable)()
    SetIncrementalScans(value []SecretScanningScanable)()
    SetPatternUpdateScans(value []SecretScanningScanable)()
}
