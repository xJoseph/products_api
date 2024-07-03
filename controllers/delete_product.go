package controllers

import (
	"clients_api/db"
	"net/http"

	"github.com/labstack/echo/v4"
)

func DeleteProduct(c echo.Context) error {
	db, err := db.ConnectDB()

	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Não foi possível deletar o produto nesse momento")
	}
	result, _ := db.Exec("UPDATE products SET AVAILABLE = 0, LAST_UPDATE_DATE = NOW() WHERE ID = ?", c.Param("id"))
	rowsAffected, err := result.RowsAffected()

	if err != nil || rowsAffected == 0 {
		return c.JSON(http.StatusNotFound, "O produto não foi encontrado para apagar")
	}

	return c.JSON(http.StatusOK, "O produto foi deletado com sucesso")
}
