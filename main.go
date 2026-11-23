package main

import (
	"bytes"
	"embed"
	"io"
	"os"
	"seegolauncher/internal"
	"time"

	"github.com/charmbracelet/log"
)

//go:embed all:frontend/dist
var assets embed.FS

//go:embed frontend/src/public/images/seego_icon.png
var icon []byte

func main() {
	var logBuffer bytes.Buffer
	currentTime := time.Now().Format("2006-01-02_15-04-05")
	logFile := "seego-launcher-" + currentTime + ".log"

	log.SetOutput(io.MultiWriter(os.Stdout, &logBuffer))
	log.SetLevel(log.DebugLevel)
	log.SetReportTimestamp(true)
	log.SetReportCaller(true)

	saveLog := func() {
		log.Infof("Created a new log file: %s", logFile)
		_ = os.WriteFile(logFile, logBuffer.Bytes(), 0o644)
	}

	log.Info("SeeGO Launcher by Fodor Bence (BXn4)")

	app := internal.CreateApp(assets, icon, logFile, saveLog)

	if err := app.Run(); err != nil {
		log.Error(err)
		saveLog()
		os.Exit(1)
	}
}
