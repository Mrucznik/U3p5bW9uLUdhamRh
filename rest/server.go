package rest

import (
	"context"
	"github.com/Mrucznik/U3p5bW9uLUdhamRh/db"
	"github.com/Mrucznik/U3p5bW9uLUdhamRh/engine"
	"github.com/Mrucznik/U3p5bW9uLUdhamRh/engine/database"
	"github.com/Mrucznik/U3p5bW9uLUdhamRh/engine/in_memory"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"net/http"
	"strconv"
	"time"
)

var service engine.IURLsService

func RunRESTServer() {
	if viper.GetBool("USE_DATABASE") {
		logrus.Infoln("Connecting to MySQL databasee.")
		mysqlConnection := db.SetUpDatabase()
		defer mysqlConnection.Close()
		logrus.Infoln("Connectd.")

		service = database.NewURLsService(mysqlConnection)
	} else {
		service = in_memory.NewURLsService()
	}

	r := chi.NewRouter()

	// A good base middleware stack
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
			r.Use(UrlCtx)
			r.Get("/history", listHistory)
			r.Delete("/", deleteUrl)
		})
	})

	err := http.ListenAndServe(":8080", r)
	if err != nil {
		logrus.Panicln(err)
	}
}

func UrlCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		urlID, err := strconv.Atoi(chi.URLParam(r, "urlID"))
		if err != nil {
			http.Error(w, http.StatusText(400), 400)
			return
		}
		ctx := context.WithValue(r.Context(), "urlID", urlID)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
