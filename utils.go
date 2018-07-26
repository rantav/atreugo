package atreugo

import (
	"unsafe"
)

func panicOnError(err error) {
	if err != nil {
		panic(err)
	}
}

// b2s convert bytes array to string without memory allocation (non safe)
func b2s(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

// int64ToInt convert int64 to int without memory allocation (non safe)
func int64ToInt(i int64) int {
	return *(*int)(unsafe.Pointer(&i))
}

// index returns the first index of the target string `t`, or
// -1 if no match is found.
func indexOf(vs []string, t string) int {
	for i, v := range vs {
		if v == t {
			return i
		}
	}
	return -1
}

// include returns `true` if the target string t is in the
// slice.
func include(vs []string, t string) bool {
	return indexOf(vs, t) >= 0
}