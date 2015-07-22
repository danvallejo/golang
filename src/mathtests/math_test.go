package math

import (
	"testing"
)

func Assert(t *testing.T, expectedValue interface{}, actualValue interface{}, message string) {
	if expectedValue != actualValue {
		t.Errorf("Msg: %s Expected: %v Actual: %v", message, expectedValue, actualValue)
	}
}

func TestAdd(t *testing.T) {
	sum := Add(1, 2)
	t.Logf("TestAdd: %v + %v", 1, 2)
	Assert(t, 31, sum, "Invalid sum")

}

func TestSubtract(t *testing.T) {
	temp := Subtract(1, 5)
	Assert(t, -4, temp, "Invalid subtraction")
}
