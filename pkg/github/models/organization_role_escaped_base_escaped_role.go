package models
// The system role from which this role inherits permissions.
type OrganizationRole_base_role int

const (
    READ_ORGANIZATIONROLE_BASE_ROLE OrganizationRole_base_role = iota
    TRIAGE_ORGANIZATIONROLE_BASE_ROLE
    WRITE_ORGANIZATIONROLE_BASE_ROLE
    MAINTAIN_ORGANIZATIONROLE_BASE_ROLE
    ADMIN_ORGANIZATIONROLE_BASE_ROLE
)

func (i OrganizationRole_base_role) String() string {
    return []string{"read", "triage", "write", "maintain", "admin"}[i]
}
func ParseOrganizationRole_base_role(v string) (any, error) {
    result := READ_ORGANIZATIONROLE_BASE_ROLE
    switch v {
        case "read":
            result = READ_ORGANIZATIONROLE_BASE_ROLE
        case "triage":
            result = TRIAGE_ORGANIZATIONROLE_BASE_ROLE
        case "write":
            result = WRITE_ORGANIZATIONROLE_BASE_ROLE
        case "maintain":
            result = MAINTAIN_ORGANIZATIONROLE_BASE_ROLE
        case "admin":
            result = ADMIN_ORGANIZATIONROLE_BASE_ROLE
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeOrganizationRole_base_role(values []OrganizationRole_base_role) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i OrganizationRole_base_role) isMultiValue() bool {
    return false
}
