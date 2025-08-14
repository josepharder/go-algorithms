package sort

func BubbleSort(nums []int) []int {
	l := len(nums)
	for i := 0; i < l-1; i++ {
		swapped := false

		for j := 0; j < l-i-1; j++ {
			if nums[j] > nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
				swapped = true
			}
		}

		if !swapped {
			break
		}
	}

	return nums
}
