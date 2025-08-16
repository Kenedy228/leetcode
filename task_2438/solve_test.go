package task2438

import (
	"slices"
	"testing"
)

func TestBinaryRepresentation(t *testing.T) {
	inputs := []int{15, 2, 6}
	answers := [][]int{{1, 1, 1, 1}, {1, 0}, {1, 1, 0}}

	for i := range inputs {
		output := getBinaryRepresentation(inputs[i])
		answer := answers[i]

		if !slices.Equal(output, answer) {
			t.Errorf("getBinaryRepresentation: test case failed for input %d: expected %v, got %v", inputs[i], answers[i], output)
		}
	}
}

func TestFindPowers(t *testing.T) {
	inputs := []int{15, 2, 6}
	answers := [][]int{{1, 2, 4, 8}, {2}, {2, 4}}

	for i := range inputs {
		output := findPowers(inputs[i])
		answer := answers[i]

		if !slices.Equal(output, answer) {
			t.Errorf("findPowers: test case failed for input %d: expected %v, got %v", inputs[i], answer, output)
		}
	}
}

func TestProductQueries(t *testing.T) {
	inputs := []int{15, 2}
	queries := [][][]int{{{0, 1}, {2, 2}, {0, 3}}, {{0, 0}}}
	answers := [][]int{{2, 4, 64}, {2}}

	for i := range inputs {
		output := productQueries(inputs[i], queries[i])
		answer := answers[i]

		if !slices.Equal(output, answer) {
			t.Errorf("productQueries: test case failed for input %d: expected %v, got %v", inputs[i], answer, output)
		}
	}
}
