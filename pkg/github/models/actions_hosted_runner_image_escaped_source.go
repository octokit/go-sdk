package models
// The image provider.
type ActionsHostedRunnerImage_source int

const (
    GITHUB_ACTIONSHOSTEDRUNNERIMAGE_SOURCE ActionsHostedRunnerImage_source = iota
    PARTNER_ACTIONSHOSTEDRUNNERIMAGE_SOURCE
    CUSTOM_ACTIONSHOSTEDRUNNERIMAGE_SOURCE
)

func (i ActionsHostedRunnerImage_source) String() string {
    return []string{"github", "partner", "custom"}[i]
}
func ParseActionsHostedRunnerImage_source(v string) (any, error) {
    result := GITHUB_ACTIONSHOSTEDRUNNERIMAGE_SOURCE
    switch v {
        case "github":
            result = GITHUB_ACTIONSHOSTEDRUNNERIMAGE_SOURCE
        case "partner":
            result = PARTNER_ACTIONSHOSTEDRUNNERIMAGE_SOURCE
        case "custom":
            result = CUSTOM_ACTIONSHOSTEDRUNNERIMAGE_SOURCE
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeActionsHostedRunnerImage_source(values []ActionsHostedRunnerImage_source) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i ActionsHostedRunnerImage_source) isMultiValue() bool {
    return false
}
