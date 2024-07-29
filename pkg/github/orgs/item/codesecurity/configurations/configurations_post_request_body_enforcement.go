package configurations
import (
    "errors"
)
// The enforcement status for a security configuration
type ConfigurationsPostRequestBody_enforcement int

const (
    ENFORCED_CONFIGURATIONSPOSTREQUESTBODY_ENFORCEMENT ConfigurationsPostRequestBody_enforcement = iota
    UNENFORCED_CONFIGURATIONSPOSTREQUESTBODY_ENFORCEMENT
)

func (i ConfigurationsPostRequestBody_enforcement) String() string {
    return []string{"enforced", "unenforced"}[i]
}
func ParseConfigurationsPostRequestBody_enforcement(v string) (any, error) {
    result := ENFORCED_CONFIGURATIONSPOSTREQUESTBODY_ENFORCEMENT
    switch v {
        case "enforced":
            result = ENFORCED_CONFIGURATIONSPOSTREQUESTBODY_ENFORCEMENT
        case "unenforced":
            result = UNENFORCED_CONFIGURATIONSPOSTREQUESTBODY_ENFORCEMENT
        default:
            return 0, errors.New("Unknown ConfigurationsPostRequestBody_enforcement value: " + v)
    }
    return &result, nil
}
func SerializeConfigurationsPostRequestBody_enforcement(values []ConfigurationsPostRequestBody_enforcement) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i ConfigurationsPostRequestBody_enforcement) isMultiValue() bool {
    return false
}
