package repos

import (
    i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e "github.com/octokit/go-sdk/github/octokit/models"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemItemHooksPostRequestBody_config key/value pairs to provide settings for this webhook.
type ItemItemHooksPostRequestBody_config struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The media type used to serialize the payloads. Supported values include `json` and `form`. The default is `form`.
    content_type *string
    // The digest property
    digest *string
    // The insecure_ssl property
    insecure_ssl i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.WebhookConfigInsecureSslable
    // If provided, the `secret` will be used as the `key` to generate the HMAC hex digest value for [delivery signature headers](https://docs.github.com/webhooks/event-payloads/#delivery-headers).
    secret *string
    // The token property
    token *string
    // The URL to which the payloads will be delivered.
    url *string
}
// NewItemItemHooksPostRequestBody_config instantiates a new ItemItemHooksPostRequestBody_config and sets the default values.
func NewItemItemHooksPostRequestBody_config()(*ItemItemHooksPostRequestBody_config) {
    m := &ItemItemHooksPostRequestBody_config{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateItemItemHooksPostRequestBody_configFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemItemHooksPostRequestBody_configFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemItemHooksPostRequestBody_config(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ItemItemHooksPostRequestBody_config) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetContentType gets the content_type property value. The media type used to serialize the payloads. Supported values include `json` and `form`. The default is `form`.
func (m *ItemItemHooksPostRequestBody_config) GetContentType()(*string) {
    return m.content_type
}
// GetDigest gets the digest property value. The digest property
func (m *ItemItemHooksPostRequestBody_config) GetDigest()(*string) {
    return m.digest
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ItemItemHooksPostRequestBody_config) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["content_type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetContentType(val)
        }
        return nil
    }
    res["digest"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDigest(val)
        }
        return nil
    }
    res["insecure_ssl"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.CreateWebhookConfigInsecureSslFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetInsecureSsl(val.(i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.WebhookConfigInsecureSslable))
        }
        return nil
    }
    res["secret"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSecret(val)
        }
        return nil
    }
    res["token"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetToken(val)
        }
        return nil
    }
    res["url"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUrl(val)
        }
        return nil
    }
    return res
}
// GetInsecureSsl gets the insecure_ssl property value. The insecure_ssl property
func (m *ItemItemHooksPostRequestBody_config) GetInsecureSsl()(i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.WebhookConfigInsecureSslable) {
    return m.insecure_ssl
}
// GetSecret gets the secret property value. If provided, the `secret` will be used as the `key` to generate the HMAC hex digest value for [delivery signature headers](https://docs.github.com/webhooks/event-payloads/#delivery-headers).
func (m *ItemItemHooksPostRequestBody_config) GetSecret()(*string) {
    return m.secret
}
// GetToken gets the token property value. The token property
func (m *ItemItemHooksPostRequestBody_config) GetToken()(*string) {
    return m.token
}
// GetUrl gets the url property value. The URL to which the payloads will be delivered.
func (m *ItemItemHooksPostRequestBody_config) GetUrl()(*string) {
    return m.url
}
// Serialize serializes information the current object
func (m *ItemItemHooksPostRequestBody_config) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("content_type", m.GetContentType())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("digest", m.GetDigest())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("insecure_ssl", m.GetInsecureSsl())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("secret", m.GetSecret())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("token", m.GetToken())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("url", m.GetUrl())
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
func (m *ItemItemHooksPostRequestBody_config) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetContentType sets the content_type property value. The media type used to serialize the payloads. Supported values include `json` and `form`. The default is `form`.
func (m *ItemItemHooksPostRequestBody_config) SetContentType(value *string)() {
    m.content_type = value
}
// SetDigest sets the digest property value. The digest property
func (m *ItemItemHooksPostRequestBody_config) SetDigest(value *string)() {
    m.digest = value
}
// SetInsecureSsl sets the insecure_ssl property value. The insecure_ssl property
func (m *ItemItemHooksPostRequestBody_config) SetInsecureSsl(value i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.WebhookConfigInsecureSslable)() {
    m.insecure_ssl = value
}
// SetSecret sets the secret property value. If provided, the `secret` will be used as the `key` to generate the HMAC hex digest value for [delivery signature headers](https://docs.github.com/webhooks/event-payloads/#delivery-headers).
func (m *ItemItemHooksPostRequestBody_config) SetSecret(value *string)() {
    m.secret = value
}
// SetToken sets the token property value. The token property
func (m *ItemItemHooksPostRequestBody_config) SetToken(value *string)() {
    m.token = value
}
// SetUrl sets the url property value. The URL to which the payloads will be delivered.
func (m *ItemItemHooksPostRequestBody_config) SetUrl(value *string)() {
    m.url = value
}
// ItemItemHooksPostRequestBody_configable 
type ItemItemHooksPostRequestBody_configable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetContentType()(*string)
    GetDigest()(*string)
    GetInsecureSsl()(i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.WebhookConfigInsecureSslable)
    GetSecret()(*string)
    GetToken()(*string)
    GetUrl()(*string)
    SetContentType(value *string)()
    SetDigest(value *string)()
    SetInsecureSsl(value i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.WebhookConfigInsecureSslable)()
    SetSecret(value *string)()
    SetToken(value *string)()
    SetUrl(value *string)()
}
