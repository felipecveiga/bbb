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

	votes, err := mock.GetAllVotes()

	assert.NoError(t, err)
	assert.Equal(t, int64(10), votes)
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
