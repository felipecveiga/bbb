package handler

import (
	"net/http"
	"strconv"

	"github.com/felipecveiga/bbb/model"
	"github.com/felipecveiga/bbb/service"
	"github.com/labstack/echo"
)
//go:generate mockgen -source=./voto_bbb.go -destination=./voto_bbb_mock.go -package=handler
type Handler interface {
	Vote(c echo.Context) error
	GetTotalVotes(c echo.Context) error
	GetParticipantVotes(c echo.Context) error
	GetVotesHour(c echo.Context) error
}

type handler struct {
	Service service.Service
}

func NewHandler(s service.Service) Handler {
	return &handler{
		Service: s,
	}
}

func (h *handler) Vote(c echo.Context) error {

	vote := new(model.HistoricoVoto)

	if err := c.Bind(vote); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	if vote.IdParticipante == 0 {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "id participante inválido"})
	}

	err := h.Service.CreateVote(vote)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, vote)
}

func (h *handler) GetTotalVotes(c echo.Context) error {

	votes, err := h.Service.GetAllVotes()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, votes)
}

func (h *handler) GetParticipantVotes(c echo.Context) error {

	participantId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "id inválido"})
	}

	if participantId <= 0 {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "id participante inválido"})
	}

	votesParticipants, err := h.Service.GetVote(participantId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, votesParticipants)
}

func (h *handler) GetVotesHour(c echo.Context) error {

	votes, err := h.Service.GetVoteHour()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, votes)
}
