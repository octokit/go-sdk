package models
type RepositoryRuleMergeQueue_type int

const (
    MERGE_QUEUE_REPOSITORYRULEMERGEQUEUE_TYPE RepositoryRuleMergeQueue_type = iota
)

func (i RepositoryRuleMergeQueue_type) String() string {
    return []string{"merge_queue"}[i]
}
func ParseRepositoryRuleMergeQueue_type(v string) (any, error) {
    result := MERGE_QUEUE_REPOSITORYRULEMERGEQUEUE_TYPE
    switch v {
        case "merge_queue":
            result = MERGE_QUEUE_REPOSITORYRULEMERGEQUEUE_TYPE
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeRepositoryRuleMergeQueue_type(values []RepositoryRuleMergeQueue_type) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i RepositoryRuleMergeQueue_type) isMultiValue() bool {
    return false
}
