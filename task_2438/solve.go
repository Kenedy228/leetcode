package task2438

import (
	"math"
	"slices"
)

const modulo int = 1e9 + 7

func getBinaryRepresentation(n int) []int {
	binaries := []int{}

	for n > 0 {
		if n%2 == 0 {
			binaries = append(binaries, 0)
		} else {
			binaries = append(binaries, 1)
		}
		n /= 2
	}

	slices.Reverse(binaries)
	return binaries
}

func findPowers(n int) []int {
	binaries := getBinaryRepresentation(n)

	powers := []int{}

	exponent := len(binaries) - 1

	for i := range binaries {
		if binaries[i] == 1 {
			power := int(math.Pow(float64(2), float64(exponent)))
			powers = append(powers, power)
		}

		exponent--
	}

	slices.Reverse(powers)

	return powers
}

func findMult(powers []int, start, end int) int {
	mult := 1

	for i := start; i <= end; i++ {
		mult = (mult * powers[i]) % modulo
	}

	return mult % modulo
}

func productQueries(n int, queries [][]int) []int {
	powers := findPowers(n)
	result := []int{}

	for _, query := range queries {
		result = append(result, findMult(powers, query[0], query[1]))
	}

	return result
}
