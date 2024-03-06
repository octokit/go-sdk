package handlers

import (
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

type rateLimitHandlerOptionsInt interface {
	abs.RequestOption
	IsRateLimited() func(req *netHttp.Request, res *netHttp.Response) bool
}

var rateLimitKeyValue = abs.RequestOptionKey{
	Key: "RateLimitHandler",
}

func (options *RateLimitHandlerOptions) GetKey() abs.RequestOptionKey {
	return rateLimitKeyValue
}

func (options *RateLimitHandlerOptions) IsRateLimited() func(req *netHttp.Request, res *netHttp.Response) bool {
	// TODO(kfcampbell): implement this method
	return func(req *netHttp.Request, res *netHttp.Response) bool {
		return false
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

func (RateLimitHandler) Intercept(pipeline kiotaHttp.Pipeline, middlewareIndex int, request *netHttp.Request) (*netHttp.Response, error) {
	// TODO: implement rate limiting in the intercept method
	panic("implement me")
}

// TODO(kfcampbell): helper method to wait and retry requests after a specified period of time
