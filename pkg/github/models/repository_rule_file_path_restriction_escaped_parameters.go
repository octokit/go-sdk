package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type RepositoryRuleFilePathRestriction_parameters struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The file paths that are restricted from being pushed to the commit graph.
    restricted_file_paths []string
}
// NewRepositoryRuleFilePathRestriction_parameters instantiates a new RepositoryRuleFilePathRestriction_parameters and sets the default values.
func NewRepositoryRuleFilePathRestriction_parameters()(*RepositoryRuleFilePathRestriction_parameters) {
    m := &RepositoryRuleFilePathRestriction_parameters{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateRepositoryRuleFilePathRestriction_parametersFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateRepositoryRuleFilePathRestriction_parametersFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewRepositoryRuleFilePathRestriction_parameters(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *RepositoryRuleFilePathRestriction_parameters) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *RepositoryRuleFilePathRestriction_parameters) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["restricted_file_paths"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
            m.SetRestrictedFilePaths(res)
        }
        return nil
    }
    return res
}
// GetRestrictedFilePaths gets the restricted_file_paths property value. The file paths that are restricted from being pushed to the commit graph.
// returns a []string when successful
func (m *RepositoryRuleFilePathRestriction_parameters) GetRestrictedFilePaths()([]string) {
    return m.restricted_file_paths
}
// Serialize serializes information the current object
func (m *RepositoryRuleFilePathRestriction_parameters) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetRestrictedFilePaths() != nil {
        err := writer.WriteCollectionOfStringValues("restricted_file_paths", m.GetRestrictedFilePaths())
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
func (m *RepositoryRuleFilePathRestriction_parameters) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetRestrictedFilePaths sets the restricted_file_paths property value. The file paths that are restricted from being pushed to the commit graph.
func (m *RepositoryRuleFilePathRestriction_parameters) SetRestrictedFilePaths(value []string)() {
    m.restricted_file_paths = value
}
type RepositoryRuleFilePathRestriction_parametersable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetRestrictedFilePaths()([]string)
    SetRestrictedFilePaths(value []string)()
}
