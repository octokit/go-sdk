package models
// Method to use when merging changes from queued pull requests.
type RepositoryRuleMergeQueue_parameters_merge_method int

const (
    MERGE_REPOSITORYRULEMERGEQUEUE_PARAMETERS_MERGE_METHOD RepositoryRuleMergeQueue_parameters_merge_method = iota
    SQUASH_REPOSITORYRULEMERGEQUEUE_PARAMETERS_MERGE_METHOD
    REBASE_REPOSITORYRULEMERGEQUEUE_PARAMETERS_MERGE_METHOD
)

func (i RepositoryRuleMergeQueue_parameters_merge_method) String() string {
    return []string{"MERGE", "SQUASH", "REBASE"}[i]
}
func ParseRepositoryRuleMergeQueue_parameters_merge_method(v string) (any, error) {
    result := MERGE_REPOSITORYRULEMERGEQUEUE_PARAMETERS_MERGE_METHOD
    switch v {
        case "MERGE":
            result = MERGE_REPOSITORYRULEMERGEQUEUE_PARAMETERS_MERGE_METHOD
        case "SQUASH":
            result = SQUASH_REPOSITORYRULEMERGEQUEUE_PARAMETERS_MERGE_METHOD
        case "REBASE":
            result = REBASE_REPOSITORYRULEMERGEQUEUE_PARAMETERS_MERGE_METHOD
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeRepositoryRuleMergeQueue_parameters_merge_method(values []RepositoryRuleMergeQueue_parameters_merge_method) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i RepositoryRuleMergeQueue_parameters_merge_method) isMultiValue() bool {
    return false
}
