package Routes

import (
	"Exercise/Controllers"
	"github.com/gin-gonic/gin"
)

//SetupRouter ... Configure routes
func SetupRouter() *gin.Engine {
	r := gin.Default()
	router := r.Group("/")
	{
		router.GET("customers", Controllers.GetCustomers)
		router.POST("customers", Controllers.AddCustomerDetails)
		router.GET("customers/:id", Controllers.GetCustomerByID)

		router.GET("customers/orders/:id", Controllers.GetOrdersByCustomerID)
		router.PUT("customers/:id", Controllers.UpdateCustomerDetails)
		router.DELETE("customers/:id", Controllers.DeleteCustomerDetails)

		router.GET("products", Controllers.GetProducts)

		router.POST("products", Controllers.AddProductDetails)
		router.GET("products/:id", Controllers.GetProductByID)
		router.GET("products/orders/:id", Controllers.GetOrdersByProductID)
		// update it to patch
		router.PUT("products/:id", Controllers.UpdateProductDetails)
		router.DELETE("products/:id", Controllers.DeleteProductDetails)

		router.GET("orders", Controllers.GetOrders)

		//fetch the customer and product
		router.POST("orders", Controllers.PlaceOrderDetails)
		router.GET("orders/:id", Controllers.GetOrderByID)
		router.PUT("order/:id", Controllers.UpdateOrderDetails)
		router.DELETE("order/:id", Controllers.DeleteOrderDetails)
	}

	return r
}
