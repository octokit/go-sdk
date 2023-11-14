package merge
import (
    "errors"
)
// The merge method to use.
type MergePutRequestBody_merge_method int

const (
    MERGE_MERGEPUTREQUESTBODY_MERGE_METHOD MergePutRequestBody_merge_method = iota
    SQUASH_MERGEPUTREQUESTBODY_MERGE_METHOD
    REBASE_MERGEPUTREQUESTBODY_MERGE_METHOD
)

func (i MergePutRequestBody_merge_method) String() string {
    return []string{"merge", "squash", "rebase"}[i]
}
func ParseMergePutRequestBody_merge_method(v string) (any, error) {
    result := MERGE_MERGEPUTREQUESTBODY_MERGE_METHOD
    switch v {
        case "merge":
            result = MERGE_MERGEPUTREQUESTBODY_MERGE_METHOD
        case "squash":
            result = SQUASH_MERGEPUTREQUESTBODY_MERGE_METHOD
        case "rebase":
            result = REBASE_MERGEPUTREQUESTBODY_MERGE_METHOD
        default:
            return 0, errors.New("Unknown MergePutRequestBody_merge_method value: " + v)
    }
    return &result, nil
}
func SerializeMergePutRequestBody_merge_method(values []MergePutRequestBody_merge_method) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i MergePutRequestBody_merge_method) isMultiValue() bool {
    return false
}
