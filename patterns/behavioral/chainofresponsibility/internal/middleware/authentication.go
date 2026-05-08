package middleware

import "github.com/martishin/go-design-patterns/patterns/behavioral/chainofresponsibility/pkg/api"

type AuthenticationHandler struct {
	api.BaseHandler
	validTokens map[string]struct{}
}

func NewAuthenticationHandler(tokens ...string) *AuthenticationHandler {
	validTokens := make(map[string]struct{}, len(tokens))
	for _, token := range tokens {
		validTokens[token] = struct{}{}
	}

	return &AuthenticationHandler{
		validTokens: validTokens,
	}
}

func (h *AuthenticationHandler) Handle(request api.Request) api.Response {
	if _, exists := h.validTokens[request.Token]; !exists {
		return api.Response{
			StatusCode: 401,
			Body:       "unauthorized",
		}
	}

	return h.HandleNext(request)
}
