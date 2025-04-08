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
		WillReturnResult(sqlmock.NewErrorResult(errors.New("some error")))

	err := repository.VotoRepository.CreateVoteFromDB(vote)

	assert.Error(t, err)
}
