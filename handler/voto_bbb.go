package handler

import (
	"net/http"

	"github.com/felipecveiga/bbb/model"
	"github.com/felipecveiga/bbb/service"
	"github.com/labstack/echo"
)

type Handler struct {
	Service *service.Service
}

func NewHandler(s *service.Service) *Handler {
	return &Handler{
		Service: s,
	}
}

func (h *Handler) RegistrarVoto(c echo.Context) error {

	voto := new(model.HistoricoVoto)
	if err := c.Bind(voto); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	err := h.Service.CreateVoto(voto)
	if err != nil {
		return c.JSON(http.StatusConflict, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, voto)
}

func (h *Handler) ObterTotalVotos(c echo.Context) error {

	return c.JSON(200, "")
}

func (h *Handler) ObterVotosPorParticipante(c echo.Context) error {
	return c.JSON(200, "")
}

func (h *Handler) ObterVotosPorHora(c echo.Context) error {
	return c.JSON(200, "")
}
