package api

type Request struct {
	Path     string
	Method   string
	Token    string
	Role     string
	ClientIP string
}

type Response struct {
	StatusCode int
	Body       string
}

type Handler interface {
	SetNext(handler Handler) Handler
	Handle(request Request) Response
}

type BaseHandler struct {
	next Handler
}

func (h *BaseHandler) SetNext(next Handler) Handler {
	h.next = next
	return next
}

func (h *BaseHandler) HandleNext(request Request) Response {
	if h.next == nil {
		return Response{
			StatusCode: 404,
			Body:       "request was not handled",
		}
	}

	return h.next.Handle(request)
}
