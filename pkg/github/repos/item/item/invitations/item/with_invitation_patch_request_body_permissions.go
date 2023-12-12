package item
import (
    "errors"
)
// The permissions that the associated user will have on the repository. Valid values are `read`, `write`, `maintain`, `triage`, and `admin`.
type WithInvitation_PatchRequestBody_permissions int

const (
    READ_WITHINVITATION_PATCHREQUESTBODY_PERMISSIONS WithInvitation_PatchRequestBody_permissions = iota
    WRITE_WITHINVITATION_PATCHREQUESTBODY_PERMISSIONS
    MAINTAIN_WITHINVITATION_PATCHREQUESTBODY_PERMISSIONS
    TRIAGE_WITHINVITATION_PATCHREQUESTBODY_PERMISSIONS
    ADMIN_WITHINVITATION_PATCHREQUESTBODY_PERMISSIONS
)

func (i WithInvitation_PatchRequestBody_permissions) String() string {
    return []string{"read", "write", "maintain", "triage", "admin"}[i]
}
func ParseWithInvitation_PatchRequestBody_permissions(v string) (any, error) {
    result := READ_WITHINVITATION_PATCHREQUESTBODY_PERMISSIONS
    switch v {
        case "read":
            result = READ_WITHINVITATION_PATCHREQUESTBODY_PERMISSIONS
        case "write":
            result = WRITE_WITHINVITATION_PATCHREQUESTBODY_PERMISSIONS
        case "maintain":
            result = MAINTAIN_WITHINVITATION_PATCHREQUESTBODY_PERMISSIONS
        case "triage":
            result = TRIAGE_WITHINVITATION_PATCHREQUESTBODY_PERMISSIONS
        case "admin":
            result = ADMIN_WITHINVITATION_PATCHREQUESTBODY_PERMISSIONS
        default:
            return 0, errors.New("Unknown WithInvitation_PatchRequestBody_permissions value: " + v)
    }
    return &result, nil
}
func SerializeWithInvitation_PatchRequestBody_permissions(values []WithInvitation_PatchRequestBody_permissions) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i WithInvitation_PatchRequestBody_permissions) isMultiValue() bool {
    return false
}
