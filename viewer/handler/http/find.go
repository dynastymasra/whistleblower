package http

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/dynastymasra/cookbook"
	"github.com/dynastymasra/whistleblower/config"
	"github.com/dynastymasra/whistleblower/infrastructure/provider"
	"github.com/dynastymasra/whistleblower/viewer"
	"github.com/sirupsen/logrus"
)

// TODO: simple http handler add instrumentation
func StatisticCountHandler(service viewer.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		requestID := r.Context().Value(cookbook.RequestID).(string)

		v := mux.Vars(r)
		articleID := v["article_id"]

		log := logrus.WithFields(logrus.Fields{
			cookbook.RequestID: requestID,
			"article_id":       articleID,
		})

		query := provider.NewQuery(config.ViewerTableName)
		query = query.Filter("article_id", provider.Equal, articleID)

		statistic, err := service.Statistic(r.Context(), query)
		if err != nil {
			log.WithError(err).Errorln("Failed find all data")

			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprint(w, cookbook.ErrorResponse(err.Error(), requestID).Stringify())
			return
		}

		w.WriteHeader(http.StatusOK)

		res := &cookbook.JSON{
			"data": cookbook.JSON{
				"article_id": articleID,
				"type":       "statistics_article_view_count",
				"attributes": statistic,
			},
		}
		fmt.Fprint(w, cookbook.SuccessDataResponse(res, nil).Stringify())
	}
}
