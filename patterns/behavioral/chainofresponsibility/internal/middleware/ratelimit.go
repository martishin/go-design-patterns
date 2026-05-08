package middleware

import "github.com/martishin/go-design-patterns/patterns/behavioral/chainofresponsibility/pkg/api"

type RateLimitHandler struct {
	api.BaseHandler
	maxRequests int
	requests    map[string]int
}

func NewRateLimitHandler(maxRequests int) *RateLimitHandler {
	return &RateLimitHandler{
		maxRequests: maxRequests,
		requests:    make(map[string]int),
	}
}

func (h *RateLimitHandler) Handle(request api.Request) api.Response {
	h.requests[request.ClientIP]++
	if h.requests[request.ClientIP] > h.maxRequests {
		return api.Response{
			StatusCode: 429,
			Body:       "too many requests",
		}
	}

	return h.HandleNext(request)
}
