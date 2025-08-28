package task53

import "testing"

func TestMaxSubArray(t *testing.T) {
	inputs := [][]int{{-2, 1, -3, 4, -1, 2, 1, -5, 4}, {1}, {5, 4, -1, 7, 8}}
	expected := []int{6, 1, 23}

	for i := range inputs {
		output := maxSubArray(inputs[i])

		if output != expected[i] {
			t.Errorf("failed maxSubArray: for input %v expected %v, got %v", inputs[i], expected[i], output)
		}
	}
}
