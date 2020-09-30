package rest

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"log"
	"net/http"
	"time"
)

func RunRESTServer() {
	r := chi.NewRouter()

	// A good base middleware stack
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	// Set a timeout value on the request context (ctx), that will signal
	// through ctx.Done() that the request has timed out and further
	// processing should be stopped.
	r.Use(middleware.Timeout(60 * time.Second))

	// RESTy routes for "fetcher" resource
	r.Route("/api/fetcher", func(r chi.Router) {
		r.Get("/", listUrls)
		r.Post("/", createUrl)

		// Subrouters:
		r.Route("/{urlID}", func(r chi.Router) {
			r.Get("/history", listHistory)
			r.Delete("/", deleteUrl)
		})
	})

	err := http.ListenAndServe(":8080", r)
	if err != nil {
		log.Panicln(err)
	}
}

func listUrls(writer http.ResponseWriter, request *http.Request) {
	// TODO: Impl
	writer.Write([]byte("listUrls"))
}

func createUrl(writer http.ResponseWriter, request *http.Request) {
	// TODO: Impl
	writer.Write([]byte("createUrl"))
}

func deleteUrl(writer http.ResponseWriter, request *http.Request) {
	// TODO: Impl
	writer.Write([]byte("deleteUrl"))
}

func listHistory(writer http.ResponseWriter, request *http.Request) {
	// TODO: Impl
	writer.Write([]byte("listHistory"))
}
