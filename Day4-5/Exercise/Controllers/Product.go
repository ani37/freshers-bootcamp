package Controllers

import (
	"Exercise/Models"
	"fmt"

	"github.com/gin-gonic/gin"
	"net/http"
)

//GetProducts ... Get all products
func GetProducts(c *gin.Context) {
	var product []Models.Product
	err := Models.GetAllProducts(&product)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, product)
	}
}

//AddProductDetails ... Create Product
func AddProductDetails(c *gin.Context) {
	var product Models.Product
	c.BindJSON(&product)
	err := Models.AddProductDetails(&product)
	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, product)
	}
}

//GetProductByID ... Get the product by id
func GetProductByID(c *gin.Context) {
	id := c.Params.ByName("id")
	var product Models.Product
	err := Models.GetProductByID(&product, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, product)
	}
}

//GetOrderByProductID ... Get the order by product id
func GetOrderByProductID(c *gin.Context) {
	id := c.Params.ByName("id")

	var order Models.Order
	err := Models.GetOrderByProductID(&order, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, order)
	}
}

//UpdateProductDetails ... Update the product information
func UpdateProductDetails(c *gin.Context) {
	var product Models.Product
	id := c.Params.ByName("id")
	err := Models.GetProductByID(&product, id)
	if err != nil {
		c.JSON(http.StatusNotFound, product)
	}
	c.BindJSON(&product)
	err = Models.UpdateProductDetails(&product, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, product)
	}
}

//DeleteProductDetails ... Delete the product
func DeleteProductDetails(c *gin.Context) {
	var product Models.Product
	id := c.Params.ByName("id")
	err := Models.DeleteProductDetails(&product, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, gin.H{"id" + id: "is deleted"})
	}
}
