package controllers

import (
	"clients_api/db"
	"clients_api/models"
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetProducts(c echo.Context) error {
	db, err := db.ConnectDB()

	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Não foi possível resgatar os produtos nesse momento")
	}
	rows, err := db.Query("SELECT ID, NAME, QTY, PRICE, STOCK FROM products WHERE AVAILABLE = 1 ORDER BY ID ASC")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Há um problema com a base de dados. Talvez a tabela não exista")
	}

	var products []models.Product
	for rows.Next() {
		var product models.Product
		err = rows.Scan(&product.ID, &product.NAME, &product.QTY, &product.PRICE, &product.STOCK)

		if err != nil {
			return c.JSON(http.StatusInternalServerError, "Erro ao tratar dados")
		}
		products = append(products, product)
	}
	return c.JSON(http.StatusOK, products)
}

func GetProductsById(c echo.Context) error {
	db, err := db.ConnectDB()

	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Não foi possível resgatar os produtos nesse momento")
	}
	rows, err := db.Query("SELECT ID, NAME, QTY, PRICE, STOCK FROM products WHERE AVAILABLE = 1 AND ID = ?", c.Param("id"))

	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Há um problema com a base de dados. Talvez a tabela não exista")
	}

	var product models.Product
	found := false
	for rows.Next() {
		err = rows.Scan(&product.ID, &product.NAME, &product.QTY, &product.PRICE, &product.STOCK)

		if err != nil {
			return c.JSON(http.StatusInternalServerError, "Erro ao tratar dados")
		}
		found = true
	}
	if !found {
		return c.JSON(http.StatusNotFound, "Produto inexistente")
	}

	return c.JSON(http.StatusOK, product)
}
