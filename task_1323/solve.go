package task1323

import (
	"math"
	"slices"
)

func maximum69Number(num int) int {
	digits := convertToArray(num)

	for i := range digits {
		if digits[i] == 6 {
			digits[i] = 9
			break
		}
	}

	return convertToNum(digits)
}

func convertToArray(num int) []int {
	result := make([]int, 0)

	for num > 0 {
		result = append(result, num%10)
		num /= 10
	}

	slices.Reverse(result)

	return result
}

func convertToNum(digits []int) int {
	num := 0
	factor := int(math.Pow(10.0, float64(len(digits)-1)))

	for _, v := range digits {
		num += factor * v
		if factor == 10 {
			factor = 1
		} else {
			factor /= 10
		}
	}

	return num
}
