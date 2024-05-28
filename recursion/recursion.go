package recursion

func Sum(num []int) int {
	if len(num) == 0 {
		return 0
	}
	return num[0] + Sum(num[1:])
}

func CountLength(items []int) int {
	if len(items) == 0 {
		return 0
	}

	return 1 + CountLength(items[1:])
}

func FindMax(num []int, prevMax int) int {
	if len(num) == 0 {
		return prevMax
	}

	max := prevMax
	first := num[0]
	if first > prevMax {
		max = first
	}

	return FindMax(num[1:], max)
}
