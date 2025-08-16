package task2264

import (
	"fmt"
	"slices"
)

func largestGoodInteger(num string) string {
	start, end := 0, 1
	seq := 0
	goods := []string{}

	for end < len(num) {
		if num[start] == num[end] {
			seq++
		} else {
			seq = 0
		}

		if seq == 2 {
			goods = append(goods, string(num[start]))
			seq = 0
		}

		start++
		end++
	}

	if len(goods) > 0 {
		slices.Sort(goods)
		return fmt.Sprintf("%v%v%v", goods[len(goods)-1], goods[len(goods)-1], goods[len(goods)-1])
	}

	return ""
}
