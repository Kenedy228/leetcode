package task2787

import (
	"slices"
	"testing"
)

func TestFindPowers(t *testing.T) {
	n := 10
	power := 2
	expected := []int{1, 4, 9}
	output := findPowers(n, power)

	if !slices.Equal(output, expected) {
		t.Errorf("findPowers: n = %v, power = %v, expected %v, got %v", n, power, expected, output)
	}
}

func TestNumberOfWays(t *testing.T) {
	n := 10
	power := 2
	output := numberOfWays(n, power)
	expected := 1

	if output != expected {
		t.Errorf("numberOfWays: n = %v, power = %v, expected %v, got %v", n, power, expected, output)
	}

	n = 4
	power = 1
	output = numberOfWays(n, power)
	expected = 2

	if output != expected {
		t.Errorf("numberOfWays: n = %v, power = %v, expected %v, got %v", n, power, expected, output)
	}

	// n = 98
	// power = 1
	// output = numberOfWays(n, power)
	// expected = 376256
	//
	// if output != expected {
	// 	t.Errorf("numberOfWays: n = %v, power = %v, expected %v, got %v", n, power, expected, output)
	// }

}
