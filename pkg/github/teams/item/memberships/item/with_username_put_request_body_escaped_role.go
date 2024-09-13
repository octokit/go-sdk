package item
// The role that this user should have in the team.
type WithUsernamePutRequestBody_role int

const (
    MEMBER_WITHUSERNAMEPUTREQUESTBODY_ROLE WithUsernamePutRequestBody_role = iota
    MAINTAINER_WITHUSERNAMEPUTREQUESTBODY_ROLE
)

func (i WithUsernamePutRequestBody_role) String() string {
    return []string{"member", "maintainer"}[i]
}
func ParseWithUsernamePutRequestBody_role(v string) (any, error) {
    result := MEMBER_WITHUSERNAMEPUTREQUESTBODY_ROLE
    switch v {
        case "member":
            result = MEMBER_WITHUSERNAMEPUTREQUESTBODY_ROLE
        case "maintainer":
            result = MAINTAINER_WITHUSERNAMEPUTREQUESTBODY_ROLE
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeWithUsernamePutRequestBody_role(values []WithUsernamePutRequestBody_role) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i WithUsernamePutRequestBody_role) isMultiValue() bool {
    return false
}
