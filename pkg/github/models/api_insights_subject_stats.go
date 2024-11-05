package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type ApiInsightsSubjectStats struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The last_rate_limited_timestamp property
    last_rate_limited_timestamp *string
    // The last_request_timestamp property
    last_request_timestamp *string
    // The rate_limited_request_count property
    rate_limited_request_count *int32
    // The subject_id property
    subject_id *int64
    // The subject_name property
    subject_name *string
    // The subject_type property
    subject_type *string
    // The total_request_count property
    total_request_count *int32
}
// NewApiInsightsSubjectStats instantiates a new ApiInsightsSubjectStats and sets the default values.
func NewApiInsightsSubjectStats()(*ApiInsightsSubjectStats) {
    m := &ApiInsightsSubjectStats{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateApiInsightsSubjectStatsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateApiInsightsSubjectStatsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewApiInsightsSubjectStats(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *ApiInsightsSubjectStats) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ApiInsightsSubjectStats) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
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
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRateLimitedRequestCount(val)
        }
        return nil
    }
    res["subject_id"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSubjectId(val)
        }
        return nil
    }
    res["subject_name"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSubjectName(val)
        }
        return nil
    }
    res["subject_type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSubjectType(val)
        }
        return nil
    }
    res["total_request_count"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
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
// GetLastRateLimitedTimestamp gets the last_rate_limited_timestamp property value. The last_rate_limited_timestamp property
// returns a *string when successful
func (m *ApiInsightsSubjectStats) GetLastRateLimitedTimestamp()(*string) {
    return m.last_rate_limited_timestamp
}
// GetLastRequestTimestamp gets the last_request_timestamp property value. The last_request_timestamp property
// returns a *string when successful
func (m *ApiInsightsSubjectStats) GetLastRequestTimestamp()(*string) {
    return m.last_request_timestamp
}
// GetRateLimitedRequestCount gets the rate_limited_request_count property value. The rate_limited_request_count property
// returns a *int32 when successful
func (m *ApiInsightsSubjectStats) GetRateLimitedRequestCount()(*int32) {
    return m.rate_limited_request_count
}
// GetSubjectId gets the subject_id property value. The subject_id property
// returns a *int64 when successful
func (m *ApiInsightsSubjectStats) GetSubjectId()(*int64) {
    return m.subject_id
}
// GetSubjectName gets the subject_name property value. The subject_name property
// returns a *string when successful
func (m *ApiInsightsSubjectStats) GetSubjectName()(*string) {
    return m.subject_name
}
// GetSubjectType gets the subject_type property value. The subject_type property
// returns a *string when successful
func (m *ApiInsightsSubjectStats) GetSubjectType()(*string) {
    return m.subject_type
}
// GetTotalRequestCount gets the total_request_count property value. The total_request_count property
// returns a *int32 when successful
func (m *ApiInsightsSubjectStats) GetTotalRequestCount()(*int32) {
    return m.total_request_count
}
// Serialize serializes information the current object
func (m *ApiInsightsSubjectStats) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
        err := writer.WriteInt32Value("rate_limited_request_count", m.GetRateLimitedRequestCount())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt64Value("subject_id", m.GetSubjectId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("subject_name", m.GetSubjectName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("subject_type", m.GetSubjectType())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("total_request_count", m.GetTotalRequestCount())
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
func (m *ApiInsightsSubjectStats) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetLastRateLimitedTimestamp sets the last_rate_limited_timestamp property value. The last_rate_limited_timestamp property
func (m *ApiInsightsSubjectStats) SetLastRateLimitedTimestamp(value *string)() {
    m.last_rate_limited_timestamp = value
}
// SetLastRequestTimestamp sets the last_request_timestamp property value. The last_request_timestamp property
func (m *ApiInsightsSubjectStats) SetLastRequestTimestamp(value *string)() {
    m.last_request_timestamp = value
}
// SetRateLimitedRequestCount sets the rate_limited_request_count property value. The rate_limited_request_count property
func (m *ApiInsightsSubjectStats) SetRateLimitedRequestCount(value *int32)() {
    m.rate_limited_request_count = value
}
// SetSubjectId sets the subject_id property value. The subject_id property
func (m *ApiInsightsSubjectStats) SetSubjectId(value *int64)() {
    m.subject_id = value
}
// SetSubjectName sets the subject_name property value. The subject_name property
func (m *ApiInsightsSubjectStats) SetSubjectName(value *string)() {
    m.subject_name = value
}
// SetSubjectType sets the subject_type property value. The subject_type property
func (m *ApiInsightsSubjectStats) SetSubjectType(value *string)() {
    m.subject_type = value
}
// SetTotalRequestCount sets the total_request_count property value. The total_request_count property
func (m *ApiInsightsSubjectStats) SetTotalRequestCount(value *int32)() {
    m.total_request_count = value
}
type ApiInsightsSubjectStatsable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetLastRateLimitedTimestamp()(*string)
    GetLastRequestTimestamp()(*string)
    GetRateLimitedRequestCount()(*int32)
    GetSubjectId()(*int64)
    GetSubjectName()(*string)
    GetSubjectType()(*string)
    GetTotalRequestCount()(*int32)
    SetLastRateLimitedTimestamp(value *string)()
    SetLastRequestTimestamp(value *string)()
    SetRateLimitedRequestCount(value *int32)()
    SetSubjectId(value *int64)()
    SetSubjectName(value *string)()
    SetSubjectType(value *string)()
    SetTotalRequestCount(value *int32)()
}
