package generalpurposeunitconverter

import (
	"ChapterTwo/tempconv"
	"bufio"
	"fmt"
	"os"
	"strconv"
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
			fmt.Printf("%s = %s, %s = %s\n",
				f, tempconv.FToC(f), c, tempconv.CToF(c))
		}
	}
}
