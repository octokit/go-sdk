package handlers

import (
	"encoding/json"
	"net/http"
	"testing"
	"time"
)

// paired with 200 status code
var happyPathTestHeaders = `{
	"Access-Control-Allow-Origin": [
		"*"
	],
	"Access-Control-Expose-Headers": [
		"ETag, Link, Location, Retry-After, X-GitHub-OTP, X-RateLimit-Limit, X-RateLimit-Remaining, X-RateLimit-Used, X-RateLimit-Resource, X-RateLimit-Reset, X-OAuth-Scopes, X-Accepted-OAuth-Scopes, X-Poll-Interval, X-GitHub-Media-Type, X-GitHub-SSO, X-GitHub-Request-Id, Deprecation, Sunset"
	],
	"Cache-Control": [
		"private, max-age=60, s-maxage=60"
	],
	"Connection": [
		"keep-alive"
	],
	"Content-Security-Policy": [
		"default-src 'none'"
	],
	"Content-Type": [
		"application/json; charset=utf-8"
	],
	"Date": [
		"Thu, 28 Mar 2024 20:20:55 GMT"
	],
	"Etag": [
		"W/\"513ca5822992932a02050bbecb8c58a75cc211fb7ef0d597b6a3119c9fb45e94\""
	],
	"Github-Authentication-Token-Expiration": [
		"2024-04-27 20:14:21 UTC"
	],
	"Referrer-Policy": [
		"origin-when-cross-origin, strict-origin-when-cross-origin"
	],
	"Server": [
		"nginx/1.18.0 (Ubuntu)"
	],
	"Vary": [
		"Accept, Authorization, Cookie, X-GitHub-OTP",
		"Accept, X-Requested-With"
	],
	"X-Accepted-Oauth-Scopes": [
		""
	],
	"X-Content-Type-Options": [
		"nosniff"
	],
	"X-Frame-Options": [
		"deny"
	],
	"X-Github-Api-Version-Selected": [
		"2022-11-28"
	],
	"X-Github-Media-Type": [
		"github.v3"
	],
	"X-Github-Request-Id": [
		"488c61a3-2ad6-4182-a16e-299d4cbd5c3c"
	],
	"X-Github-Sso": [
		"partial-results; organizations=136,137,138"
	],
	"X-Glb-Log-Append": [
		"rails_request_category=api rails_method=get rails_controller=api_repositories rails_action=/user/repos rails_catalog_service=github/repos rails_request_queued_time=691 rails_request_time=5747"
	],
	"X-Oauth-Scopes": [
		"repo"
	],
	"X-Xss-Protection": [
		"0"
	]
}`

// example primary rate-limited headers (paired with 403 status code)
var primaryRateLimitHeaders = `{
	"Access-Control-Allow-Origin": [
		"*"
	],
	"Access-Control-Expose-Headers": [
		"ETag, Link, Location, Retry-After, X-GitHub-OTP, X-RateLimit-Limit, X-RateLimit-Remaining, X-RateLimit-Used, X-RateLimit-Resource, X-RateLimit-Reset, X-OAuth-Scopes, X-Accepted-OAuth-Scopes, X-Poll-Interval, X-GitHub-Media-Type, X-GitHub-SSO, X-GitHub-Request-Id, Deprecation, Sunset"
	],
	"Connection": [
		"keep-alive"
	],
	"Content-Security-Policy": [
		"default-src 'none'"
	],
	"Content-Type": [
		"application/json; charset=utf-8"
	],
	"Date": [
		"Thu, 28 Mar 2024 20:16:11 GMT"
	],
	"Github-Authentication-Token-Expiration": [
		"2024-04-27 20:14:21 UTC"
	],
	"Referrer-Policy": [
		"origin-when-cross-origin, strict-origin-when-cross-origin"
	],
	"Server": [
		"nginx/1.18.0 (Ubuntu)"
	],
	"Vary": [
		"Accept, X-Requested-With"
	],
	"X-Accepted-Oauth-Scopes": [
		"repo"
	],
	"X-Content-Type-Options": [
		"nosniff"
	],
	"X-Frame-Options": [
		"deny"
	],
	"X-Github-Media-Type": [
		"github.v3"
	],
	"X-Github-Request-Id": [
		"d283bbfe-d024-4334-9bd1-284ff860fe3e"
	],
	"X-Glb-Log-Append": [
		"rails_request_category=api rails_method=get rails_controller=api_repositories rails_action=/user/repos rails_catalog_service=github/repos rails_request_queued_time=513 rails_request_time=13"
	],
	"X-Oauth-Scopes": [
		"repo"
	],
	"X-Ratelimit-Limit": [
		"100"
	],
	"X-Ratelimit-Remaining": [
		"0"
	],
	"X-Ratelimit-Reset": [
		"1711656978"
	],
	"X-Ratelimit-Resource": [
		"core"
	],
	"X-Ratelimit-Used": [
		"103"
	],
	"X-Xss-Protection": [
		"0"
	]
}`

