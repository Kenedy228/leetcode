package task342

import "testing"

func TestIsPowerOfFour(t *testing.T) {
	inputs := []int{16, 5, 1}
	expected := []bool{true, false, true}

	for i := range inputs {
		output := isPowerOfFour(inputs[i])
		if output != expected[i] {
			t.Errorf("error isPowerOfFour: for input %v expected %v, got %v", inputs[i], expected[i], output)
		}
	}
}
