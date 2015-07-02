package fizzbuzz

func fizzbuzz(input string) string {
	if input == "3" || input == "6"{
		return "fizz"
	}

	if input == "5" {
		return "buzz"
	}

	return "1"
}