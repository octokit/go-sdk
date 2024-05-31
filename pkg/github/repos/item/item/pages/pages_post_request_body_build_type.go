package pages
// The process in which the Page will be built. Possible values are `"legacy"` and `"workflow"`.
type PagesPostRequestBody_build_type int

const (
    LEGACY_PAGESPOSTREQUESTBODY_BUILD_TYPE PagesPostRequestBody_build_type = iota
    WORKFLOW_PAGESPOSTREQUESTBODY_BUILD_TYPE
)

func (i PagesPostRequestBody_build_type) String() string {
    return []string{"legacy", "workflow"}[i]
}
func ParsePagesPostRequestBody_build_type(v string) (any, error) {
    result := LEGACY_PAGESPOSTREQUESTBODY_BUILD_TYPE
    switch v {
        case "legacy":
            result = LEGACY_PAGESPOSTREQUESTBODY_BUILD_TYPE
        case "workflow":
            result = WORKFLOW_PAGESPOSTREQUESTBODY_BUILD_TYPE
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializePagesPostRequestBody_build_type(values []PagesPostRequestBody_build_type) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i PagesPostRequestBody_build_type) isMultiValue() bool {
    return false
}
