package varidiacfunc

import (
	"errors"
	"strings"
)

func Max(args ...int) (int, error) {
	var num int
	if len(args) < 1 {
		return 0, errors.New("no arguments provided")
	} else if len(args) == 1 {
		return args[0], nil
	} else {
		max := args[0]
		for _, i := range args {
			if i > max {
				max = i
			}
		}
		num = max
	}
	return num, nil
}

func Min(args ...int) (int, error) {
	var num int
	if len(args) < 1 {
		return 0, errors.New("No arguments provided")
	} else if len(args) == 1 {
		return args[0], nil
	} else {
		max := args[0]
		for _, i := range args {
			if i > max {
				max = i
			}
		}
		num = max
	}
	return num, nil
}

func JoinStrings(separator string, args ...string) string {
	var str string
	if len(args) < 1 {
		return ""
	}
	var formatted strings.Builder
	formatted.WriteString(args[0])
	for _, i := range args[1:] {
		formatted.WriteString(separator)
		formatted.WriteString(i)
	}

	return str
}
