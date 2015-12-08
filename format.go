package cfmt

import (
	"fmt"
)

const (
	LEFT byte = iota
	RIGHT
	CENTER
)

// Formatter with default value
var formatter func(f Format) string = func(f Format) string {
	return TextFormat(f)
}

// Interface for self-formatable entries
type Formatted interface {
	Cfmt() []interface{}
}

// Represents color formatting
type Format struct {
	Value                    interface{}
	Fg, Bg                   int
	Bold, Intense, Underline bool
	Width                    int
	Align                    byte
}

// List of formats
type FormatList []interface{}

// String representation of inner value
func (f Format) String() string {
	return fmt.Sprintf("%v", f.Value)
}

// Returns true if format contains colors
func (f Format) HasColors() bool {
	return f.Fg > 0 || f.Bg > 0 || f.Bold || f.Underline
}

// Returns true if format contains data for data modifications
func (f Format) HasModification() bool {
	return f.Width > 0
}

// Returns copy of format with provided width
func (f Format) WithWidth(width int) Format {
	return Format{
		Value: f.Value,
		Fg: f.Fg,
		Bg: f.Bg,
		Bold: f.Bold,
		Intense: f.Intense,
		Underline: f.Underline,
		Width: width,
		Align: f.Align,
	}
}