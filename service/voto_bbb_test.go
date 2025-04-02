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

func TestCreateVote_WhenReturnSucess(t *testing.T) {
	ctrl := gomock.NewController(t)

	mock := NewMockService(ctrl)

	now := time.Now()

	vote := &model.HistoricoVoto{
		ID:             1,
		IdParticipante: 2,
		Ip:             "",
		Created_at:     now,
	}

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

	now := time.Now()

	vote := &model.HistoricoVoto{
		ID:             1,
		IdParticipante: 2,
		Ip:             "",
		Created_at:     now,
	}

	mockRepo.EXPECT().
		GetParticipantFomDB(vote.IdParticipante).
		Return(false, errors.New("erro: ao consultar participante"))

	err := service.CreateVote(vote)

	assert.Error(t, err)
}
