package item
type GetSortQueryParameterType int

const (
    LAST_RATE_LIMITED_TIMESTAMP_GETSORTQUERYPARAMETERTYPE GetSortQueryParameterType = iota
    LAST_REQUEST_TIMESTAMP_GETSORTQUERYPARAMETERTYPE
    RATE_LIMITED_REQUEST_COUNT_GETSORTQUERYPARAMETERTYPE
    HTTP_METHOD_GETSORTQUERYPARAMETERTYPE
    API_ROUTE_GETSORTQUERYPARAMETERTYPE
    TOTAL_REQUEST_COUNT_GETSORTQUERYPARAMETERTYPE
)

func (i GetSortQueryParameterType) String() string {
    return []string{"last_rate_limited_timestamp", "last_request_timestamp", "rate_limited_request_count", "http_method", "api_route", "total_request_count"}[i]
}
func ParseGetSortQueryParameterType(v string) (any, error) {
    result := LAST_RATE_LIMITED_TIMESTAMP_GETSORTQUERYPARAMETERTYPE
    switch v {
        case "last_rate_limited_timestamp":
            result = LAST_RATE_LIMITED_TIMESTAMP_GETSORTQUERYPARAMETERTYPE
        case "last_request_timestamp":
            result = LAST_REQUEST_TIMESTAMP_GETSORTQUERYPARAMETERTYPE
        case "rate_limited_request_count":
            result = RATE_LIMITED_REQUEST_COUNT_GETSORTQUERYPARAMETERTYPE
        case "http_method":
            result = HTTP_METHOD_GETSORTQUERYPARAMETERTYPE
        case "api_route":
            result = API_ROUTE_GETSORTQUERYPARAMETERTYPE
        case "total_request_count":
            result = TOTAL_REQUEST_COUNT_GETSORTQUERYPARAMETERTYPE
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
