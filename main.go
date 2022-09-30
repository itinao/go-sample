package main

import (
	"log"

	"github.com/itinao/go-sample/config"
	"github.com/itinao/go-sample/utils"
)

func main() {
	utils.LoggingSettings(config.Config.LogFile)
	log.Printf("go-sample!")
}
