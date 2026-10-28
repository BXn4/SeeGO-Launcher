package utils

import (
	"github.com/charmbracelet/log"
	"github.com/gen2brain/beeep"
)

func Notify(msg string, icon []byte, key string) {
	_ = beeep.Notify("SeeGO Launcher", msg, icon)
	log.Debugf("Showing notification: %s", key)
}
