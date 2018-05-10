package goshorts

import (
	"log"
	"os"
	"runtime"
)

var (
	ExitOnErr = true
)

func ThrowError(err error) {
	_, _, line, _ := runtime.Caller(2)
	log.Println("We've encountered an unrecoverable issue. Please review the documentation and try again")
	log.Fatalf("Function: %s at line %v failed with error: %s", getCaller(), line, err.Error())
	if ExitOnErr {
		os.Exit(1)
	}
}

func ErrCheck(err error) {
	if err != nil {
		ThrowError(err)
	}
}

func getCaller() string {

	// we get the callers as uintptrs - but we just need 1
	fpcs := make([]uintptr, 1)

	// skip 3 levels to get to the caller of whoever called Caller()
	n := runtime.Callers(4, fpcs)
	if n == 0 {
		return "n/a" // proper error her would be better
	}

	// get the info of the actual function that's in the pointer
	fun := runtime.FuncForPC(fpcs[0] - 1)
	if fun == nil {
		return "n/a"
	}

	// return its name
	return fun.Name()
}
