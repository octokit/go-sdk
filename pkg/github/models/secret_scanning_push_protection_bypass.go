package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type SecretScanningPushProtectionBypass struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The time that the bypass will expire in ISO 8601 format: `YYYY-MM-DDTHH:MM:SSZ`.
    expire_at *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The reason for bypassing push protection.
    reason *SecretScanningPushProtectionBypassReason
    // The token type this bypass is for.
    token_type *string
}
// NewSecretScanningPushProtectionBypass instantiates a new SecretScanningPushProtectionBypass and sets the default values.
func NewSecretScanningPushProtectionBypass()(*SecretScanningPushProtectionBypass) {
    m := &SecretScanningPushProtectionBypass{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateSecretScanningPushProtectionBypassFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateSecretScanningPushProtectionBypassFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewSecretScanningPushProtectionBypass(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *SecretScanningPushProtectionBypass) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetExpireAt gets the expire_at property value. The time that the bypass will expire in ISO 8601 format: `YYYY-MM-DDTHH:MM:SSZ`.
// returns a *Time when successful
func (m *SecretScanningPushProtectionBypass) GetExpireAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.expire_at
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *SecretScanningPushProtectionBypass) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["expire_at"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetExpireAt(val)
        }
        return nil
    }
    res["reason"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseSecretScanningPushProtectionBypassReason)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetReason(val.(*SecretScanningPushProtectionBypassReason))
        }
        return nil
    }
    res["token_type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTokenType(val)
        }
        return nil
    }
    return res
}
// GetReason gets the reason property value. The reason for bypassing push protection.
// returns a *SecretScanningPushProtectionBypassReason when successful
func (m *SecretScanningPushProtectionBypass) GetReason()(*SecretScanningPushProtectionBypassReason) {
    return m.reason
}
// GetTokenType gets the token_type property value. The token type this bypass is for.
// returns a *string when successful
func (m *SecretScanningPushProtectionBypass) GetTokenType()(*string) {
    return m.token_type
}
// Serialize serializes information the current object
func (m *SecretScanningPushProtectionBypass) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteTimeValue("expire_at", m.GetExpireAt())
        if err != nil {
            return err
        }
    }
    if m.GetReason() != nil {
        cast := (*m.GetReason()).String()
        err := writer.WriteStringValue("reason", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("token_type", m.GetTokenType())
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
func (m *SecretScanningPushProtectionBypass) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetExpireAt sets the expire_at property value. The time that the bypass will expire in ISO 8601 format: `YYYY-MM-DDTHH:MM:SSZ`.
func (m *SecretScanningPushProtectionBypass) SetExpireAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.expire_at = value
}
// SetReason sets the reason property value. The reason for bypassing push protection.
func (m *SecretScanningPushProtectionBypass) SetReason(value *SecretScanningPushProtectionBypassReason)() {
    m.reason = value
}
// SetTokenType sets the token_type property value. The token type this bypass is for.
func (m *SecretScanningPushProtectionBypass) SetTokenType(value *string)() {
    m.token_type = value
}
type SecretScanningPushProtectionBypassable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetExpireAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetReason()(*SecretScanningPushProtectionBypassReason)
    GetTokenType()(*string)
    SetExpireAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetReason(value *SecretScanningPushProtectionBypassReason)()
    SetTokenType(value *string)()
}
