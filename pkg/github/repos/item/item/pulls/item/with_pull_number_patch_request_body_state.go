package item
// State of this Pull Request. Either `open` or `closed`.
type WithPull_numberPatchRequestBody_state int

const (
    OPEN_WITHPULL_NUMBERPATCHREQUESTBODY_STATE WithPull_numberPatchRequestBody_state = iota
    CLOSED_WITHPULL_NUMBERPATCHREQUESTBODY_STATE
)

func (i WithPull_numberPatchRequestBody_state) String() string {
    return []string{"open", "closed"}[i]
}
func ParseWithPull_numberPatchRequestBody_state(v string) (any, error) {
    result := OPEN_WITHPULL_NUMBERPATCHREQUESTBODY_STATE
    switch v {
        case "open":
            result = OPEN_WITHPULL_NUMBERPATCHREQUESTBODY_STATE
        case "closed":
            result = CLOSED_WITHPULL_NUMBERPATCHREQUESTBODY_STATE
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeWithPull_numberPatchRequestBody_state(values []WithPull_numberPatchRequestBody_state) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i WithPull_numberPatchRequestBody_state) isMultiValue() bool {
    return false
}
