package item
// The baseline permission that all organization members have on this project
type WithProject_PatchRequestBody_organization_permission int

const (
    READ_WITHPROJECT_PATCHREQUESTBODY_ORGANIZATION_PERMISSION WithProject_PatchRequestBody_organization_permission = iota
    WRITE_WITHPROJECT_PATCHREQUESTBODY_ORGANIZATION_PERMISSION
    ADMIN_WITHPROJECT_PATCHREQUESTBODY_ORGANIZATION_PERMISSION
    NONE_WITHPROJECT_PATCHREQUESTBODY_ORGANIZATION_PERMISSION
)

func (i WithProject_PatchRequestBody_organization_permission) String() string {
    return []string{"read", "write", "admin", "none"}[i]
}
func ParseWithProject_PatchRequestBody_organization_permission(v string) (any, error) {
    result := READ_WITHPROJECT_PATCHREQUESTBODY_ORGANIZATION_PERMISSION
    switch v {
        case "read":
            result = READ_WITHPROJECT_PATCHREQUESTBODY_ORGANIZATION_PERMISSION
        case "write":
            result = WRITE_WITHPROJECT_PATCHREQUESTBODY_ORGANIZATION_PERMISSION
        case "admin":
            result = ADMIN_WITHPROJECT_PATCHREQUESTBODY_ORGANIZATION_PERMISSION
        case "none":
            result = NONE_WITHPROJECT_PATCHREQUESTBODY_ORGANIZATION_PERMISSION
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeWithProject_PatchRequestBody_organization_permission(values []WithProject_PatchRequestBody_organization_permission) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i WithProject_PatchRequestBody_organization_permission) isMultiValue() bool {
    return false
}
