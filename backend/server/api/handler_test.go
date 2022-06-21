package api_test

import (
	"context"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/smockyio/smocky/backend/server/api"
	"github.com/smockyio/smocky/engine/mock"
	"github.com/smockyio/smocky/engine/persistent"
	"github.com/smockyio/smocky/engine/persistent/memory"
)

func Test_GetMocksHandler(t *testing.T) {
	db := memory.New()
	persistent.New(db)

	_ = db.SetConfig(context.Background(), &mock.Mock{
		ID: "123",
	})

	_ = db.SetConfig(context.Background(), &mock.Mock{
		ID: "456",
	})

	req := httptest.NewRequest(http.MethodGet, "/mocks", nil)
	w := httptest.NewRecorder()
	api.GetMocksHandler(w, req)

	res := w.Result()
	defer func() {
		_ = res.Body.Close()
	}()

	data, err := ioutil.ReadAll(res.Body)
	assert.NoError(t, err)
	assert.JSONEq(t, `[{"id":"123"},{"id":"456"}]`, string(data))
	assert.Equal(t, http.StatusOK, res.StatusCode)
}