package authentication

import (
	"fmt"

	abs "github.com/microsoft/kiota-abstractions-go"
)

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

// WithAuthorization sets the Authorization header to the given token,
// prepended by the AuthType
func (r *Request) WithAuthorization(token string) {
	if r.Headers.ContainsKey(AuthorizationKey) {
		r.Headers.Remove(AuthorizationKey)
	}
	r.Headers.Add(AuthorizationKey, fmt.Sprintf("%v %v", AuthType, token))
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

// WithAPIVersion sets the API version header for each request
func (r *Request) WithAPIVersion(version string) {
	if r.Headers.ContainsKey(APIVersionKey) {
		r.Headers.Remove(APIVersionKey)
	}
	r.Headers.Add(APIVersionKey, version)
}

// WithDefaultAPIVersion sets the API version header to the default (the version used
// to generate the code) for each request
func (r *Request) WithDefaultAPIVersion() {
	r.WithAPIVersion(APIVersionValue)
}
