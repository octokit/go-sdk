package models
import (
    "errors"
)
// The location type. Because secrets may be found in different types of resources (ie. code, comments, issues), this field identifies the type of resource where the secret was found.
type SecretScanningLocation_type int

const (
    COMMIT_SECRETSCANNINGLOCATION_TYPE SecretScanningLocation_type = iota
    ISSUE_TITLE_SECRETSCANNINGLOCATION_TYPE
    ISSUE_BODY_SECRETSCANNINGLOCATION_TYPE
    ISSUE_COMMENT_SECRETSCANNINGLOCATION_TYPE
)

func (i SecretScanningLocation_type) String() string {
    return []string{"commit", "issue_title", "issue_body", "issue_comment"}[i]
}
func ParseSecretScanningLocation_type(v string) (any, error) {
    result := COMMIT_SECRETSCANNINGLOCATION_TYPE
    switch v {
        case "commit":
            result = COMMIT_SECRETSCANNINGLOCATION_TYPE
        case "issue_title":
            result = ISSUE_TITLE_SECRETSCANNINGLOCATION_TYPE
        case "issue_body":
            result = ISSUE_BODY_SECRETSCANNINGLOCATION_TYPE
        case "issue_comment":
            result = ISSUE_COMMENT_SECRETSCANNINGLOCATION_TYPE
        default:
            return 0, errors.New("Unknown SecretScanningLocation_type value: " + v)
    }
    return &result, nil
}
func SerializeSecretScanningLocation_type(values []SecretScanningLocation_type) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i SecretScanningLocation_type) isMultiValue() bool {
    return false
}
