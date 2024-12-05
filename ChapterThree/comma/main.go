package main

import (
	"bytes"
)

func main() {

}

func comma(s string) string {
	var buf bytes.Buffer
	buf.WriteByte('"')
	n := len(s)
	if n <= 3 {
		return s
	}
	pre := n % 3
	if pre > 0 {
		buf.WriteString(s[:pre])
		if n > pre {
			buf.WriteString(",")
		}
	}
	for i := pre; i < n; i += 3 {
		buf.WriteString(s[i : i+3])
		if i+3 < n {
			buf.WriteString(",")
		}
	}
	return buf.String()
}

// func comma2(s string) string {
// 	n := len(s)
// 	var o string
// 	var p string
// 	for i := 0; i < n; i++ {
// 		if s[i] == '.' {
// 			o = s[i:]
// 			p = s[:i]
// 		}
// 	}
// 	if len(p) <= 3 {
// 		return p + o
// 	}
// 	return comma2(s[:n-3]) + "," + s[n-3:] + o
// }

func comma2(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}

	var integerPart, fractionalPart string
	for i := 0; i < n; i++ {
		if s[i] == '.' {
			integerPart = s[:i]
			fractionalPart = s[i:]
			break
		}
	}
	if integerPart == "" {
		integerPart = s
	}

	n = len(integerPart)
	if n < 3 {
		return integerPart + fractionalPart
	}
	return comma2(integerPart[:n-3]) + "," + integerPart[n-3:] + fractionalPart
}

// func anagram(s, t string) bool {
// 	len1 := len(s)
// 	len2 := len(t)

// 	if len1 != len2 {
// 		return false
// 	}
// 	slice1 := make([]string, 0)

// 	for i := 0; i < len1; i++ {
// 		for j := 0; j < len2; j++{
// 			if i == j {
// 				slice1 = append(slice1,string( t[j]))
// 			}
// 		}
// 	}
// 	if len(slice1) != len1 {
// 		return false
// 	}
// 	return true
// }

func anagram(s, t string) bool {
	if len(s) != len(t) {
		return false
	}
	counts := make(map[rune]int)
	for _, char := range s {
		counts[char]++
	}
	for _, char := range t {
		counts[char]--
		if counts[char] < 0 {
			return false
		}
	}
	return true
}
