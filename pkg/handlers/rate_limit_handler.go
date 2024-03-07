package handlers

import (
	"context"
	"fmt"
	netHttp "net/http"
	"time"

	abs "github.com/microsoft/kiota-abstractions-go"
	kiotaHttp "github.com/microsoft/kiota-http-go"
)

type RateLimitHandler struct {
	options RateLimitHandlerOptions
}

type RateLimitHandlerOptions struct {
	ReadDelay        time.Duration
	WriteDelay       time.Duration
	ParallelRequests bool
}

type RateLimitType int

const (
	None RateLimitType = iota
	Primary
	Secondary
)

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

// TODO(kfcampbell): should this return an iota or something that makes it clear whether it's
// a primary or secondary rate limit we hit?
func (options *RateLimitHandlerOptions) IsRateLimited() func(req *netHttp.Request, resp *netHttp.Response) RateLimitType {
	// TODO(kfcampbell): validate this method
	return func(req *netHttp.Request, resp *netHttp.Response) RateLimitType {
		if resp.StatusCode != 429 && resp.StatusCode != 403 {
			return None
		}

		// abuse (secondary) rate limit situation
		if resp.Header.Get("Retry-After") != "" {
			return Secondary
		}

		// primary rate limit situation
		if resp.Header.Get("x-ratelimit-remaining") == "0" {
			return Primary
		}

		return None
	}
}

func NewRateLimitHandler() *RateLimitHandler {
	return NewRateLimitHandlerWithOptions(RateLimitHandlerOptions{
		ReadDelay:        0 * time.Second,
		WriteDelay:       1 * time.Second,
		ParallelRequests: false,
	})
}

func NewRateLimitHandlerWithOptions(options RateLimitHandlerOptions) *RateLimitHandler {
	return &RateLimitHandler{options: options}
}

func (handler RateLimitHandler) Intercept(pipeline kiotaHttp.Pipeline, middlewareIndex int, request *netHttp.Request) (*netHttp.Response, error) {
	resp, err := pipeline.Next(request, middlewareIndex)
	if err != nil {
		return resp, err
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

// TODO(kfcampbell): helper method to wait and retry requests after a specified period of time
func (handler RateLimitHandler) retryRequest(ctx context.Context, pipeline kiotaHttp.Pipeline, middlewareIndex int,
	options rateLimitHandlerOptionsInt, rateLimitType RateLimitType, request *netHttp.Request, resp *netHttp.Response) (*netHttp.Response, error) {

	if rateLimitType == Secondary {
		fmt.Printf("<<<<<<< Abuse detection mechanism triggered, sleeping for %s before retrying\n", resp.Header.Get("Retry-After"))
		retryAfter := resp.Header.Get("Retry-After")
		retryAfterDuration, err := time.ParseDuration(retryAfter + "s")
		if err != nil {
			return nil, fmt.Errorf("failed to parse retry-after header into duration: %v", err)
		}
		time.Sleep(retryAfterDuration)
		return handler.Intercept(pipeline, middlewareIndex, request)
	}

	if rateLimitType == Primary {
		fmt.Printf("<<<<<<< Primary rate limit %s reached, sleeping for %s before retrying\n", resp.Header.Get("x-ratelimit-limit"), resp.Header.Get("Retry-After"))
	}
	panic("not implemented")
}
