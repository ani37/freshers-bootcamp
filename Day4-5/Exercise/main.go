package main

//seperate the imports
import (
	"fmt"

	"Exercise/Config"
	"Exercise/Models"
	"Exercise/Routes"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var err error

func main() {
	Config.DB, err = gorm.Open(mysql.Open(Config.DbURL(Config.BuildDBConfig())), &gorm.Config{})
	if err != nil {
		// give more better status
		fmt.Println("Status:", err)
	}

	Config.DB.AutoMigrate(&Models.Product{}, &Models.Customer{}, &Models.Order{})
	r := Routes.SetupRouter()
	//running
	r.Run()
}
