package importescaped
import (
    "errors"
)
// The originating VCS type. Without this parameter, the import job will take additional time to detect the VCS type before beginning the import. This detection step will be reflected in the response.
type ImportPutRequestBody_vcs int

const (
    SUBVERSION_IMPORTPUTREQUESTBODY_VCS ImportPutRequestBody_vcs = iota
    GIT_IMPORTPUTREQUESTBODY_VCS
    MERCURIAL_IMPORTPUTREQUESTBODY_VCS
    TFVC_IMPORTPUTREQUESTBODY_VCS
)

func (i ImportPutRequestBody_vcs) String() string {
    return []string{"subversion", "git", "mercurial", "tfvc"}[i]
}
func ParseImportPutRequestBody_vcs(v string) (any, error) {
    result := SUBVERSION_IMPORTPUTREQUESTBODY_VCS
    switch v {
        case "subversion":
            result = SUBVERSION_IMPORTPUTREQUESTBODY_VCS
        case "git":
            result = GIT_IMPORTPUTREQUESTBODY_VCS
        case "mercurial":
            result = MERCURIAL_IMPORTPUTREQUESTBODY_VCS
        case "tfvc":
            result = TFVC_IMPORTPUTREQUESTBODY_VCS
        default:
            return 0, errors.New("Unknown ImportPutRequestBody_vcs value: " + v)
    }
    return &result, nil
}
func SerializeImportPutRequestBody_vcs(values []ImportPutRequestBody_vcs) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i ImportPutRequestBody_vcs) isMultiValue() bool {
    return false
}
