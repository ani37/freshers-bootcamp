package main

import (
	"Exercise/Config"
	"Exercise/Models"
	"Exercise/Routes"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var err error

func main() {
	Config.DB, err = gorm.Open(mysql.Open(Config.DbURL(Config.BuildDBConfig())), &gorm.Config{})
	if err != nil {
		fmt.Println("Status:", err)
	}

	Config.DB.AutoMigrate(&Models.Product{}, &Models.Customer{}, &Models.Order{})
	r := Routes.SetupRouter()
	//running
	r.Run()
}
