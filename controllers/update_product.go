package controllers

import (
	"clients_api/db"
	"clients_api/models"
	"net/http"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
)

func UpdateProduct(c echo.Context) error {
	var product models.UpdateProduct
	if err := c.Bind(&product); err != nil {
		return c.JSON(http.StatusBadRequest, "Informações incorretas para atualizar o produto")
	}

	validate := validator.New()
	if err := validate.Struct(product); err != nil {
		return c.String(http.StatusBadRequest, "Informações incorretas para atualizar o produto")
	}

	db, err := db.ConnectDB()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Não foi possível atualizar o produto nesse momento")
	}

	result, _ := db.Exec("UPDATE products SET QTY = ?, PRICE = ?, STOCK = ?, LAST_UPDATE_DATE = NOW() WHERE ID = ? AND AVAILABLE = 1", product.QTY, product.PRICE, product.STOCK, c.Param("id"))
	rowsAffected, _ := result.RowsAffected()

	if rowsAffected == 0 {
		return c.JSON(http.StatusNotFound, "O produto não foi encontrado para atualizar")
	}
	return c.JSON(http.StatusOK, "O produto foi atualizado com sucesso")
}
