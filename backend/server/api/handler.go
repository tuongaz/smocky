package api

import (
	"net/http"

	log "github.com/sirupsen/logrus"

	"github.com/smockyio/smocky/backend/persistent"
)

func GetMocksHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	db := persistent.GetDefault()
	mocks, err := db.GetConfigs(r.Context())
	if err != nil {
		log.WithError(err).Error("get configs")
		responseError(w, http.StatusInternalServerError, err.Error())
		return
	}

	response(w, http.StatusOK, mocks)
}
