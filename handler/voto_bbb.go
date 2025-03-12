package handler

import (
	"net/http"
	"strconv"

	"github.com/felipecveiga/bbb/model"
	"github.com/felipecveiga/bbb/service"
	"github.com/labstack/echo"
)

type IHandler interface {
	Votar(c echo.Context) error
	ObterTotalVotos(c echo.Context) error
	ObterVotosPorParticipante(c echo.Context) error
	ObterVotosPorHora(c echo.Context) error
}

type Handler struct {
	Service *service.Service
}

func NewHandler(s *service.Service) *Handler {
	return &Handler{
		Service: s,
	}
}

func (h *Handler) Vote(c echo.Context) error {

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

func (h *Handler) GetTotalVotes(c echo.Context) error {

	votes, err := h.Service.GetAllVotes()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, votes)
}

func (h *Handler) GetParticipantVotes(c echo.Context) error {

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

func (h *Handler) GetVotesHour(c echo.Context) error {

	votes, err := h.Service.GetVoteHour()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, votes)
}
