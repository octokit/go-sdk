package teams
import (
    "errors"
)
// The level of privacy this team should have. The options are:  **For a non-nested team:**   * `secret` - only visible to organization owners and members of this team.   * `closed` - visible to all members of this organization.  Default: `secret`  **For a parent or child team:**   * `closed` - visible to all members of this organization.  Default for child team: `closed`
type TeamsPostRequestBody_privacy int

const (
    SECRET_TEAMSPOSTREQUESTBODY_PRIVACY TeamsPostRequestBody_privacy = iota
    CLOSED_TEAMSPOSTREQUESTBODY_PRIVACY
)

func (i TeamsPostRequestBody_privacy) String() string {
    return []string{"secret", "closed"}[i]
}
func ParseTeamsPostRequestBody_privacy(v string) (any, error) {
    result := SECRET_TEAMSPOSTREQUESTBODY_PRIVACY
    switch v {
        case "secret":
            result = SECRET_TEAMSPOSTREQUESTBODY_PRIVACY
        case "closed":
            result = CLOSED_TEAMSPOSTREQUESTBODY_PRIVACY
        default:
            return 0, errors.New("Unknown TeamsPostRequestBody_privacy value: " + v)
    }
    return &result, nil
}
func SerializeTeamsPostRequestBody_privacy(values []TeamsPostRequestBody_privacy) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i TeamsPostRequestBody_privacy) isMultiValue() bool {
    return false
}