// example secondary rate-limited headers (paired with 403 status code)
var secondaryRateLimitHeaders = `{
	"Access-Control-Allow-Origin": [
		"*"
	],
	"Access-Control-Expose-Headers": [
		"ETag, Link, Location, Retry-After, X-GitHub-OTP, X-RateLimit-Limit, X-RateLimit-Remaining, X-RateLimit-Used, X-RateLimit-Resource, X-RateLimit-Reset, X-OAuth-Scopes, X-Accepted-OAuth-Scopes, X-Poll-Interval, X-GitHub-Media-Type, X-GitHub-SSO, X-GitHub-Request-Id, Deprecation, Sunset"
	],
	"Connection": [
		"keep-alive"
	],
	"Content-Security-Policy": [
		"default-src 'none'; base-uri 'self'; child-src github.localhost/assets-cdn/worker/; connect-src 'self' http://alambic.github.localhost https://www.githubstatus.com http://collector.github.localhost ws://127.0.0.1:35729/livereload ws://github.localhost raw.github.localhost api.github.localhost https://github-cloud.s3.amazonaws.com https://github-development-repository-file-5c1aeb.s3.amazonaws.com https://github-development-upload-manifest-file-7fdce7.s3.amazonaws.com https://github-development-user-asset-6210df.s3.amazonaws.com http://localhost:2206 http://objects-staging-origin.githubusercontent.com; font-src 'self'; form-action 'self' github.localhost copilot-workspace.githubnext.com http://objects-staging-origin.githubusercontent.com; frame-ancestors 'none'; frame-src http://viewscreen.githubusercontent.localhost:9494 http://notebooks.githubusercontent.localhost:8888; img-src 'self' data: http://alambic.github.localhost http://github.localhost:8081 https://identicons.github.com http://alambic.github.localhost/avatars https://github-dev.s3.amazonaws.com https://objects-staging.githubusercontent.com http://localhost:7071 https://github-development-user-asset-6210df.s3.amazonaws.com https://customer-stories-feed.github.com https://spotlights-feed.github.com http://objects-staging-origin.githubusercontent.com; manifest-src 'self'; media-src http://github.localhost http://alambic.github.localhost https://github-development-user-asset-6210df.s3.amazonaws.com; script-src 'self'; style-src 'unsafe-inline' 'self'; worker-src github.localhost/assets-cdn/worker/"
	],
	"Content-Type": [
		"application/json; charset=utf-8"
	],
	"Date": [
		"Thu, 28 Mar 2024 20:22:54 GMT"
	],
	"Gh-Limited-By": [
		"time-based"
	],
	"Gh-Limited-Group": [
		"api"
	],
	"Referrer-Policy": [
		"origin-when-cross-origin, strict-origin-when-cross-origin"
	],
	"Retry-After": [
		"60"
	],
	"Server": [
		"nginx/1.18.0 (Ubuntu)"
	],
	"Vary": [
		"Accept, X-Requested-With"
	],
	"X-Content-Type-Options": [
		"nosniff"
	],
	"X-Frame-Options": [
		"deny"
	],
	"X-Github-Media-Type": [
		"github.v3; format=json"
	],
	"X-Github-Request-Id": [
		"47dae098-2903-4a0a-835d-d7de10368f0f"
	],
	"X-Glb-Log-Append": [
		"rails_request_category=api rails_method=get rails_controller=api_repositories rails_action=/user/repos rails_catalog_service=github/rest_api rails_request_queued_time=4 rails_request_time=12"
	],
	"X-Xss-Protection": [
		"0"
	]
}`

