package service

import (
	"errors"
	"testing"
	"time"

	"github.com/felipecveiga/bbb/model"
	"github.com/felipecveiga/bbb/repository"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

var (
	now = time.Now()

	vote = &model.HistoricoVoto{
		ID:             1,
		IdParticipante: 2,
		Ip:             "",
		Created_at:     now,
	}
)

func TestCreateVote_WhenReturnSucess(t *testing.T) {
	ctrl := gomock.NewController(t)
	mock := NewMockService(ctrl)

	mock.EXPECT().
		CreateVote(vote).
		Return(nil)

	err := mock.CreateVote(vote)

	assert.NoError(t, err)
}

func TestCreateVote_WhenConsultingParticipant_ReturnError(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockRepo := repository.NewMockRepository(ctrl)
	service := NewService(mockRepo)

	mockRepo.EXPECT().
		GetParticipantFomDB(vote.IdParticipante).
		Return(false, errors.New("erro: ao consultar participante"))

	err := service.CreateVote(vote)

	assert.Error(t, err)
}

func TestCreateVote_WhenConsultingStatusParticipant_ReturnError(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockRepo := repository.NewMockRepository(ctrl)
	service := NewService(mockRepo)

	mockRepo.EXPECT().
		GetParticipantFomDB(vote.IdParticipante).
		Return(true, nil)
	mockRepo.EXPECT().
		GetParticipantStatusFromDB(vote.IdParticipante).
		Return(nil, errors.New("erro: ao consultar status"))

	err := service.CreateVote(vote)

	assert.Error(t, err)
}

func TestCreateVote_WhenReturnError(t *testing.T) {
	ctrl := gomock.NewController(t)
	mock := NewMockService(ctrl)

	mock.EXPECT().
		CreateVote(vote).
		Return(errors.New("erro ao registrar voto"))

	err := mock.CreateVote(vote)

	assert.Error(t, err)
}

func TestGetAllVotes_WhenReturnSucess(t *testing.T) {
	ctrl := gomock.NewController(t)
	mock := NewMockService(ctrl)

	mock.EXPECT().
		GetAllVotes().
		Return(int64(10), nil)

	result, err := mock.GetAllVotes()

	assert.NoError(t, err)
	assert.Equal(t, int64(10), result)
}

func TestGetAllVotes_WhenReturnError(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockRepo := repository.NewMockRepository(ctrl)
	service := NewService(mockRepo)

	mockRepo.EXPECT().
		GetAllVotesFromDB().
		Return(int64(0), errors.New("Erro ao consultar todos os votos"))

	_, err := service.GetAllVotes()

	assert.Error(t, err)
}

func TestGetVote_WhenReturnSucess(t *testing.T) {
	ctrl := gomock.NewController(t)
	mock := NewMockService(ctrl)

	mock.EXPECT().
		GetVote(1).
		Return(int64(10), nil)

	result, err := mock.GetVote(1)

	assert.NoError(t, err)
	assert.Equal(t, int64(10), result)
}

func TestGetVote_WhenConsultingParticipant_ReturnError(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockRepo := repository.NewMockRepository(ctrl)
	service := NewService(mockRepo)

	mockRepo.EXPECT().
		GetParticipantFomDB(vote.IdParticipante).
		Return(false, errors.New("erro: ao consultar participante"))

	_, err := service.GetVote(vote.IdParticipante)

	assert.Error(t, err)
}

func TestGetVote_WhenConsultingStatusParticipant_ReturnError(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockRepo := repository.NewMockRepository(ctrl)
	service := NewService(mockRepo)

	mockRepo.EXPECT().
		GetParticipantFomDB(vote.IdParticipante).
		Return(true, nil)
	mockRepo.EXPECT().
		GetParticipantStatusFromDB(vote.IdParticipante).
		Return(nil, errors.New("erro: ao consultar status"))

	_, err := service.GetVote(vote.IdParticipante)

	assert.Error(t, err)
}

func TestGetVoteHour_WhenReturnSucess(t *testing.T) {
	ctrl := gomock.NewController(t)
	mock := NewMockService(ctrl)

	mock.EXPECT().
		GetVoteHour().
		Return(map[string]int{"08:00": 10, "10:00": 20}, nil)

	result, err := mock.GetVoteHour()

	assert.NoError(t, err)
	assert.Equal(t, map[string]int{"08:00": 10, "10:00": 20}, result)
}

func TestGetVoteHour_WhenVotesHour_ReturnError(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockRepo := repository.NewMockRepository(ctrl)
	service := NewService(mockRepo)

	mockRepo.EXPECT().
		GetAllVotesHourFromDB().
		Return(nil, errors.New("Erro ao consultar votos por hora"))

	_, err := service.GetVoteHour()

	assert.Error(t, err)
}
