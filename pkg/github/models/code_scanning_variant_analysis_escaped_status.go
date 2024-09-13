package models
type CodeScanningVariantAnalysis_status int

const (
    IN_PROGRESS_CODESCANNINGVARIANTANALYSIS_STATUS CodeScanningVariantAnalysis_status = iota
    SUCCEEDED_CODESCANNINGVARIANTANALYSIS_STATUS
    FAILED_CODESCANNINGVARIANTANALYSIS_STATUS
    CANCELLED_CODESCANNINGVARIANTANALYSIS_STATUS
)

func (i CodeScanningVariantAnalysis_status) String() string {
    return []string{"in_progress", "succeeded", "failed", "cancelled"}[i]
}
func ParseCodeScanningVariantAnalysis_status(v string) (any, error) {
    result := IN_PROGRESS_CODESCANNINGVARIANTANALYSIS_STATUS
    switch v {
        case "in_progress":
            result = IN_PROGRESS_CODESCANNINGVARIANTANALYSIS_STATUS
        case "succeeded":
            result = SUCCEEDED_CODESCANNINGVARIANTANALYSIS_STATUS
        case "failed":
            result = FAILED_CODESCANNINGVARIANTANALYSIS_STATUS
        case "cancelled":
            result = CANCELLED_CODESCANNINGVARIANTANALYSIS_STATUS
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeCodeScanningVariantAnalysis_status(values []CodeScanningVariantAnalysis_status) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i CodeScanningVariantAnalysis_status) isMultiValue() bool {
    return false
}
