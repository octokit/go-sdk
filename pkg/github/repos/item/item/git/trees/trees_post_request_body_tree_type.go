package trees
import (
    "errors"
)
// Either `blob`, `tree`, or `commit`.
type TreesPostRequestBody_tree_type int

const (
    BLOB_TREESPOSTREQUESTBODY_TREE_TYPE TreesPostRequestBody_tree_type = iota
    TREE_TREESPOSTREQUESTBODY_TREE_TYPE
    COMMIT_TREESPOSTREQUESTBODY_TREE_TYPE
)

func (i TreesPostRequestBody_tree_type) String() string {
    return []string{"blob", "tree", "commit"}[i]
}
func ParseTreesPostRequestBody_tree_type(v string) (any, error) {
    result := BLOB_TREESPOSTREQUESTBODY_TREE_TYPE
    switch v {
        case "blob":
            result = BLOB_TREESPOSTREQUESTBODY_TREE_TYPE
        case "tree":
            result = TREE_TREESPOSTREQUESTBODY_TREE_TYPE
        case "commit":
            result = COMMIT_TREESPOSTREQUESTBODY_TREE_TYPE
        default:
            return 0, errors.New("Unknown TreesPostRequestBody_tree_type value: " + v)
    }
    return &result, nil
}
func SerializeTreesPostRequestBody_tree_type(values []TreesPostRequestBody_tree_type) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i TreesPostRequestBody_tree_type) isMultiValue() bool {
    return false
}
