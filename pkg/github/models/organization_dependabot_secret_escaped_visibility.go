package models
// Visibility of a secret
type OrganizationDependabotSecret_visibility int

const (
    ALL_ORGANIZATIONDEPENDABOTSECRET_VISIBILITY OrganizationDependabotSecret_visibility = iota
    PRIVATE_ORGANIZATIONDEPENDABOTSECRET_VISIBILITY
    SELECTED_ORGANIZATIONDEPENDABOTSECRET_VISIBILITY
)

func (i OrganizationDependabotSecret_visibility) String() string {
    return []string{"all", "private", "selected"}[i]
}
func ParseOrganizationDependabotSecret_visibility(v string) (any, error) {
    result := ALL_ORGANIZATIONDEPENDABOTSECRET_VISIBILITY
    switch v {
        case "all":
            result = ALL_ORGANIZATIONDEPENDABOTSECRET_VISIBILITY
        case "private":
            result = PRIVATE_ORGANIZATIONDEPENDABOTSECRET_VISIBILITY
        case "selected":
            result = SELECTED_ORGANIZATIONDEPENDABOTSECRET_VISIBILITY
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeOrganizationDependabotSecret_visibility(values []OrganizationDependabotSecret_visibility) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i OrganizationDependabotSecret_visibility) isMultiValue() bool {
    return false
}
