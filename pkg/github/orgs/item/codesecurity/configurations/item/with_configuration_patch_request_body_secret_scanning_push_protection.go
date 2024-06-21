package item
import (
    "errors"
)
// The enablement status of secret scanning push protection
type WithConfiguration_PatchRequestBody_secret_scanning_push_protection int

const (
    ENABLED_WITHCONFIGURATION_PATCHREQUESTBODY_SECRET_SCANNING_PUSH_PROTECTION WithConfiguration_PatchRequestBody_secret_scanning_push_protection = iota
    DISABLED_WITHCONFIGURATION_PATCHREQUESTBODY_SECRET_SCANNING_PUSH_PROTECTION
    NOT_SET_WITHCONFIGURATION_PATCHREQUESTBODY_SECRET_SCANNING_PUSH_PROTECTION
)

func (i WithConfiguration_PatchRequestBody_secret_scanning_push_protection) String() string {
    return []string{"enabled", "disabled", "not_set"}[i]
}
func ParseWithConfiguration_PatchRequestBody_secret_scanning_push_protection(v string) (any, error) {
    result := ENABLED_WITHCONFIGURATION_PATCHREQUESTBODY_SECRET_SCANNING_PUSH_PROTECTION
    switch v {
        case "enabled":
            result = ENABLED_WITHCONFIGURATION_PATCHREQUESTBODY_SECRET_SCANNING_PUSH_PROTECTION
        case "disabled":
            result = DISABLED_WITHCONFIGURATION_PATCHREQUESTBODY_SECRET_SCANNING_PUSH_PROTECTION
        case "not_set":
            result = NOT_SET_WITHCONFIGURATION_PATCHREQUESTBODY_SECRET_SCANNING_PUSH_PROTECTION
        default:
            return 0, errors.New("Unknown WithConfiguration_PatchRequestBody_secret_scanning_push_protection value: " + v)
    }
    return &result, nil
}
func SerializeWithConfiguration_PatchRequestBody_secret_scanning_push_protection(values []WithConfiguration_PatchRequestBody_secret_scanning_push_protection) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i WithConfiguration_PatchRequestBody_secret_scanning_push_protection) isMultiValue() bool {
    return false
}
