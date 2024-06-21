package item
import (
    "errors"
)
// The enablement status of secret scanning
type WithConfiguration_PatchRequestBody_secret_scanning int

const (
    ENABLED_WITHCONFIGURATION_PATCHREQUESTBODY_SECRET_SCANNING WithConfiguration_PatchRequestBody_secret_scanning = iota
    DISABLED_WITHCONFIGURATION_PATCHREQUESTBODY_SECRET_SCANNING
    NOT_SET_WITHCONFIGURATION_PATCHREQUESTBODY_SECRET_SCANNING
)

func (i WithConfiguration_PatchRequestBody_secret_scanning) String() string {
    return []string{"enabled", "disabled", "not_set"}[i]
}
func ParseWithConfiguration_PatchRequestBody_secret_scanning(v string) (any, error) {
    result := ENABLED_WITHCONFIGURATION_PATCHREQUESTBODY_SECRET_SCANNING
    switch v {
        case "enabled":
            result = ENABLED_WITHCONFIGURATION_PATCHREQUESTBODY_SECRET_SCANNING
        case "disabled":
            result = DISABLED_WITHCONFIGURATION_PATCHREQUESTBODY_SECRET_SCANNING
        case "not_set":
            result = NOT_SET_WITHCONFIGURATION_PATCHREQUESTBODY_SECRET_SCANNING
        default:
            return 0, errors.New("Unknown WithConfiguration_PatchRequestBody_secret_scanning value: " + v)
    }
    return &result, nil
}
func SerializeWithConfiguration_PatchRequestBody_secret_scanning(values []WithConfiguration_PatchRequestBody_secret_scanning) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i WithConfiguration_PatchRequestBody_secret_scanning) isMultiValue() bool {
    return false
}
