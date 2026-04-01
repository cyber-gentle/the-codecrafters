package main

import (
	"fmt"
	"os"
)

func hexConverter() {


}

func main() {

	if len(os.Args) != 3 {
		fmt.Println("Usage: go run . input.txt output.txt")
	}
	inputFile := os.Args[1]
	outputFile := os.Args[2]

	// file, err := os.Open("input.txt")
	// if err != nil {
	// 	fmt.Print(" Failed to open file\n No file name matches yours\n \n")
	// }
	// defer file.Close()
	
	content, err := os.ReadFile(inputFile)
	if err != nil {
		fmt.Print(" Failed to read file\n \n")
	}
	err = os.WriteFile(outputFile, []byte(content), 0644)
	if err != nil {
		fmt.Println(" Failed to write file")
	}

	for _, word := range (content) {
		
	}
}
