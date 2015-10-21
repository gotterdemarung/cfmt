package cfmt

import (
	"time"
	"bytes"
)

// Returns format for nil value
func FNil() Format {
	return styles.Get(S_NIL, "<nil>")
}

// Returns format for strings
func FString(s string) Format {
	return styles.Get(S_STRING, s)
}

// Returns format for string slice
func FStringl(s []string) Format {
	if s == nil {
		return FNil()
	} else if len(s) == 0 {
		return styles.Get(S_STRING, "")
	}
	buf := bytes.NewBufferString("")
	for i, v := range s {
		if i > 0 {
			buf.WriteString(", ")
		}
		buf.WriteString(v)
	}
	return styles.Get(S_STRING, buf.String())
}

// Returns format for integers
func FInt(x int) Format {
	return styles.Get(S_INT, x)
}

// Returns format for types (structs, etc.)
func FType(x interface{}) Format {
	return styles.Get(S_TYPE, x)
}

// Returns format for map keys
func FKey(x string) Format {
	return styles.Get(S_KEY, x)
}

// Returns format for durations
func FDuration(x time.Duration) Format {
	return styles.Get(S_DURATION, x)
}
