package main

import (
	"apii_v1/config"
	"log"
	"os"
)

func main() {
	err := config.Connect()
	if err != nil {
		log.Fatal(err)
		os.Exit(2)
	}
}
