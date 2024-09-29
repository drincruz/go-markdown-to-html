package main

import (
	"log"
)

func check(err error, message string) {
	if err != nil {
		log.Fatalf("[Error] %s: %s", message, err)
		panic(err)
	}
}
