package checkruns
import (
    "errors"
)
// 
type CheckRunsPostRequestBodyMember2_status int

const (
    QUEUED_CHECKRUNSPOSTREQUESTBODYMEMBER2_STATUS CheckRunsPostRequestBodyMember2_status = iota
    IN_PROGRESS_CHECKRUNSPOSTREQUESTBODYMEMBER2_STATUS
)

func (i CheckRunsPostRequestBodyMember2_status) String() string {
    return []string{"queued", "in_progress"}[i]
}
func ParseCheckRunsPostRequestBodyMember2_status(v string) (any, error) {
    result := QUEUED_CHECKRUNSPOSTREQUESTBODYMEMBER2_STATUS
    switch v {
        case "queued":
            result = QUEUED_CHECKRUNSPOSTREQUESTBODYMEMBER2_STATUS
        case "in_progress":
            result = IN_PROGRESS_CHECKRUNSPOSTREQUESTBODYMEMBER2_STATUS
        default:
            return 0, errors.New("Unknown CheckRunsPostRequestBodyMember2_status value: " + v)
    }
    return &result, nil
}
func SerializeCheckRunsPostRequestBodyMember2_status(values []CheckRunsPostRequestBodyMember2_status) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i CheckRunsPostRequestBodyMember2_status) isMultiValue() bool {
    return false
}
