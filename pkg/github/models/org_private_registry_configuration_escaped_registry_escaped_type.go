package models
// The registry type.
type OrgPrivateRegistryConfiguration_registry_type int

const (
    MAVEN_REPOSITORY_ORGPRIVATEREGISTRYCONFIGURATION_REGISTRY_TYPE OrgPrivateRegistryConfiguration_registry_type = iota
)

func (i OrgPrivateRegistryConfiguration_registry_type) String() string {
    return []string{"maven_repository"}[i]
}
func ParseOrgPrivateRegistryConfiguration_registry_type(v string) (any, error) {
    result := MAVEN_REPOSITORY_ORGPRIVATEREGISTRYCONFIGURATION_REGISTRY_TYPE
    switch v {
        case "maven_repository":
            result = MAVEN_REPOSITORY_ORGPRIVATEREGISTRYCONFIGURATION_REGISTRY_TYPE
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeOrgPrivateRegistryConfiguration_registry_type(values []OrgPrivateRegistryConfiguration_registry_type) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i OrgPrivateRegistryConfiguration_registry_type) isMultiValue() bool {
    return false
}
