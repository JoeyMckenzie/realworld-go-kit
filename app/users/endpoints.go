package users

import (
	"net/http"

	"github.com/go-chi/chi"
)

func MakeEndpoints(router *chi.Mux, service UsersService) *chi.Mux {
	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("welcome"))
	})

	return router
}
