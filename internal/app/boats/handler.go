package boats

import (
	"context"
	"net/http"
)

type Handler struct {
	doHandle func(ctx context.Context, w http.ResponseWriter, r *http.Request)
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.doHandle(context.Background(), w, r)
}

func (s *BoatsServer) Handler(f func(ctx context.Context, w http.ResponseWriter, r *http.Request)) *Handler {
	return &Handler{f}
}
