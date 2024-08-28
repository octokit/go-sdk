package item
// **Deprecated**. The permission that new repositories will be added to the team with when none is specified.
type WithTeam_slugPatchRequestBody_permission int

const (
    PULL_WITHTEAM_SLUGPATCHREQUESTBODY_PERMISSION WithTeam_slugPatchRequestBody_permission = iota
    PUSH_WITHTEAM_SLUGPATCHREQUESTBODY_PERMISSION
    ADMIN_WITHTEAM_SLUGPATCHREQUESTBODY_PERMISSION
)

func (i WithTeam_slugPatchRequestBody_permission) String() string {
    return []string{"pull", "push", "admin"}[i]
}
func ParseWithTeam_slugPatchRequestBody_permission(v string) (any, error) {
    result := PULL_WITHTEAM_SLUGPATCHREQUESTBODY_PERMISSION
    switch v {
        case "pull":
            result = PULL_WITHTEAM_SLUGPATCHREQUESTBODY_PERMISSION
        case "push":
            result = PUSH_WITHTEAM_SLUGPATCHREQUESTBODY_PERMISSION
        case "admin":
            result = ADMIN_WITHTEAM_SLUGPATCHREQUESTBODY_PERMISSION
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeWithTeam_slugPatchRequestBody_permission(values []WithTeam_slugPatchRequestBody_permission) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i WithTeam_slugPatchRequestBody_permission) isMultiValue() bool {
    return false
}
