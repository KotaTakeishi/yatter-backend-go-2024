package main

import (
	"log"

	"yatter-backend-go/app/server"
)

func main() {
	log.Fatalf("%+v", server.Run())
}
