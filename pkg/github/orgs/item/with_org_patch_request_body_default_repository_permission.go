package item
import (
    "errors"
)
// Default permission level members have for organization repositories.
type WithOrgPatchRequestBody_default_repository_permission int

const (
    READ_WITHORGPATCHREQUESTBODY_DEFAULT_REPOSITORY_PERMISSION WithOrgPatchRequestBody_default_repository_permission = iota
    WRITE_WITHORGPATCHREQUESTBODY_DEFAULT_REPOSITORY_PERMISSION
    ADMIN_WITHORGPATCHREQUESTBODY_DEFAULT_REPOSITORY_PERMISSION
    NONE_WITHORGPATCHREQUESTBODY_DEFAULT_REPOSITORY_PERMISSION
)

func (i WithOrgPatchRequestBody_default_repository_permission) String() string {
    return []string{"read", "write", "admin", "none"}[i]
}
func ParseWithOrgPatchRequestBody_default_repository_permission(v string) (any, error) {
    result := READ_WITHORGPATCHREQUESTBODY_DEFAULT_REPOSITORY_PERMISSION
    switch v {
        case "read":
            result = READ_WITHORGPATCHREQUESTBODY_DEFAULT_REPOSITORY_PERMISSION
        case "write":
            result = WRITE_WITHORGPATCHREQUESTBODY_DEFAULT_REPOSITORY_PERMISSION
        case "admin":
            result = ADMIN_WITHORGPATCHREQUESTBODY_DEFAULT_REPOSITORY_PERMISSION
        case "none":
            result = NONE_WITHORGPATCHREQUESTBODY_DEFAULT_REPOSITORY_PERMISSION
        default:
            return 0, errors.New("Unknown WithOrgPatchRequestBody_default_repository_permission value: " + v)
    }
    return &result, nil
}
func SerializeWithOrgPatchRequestBody_default_repository_permission(values []WithOrgPatchRequestBody_default_repository_permission) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i WithOrgPatchRequestBody_default_repository_permission) isMultiValue() bool {
    return false
}
