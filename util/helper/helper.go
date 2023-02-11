package helper

import "unicode/utf8"

func FindFirstInvalidUTF8Index(encodedData []byte) int {
	for i, b := range encodedData[len(encodedData)/2:] {
		if b < utf8.RuneSelf && b != 0x09 && b != 0x0A && b != 0x0D {
			continue
		}
		_, size := utf8.DecodeRune(encodedData[i:])
		if size == 1 {
			// Invalid UTF-8 encoded rune
			return i
		}
		i += size - 1
	}
	return -1
}
