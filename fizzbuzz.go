package fizzbuzz

import (
	"strconv"
	"math"
)

func fizzbuzz(input string) string {
	floatInput, _ := strconv.ParseFloat(input, 64)

	result := input

	if math.Mod(floatInput, 15) == 0 {
		result = "fizzbuzz"
	} else if math.Mod(floatInput, 3) == 0 {
		result = "fizz"
	} else if math.Mod(floatInput, 5) == 0 {
		result = "buzz"
	}

	return result
}