package item
import (
    "errors"
)
// Specifies whether this release should be set as the latest release for the repository. Drafts and prereleases cannot be set as latest. Defaults to `true` for newly published releases. `legacy` specifies that the latest release should be determined based on the release creation date and higher semantic version.
type WithRelease_PatchRequestBody_make_latest int

const (
    TRUE_WITHRELEASE_PATCHREQUESTBODY_MAKE_LATEST WithRelease_PatchRequestBody_make_latest = iota
    FALSE_WITHRELEASE_PATCHREQUESTBODY_MAKE_LATEST
    LEGACY_WITHRELEASE_PATCHREQUESTBODY_MAKE_LATEST
)

func (i WithRelease_PatchRequestBody_make_latest) String() string {
    return []string{"true", "false", "legacy"}[i]
}
func ParseWithRelease_PatchRequestBody_make_latest(v string) (any, error) {
    result := TRUE_WITHRELEASE_PATCHREQUESTBODY_MAKE_LATEST
    switch v {
        case "true":
            result = TRUE_WITHRELEASE_PATCHREQUESTBODY_MAKE_LATEST
        case "false":
            result = FALSE_WITHRELEASE_PATCHREQUESTBODY_MAKE_LATEST
        case "legacy":
            result = LEGACY_WITHRELEASE_PATCHREQUESTBODY_MAKE_LATEST
        default:
            return 0, errors.New("Unknown WithRelease_PatchRequestBody_make_latest value: " + v)
    }
    return &result, nil
}
func SerializeWithRelease_PatchRequestBody_make_latest(values []WithRelease_PatchRequestBody_make_latest) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i WithRelease_PatchRequestBody_make_latest) isMultiValue() bool {
    return false
}
