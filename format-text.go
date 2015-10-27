package cfmt

import (
	"strings"
)

func TextFormat(f Format) string {
	str := f.String()

	if f.Width > 0 && len(str) < f.Width {
		// Padding
		str = str + strings.Repeat(" ", f.Width-len(str))
	}

	return str
}
