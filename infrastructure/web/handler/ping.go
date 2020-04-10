package handler

import (
	"fmt"
	"net/http"

	"github.com/dynastymasra/whistleblower/infrastructure/provider"

	"github.com/dynastymasra/cookbook"

	"github.com/jinzhu/gorm"

	"github.com/sirupsen/logrus"
)

func Ping(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		requestID := r.Context().Value(cookbook.RequestID).(string)
		log := logrus.WithField(cookbook.RequestID, requestID)

		if err := provider.Ping(db); err != nil {
			log.WithError(err).Errorln("Failed ping postgres")

			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprint(w, cookbook.ErrorResponse(err.Error(), requestID).Stringify())
			return
		}

		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, cookbook.SuccessResponse().Stringify())
	}
}
