// Main package of OctoSummon Prometheus alert router
package main

import (
	"github.com/newrushbolt/OctoSummon/config"
	"github.com/newrushbolt/OctoSummon/logger"
	"github.com/newrushbolt/OctoSummon/server"
)

func main() {
	logger.Init("")
	mainConfig, err := config.GetConfig()
	if err != nil {
		logger.Logger.Fatalf("Config load failed:\n%v\n", err)
	}
	server.Start(mainConfig)
}
