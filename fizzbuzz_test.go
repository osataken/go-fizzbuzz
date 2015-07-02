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