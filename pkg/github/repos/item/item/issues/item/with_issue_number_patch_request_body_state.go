package item
// The open or closed state of the issue.
type WithIssue_numberPatchRequestBody_state int

const (
    OPEN_WITHISSUE_NUMBERPATCHREQUESTBODY_STATE WithIssue_numberPatchRequestBody_state = iota
    CLOSED_WITHISSUE_NUMBERPATCHREQUESTBODY_STATE
)

func (i WithIssue_numberPatchRequestBody_state) String() string {
    return []string{"open", "closed"}[i]
}
func ParseWithIssue_numberPatchRequestBody_state(v string) (any, error) {
    result := OPEN_WITHISSUE_NUMBERPATCHREQUESTBODY_STATE
    switch v {
        case "open":
            result = OPEN_WITHISSUE_NUMBERPATCHREQUESTBODY_STATE
        case "closed":
            result = CLOSED_WITHISSUE_NUMBERPATCHREQUESTBODY_STATE
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeWithIssue_numberPatchRequestBody_state(values []WithIssue_numberPatchRequestBody_state) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i WithIssue_numberPatchRequestBody_state) isMultiValue() bool {
    return false
}
