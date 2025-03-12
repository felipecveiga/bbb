package repository

import (
	"fmt"

	"github.com/felipecveiga/bbb/model"
	"gorm.io/gorm"
)

type IRepository interface {
	CreateVoteFromDB(voto *model.HistoricoVoto) error                   // Registra o voto no BD
	StatusParticipante(idParticipante int) (*model.Participante, error) // Verifica o status do participante
	GetAllVotosFromDB() (int64, error)                                  // Retorna todos os votos
	GetVotosByIDFromDB(participanteId int) (int64, error)               // Retorna os votos pelo ID
	GetAllVotosHoraFromDB() (map[string]int, error)                     // Retorna os votos por hora
	GetParticipanteFomDB(participanteId int) (bool, error)              // Verifica se o participante existe.
}

type Repository struct {
	DB gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		DB: *db,
	}
}

func (r *Repository) CreateVoteFromDB(vote *model.HistoricoVoto) error {

	err := r.DB.Create(vote).Error

	if err != nil {
		return fmt.Errorf("erro ao salvar voto no banco de dados: %w", err)
	}

	return nil
}

func (r *Repository) GetParticipantStatusFromDB(idParticipant int) (*model.Participante, error) {

	var participant model.Participante

	err := r.DB.Model(&model.Participante{}).
		Select("status").
		Where("id = ?", idParticipant).
		First(&participant).Error

	if err != nil {
		return nil, fmt.Errorf("erro na consulta do banco de dados %w", err)
	}

	return &participant, nil
}

func (r *Repository) GetAllVotesFromDB() (int64, error) {

	var votes int64

	err := r.DB.Model(&model.HistoricoVoto{}).
		Count(&votes).Error

	if err != nil {
		return 0, fmt.Errorf("erro na consulta do banco de dados %w", err)
	}

	return votes, nil
}

func (r *Repository) GetVotesByIdFromDB(participantId int) (int64, error) {

	var votes int64

	err := r.DB.Model(&model.HistoricoVoto{}).
		Where("id_participante =?", participantId).
		Count(&votes).Error

	if err != nil {
		return 0, fmt.Errorf("erro na consulta do banco de dados %w", err)
	}

	return votes, nil
}

func (r *Repository) GetAllVotesHourFromDB() (map[string]int, error) {

	var totalVotesHour []struct {
		Hora  string
		Total int
	}

	err := r.DB.Model(&model.HistoricoVoto{}).
		Select("DATE_FORMAT(created_at, '%d-%m-%Y %H:00:00') as hora, COUNT(*) as total").
		Group("hora").
		Order("hora ASC").
		Scan(&totalVotesHour).Error

	if err != nil {
		return nil, fmt.Errorf("erro na consulta do banco de dados %w", err)
	}

	votesHour := make(map[string]int)
	for _, result := range totalVotesHour {
		votesHour[result.Hora] = result.Total
	}

	return votesHour, nil
}

func (r *Repository) GetParticipantFomDB(participantId int) (bool, error) {

	var result int64

	err := r.DB.Model(&model.Participante{}).
		Where("id =?", participantId).
		Count(&result).Error

	if err != nil {
		return false, fmt.Errorf("erro na consulta do banco de dados %w", err)
	}

	return result > 0, nil
}
