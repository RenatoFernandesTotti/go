package resources

import (
	"server/pkg/resources/v1/login"

	"github.com/go-chi/chi/v5"
)

func SetupRouter(mainRouter *chi.Mux) {
	setupV1(mainRouter)
}

func setupV1(mainRouter *chi.Mux) {
	v1Router := chi.NewRouter()

	v1Router.Mount("/login", login.LoginRouter())

	mainRouter.Mount("/v1", v1Router)
}
