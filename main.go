package main

import (
	"log"
	"os"

	"github.com/nrechn/musubi/api"
	"github.com/nrechn/musubi/utils"
)

func main() {
	if err := utils.CheckConfig(); err != nil {
		log.Println(err)
		os.Exit(1)
	}
	if err := utils.LoadConfig(); err != nil {
		log.Println(err)
		os.Exit(1)
	}

	api.Execute()
}
