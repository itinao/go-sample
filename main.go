package main

import (
	"log"

	"github.com/itinao/go-sample/app/controllers"
	"github.com/itinao/go-sample/config"
	"github.com/itinao/go-sample/utils"
)

func init() {
	utils.LoggingSettings(config.Config.LogFile)
}

func main() {
	log.Printf("go-sample!")
	//controllers.IngestionData()
	controllers.StartWebServer()
}
