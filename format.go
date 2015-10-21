package cfmt

import (
	"fmt"
)

// Formatter with default value
var formatter func(f Format) string = func(f Format) string {
	return f.String()
}

// Represents color formatting
type Format struct {
	Value                    interface{}
	Fg, Bg                   int
	Bold, Intense, Underline bool
}

// String representation of inner value
func (f Format) String() string {
	return fmt.Sprintf("%v", f.Value)
}

// Returns true if format contains no formatting
func (f Format) NoFormat() bool {
	return f.Fg == 0 && f.Bg == 0 && !f.Bold
}
