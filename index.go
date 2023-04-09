package main

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type Product struct {
	ID          int    `json:"id"`
	NamaProduct string `json:"nama_product"`
	Price       int    `json:"price"`
	Stock       int    `json:"stock"`
}

// data product sementara
var products = []Product{
	{1, "Product1", 12000, 40},
	{2, "Product2", 24500, 45},
	{3, "Product3", 68000, 23},
}

// controllers
func getAllProducts(c echo.Context) error {
	return c.JSON(http.StatusOK, products)
}

func getProduct(c echo.Context) error {
	id := c.Param("id")

	for _, p := range products {
		if strconv.Itoa(p.ID) == id {
			return c.JSON(http.StatusOK, p)
		}
	}

	return c.NoContent(http.StatusNotFound)
}

func createProduct(c echo.Context) error {
	var p Product

	if err := c.Bind(&p); err != nil {
		return err
	}

	products = append(products, p)
	return c.JSON(http.StatusCreated, p)
}

func updateProduct(c echo.Context) error {
	id := c.Param("id")

	for i, p := range products {
		if strconv.Itoa(p.ID) == id {
			var newP Product

			if err := c.Bind(&newP); err != nil {
				return err
			}

			products[i] = newP
			return c.JSON(http.StatusOK, newP)
		}
	}

	return c.NoContent(http.StatusNotFound)
}

func deleteProduct(c echo.Context) error {
	id := c.Param("id")

	for i, p := range products {
		if strconv.Itoa(p.ID) == id {
			products = append(products[:i], products[i+1:]...)
			return c.NoContent(http.StatusOK)
		}
	}

	return c.NoContent(http.StatusNotFound)
}

func main() {
	e := echo.New()

	e.GET("/products", getAllProducts)
	e.GET("/products/:id", getProduct)
	e.POST("/products", createProduct)
	e.PUT("/products/:id", updateProduct)
	e.DELETE("/products/:id", deleteProduct)

	e.Start(":8080")
}
