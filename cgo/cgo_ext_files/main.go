package main

// #cgo CFLAGS: -I${SRCDIR}/cLib
// #cgo LDFLAGS: ${SRCDIR}/cLib.a
// #include <stdlib.h>
// #include <cLib.h>
import "C"

import (
	"unsafe"
)

func main() {
	C.printHello()
	message := C.CString("we need more Go coders !!")
	defer C.free(unsafe.Pointer(message))
	C.printMessage(message)
}
