package main

import (
	"fmt"
	"strconv"
)

// func decimalToOthers(s string, base int) (float64, error) {
// 	conv, _ := strconv.Atoi(s)
// 	conv = float64(conv)
// 	return strconv.FormatInt(conv, base, 64)
// }

func baseToDecimal(s string, base int) (int64, error) {
	return strconv.ParseInt(s, base, 64)
}

func binToDecimal(s string, base int) (int64, error) {
	return strconv.ParseInt(s, base, 64)
}

func main() {
	fmt.Println("       Welcome!")
	fmt.Println(" Select Base to covert\n 1. Decimal to Binary and HexaDecimal\n 2. HexaDecimal to Decimal\n 3. Binary to Decimal")
	fmt.Println("Pick any one to continue: ")

	// fmt.Println(baseToDecimal("1E", 16))
	// fmt.Println(binToDecimal("1001", 2))
}
