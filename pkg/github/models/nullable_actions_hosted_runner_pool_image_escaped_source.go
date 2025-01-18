package models
// The image provider.
type NullableActionsHostedRunnerPoolImage_source int

const (
    GITHUB_NULLABLEACTIONSHOSTEDRUNNERPOOLIMAGE_SOURCE NullableActionsHostedRunnerPoolImage_source = iota
    PARTNER_NULLABLEACTIONSHOSTEDRUNNERPOOLIMAGE_SOURCE
    CUSTOM_NULLABLEACTIONSHOSTEDRUNNERPOOLIMAGE_SOURCE
)

func (i NullableActionsHostedRunnerPoolImage_source) String() string {
    return []string{"github", "partner", "custom"}[i]
}
func ParseNullableActionsHostedRunnerPoolImage_source(v string) (any, error) {
    result := GITHUB_NULLABLEACTIONSHOSTEDRUNNERPOOLIMAGE_SOURCE
    switch v {
        case "github":
            result = GITHUB_NULLABLEACTIONSHOSTEDRUNNERPOOLIMAGE_SOURCE
        case "partner":
            result = PARTNER_NULLABLEACTIONSHOSTEDRUNNERPOOLIMAGE_SOURCE
        case "custom":
            result = CUSTOM_NULLABLEACTIONSHOSTEDRUNNERPOOLIMAGE_SOURCE
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeNullableActionsHostedRunnerPoolImage_source(values []NullableActionsHostedRunnerPoolImage_source) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i NullableActionsHostedRunnerPoolImage_source) isMultiValue() bool {
    return false
}
