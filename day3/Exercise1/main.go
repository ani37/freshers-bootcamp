package main
import (

	"fmt"
	"github.com/jinzhu/gorm"
	"simpleAPI/Config"
	"simpleAPI/Models"
	"simpleAPI/Routers"
)
var err error
func main() {
	Config.DB, err = gorm.Open("mysql", Config.DbURL(Config.BuildDBConfig()))
	if err != nil {
		fmt.Println("Status:", err)
	}
	defer Config.DB.Close()
	Config.DB.AutoMigrate(&Models.User{})
	r := Routes.SetupRouter()
	//running
	r.Run()
}