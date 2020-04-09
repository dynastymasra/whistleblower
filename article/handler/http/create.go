package http

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	uuid "github.com/satori/go.uuid"

	"github.com/dynastymasra/whistleblower/domain"

	"github.com/dynastymasra/cookbook"
	"github.com/sirupsen/logrus"

	"github.com/dynastymasra/whistleblower/article"
)

// TODO: simple http handler add instrumentation
func CreateArticle(service article.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		requestID := r.Context().Value(cookbook.RequestID).(string)

		log := logrus.WithFields(logrus.Fields{
			cookbook.RequestID: requestID,
		})

		var article domain.Article

		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.WithError(err).Errorln("Unable to read request body")

			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprint(w, cookbook.ErrorResponse(err.Error(), requestID).Stringify())
			return
		}

		if err := json.Unmarshal(body, &article); err != nil {
			log.WithError(err).WithField("body", string(body)).Errorln("Unable to parse request body")

			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprint(w, cookbook.ErrorResponse(err.Error(), requestID).Stringify())
			return
		}

		if _, err := uuid.FromString(article.ID); err != nil {
			article.ID = uuid.NewV4().String()
		}

		res, err := service.Create(r.Context(), article)
		if err != nil {
			log.WithError(err).WithField("article", article).Errorln("Failed create new data")

			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprint(w, cookbook.ErrorResponse(err.Error(), requestID).Stringify())
			return
		}

		w.WriteHeader(http.StatusCreated)
		fmt.Fprint(w, cookbook.SuccessDataResponse(&cookbook.JSON{"article": res}, nil).Stringify())
	}
}
