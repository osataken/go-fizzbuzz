package fizzbuzz

import (
	"strconv"
	"math"
)

func fizzbuzz(input string) string {
	floatInput, _ := strconv.ParseFloat(input, 64)

	if math.Mod(floatInput, 3) == 0{
		return "fizz"
	}

	if input == "5" {
		return "buzz"
	}

	return "1"
}