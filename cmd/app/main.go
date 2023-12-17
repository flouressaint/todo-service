package main

import (
	"log"

	"github.com/flouressaint/todo-service/internal/app"
)

const configPath = "config/config.yaml"

func main() {
	a, err := app.New()
	if err != nil {
		log.Fatal(err)
	}

	err = a.Run()
	if err != nil {
		log.Fatal(err)
	}
}
