package teams
// **Closing down notice**. The permission that new repositories will be added to the team with when none is specified.
type TeamsPostRequestBody_permission int

const (
    PULL_TEAMSPOSTREQUESTBODY_PERMISSION TeamsPostRequestBody_permission = iota
    PUSH_TEAMSPOSTREQUESTBODY_PERMISSION
)

func (i TeamsPostRequestBody_permission) String() string {
    return []string{"pull", "push"}[i]
}
func ParseTeamsPostRequestBody_permission(v string) (any, error) {
    result := PULL_TEAMSPOSTREQUESTBODY_PERMISSION
    switch v {
        case "pull":
            result = PULL_TEAMSPOSTREQUESTBODY_PERMISSION
        case "push":
            result = PUSH_TEAMSPOSTREQUESTBODY_PERMISSION
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeTeamsPostRequestBody_permission(values []TeamsPostRequestBody_permission) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i TeamsPostRequestBody_permission) isMultiValue() bool {
    return false
}
