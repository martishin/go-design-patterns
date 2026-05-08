package middleware_test

import (
	"testing"

	"github.com/martishin/go-design-patterns/patterns/behavioral/chainofresponsibility/internal/middleware"
	"github.com/martishin/go-design-patterns/patterns/behavioral/chainofresponsibility/pkg/api"
)

func TestApplicationHandler_HandleReturnsStatusResponse(t *testing.T) {
	handler := middleware.NewApplicationHandler()

	got := handler.Handle(api.Request{
		Path:   "/app/status",
		Method: "GET",
	})
	want := api.Response{
		StatusCode: 200,
		Body:       "ok",
	}
	if got != want {
		t.Fatalf("got response %+v, want %+v", got, want)
	}
}

func TestApplicationHandler_HandleReturnsDeploymentResponse(t *testing.T) {
	handler := middleware.NewApplicationHandler()

	got := handler.Handle(api.Request{
		Path:   "/deployments",
		Method: "POST",
	})
	want := api.Response{
		StatusCode: 201,
		Body:       "deployment created",
	}
	if got != want {
		t.Fatalf("got response %+v, want %+v", got, want)
	}
}

func TestApplicationHandler_HandleReturnsNotFoundResponse(t *testing.T) {
	handler := middleware.NewApplicationHandler()

	got := handler.Handle(api.Request{
		Path:   "/missing",
		Method: "GET",
	})
	want := api.Response{
		StatusCode: 404,
		Body:       "not found",
	}
	if got != want {
		t.Fatalf("got response %+v, want %+v", got, want)
	}
}
