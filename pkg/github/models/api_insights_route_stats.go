package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type ApiInsightsRouteStats struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The API path's route template
    api_route *string
    // The HTTP method
    http_method *string
    // The last_rate_limited_timestamp property
    last_rate_limited_timestamp *string
    // The last_request_timestamp property
    last_request_timestamp *string
    // The total number of requests that were rate limited within the queried time period
    rate_limited_request_count *int64
    // The total number of requests within the queried time period
    total_request_count *int64
}
// NewApiInsightsRouteStats instantiates a new ApiInsightsRouteStats and sets the default values.
func NewApiInsightsRouteStats()(*ApiInsightsRouteStats) {
    m := &ApiInsightsRouteStats{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateApiInsightsRouteStatsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateApiInsightsRouteStatsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewApiInsightsRouteStats(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *ApiInsightsRouteStats) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetApiRoute gets the api_route property value. The API path's route template
// returns a *string when successful
func (m *ApiInsightsRouteStats) GetApiRoute()(*string) {
    return m.api_route
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ApiInsightsRouteStats) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["api_route"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetApiRoute(val)
        }
        return nil
    }
    res["http_method"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHttpMethod(val)
        }
        return nil
    }
    res["last_rate_limited_timestamp"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastRateLimitedTimestamp(val)
        }
        return nil
    }
    res["last_request_timestamp"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastRequestTimestamp(val)
        }
        return nil
    }
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
// GetHttpMethod gets the http_method property value. The HTTP method
// returns a *string when successful
func (m *ApiInsightsRouteStats) GetHttpMethod()(*string) {
    return m.http_method
}
// GetLastRateLimitedTimestamp gets the last_rate_limited_timestamp property value. The last_rate_limited_timestamp property
// returns a *string when successful
func (m *ApiInsightsRouteStats) GetLastRateLimitedTimestamp()(*string) {
    return m.last_rate_limited_timestamp
}
// GetLastRequestTimestamp gets the last_request_timestamp property value. The last_request_timestamp property
// returns a *string when successful
func (m *ApiInsightsRouteStats) GetLastRequestTimestamp()(*string) {
    return m.last_request_timestamp
}
// GetRateLimitedRequestCount gets the rate_limited_request_count property value. The total number of requests that were rate limited within the queried time period
// returns a *int64 when successful
func (m *ApiInsightsRouteStats) GetRateLimitedRequestCount()(*int64) {
    return m.rate_limited_request_count
}
// GetTotalRequestCount gets the total_request_count property value. The total number of requests within the queried time period
// returns a *int64 when successful
func (m *ApiInsightsRouteStats) GetTotalRequestCount()(*int64) {
    return m.total_request_count
}
// Serialize serializes information the current object
func (m *ApiInsightsRouteStats) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("api_route", m.GetApiRoute())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("http_method", m.GetHttpMethod())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("last_rate_limited_timestamp", m.GetLastRateLimitedTimestamp())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("last_request_timestamp", m.GetLastRequestTimestamp())
        if err != nil {
            return err
        }
    }
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
func (m *ApiInsightsRouteStats) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetApiRoute sets the api_route property value. The API path's route template
func (m *ApiInsightsRouteStats) SetApiRoute(value *string)() {
    m.api_route = value
}
// SetHttpMethod sets the http_method property value. The HTTP method
func (m *ApiInsightsRouteStats) SetHttpMethod(value *string)() {
    m.http_method = value
}
// SetLastRateLimitedTimestamp sets the last_rate_limited_timestamp property value. The last_rate_limited_timestamp property
func (m *ApiInsightsRouteStats) SetLastRateLimitedTimestamp(value *string)() {
    m.last_rate_limited_timestamp = value
}
// SetLastRequestTimestamp sets the last_request_timestamp property value. The last_request_timestamp property
func (m *ApiInsightsRouteStats) SetLastRequestTimestamp(value *string)() {
    m.last_request_timestamp = value
}
// SetRateLimitedRequestCount sets the rate_limited_request_count property value. The total number of requests that were rate limited within the queried time period
func (m *ApiInsightsRouteStats) SetRateLimitedRequestCount(value *int64)() {
    m.rate_limited_request_count = value
}
// SetTotalRequestCount sets the total_request_count property value. The total number of requests within the queried time period
func (m *ApiInsightsRouteStats) SetTotalRequestCount(value *int64)() {
    m.total_request_count = value
}
type ApiInsightsRouteStatsable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetApiRoute()(*string)
    GetHttpMethod()(*string)
    GetLastRateLimitedTimestamp()(*string)
    GetLastRequestTimestamp()(*string)
    GetRateLimitedRequestCount()(*int64)
    GetTotalRequestCount()(*int64)
    SetApiRoute(value *string)()
    SetHttpMethod(value *string)()
    SetLastRateLimitedTimestamp(value *string)()
    SetLastRequestTimestamp(value *string)()
    SetRateLimitedRequestCount(value *int64)()
    SetTotalRequestCount(value *int64)()
}
