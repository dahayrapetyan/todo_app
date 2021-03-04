package main

import (
	"log"
	"todo_app"
)

func main(){
	srv := new(todo_app.Server)
	if err := srv.Run("8080"); err != nil {
		log.Fatalf("erro server %s", err.Error())
	}
}
