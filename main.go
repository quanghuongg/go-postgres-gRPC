package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"go-postgres/handler"
)


const (
	host     = "localhost"
	port     = "5432"
	user     = "postgres"
	password = "huongnq"
	dbname   = "postgres"
)
// our initial migration function


func main()  {
	router := gin.Default()
	v1 := router.Group("/user")
	{
		v1.POST("/login", handler.LoginUser)
		v1.POST("/register", handler.RegisterUser)
		v1.GET("/info", handler.InfoUser)
	}
	router.Run(":3000")


}