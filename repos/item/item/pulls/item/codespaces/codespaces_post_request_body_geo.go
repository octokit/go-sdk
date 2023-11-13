package codespaces
import (
    "errors"
)
// The geographic area for this codespace. If not specified, the value is assigned by IP. This property replaces `location`, which is being deprecated.
type CodespacesPostRequestBody_geo int

const (
    EUROPEWEST_CODESPACESPOSTREQUESTBODY_GEO CodespacesPostRequestBody_geo = iota
    SOUTHEASTASIA_CODESPACESPOSTREQUESTBODY_GEO
    USEAST_CODESPACESPOSTREQUESTBODY_GEO
    USWEST_CODESPACESPOSTREQUESTBODY_GEO
)

func (i CodespacesPostRequestBody_geo) String() string {
    return []string{"EuropeWest", "SoutheastAsia", "UsEast", "UsWest"}[i]
}
func ParseCodespacesPostRequestBody_geo(v string) (any, error) {
    result := EUROPEWEST_CODESPACESPOSTREQUESTBODY_GEO
    switch v {
        case "EuropeWest":
            result = EUROPEWEST_CODESPACESPOSTREQUESTBODY_GEO
        case "SoutheastAsia":
            result = SOUTHEASTASIA_CODESPACESPOSTREQUESTBODY_GEO
        case "UsEast":
            result = USEAST_CODESPACESPOSTREQUESTBODY_GEO
        case "UsWest":
            result = USWEST_CODESPACESPOSTREQUESTBODY_GEO
        default:
            return 0, errors.New("Unknown CodespacesPostRequestBody_geo value: " + v)
    }
    return &result, nil
}
func SerializeCodespacesPostRequestBody_geo(values []CodespacesPostRequestBody_geo) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i CodespacesPostRequestBody_geo) isMultiValue() bool {
    return false
}
