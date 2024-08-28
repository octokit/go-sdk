package tags
// The type of the object we're tagging. Normally this is a `commit` but it can also be a `tree` or a `blob`.
type TagsPostRequestBody_type int

const (
    COMMIT_TAGSPOSTREQUESTBODY_TYPE TagsPostRequestBody_type = iota
    TREE_TAGSPOSTREQUESTBODY_TYPE
    BLOB_TAGSPOSTREQUESTBODY_TYPE
)

func (i TagsPostRequestBody_type) String() string {
    return []string{"commit", "tree", "blob"}[i]
}
func ParseTagsPostRequestBody_type(v string) (any, error) {
    result := COMMIT_TAGSPOSTREQUESTBODY_TYPE
    switch v {
        case "commit":
            result = COMMIT_TAGSPOSTREQUESTBODY_TYPE
        case "tree":
            result = TREE_TAGSPOSTREQUESTBODY_TYPE
        case "blob":
            result = BLOB_TAGSPOSTREQUESTBODY_TYPE
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeTagsPostRequestBody_type(values []TagsPostRequestBody_type) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i TagsPostRequestBody_type) isMultiValue() bool {
    return false
}
