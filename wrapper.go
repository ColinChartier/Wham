package main

import "C"

import (
	"unsafe"
)

//export call_callback
func call_callback(callback unsafe.Pointer) {
	(*(*func())(callback))()
}
