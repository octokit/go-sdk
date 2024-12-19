package models
// Determines if the user has a direct, indirect, or mixed relationship to a role
type UserRoleAssignment_assignment int

const (
    DIRECT_USERROLEASSIGNMENT_ASSIGNMENT UserRoleAssignment_assignment = iota
    INDIRECT_USERROLEASSIGNMENT_ASSIGNMENT
    MIXED_USERROLEASSIGNMENT_ASSIGNMENT
)

func (i UserRoleAssignment_assignment) String() string {
    return []string{"direct", "indirect", "mixed"}[i]
}
func ParseUserRoleAssignment_assignment(v string) (any, error) {
    result := DIRECT_USERROLEASSIGNMENT_ASSIGNMENT
    switch v {
        case "direct":
            result = DIRECT_USERROLEASSIGNMENT_ASSIGNMENT
        case "indirect":
            result = INDIRECT_USERROLEASSIGNMENT_ASSIGNMENT
        case "mixed":
            result = MIXED_USERROLEASSIGNMENT_ASSIGNMENT
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeUserRoleAssignment_assignment(values []UserRoleAssignment_assignment) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i UserRoleAssignment_assignment) isMultiValue() bool {
    return false
}
