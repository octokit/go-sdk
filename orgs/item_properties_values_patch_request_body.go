package orgs

import (
    i33bcfa37677dfd8b71eaf22ad3bfd140ba1702e7a8418558499c22418bf2fef2 "github.com/octokit/go-sdk/models"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemPropertiesValuesPatchRequestBody 
type ItemPropertiesValuesPatchRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // List of custom property names and associated values to apply to the repositories.
    properties []i33bcfa37677dfd8b71eaf22ad3bfd140ba1702e7a8418558499c22418bf2fef2.CustomPropertyValueable
    // The names of repositories that the custom property values will be applied to.
    repository_names []string
}
// NewItemPropertiesValuesPatchRequestBody instantiates a new ItemPropertiesValuesPatchRequestBody and sets the default values.
func NewItemPropertiesValuesPatchRequestBody()(*ItemPropertiesValuesPatchRequestBody) {
    m := &ItemPropertiesValuesPatchRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateItemPropertiesValuesPatchRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemPropertiesValuesPatchRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemPropertiesValuesPatchRequestBody(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ItemPropertiesValuesPatchRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ItemPropertiesValuesPatchRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["properties"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(i33bcfa37677dfd8b71eaf22ad3bfd140ba1702e7a8418558499c22418bf2fef2.CreateCustomPropertyValueFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]i33bcfa37677dfd8b71eaf22ad3bfd140ba1702e7a8418558499c22418bf2fef2.CustomPropertyValueable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(i33bcfa37677dfd8b71eaf22ad3bfd140ba1702e7a8418558499c22418bf2fef2.CustomPropertyValueable)
                }
            }
            m.SetProperties(res)
        }
        return nil
    }
    res["repository_names"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = *(v.(*string))
                }
            }
            m.SetRepositoryNames(res)
        }
        return nil
    }
    return res
}
// GetProperties gets the properties property value. List of custom property names and associated values to apply to the repositories.
func (m *ItemPropertiesValuesPatchRequestBody) GetProperties()([]i33bcfa37677dfd8b71eaf22ad3bfd140ba1702e7a8418558499c22418bf2fef2.CustomPropertyValueable) {
    return m.properties
}
// GetRepositoryNames gets the repository_names property value. The names of repositories that the custom property values will be applied to.
func (m *ItemPropertiesValuesPatchRequestBody) GetRepositoryNames()([]string) {
    return m.repository_names
}
// Serialize serializes information the current object
func (m *ItemPropertiesValuesPatchRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetProperties() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetProperties()))
        for i, v := range m.GetProperties() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("properties", cast)
        if err != nil {
            return err
        }
    }
    if m.GetRepositoryNames() != nil {
        err := writer.WriteCollectionOfStringValues("repository_names", m.GetRepositoryNames())
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
func (m *ItemPropertiesValuesPatchRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetProperties sets the properties property value. List of custom property names and associated values to apply to the repositories.
func (m *ItemPropertiesValuesPatchRequestBody) SetProperties(value []i33bcfa37677dfd8b71eaf22ad3bfd140ba1702e7a8418558499c22418bf2fef2.CustomPropertyValueable)() {
    m.properties = value
}
// SetRepositoryNames sets the repository_names property value. The names of repositories that the custom property values will be applied to.
func (m *ItemPropertiesValuesPatchRequestBody) SetRepositoryNames(value []string)() {
    m.repository_names = value
}
// ItemPropertiesValuesPatchRequestBodyable 
type ItemPropertiesValuesPatchRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetProperties()([]i33bcfa37677dfd8b71eaf22ad3bfd140ba1702e7a8418558499c22418bf2fef2.CustomPropertyValueable)
    GetRepositoryNames()([]string)
    SetProperties(value []i33bcfa37677dfd8b71eaf22ad3bfd140ba1702e7a8418558499c22418bf2fef2.CustomPropertyValueable)()
    SetRepositoryNames(value []string)()
}
