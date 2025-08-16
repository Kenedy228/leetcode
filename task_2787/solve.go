package task2787

import (
	"fmt"
	"math"
)

var modulo int = 1e9 + 7

func findPowers(n, power int) []int {
	powers := []int{}
	num := 1

	for {
		p := int(math.Pow(float64(num), float64(power)))

		if p > n {
			break
		}

		powers = append(powers, p)
		num++
	}

	return powers
}

func numberOfWays(n, power int) int {
	powers := findPowers(n, power)
	memo := map[string]int{}
	return findWay(n, 0, powers, memo) % modulo
}

func findWay(n, index int, powers []int, memo map[string]int) int {
	if index >= len(powers) {
		return 0
	}

	key := fmt.Sprintf("%d,%d", n, powers[index])

	if v, ok := memo[key]; ok {
		return v
	}

	if n < 0 {
		return 0
	}

	if n == 0 {
		return 1
	}

	total := 0

	for i := index; i < len(powers); i++ {
		total += findWay(n-powers[i], i+1, powers, memo)
	}

	memo[key] += total

	return total
}
