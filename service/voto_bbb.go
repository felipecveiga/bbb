package service

import (
	"errors"
	"fmt"

	"github.com/felipecveiga/bbb/model"
	"github.com/felipecveiga/bbb/repository"
)

type Service struct {
	Repository *repository.Repository
}

func NewService(r *repository.Repository) *Service {
	return &Service{
		Repository: r,
	}
}

func (s *Service) CreateVoto(voto *model.HistoricoVoto) error {

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

/*
func (s *Service) GetAllVotos() error {}

func (s *Service) GetVoto() error {}

func (s *Service) GetVotoHora() error {}
*/
