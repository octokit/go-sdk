package models
type SecurityAndAnalysis_secret_scanning_ai_detection_status int

const (
    ENABLED_SECURITYANDANALYSIS_SECRET_SCANNING_AI_DETECTION_STATUS SecurityAndAnalysis_secret_scanning_ai_detection_status = iota
    DISABLED_SECURITYANDANALYSIS_SECRET_SCANNING_AI_DETECTION_STATUS
)

func (i SecurityAndAnalysis_secret_scanning_ai_detection_status) String() string {
    return []string{"enabled", "disabled"}[i]
}
func ParseSecurityAndAnalysis_secret_scanning_ai_detection_status(v string) (any, error) {
    result := ENABLED_SECURITYANDANALYSIS_SECRET_SCANNING_AI_DETECTION_STATUS
    switch v {
        case "enabled":
            result = ENABLED_SECURITYANDANALYSIS_SECRET_SCANNING_AI_DETECTION_STATUS
        case "disabled":
            result = DISABLED_SECURITYANDANALYSIS_SECRET_SCANNING_AI_DETECTION_STATUS
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeSecurityAndAnalysis_secret_scanning_ai_detection_status(values []SecurityAndAnalysis_secret_scanning_ai_detection_status) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i SecurityAndAnalysis_secret_scanning_ai_detection_status) isMultiValue() bool {
    return false
}
