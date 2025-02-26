package main

import (
	"github.com/felipecveiga/bbb/config"
	"github.com/felipecveiga/bbb/handler"
	"github.com/felipecveiga/bbb/repository"
	"github.com/felipecveiga/bbb/service"
	"github.com/labstack/echo"
)

func main() {
	db := config.Carregar()
	repoHistoricoVoto := repository.NewRepository(db)
	ServiceHistoricoVoto := service.NewService(repoHistoricoVoto)
	HandlerHistoricoVoto := handler.NewHandler(ServiceHistoricoVoto)

	e := echo.New()
	e.POST("/votar", HandlerHistoricoVoto.RegistrarVoto)
	e.GET("/votos/:id", HandlerHistoricoVoto.ObterVotosPorParticipante)

	e.Logger.Fatal(e.Start(":8080"))
}
