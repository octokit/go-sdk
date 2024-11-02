package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type ApiInsightsUserStats struct {
    // The actor_id property
    actor_id *int64
    // The actor_name property
    actor_name *string
    // The actor_type property
    actor_type *string
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The integration_id property
    integration_id *int64
    // The last_rate_limited_timestamp property
    last_rate_limited_timestamp *string
    // The last_request_timestamp property
    last_request_timestamp *string
    // The oauth_application_id property
    oauth_application_id *int64
    // The rate_limited_request_count property
    rate_limited_request_count *int32
    // The total_request_count property
    total_request_count *int32
}
// NewApiInsightsUserStats instantiates a new ApiInsightsUserStats and sets the default values.
func NewApiInsightsUserStats()(*ApiInsightsUserStats) {
    m := &ApiInsightsUserStats{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateApiInsightsUserStatsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateApiInsightsUserStatsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewApiInsightsUserStats(), nil
}
// GetActorId gets the actor_id property value. The actor_id property
// returns a *int64 when successful
func (m *ApiInsightsUserStats) GetActorId()(*int64) {
    return m.actor_id
}
// GetActorName gets the actor_name property value. The actor_name property
// returns a *string when successful
func (m *ApiInsightsUserStats) GetActorName()(*string) {
    return m.actor_name
}
// GetActorType gets the actor_type property value. The actor_type property
// returns a *string when successful
func (m *ApiInsightsUserStats) GetActorType()(*string) {
    return m.actor_type
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *ApiInsightsUserStats) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ApiInsightsUserStats) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["actor_id"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetActorId(val)
        }
        return nil
    }
    res["actor_name"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetActorName(val)
        }
        return nil
    }
    res["actor_type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetActorType(val)
        }
        return nil
    }
    res["integration_id"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIntegrationId(val)
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
    res["oauth_application_id"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOauthApplicationId(val)
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
// GetIntegrationId gets the integration_id property value. The integration_id property
// returns a *int64 when successful
func (m *ApiInsightsUserStats) GetIntegrationId()(*int64) {
    return m.integration_id
}
// GetLastRateLimitedTimestamp gets the last_rate_limited_timestamp property value. The last_rate_limited_timestamp property
// returns a *string when successful
func (m *ApiInsightsUserStats) GetLastRateLimitedTimestamp()(*string) {
    return m.last_rate_limited_timestamp
}
// GetLastRequestTimestamp gets the last_request_timestamp property value. The last_request_timestamp property
// returns a *string when successful
func (m *ApiInsightsUserStats) GetLastRequestTimestamp()(*string) {
    return m.last_request_timestamp
}
// GetOauthApplicationId gets the oauth_application_id property value. The oauth_application_id property
// returns a *int64 when successful
func (m *ApiInsightsUserStats) GetOauthApplicationId()(*int64) {
    return m.oauth_application_id
}
// GetRateLimitedRequestCount gets the rate_limited_request_count property value. The rate_limited_request_count property
// returns a *int32 when successful
func (m *ApiInsightsUserStats) GetRateLimitedRequestCount()(*int32) {
    return m.rate_limited_request_count
}
// GetTotalRequestCount gets the total_request_count property value. The total_request_count property
// returns a *int32 when successful
func (m *ApiInsightsUserStats) GetTotalRequestCount()(*int32) {
    return m.total_request_count
}
// Serialize serializes information the current object
func (m *ApiInsightsUserStats) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteInt64Value("actor_id", m.GetActorId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("actor_name", m.GetActorName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("actor_type", m.GetActorType())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt64Value("integration_id", m.GetIntegrationId())
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
        err := writer.WriteInt64Value("oauth_application_id", m.GetOauthApplicationId())
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
// SetActorId sets the actor_id property value. The actor_id property
func (m *ApiInsightsUserStats) SetActorId(value *int64)() {
    m.actor_id = value
}
// SetActorName sets the actor_name property value. The actor_name property
func (m *ApiInsightsUserStats) SetActorName(value *string)() {
    m.actor_name = value
}
// SetActorType sets the actor_type property value. The actor_type property
func (m *ApiInsightsUserStats) SetActorType(value *string)() {
    m.actor_type = value
}
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ApiInsightsUserStats) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetIntegrationId sets the integration_id property value. The integration_id property
func (m *ApiInsightsUserStats) SetIntegrationId(value *int64)() {
    m.integration_id = value
}
// SetLastRateLimitedTimestamp sets the last_rate_limited_timestamp property value. The last_rate_limited_timestamp property
func (m *ApiInsightsUserStats) SetLastRateLimitedTimestamp(value *string)() {
    m.last_rate_limited_timestamp = value
}
// SetLastRequestTimestamp sets the last_request_timestamp property value. The last_request_timestamp property
func (m *ApiInsightsUserStats) SetLastRequestTimestamp(value *string)() {
    m.last_request_timestamp = value
}
// SetOauthApplicationId sets the oauth_application_id property value. The oauth_application_id property
func (m *ApiInsightsUserStats) SetOauthApplicationId(value *int64)() {
    m.oauth_application_id = value
}
// SetRateLimitedRequestCount sets the rate_limited_request_count property value. The rate_limited_request_count property
func (m *ApiInsightsUserStats) SetRateLimitedRequestCount(value *int32)() {
    m.rate_limited_request_count = value
}
// SetTotalRequestCount sets the total_request_count property value. The total_request_count property
func (m *ApiInsightsUserStats) SetTotalRequestCount(value *int32)() {
    m.total_request_count = value
}
type ApiInsightsUserStatsable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetActorId()(*int64)
    GetActorName()(*string)
    GetActorType()(*string)
    GetIntegrationId()(*int64)
    GetLastRateLimitedTimestamp()(*string)
    GetLastRequestTimestamp()(*string)
    GetOauthApplicationId()(*int64)
    GetRateLimitedRequestCount()(*int32)
    GetTotalRequestCount()(*int32)
    SetActorId(value *int64)()
    SetActorName(value *string)()
    SetActorType(value *string)()
    SetIntegrationId(value *int64)()
    SetLastRateLimitedTimestamp(value *string)()
    SetLastRequestTimestamp(value *string)()
    SetOauthApplicationId(value *int64)()
    SetRateLimitedRequestCount(value *int32)()
    SetTotalRequestCount(value *int32)()
}
