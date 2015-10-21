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

	y := make([]interface{}, len(x))

	for i, j := range x {
		// Real format
		if f, ok := j.(Format); ok {
			y[i] = formatter(f)
		} else if f, ok := j.(error); ok {
			y[i] = formatter(styles.Get(S_ERROR, f))
		} else {
			y[i] = j
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
		Print(x...)
		Print("\n")
	}
}

func Print(x ...interface{}) {
	Fprint(os.Stdout, x...)
}


