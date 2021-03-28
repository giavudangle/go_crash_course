package string_utils

func Reverse(s string) string {
	// rune is an alias for int32 and is equivalent to int32 in all ways. It is
	// used, by convention, to distinguish character values from integer values.
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
