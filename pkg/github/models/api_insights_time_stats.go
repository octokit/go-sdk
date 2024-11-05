package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type ApiInsightsTimeStats struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The rate_limited_request_count property
    rate_limited_request_count *int64
    // The timestamp property
    timestamp *string
    // The total_request_count property
    total_request_count *int64
}
// NewApiInsightsTimeStats instantiates a new ApiInsightsTimeStats and sets the default values.
func NewApiInsightsTimeStats()(*ApiInsightsTimeStats) {
    m := &ApiInsightsTimeStats{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateApiInsightsTimeStatsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateApiInsightsTimeStatsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewApiInsightsTimeStats(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *ApiInsightsTimeStats) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ApiInsightsTimeStats) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["rate_limited_request_count"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRateLimitedRequestCount(val)
        }
        return nil
    }
    res["timestamp"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTimestamp(val)
        }
        return nil
    }
    res["total_request_count"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTotalRequestCount(val)
        }
        return nil
    }
    return res
}
// GetRateLimitedRequestCount gets the rate_limited_request_count property value. The rate_limited_request_count property
// returns a *int64 when successful
func (m *ApiInsightsTimeStats) GetRateLimitedRequestCount()(*int64) {
    return m.rate_limited_request_count
}
// GetTimestamp gets the timestamp property value. The timestamp property
// returns a *string when successful
func (m *ApiInsightsTimeStats) GetTimestamp()(*string) {
    return m.timestamp
}
// GetTotalRequestCount gets the total_request_count property value. The total_request_count property
// returns a *int64 when successful
func (m *ApiInsightsTimeStats) GetTotalRequestCount()(*int64) {
    return m.total_request_count
}
// Serialize serializes information the current object
func (m *ApiInsightsTimeStats) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteInt64Value("rate_limited_request_count", m.GetRateLimitedRequestCount())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("timestamp", m.GetTimestamp())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt64Value("total_request_count", m.GetTotalRequestCount())
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
func (m *ApiInsightsTimeStats) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetRateLimitedRequestCount sets the rate_limited_request_count property value. The rate_limited_request_count property
func (m *ApiInsightsTimeStats) SetRateLimitedRequestCount(value *int64)() {
    m.rate_limited_request_count = value
}
// SetTimestamp sets the timestamp property value. The timestamp property
func (m *ApiInsightsTimeStats) SetTimestamp(value *string)() {
    m.timestamp = value
}
// SetTotalRequestCount sets the total_request_count property value. The total_request_count property
func (m *ApiInsightsTimeStats) SetTotalRequestCount(value *int64)() {
    m.total_request_count = value
}
type ApiInsightsTimeStatsable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetRateLimitedRequestCount()(*int64)
    GetTimestamp()(*string)
    GetTotalRequestCount()(*int64)
    SetRateLimitedRequestCount(value *int64)()
    SetTimestamp(value *string)()
    SetTotalRequestCount(value *int64)()
}
