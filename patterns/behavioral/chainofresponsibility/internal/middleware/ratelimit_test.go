package middleware_test

import (
	"testing"

	"github.com/martishin/go-design-patterns/patterns/behavioral/chainofresponsibility/internal/middleware"
	"github.com/martishin/go-design-patterns/patterns/behavioral/chainofresponsibility/pkg/api"
)

type rateLimitNextSpy struct {
	response api.Response
	calls    int
}

func (s *rateLimitNextSpy) SetNext(next api.Handler) api.Handler {
	return next
}

func (s *rateLimitNextSpy) Handle(request api.Request) api.Response {
	s.calls++
	return s.response
}

func TestRateLimitHandler_HandleDelegatesWithinLimit(t *testing.T) {
	handler := middleware.NewRateLimitHandler(2)
	next := &rateLimitNextSpy{response: api.Response{StatusCode: 200, Body: "ok"}}
	handler.SetNext(next)

	request := api.Request{ClientIP: "10.0.0.7"}

	first := handler.Handle(request)
	second := handler.Handle(request)

	if first.StatusCode != 200 {
		t.Fatalf("got first status %d, want 200", first.StatusCode)
	}
	if second.StatusCode != 200 {
		t.Fatalf("got second status %d, want 200", second.StatusCode)
	}
	if next.calls != 2 {
		t.Fatalf("got calls %d, want 2", next.calls)
	}
}

func TestRateLimitHandler_HandleStopsWhenLimitIsExceeded(t *testing.T) {
	handler := middleware.NewRateLimitHandler(1)
	next := &rateLimitNextSpy{response: api.Response{StatusCode: 200, Body: "ok"}}
	handler.SetNext(next)

	request := api.Request{ClientIP: "10.0.0.7"}

	_ = handler.Handle(request)
	got := handler.Handle(request)

	want := api.Response{
		StatusCode: 429,
		Body:       "too many requests",
	}
	if got != want {
		t.Fatalf("got response %+v, want %+v", got, want)
	}
	if next.calls != 1 {
		t.Fatalf("got calls %d, want 1", next.calls)
	}
}
