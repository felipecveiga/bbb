package handler

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/felipecveiga/bbb/model"
	"github.com/felipecveiga/bbb/service"
	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
	gomock "go.uber.org/mock/gomock"
)

type MockObjectHandler struct {
	Service  *service.MockService
	Recorder *httptest.ResponseRecorder
	Ctx      echo.Context
	Handler  *handler
}

func SetupMockHandler(url, method string, payload io.Reader, svc *service.MockService) *MockObjectHandler {
	e := echo.New()
	recorder := httptest.NewRecorder()
	request := httptest.NewRequest(method, url, payload)
	request.Header.Set("Content-Type", "application/json")
	context := e.NewContext(request, recorder)
	handler := &handler{Service: svc}
	return &MockObjectHandler{
		Recorder: recorder,
		Service:  svc,
		Ctx:      context,
		Handler:  handler,
	}
}

func TestVote_WhenReturnSucess(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := service.NewMockService(ctrl)

	url := "/votar"
	now := time.Now()

	payload, _ := json.Marshal(model.HistoricoVoto{ID: 1, IdParticipante: 2, Ip: "", Created_at: now})

	mocks := SetupMockHandler(url, http.MethodPost, bytes.NewBuffer(payload), mockService)

	mockService.EXPECT().
		CreateVote(gomock.Any()).
		Return(nil)

	err := mocks.Handler.Vote(mocks.Ctx)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, mocks.Recorder.Code)
}

func TestVote_WhenReturnErrorParsing(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := service.NewMockService(ctrl)

	url := "/votar"

	payload := bytes.NewBuffer([]byte(`{invalid json}`))

	mocks := SetupMockHandler(url, http.MethodPost, payload, mockService)

	err := mocks.Handler.Vote(mocks.Ctx)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusBadRequest, mocks.Recorder.Code)
}

func TestVote_WhenReturnErrorIdParticipante(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	now := time.Now()

	mockService := service.NewMockService(ctrl)

	url := "/votar"

	payload, _ := json.Marshal(model.HistoricoVoto{ID: 1, IdParticipante: 0, Ip: "", Created_at: now})

	mocks := SetupMockHandler(url, http.MethodPost, bytes.NewBuffer(payload), mockService)

	err := mocks.Handler.Vote(mocks.Ctx)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusBadRequest, mocks.Recorder.Code)
	assert.Contains(t, mocks.Recorder.Body.String(), "id participante inválido")
}

func TestVote_WhenServiceReturnError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := service.NewMockService(ctrl)

	url := "/votar"
	now := time.Now()

	payload, _ := json.Marshal(model.HistoricoVoto{ID: 1, IdParticipante: 2, Ip: "", Created_at: now})

	mocks := SetupMockHandler(url, http.MethodPost, bytes.NewBuffer(payload), mockService)

	mockService.EXPECT().
		CreateVote(gomock.Any()).
		Return(errors.New("some error"))

	err := mocks.Handler.Vote(mocks.Ctx)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusInternalServerError, mocks.Recorder.Code)
}

func TestGetTotalVotes_WhenReturnSucess(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := service.NewMockService(ctrl)

	url := "/votos"

	mocks := SetupMockHandler(url, http.MethodGet, nil, mockService)

	mockService.EXPECT().
		GetAllVotes().
		Return(int64(10), nil)

	err := mocks.Handler.GetTotalVotes(mocks.Ctx)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, mocks.Recorder.Code)
}

func TestGetTotalVotes_WhenReturnError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := service.NewMockService(ctrl)

	url := "/votos"

	mocks := SetupMockHandler(url, http.MethodGet, nil, mockService)

	mockService.EXPECT().
		GetAllVotes().
		Return(int64(0), errors.New("some error"))

	err := mocks.Handler.GetTotalVotes(mocks.Ctx)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusInternalServerError, mocks.Recorder.Code)
}
