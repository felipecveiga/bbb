package repository

import (
	"errors"

	"github.com/felipecveiga/bbb/model"
	"gorm.io/gorm"
)

type Repository struct {
	DB gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		DB: *db,
	}
}

func (r *Repository) CreateVotoFromDB(voto *model.HistoricoVoto) error {
	return r.DB.Create(voto).Error
}

func (r *Repository) StatusParticipante(idParticipante int) (*model.Participante, error) {

	var participante model.Participante

	err := r.DB.Model(&model.Participante{}).
		Select("status").
		Where("id = ?", idParticipante).
		First(&participante).Error

	if err != nil {
		return nil, errors.New("erro na consulta do banco de dados")
	}

	return &participante, nil
}
