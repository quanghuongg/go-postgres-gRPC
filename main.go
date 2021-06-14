package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	"go-postgres/handler"
	"go-postgres/model"
)


const (
	host     = "localhost"
	port     = "5432"
	user     = "postgres"
	password = "huongnq"
	dbname   = "postgres"
)
// our initial migration function
func initialMigration() {
	db, err :=  gorm.Open( "postgres", "host=localhost port=5432 user=postgres dbname=postgres sslmode=disable password=huongnq")
	if err != nil {
		fmt.Println(err.Error())
		panic("failed to connect database")
	}
	defer db.Close()

	// Migrate the schema
	db.AutoMigrate(&model.User{})
}

func main()  {
	initialMigration()
	router := gin.Default()
	v1 := router.Group("/user")
	{
		v1.POST("/login", handler.LoginUser)
		v1.POST("/register", handler.RegisterUser)
		v1.GET("/info", handler.InfoUser)
	}
	router.Run(":3000")


}