package item
import (
    "errors"
)
// CodeQL query suite to be used. If you specify the `query_suite` parameter, the default setup will be configured with this query suite only on all repositories that didn't have default setup already configured. It will not change the query suite on repositories that already have default setup configured.If you don't specify any `query_suite` in your request, the preferred query suite of the organization will be applied.
type WithEnablementPostRequestBody_query_suite int

const (
    DEFAULT_WITHENABLEMENTPOSTREQUESTBODY_QUERY_SUITE WithEnablementPostRequestBody_query_suite = iota
    EXTENDED_WITHENABLEMENTPOSTREQUESTBODY_QUERY_SUITE
)

func (i WithEnablementPostRequestBody_query_suite) String() string {
    return []string{"default", "extended"}[i]
}
func ParseWithEnablementPostRequestBody_query_suite(v string) (any, error) {
    result := DEFAULT_WITHENABLEMENTPOSTREQUESTBODY_QUERY_SUITE
    switch v {
        case "default":
            result = DEFAULT_WITHENABLEMENTPOSTREQUESTBODY_QUERY_SUITE
        case "extended":
            result = EXTENDED_WITHENABLEMENTPOSTREQUESTBODY_QUERY_SUITE
        default:
            return 0, errors.New("Unknown WithEnablementPostRequestBody_query_suite value: " + v)
    }
    return &result, nil
}
func SerializeWithEnablementPostRequestBody_query_suite(values []WithEnablementPostRequestBody_query_suite) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i WithEnablementPostRequestBody_query_suite) isMultiValue() bool {
    return false
}
