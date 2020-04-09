package handler

import (
	"fmt"
	"net/http"

	"github.com/dynastymasra/cookbook"

	"github.com/jinzhu/gorm"

	"github.com/sirupsen/logrus"
)

func Ping(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		log := logrus.WithField(cookbook.RequestID, r.Context().Value(cookbook.XRequestID))

		if err := db.DB().Ping(); err != nil {
			log.WithError(err).Errorln("Failed ping postgres")

			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, cookbook.ErrorResponse(err.Error(), r.Context().Value(cookbook.RequestID).(string)).Stringify())
			return
		}

		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, cookbook.SuccessResponse().Stringify())
	}
}
