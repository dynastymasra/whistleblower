package web

import (
	"fmt"
	"net/http"

	"github.com/dynastymasra/cookbook"
	"github.com/dynastymasra/cookbook/negroni/middleware"
	"github.com/dynastymasra/whistleblower/infrastructure/web/handler"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"github.com/urfave/negroni"
)

const DefaultResponseNotFound = "the requested resource doesn't exists"

type RouterInstance struct {
	port string
	name string
	db   *gorm.DB
}

func NewRouter(port, name string, db *gorm.DB) *RouterInstance {
	return &RouterInstance{
		port: port,
		name: name,
		db:   db,
	}
}

func (r *RouterInstance) Router() *mux.Router {
	router := mux.NewRouter().StrictSlash(true).UseEncodedPath()

	router.NotFoundHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w, cookbook.FailResponse(&cookbook.JSON{
			"endpoint": DefaultResponseNotFound,
		}, "").Stringify())
	})

	router.MethodNotAllowedHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprint(w, cookbook.FailResponse(&cookbook.JSON{
			"method": DefaultResponseNotFound,
		}, "").Stringify())
	})

	commonHandlers := negroni.New(
		middleware.RequestID(),
	)

	// Probes
	router.Handle("/ping", commonHandlers.With(
		negroni.WrapFunc(handler.Ping(r.db)),
	)).Methods(http.MethodGet, http.MethodHead)

	router.Handle("/ping", commonHandlers.With(
		negroni.WrapFunc(handler.Ping(r.db)),
	)).Methods(http.MethodGet, http.MethodHead)

	subRouter := router.PathPrefix("/v1/").Subrouter().UseEncodedPath()
	commonHandlers.Use(middleware.LogrusLog(r.name))

	subRouter.Handle("/ping", commonHandlers.With(
		negroni.WrapFunc(handler.Ping(r.db)),
	)).Methods(http.MethodGet, http.MethodHead)

	return router
}
