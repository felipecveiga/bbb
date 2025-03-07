package service

import (
	"errors"
	"fmt"

	"github.com/felipecveiga/bbb/model"
	"github.com/felipecveiga/bbb/repository"
)

type Iservice interface {
	CreateVoto(voto *model.HistoricoVoto) error // Cria um novo voto
	GetAllVotos() (int64, error)                // Retorna todos votos
	GetVoto(participanteId int) (int64, error)  // Retorna voto pelo ID participante
	GetParticipanteFomDB(id int) (bool, error)  // Consulta participante no DB
	GetVotoHora() (map[string]int, error)       // Retorna votos por hora
}
type Service struct {
	Repository *repository.Repository
}

func NewService(r *repository.Repository) *Service {
	return &Service{
		Repository: r,
	}
}

func (s *Service) CreateVoto(voto *model.HistoricoVoto) error {

	participanteExist, err := s.Repository.GetParticipanteFomDB(voto.IdParticipante)
	if err != nil {
		return fmt.Errorf("erro: ao consultar participante: %w", err)
	}

	if !participanteExist {
		return errors.New("participante não existe")
	}

	isValido, err := s.Repository.StatusParticipante(voto.IdParticipante)
	if err != nil {
		return fmt.Errorf("erro: ao consultar status: %w", err)
	}

	if !isValido.Status {
		return errors.New("participante não está ativo")
	}

	err = s.Repository.CreateVotoFromDB(voto)
	if err != nil {
		return fmt.Errorf("erro ao registrar voto: %w", err)
	}

	return nil
}

func (s *Service) GetAllVotos() (int64, error) {

	totalVotos, err := s.Repository.GetAllVotosFromDB()
	if err != nil {
		return 0, fmt.Errorf("erro ao consultar todos os votos: %w", err)
	}

	return totalVotos, nil
}

func (s *Service) GetVoto(participanteId int) (int64, error) {

	participanteExist, err := s.Repository.GetParticipanteFomDB(participanteId)
	if err != nil {
		return 0, fmt.Errorf("erro: ao consultar participante: %w", err)
	}

	if !participanteExist {
		return 0, errors.New("participante não existe")
	}

	isValido, err := s.Repository.StatusParticipante(participanteId)
	if err != nil {
		return 0, fmt.Errorf("erro: ao consultar status: %w", err)
	}

	if !isValido.Status {
		return 0, errors.New("participante não está ativo")
	}

	votosParticipantes, err := s.Repository.GetVotosByIDFromDB(participanteId)
	if err != nil {
		return 0, fmt.Errorf("erro ao consultar total de votos: %w", err)
	}

	return votosParticipantes, nil
}

func (s *Service) GetVotoHora() (map[string]int, error) {

	votosHora, err := s.Repository.GetAllVotosHoraFromDB()
	if err != nil {
		return nil, fmt.Errorf("erro ao consultar votos por hora: %w", err)
	}
	return votosHora, nil
}
