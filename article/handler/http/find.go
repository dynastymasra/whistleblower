package http

import (
	"fmt"
	"net/http"

	"github.com/dynastymasra/cookbook"
	"github.com/sirupsen/logrus"

	"github.com/dynastymasra/whistleblower/article"
)

// TODO: simple http handler add instrumentation
func FindAllArticle(repository article.Repository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		requestID := r.Context().Value(cookbook.RequestID).(string)

		log := logrus.WithFields(logrus.Fields{
			cookbook.RequestID: requestID,
		})

		articles, err := repository.FindAll(r.Context(), nil)
		if err != nil {
			log.WithError(err).Errorln("Failed find all data")

			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprint(w, cookbook.ErrorResponse(err.Error(), requestID).Stringify())
			return
		}

		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, cookbook.SuccessDataResponse(&cookbook.JSON{"articles": articles}, nil).Stringify())
	}
}
