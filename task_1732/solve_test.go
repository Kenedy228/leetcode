package task1732

import (
	"testing"
)

func TestLargestAltitude(t *testing.T) {
	inputs := [][]int{{-5, 1, 5, 0, -7}, {-4, -3, -2, -1, 4, 3, 2}}
	expected := []int{1, 0}

	for i := range inputs {
		output := largestAltitude(inputs[i])

		if output != expected[i] {
			t.Errorf("largestAltitude fail: for input %v expected %v, got %v", inputs[i], expected[i], output)
		}
	}
}
