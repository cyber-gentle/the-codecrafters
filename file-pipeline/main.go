// CodeCrafters — Operation Gopher Protocol
// Module: File Pipeline
// Author: Abraham David
// Squad:  [Your Squad Name]

package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Print(" Failed to open file\n No file name matches yours\n \n")
	}

	defer file.Close()

	// input_content, err := os.ReadFile("file")
	// if err != nil {
	// 	fmt.Print(" Unable to read file\n \n")
	// }

	data := make([]byte, 200)
	count, err := file.Read(data)
	 if err != nil {
              fmt.Println("Error in reading file")
	}
	fmt.Printf("%q\n", data[:count])
}
