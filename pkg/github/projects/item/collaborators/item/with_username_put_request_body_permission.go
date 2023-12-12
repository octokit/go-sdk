package item
import (
    "errors"
)
// The permission to grant the collaborator.
type WithUsernamePutRequestBody_permission int

const (
    READ_WITHUSERNAMEPUTREQUESTBODY_PERMISSION WithUsernamePutRequestBody_permission = iota
    WRITE_WITHUSERNAMEPUTREQUESTBODY_PERMISSION
    ADMIN_WITHUSERNAMEPUTREQUESTBODY_PERMISSION
)

func (i WithUsernamePutRequestBody_permission) String() string {
    return []string{"read", "write", "admin"}[i]
}
func ParseWithUsernamePutRequestBody_permission(v string) (any, error) {
    result := READ_WITHUSERNAMEPUTREQUESTBODY_PERMISSION
    switch v {
        case "read":
            result = READ_WITHUSERNAMEPUTREQUESTBODY_PERMISSION
        case "write":
            result = WRITE_WITHUSERNAMEPUTREQUESTBODY_PERMISSION
        case "admin":
            result = ADMIN_WITHUSERNAMEPUTREQUESTBODY_PERMISSION
        default:
            return 0, errors.New("Unknown WithUsernamePutRequestBody_permission value: " + v)
    }
    return &result, nil
}
func SerializeWithUsernamePutRequestBody_permission(values []WithUsernamePutRequestBody_permission) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i WithUsernamePutRequestBody_permission) isMultiValue() bool {
    return false
}
