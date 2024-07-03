package controllers

import (
	"clients_api/db"
	"clients_api/models"
	"net/http"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
)

func CreateProduct(c echo.Context) error {
	var product models.CreateProduct
	if err := c.Bind(&product); err != nil {
		return c.JSON(http.StatusBadRequest, "Informações incorretas para criar um novo produto")
	}

	validate := validator.New()
	if err := validate.Struct(product); err != nil {
		return c.String(http.StatusBadRequest, "Informações incorretas para criar um novo produto")
	}

	db, err := db.ConnectDB()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Não foi possível criar um novo produto nesse momento")
	}

	result, _ := db.Exec(
		"INSERT INTO products(`NAME`,`QTY`,`PRICE`,`TYPE_ID`,`BRAND_ID`,`MEASUREMENT_ID`,`AVAILABLE`, `REGISTRY_DATE`, `LAST_UPDATE_DATE`, `STOCK`) VALUES(?,?,?,?,?,?,1, NOW(), NOW(), ?);",
		product.NAME, product.QTY, product.PRICE, product.TYPE_ID, product.BRAND_ID, product.MEASUREMENT_ID, product.STOCK)
	rowsAffected, _ := result.RowsAffected()

	if rowsAffected == 0 {
		return c.JSON(http.StatusBadRequest, "O produto não foi criado")
	}
	return c.JSON(http.StatusOK, "O produto foi criado com sucesso")
}
