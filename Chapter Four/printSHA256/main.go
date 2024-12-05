package main

import (
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
	"io"
	"os"
)

func main() {

}

// func sha() {
// 	input := os.Args[1]
// 	cliFlag := os.Args[2]

// 	if cliFlag != "" {
// 		if cliFlag == "384" {
// 			sha := sha512.Sum384([]byte(input))
// 			fmt.Printf("SHA-384: %x\n", sha)
// 			return
// 		} else if cliFlag == "512" {
// 			sha := sha512.Sum512([]byte(input))
// 			fmt.Printf("SHA-512: %x\n", sha)
// 			return
// 		} else {
// 			fmt.Print("Invalid command-line flag. Must be 384 or 512")
// 		}
// 	} else {
// 		sha := sha256.Sum256([]byte(input))
// 		fmt.Printf("SHA-256: %x\n", sha)
// 	}
// }

func sha() {
	// Define command-line flags
	flagSHA384 := flag.Bool("384", false, "Use SHA-384 hash")
	flagSHA512 := flag.Bool("512", false, "Use SHA-512 hash")
	flag.Parse()

	input, err := io.ReadAll(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading input: %v\n", err)
		os.Exit(1)
	}

	if *flagSHA384 {
		sha := sha512.Sum384(input)
		fmt.Printf("SHA-384: %x\n", sha)
	} else if *flagSHA512 {
		sha := sha512.Sum512(input)
		fmt.Printf("SHA-512: %x\n", sha)
	} else {
		sha := sha256.Sum256(input)
		fmt.Printf("SHA-256: %x\n", sha)
	}
}
