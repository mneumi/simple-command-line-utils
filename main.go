package main

import (
	"log"

	"github.com/mneumi/simple-command-line-utils/cmd"
)

func main() {
	err := cmd.Execute()
	if err != nil {
		log.Fatalf("cmd.Execute err: %v", err)
	}
}
