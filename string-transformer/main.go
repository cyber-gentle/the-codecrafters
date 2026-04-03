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

// func capitalizeCase(s string) string {
// 	return strings.ToLower(s)
// }

func main() {

start:
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("\n SENTINEL STRING TRANSFORMER — ONLINE\n ──────────────────────────────────────\n \n")
	fmt.Print(" > Enter command: ")

	input_word, _ := reader.ReadString('\n') // User's input
	input_word = strings.TrimSpace(input_word)
	split := strings.Fields(input_word) // Splitting of User's input into slice of string
	cmd := strings.ToLower(split[0])    // getting command from User's input to lower case

	word := split[1:]                // Get word from  user input in form of slice of strings
	words := strings.Join(word, " ") // Joining word to back from slice to string

	for {
		switch cmd {
		case "upper":
			fmt.Printf(" → %v\n \n", toUpperCase(words))
			goto start

		case "lower":
			fmt.Printf(" → %v\n \n", toLowerCase(words))
			goto start

		case "exit":
			fmt.Print(" Shutting down String Transformer. Goodbye.\n \n")
			return

		default:
			fmt.Printf(" ✗ Unknown command: %q\n Valid commands: upper, lower, cap, title, snake, reverse, exit \n \n", cmd)
			goto start
		}
	} 

}
