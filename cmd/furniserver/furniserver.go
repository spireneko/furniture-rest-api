package main

import (
	"github.com/spireneko/furniture-rest-api/internal/app"

	"log"
)

func main() {
	if err := app.Run(":8080"); err != nil {
		log.Println(err)
	}
}
