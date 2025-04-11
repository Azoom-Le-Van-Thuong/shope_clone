package main

import (
	"log"
	"shope_clone/internal/app"
)

func main() {
	if err := app.Bootstrap(); err != nil {
		log.Fatal(err)
	}
}
