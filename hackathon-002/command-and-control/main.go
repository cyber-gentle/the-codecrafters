// CodeCrafters — Hackathon 002
// Squad: Gophers
// Members: David Abraham, Obeko Eunice, Ugwu Chioma, Michael Bulus, Emmanuel Eliabu, Akatu Worthy, Samuel Jireh

package main

import (
	"fmt"
	// "math"
	"strconv"
	"strings"
)

// func main() {
// 	var first float64
// 	var second float64

// 	fmt.Println("...GOPHER'S CALC...")
// Start1:
// 	fmt.Println("Input first number")
// 	_, err := fmt.Scanln(&first)
// 	if err != nil {
// 		fmt.Print("Enter digit only!\n")
// 		goto Start1
// 	}

// Start2:
// 	fmt.Println("Input second number")
// 	fmt.Scanln(&second)
// 	if err != nil {
// 		fmt.Print("Enter digit only!\n")
// 		goto Start2
// 	}

// 	var Operator int

// 	fmt.Println("1: Addition| 2: Subtraction| 3: Multiplication| 4: Division| 5: Remainder | 6: Power | 7: Last | 8: Exit")
// 	fmt.Scanln(&Operator)

// 	if Operator < 1 || Operator > 10 {
// 		fmt.Println("Not a valid Operator")
// 	}

// 	for {
// 		switch Operator {
// 		case 1:
// 			result := first + second
// 			fmt.Println(result)
// 			continue

// 		case 2:
// 			result := first - second
// 			fmt.Println(result)
// 			continue

// 		case 3:
// 			result := first * second
// 			fmt.Println(result)
// 			continue

// 		case 4:
// 			if second == 0 {
// 				fmt.Println("Not divisible by zero")
// 			}
// 			result := first / second
// 			fmt.Println(result)
// 			continue

// 		case 5:
// 			result := int64(first) % int64(second)
// 			fmt.Println(result)
// 			continue

// 		case 6:
// 			result := math.Pow(first, second)
// 			fmt.Println(result)
// 			continue

// 		case 8:
// 			fmt.Println("Exiting...")
// 			break

// 		default:
// 			fmt.Print("Out of operation range\n", "Choose from the above\n")
// 			continue
// 		}
// 		break
// 	}
// }

func main() {
	var hex, dec, bin string

	var base int
start:
	fmt.Println("1: Dec")
	fmt.Println("2: Hex")
	fmt.Println("3: Bin")
	fmt.Println("4: Exit")
	fmt.Scanln(&base)

	for {
		switch base {
		case 1:
			fmt.Print("ENTER DEC NUMBER: ")
			fmt.Scanln(&dec)

			dec, err := strconv.ParseInt(dec, 10, 64)
			if err != nil {
				fmt.Println("Invalid dec")
			}
			hexN := strconv.FormatInt(dec, 16)
			binN := strconv.FormatInt(dec, 2)
			fmt.Printf("HexaDecimal: %v\n", strings.ToUpper(hexN))
			fmt.Printf("Binary: %v\n", binN)
			continue

		case 2:

			fmt.Print("Enter Number: ")
			fmt.Scanln(&hex)
			dec, err := strconv.ParseInt(hex, 16, 64)
			if err != nil {
				fmt.Println("Invalid Hex")
			}
			fmt.Println(dec)
			continue

		case 3:

			fmt.Print("Enter Number: ")
			fmt.Scanln(&bin)

			dec, err := strconv.ParseInt(bin, 2, 64)
			if err != nil {
				fmt.Println("Invalid Bin")
			}
			fmt.Println(dec)
			continue

		case 4:
			fmt.Println("Exiting...")
			return

		default:
			fmt.Println("Invalid Option")
			fmt.Println("Choose A Valid Number")
			goto start
		}

	}
}
