package main

import (
	"embed"
	"seegolauncher/internal"
	"seegolauncher/internal/services"

	"github.com/charmbracelet/log"
	"github.com/joho/godotenv"
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
	envFile, err := godotenv.Read(".env")

	hasConfig := err == nil

	if !hasConfig {
		log.Warnf("Cannot find .env file!")
	}

	if services.OA == "" {
		services.OA = envFile["OA"]
	}

	log.Info("SeeGO Launcher by BXn4")

	app := internal.CreateApp(assets, icon)

	if err := app.Run(); err != nil {
		log.Fatal(err)
	}
}
