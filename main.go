package main

import (
	"github.com/nomionz/currency-converter/cmd"
	"log"
)

func main() {
	if err := cmd.RunApi(); err != nil {
		log.Fatal(err)
	}
}
