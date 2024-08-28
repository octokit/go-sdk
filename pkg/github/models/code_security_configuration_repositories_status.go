package models
// The attachment status of the code security configuration on the repository.
type CodeSecurityConfigurationRepositories_status int

const (
    ATTACHED_CODESECURITYCONFIGURATIONREPOSITORIES_STATUS CodeSecurityConfigurationRepositories_status = iota
    ATTACHING_CODESECURITYCONFIGURATIONREPOSITORIES_STATUS
    DETACHED_CODESECURITYCONFIGURATIONREPOSITORIES_STATUS
    REMOVED_CODESECURITYCONFIGURATIONREPOSITORIES_STATUS
    ENFORCED_CODESECURITYCONFIGURATIONREPOSITORIES_STATUS
    FAILED_CODESECURITYCONFIGURATIONREPOSITORIES_STATUS
    UPDATING_CODESECURITYCONFIGURATIONREPOSITORIES_STATUS
    REMOVED_BY_ENTERPRISE_CODESECURITYCONFIGURATIONREPOSITORIES_STATUS
)

func (i CodeSecurityConfigurationRepositories_status) String() string {
    return []string{"attached", "attaching", "detached", "removed", "enforced", "failed", "updating", "removed_by_enterprise"}[i]
}
func ParseCodeSecurityConfigurationRepositories_status(v string) (any, error) {
    result := ATTACHED_CODESECURITYCONFIGURATIONREPOSITORIES_STATUS
    switch v {
        case "attached":
            result = ATTACHED_CODESECURITYCONFIGURATIONREPOSITORIES_STATUS
        case "attaching":
            result = ATTACHING_CODESECURITYCONFIGURATIONREPOSITORIES_STATUS
        case "detached":
            result = DETACHED_CODESECURITYCONFIGURATIONREPOSITORIES_STATUS
        case "removed":
            result = REMOVED_CODESECURITYCONFIGURATIONREPOSITORIES_STATUS
        case "enforced":
            result = ENFORCED_CODESECURITYCONFIGURATIONREPOSITORIES_STATUS
        case "failed":
            result = FAILED_CODESECURITYCONFIGURATIONREPOSITORIES_STATUS
        case "updating":
            result = UPDATING_CODESECURITYCONFIGURATIONREPOSITORIES_STATUS
        case "removed_by_enterprise":
            result = REMOVED_BY_ENTERPRISE_CODESECURITYCONFIGURATIONREPOSITORIES_STATUS
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeCodeSecurityConfigurationRepositories_status(values []CodeSecurityConfigurationRepositories_status) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i CodeSecurityConfigurationRepositories_status) isMultiValue() bool {
    return false
}
