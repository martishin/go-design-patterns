package middleware_test

import (
	"testing"

	"github.com/martishin/go-design-patterns/patterns/behavioral/chainofresponsibility/internal/middleware"
	"github.com/martishin/go-design-patterns/patterns/behavioral/chainofresponsibility/pkg/api"
)

type authenticationNextSpy struct {
	response api.Response
	calls    int
}

func (s *authenticationNextSpy) SetNext(next api.Handler) api.Handler {
	return next
}

func (s *authenticationNextSpy) Handle(request api.Request) api.Response {
	s.calls++
	return s.response
}

func TestAuthenticationHandler_HandleDelegatesForValidToken(t *testing.T) {
	handler := middleware.NewAuthenticationHandler("deploy-token")
	next := &authenticationNextSpy{response: api.Response{StatusCode: 200, Body: "ok"}}
	handler.SetNext(next)

	got := handler.Handle(api.Request{Token: "deploy-token"})

	if got.StatusCode != 200 {
		t.Fatalf("got status %d, want 200", got.StatusCode)
	}
	if next.calls != 1 {
		t.Fatalf("got calls %d, want 1", next.calls)
	}
}

func TestAuthenticationHandler_HandleStopsForInvalidToken(t *testing.T) {
	handler := middleware.NewAuthenticationHandler("deploy-token")
	next := &authenticationNextSpy{response: api.Response{StatusCode: 200, Body: "ok"}}
	handler.SetNext(next)

	got := handler.Handle(api.Request{Token: "bad-token"})
	want := api.Response{
		StatusCode: 401,
		Body:       "unauthorized",
	}
	if got != want {
		t.Fatalf("got response %+v, want %+v", got, want)
	}
	if next.calls != 0 {
		t.Fatalf("got calls %d, want 0", next.calls)
	}
}
