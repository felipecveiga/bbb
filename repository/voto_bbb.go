package repository

import (
	"fmt"

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

	err := r.DB.Create(voto).Error

	if err != nil {
		return fmt.Errorf("erro ao salvar voto no banco de dados: %w", err)
	}

	return nil
}

func (r *Repository) StatusParticipante(idParticipante int) (*model.Participante, error) {

	var participante model.Participante

	err := r.DB.Model(&model.Participante{}).
		Select("status").
		Where("id = ?", idParticipante).
		First(&participante).Error

	if err != nil {
		return nil, fmt.Errorf("erro na consulta do banco de dados %w", err)
	}

	return &participante, nil
}

func (r *Repository) GetAllVotosFromDB() (int64, error) {

	var votos int64

	err := r.DB.Model(&model.HistoricoVoto{}).
		Count(&votos).Error

	if err != nil {
		return 0, fmt.Errorf("erro na consulta do banco de dados %w", err)
	}

	return votos, nil
}

func (r *Repository) GetVotosByIDFromDB(participanteId int) (int64, error) {

	var votos int64

	err := r.DB.Model(&model.HistoricoVoto{}).
		Where("id_participante =?", participanteId).
		Count(&votos).Error

	if err != nil {
		return 0, fmt.Errorf("erro na consulta do banco de dados %w", err)
	}

	return votos, nil
}

func (r *Repository) GetAllVotosHoraFromDB() (map[string]int, error) {

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
		return nil, fmt.Errorf("erro na consulta do banco de dados %w", err)
	}

	votosPorHora := make(map[string]int)
	for _, resultado := range resultados {
		votosPorHora[resultado.Hora] = resultado.Total
	}

	return votosPorHora, nil
}

func (r *Repository) GetParticipanteFomDB(participanteId int) (bool, error) {

	var result int64

	err := r.DB.Model(&model.Participante{}).
		Where("id =?", participanteId).
		Count(&result).Error

	if err != nil {
		return false, fmt.Errorf("erro na consulta do banco de dados %w", err)
	}

	return result > 0, nil
}
