package utils

import "hash/crc32"

// Base62 characters
const base62Chars = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"

// encodeBase62 converts an integer to a Base62 string
func encodeBase62(number uint32) string {
	if number == 0 {
		return "0"
	}
	encoded := ""
	base := uint32(len(base62Chars))
	for number > 0 {
		remainder := number % base
		encoded = string(base62Chars[remainder]) + encoded
		number /= base
	}
	return encoded
}

func GenerateShortString(longString string) string {
	hash := crc32.ChecksumIEEE([]byte(longString))
	shortStr := encodeBase62(hash)
	return shortStr
}