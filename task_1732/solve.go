package task1732

func largestAltitude(gain []int) int {
	prefix := 0
	highestScore := 0

	for i := range gain {
		prefix += gain[i]

		if prefix > highestScore {
			highestScore = prefix
		}
	}

	return highestScore
}
