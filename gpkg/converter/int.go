package converter

import "strconv"

func IntToStr[T ~int | ~int32 | ~int64](i T) string {
	return strconv.FormatInt(int64(i), 10)
}
