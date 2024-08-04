package main

import (
	"strings"
)

const shift = 6

func caesarCipher(input string, shift int32) string {
	var sb strings.Builder
	for _, r := range input {
		switch {
		case 'A' <= r && r <= 'Z':
			sb.WriteRune('A' + (r-'A'+shift)%26)
		case 'a' <= r && r <= 'z':
			sb.WriteRune('a' + (r-'a'+shift)%26)
		default:
			sb.WriteRune(r)
		}
	}

	return sb.String()
}
