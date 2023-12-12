package item
import (
    "errors"
)
// Action to apply to the fine-grained personal access token.
type WithPat_PostRequestBody_action int

const (
    REVOKE_WITHPAT_POSTREQUESTBODY_ACTION WithPat_PostRequestBody_action = iota
)

func (i WithPat_PostRequestBody_action) String() string {
    return []string{"revoke"}[i]
}
func ParseWithPat_PostRequestBody_action(v string) (any, error) {
    result := REVOKE_WITHPAT_POSTREQUESTBODY_ACTION
    switch v {
        case "revoke":
            result = REVOKE_WITHPAT_POSTREQUESTBODY_ACTION
        default:
            return 0, errors.New("Unknown WithPat_PostRequestBody_action value: " + v)
    }
    return &result, nil
}
func SerializeWithPat_PostRequestBody_action(values []WithPat_PostRequestBody_action) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i WithPat_PostRequestBody_action) isMultiValue() bool {
    return false
}
