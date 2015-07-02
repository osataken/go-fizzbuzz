package fizzbuzz

import (
	"testing"
)

func Test_input_1_should_return_1(t *testing.T) {
	expected := "1"

	result := fizzbuzz("1")

	if result != expected {
		t.Errorf("Expected %V but was %V", expected, result)
	}
}

func Test_input_3_should_return_fizz(t *testing.T) {
	expected := "fizz"

	result := fizzbuzz("3")

	if result != expected {
		t.Errorf("Expected %V but was %V", expected, result)
	}
}

func Test_input_5_should_return_buzz(t *testing.T) {
	expected := "buzz"

	result := fizzbuzz("5")

	if result != expected {
		t.Errorf("Expected %V but was %V", expected, result)
	}
}

func Test_input_6_should_return_fizz(t *testing.T) {
	expected := "fizz"

	result :=fizzbuzz("6")

	if result != expected {
		t.Errorf("Expected %V but was %V", expected, result)
	}
}

func Test_input_15_should_return_fizzbuzz(t *testing.T) {
	expected := "fizzbuzz"

	result := fizzbuzz("15")

	if result != expected {
		t.Errorf("Expected %V but was %V", expected, result)
	}
}