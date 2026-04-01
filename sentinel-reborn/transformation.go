package main

import (
	"strconv"
	"strings"
)

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
