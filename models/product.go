package models

type Product struct {
	ID    int     `json:"id"`
	NAME  string  `json:"name"`
	QTY   int     `json:"qty"`
	PRICE float64 `json:"price"`
	STOCK int     `json:"stock"`
}

type UpdateProduct struct {
	QTY   int     `json:"qty" validate:"required"`
	PRICE float64 `json:"price" validate:"required"`
	STOCK int     `json:"stock"  validate:"required"`
}

type CreateProduct struct {
	NAME           string  `json:"name" validate:"required"`
	QTY            int     `json:"qty" validate:"required"`
	PRICE          float64 `json:"price" validate:"required"`
	STOCK          int     `json:"stock"  validate:"required"`
	TYPE_ID        int     `json:"type_id" validate:"required"`
	BRAND_ID       int     `json:"brand_id" validate:"required"`
	MEASUREMENT_ID int     `json:"measurement_id" validate:"required"`
}
