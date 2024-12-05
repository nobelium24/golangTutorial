package main

import (
	"unicode"
)

func main() {

}

func squash(s []byte) []byte {
	i := 0
	for _, j := range s {
		if unicode.IsSpace(rune(j)) {
			s[i] = ' '
			i++
		} else {
			s[i] = j
			i++
		}
	}
	return s[:i]
}
