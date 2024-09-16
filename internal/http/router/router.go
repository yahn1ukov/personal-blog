package router

import (
	"net/http"

	"github.com/yahn1ukov/personal-blog/internal/http/handler"
)

type Router struct {
	*http.ServeMux
}

func New(handler *handler.Handler) *Router {
	mux := http.NewServeMux()

	mux.HandleFunc("POST /blogs", handler.Create)
	mux.HandleFunc("GET /blogs", handler.GetAll)
	mux.HandleFunc("GET /blogs/{id}", handler.GetByID)
	mux.HandleFunc("PATCH /blogs/{id}", handler.Update)
	mux.HandleFunc("DELETE /blogs/{id}", handler.Delete)

	return &Router{
		mux,
	}
}
