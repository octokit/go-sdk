package transport

import (
	"bytes"
	"io"
	"log"
	"net/http"
	"strings"
	"sync"
	"time"
)

// RateLimitTransport and associated code mostly taken from integrations/terraform-provider-github and modified
// WARNING: UNTESTED! this code needs significant validation before it's ready for merging
type RateLimitTransport struct {
	transport        http.RoundTripper
	nextRequestDelay time.Duration
	writeDelay       time.Duration
	readDelay        time.Duration
	parallelRequests bool

	m sync.Mutex
}

func (rlt *RateLimitTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	// Make requests for a single user or client ID serially when parallel_requests is false.
	// If parallel_requests is true skips the lock and allow the parallelism defined by terraform itself.
	rlt.smartLock(true)

	// Sleep for the delay that the last request defined. This delay might be different
	// for read and write requests. See isWriteMethod for the distinction between them.
	if rlt.nextRequestDelay > 0 {
		log.Printf("[DEBUG] Sleeping %s between operations", rlt.nextRequestDelay)
		time.Sleep(rlt.nextRequestDelay)
	}

	rlt.nextRequestDelay = rlt.calculateNextDelay(req.Method)

	resp, err := rlt.transport.RoundTrip(req)
	if err != nil {
		rlt.smartLock(false)
		return resp, err
	}

	// Make response body accessible for retries & debugging
	// See https://github.com/google/go-github/pull/986
	r1, r2, err := drainBody(resp.Body)
	if err != nil {
		return nil, err
	}
	resp.Body = r1
	resp.Body = r2

	// When you have been limited, use the Retry-After response header to slow down.
	// according to google/go-github,
	// AbuseRateLimitError occurs when GitHub returns 403 Forbidden response with the
	// "documentation_url" field value equal to "https://docs.github.com/rest/overview/rate-limits-for-the-rest-api#about-secondary-rate-limits".
	// RetryAfter is provided with some abuse rate limit errors. If present,
	// it is the amount of time that the client should wait before retrying.
	// Otherwise, the client should try again later (after an unspecified amount of time).

	// 403 and 429 are the possible rate-limiting status codes
	// TODO(kfcampbell): this code branch needs validation
	if resp.StatusCode == 403 || resp.StatusCode == 429 {
		var buf bytes.Buffer
		_, err = io.Copy(&buf, resp.Body)
		if err != nil {
			return nil, err
		}
		bufStr := buf.String()
		log.Printf("[DEBUG] Response body: %s", bufStr)
		// TODO(kfcampbell): this code branch needs validation
		if strings.Contains(bufStr, "documentation_url") && strings.Contains(bufStr, "https://docs.github.com/rest/overview/rate-limits-for-the-rest-api#about-secondary-rate-limits") {
			rlt.nextRequestDelay = 0
			// TODO(kfcampbell): this header retrieval and parsing needs validation
			retryAfter := resp.Header.Get("retry-after")
			retryAfterDuration, err := time.ParseDuration(retryAfter + "s")
			if err != nil {
				log.Printf("[WARN] Failed to parse retry-after header: %s", err)
				retryAfterDuration = 1 * time.Second
			}
			log.Printf("[WARN] Abuse detection mechanism triggered, sleeping for %s before retrying",
				retryAfter)
			time.Sleep(retryAfterDuration)
			rlt.smartLock(false)
			return rlt.RoundTrip(req)
		}

		// just a regular rate limit situation, not an abusive one
		// TODO(kfcampbell): this code branch needs validation
		if resp.Header.Get("x-ratelimit-remaining") == "0" {
			rlt.nextRequestDelay = 0
			// TODO(kfcampbell): this header retrieval and parsing needs validation
			// get "retry-after" header (a time value represented in UTC epoch seconds)
			// and convert it to a time.Duration between then and now
			retryAfterTime := resp.Header.Get("retry-after")

			// convert retryAfterTime (a string representing UTC epoch seconds) to a time.Duration
			// representing the difference between then and now
			retryAfter, err := time.ParseDuration(retryAfterTime + "s")
			if err != nil {
				log.Printf("[WARN] Failed to parse retry-after header: %s", err)
				retryAfter = 1 * time.Second // default to 1 second
			}
			retryAfter = retryAfter - time.Since(time.Now().UTC())

			log.Printf("[WARN] Rate limit %s reached, sleeping for %s (until %s) before retrying",
				resp.Header.Get("x-ratelimit-limit"), retryAfter, time.Now().Add(retryAfter))
			time.Sleep(retryAfter)
			rlt.smartLock(false)
			return rlt.RoundTrip(req)
		}
	}
	rlt.smartLock(false)

	return resp, nil
}

// smartLock wraps the mutex locking system and performs its operation via a boolean input for locking and unlocking.
// It also skips the locking when parallelRequests is set to true since, in this case, the lock is not needed.
func (rlt *RateLimitTransport) smartLock(lock bool) {
	if rlt.parallelRequests {
		return
	}
	if lock {
		rlt.m.Lock()
		return
	}
	rlt.m.Unlock()
}

// calculateNextDelay returns a time.Duration specifying the backoff before the next request
// the actual value depends on the current method being a write or a read request
func (rlt *RateLimitTransport) calculateNextDelay(method string) time.Duration {
	if isWriteMethod(method) {
		return rlt.writeDelay
	}
	return rlt.readDelay
}

type RateLimitTransportOption func(*RateLimitTransport)

// NewRateLimitTransport takes in an http.RoundTripper and a variadic list of
// optional functions that modify the RateLimitTransport struct itself. This
// may be used to alter the write delay in between requests, for example.
func NewRateLimitTransport(rt http.RoundTripper, options ...RateLimitTransportOption) *RateLimitTransport {
	// Default to 1 second of write delay if none is provided
	// Default to no read delay if none is provided
	rlt := &RateLimitTransport{transport: rt, writeDelay: 1 * time.Second, readDelay: 0 * time.Second, parallelRequests: false}

	for _, opt := range options {
		opt(rlt)
	}

	return rlt
}

// WithWriteDelay is used to set the write delay between requests
func WithWriteDelay(d time.Duration) RateLimitTransportOption {
	return func(rlt *RateLimitTransport) {
		rlt.writeDelay = d
	}
}

// WithReadDelay is used to set the delay between read requests
func WithReadDelay(d time.Duration) RateLimitTransportOption {
	return func(rlt *RateLimitTransport) {
		rlt.readDelay = d
	}
}

// WithParallelRequests is used to enforce serial api requests for rate limits
func WithParallelRequests(p bool) RateLimitTransportOption {
	return func(rlt *RateLimitTransport) {
		rlt.parallelRequests = p
	}
}

// drainBody reads all of b to memory and then returns two equivalent
// ReadClosers yielding the same bytes.
func drainBody(b io.ReadCloser) (r1, r2 io.ReadCloser, err error) {
	if b == http.NoBody {
		// No copying needed. Preserve the magic sentinel meaning of NoBody.
		return http.NoBody, http.NoBody, nil
	}
	var buf bytes.Buffer
	if _, err = buf.ReadFrom(b); err != nil {
		return nil, b, err
	}
	if err = b.Close(); err != nil {
		return nil, b, err
	}
	return io.NopCloser(&buf), io.NopCloser(bytes.NewReader(buf.Bytes())), nil
}

func isWriteMethod(method string) bool {
	switch method {
	case "POST", "PATCH", "PUT", "DELETE":
		return true
	}
	return false
}
