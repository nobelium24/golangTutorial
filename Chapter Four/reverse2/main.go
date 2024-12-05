package main

import "unicode/utf8"

func main() {

}

func reverse2(s []byte) []byte {
	var k []int
	for i := 0; i < len(s); i++ {
		_, size := utf8.DecodeRune(s[i:])
		k = append(k, size)
		i += size
	}
	for i, j := 0, len(k)-2; i < j; i, j = i+1, j-1 {
		reverse(s[k[i]:k[i+1]], s[k[j]:k[j+1]])
	}
	return s
}

func reverse(a, b []byte) {
	if len(a) != len(b) {
		panic("Length must eb the same")
	}
	for i := range a {
		a[i], b[i] = b[i], a[i]
	}

}
