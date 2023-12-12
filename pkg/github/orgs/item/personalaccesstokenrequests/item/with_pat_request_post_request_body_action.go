package item
import (
    "errors"
)
// Action to apply to the request.
type WithPat_request_PostRequestBody_action int

const (
    APPROVE_WITHPAT_REQUEST_POSTREQUESTBODY_ACTION WithPat_request_PostRequestBody_action = iota
    DENY_WITHPAT_REQUEST_POSTREQUESTBODY_ACTION
)

func (i WithPat_request_PostRequestBody_action) String() string {
    return []string{"approve", "deny"}[i]
}
func ParseWithPat_request_PostRequestBody_action(v string) (any, error) {
    result := APPROVE_WITHPAT_REQUEST_POSTREQUESTBODY_ACTION
    switch v {
        case "approve":
            result = APPROVE_WITHPAT_REQUEST_POSTREQUESTBODY_ACTION
        case "deny":
            result = DENY_WITHPAT_REQUEST_POSTREQUESTBODY_ACTION
        default:
            return 0, errors.New("Unknown WithPat_request_PostRequestBody_action value: " + v)
    }
    return &result, nil
}
func SerializeWithPat_request_PostRequestBody_action(values []WithPat_request_PostRequestBody_action) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i WithPat_request_PostRequestBody_action) isMultiValue() bool {
    return false
}
