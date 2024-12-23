package models
// The registry type.
type OrgPrivateRegistryConfigurationWithSelectedRepositories_registry_type int

const (
    MAVEN_REPOSITORY_ORGPRIVATEREGISTRYCONFIGURATIONWITHSELECTEDREPOSITORIES_REGISTRY_TYPE OrgPrivateRegistryConfigurationWithSelectedRepositories_registry_type = iota
)

func (i OrgPrivateRegistryConfigurationWithSelectedRepositories_registry_type) String() string {
    return []string{"maven_repository"}[i]
}
func ParseOrgPrivateRegistryConfigurationWithSelectedRepositories_registry_type(v string) (any, error) {
    result := MAVEN_REPOSITORY_ORGPRIVATEREGISTRYCONFIGURATIONWITHSELECTEDREPOSITORIES_REGISTRY_TYPE
    switch v {
        case "maven_repository":
            result = MAVEN_REPOSITORY_ORGPRIVATEREGISTRYCONFIGURATIONWITHSELECTEDREPOSITORIES_REGISTRY_TYPE
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeOrgPrivateRegistryConfigurationWithSelectedRepositories_registry_type(values []OrgPrivateRegistryConfigurationWithSelectedRepositories_registry_type) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i OrgPrivateRegistryConfigurationWithSelectedRepositories_registry_type) isMultiValue() bool {
    return false
}
