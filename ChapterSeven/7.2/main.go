package main

import (
	"./reader"
	"fmt"
	"io"
)

func main() {
	input := "Hello, World!"
	reader := reader.NewReader(input)
	buf := make([]byte, 5)

	for {
		n, err := reader.Read(buf)
		fmt.Printf("Read %d bytes: %q\n", n, buf[:n])

		if err == io.EOF {
			fmt.Println("End of string reached")
			break
		}
	}
}
