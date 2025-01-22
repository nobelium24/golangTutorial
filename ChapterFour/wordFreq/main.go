package main

import (
	"bufio"
	"fmt"
	"os"
)

// Write a program wordfreq to report the frequency of each word in an input text
// file. Call input.Split(bufio.ScanWords) before the first call to Scan to break the input into
// words instead of lines.
func main() {
	words := make(map[string]int)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	for scanner.Scan() {
		word := scanner.Text()
		words[word]++
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "Error reading file", err)
		os.Exit(1)
	}

	for word, count := range words {
		fmt.Printf("%s: %d\n", word, count)
	}
}
