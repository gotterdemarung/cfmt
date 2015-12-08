package cfmt

import (
	"fmt"
	"io"
	"os"
)

// Replaces entities with known ones
func replace(x []interface{}) []interface{} {
	if x == nil {
		return nil
	} else if len(x) == 0 {
		return x
	}

	y := []interface{}{}

	for _, j := range x {
		// Real format
		if f, ok := j.(Format); ok {
			y = append(y, formatter(f))
		} else if f, ok := j.(FormatList); ok {
			y = append(y, replace(f)...)
		} else if f, ok := j.(Formatted); ok {
			y = append(y, replace(f.Cfmt())...)
		} else if f, ok := j.(error); ok {
			y = append(y, formatter(styles.Get(S_ERROR, f)))
		} else {
			y = append(y, j)
		}
	}

	return y
}

func Fprint(w io.Writer, x ...interface{}) {
	fmt.Fprint(w, replace(x)...)
}

func Fprintf(w io.Writer, pattern string, x ...interface{}) {
	fmt.Fprintf(w, pattern, replace(x)...)
}

func Printf(pattern string, x ...interface{}) {
	Fprintf(os.Stdout, pattern, replace(x)...)
}

func Println(x ...interface{}) {
	if len(x) == 0 {
		fmt.Fprintf(os.Stdout, "\n")
	} else {
		y := append(x, "\n")
		Print(y...)
	}
}

func Print(x ...interface{}) {
	os.Stdout.Fd()
	Fprint(os.Stdout, x...)
}
