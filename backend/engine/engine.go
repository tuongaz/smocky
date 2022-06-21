package engine

import (
	"net/http"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/smockyio/smocky/backend/engine/matcher"
	"github.com/smockyio/smocky/backend/engine/mock"
	"github.com/smockyio/smocky/backend/persistent"
)

type Mock struct {
	mockID string
}

func New(configID string) (*Mock, error) {
	return &Mock{
		mockID: configID,
	}, nil
}

func (m *Mock) Engine(req *http.Request) *mock.Response {
	ctx := req.Context()
	db := persistent.GetDefault()
	cfg, err := db.GetConfig(ctx, m.mockID)
	if err != nil {
		log.WithError(err).Error("loading mock")
		return nil
	}

	sessionID, err := db.GetActiveSession(ctx, m.mockID)
	if err != nil {
		log.WithError(err).WithField("config_id", m.mockID).Error("get active session")
	}

	for _, route := range cfg.Routes {
		log.Debugf("Matching route: %v", route.Request)
		response, err := matcher.NewRouteMatcher(route, matcher.Context{
			HTTPRequest: req,
			SessionID:   sessionID,
		}).Match()
		if err != nil {
			log.WithError(err).Error("error while matching route")
			continue
		}

		if response == nil {
			log.Debug("no route matched")
			continue
		}

		if response.Delay > 0 {
			time.Sleep(time.Millisecond * time.Duration(response.Delay))
		}

		return response
	}

	return nil
}

func (m *Mock) Handler(w http.ResponseWriter, r *http.Request) {
	response := m.Engine(r)
	if response == nil {
		// TODO: no matched? What will be the response?
		w.WriteHeader(http.StatusNotFound)
		return
	}

	for k, v := range response.Headers {
		w.Header().Add(k, v)
	}

	if response.Status == 0 {
		response.Status = 200
	}
	w.WriteHeader(response.Status)
	_, _ = w.Write([]byte(response.Body))
}