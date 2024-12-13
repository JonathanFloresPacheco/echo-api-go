package main

import (
	"fmt"
	"myapp/src/function"

	"github.com/labstack/echo"
)

func main() {
	fmt.Println("Inicializando servidor......")
	e := echo.New()
	function.GetEcho(e)
	function.PostEcho(e)
	// Iniciar el servidor en un puerto
	e.Logger.Fatal(e.Start(":1323"))
}
