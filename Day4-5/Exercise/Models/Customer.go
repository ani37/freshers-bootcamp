package Models

import (
	"Exercise/Config"
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

//GetOrdersByCustomerID ... Fetch only one order by ID
func GetOrdersByCustomerID(order *[]Order, CustomerId string) (err error) {
	if err = Config.DB.Where("customer_id = ?", CustomerId).Find(order).Error; err != nil {
		return err
	}
	return nil
}

//UpdateCustomerDetails ... Update customer
func UpdateCustomerDetails(customer *Customer, id string) (err error) {
	//fmt.Println(customer)
	//handle errors
	Config.DB.Save(customer)
	return nil
}

//DeleteCustomerDetails ... Delete customer
func DeleteCustomerDetails(customer *Customer, id string) (err error) {
	// handle errors
	Config.DB.Where("id = ?", id).Delete(customer)
	return nil
}
