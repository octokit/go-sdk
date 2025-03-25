package models
type RepositoryRulePullRequest_parameters_allowed_merge_methods int

const (
    MERGE_REPOSITORYRULEPULLREQUEST_PARAMETERS_ALLOWED_MERGE_METHODS RepositoryRulePullRequest_parameters_allowed_merge_methods = iota
    SQUASH_REPOSITORYRULEPULLREQUEST_PARAMETERS_ALLOWED_MERGE_METHODS
    REBASE_REPOSITORYRULEPULLREQUEST_PARAMETERS_ALLOWED_MERGE_METHODS
)

func (i RepositoryRulePullRequest_parameters_allowed_merge_methods) String() string {
    return []string{"merge", "squash", "rebase"}[i]
}
func ParseRepositoryRulePullRequest_parameters_allowed_merge_methods(v string) (any, error) {
    result := MERGE_REPOSITORYRULEPULLREQUEST_PARAMETERS_ALLOWED_MERGE_METHODS
    switch v {
        case "merge":
            result = MERGE_REPOSITORYRULEPULLREQUEST_PARAMETERS_ALLOWED_MERGE_METHODS
        case "squash":
            result = SQUASH_REPOSITORYRULEPULLREQUEST_PARAMETERS_ALLOWED_MERGE_METHODS
        case "rebase":
            result = REBASE_REPOSITORYRULEPULLREQUEST_PARAMETERS_ALLOWED_MERGE_METHODS
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeRepositoryRulePullRequest_parameters_allowed_merge_methods(values []RepositoryRulePullRequest_parameters_allowed_merge_methods) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i RepositoryRulePullRequest_parameters_allowed_merge_methods) isMultiValue() bool {
    return false
}
