package main

import (
	"embed"
	"io"
	"os"
	"seegolauncher/internal"
	"seegolauncher/internal/utils"

	"github.com/charmbracelet/log"
)

//go:embed all:frontend/dist
var assets embed.FS

//go:embed frontend/src/public/images/seego_icon.png
var icon []byte

func main() {
	log.SetOutput(io.MultiWriter(os.Stdout, &utils.LogBuffer))
	log.SetLevel(log.DebugLevel)
	log.SetReportTimestamp(true)
	log.SetReportCaller(true)

	log.Info("SeeGO Launcher by Fodor Bence (BXn4)")

	app := internal.CreateApp(assets, icon)

	if err := app.Run(); err != nil {
		log.Error(err)
		utils.SaveLog()
		os.Exit(1)
	}
}
