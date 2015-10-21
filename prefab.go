package cfmt

import "time"

// Returns format for nil value
func FNil() Format {
	return styles.Get(S_NIL, "<nil>")
}

// Returns format for strings
func FString(s string) Format {
	return styles.Get(S_STRING, s)
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
