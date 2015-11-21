package cfmt

import (
	"bytes"
	"strings"
	"time"
)

// Detects type of provided value and returns most suitable format
// for it
func FAuto(in interface{}) Format {
	if in == nil {
		return FNil()
	}
	if x, ok := in.(string); ok {
		return FString(x)
	}
	if x, ok := in.(int); ok {
		return FInt(x)
	}
	if x, ok := in.(bool); ok {
		return FBool(x)
	}
	if x, ok := in.(error); ok {
		return FError(x)
	}
	if x, ok := in.(Format); ok {
		return x
	}

	return Format{
		Value: in,
	}
}

func FHeader(h string) FormatList {
	return []interface{}{
		" ",
		Format{Value: h, Fg: 70},
		"\n",
		" ",
		Format{Value: strings.Repeat("─", len(h)), Fg: 64},
		"\n",
	}
}

// Returns format for nil value
func FNil() Format {
	return styles.Get(S_NIL, "<nil>")
}

// Returns format for errors
func FError(value interface{}) Format {
	if x, ok := value.(error); ok {
		return styles.Get(S_ERROR, x.Error())
	} else {
		return styles.Get(S_ERROR, value)
	}
}

// Returns true/false formatting for bools
func FBool(value bool) Format {
	if value {
		return styles.Get(S_BOOL, "true")
	} else {
		return styles.Get(S_BOOL, "false")
	}
}

// Returns yes/no formatting for bools
func FYesNo(value bool) Format {
	if value {
		return styles.Get(S_BOOL, "yes")
	} else {
		return styles.Get(S_BOOL, "no")
	}
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
