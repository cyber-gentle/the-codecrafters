// CodeCrafters — Operation Gopher Protocol
// Module: String Transformer
// Author: Abraham David
// Squad:  The Gopher's

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func toUpperCase(s string) string {
	return strings.ToUpper(s)
}

func toLowerCase(s string) string {
	return strings.ToLower(s)
}

func toSnakeCase(s string) string {
	word := strings.ReplaceAll(s, " ", "_")

	var result strings.Builder

	for _, char := range word {
		isLetter := (char >= 'a' && char <= 'z') || (char >= 'A' && char <= 'Z')
		isDigit := char >= '0' && char <= '9'
		isUnderscore := char == '_'

		if isLetter || isDigit || isUnderscore {
			result.WriteRune(char)
		}
	}
	return strings.ToLower(result.String())
}

func titleCase(s string) string {
	smallWords := map[string]bool{
		"a": true, "an": true, "the": true, "and": true,
		"but": true, "or": true, "for": true, "nor": true,
		"on": true, "at": true, "to": true, "by": true,
		"in": true, "of": true, "up": true, "as": true,
		"is": true, "it": true,
	}

	words := strings.Fields(s)

	for i, word := range words {
		if len(word) == 0 {
			continue
		}
		lowerWord := strings.ToLower(word)

		if i == 0 || !smallWords[lowerWord] {
			words[i] = strings.ToUpper(string(lowerWord[0])) + lowerWord[1:]
		} else {
			words[i] = lowerWord
		}
	}

	return strings.Join(words, " ")
}

func reverseWords(s string) string {
	words := strings.Fields(s)

	for i, word := range words {
		runes := []rune(word)

		left := 0
		right := len(runes) - 1
		for left < right {
			runes[left], runes[right] = runes[right], runes[left]
			left++
			right--
		}

		words[i] = string(runes)
	}

	return strings.Join(words, " ")
}

func palindrome(s string) string {
	cleaned := strings.ToLower(strings.ReplaceAll(s, " ", ""))

	runes := []rune(cleaned)
	left := 0
	right := len(runes) - 1
	for left < right {
		runes[left], runes[right] = runes[right], runes[left]
		left++
		right--
	}
	reversed := string(runes)

	if cleaned == reversed {
		return fmt.Sprintf("✦ %q is a palindrome!", s)
	}
	return fmt.Sprintf("✗ %q is not a palindrome.", s)
}

func main() {

start:
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("\n SENTINEL STRING TRANSFORMER — ONLINE\n ──────────────────────────────────────\n \n")
	fmt.Print(" > Enter command: ")

	input_word, _ := reader.ReadString('\n') // User's input
	input_word = strings.TrimSpace(input_word)
	split := strings.Fields(input_word) // Splitting of User's input into slice of string
	cmd := strings.ToLower(split[0])    // getting command from User's input to lower case

	word := split[1:] // Get word from  user input in form of slice of strings
	if len(word) == 0 {
		fmt.Printf("✗ No text provided. Usage: %v <text>\n \n", cmd)
		goto start

	}
	words := strings.Join(word, " ") // Joining word to back from slice to string

	for {
		switch cmd {
		case "upper":
			fmt.Printf(" → %v\n \n", toUpperCase(words))

		case "lower":
			fmt.Printf(" → %v\n \n", toLowerCase(words))

		case "snake":
			fmt.Printf(" → %v\n \n", toSnakeCase(words))

		case "title":
			fmt.Printf(" → %v\n \n", titleCase(words))

		case "reverse":
			fmt.Printf(" → %v\n \n", reverseWords(words))

		case "palindrome":
			fmt.Printf(" → %q %v\n \n", words, palindrome(words))

		case "exit":
			fmt.Print(" Shutting down String Transformer. Goodbye.\n \n")
			return

		default:
			fmt.Printf(" ✗ Unknown command: %q\n Valid commands: upper, lower, cap, title, snake, reverse, exit \n \n", cmd)
		}
	}

}
