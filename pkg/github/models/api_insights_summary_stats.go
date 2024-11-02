package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ApiInsightsSummaryStats aPI Insights usage summary stats for an organization
type ApiInsightsSummaryStats struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The total number of requests that were rate limited within the queried time period
    rate_limited_request_count *int64
    // The total number of requests within the queried time period
    total_request_count *int64
}
// NewApiInsightsSummaryStats instantiates a new ApiInsightsSummaryStats and sets the default values.
func NewApiInsightsSummaryStats()(*ApiInsightsSummaryStats) {
    m := &ApiInsightsSummaryStats{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateApiInsightsSummaryStatsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateApiInsightsSummaryStatsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewApiInsightsSummaryStats(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *ApiInsightsSummaryStats) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ApiInsightsSummaryStats) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
// GetRateLimitedRequestCount gets the rate_limited_request_count property value. The total number of requests that were rate limited within the queried time period
// returns a *int64 when successful
func (m *ApiInsightsSummaryStats) GetRateLimitedRequestCount()(*int64) {
    return m.rate_limited_request_count
}
// GetTotalRequestCount gets the total_request_count property value. The total number of requests within the queried time period
// returns a *int64 when successful
func (m *ApiInsightsSummaryStats) GetTotalRequestCount()(*int64) {
    return m.total_request_count
}
// Serialize serializes information the current object
func (m *ApiInsightsSummaryStats) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteInt64Value("rate_limited_request_count", m.GetRateLimitedRequestCount())
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
func (m *ApiInsightsSummaryStats) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetRateLimitedRequestCount sets the rate_limited_request_count property value. The total number of requests that were rate limited within the queried time period
func (m *ApiInsightsSummaryStats) SetRateLimitedRequestCount(value *int64)() {
    m.rate_limited_request_count = value
}
// SetTotalRequestCount sets the total_request_count property value. The total number of requests within the queried time period
func (m *ApiInsightsSummaryStats) SetTotalRequestCount(value *int64)() {
    m.total_request_count = value
}
type ApiInsightsSummaryStatsable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetRateLimitedRequestCount()(*int64)
    GetTotalRequestCount()(*int64)
    SetRateLimitedRequestCount(value *int64)()
    SetTotalRequestCount(value *int64)()
}
