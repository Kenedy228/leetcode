package task2264

import "testing"

func TestLargestGoodInteger(t *testing.T) {
	inputs := []string{"6777133339", "2300019", "42352338", "222"}
	expected := []string{"777", "000", "", "222"}

	for i := range inputs {
		output := largestGoodInteger(inputs[i])

		if output != expected[i] {
			t.Errorf("error largestGoodInteger: for value %v expected %v, got %v", inputs[i], expected[i], output)
		}
	}
}
