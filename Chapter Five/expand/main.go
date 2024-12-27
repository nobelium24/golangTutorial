// package main

// import "strings"

// func main() {

// }

// func expand(s string, f func(string) string) string {
// 	// newStrArr := strings.Split(s, "$")
// 	// newStr := newStrArr[0]
// 	// return f(newStr)
// 	newS := strings.Split(s, " ")

// 	for index, word := range newS {
// 		if strings.Contains(word, "$") {
// 			newWord := strings.Replace(word, "$", "", -1)
// 			newWord = f(newWord)
// 			newS[index] = newWord
// 		}
// 	}
// 	return strings.Join(newS, " ")
// }

// func f(input string) string {
// 	if input == "name" {
// 		return "Alice"
// 	} else if input == "place" {
// 		return "Wonderland"
// 	}
// 	return ""
// }

package main

import (
	"strings"
	"unicode"
)

func main() {
	s := "Hello $name, welcome to $place!"
	result := expand(s, f)
	println(result) // Expected: "Hello Alice, welcome to Wonderland!"
}

func expand(s string, f func(string) string) string {
	var result strings.Builder      // Efficiently build the resulting string.
	var placeholder strings.Builder // Temporarily store the placeholder name.
	inPlaceholder := false          // Track if we're processing a placeholder.

	for _, r := range s { // Iterate through each rune (character) in the string.
		if r == '$' { // Check if the character starts a placeholder.
			inPlaceholder = true       // Start tracking the placeholder.
			if placeholder.Len() > 0 { // If there's an unprocessed placeholder, handle it.
				result.WriteString(f(placeholder.String())) // Replace it using `f`.
				placeholder.Reset()                         // Clear the placeholder buffer.
			}
		} else if inPlaceholder && (unicode.IsLetter(r) || unicode.IsDigit(r)) {
			// If we're in a placeholder and the character is alphanumeric, add it.
			placeholder.WriteRune(r)
		} else {
			// If we encounter a non-placeholder character, finalize the placeholder.
			if inPlaceholder {
				inPlaceholder = false                       // Exit placeholder mode.
				result.WriteString(f(placeholder.String())) // Replace using `f`.
				placeholder.Reset()                         // Clear the buffer for the next placeholder.
			}
			result.WriteRune(r) // Add the current character to the result.
		}
	}

	// Handle any leftover placeholder at the end of the string.
	if placeholder.Len() > 0 {
		result.WriteString(f(placeholder.String()))
	}

	return result.String() // Convert the builder content to a string and return it.
}

func f(input string) string {
	// Replacement logic based on the placeholder name.
	if input == "name" {
		return "Alice"
	} else if input == "place" {
		return "Wonderland"
	}
	return "" // Default to an empty string if no match is found.
}
