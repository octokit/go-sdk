package authentication

import abs "github.com/microsoft/kiota-abstractions-go"

const AuthorizationKey = "Authorization"
const AuthType = "bearer"
const UserAgentKey = "User-Agent"

// TODO(kfcampbell): get the version and binary name from build settings rather than hard-coding
const UserAgentValue = "go-sdk@v0.0.0"

const APIVersionKey = "X-GitHub-Api-Version"

// TODO(kfcampbell): get the version from the generated code somehow
const APIVersionValue = "2022-11-28"

// Request provides a wrapper around Kiota's abs.RequestInformation type
type Request struct {
	*abs.RequestInformation
}

// WithUserAgent allows the caller to set the User-Agent string for each request
func (r *Request) WithUserAgent(userAgent string) {
	if r.Headers.ContainsKey(UserAgentKey) {
		r.Headers.Remove(UserAgentKey)
	}
	r.Headers.Add(UserAgentKey, userAgent)
}

// WithDefaultUserAgent sets the default User-Agent string for each request
func (r *Request) WithDefaultUserAgent() {
	r.WithUserAgent(UserAgentValue)
}
