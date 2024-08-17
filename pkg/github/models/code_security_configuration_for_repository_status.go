package models
import (
    "errors"
)
// The attachment status of the code security configuration on the repository.
type CodeSecurityConfigurationForRepository_status int

const (
    ATTACHED_CODESECURITYCONFIGURATIONFORREPOSITORY_STATUS CodeSecurityConfigurationForRepository_status = iota
    ATTACHING_CODESECURITYCONFIGURATIONFORREPOSITORY_STATUS
    DETACHED_CODESECURITYCONFIGURATIONFORREPOSITORY_STATUS
    REMOVED_CODESECURITYCONFIGURATIONFORREPOSITORY_STATUS
    ENFORCED_CODESECURITYCONFIGURATIONFORREPOSITORY_STATUS
    FAILED_CODESECURITYCONFIGURATIONFORREPOSITORY_STATUS
    UPDATING_CODESECURITYCONFIGURATIONFORREPOSITORY_STATUS
    REMOVED_BY_ENTERPRISE_CODESECURITYCONFIGURATIONFORREPOSITORY_STATUS
)

func (i CodeSecurityConfigurationForRepository_status) String() string {
    return []string{"attached", "attaching", "detached", "removed", "enforced", "failed", "updating", "removed_by_enterprise"}[i]
}
func ParseCodeSecurityConfigurationForRepository_status(v string) (any, error) {
    result := ATTACHED_CODESECURITYCONFIGURATIONFORREPOSITORY_STATUS
    switch v {
        case "attached":
            result = ATTACHED_CODESECURITYCONFIGURATIONFORREPOSITORY_STATUS
        case "attaching":
            result = ATTACHING_CODESECURITYCONFIGURATIONFORREPOSITORY_STATUS
        case "detached":
            result = DETACHED_CODESECURITYCONFIGURATIONFORREPOSITORY_STATUS
        case "removed":
            result = REMOVED_CODESECURITYCONFIGURATIONFORREPOSITORY_STATUS
        case "enforced":
            result = ENFORCED_CODESECURITYCONFIGURATIONFORREPOSITORY_STATUS
        case "failed":
            result = FAILED_CODESECURITYCONFIGURATIONFORREPOSITORY_STATUS
        case "updating":
            result = UPDATING_CODESECURITYCONFIGURATIONFORREPOSITORY_STATUS
        case "removed_by_enterprise":
            result = REMOVED_BY_ENTERPRISE_CODESECURITYCONFIGURATIONFORREPOSITORY_STATUS
        default:
            return 0, errors.New("Unknown CodeSecurityConfigurationForRepository_status value: " + v)
    }
    return &result, nil
}
func SerializeCodeSecurityConfigurationForRepository_status(values []CodeSecurityConfigurationForRepository_status) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i CodeSecurityConfigurationForRepository_status) isMultiValue() bool {
    return false
}
