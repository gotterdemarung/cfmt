package cfmt

import (
	"fmt"
	"os"
)

// Shows error and stops script
func Panic(err error) {
	Print(err, "\n")
	os.Exit(1)
}

func Panicf(message string, args ...interface{}) {
	Panic(fmt.Errorf(message, args...))
}
