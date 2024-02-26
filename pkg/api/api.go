package api

import (
	"fmt"
	"net/http"
	"server/pkg/resources"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func api() {
	mainRouter := chi.NewRouter()
	mainRouter.Use(middleware.Logger)
	mainRouter.Use(middleware.Recoverer)

	resources.SetupRouter(mainRouter)
	mainRouter.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("welcome"))
	})

	fmt.Println("Server running!")

	http.ListenAndServe(":3000", mainRouter)
}

func StartServer() {
	api()
}
