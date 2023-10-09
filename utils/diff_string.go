package utils

import "strings"

func DiffString(a string, b string) {

}

func DiffRunicString(a []rune, b []rune) {
	if len(a) > len(b) {
		a, b = b, a
	}
	sb := strings.Builder{}
	for i, r := range a {
		if r != b[i] {
			//
		} else {
			//
		}
	}
}
