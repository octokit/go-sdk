package headers

const AuthorizationKey = "Authorization"
const AuthType = "bearer"
const UserAgentKey = "User-Agent"

// TODO(kfcampbell): get the version and binary name from build settings rather than hard-coding
const UserAgentValue = "go-sdk@v0.0.0"

const APIVersionKey = "X-GitHub-Api-Version"

// TODO(kfcampbell): get the version from the generated code somehow
const APIVersionValue = "2022-11-28"

const XRateLimitRemainingKey = "X-Ratelimit-Remaining"
const XRateLimitResetKey = "X-Ratelimit-Reset"
const RetryAfterKey = "Retry-After"
