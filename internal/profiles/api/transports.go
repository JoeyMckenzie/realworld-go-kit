package api

import (
	"github.com/go-chi/chi"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/go-kit/log"
	"github.com/joeymckenzie/realworld-go-kit/internal/profiles/core"
	"github.com/joeymckenzie/realworld-go-kit/internal/shared"
)

func MakeProfileRoutes(logger log.Logger, router *chi.Mux, service core.ProfilesService) *chi.Mux {
	getProfileHandler := httptransport.NewServer(
		makeGetProfileEndpoint(service),
		shared.DecodeNilPayload,
		shared.EncodeSuccessfulOkResponse,
		shared.HandlerOptions(logger)...,
	)

	followUserHandler := httptransport.NewServer(
		makeFollowUserEndpoint(service),
		shared.DecodeNilPayload,
		shared.EncodeSuccessfulOkResponse,
		shared.HandlerOptions(logger)...,
	)

	unfollowUserHandler := httptransport.NewServer(
		makeUnfollowUserEndpoint(service),
		shared.DecodeNilPayload,
		shared.EncodeSuccessfulOkResponse,
		shared.HandlerOptions(logger)...,
	)

	router.Route("/profiles/{username}", func(r chi.Router) {
		r.Use(shared.UsernameRequired)

		r.Group(func(r chi.Router) {
			r.Use(shared.AuthorizationOptional)
			r.Get("/", getProfileHandler.ServeHTTP)
		})

		r.Group(func(r chi.Router) {
			r.Use(shared.AuthorizationRequired)
			r.Post("/follow", followUserHandler.ServeHTTP)
			r.Delete("/follow", unfollowUserHandler.ServeHTTP)
		})
	})

	return router
}
