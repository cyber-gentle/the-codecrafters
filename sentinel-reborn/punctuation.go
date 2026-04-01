package main

import (
	"strings"
)

func fixArticle(word string) string {
 words := strings.Fields(word)
 vowels := "aeiouhAEIOUH"
 for i := 0; i < len(words)-1; i++ {
     if (words[i] == "a" || words[i] == "A") && strings.ContainsRune(vowels, rune(words[i+1][0])) {
         if words[i] == "A" {
             words[i] = "An"
         } else {
             words[i] = "an"
         }
     }
 }
 return strings.Join(words, " ")
}