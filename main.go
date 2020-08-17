package main

import (
	"fmt"
	"github.com/mariojuzar/kitara-store/api"
	"github.com/mariojuzar/kitara-store/api/configuration"
	"github.com/mariojuzar/kitara-store/api/service"
	"os"
)

func main() {
	cmdString := command()
	fmt.Println("command " + cmdString)

	if cmdString == "migrate" {
		var dbSvc = service.NewDatabaseService()
		_ = dbSvc.Initialize()
		_ = dbSvc.Migrate()
	} else {
		engine := api.Run()
		_ = engine.Run(":" + configuration.Port)
	}
}

func command() string {
	args := os.Args[1:]

	if len(args) > 0 {
		return args[0]
	}
	return ""
}
