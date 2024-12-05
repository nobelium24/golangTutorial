package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"unicode"
)

func main() {
	letter, number, invalid := make(map[rune]int), make(map[rune]int), 0
	in := bufio.NewReader(os.Stdin)

	for {
		r, n, err := in.ReadRune()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "charcount: %v\n", err)
			os.Exit(1)
		}
		if r == unicode.ReplacementChar && n == 1 {
			invalid += 1
		}
		if unicode.IsLetter(r) {
			letter[r]++
		} else if unicode.IsNumber(r) {
			number[r]++
		}
	}

	fmt.Print("letter \t count \n")
	for i, j := range letter {
		fmt.Printf("%q \t %d \n", i, j)
	}

	fmt.Print("number \t count \n")
	for k, l := range number {
		fmt.Printf("%q \t %d \n", k, l)
	}
}

func charCount() {

}
