package main

import (
	"log"

	server "github.com/sam-val/text-to-img/api"
)

func main() {
	err := server.Run("localhost:8080")
	if err != nil {
		log.Fatal("Can not run server ", err)
	}
}
