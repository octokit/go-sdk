package item
import (
    "errors"
)
// The level of privacy this team should have. Editing teams without specifying this parameter leaves `privacy` intact. When a team is nested, the `privacy` for parent teams cannot be `secret`. The options are:  **For a non-nested team:**   * `secret` - only visible to organization owners and members of this team.   * `closed` - visible to all members of this organization.  **For a parent or child team:**   * `closed` - visible to all members of this organization.
type WithTeam_slugPatchRequestBody_privacy int

const (
    SECRET_WITHTEAM_SLUGPATCHREQUESTBODY_PRIVACY WithTeam_slugPatchRequestBody_privacy = iota
    CLOSED_WITHTEAM_SLUGPATCHREQUESTBODY_PRIVACY
)

func (i WithTeam_slugPatchRequestBody_privacy) String() string {
    return []string{"secret", "closed"}[i]
}
func ParseWithTeam_slugPatchRequestBody_privacy(v string) (any, error) {
    result := SECRET_WITHTEAM_SLUGPATCHREQUESTBODY_PRIVACY
    switch v {
        case "secret":
            result = SECRET_WITHTEAM_SLUGPATCHREQUESTBODY_PRIVACY
        case "closed":
            result = CLOSED_WITHTEAM_SLUGPATCHREQUESTBODY_PRIVACY
        default:
            return 0, errors.New("Unknown WithTeam_slugPatchRequestBody_privacy value: " + v)
    }
    return &result, nil
}
func SerializeWithTeam_slugPatchRequestBody_privacy(values []WithTeam_slugPatchRequestBody_privacy) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i WithTeam_slugPatchRequestBody_privacy) isMultiValue() bool {
    return false
}
