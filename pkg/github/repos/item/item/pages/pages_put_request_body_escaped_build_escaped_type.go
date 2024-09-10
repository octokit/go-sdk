package pages
// The process by which the GitHub Pages site will be built. `workflow` means that the site is built by a custom GitHub Actions workflow. `legacy` means that the site is built by GitHub when changes are pushed to a specific branch.
type PagesPutRequestBody_build_type int

const (
    LEGACY_PAGESPUTREQUESTBODY_BUILD_TYPE PagesPutRequestBody_build_type = iota
    WORKFLOW_PAGESPUTREQUESTBODY_BUILD_TYPE
)

func (i PagesPutRequestBody_build_type) String() string {
    return []string{"legacy", "workflow"}[i]
}
func ParsePagesPutRequestBody_build_type(v string) (any, error) {
    result := LEGACY_PAGESPUTREQUESTBODY_BUILD_TYPE
    switch v {
        case "legacy":
            result = LEGACY_PAGESPUTREQUESTBODY_BUILD_TYPE
        case "workflow":
            result = WORKFLOW_PAGESPUTREQUESTBODY_BUILD_TYPE
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializePagesPutRequestBody_build_type(values []PagesPutRequestBody_build_type) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i PagesPutRequestBody_build_type) isMultiValue() bool {
    return false
}
