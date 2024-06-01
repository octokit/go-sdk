package releases
import (
    "errors"
)
// Specifies whether this release should be set as the latest release for the repository. Drafts and prereleases cannot be set as latest. Defaults to `true` for newly published releases. `legacy` specifies that the latest release should be determined based on the release creation date and higher semantic version.
type ReleasesPostRequestBody_make_latest int

const (
    TRUE_RELEASESPOSTREQUESTBODY_MAKE_LATEST ReleasesPostRequestBody_make_latest = iota
    FALSE_RELEASESPOSTREQUESTBODY_MAKE_LATEST
    LEGACY_RELEASESPOSTREQUESTBODY_MAKE_LATEST
)

func (i ReleasesPostRequestBody_make_latest) String() string {
    return []string{"true", "false", "legacy"}[i]
}
func ParseReleasesPostRequestBody_make_latest(v string) (any, error) {
    result := TRUE_RELEASESPOSTREQUESTBODY_MAKE_LATEST
    switch v {
        case "true":
            result = TRUE_RELEASESPOSTREQUESTBODY_MAKE_LATEST
        case "false":
            result = FALSE_RELEASESPOSTREQUESTBODY_MAKE_LATEST
        case "legacy":
            result = LEGACY_RELEASESPOSTREQUESTBODY_MAKE_LATEST
        default:
            return 0, errors.New("Unknown ReleasesPostRequestBody_make_latest value: " + v)
    }
    return &result, nil
}
func SerializeReleasesPostRequestBody_make_latest(values []ReleasesPostRequestBody_make_latest) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i ReleasesPostRequestBody_make_latest) isMultiValue() bool {
    return false
}
