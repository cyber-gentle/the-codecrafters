package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var ()

// func decimalToOthers(s string, base int) (float64, error) {
// 	conv, _ := strconv.Atoi(s)
// 	conv = float64(conv)
// 	return strconv.FormatInt(conv, base, 64)
// }

func hexToDecimal(s string, base int) (int64, error) {
	return strconv.ParseInt(s, base, 64)
}

func binToDecimal(n string, base int) (int64, error) {
	return strconv.ParseInt(n, base, 64)
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("       Welcome!")
	fmt.Println(" Select Base to covert\n 1. Decimal to Binary and HexaDecimal\n 2. HexaDecimal to Decimal\n 3. Binary to Decimal")
	fmt.Println("Pick any one to continue: ")
	choice, _ := reader.ReadString('\n')
	choice = strings.TrimSpace(choice)
	c_choice, _ := strconv.Atoi(choice) // c_choice stands for cleaned choice after Atoi

	for {
		switch c_choice {
		case 1:
			fmt.Println("Enter a binary number to ")

		case 2:
			fmt.Println("Enter an HexaDecimal Number to convert to Decimal")
			number, _ := reader.ReadString('\n')
			number = strings.TrimSpace(number)

			fmt.Println("Enter current base")
			base, _ := reader.ReadString('\n')
			base = strings.TrimSpace(base)
			c_base, _ := strconv.Atoi(base) // c_base stands for cleaned base after Atoi
			fmt.Println(hexToDecimal(number, c_base))
			return

		}
	}

	// fmt.Println(hexToDecimal("1E", 16))
	// fmt.Println(binToDecimal("1001", 2))
}
