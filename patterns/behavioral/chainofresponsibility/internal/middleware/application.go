package middleware

import "github.com/martishin/go-design-patterns/patterns/behavioral/chainofresponsibility/pkg/api"

type ApplicationHandler struct{}

func NewApplicationHandler() *ApplicationHandler {
	return &ApplicationHandler{}
}

func (h *ApplicationHandler) SetNext(next api.Handler) api.Handler {
	return next
}

func (h *ApplicationHandler) Handle(request api.Request) api.Response {
	if request.Path == "/app/status" && request.Method == "GET" {
		return api.Response{
			StatusCode: 200,
			Body:       "ok",
		}
	}

	if request.Path == "/deployments" && request.Method == "POST" {
		return api.Response{
			StatusCode: 201,
			Body:       "deployment created",
		}
	}

	return api.Response{
		StatusCode: 404,
		Body:       "not found",
	}
}
