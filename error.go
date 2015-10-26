package cfmt

import (
	"fmt"
	"os"
)

// Shows error and stops application
func Panic(err error) {
	Print(err, "\n")
	os.Exit(1)
}

// Shows error and stops application
func Panicf(message string, args ...interface{}) {
	Panic(fmt.Errorf(message, args...))
}
