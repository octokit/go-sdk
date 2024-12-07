package models
// Whether to use labeled runners or standard GitHub runners.
type CodeScanningDefaultSetupOptions_runner_type int

const (
    STANDARD_CODESCANNINGDEFAULTSETUPOPTIONS_RUNNER_TYPE CodeScanningDefaultSetupOptions_runner_type = iota
    LABELED_CODESCANNINGDEFAULTSETUPOPTIONS_RUNNER_TYPE
    NOT_SET_CODESCANNINGDEFAULTSETUPOPTIONS_RUNNER_TYPE
)

func (i CodeScanningDefaultSetupOptions_runner_type) String() string {
    return []string{"standard", "labeled", "not_set"}[i]
}
func ParseCodeScanningDefaultSetupOptions_runner_type(v string) (any, error) {
    result := STANDARD_CODESCANNINGDEFAULTSETUPOPTIONS_RUNNER_TYPE
    switch v {
        case "standard":
            result = STANDARD_CODESCANNINGDEFAULTSETUPOPTIONS_RUNNER_TYPE
        case "labeled":
            result = LABELED_CODESCANNINGDEFAULTSETUPOPTIONS_RUNNER_TYPE
        case "not_set":
            result = NOT_SET_CODESCANNINGDEFAULTSETUPOPTIONS_RUNNER_TYPE
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeCodeScanningDefaultSetupOptions_runner_type(values []CodeScanningDefaultSetupOptions_runner_type) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i CodeScanningDefaultSetupOptions_runner_type) isMultiValue() bool {
    return false
}
