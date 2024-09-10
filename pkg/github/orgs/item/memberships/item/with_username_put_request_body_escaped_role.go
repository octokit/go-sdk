package item
// The role to give the user in the organization. Can be one of:   * `admin` - The user will become an owner of the organization.   * `member` - The user will become a non-owner member of the organization.
type WithUsernamePutRequestBody_role int

const (
    ADMIN_WITHUSERNAMEPUTREQUESTBODY_ROLE WithUsernamePutRequestBody_role = iota
    MEMBER_WITHUSERNAMEPUTREQUESTBODY_ROLE
)

func (i WithUsernamePutRequestBody_role) String() string {
    return []string{"admin", "member"}[i]
}
func ParseWithUsernamePutRequestBody_role(v string) (any, error) {
    result := ADMIN_WITHUSERNAMEPUTREQUESTBODY_ROLE
    switch v {
        case "admin":
            result = ADMIN_WITHUSERNAMEPUTREQUESTBODY_ROLE
        case "member":
            result = MEMBER_WITHUSERNAMEPUTREQUESTBODY_ROLE
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
