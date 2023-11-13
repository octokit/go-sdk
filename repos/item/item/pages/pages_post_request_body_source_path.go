package pages
import (
    "errors"
)
// The repository directory that includes the source files for the Pages site. Allowed paths are `/` or `/docs`. Default: `/`
type PagesPostRequestBody_source_path int

const (
    SLASH_PAGESPOSTREQUESTBODY_SOURCE_PATH PagesPostRequestBody_source_path = iota
    DOCS_PAGESPOSTREQUESTBODY_SOURCE_PATH
)

func (i PagesPostRequestBody_source_path) String() string {
    return []string{"/", "/docs"}[i]
}
func ParsePagesPostRequestBody_source_path(v string) (any, error) {
    result := SLASH_PAGESPOSTREQUESTBODY_SOURCE_PATH
    switch v {
        case "/":
            result = SLASH_PAGESPOSTREQUESTBODY_SOURCE_PATH
        case "/docs":
            result = DOCS_PAGESPOSTREQUESTBODY_SOURCE_PATH
        default:
            return 0, errors.New("Unknown PagesPostRequestBody_source_path value: " + v)
    }
    return &result, nil
}
func SerializePagesPostRequestBody_source_path(values []PagesPostRequestBody_source_path) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i PagesPostRequestBody_source_path) isMultiValue() bool {
    return false
}
