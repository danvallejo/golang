package math

import (
	"fmt"
	"testing"
)

func Assert(t *testing.T, expectedValue int, actualValue int, message string) {
	if expectedValue != actualValue {
		fmt.Println("Msg:", message, "Expected:", expectedValue, "Actual:", actualValue)
	}
}

func TestAdd(t *testing.T) {
	sum := Add(1, 2)
	Assert(t, 3, sum, "Invalid sum")

}

func TestSubtract(t *testing.T) {
	temp := Subtract(1, 5)
	Assert(t, -4, temp, "Invalid subtraction")
}
