package main

import (
	"os"
)

const EXIT_SUCCESS = 0
const EXIT_FAILURE_DISPLAY_CREATION = 11
const EXIT_FAILURE_DISPLAY_SOCKET = 12
const EXIT_FAILURE_SWC_INITIALIZE = 13

func main() {
	display := CreateDisplay()
	if display == nil {
		os.Exit(EXIT_FAILURE_DISPLAY_CREATION)
	}
	if AddDisplaySocket(display) != 0 {
		os.Exit(EXIT_FAILURE_DISPLAY_SOCKET)
	}
	if !InitializeDisplay(display) {
		os.Exit(EXIT_FAILURE_SWC_INITIALIZE)
	}
	AddDebugKeyBindings()
	RunDisplay(display)
	DestroyDisplay(display)
	os.Exit(EXIT_SUCCESS)
}
