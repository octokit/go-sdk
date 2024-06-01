package item
import (
    "errors"
)
// The permission to grant the team on this repository. If no permission is specified, the team's `permission` attribute will be used to determine what permission to grant the team on this repository.
type WithRepoPutRequestBody_permission int

const (
    PULL_WITHREPOPUTREQUESTBODY_PERMISSION WithRepoPutRequestBody_permission = iota
    PUSH_WITHREPOPUTREQUESTBODY_PERMISSION
    ADMIN_WITHREPOPUTREQUESTBODY_PERMISSION
)

func (i WithRepoPutRequestBody_permission) String() string {
    return []string{"pull", "push", "admin"}[i]
}
func ParseWithRepoPutRequestBody_permission(v string) (any, error) {
    result := PULL_WITHREPOPUTREQUESTBODY_PERMISSION
    switch v {
        case "pull":
            result = PULL_WITHREPOPUTREQUESTBODY_PERMISSION
        case "push":
            result = PUSH_WITHREPOPUTREQUESTBODY_PERMISSION
        case "admin":
            result = ADMIN_WITHREPOPUTREQUESTBODY_PERMISSION
        default:
            return 0, errors.New("Unknown WithRepoPutRequestBody_permission value: " + v)
    }
    return &result, nil
}
func SerializeWithRepoPutRequestBody_permission(values []WithRepoPutRequestBody_permission) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i WithRepoPutRequestBody_permission) isMultiValue() bool {
    return false
}
