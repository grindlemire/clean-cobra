package main

import (
	"github.com/grindlemire/clean-cobra/cmd"
	"github.com/grindlemire/log"
)

func main() {
	log.Init(log.Default)

	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
}
