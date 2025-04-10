package converter

import "unsafe"

func ZeroCopyByteToString(b []byte) string {
	return *((*string)(unsafe.Pointer(&b)))
}

func ZeroCopyStringToBytes(s string) []byte {
	return unsafe.Slice(unsafe.StringData(s), len(s))
}
