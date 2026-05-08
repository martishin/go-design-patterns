package api_test

import (
	"testing"

	"github.com/martishin/go-design-patterns/patterns/behavioral/chainofresponsibility/pkg/api"
)

type handlerSpy struct {
	gotRequest api.Request
	response   api.Response
	calls      int
}

func (s *handlerSpy) SetNext(next api.Handler) api.Handler {
	return next
}

func (s *handlerSpy) Handle(request api.Request) api.Response {
	s.gotRequest = request
	s.calls++
	return s.response
}

func TestBaseHandler_SetNextReturnsNextHandler(t *testing.T) {
	var base api.BaseHandler
	next := &handlerSpy{}

	got := base.SetNext(next)
	if got != next {
		t.Fatalf("got next handler %v, want %v", got, next)
	}
}

func TestBaseHandler_HandleNextDelegatesToNextHandler(t *testing.T) {
	request := api.Request{
		Path:   "/app/status",
		Method: "GET",
	}
	want := api.Response{
		StatusCode: 200,
		Body:       "ok",
	}
	next := &handlerSpy{response: want}

	var base api.BaseHandler
	base.SetNext(next)

	got := base.HandleNext(request)
	if got != want {
		t.Fatalf("got response %+v, want %+v", got, want)
	}
	if next.calls != 1 {
		t.Fatalf("got calls %d, want 1", next.calls)
	}
	if next.gotRequest != request {
		t.Fatalf("got request %+v, want %+v", next.gotRequest, request)
	}
}

func TestBaseHandler_HandleNextReturnsUnhandledResponseWithoutNextHandler(t *testing.T) {
	var base api.BaseHandler

	got := base.HandleNext(api.Request{})
	want := api.Response{
		StatusCode: 404,
		Body:       "request was not handled",
	}
	if got != want {
		t.Fatalf("got response %+v, want %+v", got, want)
	}
}
