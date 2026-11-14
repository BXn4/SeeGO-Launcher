package main

import (
	"embed"
	"seegolauncher/internal"

	"github.com/charmbracelet/log"
	"github.com/wailsapp/wails/v3/pkg/application"
)

//go:embed all:frontend/dist
var assets embed.FS

//go:embed frontend/src/public/images/seego_icon.png
var icon []byte

func init() {
	application.RegisterEvent[string]("time")
}

func main() {
	log.SetLevel(log.DebugLevel)

	log.Info("SeeGO Launcher by Fodor Bence (BXn4)")

	app := internal.CreateApp(assets, icon)

	if err := app.Run(); err != nil {
		log.Fatal(err)
	}
}
