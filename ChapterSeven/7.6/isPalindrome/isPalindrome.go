package ispalindrome

import "sort"

type Palindrome []rune

func (p Palindrome) Len() int {
	return len(p)
}

func (p Palindrome) Less(i, j int) bool {
	return p[i] < p[j]
}

func (p Palindrome) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func IsPalindrome(s sort.Interface) bool {
	one := 0
	two := s.Len() - 1
	for one < two {
		if s.Less(one, two) || s.Less(two, one) {
			return false
		}
		one++
		two--
	}
	return true
}
