package app

import (
	"ahti/app/common/logger"
	"ahti/app/config"
	"ahti/app/internal/web"
	"github.com/common-nighthawk/go-figure"
)

// on start print a ascii art
func init() {
	heading := figure.NewFigure("AHTI.GO", "", true)
	heading.Print()
}

func Start() {

	config.InitializeConfigs()

	LOGGER := logger.Logger()
	LOGGER.Infof("Initializing Server.......")

	server := web.NewServer(LOGGER)

	err := server.Start()
	if err != nil {
		panic("Server Shutting Down ... " + err.Error())
	}

}
