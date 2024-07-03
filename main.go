package main

import (
	"clients_api/routes"
	"fmt"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		panic("Erro ao carregar arquivo .env")
	}

	server := echo.New()

	server.Use(middleware.Logger())
	server.Use(middleware.Recover())
	server.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.PUT, echo.POST, echo.DELETE},
	}))

	routes.ProductsRouter(server)

	server.GET(fmt.Sprintf("/api/%s", os.Getenv("API_VERSION")), func(c echo.Context) error {
		return c.JSON(http.StatusOK, "Api Online")
	})

	server.Logger.Fatal(server.Start(":" + os.Getenv("API_PORT")))
}
