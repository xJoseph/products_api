package routes

import (
	"clients_api/controllers"
	"fmt"
	"os"

	"github.com/labstack/echo/v4"
)

func ProductsRouter(e *echo.Echo) {
	var apiRoute = fmt.Sprintf("/api/%s", os.Getenv("API_VERSION"))

	e.GET(apiRoute+"/products", controllers.GetProducts)
	e.GET(apiRoute+"/products/:id", controllers.GetProductsById)
	e.POST(apiRoute+"/products", controllers.CreateProduct)
	e.PUT(apiRoute+"/products/:id", controllers.UpdateProduct)
	e.DELETE(apiRoute+"/products/:id", controllers.DeleteProduct)
}
