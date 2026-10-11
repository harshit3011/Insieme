package main

import (
	"context"
	"fmt"
	"log"

	"github.com/harshit3011/Insieme/database"
)

func main() {
	db, err:= database.Connect()
	if err!=nil{
		log.Fatal("Couldn't connect to the db",err)
	}
	var result string
	defer db.Close()
	err=db.QueryRow(context.Background(),"SELECT current_database()::text").Scan(&result)
	if err!=nil{
		log.Fatal("Couldn't perform the query",err)
	}
	fmt.Println("Connected to the Database!!")
}