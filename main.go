package main

import (
	"github.com/itinao/go-sample/app/controllers"
	"github.com/itinao/go-sample/config"
	"github.com/itinao/go-sample/utils"
)

func init() {
	utils.LoggingSettings(config.Config.LogFile)
}

func main() {
	controllers.IngestionData()
	controllers.StartWebServer()
}
