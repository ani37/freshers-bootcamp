**Task**

Build service to support a hypothetical retailer whose requirements are as follows:
The retailer has some products which he wants to sell. The retailer should be able to add a product, update the available quantity and price of the product as and when required.
Customers should be able to view all the available products along with their prices using the REST API.
A customer should be able to place an order. And later should be able to view his order history.
The retailer should be able to view a detailed history of all the transactions of his business.
Since our retailer has many customers, He wants you to design the solution in such a way that once an order of the customer is executed, there should be a gap or cool-down period of 5 mins for the next order to be executed. (Cool down is at customer level)

**Tech Stack:**

golang - gin (microframework)

Mysql

Gorm as orm library


**Api contract**

```GET    /customers                --> Exercise/Controllers.GetCustomers (3 handlers) 

 POST   /customers                --> Exercise/Controllers.AddCustomerDetails (3 handlers)

 GET    /customers/:id            --> Exercise/Controllers.GetCustomerByID (3 handlers)

 GET    /customers/orders/:id     --> Exercise/Controllers.GetOrdersByCustomerID (3 handlers)

 PUT    /customers/:id            --> Exercise/Controllers.UpdateCustomerDetails (3 handlers)

 DELETE /customers/:id            --> Exercise/Controllers.DeleteCustomerDetails (3 handlers)

 GET    /products                 --> Exercise/Controllers.GetProducts (3 handlers)

 POST   /products                 --> Exercise/Controllers.AddProductDetails (3 handlers)

 GET    /products/:id             --> Exercise/Controllers.GetProductByID (3 handlers)

 GET    /products/orders/:id      --> Exercise/Controllers.GetOrdersByProductID (3 handlers)

 PUT    /products/:id             --> Exercise/Controllers.UpdateProductDetails (3 handlers)

 DELETE /products/:id             --> Exercise/Controllers.DeleteProductDetails (3 handlers)

 GET    /orders                   --> Exercise/Controllers.GetOrders (3 handlers)

 POST   /orders                   --> Exercise/Controllers.PlaceOrderDetails (3 handlers)

 GET    /orders/:id               --> Exercise/Controllers.GetOrderByID (3 handlers)

 PUT    /order/:id                --> Exercise/Controllers.UpdateOrderDetails (3 handlers)

 DELETE /order/:id                --> Exercise/Controllers.DeleteOrderDetails (3 handlers)
```