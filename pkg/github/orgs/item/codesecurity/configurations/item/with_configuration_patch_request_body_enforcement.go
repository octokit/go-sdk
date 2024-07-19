package item
import (
    "errors"
)
// The status of enforcement
type WithConfiguration_PatchRequestBody_enforcement int

const (
    ENFORCED_WITHCONFIGURATION_PATCHREQUESTBODY_ENFORCEMENT WithConfiguration_PatchRequestBody_enforcement = iota
    UNENFORCED_WITHCONFIGURATION_PATCHREQUESTBODY_ENFORCEMENT
)

func (i WithConfiguration_PatchRequestBody_enforcement) String() string {
    return []string{"enforced", "unenforced"}[i]
}
func ParseWithConfiguration_PatchRequestBody_enforcement(v string) (any, error) {
    result := ENFORCED_WITHCONFIGURATION_PATCHREQUESTBODY_ENFORCEMENT
    switch v {
        case "enforced":
            result = ENFORCED_WITHCONFIGURATION_PATCHREQUESTBODY_ENFORCEMENT
        case "unenforced":
            result = UNENFORCED_WITHCONFIGURATION_PATCHREQUESTBODY_ENFORCEMENT
        default:
            return 0, errors.New("Unknown WithConfiguration_PatchRequestBody_enforcement value: " + v)
    }
    return &result, nil
}
func SerializeWithConfiguration_PatchRequestBody_enforcement(values []WithConfiguration_PatchRequestBody_enforcement) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i WithConfiguration_PatchRequestBody_enforcement) isMultiValue() bool {
    return false
}
