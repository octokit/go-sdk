package models
// Runner type to be used.
type CodeScanningDefaultSetup_runner_type int

const (
    STANDARD_CODESCANNINGDEFAULTSETUP_RUNNER_TYPE CodeScanningDefaultSetup_runner_type = iota
    LABELED_CODESCANNINGDEFAULTSETUP_RUNNER_TYPE
)

func (i CodeScanningDefaultSetup_runner_type) String() string {
    return []string{"standard", "labeled"}[i]
}
func ParseCodeScanningDefaultSetup_runner_type(v string) (any, error) {
    result := STANDARD_CODESCANNINGDEFAULTSETUP_RUNNER_TYPE
    switch v {
        case "standard":
            result = STANDARD_CODESCANNINGDEFAULTSETUP_RUNNER_TYPE
        case "labeled":
            result = LABELED_CODESCANNINGDEFAULTSETUP_RUNNER_TYPE
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeCodeScanningDefaultSetup_runner_type(values []CodeScanningDefaultSetup_runner_type) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i CodeScanningDefaultSetup_runner_type) isMultiValue() bool {
    return false
}
