package main

import (
	"fmt"
	"go-rest-api/db"
	"go-rest-api/model"
)

func main() {
	dbConn := db.NewDb()
	defer fmt.Println("Sucessfully migrated")
	defer db.CloseDB(dbConn)
	dbConn.AutoMigrate(&model.User{},&model.Task{})
}