package models
// The frequency of the periodic analysis.
type CodeScanningDefaultSetup_schedule int

const (
    WEEKLY_CODESCANNINGDEFAULTSETUP_SCHEDULE CodeScanningDefaultSetup_schedule = iota
)

func (i CodeScanningDefaultSetup_schedule) String() string {
    return []string{"weekly"}[i]
}
func ParseCodeScanningDefaultSetup_schedule(v string) (any, error) {
    result := WEEKLY_CODESCANNINGDEFAULTSETUP_SCHEDULE
    switch v {
        case "weekly":
            result = WEEKLY_CODESCANNINGDEFAULTSETUP_SCHEDULE
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeCodeScanningDefaultSetup_schedule(values []CodeScanningDefaultSetup_schedule) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i CodeScanningDefaultSetup_schedule) isMultiValue() bool {
    return false
}
