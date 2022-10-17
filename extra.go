package cimgui

// #include "cimgui/cimgui.h"
// #include "extra.h"
import "C"
import (
	"unsafe"
)

var getClipboardTextFn = func() string { return "" }
var setClipboardTextFn = func(text string) {}
var dropLastClipboardText = func() {}

func (self ImGuiIO) SetClipboardTextFn(fn func(text string)) {
	setClipboardTextFn = fn
	C.RegisterClipboardFunctions(self.handle())
}

func (self ImGuiIO) GetClipboardTextFn(fn func() string) {
	getClipboardTextFn = fn
	C.RegisterClipboardFunctions(self.handle())
}

//export setClipboardText
func setClipboardText(userData unsafe.Pointer, text *C.char) {
	dropLastClipboardText()
	setClipboardTextFn(C.GoString(text))
}

//export getClipboardText
func getClipboardText(iuserData unsafe.Pointer) *C.char {
	dropLastClipboardText()
	text := getClipboardTextFn()
	textPtr, textFin := wrapString(text)
	dropLastClipboardText = func() {
		dropLastClipboardText = func() {}
		textFin()
	}
	return textPtr
}
