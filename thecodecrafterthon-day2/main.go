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
	dec, err := strconv.ParseInt(s, base, 64)
	if err != nil {
		fmt.Println("Invalid Hexadecimal number:", err)
	}
	return dec, err
}

func binToDecimal(n string, base int) (int64, error) {
	return strconv.ParseInt(n, base, 64)
}

func main() {
	reader := bufio.NewReader(os.Stdin)
start:
	fmt.Println("       Welcome!")
	fmt.Print(" Select Base to covert\n 1. Decimal to Binary and HexaDecimal\n 2. HexaDecimal to Decimal\n 3. Binary to Decimal\n 4. Exit\n")
	fmt.Print("Pick any one to continue: ")
	choice, _ := reader.ReadString('\n')
	choice = strings.TrimSpace(choice)
	c_choice, _ := strconv.Atoi(choice) // c_choice stands for cleaned choice after Atoi

	for {
		switch c_choice {
		case 1:
			fmt.Println("Enter a binary number to ")

		case 2:
			fmt.Println()
			fmt.Print("Enter a valid HexaDecimal Number to convert to Decimal: ")
			number, _ := reader.ReadString('\n')
			number = strings.TrimSpace(number)

			base := 16
			fmt.Println(hexToDecimal(number, base))

			fmt.Println()
			fmt.Print("Do you have another HexaDecimal number to convert\n 1. Yes\n 2. No, Quit\n 3. Go to Menu\n")
			choice, _ := reader.ReadString('\n')
			choice = strings.TrimSpace(choice)
			choice_number, _ := strconv.Atoi(choice) // cleaned choice after atoi

			if choice_number == 1 {
				continue
			} else if choice_number == 2 {
				return
			} else if choice_number == 3 {
				goto start
			}
		case 3:
			fmt.Println()
			fmt.Print("Enter a valid Binary Number to convert to Decimal: ")
			number, _ := reader.ReadString('\n')
			number = strings.TrimSpace(number)

			base := 2
			fmt.Println(binToDecimal(number, base))

			fmt.Println()
			fmt.Print("Do you have another Binary number to convert\n 1. Yes\n 2. No, Quit\n 3. Go to Menu\n")
			choice, _ := reader.ReadString('\n')
			choice = strings.TrimSpace(choice)
			choice_number, _ := strconv.Atoi(choice) // cleaned choice after atoi

			if choice_number == 1 {
				continue
			} else if choice_number == 2 {
				return
			} else if choice_number == 3 {
				goto start
			}

		}
	}

	// fmt.Println(hexToDecimal("1E", 16))
	// fmt.Println(binToDecimal("1001", 2))
}
