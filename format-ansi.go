// +build darwin linux

package cfmt

import (
	"bytes"
	"strconv"
	"strings"
)

const (
	// Start ANSI color sequence
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

	if f.Width > 0 && len(str) < f.Width {
		// Padding
		str = str + strings.Repeat(" ", f.Width-len(str))
	}

	if !f.HasColors() {
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
	}
	if f.Bg > 0 {
		if f.Fg > 0 {
			buf.WriteString(";")
		}
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
