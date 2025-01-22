package main

import "fmt"

func main() {

}

func reverse(s *[]string) {
	for i, j := 0, len(*s)-1; i < j; i, j = i+1, j-1 {
		(*s)[i], (*s)[j] = (*s)[j], (*s)[i]
	}
}

func rotate(s []int, n int) {
	if n > len(s) {
		fmt.Printf("%v is invalid", n)
		return
	}
	temp := make([]int, n)
	copy(temp, s[:n])

	copy(s, s[n:])

	copy(s[len(s)-n:], temp)
}

func deleteItem(s []string, i int) []string {
	return append(s[:i], s[i+1:]...)
}

func elimDup(s []string) []string {
	j := 0
	for i := 0; i < len(s); i++ {
		if s[j] != s[i] {
			j++
			s[j] = s[i]
		}
	}
	return s[:j+1]
}
