package models
// Which type of organization repositories have access to the private registry. `selected` means only the repositories specified by `selected_repository_ids` can access the private registry.
type OrgPrivateRegistryConfigurationWithSelectedRepositories_visibility int

const (
    ALL_ORGPRIVATEREGISTRYCONFIGURATIONWITHSELECTEDREPOSITORIES_VISIBILITY OrgPrivateRegistryConfigurationWithSelectedRepositories_visibility = iota
    PRIVATE_ORGPRIVATEREGISTRYCONFIGURATIONWITHSELECTEDREPOSITORIES_VISIBILITY
    SELECTED_ORGPRIVATEREGISTRYCONFIGURATIONWITHSELECTEDREPOSITORIES_VISIBILITY
)

func (i OrgPrivateRegistryConfigurationWithSelectedRepositories_visibility) String() string {
    return []string{"all", "private", "selected"}[i]
}
func ParseOrgPrivateRegistryConfigurationWithSelectedRepositories_visibility(v string) (any, error) {
    result := ALL_ORGPRIVATEREGISTRYCONFIGURATIONWITHSELECTEDREPOSITORIES_VISIBILITY
    switch v {
        case "all":
            result = ALL_ORGPRIVATEREGISTRYCONFIGURATIONWITHSELECTEDREPOSITORIES_VISIBILITY
        case "private":
            result = PRIVATE_ORGPRIVATEREGISTRYCONFIGURATIONWITHSELECTEDREPOSITORIES_VISIBILITY
        case "selected":
            result = SELECTED_ORGPRIVATEREGISTRYCONFIGURATIONWITHSELECTEDREPOSITORIES_VISIBILITY
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeOrgPrivateRegistryConfigurationWithSelectedRepositories_visibility(values []OrgPrivateRegistryConfigurationWithSelectedRepositories_visibility) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i OrgPrivateRegistryConfigurationWithSelectedRepositories_visibility) isMultiValue() bool {
    return false
}
