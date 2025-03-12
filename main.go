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
	e.POST("/votar", HandlerHistoricoVoto.Vote)
	e.GET("/votos/:id", HandlerHistoricoVoto.GetParticipantVotes)
	e.GET("/votos", HandlerHistoricoVoto.GetTotalVotes)
	e.GET("/votos/hora", HandlerHistoricoVoto.GetVotesHour)

	e.Logger.Fatal(e.Start(":8080"))
}
