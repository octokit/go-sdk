package item
import (
    "errors"
)
// The permission to grant to the team for this project. Default: the team's `permission` attribute will be used to determine what permission to grant the team on this project. Note that, if you choose not to pass any parameters, you'll need to set `Content-Length` to zero when calling this endpoint. For more information, see "[HTTP verbs](https://docs.github.com/rest/overview/resources-in-the-rest-api#http-verbs)."
type WithProject_PutRequestBody_permission int

const (
    READ_WITHPROJECT_PUTREQUESTBODY_PERMISSION WithProject_PutRequestBody_permission = iota
    WRITE_WITHPROJECT_PUTREQUESTBODY_PERMISSION
    ADMIN_WITHPROJECT_PUTREQUESTBODY_PERMISSION
)

func (i WithProject_PutRequestBody_permission) String() string {
    return []string{"read", "write", "admin"}[i]
}
func ParseWithProject_PutRequestBody_permission(v string) (any, error) {
    result := READ_WITHPROJECT_PUTREQUESTBODY_PERMISSION
    switch v {
        case "read":
            result = READ_WITHPROJECT_PUTREQUESTBODY_PERMISSION
        case "write":
            result = WRITE_WITHPROJECT_PUTREQUESTBODY_PERMISSION
        case "admin":
            result = ADMIN_WITHPROJECT_PUTREQUESTBODY_PERMISSION
        default:
            return 0, errors.New("Unknown WithProject_PutRequestBody_permission value: " + v)
    }
    return &result, nil
}
func SerializeWithProject_PutRequestBody_permission(values []WithProject_PutRequestBody_permission) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i WithProject_PutRequestBody_permission) isMultiValue() bool {
    return false
}
