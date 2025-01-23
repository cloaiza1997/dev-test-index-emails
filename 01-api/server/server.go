package server

import (
	"fmt"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"github.com/cloaiza1997/dev-test-index-emails/config"
	"github.com/cloaiza1997/dev-test-index-emails/src/controllers"
)

func Start() {
	port := config.ApiConfig.Port

	if port == "" {
		fmt.Println("Port not defined")
		os.Exit(1)
	}

	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.AllowContentType("application/json"))

	routes(r)

	fmt.Printf("Starting server on port %s...\n", port)

	if err := http.ListenAndServe(":"+port, r); err != nil {
		fmt.Printf("error starting server: %v\n", err.Error())
	}
}

func routes(r *chi.Mux) {
	r.Mount("/", middleware.Profiler())

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Welcome to the email searcher API!"))
	})

	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("OK"))
	})

	r.Route("/v1/emails", func(r chi.Router) {
		r.Get("/", controllers.GetEmails)
	})
}
