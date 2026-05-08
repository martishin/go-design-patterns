package main

import (
	"fmt"

	"github.com/martishin/go-design-patterns/patterns/behavioral/chainofresponsibility/internal/middleware"
	"github.com/martishin/go-design-patterns/patterns/behavioral/chainofresponsibility/pkg/api"
)

func main() {
	rateLimit := middleware.NewRateLimitHandler(1)
	authentication := middleware.NewAuthenticationHandler("deploy-token")
	authorization := middleware.NewAuthorizationHandler("operator", "admin")
	application := middleware.NewApplicationHandler()

	rateLimit.
		SetNext(authentication).
		SetNext(authorization).
		SetNext(application)

	request := api.Request{
		Path:     "/app/status",
		Method:   "GET",
		Token:    "deploy-token",
		Role:     "operator",
		ClientIP: "10.0.0.7",
	}

	firstResponse := rateLimit.Handle(request)
	secondResponse := rateLimit.Handle(request)
	unauthenticatedResponse := rateLimit.Handle(api.Request{
		Path:     "/deployments",
		Method:   "POST",
		Token:    "bad-token",
		Role:     "operator",
		ClientIP: "10.0.0.8",
	})
	forbiddenResponse := rateLimit.Handle(api.Request{
		Path:     "/deployments",
		Method:   "POST",
		Token:    "deploy-token",
		Role:     "viewer",
		ClientIP: "10.0.0.9",
	})

	fmt.Printf("First status request: %d %s\n", firstResponse.StatusCode, firstResponse.Body)
	fmt.Printf("Second status request: %d %s\n", secondResponse.StatusCode, secondResponse.Body)
	fmt.Printf("Unauthenticated deployment request: %d %s\n", unauthenticatedResponse.StatusCode, unauthenticatedResponse.Body)
	fmt.Printf("Forbidden deployment request: %d %s\n", forbiddenResponse.StatusCode, forbiddenResponse.Body)
}
