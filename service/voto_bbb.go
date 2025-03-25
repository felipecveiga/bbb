package service

import (
	"errors"
	"fmt"

	"github.com/felipecveiga/bbb/model"
	"github.com/felipecveiga/bbb/repository"
)

//go:generate mockgen -source=./voto_bbb.go -destination=./voto_bbb_mock.go -package=service
type Service interface {
	CreateVote(voto *model.HistoricoVoto) error // Cria um novo voto
	GetAllVotes() (int64, error)                // Retorna todos votos
	GetVote(participanteId int) (int64, error)  // Retorna voto pelo ID participante
	GetVoteHour() (map[string]int, error)       // Retorna votos por hora
}

type service struct {
	Repository repository.Repository
}

func NewService(r repository.Repository) Service {
	return &service{
		Repository: r,
	}
}

func (s *service) CreateVote(vote *model.HistoricoVoto) error {

	participantExist, err := s.Repository.GetParticipantFomDB(vote.IdParticipante)
	if err != nil {
		return fmt.Errorf("erro: ao consultar participante: %w", err)
	}

	if !participantExist {
		return errors.New("participante não existe")
	}

	isValid, err := s.Repository.GetParticipantStatusFromDB(vote.IdParticipante)
	if err != nil {
		return fmt.Errorf("erro: ao consultar status: %w", err)
	}

	if !isValid.Status {
		return errors.New("participante não está ativo")
	}

	err = s.Repository.CreateVoteFromDB(vote)
	if err != nil {
		return fmt.Errorf("erro ao registrar voto: %w", err)
	}

	return nil
}

func (s *service) GetAllVotes() (int64, error) {

	totalVotes, err := s.Repository.GetAllVotesFromDB()
	if err != nil {
		return 0, fmt.Errorf("erro ao consultar todos os votos: %w", err)
	}

	return totalVotes, nil
}

func (s *service) GetVote(participantId int) (int64, error) {

	participantExist, err := s.Repository.GetParticipantFomDB(participantId)
	if err != nil {
		return 0, fmt.Errorf("erro: ao consultar participante: %w", err)
	}

	if !participantExist {
		return 0, errors.New("participante não existe")
	}

	isValid, err := s.Repository.GetParticipantStatusFromDB(participantId)
	if err != nil {
		return 0, fmt.Errorf("erro: ao consultar status: %w", err)
	}

	if !isValid.Status {
		return 0, errors.New("participante não está ativo")
	}

	votesParticipants, err := s.Repository.GetVotesByIdFromDB(participantId)
	if err != nil {
		return 0, fmt.Errorf("erro ao consultar total de votos: %w", err)
	}

	return votesParticipants, nil
}

func (s *service) GetVoteHour() (map[string]int, error) {

	votesHour, err := s.Repository.GetAllVotesHourFromDB()
	if err != nil {
		return nil, fmt.Errorf("erro ao consultar votos por hora: %w", err)
	}
	return votesHour, nil
}
