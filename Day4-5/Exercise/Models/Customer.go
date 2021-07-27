package Models

import (
	"Exercise/Config"
	"fmt"
)

//GetAllCustomers Fetch all customers data
func GetAllCustomers(customer *[]Customer) (err error) {
	if err = Config.DB.Find(customer).Error; err != nil {
		return err
	}
	return nil
}

//AddCustomerDetails ... Insert New data
func AddCustomerDetails(customer *Customer) (err error) {
	if err = Config.DB.Create(customer).Error; err != nil {
		return err
	}
	return nil
}

//GetCustomerByID ... Fetch only one customer by ID
func GetCustomerByID(customer *Customer, id string) (err error) {
	if err = Config.DB.Where("id = ?", id).First(customer).Error; err != nil {
		return err
	}
	return nil
}

//GetOrderByCustomerID ... Fetch only one order by ID
func GetOrderByCustomerID(order *Order, CustomerId string) (err error) {
	if err = Config.DB.Where("customer_id = ?", CustomerId).First(order).Error; err != nil {
		return err
	}
	return nil
}

//UpdateCustomerDetails ... Update customer
func UpdateCustomerDetails(customer *Customer, id string) (err error) {
	fmt.Println(customer)
	Config.DB.Save(customer)
	return nil
}

//DeleteCustomerDetails ... Delete customer
func DeleteCustomerDetails(customer *Customer, id string) (err error) {
	Config.DB.Where("id = ?", id).Delete(customer)
	return nil
}
