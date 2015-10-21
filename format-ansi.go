package cfmt

import (
	"bytes"
	"strconv"
)

const (
	// Start ANSI color sequenct
	Start = "\033["
	// Bold
	Bold = "1;"
	// Reset is the ANSI reset escape sequence
	Reset = "\033[0m"
	// DefaultBG is the default background
	DefaultBG = "\033[49m"
	// DefaultFG is the default foreground
	DefaultFG = "\033[39m"
)

// Performs ANSI formatting of provided data
func AnsiFormat(f Format) string {
	str := f.String()
	if f.NoFormat() {
		return str
	}

	buf := bytes.NewBufferString(Start)
	if f.Bold {
		buf.WriteString("1;")
	} else if f.Underline {
		buf.WriteString("4;")
	} else {
		buf.WriteString("0;")
	}
	if f.Fg > 0 {
		if f.Fg > 15 {
			buf.WriteString("38;5;")
			buf.WriteString(strconv.Itoa(f.Fg))
		} else {
			clr := f.Fg + 29
			if f.Intense {
				clr += 60
			}
			buf.WriteString(strconv.Itoa(clr))
		}
		buf.WriteString(";")
	}
	if f.Bg > 0 {
		clr := f.Bg + 39
		buf.WriteString(strconv.Itoa(clr))
		buf.WriteString(";")
	}
	buf.WriteString("m")

	buf.WriteString(str)
	buf.WriteString(Reset)

	return buf.String()
}

func init() {
	formatter = AnsiFormat
}