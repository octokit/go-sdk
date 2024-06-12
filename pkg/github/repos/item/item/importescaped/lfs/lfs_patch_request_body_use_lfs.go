package lfs
import (
    "errors"
)
// Whether to store large files during the import. `opt_in` means large files will be stored using Git LFS. `opt_out` means large files will be removed during the import.
type LfsPatchRequestBody_use_lfs int

const (
    OPT_IN_LFSPATCHREQUESTBODY_USE_LFS LfsPatchRequestBody_use_lfs = iota
    OPT_OUT_LFSPATCHREQUESTBODY_USE_LFS
)

func (i LfsPatchRequestBody_use_lfs) String() string {
    return []string{"opt_in", "opt_out"}[i]
}
func ParseLfsPatchRequestBody_use_lfs(v string) (any, error) {
    result := OPT_IN_LFSPATCHREQUESTBODY_USE_LFS
    switch v {
        case "opt_in":
            result = OPT_IN_LFSPATCHREQUESTBODY_USE_LFS
        case "opt_out":
            result = OPT_OUT_LFSPATCHREQUESTBODY_USE_LFS
        default:
            return 0, errors.New("Unknown LfsPatchRequestBody_use_lfs value: " + v)
    }
    return &result, nil
}
func SerializeLfsPatchRequestBody_use_lfs(values []LfsPatchRequestBody_use_lfs) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i LfsPatchRequestBody_use_lfs) isMultiValue() bool {
    return false
}
