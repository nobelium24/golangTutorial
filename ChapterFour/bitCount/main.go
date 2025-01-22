package main

import "crypto/sha256"

func main() {

}

// func bitCount(s, t string) int {
// 	c1 := sha256.Sum256([]byte(s))
// 	c2 := sha256.Sum256([]byte(t))
// 	x := 0

// 	counts := make(map[byte]int)
// 	for _, item := range c1 {
// 		counts[item]++
// 	}

// 	for _, item := range c2 {
// 		counts[item]--
// 		if counts[item] >= 0 {
// 			x++
// 		}
// 	}
// 	return x
// }

func bitCount(s, t string) int {
	c1 := sha256.Sum256([]byte(s))
	c2 := sha256.Sum256([]byte(t))

	count := 0
	for i := 0; i < len(c1); i++ {
		diff := c1[i] ^ c2[i]
		for diff > 0 {
			count += int(diff & 1)
			diff >>= 1
		}
	}
	return count
}
