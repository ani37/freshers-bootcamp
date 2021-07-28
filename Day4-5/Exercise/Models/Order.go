package Models

import (
	"fmt"
	"strconv"
	"sync"

	"Exercise/Config"
)

var mutex sync.Mutex

//GetAllOrders Fetch all orders data
func GetAllOrders(order *[]Order) (err error) {
	if err = Config.DB.Find(order).Error; err != nil {
		return err
	}
	return nil
}

//PlaceOrderDetails ... Insert New data
func PlaceOrderDetails(order *Order) (err error) {

	order.Status = "Failed!"
	defer mutex.Unlock()
	err = GetProductByID(&order.Product, strconv.Itoa(int(order.ProductID)))
	if err != nil {
		return err
	}
	err = GetCustomerByID(&order.Customer, strconv.Itoa(int(order.CustomerID)))
	if err != nil {
		return err
	}
	mutex.Lock()
	_ = GetProductByID(&order.Product, strconv.Itoa(int(order.ProductID)))
	if order.Product.Quantity < order.Quantity {
		order.Status = "Processed!"
		if err = Config.DB.Create(order).Error; err != nil {
			return err
		}
		return nil
	}
	order.Product.Quantity = order.Product.Quantity - order.Quantity
	UpdateProductDetails(&order.Product, strconv.Itoa(int(order.ProductID)))
	order.Status = "Successful!"
	if err = Config.DB.Create(order).Error; err != nil {
		return err
	}
	return nil
}

//GetOrderByID ... Fetch only one order by ID
func GetOrderByID(order *Order, id string) (err error) {
	if err = Config.DB.Where("id = ?", id).First(order).Error; err != nil {
		return err
	}
	return nil
}

//UpdateOrderDetails ... Update order
func UpdateOrderDetails(order *Order, id string) (err error) {
	fmt.Println(order)
	Config.DB.Save(order)
	return nil
}

//DeleteOrderDetails ... Delete order
func DeleteOrderDetails(order *Order, id string) (err error) {
	Config.DB.Where("id = ?", id).Delete(order)
	return nil
}
