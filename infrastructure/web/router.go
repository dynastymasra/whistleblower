package web

import (
	"fmt"
	"net/http"

	"github.com/dynastymasra/whistleblower/viewer"
	viewerHandler "github.com/dynastymasra/whistleblower/viewer/handler/http"

	"github.com/dynastymasra/whistleblower/article"
	articleHandler "github.com/dynastymasra/whistleblower/article/handler/http"

	"github.com/dynastymasra/cookbook"
	"github.com/dynastymasra/cookbook/negroni/middleware"
	"github.com/dynastymasra/whistleblower/infrastructure/web/handler"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"github.com/urfave/negroni"
)

const DefaultResponseNotFound = "the requested resource doesn't exists"

type RouterInstance struct {
	port           string
	name           string
	db             *gorm.DB
	articleService article.Service
	articleRepo    article.Repository
	viewerService  viewer.Service
	viewerRepo     viewer.Repository
}

// TODO: move all instance to new struct if params is too long
func NewRouter(port, name string, db *gorm.DB,
	articleService article.Service, articleRepo article.Repository,
	viewerService viewer.Service, viewerRepo viewer.Repository) *RouterInstance {
	return &RouterInstance{
		port:           port,
		name:           name,
		db:             db,
		articleService: articleService,
		articleRepo:    articleRepo,
		viewerService:  viewerService,
		viewerRepo:     viewerRepo,
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

	articleRouter := router.PathPrefix("/v1/").Subrouter().UseEncodedPath()
	commonHandlers.Use(middleware.LogrusLog(r.name))

	articleRouter.Handle("/articles", commonHandlers.With(
		negroni.WrapFunc(articleHandler.CreateArticle(r.articleService)),
	)).Methods(http.MethodPost)

	articleRouter.Handle("/articles", commonHandlers.With(
		negroni.WrapFunc(articleHandler.FindAllArticle(r.articleRepo)),
	)).Methods(http.MethodGet)

	statisticRouter := router.PathPrefix("/counter/v1/").Subrouter().UseEncodedPath()

	statisticRouter.Handle("/statistics", commonHandlers.With(
		negroni.WrapFunc(viewerHandler.CountViewer(r.viewerRepo)),
	)).Methods(http.MethodPost)

	statisticRouter.Handle("/statistics/article_id/{article_id}", commonHandlers.With(
		negroni.WrapFunc(viewerHandler.StatisticCount(r.viewerService)),
	)).Methods(http.MethodGet)

	return router
}
