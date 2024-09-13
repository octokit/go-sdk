package migrations
// Allowed values that can be passed to the exclude param.
type MigrationsPostRequestBody_exclude int

const (
    REPOSITORIES_MIGRATIONSPOSTREQUESTBODY_EXCLUDE MigrationsPostRequestBody_exclude = iota
)

func (i MigrationsPostRequestBody_exclude) String() string {
    return []string{"repositories"}[i]
}
func ParseMigrationsPostRequestBody_exclude(v string) (any, error) {
    result := REPOSITORIES_MIGRATIONSPOSTREQUESTBODY_EXCLUDE
    switch v {
        case "repositories":
            result = REPOSITORIES_MIGRATIONSPOSTREQUESTBODY_EXCLUDE
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeMigrationsPostRequestBody_exclude(values []MigrationsPostRequestBody_exclude) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i MigrationsPostRequestBody_exclude) isMultiValue() bool {
    return false
}
