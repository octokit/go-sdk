package advisories
type GetSortQueryParameterType int

const (
    UPDATED_GETSORTQUERYPARAMETERTYPE GetSortQueryParameterType = iota
    PUBLISHED_GETSORTQUERYPARAMETERTYPE
    EPSS_PERCENTAGE_GETSORTQUERYPARAMETERTYPE
    EPSS_PERCENTILE_GETSORTQUERYPARAMETERTYPE
)

func (i GetSortQueryParameterType) String() string {
    return []string{"updated", "published", "epss_percentage", "epss_percentile"}[i]
}
func ParseGetSortQueryParameterType(v string) (any, error) {
    result := UPDATED_GETSORTQUERYPARAMETERTYPE
    switch v {
        case "updated":
            result = UPDATED_GETSORTQUERYPARAMETERTYPE
        case "published":
            result = PUBLISHED_GETSORTQUERYPARAMETERTYPE
        case "epss_percentage":
            result = EPSS_PERCENTAGE_GETSORTQUERYPARAMETERTYPE
        case "epss_percentile":
            result = EPSS_PERCENTILE_GETSORTQUERYPARAMETERTYPE
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeGetSortQueryParameterType(values []GetSortQueryParameterType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i GetSortQueryParameterType) isMultiValue() bool {
    return false
}
