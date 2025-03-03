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

func (r *Repository) GetAllVotosFromDB() (int64, error) {

	var votos int64

	err := r.DB.Model(&model.HistoricoVoto{}).
		Count(&votos).Error

	if err != nil {
		return 0, errors.New("erro na consulta do banco de dados")
	}

	return votos, nil
}

func (r *Repository) VotosParticipantes(participanteId int) (int64, error) {

	var votos int64

	err := r.DB.Model(&model.HistoricoVoto{}).
		Where("id_participante =?", participanteId).
		Count(&votos).Error

	if err != nil {
		return 0, errors.New("erro na consulta do banco de dados")
	}

	return votos, nil
}

func (r *Repository) GetAllVotosPorHoraFromDB() (map[string]int, error) {

	var resultados []struct {
		Hora  string
		Total int
	}

	err := r.DB.Model(&model.HistoricoVoto{}).
		Select("DATE_FORMAT(created_at, '%d-%m-%Y %H:00:00') as hora, COUNT(*) as total").
		Group("hora").
		Order("hora ASC").
		Scan(&resultados).Error

	if err != nil {
		return nil, errors.New("erro na consulta do banco de dados")
	}

	votosPorHora := make(map[string]int)
	for _, resultado := range resultados {
		votosPorHora[resultado.Hora] = resultado.Total
	}

	return votosPorHora, nil
}
