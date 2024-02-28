package checkruns
import (
    "errors"
)
type CheckRunsPostRequestBodyMember1_status int

const (
    COMPLETED_CHECKRUNSPOSTREQUESTBODYMEMBER1_STATUS CheckRunsPostRequestBodyMember1_status = iota
)

func (i CheckRunsPostRequestBodyMember1_status) String() string {
    return []string{"completed"}[i]
}
func ParseCheckRunsPostRequestBodyMember1_status(v string) (any, error) {
    result := COMPLETED_CHECKRUNSPOSTREQUESTBODYMEMBER1_STATUS
    switch v {
        case "completed":
            result = COMPLETED_CHECKRUNSPOSTREQUESTBODYMEMBER1_STATUS
        default:
            return 0, errors.New("Unknown CheckRunsPostRequestBodyMember1_status value: " + v)
    }
    return &result, nil
}
func SerializeCheckRunsPostRequestBodyMember1_status(values []CheckRunsPostRequestBodyMember1_status) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i CheckRunsPostRequestBodyMember1_status) isMultiValue() bool {
    return false
}
