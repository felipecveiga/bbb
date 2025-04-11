package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/felipecveiga/bbb/model"
	"github.com/stretchr/testify/assert"
	gormMysql "gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type testVotoRepository struct {
	VotoRepository Repository
}

var (
	now = time.Now()

	vote = &model.HistoricoVoto{
		ID:             1,
		IdParticipante: 2,
		Ip:             "",
		Created_at:     now,
	}

	participante = &model.Participante{
		ID:         2,
		Nome:       "",
		Residencia: "",
		Ocupacao:   "",
		Status:     true,
	}
)

func getMockRepository() (*testVotoRepository, sqlmock.Sqlmock, *sql.DB) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		panic(fmt.Sprintf("an error '%s' was not expected when opening a stub database connection", err))
	}

	dbGorm, err := gorm.Open(gormMysql.New(gormMysql.Config{
		Conn:                      db,
		SkipInitializeWithVersion: true,
	}), &gorm.Config{Logger: logger.Default.LogMode(logger.Silent)})
	if err != nil {
		panic(fmt.Sprintf("an error '%s' was not expected when opening a stub database connection", err))
	}

	voto := NewRepository(dbGorm)
	return &testVotoRepository{
		VotoRepository: voto,
	}, mock, db
}

func TestCreateVoteFromDB_WhenReturnSucess(t *testing.T) {
	repository, mock, db := getMockRepository()
	defer db.Close()

	expectedSQL := "INSERT INTO `historico_votos` (`id_participante`,`ip`,`created_at`,`id`) VALUES (?,?,?,?)"

	mock.ExpectBegin()
	mock.ExpectExec(expectedSQL).
		WithArgs(vote.IdParticipante, vote.Ip, sqlmock.AnyArg(), vote.ID).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	err := repository.VotoRepository.CreateVoteFromDB(vote)

	assert.NoError(t, err)
}

func TestCreateVoteFromDB_WhenReturnError(t *testing.T) {
	repository, mock, db := getMockRepository()
	defer db.Close()

	expectedSQL := "INSERT INTO `historico_votos` (`id_participante`,`ip`,`created_at`,`id`) VALUES (?,?,?,?)"

	mock.ExpectExec(expectedSQL).
		WithArgs(vote.IdParticipante, vote.Ip, sqlmock.AnyArg(), vote.ID).
		WillReturnError(errors.New("some error"))

	err := repository.VotoRepository.CreateVoteFromDB(vote)

	assert.Error(t, err)
}

func TestGetParticipantStatusFromDB_WhenReturnSucess(t *testing.T) {
	repository, mock, db := getMockRepository()
	defer db.Close()

	rows := []string{"status"}
	expectedSQL := "SELECT `status` FROM `participantes` WHERE id = ? ORDER BY `participantes`.`id` LIMIT ?"

	mock.ExpectQuery(expectedSQL).
		WithArgs(participante.ID, 1).
		WillReturnRows(
			sqlmock.NewRows(rows).
				AddRow(participante.Status))

	response, err := repository.VotoRepository.GetParticipantStatusFromDB(participante.ID)

	assert.NoError(t, err)
	assert.Equal(t, participante.Status, response.Status)
}

func TestGetParticipantStatusFromDB_WhenReturnError(t *testing.T) {
	repository, mock, db := getMockRepository()
	defer db.Close()

	expectedSQL := "SELECT `status` FROM `participantes` WHERE id = ? ORDER BY `participantes`.`id` LIMIT ?"

	mock.ExpectQuery(expectedSQL).
		WithArgs(participante.ID, 1).
		WillReturnError(errors.New("some error"))

	_, err := repository.VotoRepository.GetParticipantStatusFromDB(participante.ID)

	assert.Error(t, err)
}

func TestGetAllVotesFromDB_WhenReturnSucess(t *testing.T) {
	repository, mock, db := getMockRepository()
	defer db.Close()

	rows := sqlmock.NewRows([]string{"count"}).AddRow(4)
	expectedSQL := "SELECT count(*) FROM `historico_votos`"

	mock.ExpectQuery(expectedSQL).
		WillReturnRows(rows)

	result, err := repository.VotoRepository.GetAllVotesFromDB()

	assert.NoError(t, err)
	assert.Equal(t, int64(4), result)
}

func TestGetAllVotesFromDB_WhenReturnError(t *testing.T) {
	repository, mock, db := getMockRepository()
	defer db.Close()

	expectedSQL := "SELECT count(*) FROM `historico_votos`"

	mock.ExpectQuery(expectedSQL).
		WillReturnError(errors.New("some error"))

	_, err := repository.VotoRepository.GetAllVotesFromDB()

	assert.Error(t, err)
}

func TestGetVotesByIdFromDB_WhenReturnSucess(t *testing.T) {
	repository, mock, db := getMockRepository()
	defer db.Close()

	rows := sqlmock.NewRows([]string{"count"}).AddRow(4)
	expectedSQL := "SELECT count(*) FROM `historico_votos` WHERE id_participante =?"

	mock.ExpectQuery(expectedSQL).
		WithArgs(participante.ID).
		WillReturnRows(rows)

	result, err := repository.VotoRepository.GetVotesByIdFromDB(participante.ID)

	assert.NoError(t, err)
	assert.Equal(t, int64(4), result)
}

func TestGetVotesByIdFromDB_WhenReturnError(t *testing.T) {
	repository, mock, db := getMockRepository()
	defer db.Close()

	expectedSQL := "SELECT count(*) FROM `historico_votos` WHERE id_participante =?"

	mock.ExpectQuery(expectedSQL).
		WithArgs(participante.ID).
		WillReturnError(errors.New("some error"))

	_, err := repository.VotoRepository.GetVotesByIdFromDB(participante.ID)

	assert.Error(t, err)
}

func TestGetAllVotesHourFromDB_WhenReturnSucess(t *testing.T) {
	repository, mock, db := getMockRepository()
	defer db.Close()

	rows := sqlmock.NewRows([]string{"hora", "total"}).
		AddRow("20-10-2025", 4).
		AddRow("21-11-2025", 11)
	expectedSQL := "SELECT DATE_FORMAT(created_at, '%d-%m-%Y %H:00:00') as hora, COUNT(*) as total FROM `historico_votos` GROUP BY `hora` ORDER BY hora ASC"

	mock.ExpectQuery(expectedSQL).
		WillReturnRows(rows)

	result, err := repository.VotoRepository.GetAllVotesHourFromDB()

	assert.NoError(t, err)
	assert.Equal(t, map[string]int{
		"20-10-2025": 4,
		"21-11-2025": 11,
	}, result)
}

func TestGetAllVotesHourFromDB_WhenReturnError(t *testing.T) {
	repository, mock, db := getMockRepository()
	defer db.Close()

	expectedSQL := "SELECT DATE_FORMAT(created_at, '%d-%m-%Y %H:00:00') as hora, COUNT(*) as total FROM `historico_votos` GROUP BY `hora` ORDER BY hora ASC"

	mock.ExpectQuery(expectedSQL).
		WillReturnError(errors.New("some error"))

	_, err := repository.VotoRepository.GetAllVotesHourFromDB()

	assert.Error(t, err)
}
