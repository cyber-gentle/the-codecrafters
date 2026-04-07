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

func toUpperCase(s, cmd string) string {
	if s == "" {
		fmt.Printf(" ✗ No text provided. Usage: %v <text>\n \n", cmd)
	}
	return strings.ToUpper(s)
}

func toLowerCase(s, cmd string) string {
	if s == "" {
		fmt.Printf(" ✗ No text provided. Usage: %v <text>\n \n", cmd)
	}
	return strings.ToLower(s)
}

func toSnakeCase(s, cmd string) string {
	if s == "" {
		fmt.Printf(" ✗ No text provided. Usage: %v <text>\n \n", cmd)
	}
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

func titleCase(s, cmd string) string {
	if s == "" {
		fmt.Printf(" ✗ No text provided. Usage: %v <text>\n \n", cmd)
	}
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

func reverseWords(s, cmd string) string {
	if s == "" {
		fmt.Printf(" ✗ No text provided. Usage: %v <text>\n \n", cmd)
	}
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

func capitalizeCase(s, cmd string) string {
	if s == "" {
		fmt.Printf(" ✗ No text provided. Usage: %v <text>\n \n", cmd)
	}
	words := strings.Fields(s)

	for i, word := range words {
		if len(word) == 0 {
			continue
		}
		words[i] = strings.ToUpper(string(word[0])) + strings.ToLower(word[1:])
	}
	return strings.Join(words, " ")
}

func palindrome(s, cmd string) string {
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
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("\n SENTINEL STRING TRANSFORMER — ONLINE\n ──────────────────────────────────────\n \n")

	for {
	start:
		fmt.Print(" > Enter command: ")
		inputWord, _ := reader.ReadString('\n')
		inputWord = strings.TrimSpace(inputWord)

		if inputWord == "" {
			goto start
			//continue
		}

		split := strings.Fields(inputWord)

		cmd := strings.ToLower(split[0])

		wordSlice := split[1:]

		words := strings.Join(wordSlice, " ")

		switch cmd {
		case "upper":
			fmt.Printf(" → %v\n \n", toUpperCase(words, cmd))

		case "lower":
			fmt.Printf(" → %v\n \n", toLowerCase(words, cmd))

		case "cap":
			fmt.Printf(" → %v\n \n", capitalizeCase(words, cmd))

		case "title":
			fmt.Printf(" → %v\n \n", titleCase(words, cmd))

		case "snake":
			fmt.Printf(" → %v\n \n", toSnakeCase(words, cmd))

		case "reverse":
			fmt.Printf(" → %v\n \n", reverseWords(words, cmd))

		case "palindrome":
			fmt.Printf(" → %v\n \n", palindrome(words, cmd))

		case "exit":
			fmt.Print(" Shutting down String Transformer. Goodbye.\n \n")
			return

		default:
			fmt.Printf(" ✗ Unknown command: %q\n", cmd)
			fmt.Print("   Valid commands: upper, lower, cap, title, snake, reverse, palindrome, exit\n \n")
		}
	}
}
