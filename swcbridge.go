package main

// #cgo LDFLAGS: -L./lib/ -L./wrapper -lswc -lwayland-server -lm
// #include "lib/wm.c"
// #include <swc.h>
// #include <wayland-server.h>
import "C"

import (
	"unsafe"
	"os/exec"
)

func AddKeyBinding(modifiers,value int, callback func()) {
	C.swc_add_binding(C.SWC_BINDING_KEY, C.uint32_t(modifiers), C.uint32_t(value), C.swc_binding_handler(unsafe.Pointer(C.execute_binding_callback)), unsafe.Pointer(&callback))
}

type StackingType int
const (
	HORIZONTAL StackingType = 0
	VERTICAL
	TABBED
)

type Container struct {
	values []Container
	orientation StackingType
}

type Screen struct {
	swc_screen unsafe.Pointer
	windows unsafe.Pointer
	windowcount int
}

type Window struct {
	swc_window unsafe.Pointer
	screen *Screen
	link unsafe.Pointer
}

type Display unsafe.Pointer

func CreateDisplay() Display {
	var display Display = Display(C.wl_display_create())
	return display
}

func AddDisplaySocket(display Display) int {
	return int(C.wl_display_add_socket((*C.struct_wl_display)(display), nil))
}

func InitializeDisplay(display Display) bool {
	C.display = (*C.struct_wl_display)(display)
	return bool(C.swc_initialize((*C.struct_wl_display)(display), nil, &C.manager))
}

func RunDisplay(display Display) {
	C.wl_display_run((*C.struct_wl_display)(display))
}

func TerminateDisplay(display Display) {
	C.wl_display_terminate((*C.struct_wl_display)(display))
}

func DestroyDisplay(display Display) {
	C.wl_display_destroy((*C.struct_wl_display)(display))
}

func StartTerminal() {
	cmd := exec.Command("st")
	cmd.Start()
}

func StopDisplay() {
	C.wl_display_terminate(C.display) //TODO
}

func AddDebugKeyBindings() {
	AddKeyBinding(C.SWC_MOD_LOGO, C.XKB_KEY_Return, StartTerminal)
	AddKeyBinding(C.SWC_MOD_LOGO, C.XKB_KEY_x, StopDisplay)
	//event_loop := C.wl_display_get_event_loop((*wl_display)display)
}
