// CodeCrafters — Operation Gopher Protocol
// Module: File Pipeline
// Author: Abraham David
// Squad: The Gophers

// ═══════════════════════════════════════════
// SQUAD PIPELINE CONTRACT
// Squad: Gophers
// ───────────────────────────────────────────
// Input line types:
// Number of lines: 20
// Normal report lines
// Lines in ALL CAPS
// Lines in all lowercase
// Lines starting with TODO:
// Lines with extra leading/trailing spaces

// Transformation rules (in order):
// 1. Trim all leading and trailing whitespace
// 2. Replace TODO: with ✦ ACTION:
// 3. Convert ALL CAPS lines to Title Case
// 4. Convert all lowercase lines to uppercase
// 5. Reverse the words in any line that contains the word REVERSE

// Output format:
// Header: Yes, Exact Text: "Gopher's Sentinel Field Report - Processed"
// Line numbering format : "1."
// Summary block: yes
//  	Fields :
//		✦ Lines read    :
//		✦ Lines written :
//		✦ Lines removed :
//		✦ Rules applied : [our 5 rules]
//
//
// Terminal summary fields:
//		✦ Lines read    :
//		✦ Lines written :
//		✦ Lines removed :
//		✦ Rules applied : [our 5 rules]
// ═══════════════════════════════════════════

// CodeCrafters — Operation Gopher Protocol
// Module: File Pipeline
// Author: Abraham David
// Squad: The Gophers

// ═══════════════════════════════════════════
// SQUAD PIPELINE CONTRACT
// Squad: Gophers
// ───────────────────────────────────────────
// Input line types:
// Number of lines: 20
// Normal report lines
// Lines in ALL CAPS
// Lines in all lowercase
// Lines starting with TODO:
// Lines with extra leading/trailing spaces
//
// Transformation rules (in order):
// 1. Trim all leading and trailing whitespace
// 2. Replace TODO: with ✦ ACTION:
// 3. Convert ALL CAPS lines to Title Case
// 4. Convert all lowercase lines to uppercase
// 5. Reverse the words in any line that contains the word REVERSE
//
// Output format:
// Header: Yes, Exact Text: "Gopher's Sentinel Field Report - Processed"
// Line numbering format : "1."
// Summary block: yes
//  	Fields :
//		✦ Lines read    :
//		✦ Lines written :
//		✦ Lines removed :
//		✦ Rules applied : [our 5 rules]
//
// Terminal summary fields:
//		✦ Lines read    :
//		✦ Lines written :
//		✦ Lines removed :
//		✦ Rules applied : [our 5 rules]
// ═══════════════════════════════════════════

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func trimWhitespace(line string) string {
	return strings.TrimSpace(line)
}

func replaceTodo(line string) string {
	if strings.HasPrefix(strings.ToUpper(line), "TODO:") {
		rest := line[5:]
		line = "✦ ACTION:" + rest
	}
	return line
}

func convertAllCapsToTitleCase(line string) string {
	if isAllCaps(line) {
		return toTitleCase(line)
	}
	return line
}

func isAllCaps(line string) bool {
	hasLetter := false

	for _, ch := range line {
		if unicode.IsLetter(ch) {
			hasLetter = true
			if unicode.IsLower(ch) {
				return false
			}
		}
	}
	return hasLetter
}

func toTitleCase(line string) string {
	words := strings.Fields(line)

	for i, word := range words {
		if len(word) == 0 {
			continue
		}
		words[i] = strings.ToUpper(string(word[0])) + strings.ToLower(word[1:])
	}

	return strings.Join(words, " ")
}

func convertAllLowerToUpper(line string) string {
	if isAllLower(line) {
		return strings.ToUpper(line)
	}
	return line
}

func isAllLower(line string) bool {
	hasLetter := false

	for _, ch := range line {
		if unicode.IsLetter(ch) {
			hasLetter = true
			if unicode.IsUpper(ch) {
				return false
			}
		}
	}

	return hasLetter
}

