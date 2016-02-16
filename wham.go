package main

// #cgo LDFLAGS: -L./lib/ -lswc -lwayland-server -lm
// #include "lib/wm.c"
import "C"

import (
	"os"
)

const EXIT_SUCCESS = 0
const EXIT_FAILURE_DISPLAY_CREATION = 11
const EXIT_FAILURE_DISPLAY_SOCKET = 12
const EXIT_FAILURE_SWC_INITIALIZE = 13

func main() {
	os.Exit(int(C.wololo()))
}
