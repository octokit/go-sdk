package models
// Which type of organization repositories have access to the private registry.
type OrgPrivateRegistryConfiguration_visibility int

const (
    ALL_ORGPRIVATEREGISTRYCONFIGURATION_VISIBILITY OrgPrivateRegistryConfiguration_visibility = iota
    PRIVATE_ORGPRIVATEREGISTRYCONFIGURATION_VISIBILITY
    SELECTED_ORGPRIVATEREGISTRYCONFIGURATION_VISIBILITY
)

func (i OrgPrivateRegistryConfiguration_visibility) String() string {
    return []string{"all", "private", "selected"}[i]
}
func ParseOrgPrivateRegistryConfiguration_visibility(v string) (any, error) {
    result := ALL_ORGPRIVATEREGISTRYCONFIGURATION_VISIBILITY
    switch v {
        case "all":
            result = ALL_ORGPRIVATEREGISTRYCONFIGURATION_VISIBILITY
        case "private":
            result = PRIVATE_ORGPRIVATEREGISTRYCONFIGURATION_VISIBILITY
        case "selected":
            result = SELECTED_ORGPRIVATEREGISTRYCONFIGURATION_VISIBILITY
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeOrgPrivateRegistryConfiguration_visibility(values []OrgPrivateRegistryConfiguration_visibility) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i OrgPrivateRegistryConfiguration_visibility) isMultiValue() bool {
    return false
}
