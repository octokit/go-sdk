package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	netHttp "net/http"
	"strconv"
	"time"

	abs "github.com/microsoft/kiota-abstractions-go"
	kiotaHttp "github.com/microsoft/kiota-http-go"
)

type RateLimitHandler struct {
	options RateLimitHandlerOptions
}

type RateLimitHandlerOptions struct {
}

type RateLimitType int

const (
	None RateLimitType = iota
	Primary
	Secondary
)

// TODO(kfcampbell): should IsRateLimited be a method on RateLimitHandlerOptions? it feels like this could be
// better expressed on the handler itself
type rateLimitHandlerOptionsInt interface {
	abs.RequestOption
	IsRateLimited() func(req *netHttp.Request, res *netHttp.Response) RateLimitType
}

var rateLimitKeyValue = abs.RequestOptionKey{
	Key: "RateLimitHandler",
}

func (options *RateLimitHandlerOptions) GetKey() abs.RequestOptionKey {
	return rateLimitKeyValue
}

func (options *RateLimitHandlerOptions) IsRateLimited() func(req *netHttp.Request, resp *netHttp.Response) RateLimitType {
	return func(req *netHttp.Request, resp *netHttp.Response) RateLimitType {
		if resp.StatusCode != 429 && resp.StatusCode != 403 {
			return None
		}

		// TODO(kfcampbell): i don't think that secondary rate limits actually come with
		// x-ratelimit-remaining headers; additional validation is required here
		if resp.Header.Get("Retry-After") != "" /*&& resp.Header.Get("x-ratelimit-remaining") != "0"*/ {
			return Secondary // secondary rate limits are abuse limits
		}

		if resp.Header.Get("x-ratelimit-remaining") == "0" {
			return Primary
		}

		return None
	}
}

func NewRateLimitHandler() *RateLimitHandler {
	return NewRateLimitHandlerWithOptions(RateLimitHandlerOptions{})
}

func NewRateLimitHandlerWithOptions(options RateLimitHandlerOptions) *RateLimitHandler {
	return &RateLimitHandler{options: options}
}

func (handler RateLimitHandler) Intercept(pipeline kiotaHttp.Pipeline, middlewareIndex int, request *netHttp.Request) (*netHttp.Response, error) {
	resp, err := pipeline.Next(request, middlewareIndex)
	if err != nil {
		return resp, err
	}

	// temp: json-stringify and log response headers so they may be saved and captured
	respJson, err := json.Marshal(resp.Header)
	if err != nil {
		log.Printf("failed to marshal response: %v", err)
	} else {
		log.Printf("response: %s", string(respJson))
	}

	rateLimit := handler.options.IsRateLimited()(request, resp)

	if rateLimit == Primary || rateLimit == Secondary {
		reqOption, ok := request.Context().Value(rateLimitKeyValue).(rateLimitHandlerOptionsInt)
		if !ok {
			reqOption = &handler.options
		}
		return handler.retryRequest(request.Context(), pipeline, middlewareIndex, reqOption, rateLimit, request, resp)
	}
	return resp, nil
}

func (handler RateLimitHandler) retryRequest(ctx context.Context, pipeline kiotaHttp.Pipeline, middlewareIndex int,
	options rateLimitHandlerOptionsInt, rateLimitType RateLimitType, request *netHttp.Request, resp *netHttp.Response) (*netHttp.Response, error) {

	if rateLimitType == Secondary || rateLimitType == Primary {
		retryAfterDuration, err := parseRateLimit(resp)
		if err != nil {
			return nil, fmt.Errorf("failed to parse retry-after header into duration: %v", err)
		}
		if *retryAfterDuration < 0 {
			log.Printf("retry-after duration is negative: %s", *retryAfterDuration)
		}
		if rateLimitType == Secondary {
			log.Printf("Abuse detection mechanism (secondary rate limit) triggered, sleeping for %s before retrying\n", *retryAfterDuration)
		} else if rateLimitType == Primary {
			log.Printf("Primary rate limit %s reached, sleeping for %s before retrying\n", resp.Header.Get("x-ratelimit-limit"), *retryAfterDuration)
		}
		time.Sleep(*retryAfterDuration)
		log.Printf("Retrying request after rate limit sleep\n")
		return handler.Intercept(pipeline, middlewareIndex, request)
	}

	return handler.retryRequest(ctx, pipeline, middlewareIndex, options, rateLimitType, request, resp)
}

// code stolen from https://github.com/google/go-github/blob/0e3ab5807f0e9bc6ea690f1b49e94b78259f3681/github/github.go#L1096
// TODO(kfcampbell): validate/give credit/import appropriately if possible
// parseRateLimit parses the secondary rate related headers,
// and returns the time to retry after.
func parseRateLimit(r *netHttp.Response) (*time.Duration, error) {
	// Retry-After corresponds to secondary rate limits and x-ratelimit-reset to primary rate limits.

	// According to GitHub support, the "Retry-After" header value will be
	// an integer which represents the number of seconds that one should
	// wait before resuming making requests.
	if v := r.Header.Get("Retry-After"); v != "" {
		return parseRetryAfter(v)
	}

	// According to GitHub support, endpoints might return x-ratelimit-reset instead,
	// as an integer which represents the number of seconds since epoch UTC,
	// representing the time to resume making requests.
	if v := r.Header.Get("X-RateLimit-Reset"); v != "" {
		secondsSinceEpoch, err := strconv.ParseInt(v, 10, 64) // Error handling is noop.
		if err != nil {
			return nil, fmt.Errorf("failed to parse x-ratelimit-reset header into duration: %v", err)
		}
		retryAfter := time.Until(time.Unix(secondsSinceEpoch, 0))
		return &retryAfter, nil
	}

	return nil, nil
}

func parseRetryAfter(retryAfterValue string) (*time.Duration, error) {
	if retryAfterValue == "" {
		return nil, fmt.Errorf("could not parse emtpy RetryAfter string")
	}

	retryAfterSeconds, err := strconv.ParseInt(retryAfterValue, 10, 64) // Error handling is noop.
	if err != nil {
		return nil, fmt.Errorf("failed to parse retry-after header into duration: %v", err)
	}
	retryAfter := time.Duration(retryAfterSeconds) * time.Second
	if retryAfter < 0 {
		return nil, fmt.Errorf("retry-after duration is negative: %s, retryAfterValue: %s", retryAfter, retryAfterValue)
	}
	return &retryAfter, nil
}