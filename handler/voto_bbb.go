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

func (h *Handler) Votar(c echo.Context) error {

	voto := new(model.HistoricoVoto)

	if err := c.Bind(voto); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	if voto.IdParticipante == 0 {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "id participante inválido"})
	}

	err := h.Service.CreateVoto(voto)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, voto)
}

func (h *Handler) ObterTotalVotos(c echo.Context) error {

	votos, err := h.Service.GetAllVotos()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, votos)
}

func (h *Handler) ObterVotosPorParticipante(c echo.Context) error {

	participanteId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "id inválido"})
	}

	if participanteId <= 0 {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "id participante inválido"})
	}

	votosParticipantes, err := h.Service.GetVoto(participanteId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, votosParticipantes)
}

func (h *Handler) ObterVotosPorHora(c echo.Context) error {

	votos, err := h.Service.GetVotoHora()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, votos)
}
