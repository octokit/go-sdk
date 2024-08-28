package models
// When set to ALLGREEN, the merge commit created by merge queue for each PR in the group must pass all required checks to merge. When set to HEADGREEN, only the commit at the head of the merge group, i.e. the commit containing changes from all of the PRs in the group, must pass its required checks to merge.
type RepositoryRuleMergeQueue_parameters_grouping_strategy int

const (
    ALLGREEN_REPOSITORYRULEMERGEQUEUE_PARAMETERS_GROUPING_STRATEGY RepositoryRuleMergeQueue_parameters_grouping_strategy = iota
    HEADGREEN_REPOSITORYRULEMERGEQUEUE_PARAMETERS_GROUPING_STRATEGY
)

func (i RepositoryRuleMergeQueue_parameters_grouping_strategy) String() string {
    return []string{"ALLGREEN", "HEADGREEN"}[i]
}
func ParseRepositoryRuleMergeQueue_parameters_grouping_strategy(v string) (any, error) {
    result := ALLGREEN_REPOSITORYRULEMERGEQUEUE_PARAMETERS_GROUPING_STRATEGY
    switch v {
        case "ALLGREEN":
            result = ALLGREEN_REPOSITORYRULEMERGEQUEUE_PARAMETERS_GROUPING_STRATEGY
        case "HEADGREEN":
            result = HEADGREEN_REPOSITORYRULEMERGEQUEUE_PARAMETERS_GROUPING_STRATEGY
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeRepositoryRuleMergeQueue_parameters_grouping_strategy(values []RepositoryRuleMergeQueue_parameters_grouping_strategy) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i RepositoryRuleMergeQueue_parameters_grouping_strategy) isMultiValue() bool {
    return false
}
