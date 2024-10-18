package runnergroups
// Visibility of a runner group. You can select all repositories, select individual repositories, or limit access to private repositories.
type RunnerGroupsPostRequestBody_visibility int

const (
    SELECTED_RUNNERGROUPSPOSTREQUESTBODY_VISIBILITY RunnerGroupsPostRequestBody_visibility = iota
    ALL_RUNNERGROUPSPOSTREQUESTBODY_VISIBILITY
    PRIVATE_RUNNERGROUPSPOSTREQUESTBODY_VISIBILITY
)

func (i RunnerGroupsPostRequestBody_visibility) String() string {
    return []string{"selected", "all", "private"}[i]
}
func ParseRunnerGroupsPostRequestBody_visibility(v string) (any, error) {
    result := SELECTED_RUNNERGROUPSPOSTREQUESTBODY_VISIBILITY
    switch v {
        case "selected":
            result = SELECTED_RUNNERGROUPSPOSTREQUESTBODY_VISIBILITY
        case "all":
            result = ALL_RUNNERGROUPSPOSTREQUESTBODY_VISIBILITY
        case "private":
            result = PRIVATE_RUNNERGROUPSPOSTREQUESTBODY_VISIBILITY
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeRunnerGroupsPostRequestBody_visibility(values []RunnerGroupsPostRequestBody_visibility) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i RunnerGroupsPostRequestBody_visibility) isMultiValue() bool {
    return false
}
