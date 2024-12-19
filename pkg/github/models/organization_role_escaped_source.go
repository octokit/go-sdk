package models
// Source answers the question, "where did this role come from?"
type OrganizationRole_source int

const (
    ORGANIZATION_ORGANIZATIONROLE_SOURCE OrganizationRole_source = iota
    ENTERPRISE_ORGANIZATIONROLE_SOURCE
    PREDEFINED_ORGANIZATIONROLE_SOURCE
)

func (i OrganizationRole_source) String() string {
    return []string{"Organization", "Enterprise", "Predefined"}[i]
}
func ParseOrganizationRole_source(v string) (any, error) {
    result := ORGANIZATION_ORGANIZATIONROLE_SOURCE
    switch v {
        case "Organization":
            result = ORGANIZATION_ORGANIZATIONROLE_SOURCE
        case "Enterprise":
            result = ENTERPRISE_ORGANIZATIONROLE_SOURCE
        case "Predefined":
            result = PREDEFINED_ORGANIZATIONROLE_SOURCE
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeOrganizationRole_source(values []OrganizationRole_source) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i OrganizationRole_source) isMultiValue() bool {
    return false
}
