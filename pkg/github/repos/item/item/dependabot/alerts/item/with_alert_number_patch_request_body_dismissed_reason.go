package item
import (
    "errors"
)
// **Required when `state` is `dismissed`.** A reason for dismissing the alert.
type WithAlert_numberPatchRequestBody_dismissed_reason int

const (
    FIX_STARTED_WITHALERT_NUMBERPATCHREQUESTBODY_DISMISSED_REASON WithAlert_numberPatchRequestBody_dismissed_reason = iota
    INACCURATE_WITHALERT_NUMBERPATCHREQUESTBODY_DISMISSED_REASON
    NO_BANDWIDTH_WITHALERT_NUMBERPATCHREQUESTBODY_DISMISSED_REASON
    NOT_USED_WITHALERT_NUMBERPATCHREQUESTBODY_DISMISSED_REASON
    TOLERABLE_RISK_WITHALERT_NUMBERPATCHREQUESTBODY_DISMISSED_REASON
)

func (i WithAlert_numberPatchRequestBody_dismissed_reason) String() string {
    return []string{"fix_started", "inaccurate", "no_bandwidth", "not_used", "tolerable_risk"}[i]
}
func ParseWithAlert_numberPatchRequestBody_dismissed_reason(v string) (any, error) {
    result := FIX_STARTED_WITHALERT_NUMBERPATCHREQUESTBODY_DISMISSED_REASON
    switch v {
        case "fix_started":
            result = FIX_STARTED_WITHALERT_NUMBERPATCHREQUESTBODY_DISMISSED_REASON
        case "inaccurate":
            result = INACCURATE_WITHALERT_NUMBERPATCHREQUESTBODY_DISMISSED_REASON
        case "no_bandwidth":
            result = NO_BANDWIDTH_WITHALERT_NUMBERPATCHREQUESTBODY_DISMISSED_REASON
        case "not_used":
            result = NOT_USED_WITHALERT_NUMBERPATCHREQUESTBODY_DISMISSED_REASON
        case "tolerable_risk":
            result = TOLERABLE_RISK_WITHALERT_NUMBERPATCHREQUESTBODY_DISMISSED_REASON
        default:
            return 0, errors.New("Unknown WithAlert_numberPatchRequestBody_dismissed_reason value: " + v)
    }
    return &result, nil
}
func SerializeWithAlert_numberPatchRequestBody_dismissed_reason(values []WithAlert_numberPatchRequestBody_dismissed_reason) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i WithAlert_numberPatchRequestBody_dismissed_reason) isMultiValue() bool {
    return false
}
