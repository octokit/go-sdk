package models
// Determines if the team has a direct, indirect, or mixed relationship to a role
type TeamRoleAssignment_assignment int

const (
    DIRECT_TEAMROLEASSIGNMENT_ASSIGNMENT TeamRoleAssignment_assignment = iota
    INDIRECT_TEAMROLEASSIGNMENT_ASSIGNMENT
    MIXED_TEAMROLEASSIGNMENT_ASSIGNMENT
)

func (i TeamRoleAssignment_assignment) String() string {
    return []string{"direct", "indirect", "mixed"}[i]
}
func ParseTeamRoleAssignment_assignment(v string) (any, error) {
    result := DIRECT_TEAMROLEASSIGNMENT_ASSIGNMENT
    switch v {
        case "direct":
            result = DIRECT_TEAMROLEASSIGNMENT_ASSIGNMENT
        case "indirect":
            result = INDIRECT_TEAMROLEASSIGNMENT_ASSIGNMENT
        case "mixed":
            result = MIXED_TEAMROLEASSIGNMENT_ASSIGNMENT
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeTeamRoleAssignment_assignment(values []TeamRoleAssignment_assignment) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i TeamRoleAssignment_assignment) isMultiValue() bool {
    return false
}