func TestParseRateLimitHappyPath(t *testing.T) {
	resp, err := setupHeaderMap(happyPathTestHeaders, 200)
	if err != nil {
		t.Errorf("Failed to set up headers: %v", err)
	}

	result, err := parseRateLimit(resp)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if result != nil {
		t.Errorf("Expected nil result, got %v", result)
	}
}

func TestParseRateLimitPrimaryRateLimit(t *testing.T) {
	resp, err := setupHeaderMap(primaryRateLimitHeaders, 403)
	if err != nil {
		t.Errorf("Failed to set up headers: %v", err)
	}

	result, err := parseRateLimit(resp)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	// TODO(kfcampbell): fix tests to use fixed times, not current times (will fix negative time issue in tests)
	if result == nil {
		t.Errorf("Expected non-nil result, got %v", result)
	}
}

func TestParseRateLimitSecondaryRateLimit(t *testing.T) {
	resp, err := setupHeaderMap(secondaryRateLimitHeaders, 403)
	if err != nil {
		t.Errorf("Failed to set up headers: %v", err)
	}

	result, err := parseRateLimit(resp)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if *result != time.Minute {
		t.Errorf("Expected one minute, got %v", result)
	}
}

func TestIsRateLimitedHappyPath(t *testing.T) {
	resp, err := setupHeaderMap(happyPathTestHeaders, 200)
	if err != nil {
		t.Errorf("Failed to set up headers: %v", err)
	}

	handler := NewRateLimitHandler()

	rateLimit := handler.options.IsRateLimited()(nil, resp)
	if rateLimit != None {
		t.Errorf("Expected no rate limit, got %v", rateLimit)
	}
}

func TestIsRateLimitedPrimaryRateLimit(t *testing.T) {
	resp, err := setupHeaderMap(primaryRateLimitHeaders, 403)
	if err != nil {
		t.Errorf("Failed to set up headers: %v", err)
	}

	handler := NewRateLimitHandler()

	rateLimit := handler.options.IsRateLimited()(nil, resp)
	if rateLimit != Primary {
		t.Errorf("Expected primary rate limit, got %v", rateLimit)
	}
}

func TestIsRateLimitedSecondaryRateLimit(t *testing.T) {
	resp, err := setupHeaderMap(secondaryRateLimitHeaders, 403)
	if err != nil {
		t.Errorf("Failed to set up headers: %v", err)
	}

	handler := NewRateLimitHandler()

	rateLimit := handler.options.IsRateLimited()(nil, resp)
	if rateLimit != Secondary {
		t.Errorf("Expected secondary rate limit, got %v", rateLimit)
	}
}

func TestParseRetryAfterBlankString(t *testing.T) {
	retryAfterValue := ""
	result, err := parseRetryAfter(retryAfterValue)
	if err == nil {
		t.Errorf("Expected error, got %v", result)
	}
}

func TestParseRetryAfterInvalidString(t *testing.T) {
	retryAfterValue := "xxx_invalid_string_xxx"
	result, err := parseRetryAfter(retryAfterValue)
	if err == nil {
		t.Errorf("Expected error, got %v", result)
	}
}

func TestParseRetryAfterNegativeValue(t *testing.T) {
	retryAfterValue := "-1"
	result, err := parseRetryAfter(retryAfterValue)
	if err == nil {
		t.Errorf("Expected error, got %v", result)
	}
}

// setupHeaderMap is a utility function that takes in a JSON string of headers and
// a status code and returns a hydrated http.Response object.
func setupHeaderMap(headers string, statusCode int) (*http.Response, error) {
	headerMap := make(map[string][]string)
	err := json.Unmarshal([]byte(headers), &headerMap)
	if err != nil {
		return nil, err
	}

	resp := &http.Response{
		Header:     make(http.Header),
		StatusCode: statusCode,
	}

	for key, values := range headerMap {
		for _, value := range values {
			resp.Header.Add(key, value)
		}
	}

	return resp, nil
}
