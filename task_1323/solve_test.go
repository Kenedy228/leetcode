package task1323

import (
	"slices"
	"testing"
)

func TestConvertToArray(t *testing.T) {
	inputs := []int{9669, 9996, 9999}
	expected := [][]int{{9, 6, 6, 9}, {9, 9, 9, 6}, {9, 9, 9, 9}}

	for i := range inputs {
		output := convertToArray(inputs[i])
		if !slices.Equal(output, expected[i]) {
			t.Errorf("error convertToArray: for input %v expected %v, got %v", inputs[i], expected[i], output)
		}
	}
}

func TestConvertToNum(t *testing.T) {
	inputs := [][]int{{9, 6, 6, 9}, {9, 9, 9, 6}, {9, 9, 9, 9}}
	expected := []int{9669, 9996, 9999}

	for i := range inputs {
		output := convertToNum(inputs[i])
		if output != expected[i] {
			t.Errorf("error convertToNum: for input %v expeceted %v, got %v", inputs[i], expected[i], output)
		}
	}
}

func TestMaximum69Number(t *testing.T) {
	inputs := []int{9669, 9996, 9999}
	expected := []int{9969, 9999, 9999}

	for i := range inputs {
		output := maximum69Number(inputs[i])
		if output != expected[i] {
			t.Errorf("error maximum69Number: for input %v expected %v, got %v", inputs[i], expected[i], output)
		}
	}
}
