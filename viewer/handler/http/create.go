package http

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/dynastymasra/cookbook"
	"github.com/dynastymasra/whistleblower/domain"
	"github.com/dynastymasra/whistleblower/viewer"
	uuid "github.com/satori/go.uuid"
	"github.com/sirupsen/logrus"
)

// TODO: simple http handler add instrumentation
func CountViewer(repo viewer.Repository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		requestID := r.Context().Value(cookbook.RequestID).(string)

		log := logrus.WithFields(logrus.Fields{
			cookbook.RequestID: requestID,
		})

		var model domain.Viewer

		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.WithError(err).Errorln("Unable to read request body")

			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprint(w, cookbook.ErrorResponse(err.Error(), requestID).Stringify())
			return
		}

		if err := json.Unmarshal(body, &model); err != nil {
			log.WithError(err).WithField("body", string(body)).Errorln("Unable to parse request body")

			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprint(w, cookbook.ErrorResponse(err.Error(), requestID).Stringify())
			return
		}

		if _, err := uuid.FromString(model.ID); err != nil {
			model.ID = uuid.NewV4().String()
		}

		go func(v domain.Viewer) {
			if err := repo.Create(r.Context(), v); err != nil {
				log.WithError(err).WithField("viewer", model).Errorln("Failed create new data")
			}
		}(model)

		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, cookbook.SuccessResponse().Stringify())
	}
}
