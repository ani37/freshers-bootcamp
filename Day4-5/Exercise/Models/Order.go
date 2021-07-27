package Models

import (
	"Exercise/Config"
	"fmt"
	"gorm.io/gorm"
)

//GetAllOrders Fetch all orders data
func GetAllOrders(order *[]Order) (err error) {
	if err = Config.DB.Find(order).Error; err != nil {
		return err
	}
	return nil
}

//AddOrderDetails ... Insert New data
func AddOrderDetails(order *Order) (err error) {

	if err := Config.DB.Table("product").Where("id = ? AND quantity >= ?",
		order.ProductID, order.Quantity).
		UpdateColumn("quantity", gorm.Expr("quantity - ?", order.Quantity)).Error; err != nil {
		return err
	}
	if err := Config.DB.Table("order").Create(order).Error; err != nil {
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
