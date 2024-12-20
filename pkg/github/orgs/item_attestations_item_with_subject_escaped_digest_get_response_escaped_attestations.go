package orgs

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type ItemAttestationsItemWithSubject_digestGetResponse_attestations struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The attestation's Sigstore Bundle.Refer to the [Sigstore Bundle Specification](https://github.com/sigstore/protobuf-specs/blob/main/protos/sigstore_bundle.proto) for more information.
    bundle ItemAttestationsItemWithSubject_digestGetResponse_attestations_bundleable
    // The bundle_url property
    bundle_url *string
    // The repository_id property
    repository_id *int32
}
// NewItemAttestationsItemWithSubject_digestGetResponse_attestations instantiates a new ItemAttestationsItemWithSubject_digestGetResponse_attestations and sets the default values.
func NewItemAttestationsItemWithSubject_digestGetResponse_attestations()(*ItemAttestationsItemWithSubject_digestGetResponse_attestations) {
    m := &ItemAttestationsItemWithSubject_digestGetResponse_attestations{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateItemAttestationsItemWithSubject_digestGetResponse_attestationsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateItemAttestationsItemWithSubject_digestGetResponse_attestationsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemAttestationsItemWithSubject_digestGetResponse_attestations(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *ItemAttestationsItemWithSubject_digestGetResponse_attestations) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetBundle gets the bundle property value. The attestation's Sigstore Bundle.Refer to the [Sigstore Bundle Specification](https://github.com/sigstore/protobuf-specs/blob/main/protos/sigstore_bundle.proto) for more information.
// returns a ItemAttestationsItemWithSubject_digestGetResponse_attestations_bundleable when successful
func (m *ItemAttestationsItemWithSubject_digestGetResponse_attestations) GetBundle()(ItemAttestationsItemWithSubject_digestGetResponse_attestations_bundleable) {
    return m.bundle
}
// GetBundleUrl gets the bundle_url property value. The bundle_url property
// returns a *string when successful
func (m *ItemAttestationsItemWithSubject_digestGetResponse_attestations) GetBundleUrl()(*string) {
    return m.bundle_url
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ItemAttestationsItemWithSubject_digestGetResponse_attestations) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["bundle"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateItemAttestationsItemWithSubject_digestGetResponse_attestations_bundleFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBundle(val.(ItemAttestationsItemWithSubject_digestGetResponse_attestations_bundleable))
        }
        return nil
    }
    res["bundle_url"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBundleUrl(val)
        }
        return nil
    }
    res["repository_id"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRepositoryId(val)
        }
        return nil
    }
    return res
}
// GetRepositoryId gets the repository_id property value. The repository_id property
// returns a *int32 when successful
func (m *ItemAttestationsItemWithSubject_digestGetResponse_attestations) GetRepositoryId()(*int32) {
    return m.repository_id
}
// Serialize serializes information the current object
func (m *ItemAttestationsItemWithSubject_digestGetResponse_attestations) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("bundle", m.GetBundle())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("bundle_url", m.GetBundleUrl())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("repository_id", m.GetRepositoryId())
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
func (m *ItemAttestationsItemWithSubject_digestGetResponse_attestations) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetBundle sets the bundle property value. The attestation's Sigstore Bundle.Refer to the [Sigstore Bundle Specification](https://github.com/sigstore/protobuf-specs/blob/main/protos/sigstore_bundle.proto) for more information.
func (m *ItemAttestationsItemWithSubject_digestGetResponse_attestations) SetBundle(value ItemAttestationsItemWithSubject_digestGetResponse_attestations_bundleable)() {
    m.bundle = value
}
// SetBundleUrl sets the bundle_url property value. The bundle_url property
func (m *ItemAttestationsItemWithSubject_digestGetResponse_attestations) SetBundleUrl(value *string)() {
    m.bundle_url = value
}
// SetRepositoryId sets the repository_id property value. The repository_id property
func (m *ItemAttestationsItemWithSubject_digestGetResponse_attestations) SetRepositoryId(value *int32)() {
    m.repository_id = value
}
type ItemAttestationsItemWithSubject_digestGetResponse_attestationsable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBundle()(ItemAttestationsItemWithSubject_digestGetResponse_attestations_bundleable)
    GetBundleUrl()(*string)
    GetRepositoryId()(*int32)
    SetBundle(value ItemAttestationsItemWithSubject_digestGetResponse_attestations_bundleable)()
    SetBundleUrl(value *string)()
    SetRepositoryId(value *int32)()
}
