package main

import (
	"fmt"
	//"os"
	"strings"
	
	"strconv"
)

func hexConverter() {


}

func main() {
	words := "1E (hex) files were added"

	hexCoverter(words)

}

func hexCoverter(s string) {
	words := strings.Split(s, " ")
	for i, word := range (words) {
		if word == "(hex)" {
			n, err := strconv.ParseInt(words[i-1], 16, 64)
			if err != nil {
			continue
			}
			words[i] = words[i-1:i] + words[i+1:]
		}

	}   
}

