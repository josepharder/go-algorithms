package sort

func QuickSort(arr []int, start, end int) []int {
	if len(arr) == 3 {
		return InsertionSort(arr)
	} else {
		pivot := partision(arr, start, end)
		arr = QuickSort(arr, start, pivot-1)
		arr = QuickSort(arr, pivot+1, end)
	}

	return arr
}

func partision(arr []int, start, end int) int {

	middle := getMidPoint(arr, start, end)
	sortFirstMidLast(arr, start, middle, end)

	arr[len(arr)-2], arr[middle] = arr[middle], arr[len(arr)-2] // swap middle to a safe spot position

	pivotIndex := len(arr) - 2
	pivotValue := arr[len(arr)-2]

	fromLeft := start + 1
	fromRight := end - 2

	done := false
	for !done {
		for arr[fromLeft] < pivotValue {
			fromLeft++
		}

		for arr[fromRight] > pivotValue {
			fromRight--
		}

		if fromLeft < fromRight {
			arr[fromLeft], arr[fromRight] = arr[fromRight], arr[fromLeft]
			fromLeft++
			fromRight--
		} else {
			done = true
		}
	}

	arr[fromLeft], arr[pivotIndex] = arr[pivotIndex], arr[fromLeft] // return pivot to its middle position

	return fromLeft

}

func getMidPoint(arr []int, start, end int) int {
	return len(arr[start:end]) / 2
}

func sortFirstMidLast(arr []int, start, middle, end int) {
	smallest, mid, largest := arr[start], arr[middle], arr[end]

	if smallest > mid {
		mid, smallest = smallest, mid
	} else {
		smallest, mid = mid, smallest
	}

	if mid > largest {
		mid, largest = largest, mid
	} else {
		largest, mid = mid, largest
	}

	arr[start], arr[middle], arr[end] = smallest, mid, largest
}
