package main

import (
	"strconv"
	"strings"
)


// DONE
func converter(input string) string {
	words := strings.Fields(input)

	for i := 1; i < len(words); i++ {
		if words[i] == "(hex)" {
			if n, err := strconv.ParseInt(words[i-1], 16, 64); err == nil {
				words[i-1] = strconv.FormatInt(n, 10)
			}
			words = append(words[:i], words[i+1:]...)
			i--

		} else if words[i] == "(bin)" {
			if n, err := strconv.ParseInt(words[i-1], 2, 64); err == nil {
				words[i-1] = strconv.FormatInt(n, 10)
			}
			words = append(words[:i], words[i+1:]...)
			i--
		}
	}

	return strings.Join(words, " ")
}

func lowerCase(text string) string {
	words := strings.Fields(text)
	for i, s := range words {
		if s == "(low," {
			n := int(words[i+1][0] - '0')

			for j := i - n; j < i; j++ {
				if j >= 0 {
					words[j] = strings.ToLower(words[j])
				}
			}
			words = append(words[:i], words[i+2:]...)
			break
		}

	}
	return strings.Join(words, " ")

}

func ToUpperCase(words string) string {
	word := strings.Fields(words)
	for i, c := range word {
		if c == "(up)" {
			word[i-1] = strings.ToUpper(word[i-1])
			word = append(word[:i], word[i+1:]...)
		}
	}
	return strings.Join(word, " ")
}

func applyTransformation(s string) string {
	s = converter(s)
	s = fixArticle(s)
	s = fixPunctuation(s)
	s = ToUpperCase(s)
	s = lowerCase(s)
	

	return s
}
