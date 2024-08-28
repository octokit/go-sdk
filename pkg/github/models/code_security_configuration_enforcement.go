package models
// The enforcement status for a security configuration
type CodeSecurityConfiguration_enforcement int

const (
    ENFORCED_CODESECURITYCONFIGURATION_ENFORCEMENT CodeSecurityConfiguration_enforcement = iota
    UNENFORCED_CODESECURITYCONFIGURATION_ENFORCEMENT
)

func (i CodeSecurityConfiguration_enforcement) String() string {
    return []string{"enforced", "unenforced"}[i]
}
func ParseCodeSecurityConfiguration_enforcement(v string) (any, error) {
    result := ENFORCED_CODESECURITYCONFIGURATION_ENFORCEMENT
    switch v {
        case "enforced":
            result = ENFORCED_CODESECURITYCONFIGURATION_ENFORCEMENT
        case "unenforced":
            result = UNENFORCED_CODESECURITYCONFIGURATION_ENFORCEMENT
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeCodeSecurityConfiguration_enforcement(values []CodeSecurityConfiguration_enforcement) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i CodeSecurityConfiguration_enforcement) isMultiValue() bool {
    return false
}