func reverseWords(line string) string {
	if strings.Contains(line, "REVERSE") {
		words := strings.Fields(line)

		left := 0
		right := len(words) - 1
		for left < right {
			words[left], words[right] = words[right], words[left]
			left++
			right--
		}

		return strings.Join(words, " ")
	}
	return line
}

func applyAllRules(line string) string {
	line = trimWhitespace(line)            // Rule 1
	line = replaceTodo(line)               // Rule 2
	line = convertAllCapsToTitleCase(line) // Rule 3
	line = convertAllLowerToUpper(line)    // Rule 4
	line = reverseWords(line)              // Rule 5
	return line
}

func main() {

	if len(os.Args) < 3 {
		fmt.Println("Usage: go run . <input.txt> <output.txt>")
		os.Exit(1)
	}

	inputFile := os.Args[1]
	outputFile := os.Args[2]

	if inputFile == outputFile {
		fmt.Println("✗ Input and output cannot be the same file.")
		os.Exit(1)
	}

	info, err := os.Stat(outputFile)
	if err == nil && info.IsDir() {
		fmt.Println("✗ Cannot write to output: path is a directory, not a file.")
		os.Exit(1)
	}

	file, err := os.Open(inputFile)
	if err != nil {
		fmt.Println("✗ File not found:", inputFile)
		os.Exit(1)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file) // bufio.Scanner is better than os.ReadFile because it reads one line at a time — it doesn't load the whole file into memory.

	var processedLines []string

	linesRead := 0
	linesRemoved := 0

	for scanner.Scan() {
		line := scanner.Text()
		linesRead++

		transformed := applyAllRules(line)

		if transformed == "" {
			linesRemoved++
			continue
		}
		processedLines = append(processedLines, transformed)
	}

	if linesRead == 0 {
		fmt.Println("⚠ Input file is empty. Nothing to process.")

		emptyOutput, err := os.Create(outputFile)
		if err != nil {
			fmt.Println("✗ Could not write output file:", err)
			os.Exit(1)
		}
		emptyOutput.Close()

		fmt.Println("✦ Lines read    :", 0)
		fmt.Println("✦ Lines written :", 0)
		fmt.Println("✦ Lines removed :", 0)
		fmt.Println("✦ Rules applied : Trim whitespace, Replace TODO, ALL CAPS → Title Case, all lower → UPPER, Reverse words")
		return
	}

	output, err := os.Create(outputFile)
	if err != nil {
		fmt.Println("✗ Could not write output file:", err)
		os.Exit(1)
	}
	defer output.Close()

	writer := bufio.NewWriter(output)

	fmt.Fprintln(writer, "Gopher's Sentinel Field Report - Processed")
	fmt.Fprintln(writer, "───────────────────────────────────────────")

	for i, line := range processedLines {
		fmt.Fprintf(writer, "%d. %s\n", i+1, line)
	}

	fmt.Fprintln(writer, "───────────────────────────────────────────")
	fmt.Fprintln(writer, "SUMMARY")
	fmt.Fprintf(writer, "✦ Lines read    : %d\n", linesRead)
	fmt.Fprintf(writer, "✦ Lines written : %d\n", len(processedLines))
	fmt.Fprintf(writer, "✦ Lines removed : %d\n", linesRemoved)
	fmt.Fprintln(writer, "✦ Rules applied : Trim whitespace | Replace TODO | ALL CAPS → Title Case | all lower → UPPER | Reverse words")

	writer.Flush()

	fmt.Println()
	fmt.Println("✦ Lines read    :", linesRead)
	fmt.Println("✦ Lines written :", len(processedLines))
	fmt.Println("✦ Lines removed :", linesRemoved)
	fmt.Println("✦ Rules applied : Trim whitespace | Replace TODO | ALL CAPS → Title Case | all lower → UPPER | Reverse words")
}
