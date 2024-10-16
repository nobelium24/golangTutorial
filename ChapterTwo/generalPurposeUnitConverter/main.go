package generalpurposeunitconverter

import (
	"ChapterTwo/tempconv"
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {
	args := os.Args[1:]
	length := len(args)
	if length != 0 {
		for _, i := range args {
			t, err := strconv.ParseFloat(i, 64)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error: %v\n", err)
				continue
			}
			f := tempconv.Fahrenheit(t)
			c := tempconv.Celsius(t)
			fmt.Printf("%s = %s, %s = %s\n",
				f, tempconv.FToC(f), c, tempconv.CToF(c))
		}
	} else {
		input := bufio.NewScanner(os.Stdin)
		for input.Scan() {
			t, err := strconv.ParseFloat(input.Text(), 64)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error: %v\n", err)
				continue
			}
			f := tempconv.Fahrenheit(t)
			c := tempconv.Celsius(t)

			fileName := "unitConverter.txt"

			file, err := os.Create(fileName)
			if err != nil {
				fmt.Printf("An error occurred: %v\n", err)
				continue
			}
			defer file.Close()
			sentence := fmt.Sprintf("%s = %s, %s = %s\n",
				f, tempconv.FToC(f), c, tempconv.CToF(c))
			reader := strings.NewReader(sentence)
			bytesWritten, err := io.Copy(file, reader)
			if err != nil {
				fmt.Printf("An error occurred while writing to the file: %v\n", err)
				continue
			}
			fmt.Printf("Bytes written: %d\n", bytesWritten)
		}
	}
}
