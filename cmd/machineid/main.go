package main

import (
	"log"

	"github.com/denisbrodbeck/machineid"
)

func main() {
	id, err := machineid.ProtectedID("server")

	if err != nil {
		log.Fatal(err)
	}

	log.Println(id)
}
