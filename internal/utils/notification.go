package utils

import (
	"github.com/gen2brain/beeep"
)

func Notify(msg string, icon []byte) {
	_ = beeep.Notify("SeeGO Launcher", msg, icon)
}
