package main

import (
	"fmt"
	"go-postgres/dao/daoimpl"
	"go-postgres/driver"
	"go-postgres/model"
	 "log"
)

const (
	host ="localhost"
	port ="5432"
	user ="postgres"
	password ="huongnq"
	dbname ="postgres"
)

func main()  {
	db:=driver.Connect(host,port,user,password,dbname)
	err:= db.SQL.Ping()
	if err!=nil{
		panic(err)
	}
	fmt.Println("Connected")
	userDao:= daoimpl.NewUseDao(db.SQL)

	person:=model.User{
		Name: "golang",
	}

	err = userDao.Insert(person)
	if err != nil {
		log.Fatal(err)
	}
	users, err :=userDao.Select()
	if err!=nil{
		log.Fatal(err)
	}
	fmt.Println(users)
	
}