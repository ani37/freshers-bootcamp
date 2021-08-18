package Controllers

import (
	"fmt"

	"Exercise/Models"

	"github.com/gin-gonic/gin"
	"net/http"
)

//GetCustomers ... Get all customers
func GetCustomers(c *gin.Context) {
	var customer []Models.Customer
	err := Models.GetAllCustomers(&customer)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, customer)
	}
}

//AddCustomerDetails ... Create Customer
func AddCustomerDetails(c *gin.Context) {
	var customer Models.Customer
	c.BindJSON(&customer)
	err := Models.AddCustomerDetails(&customer)
	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, customer)
	}
}

//GetCustomerByID ... Get the customer by id
func GetCustomerByID(c *gin.Context) {
	id := c.Params.ByName("id")
	var customer Models.Customer
	err := Models.GetCustomerByID(&customer, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, customer)
	}
}

//GetOrdersByCustomerID ... Get the order by customer id
func GetOrdersByCustomerID(c *gin.Context) {
	id := c.Params.ByName("id")

	var order []Models.Order
	err := Models.GetOrdersByCustomerID(&order, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, order)
	}
}

//UpdateCustomerDetails ... Update the customer information
func UpdateCustomerDetails(c *gin.Context) {
	var customer Models.Customer
	id := c.Params.ByName("id")
	err := Models.GetCustomerByID(&customer, id)
	if err != nil {
		c.JSON(http.StatusNotFound, customer)
	}
	c.BindJSON(&customer)
	err = Models.UpdateCustomerDetails(&customer, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, customer)
	}
}

//DeleteCustomerDetails ... Delete the customer
func DeleteCustomerDetails(c *gin.Context) {
	var customer Models.Customer
	id := c.Params.ByName("id")
	err := Models.DeleteCustomerDetails(&customer, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, gin.H{"id" + id: "is deleted"})
	}
}
