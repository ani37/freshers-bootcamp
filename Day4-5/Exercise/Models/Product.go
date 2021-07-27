package Models

import (
	"Exercise/Config"
	"fmt"
)

//GetAllProducts Fetch all products data
func GetAllProducts(product *[]Product) (err error) {
	if err = Config.DB.Find(product).Error; err != nil {
		return err
	}
	return nil
}

//AddProductDetails ... Insert New data
func AddProductDetails(product *Product) (err error) {
	if err = Config.DB.Create(product).Error; err != nil {
		return err
	}
	return nil
}

//GetProductByID ... Fetch only one product by ID
func GetProductByID(product *Product, id string) (err error) {
	if err = Config.DB.Where("id = ?", id).First(product).Error; err != nil {
		return err
	}
	return nil
}

//GetOrderByProductID ... Fetch only one order by ID
func GetOrderByProductID(order *Order, ProductId string) (err error) {
	if err = Config.DB.Where("product_id = ?", ProductId).First(order).Error; err != nil {
		return err
	}
	return nil
}

//UpdateProductDetails ... Update product
func UpdateProductDetails(product *Product, id string) (err error) {
	fmt.Println(product)
	Config.DB.Save(product)
	return nil
}

//DeleteProductDetails ... Delete product
func DeleteProductDetails(product *Product, id string) (err error) {
	Config.DB.Where("id = ?", id).Delete(product)
	return nil
}
