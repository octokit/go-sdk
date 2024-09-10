package importEscaped
// The type of version control system you are migrating from.
type ImportPatchRequestBody_vcs int

const (
    SUBVERSION_IMPORTPATCHREQUESTBODY_VCS ImportPatchRequestBody_vcs = iota
    TFVC_IMPORTPATCHREQUESTBODY_VCS
    GIT_IMPORTPATCHREQUESTBODY_VCS
    MERCURIAL_IMPORTPATCHREQUESTBODY_VCS
)

func (i ImportPatchRequestBody_vcs) String() string {
    return []string{"subversion", "tfvc", "git", "mercurial"}[i]
}
func ParseImportPatchRequestBody_vcs(v string) (any, error) {
    result := SUBVERSION_IMPORTPATCHREQUESTBODY_VCS
    switch v {
        case "subversion":
            result = SUBVERSION_IMPORTPATCHREQUESTBODY_VCS
        case "tfvc":
            result = TFVC_IMPORTPATCHREQUESTBODY_VCS
        case "git":
            result = GIT_IMPORTPATCHREQUESTBODY_VCS
        case "mercurial":
            result = MERCURIAL_IMPORTPATCHREQUESTBODY_VCS
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeImportPatchRequestBody_vcs(values []ImportPatchRequestBody_vcs) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i ImportPatchRequestBody_vcs) isMultiValue() bool {
    return false
}
