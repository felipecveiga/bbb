package service

import (
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

	return s.Repository.CreateVotoFromDB(voto)

}

/*
func (s *Service) GetAllVotos() error {}

func (s *Service) GetVoto() error {}

func (s *Service) GetVotoHora() error {}
*/