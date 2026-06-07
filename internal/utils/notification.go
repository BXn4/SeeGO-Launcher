package utils

import (
	"github.com/gen2brain/beeep"
)

func Notify(msg string) {
	_ = beeep.Notify("SeeGO Launcher", msg, "")
}
