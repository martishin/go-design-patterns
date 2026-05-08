package middleware

import "github.com/martishin/go-design-patterns/patterns/behavioral/chainofresponsibility/pkg/api"

type AuthorizationHandler struct {
	api.BaseHandler
	allowedRoles map[string]struct{}
}

func NewAuthorizationHandler(roles ...string) *AuthorizationHandler {
	allowedRoles := make(map[string]struct{}, len(roles))
	for _, role := range roles {
		allowedRoles[role] = struct{}{}
	}

	return &AuthorizationHandler{
		allowedRoles: allowedRoles,
	}
}

func (h *AuthorizationHandler) Handle(request api.Request) api.Response {
	if _, exists := h.allowedRoles[request.Role]; !exists {
		return api.Response{
			StatusCode: 403,
			Body:       "forbidden",
		}
	}

	return h.HandleNext(request)
}
