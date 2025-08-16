package sort

func InsertionSort(nums []int, start, end int) {
	for i := start + 1; i <= end; i++ {
		key := nums[i]
		j := i - 1

		for j >= start && key < nums[j] {
			nums[j+1] = nums[j]
			j--
		}

		nums[j+1] = key
	}
}
