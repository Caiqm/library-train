package str

import "strings"

func Stripos(haystack string, needle string, offset ...int) int {
	off := 0
	if len(offset) > 0 {
		off = offset[0]
	}
	if off > len(haystack) || off < 0 {
		return -1
	}
	haystack = strings.ToLower(haystack[off:])
	needle = strings.ToLower(needle)
	index := strings.Index(haystack, needle)
	if index != -1 {
		return off + index
	}
	return index
}
