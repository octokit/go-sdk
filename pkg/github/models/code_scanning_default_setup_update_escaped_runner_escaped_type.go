package models
// Runner type to be used.
type CodeScanningDefaultSetupUpdate_runner_type int

const (
    STANDARD_CODESCANNINGDEFAULTSETUPUPDATE_RUNNER_TYPE CodeScanningDefaultSetupUpdate_runner_type = iota
    LABELED_CODESCANNINGDEFAULTSETUPUPDATE_RUNNER_TYPE
)

func (i CodeScanningDefaultSetupUpdate_runner_type) String() string {
    return []string{"standard", "labeled"}[i]
}
func ParseCodeScanningDefaultSetupUpdate_runner_type(v string) (any, error) {
    result := STANDARD_CODESCANNINGDEFAULTSETUPUPDATE_RUNNER_TYPE
    switch v {
        case "standard":
            result = STANDARD_CODESCANNINGDEFAULTSETUPUPDATE_RUNNER_TYPE
        case "labeled":
            result = LABELED_CODESCANNINGDEFAULTSETUPUPDATE_RUNNER_TYPE
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeCodeScanningDefaultSetupUpdate_runner_type(values []CodeScanningDefaultSetupUpdate_runner_type) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i CodeScanningDefaultSetupUpdate_runner_type) isMultiValue() bool {
    return false
}
