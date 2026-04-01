package main

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

func capitalizeWord(word string) string {
	runes := []rune(strings.ToLower(word))
	if len(runes) > 0 {
		runes[0] = unicode.ToUpper(runes[0])
	}
	return string(runes)
}

func toLowerCase(text string) string {
	words := strings.Fields(text)

	for i := 0; i < len(words); i++ {
		if words[i] == "(low)" {
			if i-1 >= 0 {
				words[i-1] = strings.ToLower(words[i-1])
			}
			words = append(words[:i], words[i+1:]...)
			i--
			continue
		}

		if words[i] == "(low," && i+1 < len(words) {
			countStr := strings.TrimSuffix(words[i+1], ")")
			n, err := strconv.Atoi(countStr)
			if err == nil {
				start := i - n
				if start < 0 {
					start = 0
				}
				for j := start; j < i; j++ {
					words[j] = strings.ToLower(words[j])
				}
			}
			words = append(words[:i], words[i+2:]...)
			i--
		}
	}

	return strings.Join(words, " ")
}

func toUpperCase(text string) string {
	words := strings.Fields(text)

	for i := 0; i < len(words); i++ {
		if words[i] == "(up)" {
			if i-1 >= 0 {
				words[i-1] = strings.ToUpper(words[i-1])
			}
			words = append(words[:i], words[i+1:]...)
			i--
			continue
		}

		if words[i] == "(up," && i+1 < len(words) {
			countStr := strings.TrimSuffix(words[i+1], ")")
			n, err := strconv.Atoi(countStr)
			if err == nil {
				start := i - n
				if start < 0 {
					start = 0
				}
				for j := start; j < i; j++ {
					words[j] = strings.ToUpper(words[j])
				}
			}
			words = append(words[:i], words[i+2:]...)
			i--
		}
	}

	return strings.Join(words, " ")
}

func toTitleCase(text string) string {
	words := strings.Fields(text)

	for i := 0; i < len(words); i++ {
		if words[i] == "(cap)" {
			if i-1 >= 0 {
				words[i-1] = capitalizeWord(words[i-1])
			}
			words = append(words[:i], words[i+1:]...)
			i--
			continue
		}

		if words[i] == "(cap," && i+1 < len(words) {
			countStr := strings.TrimSuffix(words[i+1], ")")
			num, err := strconv.Atoi(countStr)
			if err == nil {
				start := i - num
				if start < 0 {
					start = 0
				}
				for j := start; j < i; j++ {
					words[j] = capitalizeWord(words[j])
				}
			}
			words = append(words[:i], words[i+2:]...)
			i--
		}
	}

	return strings.Join(words, " ")
}

func transforms(text string) string {
	text = toLowerCase(text)
	text = toUpperCase(text)
	text = toTitleCase(text)
	return text
}

func main() {
	sent := "it (cap) was the best of times, it was the worst of times (up) , it was the age of wisdom, it was the age of foolishness (cap, 5) , it was the epoch of belief, it was the epoch of incredulity (up, 2), it was the SEASON OF LIGHT (low, 3), it was the season of darkness, it was the SPRING (low) of hope, IT WAS THE (low, 2) winter of despair."
	fmt.Println(transforms(sent))
}
