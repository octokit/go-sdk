package models
// The status of an autofix.
type CodeScanningAutofixStatus int

const (
    PENDING_CODESCANNINGAUTOFIXSTATUS CodeScanningAutofixStatus = iota
    ERROR_CODESCANNINGAUTOFIXSTATUS
    SUCCESS_CODESCANNINGAUTOFIXSTATUS
    OUTDATED_CODESCANNINGAUTOFIXSTATUS
)

func (i CodeScanningAutofixStatus) String() string {
    return []string{"pending", "error", "success", "outdated"}[i]
}
func ParseCodeScanningAutofixStatus(v string) (any, error) {
    result := PENDING_CODESCANNINGAUTOFIXSTATUS
    switch v {
        case "pending":
            result = PENDING_CODESCANNINGAUTOFIXSTATUS
        case "error":
            result = ERROR_CODESCANNINGAUTOFIXSTATUS
        case "success":
            result = SUCCESS_CODESCANNINGAUTOFIXSTATUS
        case "outdated":
            result = OUTDATED_CODESCANNINGAUTOFIXSTATUS
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeCodeScanningAutofixStatus(values []CodeScanningAutofixStatus) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i CodeScanningAutofixStatus) isMultiValue() bool {
    return false
}
