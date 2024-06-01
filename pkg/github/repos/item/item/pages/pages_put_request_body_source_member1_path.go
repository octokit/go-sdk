package pages
import (
    "errors"
)
// The repository directory that includes the source files for the Pages site. Allowed paths are `/` or `/docs`.
type PagesPutRequestBody_sourceMember1_path int

const (
    SLASH_PAGESPUTREQUESTBODY_SOURCEMEMBER1_PATH PagesPutRequestBody_sourceMember1_path = iota
    DOCS_PAGESPUTREQUESTBODY_SOURCEMEMBER1_PATH
)

func (i PagesPutRequestBody_sourceMember1_path) String() string {
    return []string{"/", "/docs"}[i]
}
func ParsePagesPutRequestBody_sourceMember1_path(v string) (any, error) {
    result := SLASH_PAGESPUTREQUESTBODY_SOURCEMEMBER1_PATH
    switch v {
        case "/":
            result = SLASH_PAGESPUTREQUESTBODY_SOURCEMEMBER1_PATH
        case "/docs":
            result = DOCS_PAGESPUTREQUESTBODY_SOURCEMEMBER1_PATH
        default:
            return 0, errors.New("Unknown PagesPutRequestBody_sourceMember1_path value: " + v)
    }
    return &result, nil
}
func SerializePagesPutRequestBody_sourceMember1_path(values []PagesPutRequestBody_sourceMember1_path) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i PagesPutRequestBody_sourceMember1_path) isMultiValue() bool {
    return false
}
