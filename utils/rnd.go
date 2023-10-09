package rnd

import (
	"math/rand"
	"strings"
)

const (
	lowercase = "abcdefghijklmnopqrstuvwxyz"
	uppercase = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	numbers   = "0123456789"
)

const (
	LOWER_CASE = 1 << iota
	UPPER_CASE
	NUMBERS
)

func getCharset(flags uint) string {
	sb := strings.Builder{}
	if flags&LOWER_CASE != 0 {
		sb.WriteString(lowercase)
	}
	if flags&UPPER_CASE != 0 {
		sb.WriteString(uppercase)
	}
	if flags&NUMBERS != 0 {
		sb.WriteString(numbers)
	}
	return sb.String()
}

func RandomString(length int, flags uint) string {
	charset := getCharset(flags)
	result := make([]byte, length)
	for i := 0; i < length; i++ {
		result[i] = charset[rand.Intn(len(charset))]
	}
	return string(result)
}
