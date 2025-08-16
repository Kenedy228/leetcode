package task342

const power int = 4

func isPowerOfFour(n int) bool {
	x := 1

	for x < n {
		x *= power
	}

	return x == n
}
