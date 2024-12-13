package function

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"
)

func GetEcho(e *echo.Echo) {
	fmt.Println("GetEcho")
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
}

func PostEcho(e *echo.Echo) {
	fmt.Println("PostEcho")
	e.POST("/", func(c echo.Context) error {
		text := c.FormValue("text")
		if text == "" {
			text = "No se recibió ningún texto"
		}
		return c.String(http.StatusOK, fmt.Sprintf("Recibido: %s", text))
	})
}
