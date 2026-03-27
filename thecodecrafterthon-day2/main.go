package main

import (
	"fmt"
	"strconv"
)

func baseToDecimal(s string, base int) (int64, error) {
	return strconv.ParseInt(s, base, 64)
}

func binToDecimal(s string, base int) (int64, error) {
	return strconv.ParseInt(s, base, 64)
}

func main() {
	fmt.Println(baseToDecimal("1E", 16))
	fmt.Println(binToDecimal("1001", 2))
}
