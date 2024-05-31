package item
// **Deprecated**. The permission that new repositories will be added to the team with when none is specified.
type WithTeam_PatchRequestBody_permission int

const (
    PULL_WITHTEAM_PATCHREQUESTBODY_PERMISSION WithTeam_PatchRequestBody_permission = iota
    PUSH_WITHTEAM_PATCHREQUESTBODY_PERMISSION
    ADMIN_WITHTEAM_PATCHREQUESTBODY_PERMISSION
)

func (i WithTeam_PatchRequestBody_permission) String() string {
    return []string{"pull", "push", "admin"}[i]
}
func ParseWithTeam_PatchRequestBody_permission(v string) (any, error) {
    result := PULL_WITHTEAM_PATCHREQUESTBODY_PERMISSION
    switch v {
        case "pull":
            result = PULL_WITHTEAM_PATCHREQUESTBODY_PERMISSION
        case "push":
            result = PUSH_WITHTEAM_PATCHREQUESTBODY_PERMISSION
        case "admin":
            result = ADMIN_WITHTEAM_PATCHREQUESTBODY_PERMISSION
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeWithTeam_PatchRequestBody_permission(values []WithTeam_PatchRequestBody_permission) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i WithTeam_PatchRequestBody_permission) isMultiValue() bool {
    return false
}
