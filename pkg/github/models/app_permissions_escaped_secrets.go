package models
// The level of permission to grant the access token to manage repository secrets.
type AppPermissions_secrets int

const (
    READ_APPPERMISSIONS_SECRETS AppPermissions_secrets = iota
    WRITE_APPPERMISSIONS_SECRETS
)

func (i AppPermissions_secrets) String() string {
    return []string{"read", "write"}[i]
}
func ParseAppPermissions_secrets(v string) (any, error) {
    result := READ_APPPERMISSIONS_SECRETS
    switch v {
        case "read":
            result = READ_APPPERMISSIONS_SECRETS
        case "write":
            result = WRITE_APPPERMISSIONS_SECRETS
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeAppPermissions_secrets(values []AppPermissions_secrets) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i AppPermissions_secrets) isMultiValue() bool {
    return false
}
