package item
import (
    "errors"
)
// The default value for a merge commit message.- `PR_TITLE` - default to the pull request's title.- `PR_BODY` - default to the pull request's body.- `BLANK` - default to a blank commit message.
type OwnerPatchRequestBody_merge_commit_message int

const (
    PR_BODY_OWNERPATCHREQUESTBODY_MERGE_COMMIT_MESSAGE OwnerPatchRequestBody_merge_commit_message = iota
    PR_TITLE_OWNERPATCHREQUESTBODY_MERGE_COMMIT_MESSAGE
    BLANK_OWNERPATCHREQUESTBODY_MERGE_COMMIT_MESSAGE
)

func (i OwnerPatchRequestBody_merge_commit_message) String() string {
    return []string{"PR_BODY", "PR_TITLE", "BLANK"}[i]
}
func ParseOwnerPatchRequestBody_merge_commit_message(v string) (any, error) {
    result := PR_BODY_OWNERPATCHREQUESTBODY_MERGE_COMMIT_MESSAGE
    switch v {
        case "PR_BODY":
            result = PR_BODY_OWNERPATCHREQUESTBODY_MERGE_COMMIT_MESSAGE
        case "PR_TITLE":
            result = PR_TITLE_OWNERPATCHREQUESTBODY_MERGE_COMMIT_MESSAGE
        case "BLANK":
            result = BLANK_OWNERPATCHREQUESTBODY_MERGE_COMMIT_MESSAGE
        default:
            return 0, errors.New("Unknown OwnerPatchRequestBody_merge_commit_message value: " + v)
    }
    return &result, nil
}
func SerializeOwnerPatchRequestBody_merge_commit_message(values []OwnerPatchRequestBody_merge_commit_message) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i OwnerPatchRequestBody_merge_commit_message) isMultiValue() bool {
    return false
}
