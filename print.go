package cfmt

import "fmt"

// Replaces entities with known ones
func replace(x []interface{}) []interface{} {
	if x == nil {
		return nil
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

func Print(x ...interface{}) {
	fmt.Print(replace(x)...)
}

func Printf(pattern string, x ...interface{}) {
	fmt.Printf(pattern, replace(x)...)
}
