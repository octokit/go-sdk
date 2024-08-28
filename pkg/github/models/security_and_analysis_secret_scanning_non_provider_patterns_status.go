package models
type SecurityAndAnalysis_secret_scanning_non_provider_patterns_status int

const (
    ENABLED_SECURITYANDANALYSIS_SECRET_SCANNING_NON_PROVIDER_PATTERNS_STATUS SecurityAndAnalysis_secret_scanning_non_provider_patterns_status = iota
    DISABLED_SECURITYANDANALYSIS_SECRET_SCANNING_NON_PROVIDER_PATTERNS_STATUS
)

func (i SecurityAndAnalysis_secret_scanning_non_provider_patterns_status) String() string {
    return []string{"enabled", "disabled"}[i]
}
func ParseSecurityAndAnalysis_secret_scanning_non_provider_patterns_status(v string) (any, error) {
    result := ENABLED_SECURITYANDANALYSIS_SECRET_SCANNING_NON_PROVIDER_PATTERNS_STATUS
    switch v {
        case "enabled":
            result = ENABLED_SECURITYANDANALYSIS_SECRET_SCANNING_NON_PROVIDER_PATTERNS_STATUS
        case "disabled":
            result = DISABLED_SECURITYANDANALYSIS_SECRET_SCANNING_NON_PROVIDER_PATTERNS_STATUS
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeSecurityAndAnalysis_secret_scanning_non_provider_patterns_status(values []SecurityAndAnalysis_secret_scanning_non_provider_patterns_status) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i SecurityAndAnalysis_secret_scanning_non_provider_patterns_status) isMultiValue() bool {
    return false
}
