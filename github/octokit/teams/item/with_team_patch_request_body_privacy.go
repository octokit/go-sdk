package item
import (
    "errors"
)
// The level of privacy this team should have. Editing teams without specifying this parameter leaves `privacy` intact. The options are:  **For a non-nested team:**   * `secret` - only visible to organization owners and members of this team.   * `closed` - visible to all members of this organization.  **For a parent or child team:**   * `closed` - visible to all members of this organization.
type WithTeam_PatchRequestBody_privacy int

const (
    SECRET_WITHTEAM_PATCHREQUESTBODY_PRIVACY WithTeam_PatchRequestBody_privacy = iota
    CLOSED_WITHTEAM_PATCHREQUESTBODY_PRIVACY
)

func (i WithTeam_PatchRequestBody_privacy) String() string {
    return []string{"secret", "closed"}[i]
}
func ParseWithTeam_PatchRequestBody_privacy(v string) (any, error) {
    result := SECRET_WITHTEAM_PATCHREQUESTBODY_PRIVACY
    switch v {
        case "secret":
            result = SECRET_WITHTEAM_PATCHREQUESTBODY_PRIVACY
        case "closed":
            result = CLOSED_WITHTEAM_PATCHREQUESTBODY_PRIVACY
        default:
            return 0, errors.New("Unknown WithTeam_PatchRequestBody_privacy value: " + v)
    }
    return &result, nil
}
func SerializeWithTeam_PatchRequestBody_privacy(values []WithTeam_PatchRequestBody_privacy) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i WithTeam_PatchRequestBody_privacy) isMultiValue() bool {
    return false
}
