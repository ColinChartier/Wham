package main

import (
	"os"
	"os/exec"
)

const EXIT_SUCCESS = 0
const EXIT_FAILURE_DISPLAY_CREATION = 11
const EXIT_FAILURE_DISPLAY_SOCKET = 12
const EXIT_FAILURE_SWC_INITIALIZE = 13

var display Display

func AddDefaultKeyBindings() {
	AddKeyBinding(KEY_MODIFIER_LOGO, 65293, func() {
	cmd := exec.Command("st")
	cmd.Start()
	})
	AddKeyBinding(KEY_MODIFIER_LOGO | KEY_MODIFIER_ALT, 113, func() { //alt+win+q to qit
	TerminateDisplay(display)
	})
}


func main() {
	display = CreateDisplay()
	if display == nil {
		os.Exit(EXIT_FAILURE_DISPLAY_CREATION)
	}
	if AddDisplaySocket(display) != 0 {
		os.Exit(EXIT_FAILURE_DISPLAY_SOCKET)
	}
	if !InitializeDisplay(display) {
		os.Exit(EXIT_FAILURE_SWC_INITIALIZE)
	}
	AddDefaultKeyBindings()
	RunDisplay(display)
	DestroyDisplay(display)
	os.Exit(EXIT_SUCCESS)
}
