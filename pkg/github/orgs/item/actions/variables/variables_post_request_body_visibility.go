package variables
import (
    "errors"
)
// The type of repositories in the organization that can access the variable. `selected` means only the repositories specified by `selected_repository_ids` can access the variable.
type VariablesPostRequestBody_visibility int

const (
    ALL_VARIABLESPOSTREQUESTBODY_VISIBILITY VariablesPostRequestBody_visibility = iota
    PRIVATE_VARIABLESPOSTREQUESTBODY_VISIBILITY
    SELECTED_VARIABLESPOSTREQUESTBODY_VISIBILITY
)

func (i VariablesPostRequestBody_visibility) String() string {
    return []string{"all", "private", "selected"}[i]
}
func ParseVariablesPostRequestBody_visibility(v string) (any, error) {
    result := ALL_VARIABLESPOSTREQUESTBODY_VISIBILITY
    switch v {
        case "all":
            result = ALL_VARIABLESPOSTREQUESTBODY_VISIBILITY
        case "private":
            result = PRIVATE_VARIABLESPOSTREQUESTBODY_VISIBILITY
        case "selected":
            result = SELECTED_VARIABLESPOSTREQUESTBODY_VISIBILITY
        default:
            return 0, errors.New("Unknown VariablesPostRequestBody_visibility value: " + v)
    }
    return &result, nil
}
func SerializeVariablesPostRequestBody_visibility(values []VariablesPostRequestBody_visibility) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i VariablesPostRequestBody_visibility) isMultiValue() bool {
    return false
}
