package trees
// The file mode; one of `100644` for file (blob), `100755` for executable (blob), `040000` for subdirectory (tree), `160000` for submodule (commit), or `120000` for a blob that specifies the path of a symlink.
type TreesPostRequestBody_tree_mode int

const (
    ONEZEROZEROSIXFOURFOUR_TREESPOSTREQUESTBODY_TREE_MODE TreesPostRequestBody_tree_mode = iota
    ONEZEROZEROSEVENFIVEFIVE_TREESPOSTREQUESTBODY_TREE_MODE
    ZEROFOURZEROZEROZEROZERO_TREESPOSTREQUESTBODY_TREE_MODE
    ONESIXZEROZEROZEROZERO_TREESPOSTREQUESTBODY_TREE_MODE
    ONETWOZEROZEROZEROZERO_TREESPOSTREQUESTBODY_TREE_MODE
)

func (i TreesPostRequestBody_tree_mode) String() string {
    return []string{"100644", "100755", "040000", "160000", "120000"}[i]
}
func ParseTreesPostRequestBody_tree_mode(v string) (any, error) {
    result := ONEZEROZEROSIXFOURFOUR_TREESPOSTREQUESTBODY_TREE_MODE
    switch v {
        case "100644":
            result = ONEZEROZEROSIXFOURFOUR_TREESPOSTREQUESTBODY_TREE_MODE
        case "100755":
            result = ONEZEROZEROSEVENFIVEFIVE_TREESPOSTREQUESTBODY_TREE_MODE
        case "040000":
            result = ZEROFOURZEROZEROZEROZERO_TREESPOSTREQUESTBODY_TREE_MODE
        case "160000":
            result = ONESIXZEROZEROZEROZERO_TREESPOSTREQUESTBODY_TREE_MODE
        case "120000":
            result = ONETWOZEROZEROZEROZERO_TREESPOSTREQUESTBODY_TREE_MODE
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeTreesPostRequestBody_tree_mode(values []TreesPostRequestBody_tree_mode) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i TreesPostRequestBody_tree_mode) isMultiValue() bool {
    return false
}
