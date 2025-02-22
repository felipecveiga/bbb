package main

import (
	"github.com/felipecveiga/bbb/config"
	"github.com/labstack/echo"
)

func main() {
	//db := config.Carregar()
	//repoHistoricoVoto := NewHistoricoVoto(db)

	e := echo.New()

	e.Logger.Fatal(e.Start(":8080"))
}
