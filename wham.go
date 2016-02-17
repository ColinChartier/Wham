package main

// #cgo LDFLAGS: -L./lib/ -lswc -lwayland-server -lm
// #include "lib/wm.c"
// #include <swc.h>
// #include <wayland-server.h>
import "C"

import (
	"os"
	"unsafe"
)

const EXIT_SUCCESS = 0
const EXIT_FAILURE_DISPLAY_CREATION = 11
const EXIT_FAILURE_DISPLAY_SOCKET = 12
const EXIT_FAILURE_SWC_INITIALIZE = 13

var terminal_command = [...]*C.char {C.CString("st"), nil}

func AddKeyBinding(modifiers,value int, callback,data unsafe.Pointer) {
	C.swc_add_binding(C.SWC_BINDING_KEY, C.uint32_t(modifiers), C.uint32_t(value), C.swc_binding_handler(callback), data)
}

func main() {
	display := C.wl_display_create()
	if display == nil {
		os.Exit(EXIT_FAILURE_DISPLAY_CREATION)
	}
	if C.wl_display_add_socket(display, nil) != 0 {
		os.Exit(EXIT_FAILURE_DISPLAY_SOCKET)
	}
	if !C.swc_initialize(display, nil, &C.manager) {
		os.Exit(EXIT_FAILURE_SWC_INITIALIZE)
	}
	C.display = display
	AddKeyBinding(C.SWC_MOD_LOGO, C.XKB_KEY_Return, unsafe.Pointer(C.spawn), unsafe.Pointer(&terminal_command))
	AddKeyBinding(C.SWC_MOD_LOGO, C.XKB_KEY_q, unsafe.Pointer(C.quit), nil)
	//event_loop := C.wl_display_get_event_loop(display)
	C.wl_display_run(display)
	C.wl_display_destroy(display)
	os.Exit(EXIT_SUCCESS)
}
