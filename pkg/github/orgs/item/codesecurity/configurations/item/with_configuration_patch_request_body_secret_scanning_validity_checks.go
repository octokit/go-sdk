package item
import (
    "errors"
)
// The enablement status of secret scanning validity checks
type WithConfiguration_PatchRequestBody_secret_scanning_validity_checks int

const (
    ENABLED_WITHCONFIGURATION_PATCHREQUESTBODY_SECRET_SCANNING_VALIDITY_CHECKS WithConfiguration_PatchRequestBody_secret_scanning_validity_checks = iota
    DISABLED_WITHCONFIGURATION_PATCHREQUESTBODY_SECRET_SCANNING_VALIDITY_CHECKS
    NOT_SET_WITHCONFIGURATION_PATCHREQUESTBODY_SECRET_SCANNING_VALIDITY_CHECKS
)

func (i WithConfiguration_PatchRequestBody_secret_scanning_validity_checks) String() string {
    return []string{"enabled", "disabled", "not_set"}[i]
}
func ParseWithConfiguration_PatchRequestBody_secret_scanning_validity_checks(v string) (any, error) {
    result := ENABLED_WITHCONFIGURATION_PATCHREQUESTBODY_SECRET_SCANNING_VALIDITY_CHECKS
    switch v {
        case "enabled":
            result = ENABLED_WITHCONFIGURATION_PATCHREQUESTBODY_SECRET_SCANNING_VALIDITY_CHECKS
        case "disabled":
            result = DISABLED_WITHCONFIGURATION_PATCHREQUESTBODY_SECRET_SCANNING_VALIDITY_CHECKS
        case "not_set":
            result = NOT_SET_WITHCONFIGURATION_PATCHREQUESTBODY_SECRET_SCANNING_VALIDITY_CHECKS
        default:
            return 0, errors.New("Unknown WithConfiguration_PatchRequestBody_secret_scanning_validity_checks value: " + v)
    }
    return &result, nil
}
func SerializeWithConfiguration_PatchRequestBody_secret_scanning_validity_checks(values []WithConfiguration_PatchRequestBody_secret_scanning_validity_checks) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i WithConfiguration_PatchRequestBody_secret_scanning_validity_checks) isMultiValue() bool {
    return false
}
