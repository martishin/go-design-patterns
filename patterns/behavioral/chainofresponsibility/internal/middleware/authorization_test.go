package middleware_test

import (
	"testing"

	"github.com/martishin/go-design-patterns/patterns/behavioral/chainofresponsibility/internal/middleware"
	"github.com/martishin/go-design-patterns/patterns/behavioral/chainofresponsibility/pkg/api"
)

type authorizationNextSpy struct {
	response api.Response
	calls    int
}

func (s *authorizationNextSpy) SetNext(next api.Handler) api.Handler {
	return next
}

func (s *authorizationNextSpy) Handle(request api.Request) api.Response {
	s.calls++
	return s.response
}

func TestAuthorizationHandler_HandleDelegatesForAllowedRole(t *testing.T) {
	handler := middleware.NewAuthorizationHandler("operator", "admin")
	next := &authorizationNextSpy{response: api.Response{StatusCode: 200, Body: "ok"}}
	handler.SetNext(next)

	got := handler.Handle(api.Request{Role: "operator"})

	if got.StatusCode != 200 {
		t.Fatalf("got status %d, want 200", got.StatusCode)
	}
	if next.calls != 1 {
		t.Fatalf("got calls %d, want 1", next.calls)
	}
}

func TestAuthorizationHandler_HandleStopsForForbiddenRole(t *testing.T) {
	handler := middleware.NewAuthorizationHandler("operator", "admin")
	next := &authorizationNextSpy{response: api.Response{StatusCode: 200, Body: "ok"}}
	handler.SetNext(next)

	got := handler.Handle(api.Request{Role: "viewer"})
	want := api.Response{
		StatusCode: 403,
		Body:       "forbidden",
	}
	if got != want {
		t.Fatalf("got response %+v, want %+v", got, want)
	}
	if next.calls != 0 {
		t.Fatalf("got calls %d, want 0", next.calls)
	}
}
