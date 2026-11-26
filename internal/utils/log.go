package utils

import (
	"bytes"
	"os"
	"time"

	"github.com/charmbracelet/log"
)

var LogFileName string
var LogBuffer bytes.Buffer

func SaveLog() {
	currentTime := time.Now().Format("2006-01-02_15-04-05")
	LogFileName = "seego-launcher-" + currentTime + ".log"

	log.Infof("Created a new log file: %s", LogFileName)
	_ = os.WriteFile(LogFileName, LogBuffer.Bytes(), 0o644)
}
