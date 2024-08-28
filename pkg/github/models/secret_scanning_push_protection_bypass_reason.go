package models
// The reason for bypassing push protection.
type SecretScanningPushProtectionBypassReason int

const (
    FALSE_POSITIVE_SECRETSCANNINGPUSHPROTECTIONBYPASSREASON SecretScanningPushProtectionBypassReason = iota
    USED_IN_TESTS_SECRETSCANNINGPUSHPROTECTIONBYPASSREASON
    WILL_FIX_LATER_SECRETSCANNINGPUSHPROTECTIONBYPASSREASON
)

func (i SecretScanningPushProtectionBypassReason) String() string {
    return []string{"false_positive", "used_in_tests", "will_fix_later"}[i]
}
func ParseSecretScanningPushProtectionBypassReason(v string) (any, error) {
    result := FALSE_POSITIVE_SECRETSCANNINGPUSHPROTECTIONBYPASSREASON
    switch v {
        case "false_positive":
            result = FALSE_POSITIVE_SECRETSCANNINGPUSHPROTECTIONBYPASSREASON
        case "used_in_tests":
            result = USED_IN_TESTS_SECRETSCANNINGPUSHPROTECTIONBYPASSREASON
        case "will_fix_later":
            result = WILL_FIX_LATER_SECRETSCANNINGPUSHPROTECTIONBYPASSREASON
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeSecretScanningPushProtectionBypassReason(values []SecretScanningPushProtectionBypassReason) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i SecretScanningPushProtectionBypassReason) isMultiValue() bool {
    return false
}
